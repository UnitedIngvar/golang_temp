package reader

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/encoding"
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/models"
	"encoding/json"
	"os"
)

// TODO: make private may be?
type jsonDBReader struct {
	filename string
}

func NewJsonReader(filename string) *jsonDBReader {
	return &jsonDBReader{filename: filename}
}

func (reader *jsonDBReader) ReadRecipes() (*models.Recipes, error) {
	file, err := os.OpenFile(reader.filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	recipes := &models.Recipes{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (reader *jsonDBReader) GetFormat() string {
	return encoding.JsonFormat
}
