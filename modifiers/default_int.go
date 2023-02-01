package modifiers

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DefaultInt(def int64) defaultIntModifier {
	return defaultIntModifier{Default: &def}
}

func NullableInt() defaultIntModifier {
	return defaultIntModifier{}
}

type defaultIntModifier struct {
	Default *int64
}

func (m defaultIntModifier) String() string {
	if m.Default == nil {
		return "null"
	}

	return fmt.Sprintf("%f", *m.Default)
}

func (m defaultIntModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultIntModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultIntModifier) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	var str types.Int64
	diags := tfsdk.ValueAs(ctx, req.PlanValue, &str)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if !str.IsNull() {
		return
	}

	resp.PlanValue = types.Int64Value(*m.Default)
}
