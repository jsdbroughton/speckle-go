package models

type AutomateFunctionTemplateLanguage string

const (
	TypeScript AutomateFunctionTemplateLanguage = "TYPESCRIPT"
	Python     AutomateFunctionTemplateLanguage = "PYTHON"
	DotNet     AutomateFunctionTemplateLanguage = "DOT_NET"
)

type AutomateFunctionTemplate struct {
	ID    AutomateFunctionTemplateLanguage `json:"id,omitempty"`
	Logo  *string                          `json:"logo,omitempty"`
	Title *string                          `json:"title,omitempty"`
	URL   *string                          `json:"url,omitempty"`
}

type ServerAutomateInfo struct {
	AvailableFunctionTemplates []AutomateFunctionTemplate `json:"availableFunctionTemplates,omitempty"`
}
