package version

import (
	"io"

	"github.com/spf13/cobra"
)

const (
	// Version ...
	Version = "0.1.0"
)

// NewVersionCmd ...
func NewVersionCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			return
		},
	}

	return cmd
}
