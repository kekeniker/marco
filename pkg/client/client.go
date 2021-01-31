package client

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/kekeniker/marco/pkg/types"
	"github.com/kekeniker/marco/pkg/validator"
	"github.com/spf13/pflag"
	"github.com/spinnaker/spin/cmd/gateclient"
	"github.com/spinnaker/spin/cmd/output"
	gate "github.com/spinnaker/spin/gateapi"
)

type AppOptions struct{}
type AppListOptions struct {
	AppOptions *AppOptions
	Validate   bool
}

type PipelineOptions struct{}
type PipelineGetOptions struct {
	PipelineOptions *PipelineOptions
	App             string
	Expand          bool
}
type PipelineListOptions struct {
	PipelineOptions *PipelineOptions
	App             string
	All             bool
	Expand          bool
}

type PipelineTemplateOptions struct{}
type PipelineTemplateListOptions struct {
	PipelineTemplateOptions *PipelineTemplateOptions
}
type PipelineTemplateGetOptions struct {
	PipelineTemplateOptions *PipelineTemplateOptions
}

// Client is an interface for REST PRC for Spinnaker Gateway.
type Client interface {
	ListApplications(*AppListOptions) ([]types.Application, error)
	GetPipeline(string, *PipelineGetOptions) (*types.Pipeline, error)
	ListPipelines(*PipelineListOptions) ([]types.Pipeline, error)
	GetPipelineTemplate(string, *PipelineTemplateGetOptions) (*types.PipelineTemplate, error)
	ListPipelineTemplates(*PipelineTemplateListOptions) ([]types.PipelineTemplate, error)
}

// ClientImpl is the implement of the GateClient ...
type ClientImpl struct {
	spinnaker *gateclient.GatewayClient
}

var _ Client = (*ClientImpl)(nil)

// ClientOption is what you can configure for the your client
type ClientOption func(*options)

type options struct {
}

// New creates a Client that implementates GateClient
func New(flags *pflag.FlagSet, opts ...ClientOption) (Client, error) {
	var o *options
	for _, opt := range opts {
		opt(o)
	}

	outputFormater, err := output.ParseOutputFormat("")
	if err != nil {
		return nil, err
	}
	ui := output.NewUI(false, true, outputFormater, os.Stdout, os.Stderr)
	endpoint, err := flags.GetString("gate-endpoint")
	if err != nil {
		return nil, err
	}
	headers, err := flags.GetString("default-headers")
	if err != nil {
		return nil, err
	}
	config, err := flags.GetString("config")
	if err != nil {
		return nil, err
	}
	ignoreCertErrors, err := flags.GetBool("insecure")
	if err != nil {
		return nil, err
	}

	gateClient, err := gateclient.NewGateClient(ui, endpoint, headers, config, ignoreCertErrors)
	if err != nil {
		return nil, err
	}

	return &ClientImpl{
		spinnaker: gateClient,
	}, nil
}

// ListApplications list all applications
func (c *ClientImpl) ListApplications(options *AppListOptions) ([]types.Application, error) {
	opts := &gate.ApplicationControllerApiGetAllApplicationsUsingGETOpts{}
	apps, resp, err := c.spinnaker.ApplicationControllerApi.GetAllApplicationsUsingGET(c.spinnaker.Context, opts)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("encountered an error listing application, status code: %d", resp.StatusCode)
	}

	appsRead := make([]types.Application, len(apps))

	for i, app := range apps {
		var v types.Application
		if err := types.ParseApplication(app, &v); err != nil {
			return nil, err
		}

		if options.Validate {
			v.NameCheckResults = map[string]*types.NameCheckResult{}
			for provider := range validator.CloudProviders {
				valid, err := validator.ValidApplicationNameByCloudProvider(v.Name, provider)
				v.NameCheckResults[provider] = &types.NameCheckResult{
					Valid: valid,
					Error: err,
				}
			}
		}

		appsRead[i] = v
	}

	sort.Sort(types.Applications(appsRead))
	return appsRead, nil
}

// GetPipeline gets a pipeline by ID
func (c *ClientImpl) GetPipeline(name string, options *PipelineGetOptions) (*types.Pipeline, error) {
	pipeline, resp, err := c.spinnaker.ApplicationControllerApi.GetPipelineConfigUsingGET(c.spinnaker.Context, options.App, name)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("encountered an error get pipeline, status code: %d", resp.StatusCode)
	}

	v := &types.Pipeline{}
	if err := types.ParsePipeline(pipeline, v); err != nil {
		return nil, err
	}

	if options.Expand && v.Template != nil {
		refArr := strings.Split(v.Template.Reference, "://")
		if len(refArr) != 2 {
			return nil, err
		}

		id := refArr[1]
		getOption := &PipelineTemplateGetOptions{}
		pt, err := c.GetPipelineTemplate(id, getOption)
		if err != nil {
			return nil, err
		}

		v.Stages = pt.Pipeline.Stages
	}

	return v, nil
}

// ListPipelines gets all pipeline of the application
func (c *ClientImpl) ListPipelines(options *PipelineListOptions) ([]types.Pipeline, error) {
	var pipelines []interface{}
	var resp *http.Response
	var err error

	if options.All {
		listAppOptions := &AppListOptions{}
		apps, err := c.ListApplications(listAppOptions)
		if err != nil {
			return nil, err
		}

		for _, app := range apps {
			var appPipelines []interface{}
			appPipelines, resp, err = c.spinnaker.ApplicationControllerApi.GetPipelineConfigsForApplicationUsingGET(c.spinnaker.Context, app.Name)
			if err != nil {
				return nil, err
			}

			if resp.StatusCode != http.StatusOK {
				return nil, fmt.Errorf("encountered an error get pipeline, status code: %d, app: %s", resp.StatusCode, app.Name)
			}

			pipelines = append(pipelines, appPipelines...)
		}
	} else {
		pipelines, resp, err = c.spinnaker.ApplicationControllerApi.GetPipelineConfigsForApplicationUsingGET(c.spinnaker.Context, options.App)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("encountered an error get pipeline, status code: %d, app: %s", resp.StatusCode, options.App)
		}
	}

	pipelinesRead := make([]types.Pipeline, len(pipelines))

	for i, pipeline := range pipelines {
		var v types.Pipeline
		if err := types.ParsePipeline(pipeline, &v); err != nil {
			return nil, err
		}

		if options.Expand && v.Template != nil {
			refArr := strings.Split(v.Template.Reference, "://")
			if len(refArr) != 2 {
				return nil, err
			}

			id := refArr[1]
			getOption := &PipelineTemplateGetOptions{}
			pt, err := c.GetPipelineTemplate(id, getOption)
			if err != nil {
				return nil, err
			}

			v.Stages = pt.Pipeline.Stages
		}

		pipelinesRead[i] = v
	}

	return pipelinesRead, nil
}

// GetPipelineTemplate gets a pipeline by ID
func (c *ClientImpl) GetPipelineTemplate(id string, options *PipelineTemplateGetOptions) (*types.PipelineTemplate, error) {
	opts := &gate.V2PipelineTemplatesControllerApiGetUsingGET2Opts{}
	pipeline, resp, err := c.spinnaker.V2PipelineTemplatesControllerApi.GetUsingGET2(c.spinnaker.Context, id, opts)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("encountered an error get pipeline, status code: %d", resp.StatusCode)
	}

	v := &types.PipelineTemplate{}
	if err := types.ParsePipelineTemplate(pipeline, v); err != nil {
		return nil, err
	}

	return v, nil
}

// ListPipelineTemplates gets all pipeline of the application
func (c *ClientImpl) ListPipelineTemplates(options *PipelineTemplateListOptions) ([]types.PipelineTemplate, error) {
	opts := &gate.V2PipelineTemplatesControllerApiListUsingGET1Opts{}
	pts, resp, err := c.spinnaker.V2PipelineTemplatesControllerApi.ListUsingGET1(
		c.spinnaker.Context,
		opts,
	)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("encountered an error get pipeline templates, status code: %d", resp.StatusCode)
	}

	ptsRead := make([]types.PipelineTemplate, len(pts))
	for i, pt := range pts {
		var v types.PipelineTemplate
		if err := types.ParsePipelineTemplate(pt, &v); err != nil {
			return nil, err
		}

		ptsRead[i] = v
	}

	return ptsRead, nil
}
