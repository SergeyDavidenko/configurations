package client

// ConfigData ...
type ConfigData struct {
	Name            string            `json:"name"`
	Profiles        []string          `json:"profiles"`
	Label           interface{}       `json:"label"`
	Version         interface{}       `json:"version"`
	State           interface{}       `json:"state"`
	PropertySources []PropertySources `json:"propertySources"`
}

// PropertySources ...
type PropertySources struct {
	Name   string `json:"name"`
	Source Source `json:"source"`
}

// Source ...
type Source map[string]string