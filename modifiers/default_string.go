package modifiers

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DefaultString(def string) defaultStringModifier {
	return defaultStringModifier{Default: &def}
}

func NullableString() defaultStringModifier {
	return defaultStringModifier{}
}

type defaultStringModifier struct {
	Default *string
}

func (m defaultStringModifier) String() defaultStringModifier {
	defaultStr := "null"
	if m.Default == nil {
		return defaultStringModifier{
			Default: &defaultStr,
		}
	}
	return defaultStringModifier{
		Default: m.Default,
	}
}

func (m defaultStringModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultStringModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultStringModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	var str types.String
	diags := tfsdk.ValueAs(ctx, req.PlanValue, &str)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if !str.IsNull() {
		return
	}

	resp.PlanValue = types.StringValue(*m.Default)
}
