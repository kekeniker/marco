package types

import (
	"testing"
)

func TestParsePipeline(t *testing.T) {
	tcs := map[string]struct {
		input interface{}
		want  *Pipeline
	}{
		"pass": {nil, nil},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {

		})
	}
}
