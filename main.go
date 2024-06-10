package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// struct for parsing a actions.yml file
type Action struct {
	Inputs  map[string]Input  `yaml:"inputs"`
	Outputs map[string]Output `yaml:"outputs"`
	Name    string            `yaml:"name"`
}

type Input struct {
	Description string      `yaml:"description"`
	Required    bool        `yaml:"required"`
	Default     interface{} `yaml:"default,omitempty`
}

type Output struct {
	Description string `yaml:"description"`
}

func main() {
	// Define flags
	var action Action
	file := flag.String("file", "", "File to read")
	flag.Parse()
	if *file == "" {
		fmt.Println("Please provide a file to read")
		return
	}
	actions, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(actions, &action)

	if action.Name != "" {
		fmt.Printf("# %s\n\n", action.Name)
	}

	// Print the markdown table headers
	fmt.Println("## Inputs")
	fmt.Println("| Name | Description | Required | Default |")
	fmt.Println("| ---- | ----------- | -------- | ------- |")

	for name, input := range action.Inputs {
		if input.Default == nil {
			fmt.Printf("| %s | %s | %t |  |\n", name, input.Description, input.Required)
		} else {
			fmt.Printf("| %s | %s | %t | %v |\n", name, input.Description, input.Required, input.Default)
		}
	}

	fmt.Print("\n\n")
	fmt.Println("## Outputs")
	fmt.Println("| Name | Description |")
	fmt.Println("| ---- | ----------- |")
	for name, output := range action.Outputs {
		fmt.Printf("| %s | %s |\n", name, output.Description)
	}
}
