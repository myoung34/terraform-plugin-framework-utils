package validators

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	matchErr = "value must match regex %s"
)

type matchValidator struct {
	regex *regexp.Regexp
}

func Match(regex *regexp.Regexp) tfsdk.AttributeValidator {
	return matchValidator{
		regex: regex,
	}
}

func (v matchValidator) Description(context.Context) string {
	return "Ensure that the attribute value matches the provided regex."
}

func (v matchValidator) MarkdownDescription(context.Context) string {
	return "Ensure that the attribute value matches the provided regex."
}

func (v matchValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var str types.String
	{
		diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &str)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			return
		}
	}

	if str.Unknown || str.Null {
		return
	}

	if v.regex != nil && !v.regex.MatchString(str.Value) {
		resp.Diagnostics.AddAttributeError(
			req.AttributePath,
			"Invalid String",
			fmt.Sprintf(matchErr, v.regex.String()),
		)
	}
}
