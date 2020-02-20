package lsm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
)

func writeInt(file *os.File, num int) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(num))
	file.Write(bytes)
}

func readInt(file *os.File, offset int64) int {
	bytes := make([]byte, 4)
	switch offset {
	case -1:
		file.Read(bytes)
	default:
		file.ReadAt(bytes, offset)
	}
	return int(binary.LittleEndian.Uint32(bytes))
}

func dataFile(id uuid.UUID) string {
	return fmt.Sprintf("./db/%s.sstable", id.String())
}

func openDbFile(dbName string) (*os.File, error) {
	_ = os.Mkdir("./db", 0755)
	return os.OpenFile(dbName, os.O_RDWR|os.O_CREATE, 0660)
}

func (table SSTable) openStorage() (*os.File, error) {
	return openDbFile(dataFile(table.id))
}

func (table SSTable) makeStorage() (*os.File, error) {
	file, err := table.openStorage()
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
