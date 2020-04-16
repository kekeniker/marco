package types

import "fmt"

// Pipeline ...
type Pipeline struct {
	App             string
	ID              string
	LastyModifiedBy string
	Name            string
	Stages          []*Stage
	Template        *Template
	Triggers        []*Trigger

	// Additional
	V1                  bool
	ConfiguredVariables ConfiguredVariables
}

// Stage ...
type Stage struct {
	Clusters []*Cluster
	Name     string
	RefID    string
	Type     string
}

// Trigger ...
type Trigger struct {
	Account      string
	Enabled      bool
	Organization string
	Registory    string
	Repository   string
	Type         string
}

// ConfiguredVariables ...
type ConfiguredVariables map[string]interface{}

// Cluster ...
type Cluster struct {
	Account                        string
	Application                    string
	CloudProvider                  string
	Containers                     []*Container
	Deployment                     *Deployment
	DNSPolicy                      string
	InitContainers                 []*InitContainer
	InterestingHealthProviderNames []string
	LoadBalancers                  []string
	Namespace                      string
	NodeSelector                   *NodeSelector
	PodAnnotations                 *PodAnnotations
	Provider                       string
	Region                         string
	ReplicaSetAnnocations          map[string]string
	Stack                          string
	Strategy                       string
	TargetSize                     int
	TerminationGracePeriodSeconds  int
	UseSourceCapacity              bool
	VolumnSources                  []*VolumnSource
}

// Container ...
type Container struct{}

// Deployment ...
type Deployment struct{}

// InitContainer ...
type InitContainer struct{}

// NodeSelector ...
type NodeSelector struct{}

// PodAnnotations ...
type PodAnnotations struct{}

// VolumnSource ...
type VolumnSource struct{}

// Template ...
type Template struct {
	ArtifactAccount string
	Reference       string
	Type            string
}

// Moniker ...
type Moniker map[string]interface{}

// ParsePipeline ...
func ParsePipeline(input interface{}, pl *Pipeline) error {
	switch plRead := input.(type) {
	case map[string]interface{}:
		if v, ok := plRead["application"].(string); ok {
			pl.App = v
		}

		if v, ok := plRead["id"].(string); ok {
			pl.ID = v
		}

		if v, ok := plRead["name"].(string); ok {
			pl.Name = v
		}

		if v, ok := plRead["stages"].([]interface{}); ok {
			stages := make([]*Stage, len(v))
			for i, stageRead := range v {
				stage := &Stage{}
				if err := ParseStage(stageRead, stage); err != nil {
					return err
				}

				stages[i] = stage
			}

			pl.Stages = stages
		}

		if v, ok := plRead["template"].(map[string]interface{}); ok {
			template := &Template{}
			if err := parseTemplate(v, template); err != nil {
				return err
			}

			pl.Template = template
		}

		if v, ok := plRead["variables"].(ConfiguredVariables); ok {
			val := ConfiguredVariables{}
			for k, value := range v {
				v[k] = value
			}

			pl.ConfiguredVariables = val
		}

		return nil
	default:
		return fmt.Errorf("pipeline was not expected map[string]interface was: %T", plRead)
	}
}

// ParseStage ...
func ParseStage(input interface{}, stage *Stage) error {
	switch stageRead := input.(type) {
	case map[string]interface{}:
		if v, ok := stageRead["refId"].(string); ok {
			stage.RefID = v
		}

		if v, ok := stageRead["name"].(string); ok {
			stage.Name = v
		}

		if v, ok := stageRead["type"].(string); ok {
			stage.Type = v
		}

		return nil
	default:
		return fmt.Errorf("stage was not expected map[string]interface was: %T", stage)
	}
}

// ParseTemplate ...
func parseTemplate(input interface{}, template *Template) error {
	switch templateRead := input.(type) {
	case map[string]interface{}:
		if v, ok := templateRead["artifactAccount"].(string); ok {
			template.ArtifactAccount = v
		}

		if v, ok := templateRead["reference"].(string); ok {
			template.Reference = v
		}

		if v, ok := templateRead["type"].(string); ok {
			template.Type = v
		}

		return nil
	default:
		return fmt.Errorf("template was not expected map[string]interface was: %T", template)
	}
}
