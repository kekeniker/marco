package application

import (
	"io"

	"github.com/kekeniker/marco/pkg/client"
	"github.com/kekeniker/marco/pkg/output"
	"github.com/kekeniker/marco/pkg/validator"
	"github.com/spf13/cobra"
)

var (
	defaultHeaders = []string{"name", "email", "cloud providers", "instance port"}
)

// newApplicationListCmd ...
func newApplicationListCmd(out io.Writer, appOptions *client.AppOptions) *cobra.Command {
	options := &client.AppListOptions{
		AppOptions: appOptions,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		RunE:    listApplication(out, options),
	}

	cmd.PersistentFlags().BoolVarP(&options.Validate, "validate", "", false, "validate Spinnaker application")
	return cmd
}

// listApplication ...
func listApplication(out io.Writer, options *client.AppListOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c, err := client.New(cmd.InheritedFlags())
		if err != nil {
			return err
		}

		apps, err := c.ListApplications(options)
		if err != nil {
			return err
		}

		outputManager, err := output.NewOutputManager("table", out)
		if err != nil {
			return err
		}

		header := getHeader(options)
		outputManager.SetHeader(header)
		for _, app := range apps {
			if err := outputManager.Put(&app); err != nil {
				return err
			}
		}

		outputManager.Flush()
		return nil
	}
}

func getHeader(options *client.AppListOptions) []string {
	res := defaultHeaders

	if options.Validate {
		res = append(res, validator.SortedCloudProviders()...)
	}
	return res
}
