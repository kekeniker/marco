package pipeline_template

import (
	"io"

	"github.com/KeisukeYamashita/marco/pkg/client"
	"github.com/KeisukeYamashita/marco/pkg/output"
	"github.com/spf13/cobra"
)

func newPipelineTemplateGetCmd(out io.Writer, options *client.PipelineTemplateOptions) *cobra.Command {
	getOptions := &client.PipelineTemplateGetOptions{
		PipelineTemplateOptions: options,
	}

	cmd := &cobra.Command{
		Use:  "get",
		RunE: getPipelineTemplateCmd(out, getOptions),
	}

	return cmd
}

func getPipelineTemplateCmd(out io.Writer, options *client.PipelineTemplateGetOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c, err := client.New(cmd.InheritedFlags())
		if err != nil {
			return err
		}

		pt, err := c.GetPipelineTemplate(args[0], options)
		if err != nil {
			return err
		}

		outputManager, err := output.NewOutputManager("table", out)
		if err != nil {
			return err
		}

		header := getPipelineTemplateGetHeader(options)
		outputManager.SetHeader(header)

		outputManager.Put(pt)
		outputManager.Flush()
		return nil
	}
}

func getPipelineTemplateGetHeader(options *client.PipelineTemplateGetOptions) []string {
	return []string{"id", "proteced"}
}
