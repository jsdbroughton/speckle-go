package schema

import "encoding/json"

// AutomationStatus represents the status of the automation
type AutomationStatus string

const (
	StatusInitializing AutomationStatus = "INITIALIZING"
	StatusRunning      AutomationStatus = "RUNNING"
	StatusFailed       AutomationStatus = "FAILED"
	StatusSucceeded    AutomationStatus = "SUCCEEDED"
	StatusException    AutomationStatus = "EXCEPTION"
)

// AutomationResult represents the schema accepted by the Speckle server as a result for an automation run
type AutomationResult struct {
	Elapsed        float64          `json:"elapsed"`
	ResultView     *string          `json:"resultView,omitempty"`
	ResultVersions []string         `json:"resultVersions"`
	Blobs          []string         `json:"blobs"`
	RunStatus      AutomationStatus `json:"runStatus"`
	StatusMessage  *string          `json:"statusMessage,omitempty"`
	ObjectResults  []ResultCase     `json:"objectResults"`
}

// ResultGenerator defines the interface for generating automation results
type ResultGenerator interface {
	GenerateResult() AutomationResult
}

// ObjectResultReporter defines the interface for reporting object results
type ObjectResultReporter interface {
	ReportResult(ResultCase) error
}

// ResultCase represents a result case
type ResultCase struct {
	Category        string                 `json:"category"`
	Level           ObjectResultLevel      `json:"level"`
	ObjectIds       []string               `json:"objectIds"`
	Message         *string                `json:"message,omitempty"`
	Metadata        *ResultMetadata        `json:"metadata,omitempty"`
	VisualOverrides map[string]interface{} `json:"visualOverrides,omitempty"`
}

// ObjectResultLevel represents possible status message levels for object reports
type ObjectResultLevel string

const (
	LevelInfo    ObjectResultLevel = "INFO"
	LevelWarning ObjectResultLevel = "WARNING"
	LevelError   ObjectResultLevel = "ERROR"
)

// ResultMetadata represents the metadata structure for a result
type ResultMetadata struct {
	Gradient       *bool                    `json:"gradient,omitempty"`
	GradientValues map[string]GradientValue `json:"gradientValues,omitempty"`
	ExtraData      map[string]interface{}   `json:"-"`
}

// GradientValue represents a gradient value
type GradientValue struct {
	GradientValue float64 `json:"gradientValue"`
}

// UnmarshalJSON implements custom unmarshalling for ResultMetadata
func (r *ResultMetadata) UnmarshalJSON(data []byte) error {
	type Alias ResultMetadata
	aux := &struct {
		*Alias
		ExtraData map[string]interface{} `json:"extraData,omitempty"`
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.ExtraData = aux.ExtraData
	return nil
}

// MarshalJSON implements custom marshalling for ResultMetadata
func (r *ResultMetadata) MarshalJSON() ([]byte, error) {
	type Alias ResultMetadata
	aux := &struct {
		*Alias
		ExtraData map[string]interface{} `json:"extraData,omitempty"`
	}{
		Alias:     (*Alias)(r),
		ExtraData: r.ExtraData,
	}
	return json.Marshal(aux)
}
