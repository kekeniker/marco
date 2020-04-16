package pipeline

import (
	"io"

	"github.com/KeisukeYamashita/marco/pkg/client"
	"github.com/KeisukeYamashita/marco/pkg/output"
	"github.com/spf13/cobra"
)

// newPipelineListCmd ...
func newPipelineListCmd(out io.Writer, options *client.PipelineOptions) *cobra.Command {
	listOptions := &client.PipelineListOptions{
		PipelineOptions: options,
	}

	cmd := &cobra.Command{
		Use:  "list",
		Aliases: []string{"ls"},
		RunE: listPipeline(out, listOptions),
	}

	cmd.PersistentFlags().BoolVar(&listOptions.Expand, "expand", true, "expand pipeline template")
	cmd.PersistentFlags().StringVarP(&listOptions.App, "application", "a", "", "specify Spinnaker application")
	cmd.PersistentFlags().BoolVar(&listOptions.All, "all", false, "run for all Spinnaker applications")
	return cmd
}

// listPipeline ...
func listPipeline(out io.Writer, options *client.PipelineListOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c, err := client.New(cmd.InheritedFlags())
		if err != nil {
			return err
		}

		pipelines, err := c.ListPipelines(options)
		if err != nil {
			return err
		}

		opt := output.WithAutoMergeCells()
		outputManager, err := output.NewOutputManager("table", out, opt)
		header := getPipelineListHeader(options)
		outputManager.SetHeader(header)

		for _, pipeline := range pipelines {
			outputManager.Put(&pipeline)
		}

		outputManager.Flush()
		return nil
	}
}

func getPipelineListHeader(options *client.PipelineListOptions) []string {
	return []string{"app", "Pipeline id", "name", "template", "stage Refid", "stage", "type"}
}
