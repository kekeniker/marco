package pipeline_template

import (
	"io"

	"github.com/KeisukeYamashita/marco/pkg/client"
	"github.com/spf13/cobra"
)

// NewPipelineTemplateCmd ...
func NewPipelineTemplateCmd(out io.Writer) *cobra.Command {
	options := &client.PipelineTemplateOptions{}

	cmd := &cobra.Command{
		Use:     "pipeline-template",
		Aliases: []string{"pt"},
	}

	cmd.AddCommand(newPipelineTemplateGetCmd(out, options))
	cmd.AddCommand(newPipelineTemplateListCmd(out, options))
	return cmd
}
