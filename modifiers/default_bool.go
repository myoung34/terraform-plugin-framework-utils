package modifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DefaultBool(def bool) defaultBoolModifier {
	return defaultBoolModifier{Default: &def}
}

func NullableBool() defaultBoolModifier {
	return defaultBoolModifier{}
}

type defaultBoolModifier struct {
	Default *bool
}

func (m defaultBoolModifier) String() string {
	if m.Default == nil {
		return "null"
	}

	return fmt.Sprintf("%t", *m.Default)
}

func (m defaultBoolModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultBoolModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultBoolModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if !req.AttributeConfig.IsNull() {
		return
	}

	if m.Default == nil {
		resp.AttributePlan = types.Bool{Null: true}
	} else {
		resp.AttributePlan = types.Bool{Value: *m.Default}
	}
}
