package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ethanholz/action-table/lib"
)

var Version = "dev"

func main() {
	// Provide a file to read
	file := flag.String("file", "", "File to read")
	version := flag.Bool("version", false, "Print the version")
	flag.Parse()
	if *version {
		fmt.Printf("action-table %s", Version)
		os.Exit(0)
	}
	if *file == "" {
		fmt.Println("Please provide a file to read")
		os.Exit(1)
	}
	actions, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	out := lib.GenerateTable(actions)
	print(out)
}
