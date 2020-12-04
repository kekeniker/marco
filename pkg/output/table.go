package output

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kekeniker/marco/pkg/types"
	"github.com/kekeniker/marco/pkg/validator"
	"github.com/olekukonko/tablewriter"
)

// TableOutputManager ...
type TableOutputManager struct {
	tableWriter *tablewriter.Table
}

// SetHeader ...
func (mgr *TableOutputManager) SetHeader(header []string) {
	mgr.tableWriter.SetHeader(header)
}

// Put ...
func (mgr *TableOutputManager) Put(input interface{}) error {
	switch obj := input.(type) {
	case *types.Application:
		return mgr.putApplication(obj)
	case *types.Pipeline:
		return mgr.putPipeline(obj)
	case *types.PipelineTemplate:
		return mgr.putPipelineTemplate(obj)
	default:
		return nil
	}
}

// Flush ...
func (mgr *TableOutputManager) Flush() {
	mgr.tableWriter.Render()
}

// newDefaultTableOutputManager ...
func newDefaultTableOutputManager(out io.Writer, opts ...outputOptions) *TableOutputManager {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	table := tablewriter.NewWriter(out)

	if o.AutoMergeCells {
		table.SetAutoMergeCells(true)
	}

	table.SetRowLine(true)

	return &TableOutputManager{
		tableWriter: table,
	}
}

func (mgr *TableOutputManager) putApplication(app *types.Application) error {
	row := []string{
		app.Name,
		app.Email,
		strings.Join(app.CloudProviders, ","),
		strconv.Itoa(app.InstancePort),
	}

	if app.NameCheckResults != nil {
		providers := validator.SortedCloudProviders()
		for _, provider := range providers {
			var msg string
			result := app.NameCheckResults[provider]
			if result.Valid {
				msg = "✅"
			} else {
				msg = "❌"
			}

			row = append(row, msg)
		}
	}

	mgr.tableWriter.Append(row)
	return nil
}

func (mgr *TableOutputManager) putPipeline(pl *types.Pipeline) error {
	row := []string{
		pl.App,
		pl.ID,
		pl.Name,
	}

	if len(pl.Stages) == 0 {
		if pl.Template != nil {
			template := strings.Split(pl.Template.Reference, "://")[1]
			mergedRow := append(row, []string{template, "(template)", "(template)", "(template)"}...)
			mgr.tableWriter.Append(mergedRow)
			return nil
		}

		mergedRow := append(row, []string{"(none)", "(none)", "(none)", "(none)"}...)
		mgr.tableWriter.Append(mergedRow)
		return nil
	}

	for _, stage := range pl.Stages {
		template := "(none)"
		mergedRow := append(row, []string{template, stage.RefID, stage.Name, stage.Type}...)
		mgr.tableWriter.Append(mergedRow)
	}

	return nil
}

func (mgr *TableOutputManager) putPipelineTemplate(pl *types.PipelineTemplate) error {
	row := []string{
		fmt.Sprintf("%s:%s", pl.ID, pl.Tag),
		strconv.FormatBool(pl.Protect),
	}

	mgr.tableWriter.Append(row)
	return nil
}
