package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	targetDirName string
	genTargetDir  string
	Version       = "develop@v1"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("target dir required")
	}
	targetDirName, err := filepath.Abs(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal("error occurred during parsing directory: ", err)
	}
	fInfo, err := os.Stat(targetDirName)
	if err != nil {
		log.Fatal("error occurred during check path: ", err)
	}
	if !fInfo.IsDir() {
		log.Fatal("error: ", targetDirName, "is not directory")
	}
	fmt.Println(targetDirName)
	genTargetDir = targetDirName + "_dao"

	cols := getAllCollections(targetDirName)

	err = generateClientFile(cols)
	if err != nil {
		log.Fatal(err)
	}

	for _, col := range cols {
		err := generateCollectionBaseFile(col)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = generateQueryFile()
	if err != nil {
		log.Fatal(err)
	}
	for _, col := range cols {
		for fName, field := range col.Fields {
			if fName == col.IdFieldName {
				continue
			}
			var err error
			switch {
			//array
			case strings.HasPrefix(field.FieldType, "[]"):
				err = generateFieldArrayTypeFile(field)
			//map
			case strings.HasPrefix(field.FieldType, "map["):
				err = generateFieldMapTypeFile(field)
			//default(int, str...etc
			case isKnownType(field.FieldType):
				err = generateFieldKnownTypeFile(field)
			//struct
			default:
				err = generateFieldStructTypeFile(field)
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func getVersion() string {
	return Version
}
