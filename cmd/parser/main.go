package main

import (
	"blockchain-parser/internal/app"
	"log"
	"os"
)

func main() {
	if err := app.RunParser(); err != nil {
		log.Printf("Fatal Error %v\n", err)
		os.Exit(1)
	}
}
