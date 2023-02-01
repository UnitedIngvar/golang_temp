package marshaller

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/models"
	"encoding/xml"
)

type xmlMarshaller struct {
}

func NewXmlMarshaller() *xmlMarshaller {
	return &xmlMarshaller{}
}

func (marshaller *xmlMarshaller) Marshall(recipes *models.Recipes) (string, error) {
	xmlBytes, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return "", nil
	}

	return string(xmlBytes), nil
}
