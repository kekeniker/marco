package pipeline_template

import (
	"io"

	"github.com/KeisukeYamashita/marco/pkg/client"
	"github.com/KeisukeYamashita/marco/pkg/output"
	"github.com/spf13/cobra"
)

// newPipelineTemplateListCmd ...
func newPipelineTemplateListCmd(out io.Writer, options *client.PipelineTemplateOptions) *cobra.Command {
	listOptions := &client.PipelineTemplateListOptions{
		PipelineTemplateOptions: options,
	}

	cmd := &cobra.Command{
		Use:  "list",
		Aliases: []string{"ls"},
		RunE: listPipelineTemplateCmd(out, listOptions),
	}

	return cmd
}

func listPipelineTemplateCmd(out io.Writer, options *client.PipelineTemplateListOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c, err := client.New(cmd.InheritedFlags())
		if err != nil {
			return err
		}

		pts, err := c.ListPipelineTemplates(options)
		if err != nil {
			return err
		}

		outputManager, err := output.NewOutputManager("table", out)
		if err != nil {
			return err
		}

		header := getPipelineTemplateListHeader(options)
		outputManager.SetHeader(header)

		for _, pt := range pts {
			outputManager.Put(&pt)
		}

		outputManager.Flush()
		return nil
	}
}

func getPipelineTemplateListHeader(options *client.PipelineTemplateListOptions) []string {
	return []string{"id", "protected"}
}
