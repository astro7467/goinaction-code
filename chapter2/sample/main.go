package main

import (
	"log"
	// "m/v2/chapter2/sample/search"
	"os"

	_ "chapter2/sample/m/v2/matchers"
	"chapter2/sample/m/v2/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("mars")
}
