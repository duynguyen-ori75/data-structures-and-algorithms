package lsm

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"sort"
)

/*
We will use utf-8 encoding to convert string to byte array and LittleEndian represent int in bytes

Disk format:
First 12 bytes (three integers - 4 bytes each): keySize - valueSize - numberOfKeys
Next numberOfKeys * (keySize + 4 bytes) bytes: a list of tuple (key, tupleOffset)
Next numberOfKeys * (keySize + valueSize) bytes: a list of tuple (key, value)
*/

type SSTable struct {
	id uuid.UUID
	// maximum tuple size in bytes
	keySize   int
	valueSize int
	// ordered map
	data map[string]interface{}
}

func NewTable(keySize int, valueSize int) (*SSTable, error) {
	result := SSTable{id: uuid.New(), keySize: keySize, valueSize: valueSize, data: make(map[string]interface{})}
	_, err := result.makeStorage()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewTableFromID(id string) (*SSTable, error) {
	tableId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	file, err := openDbFile(dataFile(tableId))
	if err != nil {
		return nil, err
	}
	result := SSTable{id: tableId,
		keySize:   readInt(file, -1),
		valueSize: readInt(file, -1),
		data:      make(map[string]interface{})}
	numberOfKeys := readInt(file, -1)
	keyInByte := make([]byte, result.keySize)
	for idx := 0; idx < numberOfKeys; idx++ {
		file.Read(keyInByte)
		key, offset := string(bytes.Trim(keyInByte, "\x00")), readInt(file, -1)
		dataKeyInByte, dataValueInByte := make([]byte, result.keySize), make([]byte, result.valueSize)
		// ReadAt does not affect file's current cursor
		file.ReadAt(dataKeyInByte, int64(offset))
		file.ReadAt(dataValueInByte, int64(offset+result.keySize))
		dataKey, dataValue := string(bytes.Trim(dataKeyInByte, "\x00")), string(bytes.Trim(dataValueInByte, "\x00"))
		if key != dataKey {
			return nil, fmt.Errorf("The index is wrong. Key in index is %s point to tuple (%s, %s)", key, dataKey, dataValue)
		}
		result.data[dataKey] = dataValue
	}
	return &result, nil
}

func (table *SSTable) Update(key string, value string) error {
	err := table.validateInput(key, value, true)
	if err != nil {
		return err
	}
	table.data[key] = value
	return nil
}

func (table *SSTable) Delete(key string) error {
	err := table.validateInput(key, "", false)
	if err != nil {
		return err
	}
	table.data[key] = nil
	return nil
}

func (table SSTable) Read(key string) (string, error) {
	err := table.validateInput(key, "", false)
	if err != nil {
		return "", err
	}
	if val, ok := table.data[key]; ok {
		if val == nil {
			return "", fmt.Errorf("Key %s not found", key)
		}
		return val.(string), nil
	}
	return "", fmt.Errorf("Key %s not found", key)
}

func (table SSTable) Flush() error {
	file, err := table.makeStorage()
	if err != nil {
		return err
	}
	writeInt(file, table.keySize)
	writeInt(file, table.valueSize)
	// Flush all keys and values
	// First: Build & flush the index
	keys := make([]string, len(table.data))
	for idx, key := range reflect.ValueOf(table.data).MapKeys() {
		keys[idx] = key.Interface().(string)
	}
	sort.Strings(keys)
	// append number of keys, so we can simply skip the index later
	writeInt(file, len(keys))
	// for simplicity, I won't implement any good data structure for the index
	// I will just use a list of key
	for index, key := range keys {
		keyInBytes := []byte(key)
		// padding to byte array
		padding := make([]byte, table.keySize-len(keyInBytes))
		keyInBytes = append(keyInBytes, padding...)
		file.Write(keyInBytes)
		// current byte position of the expected tuple
		// position = 12 (ignore first 12 bytes - which a three metadata)
		// 			+ len(keys) * (table.keySize + 4) (the size of index block)
		//			+ index * (table.keySize + table.valueSize) (the position of the tuple in data block)
		offset := 12 + len(keys)*(table.keySize+4) + index*(table.keySize+table.valueSize)
		writeInt(file, offset)
	}
	// Second: Build & flush the data
	for _, key := range keys {
		// padding to byte array
		keyInBytes := []byte(key)
		padding := make([]byte, table.keySize-len(keyInBytes))
		keyInBytes = append(keyInBytes, padding...)
		var valueInBytes []byte
		switch value := table.data[key]; value {
		case nil:
			valueInBytes = make([]byte, table.valueSize)
		default:
			valueInBytes = []byte(value.(string))
			padding = make([]byte, table.valueSize-len(valueInBytes))
			valueInBytes = append(valueInBytes, padding...)
		}
		file.Write(keyInBytes)
		file.Write(valueInBytes)
	}
	return file.Close()
}
