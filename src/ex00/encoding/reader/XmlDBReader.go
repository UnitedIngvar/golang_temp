package reader

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/encoding"
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/models"
	"encoding/xml"
	"os"
)

type xmlDBReader struct {
	filename string
}

func NewXmlReader(filename string) *xmlDBReader {
	return &xmlDBReader{filename: filename}
}

func (reader *xmlDBReader) ReadRecipes() (*models.Recipes, error) {
	file, err := os.OpenFile(reader.filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	recipes := &models.Recipes{}
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (reader *xmlDBReader) GetFormat() string {
	return encoding.XmlFormat
}
