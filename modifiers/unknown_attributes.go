package modifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func UnknownAttributesOnUnknown() unknownAttributesOnUnknownModifier {
	return unknownAttributesOnUnknownModifier{}
}

type unknownAttributesOnUnknownModifier struct{}

func (m unknownAttributesOnUnknownModifier) Description(ctx context.Context) string {
	return "If value is unknown, defaults to an object with all attributes set to unknown."
}

func (m unknownAttributesOnUnknownModifier) MarkdownDescription(ctx context.Context) string {
	return "If value is unknown, defaults to an object with all attributes set to unknown."
}

func (m unknownAttributesOnUnknownModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if !req.AttributePlan.IsUnknown() {
		return
	}

	var object types.Object

	diags := tfsdk.ValueAs(ctx, req.AttributePlan, &object)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	object.Attrs = make(map[string]attr.Value, len(object.AttrTypes))
	object.Unknown = false

	for name, attrType := range object.AttrTypes {
		attrValue, err := attrType.ValueFromTerraform(ctx, tftypes.NewValue(attrType.TerraformType(ctx), tftypes.UnknownValue))

		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.AttributePath.AtName(name),
				"Attribute Plan Modification Error",
				"While creating unknown values for object attributes, an unexpected error occurred. Please report the following to the provider developers.\n\n"+
					"Error: "+err.Error(),
			)
		}

		object.Attrs[name] = attrValue
	}

	resp.AttributePlan = object
}
