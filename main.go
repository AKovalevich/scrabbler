package main

import (
	"os"

	"github.com/AKovalevich/scrabbler/cmd/scrabbler"
)

// Scrawl
func main() {
	os.Exit(scrabbler.Run(os.Args[1:]))
}
