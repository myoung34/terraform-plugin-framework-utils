package modifiers

import (
	"context"
	"fmt"

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

func (m defaultStringModifier) String() string {
	if m.Default == nil {
		return "null"
	}

	return *m.Default
}

func (m defaultStringModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultStringModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultStringModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if !req.AttributeConfig.IsNull() {
		return
	}

	if m.Default == nil {
		resp.AttributePlan = types.String{Null: true}
	} else {
		resp.AttributePlan = types.String{Value: *m.Default}
	}
}
