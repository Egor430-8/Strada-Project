package storage

import (
	"archive/zip"
	"io"
	"os"

	"github.com/Egor430-8/project/validation"
)

type ZipStorage struct {
	*Storage
}

func NewZipStorage(filename string) *ZipStorage {
	return new(ZipStorage{
		new(Storage{
			filename: filename,
		}),
	})
}

func (z *ZipStorage) Save(data []byte) error {
	file, err := os.Create(z.GetFilename())
	if err != nil {
		return err
	}
	defer file.Close()
	archive := zip.NewWriter(file)
	defer archive.Close()
	w, err := archive.Create("dataFile")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (z *ZipStorage) Load() ([]byte, error) {
	r, err := zip.OpenReader(z.GetFilename())
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if len(r.File) == 0 {
		return nil, validation.EmptyArchiveError
	}
	file := r.File[0]
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}
