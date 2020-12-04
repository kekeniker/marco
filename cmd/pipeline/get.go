package pipeline

import (
	"io"

	"github.com/kekeniker/marco/pkg/client"
	"github.com/kekeniker/marco/pkg/output"
	"github.com/spf13/cobra"
)

// newPipelineGetCmd ...
func newPipelineGetCmd(out io.Writer, options *client.PipelineOptions) *cobra.Command {
	getOptions := &client.PipelineGetOptions{
		PipelineOptions: options,
	}

	cmd := &cobra.Command{
		Use:  "get",
		RunE: getPipeline(out, getOptions),
	}

	cmd.PersistentFlags().BoolVar(&getOptions.Expand, "expand", true, "expand pipeline template")
	cmd.PersistentFlags().StringVarP(&getOptions.App, "application", "a", "", "specify Spinnaker application")
	return cmd
}

// getPipeline ...
func getPipeline(out io.Writer, options *client.PipelineGetOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c, err := client.New(cmd.InheritedFlags())
		if err != nil {
			return err
		}

		pipeline, err := c.GetPipeline(args[0], options)
		if err != nil {
			return err
		}

		opt := output.WithAutoMergeCells()
		outputManager, err := output.NewOutputManager("table", out, opt)
		if err != nil {
			return err
		}

		header := getPipelineGetHeader(options)
		outputManager.SetHeader(header)
		outputManager.Put(pipeline)

		outputManager.Flush()
		return nil
	}
}

func getPipelineGetHeader(options *client.PipelineGetOptions) []string {
	return []string{"app", "Pipeline id", "name", "template", "stage Refid", "stage", "type"}
}
