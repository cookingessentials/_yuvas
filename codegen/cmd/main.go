package main

import (
	"log"

	"github.com/lkisby/codegen"
)

func main() {
	codeGenerator := codegen.New()
	if err := codeGenerator.Recipes(); err != nil {
		log.Fatal(err)
	}
}
