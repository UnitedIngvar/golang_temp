package main

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/encoding/marshaller"
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex00/encoding/reader"
	"fmt"
	"log"
	"os"
)

func main() {
	filename, err := getFilename()
	if err != nil {
		log.Print(err)
		return
	}

	err = readRecipesAndPrintInOppositeFormat(err, filename)
	if err != nil {
		log.Print(err)
		return
	}
}

func readRecipesAndPrintInOppositeFormat(err error, filename string) error {
	dbReader, err := reader.GetReader(filename)
	if err != nil {
		return err
	}

	recipes, err := dbReader.ReadRecipes()
	if err != nil {
		return err
	}

	marshaller := marshaller.GetMarshallerOfOppositeFormat(dbReader.GetFormat())
	encodedMessage, err := marshaller.Marshall(recipes)
	if err != nil {
		return err
	}

	fmt.Println(encodedMessage)
	return nil
}

func getFilename() (string, error) {
	args := os.Args[1:]

	if len(args) != 2 {
		return "", fmt.Errorf("wrong number of arguments (should be 2)")
	}

	if args[0] != "-f" {
		return "", fmt.Errorf("invalid input format: ./program -f [filename]")
	}

	return args[1], nil
}
