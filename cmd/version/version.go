package version

import (
	"fmt"
	"io"

	"github.com/kekeniker/pkg/version"
	"github.com/spf13/cobra"
)

// NewVersionCmd ...
func NewVersionCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Version)
		},
	}

	return cmd
}
