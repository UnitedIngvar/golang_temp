package reader

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex01/encoding"
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex01/models"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type DBReader interface {
	ReadRecipes() (*models.Recipes, error)
	GetFormat() string
}

func GetReader(filename string) (DBReader, error) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("file %v does not exist", filename)
	}

	fileExtension := filepath.Ext(filename)
	if fileExtension != encoding.JsonFormat && fileExtension != encoding.XmlFormat {
		return nil, fmt.Errorf("file has extension is %v, should be '.json' or '.xml'", fileExtension)
	}

	if fileExtension == encoding.JsonFormat {
		return NewJsonReader(filename), nil
	} else {
		return NewXmlReader(filename), nil
	}
}
