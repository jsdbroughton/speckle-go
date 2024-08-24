package main

import "github.com/jsdbroughton/speckle-go/pkg/speckle_automate"

// FunctionInputs represents the inputs for your specific Automate function
type FunctionInputs struct {
	Param1 string `json:"param1"`
	Param2 int    `json:"param2"`
}

func main() {
	speckle_automate.RunAutomate(&speckle_automate.AutomateFunction{
		Inputs: &FunctionInputs{},
		Run: func(inputs interface{}) error {
			// Type assert to access the specific input fields
			// typedInputs := inputs.(*FunctionInputs)

			// Your function logic here
			// For example:
			// fmt.Printf("Received Param1: %s, Param2: %d\n", typedInputs.Param1, typedInputs.Param2)

			return nil // or return an error if something goes wrong
		},
	})
}
