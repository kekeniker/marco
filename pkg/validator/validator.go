package validator

import (
	"fmt"
	"regexp"
	"sort"
)

var (
	// CloudProviders ...
	// See details in Spinnaker Orca
	// ref: https://github.com/spinnaker/orca/blob/master/orca-applications/src/main/groovy/com/netflix/spinnaker/orca/applications/utils/ApplicationNameValidator.groovy
	CloudProviders = map[string]applicationNameConstraint{
		"appengine":    applicationNameConstraint{58, `^[a-z0-9]*$`},
		"aws":          applicationNameConstraint{250, `^[a-zA-Z_0-9.]*$`},
		"dcos":         applicationNameConstraint{127, `^[a-z0-9]*$`},
		"kubernetes":   applicationNameConstraint{63, `^([a-zA-Z][a-zA-Z0-9-]*)$`},
		"gce":          applicationNameConstraint{63, `^([a-zA-Z][a-zA-Z0-9]*)?$`},
		"openstack":    applicationNameConstraint{250, `^[a-zA-Z_0-9.]*$`},
		"tencentcloud": applicationNameConstraint{50, `^[a-zA-Z_0-9.-]*$`},
		"titus":        applicationNameConstraint{250, `^[a-zA-Z_0-9.]*$`},
	}
)

// applicationNameConstraint ...
type applicationNameConstraint struct {
	maxLength int
	regex     string
}

// ValidApplicationNameByCloudProvider validates te Spinnaker application name by provider rules
func ValidApplicationNameByCloudProvider(appName, provider string) (bool, error) {
	if regex, ok := CloudProviders[provider]; ok {
		if !regexp.MustCompile(regex.regex).MatchString(appName) {
			return false, nil
		}

		if c := len(appName); c > regex.maxLength {
			return false, nil
		}

		return true, nil
	}

	return false, fmt.Errorf("cloud provider %s is not supported", provider)
}

// SortedCloudProviders ...
func SortedCloudProviders() []string {
	keys := make([]string, len(CloudProviders))
	i := 0
	for k := range CloudProviders {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
