package lsm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"reflect"
	"sort"
)

/*
We will use utf-8 encoding to convert string to byte array and LittleEndian represent int in bytes

Disk format:
First 12 bytes: keySize - valueSize - numberOfKeys
Next numberOfKeys * (keySize + 4 bytes) bytes: a list of tuple (key, keyIndex)
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

func (table SSTable) dataFile() string {
	return fmt.Sprintf("./db/%s.sstable", table.id.String())
}

func (table SSTable) makeStorage() (*os.File, error) {
	_ = os.Mkdir("./db", 0755)
	file, err := os.OpenFile(table.dataFile(), os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	file.Truncate(0)
	file.Sync()
	return file, nil
}

func (table SSTable) validateInput(key string, value string, doValidateValue bool) error {
	if len(key) == 0 {
		return errors.New("Key can not be empty")
	}
	keyInBytes := []byte(key)
	if len(keyInBytes) > table.keySize {
		return errors.New("Key is bigger than expected")
	}
	if doValidateValue {
		if len(value) == 0 {
			return errors.New("Value can not be empty")
		}
		valueInBytes := []byte(value)
		if len(valueInBytes) > table.valueSize {
			return errors.New("Value is bigger than expected")
		}
	}
	return nil
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

func writeInt(file *os.File, num int) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(num))
	file.Write(bytes)
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
	for _, key := range reflect.ValueOf(table.data).MapKeys() {
		keys = append(keys, key.Interface().(string))
	}
	sort.Strings(keys)
	// append number of keys, so we can simply skip the index later
	writeInt(file, len(keys))
	for index, key := range keys {
		keyInBytes := []byte(key)
		// padding to byte array
		padding := make([]byte, table.keySize-len(keyInBytes))
		keyInBytes = append(keyInBytes, padding...)
		file.Write(keyInBytes)
		binary.Write(file, binary.LittleEndian, index)
	}
	// Second: Build & flush the data
	for _, key := range keys {
		keyInBytes := []byte(key)
		var valueInBytes []byte
		switch value := table.data[key]; value {
		case nil:
			valueInBytes = make([]byte, table.valueSize)
		default:
			valueInBytes = []byte(value.(string))
			padding := make([]byte, table.valueSize-len(valueInBytes))
			valueInBytes = append(valueInBytes, padding...)
		}
		file.Write(keyInBytes)
		file.Write(valueInBytes)
	}
	return file.Close()
}
