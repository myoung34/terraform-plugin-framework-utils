package modifiers

import (
	"context"
	"fmt"

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

func (m defaultFloatModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if !req.AttributeConfig.IsNull() {
		return
	}

	if m.Default == nil {
		resp.AttributePlan = types.Float64{Null: true}
	} else {
		resp.AttributePlan = types.Float64{Value: *m.Default}
	}
}
