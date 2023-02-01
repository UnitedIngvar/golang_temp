package main

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex01/encoding/reader"
	"flag"
	"log"
	"os"
)

const newFlag = "new"
const oldFlag = "old"

func main() {
	args := os.Args[1:]
	if len(args) != 4 {
		log.Fatalln("wrong number of arguments (should be 4)")
		return
	}
	oldFilename := getOldFilename()
	newFileName := getNewFilename()
	flag.Parse()
	if !isFlagPassed(newFlag) || !isFlagPassed(oldFlag) {
		log.Fatalln("incorrect input format. usage: --new [filename] --old [filename]")
		return
	}

	err := compareFiles(oldFilename, newFileName)
	if err != nil {
		log.Fatalln(err)
	}
}

func compareFiles(oldFilename *string, newFileName *string) error {
	oldDatabaseReader, err := reader.GetReader(*oldFilename)
	if err != nil {
		return err
	}
	newDatabaseReader, err := reader.GetReader(*newFileName)
	if err != nil {
		return err
	}

	oldRecipes, err := oldDatabaseReader.ReadRecipes()
	if err != nil {
		return err
	}
	newRecipes, err := newDatabaseReader.ReadRecipes()
	if err != nil {
		return err
	}

	err = CompareRecipes(oldRecipes, newRecipes)
	if err != nil {
		return err
	}

	return nil
}

func getOldFilename() *string {
	oldFilename := flag.String(oldFlag, "", "old database filename")

	return oldFilename
}

func getNewFilename() *string {
	newFilename := flag.String(newFlag, "", "new database filename")

	return newFilename
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
