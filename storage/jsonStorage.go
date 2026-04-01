package storage

import (
	"os"

	"github.com/Egor430-8/project/validation"
)

type JsonStorage struct {
	*Storage
}

func NewJsonStorage(filename string) *JsonStorage {
	return new(JsonStorage{
		new(Storage{
			filename: filename,
		}),
	})
}

func (s *JsonStorage) Save(data []byte) error {
	err := os.WriteFile(s.GetFilename(), data, 0644)
	if err != nil {
		return validation.DataSavingError
	}
	return nil
}

func (s *JsonStorage) Load() ([]byte, error) {
	data, err := os.ReadFile(s.GetFilename())
	if err != nil {
		return nil, validation.DataUploadError
	}
	return data, nil
}
