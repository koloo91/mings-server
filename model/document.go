package model

type DocumentType string

const (
	ServiceType DocumentType = "Service"
)

type Documents struct {
	Documents []Document `json:"documents"`
}

type Document struct {
	Id          string            `json:"id" yaml:"id"`
	Name        string            `json:"name" yaml:"name"`
	Type        DocumentType      `json:"type" yaml:"type"`
	Owner       string            `json:"owner" yaml:"owner"`
	Description string            `json:"description" yaml:"description"`
	ShortName   string            `json:"shortName" yaml:"short_name"`
	Contact     string            `json:"contact" yaml:"contact"`
	Tags        []string          `json:"tags" yaml:"tags"`
	Links       map[string]string `json:"links" yaml:"links"`
	Service     Service           `json:"service" yaml:"service"`
}

type Service struct {
	Provides  []Provide `json:"provides" yaml:"provides"`
	DependsOn DependsOn `json:"dependsOn" yaml:"depends_on"`
}

type Provide struct {
	Description       string `json:"description" yaml:"description"`
	ServiceName       string `json:"serviceName" yaml:"service_name"`
	Protocol          string `json:"protocol" yaml:"protocol"`
	Port              int    `json:"port" yaml:"port"`
	TransportProtocol string `json:"transportProtocol" yaml:"transport_protocol"`
}

type DependsOn struct {
	Internal []DependsOnService `json:"internal" yaml:"internal"`
	External []DependsOnService `json:"external" yaml:"external"`
}

type DependsOnService struct {
	ServiceName string `json:"serviceName" yaml:"service_name"`
	Why         string `json:"why" yaml:"why"`
}
