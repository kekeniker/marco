package types

import (
	"fmt"
	"strings"
)

// Application ...
type Application struct {
	CloudProviders []string
	Email          string
	InstancePort   int
	Name           string

	// Additional field to API ObJect
	NameCheckResults map[string]*NameCheckResult
}
type NameCheckResult struct {
	Valid bool
	Error error
}

// Applications ...
type Applications []Application

// Len returns the length of the sotring object
func (apps Applications) Len() int {
	return len(apps)
}

// Less compares the two objects and determine the order
func (apps Applications) Less(i, j int) bool {
	if apps[i].Name < apps[j].Name {
		return true
	}

	return false
}

// Swap swaps the two objects
func (apps Applications) Swap(i, j int) {
	apps[i], apps[j] = apps[j], apps[i]
}

// ParseApplication ...
func ParseApplication(input interface{}, app *Application) error {
	switch appRead := input.(type) {
	case map[string]interface{}:
		if v, ok := appRead["cloudProviders"].(string); ok {
			app.CloudProviders = strings.Split(v, ",")
		}

		if v, ok := appRead["email"].(string); ok {
			app.Email = v
		}

		if v, ok := appRead["name"].(string); ok {
			app.Name = v
		}

		if v, ok := appRead["instancePort"].(int); ok {
			app.InstancePort = v
		}

		return nil
	default:
		return fmt.Errorf("application was not expected map[string]interface was: %T", appRead)
	}
}
