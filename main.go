package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	targetDirName string
	genTargetDir  string
	Version       = "develop@v1"
)

func main() {
	if len(os.Args) < 0 {
		log.Fatal("target dir required")
	}
	targetDirName, err := filepath.Abs(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal("error occurred during parsing directory: ", err)
	}
	fInfo, _ := os.Stat(targetDirName)
	if !fInfo.IsDir() {
		log.Fatal("error: ", targetDirName, "is not directory")
	}
	genTargetDir = targetDirName + "_dao"

	cols := getAllCollections()
	fmt.Printf("%#v", cols)
}

func getVersion() string {
	return Version
}
