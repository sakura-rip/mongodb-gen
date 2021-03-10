package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 0 {
		log.Fatal("target dir required")
	}

}
