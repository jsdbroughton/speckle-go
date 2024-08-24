package schema

import (
	"github.com/jsdbroughton/speckle-go/pkg/speckle_automate/automation_context"
)

// AutomateFunction represents a function that can be executed by the automation system
type AutomateFunction interface {
	Run(ctx *automation_context.AutomationContext, inputs interface{}) error
}

// AutomateFunctionWithoutInputs represents a function without inputs that can be executed by the automation system
type AutomateFunctionWithoutInputs interface {
	Run(ctx *automation_context.AutomationContext) error
}

// FunctionRunner represents the interface for running automate functions
type FunctionRunner interface {
	ExecuteAutomateFunction(automateFunction interface{}, inputSchema interface{}) error
}
