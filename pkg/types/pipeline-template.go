package types

import (
	"fmt"
)

// PipelineTemplate ...
type PipelineTemplate struct {
	LastModifiedBy string
	ID             string
	Metadata       *Metadata
	Pipeline       *Pipeline
	Protect        bool
	Schema         string
	Tag            string
	Variables      []*Variable
	UpdateTs       string
}

// Metadata ...
type Metadata struct {
	Description string
	Name        string
	Owner       string
	Scopes      []string
}

// Variable ...
type Variable struct {
	DefaultValue interface{}
	Description  string
	Name         string
	Type         string
}

// ParsePipelineTemplate ...
func ParsePipelineTemplate(input interface{}, pt *PipelineTemplate) error {
	switch ptRead := input.(type) {
	case map[string]interface{}:
		if v, ok := ptRead["lastModifiedBy"].(string); ok {
			pt.ID = v
		}

		if v, ok := ptRead["id"].(string); ok {
			pt.ID = v
		}

		if v, ok := ptRead["metadata"]; ok {
			mt := new(Metadata)
			if err := parseMetadata(v, mt); err != nil {
				return err
			}

			pt.Metadata = mt
		}

		if v, ok := ptRead["pipeline"]; ok {
			pl := new(Pipeline)
			if err := ParsePipeline(v, pl); err != nil {
				return err
			}

			pt.Pipeline = pl
		}

		if v, ok := ptRead["protect"].(bool); ok {
			pt.Protect = v
		}

		if v, ok := ptRead["schema"].(string); ok {
			pt.Schema = v
		}

		if v, ok := ptRead["tag"].(string); ok {
			pt.Tag = v
		}

		if v, ok := ptRead["updateTs"].(string); ok {
			pt.UpdateTs = v
		}

		if vals, ok := ptRead["variables"].([]interface{}); ok {
			variables := make([]*Variable, len(vals))
			for _, val := range vals {
				if input, ok := val.(map[string]interface{}); ok {
					variable := &Variable{}
					if err := parseVariable(input, variable); err != nil {
						variables = append(variables, variable)
					}
				}
			}

			pt.Variables = variables
		}

		return nil
	default:
		return fmt.Errorf("pipeline template was not expected map[string]interface was: %T", ptRead)
	}
}

func parseMetadata(input interface{}, mt *Metadata) error {
	switch input := input.(type) {
	case map[string]interface{}:
		if v, ok := input["description"].(string); ok {
			mt.Description = v
		}

		if v, ok := input["name"].(string); ok {
			mt.Name = v
		}

		if v, ok := input["owner"].(string); ok {
			mt.Owner = v
		}

		if v, ok := input["scopes"].([]string); ok {
			mt.Scopes = v
		}

		return nil
	default:
		return fmt.Errorf("metadata template was not expected map[string]interface was: %T", input)
	}
}

func parseVariable(input map[string]interface{}, val *Variable) error {
	if v, ok := input["defaultValue"]; ok {
		val.DefaultValue = v
	}

	if v, ok := input["description"].(string); ok {
		val.Description = v
	}

	if v, ok := input["name"].(string); ok {
		val.Name = v
	}

	if v, ok := input["type"].(string); ok {
		val.Type = v
	}

	return nil
}
