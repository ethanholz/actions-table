package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ethanholz/action-table/lib"
)

func main() {
	// Provide a file to read
	file := flag.String("file", "", "File to read")
	_ = file
	flag.Parse()
	if *file == "" {
		fmt.Println("Please provide a file to read")
		return
	}
	actions, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	out := lib.GenerateTable(actions)
	print(out)
}
