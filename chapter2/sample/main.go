package main

import (
	"log"
	// "m/v2/chapter2/sample/search"
	"os"

	_ "m/v2/chapter2/sample/matchers"
	"m/v2/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("earth")
}
