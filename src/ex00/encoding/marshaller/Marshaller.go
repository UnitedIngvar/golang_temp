package marshaller

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/encoding"
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/models"
)

type Marshaller interface {
	Marshall(recipes *models.Recipes) (string, error)
}

func GetMarshallerOfOppositeFormat(formatExtension string) Marshaller {
	if formatExtension == encoding.JsonFormat {
		return NewXmlMarshaller()
	} else {
		return NewJsonMarshaller()
	}
}
