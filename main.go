package main

import (
	"log"
	"os"
)

var (
	GenDirName string
)

func main() {
	if len(os.Args) < 0 {
		log.Fatal("target dir required")
	}
	GenDirName = os.Args[len(os.Args)-1]

}
