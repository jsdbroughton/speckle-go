package speckle_automate

import (
	"encoding/json"
	"fmt"
	"github.com/jsdbroughton/speckle-go/internal/schema"
	"os"
)

// AutomateFunction represents a Speckle Automate function
type AutomateFunction struct {
	Inputs interface{}
	Run    func(inputs interface{}) error
}

// Execute runs the Automate function
func (af *AutomateFunction) Execute() error {
	if err := json.Unmarshal([]byte(os.Getenv("SPECKLE_FUNCTION_INPUT")), af.Inputs); err != nil {
		return fmt.Errorf("failed to parse function inputs: %v", err)
	}
	return af.Run(af.Inputs)
}

// GenerateSchema generates the JSON schema for the function inputs
func (af *AutomateFunction) GenerateSchema(outputPath string) error {
	generator := &schema.GenerateAutomateJsonSchema{
		SchemaDialect: "http://json-schema.org/draft-07/schema#",
	}

	generatedSchema, err := generator.Generate(af.Inputs, "validation")
	if err != nil {
		return fmt.Errorf("failed to generate schema: %v", err)
	}

	data, err := json.MarshalIndent(generatedSchema, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal schema: %v", err)
	}

	return os.WriteFile(outputPath, data, 0644)
}

// RunAutomate is the main entry point for Automate functions
func RunAutomate(af *AutomateFunction) {
	if len(os.Args) > 1 && os.Args[1] == "generate-schema" {
		if len(os.Args) != 3 {
			fmt.Println("Usage: <binary> generate-schema <output file path>")
			os.Exit(1)
		}
		if err := af.GenerateSchema(os.Args[2]); err != nil {
			fmt.Printf("Error generating schema: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Schema generated successfully:", os.Args[2])
	} else {
		if err := af.Execute(); err != nil {
			fmt.Printf("Error executing function: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Function executed successfully")
	}
}
