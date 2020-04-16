package application

import (
	"io"

	"github.com/KeisukeYamashita/marco/pkg/client"
	"github.com/spf13/cobra"
)

// NewApplicationCmd ...
func NewApplicationCmd(out io.Writer) *cobra.Command {
	options := &client.AppOptions{}

	cmd := &cobra.Command{
		Use:     "application",
		Aliases: []string{"app"},
	}

	cmd.AddCommand(newApplicationListCmd(out, options))
	return cmd
}
