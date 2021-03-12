package main

import (
	"fmt"
	"log"
	"os"
)

var (
	targetDirName string
)

func main() {
	if len(os.Args) < 0 {
		log.Fatal("target dir required")
	}
	targetDirName = os.Args[len(os.Args)-1]

	cols := getAllCollections(targetDirName)
	fmt.Printf("%#v", cols)
}
