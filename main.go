package main

import (
	"fmt"
	"log"
	"os"
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
	targetDirName = os.Args[len(os.Args)-1]

	genTargetDir = targetDirName + "_dao"

	cols := getAllCollections(targetDirName)
	fmt.Printf("%#v", cols)
}

func getVersion() string {
	return Version
}
