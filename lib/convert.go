package lib

import (
	"fmt"
	"strings"

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

func GenerateTable(input []byte) string {
	var action Action
	var builder strings.Builder
	yaml.Unmarshal(input, &action)
	if action.Name != "" {
		fmt.Fprintf(&builder, "# %s\n\n", action.Name)
	}

	builder.WriteString("## Inputs\n")
	builder.WriteString("| Name | Description | Required | Default |\n")
	builder.WriteString("| ---- | ----------- | -------- | ------- |\n")
	for name, input := range action.Inputs {
		if input.Default == nil {
			fmt.Fprintf(&builder, "| %s | %s | %t |  |\n", name, input.Description, input.Required)
		} else {
			fmt.Fprintf(&builder, "| %s | %s | %t | %v |\n", name, input.Description, input.Required, input.Default)
		}
	}
	builder.WriteString("\n\n")
	builder.WriteString("## Outputs\n")
	builder.WriteString("| Name | Description |\n")
	builder.WriteString("| ---- | ----------- |\n")
	for name, output := range action.Outputs {
		fmt.Fprintf(&builder, "| %s | %s |\n", name, output.Description)
	}

	return builder.String()
}
