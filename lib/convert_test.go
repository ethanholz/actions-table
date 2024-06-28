package lib

import "testing"

func TestGenerateTable(t *testing.T) {
	// Define the input YAML
	inputYAML := `
name: Sample Action
inputs:
  input1:
    description: The first input
    required: true
  input2:
    description: The second input
    required: false
    default: some value
outputs:
  output1:
    description: The first output
`

	// Define the expected output
	expectedOutput := `# Sample Action

## Inputs
| Name | Description | Required | Default |
| ---- | ----------- | -------- | ------- |
| input1 | The first input | true |  |
| input2 | The second input | false | some value |


## Outputs
| Name | Description |
| ---- | ----------- |
| output1 | The first output |
`

	// Convert the input YAML to bytes
	inputBytes := []byte(inputYAML)

	// Call the GenerateTable function
	actualOutput := GenerateTable(inputBytes)

	// Compare the actual output with the expected output
	if actualOutput != expectedOutput {
		t.Errorf("Output mismatch.\nGot:\n%s\n\nWant:\n%s", actualOutput, expectedOutput)
	}
}
