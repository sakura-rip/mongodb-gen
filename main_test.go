package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	os.Args = []string{"sample"}
	if len(os.Args) < 0 {
		log.Fatal("target dir required")
	}
	targetDirName = os.Args[len(os.Args)-1]

	cols := getAllCollections(targetDirName)
	fmt.Printf("%#v", cols)
}
