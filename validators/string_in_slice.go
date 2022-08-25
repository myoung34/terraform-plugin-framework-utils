package validators

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	stringInSliceErr = "value must be one of [%s]"
)

type stringInSliceValidator struct {
	caseSensitive bool
	values        []string
}

func StringInSlice(caseSensitive bool, values ...string) tfsdk.AttributeValidator {
	return stringInSliceValidator{
		caseSensitive: caseSensitive,
		values:        values,
	}
}

func (v stringInSliceValidator) Description(context.Context) string {
	return "Ensure that the attribute value is one of the provided values."
}

func (v stringInSliceValidator) MarkdownDescription(context.Context) string {
	return "Ensure that the attribute value is one of the provided values."
}

func (v stringInSliceValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
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

	for _, val := range v.values {
		if val == str.Value {
			return
		}
	}

	resp.Diagnostics.AddAttributeError(
		req.AttributePath,
		"Invalid Value",
		fmt.Sprintf(stringInSliceErr, strings.Join(v.values, ", ")),
	)
}
