package disk

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
)

var ErrInvalidArg = errors.New("read output arg must be a pointer")

type Store struct {
	filePath string
}

func New(filePath string) *Store {
	return &Store{
		filePath: filePath,
	}
}

func (c *Store) FileExists() bool {
	if _, err := os.Stat(c.filePath); err != nil {
		return false
	}

	return true
}

func (c *Store) Read(v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return ErrInvalidArg
	}

	f, err := os.Open(c.filePath)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(v)
	if err != nil {
		return fmt.Errorf("reading from the file: %w", err)
	}

	return nil
}

func (c *Store) Write(v interface{}) error {
	f, err := os.Create(c.filePath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(v)
	if err != nil {
		return fmt.Errorf("writing to the file: %w", err)
	}

	return nil
}
