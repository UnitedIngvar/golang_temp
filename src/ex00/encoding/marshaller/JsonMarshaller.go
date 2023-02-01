package marshaller

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/models"
	"encoding/json"
)

type jsonMarshaller struct {
}

func NewJsonMarshaller() *jsonMarshaller {
	return &jsonMarshaller{}
}

func (marshaller *jsonMarshaller) Marshall(recipes *models.Recipes) (string, error) {
	jsonBytes, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return "", nil
	}

	return string(jsonBytes), nil
}
