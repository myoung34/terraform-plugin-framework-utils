package modifiers

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DefaultFloat(def float64) defaultFloatModifier {
	return defaultFloatModifier{Default: &def}
}

func NullableFloat() defaultFloatModifier {
	return defaultFloatModifier{}
}

type defaultFloatModifier struct {
	Default *float64
}

func (m defaultFloatModifier) String() string {
	if m.Default == nil {
		return "null"
	}

	return fmt.Sprintf("%f", *m.Default)
}

func (m defaultFloatModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultFloatModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultFloatModifier) PlanModifyFloat(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	var str types.Float64
	diags := tfsdk.ValueAs(ctx, req.PlanValue, &str)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if !str.IsNull() {
		return
	}

	resp.PlanValue = types.Float64Value(*m.Default)
}
