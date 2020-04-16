package pipeline

import (
	"io"

	"github.com/KeisukeYamashita/marco/pkg/client"
	"github.com/spf13/cobra"
)

// NewPipelineCmd ...
func NewPipelineCmd(out io.Writer) *cobra.Command {
	options := &client.PipelineOptions{}

	cmd := &cobra.Command{
		Use:     "pipeline",
		Aliases: []string{"pi"},
	}

	cmd.AddCommand(newPipelineGetCmd(out, options))
	cmd.AddCommand(newPipelineListCmd(out, options))
	return cmd
}
