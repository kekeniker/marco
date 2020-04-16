package output

import (
	"fmt"
	"io"
)

const (
	outputTable = "table"
)

// OutputManager controls how results of an evaluation will be recorded and reported to the end user
type OutputManager interface {
	SetHeader([]string)
	Put(interface{}) error
	Flush()
}

type options struct {
	AutoMergeCells bool
}
type outputOptions func(*options)

// WithAutoMergeCells  ...
func WithAutoMergeCells() outputOptions {
	return func(o *options) {
		o.AutoMergeCells = true
	}
}

// NewOutputManager ...
func NewOutputManager(format string, out io.Writer, opts ...outputOptions) (OutputManager, error) {
	switch format {
	case outputTable:
		return newDefaultTableOutputManager(out, opts...), nil
	default:
		return nil, fmt.Errorf("output %s is not supported", format)
	}
}
