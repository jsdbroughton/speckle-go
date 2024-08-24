package schema

// AutomationRunData represents the values of the project/model that triggered the run of this function
type AutomationRunData struct {
	ProjectID        string                   `json:"projectId"`
	SpeckleServerURL string                   `json:"speckleServerUrl"`
	AutomationID     string                   `json:"automationId"`
	AutomationRunID  string                   `json:"automationRunId"`
	FunctionRunID    string                   `json:"functionRunId"`
	Triggers         []VersionCreationTrigger `json:"triggers"`
}

// VersionCreationTrigger represents a single version creation trigger for the automation run
type VersionCreationTrigger struct {
	TriggerType string                        `json:"triggerType"`
	Payload     VersionCreationTriggerPayload `json:"payload"`
}

// VersionCreationTriggerPayload represents the payload of a version creation trigger
type VersionCreationTriggerPayload struct {
	ModelID   string `json:"modelId"`
	VersionID string `json:"versionId"`
}

// TestAutomationRunData contains values of the run created in the test automation for local test results
type TestAutomationRunData struct {
	AutomationRunID string                   `json:"automationRunId"`
	FunctionRunID   string                   `json:"functionRunId"`
	Triggers        []VersionCreationTrigger `json:"triggers"`
}
