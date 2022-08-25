package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	envVarName          = "TEST_TOKEN"
	errMissingAuthToken = "Required token could not be found. Please set the token using an input variable in the provider configuration block or by using the `" + envVarName + "` environment variable."
)

// Ensure provider defined types fully satisfy framework interfaces
var _ provider.Provider = &testProvider{}

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type testProvider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	client graphql.Client

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Token types.String `tfsdk:"token"`
}

type authedTransport struct {
	token   string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("apikey", t.token)
	return t.wrapped.RoundTrip(req)
}

func (p *testProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data providerData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	token := ""

	if !data.Token.IsNull() {
		token = data.Token.Value
	}

	// If a token wasn't set in the provider configuration block, try and fetch it
	// from the environment variable.
	if token == "" {
		token = os.Getenv(envVarName)
	}

	// If we still don't have a token at this point, we return an error.
	if token == "" {
		resp.Diagnostics.AddError("Missing API token", errMissingAuthToken)
		return
	}

	httpClient := http.Client{
		Transport: &authedTransport{
			token:   token,
			wrapped: http.DefaultTransport,
		},
	}

	p.client = graphql.NewClient("https://tkmvizljskjrpnrapgnr.supabase.co/graphql/v1", &httpClient)
	p.configured = true
}

func (p *testProvider) GetResources(ctx context.Context) (map[string]provider.ResourceType, diag.Diagnostics) {
	return map[string]provider.ResourceType{
		"test_team": teamResourceType{},
	}, nil
}

func (p *testProvider) GetDataSources(ctx context.Context) (map[string]provider.DataSourceType, diag.Diagnostics) {
	return map[string]provider.DataSourceType{}, nil
}

func (p *testProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"token": {
				MarkdownDescription: "The token used to authenticate with Test.",
				Optional:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &testProvider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in provider.Provider) (testProvider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*testProvider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return testProvider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return testProvider{}, diags
	}

	return *p, diags
}
