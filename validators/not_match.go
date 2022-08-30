package validators

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	notMatchErr = "value must not match regex %s"
)

type notMatchValidator struct {
	regex *regexp.Regexp
}

func NotMatch(regex *regexp.Regexp) tfsdk.AttributeValidator {
	return notMatchValidator{
		regex: regex,
	}
}

func (v notMatchValidator) Description(context.Context) string {
	return "Ensure that the attribute value does not match the provided regex."
}

func (v notMatchValidator) MarkdownDescription(context.Context) string {
	return "Ensure that the attribute value does not match the provided regex."
}

func (v notMatchValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
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

	if v.regex != nil && v.regex.MatchString(str.Value) {
		resp.Diagnostics.AddAttributeError(
			req.AttributePath,
			"Invalid String",
			fmt.Sprintf(notMatchErr, v.regex.String()),
		)
	}
}
