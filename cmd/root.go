package cmd

import (
	"io"

	"github.com/kekeniker/marco/cmd/application"
	"github.com/kekeniker/marco/cmd/pipeline"
	template "github.com/kekeniker/marco/cmd/pipeline-template"
	vcmd "github.com/kekeniker/marco/cmd/version"
	"github.com/kekeniker/marco/pkg/version"
	"github.com/spf13/cobra"
)

// RootOptions is a options for all commands
type RootOptions struct {
	configFile       string
	GateEndpoint     string
	ignoreCertErrors bool
	quiet            bool
	color            bool
	outputFormat     string
	defaultHeaders   string

	// Duplicated to outputFormat. Output Format is used by spinnaker/spin
	output string
}

// Execute ...
func Execute(out io.Writer) error {
	cmd := newRootCmd(out)
	return cmd.Execute()
}

func newRootCmd(out io.Writer) *cobra.Command {
	options := new(RootOptions)
	cmd := &cobra.Command{
		Use:     "marco",
		Version: version.Version,
	}

	cmd.PersistentFlags().StringVar(&options.configFile, "config", "", "path to config file (default $HOME/.spin/config)")
	cmd.PersistentFlags().StringVar(&options.GateEndpoint, "gate-endpoint", "", "Gate (API server) endpoint (default http://localhost:8084)")
	cmd.PersistentFlags().BoolVarP(&options.ignoreCertErrors, "insecure", "k", false, "ignore certificate errors")
	cmd.PersistentFlags().BoolVarP(&options.quiet, "quiet", "q", true, "squelch non-essential output")
	cmd.PersistentFlags().BoolVar(&options.color, "no-color", true, "disable color")
	cmd.PersistentFlags().StringVar(&options.outputFormat, "log-output", "", "configure spin's log output formatting")
	cmd.PersistentFlags().StringVar(&options.defaultHeaders, "default-headers", "", "configure default headers for gate client as comma separated list (e.g. key1=value1,key2=value2)")
	cmd.PersistentFlags().StringVarP(&options.output, "output", "o", "", "configure the output format. stdout(default), table is supported")

	// Add application subcommand
	cmd.AddCommand(application.NewApplicationCmd(out))

	// Add pipeline subcommand
	cmd.AddCommand(pipeline.NewPipelineCmd(out))

	// Add pipeline-template subcommand
	cmd.AddCommand(template.NewPipelineTemplateCmd(out))

	// Add Version subcommand
	cmd.AddCommand(vcmd.NewVersionCmd(out))

	return cmd
}
