package provider

import (
	"context"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-community-providers/terraform-plugin-framework-utils/modifiers"
	"github.com/terraform-community-providers/terraform-plugin-framework-utils/validators"
)

var _ resource.Resource = &TeamResource{}
var _ resource.ResourceWithImportState = &TeamResource{}

func NewTeamResource() resource.Resource {
	return &TeamResource{}
}

type TeamResource struct {
	client *graphql.Client
}

type TeamResourceNestedModel struct {
	BoolEmptyDefault types.Bool `tfsdk:"bool_empty_default"`
}

type TeamResourceModel struct {
	Name                        types.String  `tfsdk:"id"`
	BoolEmptyDefault            types.Bool    `tfsdk:"bool_empty_default"`
	BoolKnownDefault            types.Bool    `tfsdk:"bool_known_default"`
	NullableBool                types.Bool    `tfsdk:"nullable_bool"`
	NullableBoolEmptyDefault    types.Bool    `tfsdk:"nullable_bool_empty_default"`
	NullableBoolKnownDefault    types.Bool    `tfsdk:"nullable_bool_known_default"`
	StringEmptyDefault          types.String  `tfsdk:"string_empty_default"`
	StringKnownDefault          types.String  `tfsdk:"string_known_default"`
	StringRandomDefault         types.String  `tfsdk:"string_random_default"`
	NullableString              types.String  `tfsdk:"nullable_string"`
	NullableStringEmptyDefault  types.String  `tfsdk:"nullable_string_empty_default"`
	NullableStringKnownDefault  types.String  `tfsdk:"nullable_string_known_default"`
	NullableStringRandomDefault types.String  `tfsdk:"nullable_string_random_default"`
	FloatEmptyDefault           types.Float64 `tfsdk:"float_empty_default"`
	FloatKnownDefault           types.Float64 `tfsdk:"float_known_default"`
	FloatRandomDefault          types.Float64 `tfsdk:"float_random_default"`
	NullableFloat               types.Float64 `tfsdk:"nullable_float"`
	NullableFloatEmptyDefault   types.Float64 `tfsdk:"nullable_float_empty_default"`
	NullableFloatKnownDefault   types.Float64 `tfsdk:"nullable_float_known_default"`
	NullableFloatRandomDefault  types.Float64 `tfsdk:"nullable_float_random_default"`
	Nested                      types.Object  `tfsdk:"nested"`
}

func (r *TeamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (t *TeamResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		MarkdownDescription: "Test team.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "Name of the team.",
				Type:                types.StringType,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					validators.MinLength(2),
				},
			},
			"bool_empty_default": {
				MarkdownDescription: "Boolean with empty default. **Default** `false`.",
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultBool(false),
				},
			},
			"bool_known_default": {
				MarkdownDescription: "Boolean with known default. **Default** `true`.",
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultBool(true),
				},
			},
			"nullable_bool": {
				MarkdownDescription: "Nullable boolean. **Default** `null`.",
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.NullableBool(),
				},
			},
			"nullable_bool_empty_default": {
				MarkdownDescription: "Nullable boolean with empty default. **Default** `false`.",
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultBool(false),
				},
			},
			"nullable_bool_known_default": {
				MarkdownDescription: "Nullable boolean with known default. **Default** `true`.",
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultBool(true),
				},
			},
			"string_empty_default": {
				MarkdownDescription: "String with empty default. **Default** ``.",
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultString(""),
				},
			},
			"string_known_default": {
				MarkdownDescription: "String with known default. **Default** `One`.",
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultString("One"),
				},
			},
			"string_random_default": {
				MarkdownDescription: "String with random default.",
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"nullable_string": {
				MarkdownDescription: "Nullable string. **Default** `null`.",
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.NullableString(),
				},
			},
			"nullable_string_empty_default": {
				MarkdownDescription: "Nullable string with empty default. **Default** ``.",
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultString(""),
				},
			},
			"nullable_string_known_default": {
				MarkdownDescription: "Nullable string with known default. **Default** `Two`.",
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultString("Two"),
				},
			},
			"nullable_string_random_default": {
				MarkdownDescription: "Nullable string with random default.",
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"float_empty_default": {
				MarkdownDescription: "Float with empty default. **Default** ``.",
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultFloat(0),
				},
			},
			"float_known_default": {
				MarkdownDescription: "Float with known default. **Default** `One`.",
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultFloat(1),
				},
			},
			"float_random_default": {
				MarkdownDescription: "Float with random default.",
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"nullable_float": {
				MarkdownDescription: "Nullable float. **Default** `null`.",
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.NullableFloat(),
				},
			},
			"nullable_float_empty_default": {
				MarkdownDescription: "Nullable float with empty default. **Default** ``.",
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultFloat(0),
				},
			},
			"nullable_float_known_default": {
				MarkdownDescription: "Nullable float with known default. **Default** `Two`.",
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.DefaultFloat(2),
				},
			},
			"nullable_float_random_default": {
				MarkdownDescription: "Nullable float with random default.",
				Type:                types.Float64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"nested": {
				MarkdownDescription: "Nested block.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					modifiers.UnknownAttributesOnUnknown(),
				},
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"bool_empty_default": {
						MarkdownDescription: "Bool with empty default. **Default** `false`.",
						Type:                types.BoolType,
						Optional:            true,
						Computed:            true,
						PlanModifiers: tfsdk.AttributePlanModifiers{
							modifiers.DefaultBool(false),
						},
					},
					// "bool_known_default": {
					// 	MarkdownDescription: "Bool with known default. **Default** `true`.",
					// 	Type:                types.BoolType,
					// 	Optional:            true,
					// 	Computed:            true,
					// 	PlanModifiers: tfsdk.AttributePlanModifiers{
					// 		modifiers.DefaultBool(true),
					// 	},
					// },
				}),
			},
		},
	}, nil
}

func (r *TeamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*graphql.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *graphql.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *TeamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TeamResourceModel
	var nestedData *TeamResourceNestedModel

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	input := TeamsInsertInput{
		Name:               data.Name.Value,
		BoolEmptyDefault:   data.BoolEmptyDefault.Value,
		BoolKnownDefault:   data.BoolKnownDefault.Value,
		StringEmptyDefault: data.StringEmptyDefault.Value,
		StringKnownDefault: data.StringKnownDefault.Value,
		FloatEmptyDefault:  data.FloatEmptyDefault.Value,
		FloatKnownDefault:  data.FloatKnownDefault.Value,
	}

	if !data.NullableBool.IsNull() {
		input.NullableBool = &data.NullableBool.Value
	}

	if !data.NullableBoolEmptyDefault.IsNull() {
		input.NullableBoolEmptyDefault = &data.NullableBoolEmptyDefault.Value
	}

	if !data.NullableBoolKnownDefault.IsNull() {
		input.NullableBoolKnownDefault = &data.NullableBoolKnownDefault.Value
	}

	if !data.StringRandomDefault.IsUnknown() {
		input.StringRandomDefault = &data.StringRandomDefault.Value
	}

	if !data.NullableString.IsNull() {
		input.NullableString = &data.NullableString.Value
	}

	if !data.NullableStringEmptyDefault.IsNull() {
		input.NullableStringEmptyDefault = &data.NullableStringEmptyDefault.Value
	}

	if !data.NullableStringKnownDefault.IsNull() {
		input.NullableStringKnownDefault = &data.NullableStringKnownDefault.Value
	}

	if !data.NullableStringRandomDefault.IsUnknown() {
		input.NullableStringRandomDefault = &data.NullableStringRandomDefault.Value
	}

	if !data.FloatRandomDefault.IsUnknown() {
		input.FloatRandomDefault = &data.FloatRandomDefault.Value
	}

	if !data.NullableFloat.IsNull() {
		input.NullableFloat = &data.NullableFloat.Value
	}

	if !data.NullableFloatEmptyDefault.IsNull() {
		input.NullableFloatEmptyDefault = &data.NullableFloatEmptyDefault.Value
	}

	if !data.NullableFloatKnownDefault.IsNull() {
		input.NullableFloatKnownDefault = &data.NullableFloatKnownDefault.Value
	}

	if !data.NullableFloatRandomDefault.IsUnknown() {
		input.NullableFloatRandomDefault = &data.NullableFloatRandomDefault.Value
	}

	diags = data.Nested.As(ctx, &nestedData, types.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	input.NestedBoolEmptyDefault = nestedData.BoolEmptyDefault.Value

	response, err := createTeam(context.Background(), *r.client, input)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create team, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "created a team")

	team := response.InsertIntoTeamsCollection.Records[0]

	data.BoolEmptyDefault = types.Bool{Value: team.BoolEmptyDefault}
	data.BoolKnownDefault = types.Bool{Value: team.BoolKnownDefault}
	data.StringEmptyDefault = types.String{Value: team.StringEmptyDefault}
	data.StringKnownDefault = types.String{Value: team.StringKnownDefault}
	data.StringRandomDefault = types.String{Value: team.StringRandomDefault}
	data.FloatEmptyDefault = types.Float64{Value: float64(team.FloatEmptyDefault)}
	data.FloatKnownDefault = types.Float64{Value: float64(team.FloatKnownDefault)}
	data.FloatRandomDefault = types.Float64{Value: float64(team.FloatRandomDefault)}

	if team.NullableBool != nil {
		data.NullableBool = types.Bool{Value: *team.NullableBool}
	}

	if team.NullableBoolEmptyDefault != nil {
		data.NullableBoolEmptyDefault = types.Bool{Value: *team.NullableBoolEmptyDefault}
	}

	if team.NullableBoolKnownDefault != nil {
		data.NullableBoolKnownDefault = types.Bool{Value: *team.NullableBoolKnownDefault}
	}

	if team.NullableString != nil {
		data.NullableString = types.String{Value: *team.NullableString}
	}

	if team.NullableStringEmptyDefault != nil {
		data.NullableStringEmptyDefault = types.String{Value: *team.NullableStringEmptyDefault}
	}

	if team.NullableStringKnownDefault != nil {
		data.NullableStringKnownDefault = types.String{Value: *team.NullableStringKnownDefault}
	}

	if team.NullableStringRandomDefault != nil {
		data.NullableStringRandomDefault = types.String{Value: *team.NullableStringRandomDefault}
	}

	if team.NullableFloat != nil {
		data.NullableFloat = types.Float64{Value: float64(*team.NullableFloat)}
	}

	if team.NullableFloatEmptyDefault != nil {
		data.NullableFloatEmptyDefault = types.Float64{Value: float64(*team.NullableFloatEmptyDefault)}
	}

	if team.NullableFloatKnownDefault != nil {
		data.NullableFloatKnownDefault = types.Float64{Value: float64(*team.NullableFloatKnownDefault)}
	}

	if team.NullableFloatRandomDefault != nil {
		data.NullableFloatRandomDefault = types.Float64{Value: float64(*team.NullableFloatRandomDefault)}
	}

	data.Nested = types.Object{
		AttrTypes: map[string]attr.Type{
			"bool_empty_default": types.BoolType,
		},
		Attrs: map[string]attr.Value{
			"bool_empty_default": types.Bool{Value: team.NestedBoolEmptyDefault},
		},
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *TeamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TeamResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	response, err := getTeam(context.Background(), *r.client, data.Name.Value)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read team, got error: %s", err))
		return
	}

	team := response.TeamsCollection.Edges[0].Node

	data.Name = types.String{Value: team.Name}

	data.BoolEmptyDefault = types.Bool{Value: team.BoolEmptyDefault}
	data.BoolKnownDefault = types.Bool{Value: team.BoolKnownDefault}
	data.StringEmptyDefault = types.String{Value: team.StringEmptyDefault}
	data.StringKnownDefault = types.String{Value: team.StringKnownDefault}
	data.StringRandomDefault = types.String{Value: team.StringRandomDefault}
	data.FloatEmptyDefault = types.Float64{Value: float64(team.FloatEmptyDefault)}
	data.FloatKnownDefault = types.Float64{Value: float64(team.FloatKnownDefault)}
	data.FloatRandomDefault = types.Float64{Value: float64(team.FloatRandomDefault)}

	if team.NullableBool != nil {
		data.NullableBool = types.Bool{Value: *team.NullableBool}
	}

	if team.NullableBoolEmptyDefault != nil {
		data.NullableBoolEmptyDefault = types.Bool{Value: *team.NullableBoolEmptyDefault}
	}

	if team.NullableBoolKnownDefault != nil {
		data.NullableBoolKnownDefault = types.Bool{Value: *team.NullableBoolKnownDefault}
	}

	if team.NullableString != nil {
		data.NullableString = types.String{Value: *team.NullableString}
	}

	if team.NullableStringEmptyDefault != nil {
		data.NullableStringEmptyDefault = types.String{Value: *team.NullableStringEmptyDefault}
	}

	if team.NullableStringKnownDefault != nil {
		data.NullableStringKnownDefault = types.String{Value: *team.NullableStringKnownDefault}
	}

	if team.NullableStringRandomDefault != nil {
		data.NullableStringRandomDefault = types.String{Value: *team.NullableStringRandomDefault}
	}

	if team.NullableFloat != nil {
		data.NullableFloat = types.Float64{Value: float64(*team.NullableFloat)}
	}

	if team.NullableFloatEmptyDefault != nil {
		data.NullableFloatEmptyDefault = types.Float64{Value: float64(*team.NullableFloatEmptyDefault)}
	}

	if team.NullableFloatKnownDefault != nil {
		data.NullableFloatKnownDefault = types.Float64{Value: float64(*team.NullableFloatKnownDefault)}
	}

	if team.NullableFloatRandomDefault != nil {
		data.NullableFloatRandomDefault = types.Float64{Value: float64(*team.NullableFloatRandomDefault)}
	}

	data.Nested = types.Object{
		AttrTypes: map[string]attr.Type{
			"bool_empty_default": types.BoolType,
		},
		Attrs: map[string]attr.Value{
			"bool_empty_default": types.Bool{Value: team.NestedBoolEmptyDefault},
		},
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TeamResourceModel
	var nestedData *TeamResourceNestedModel
	var state *TeamResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	input := TeamsUpdateInput{
		BoolEmptyDefault:   data.BoolEmptyDefault.Value,
		BoolKnownDefault:   data.BoolKnownDefault.Value,
		StringEmptyDefault: data.StringEmptyDefault.Value,
		StringKnownDefault: data.StringKnownDefault.Value,
		FloatEmptyDefault:  data.FloatEmptyDefault.Value,
		FloatKnownDefault:  data.FloatKnownDefault.Value,
	}

	if data.Name.Value != state.Name.Value {
		input.Name = data.Name.Value
	}

	if !data.NullableBool.IsNull() {
		input.NullableBool = &data.NullableBool.Value
	}

	if !data.NullableBoolEmptyDefault.IsNull() {
		input.NullableBoolEmptyDefault = &data.NullableBoolEmptyDefault.Value
	}

	if !data.NullableBoolKnownDefault.IsNull() {
		input.NullableBoolKnownDefault = &data.NullableBoolKnownDefault.Value
	}

	if !data.StringRandomDefault.IsNull() {
		input.StringRandomDefault = &data.StringRandomDefault.Value
	}

	if !data.NullableString.IsNull() {
		input.NullableString = &data.NullableString.Value
	}

	if !data.NullableStringEmptyDefault.IsNull() {
		input.NullableStringEmptyDefault = &data.NullableStringEmptyDefault.Value
	}

	if !data.NullableStringKnownDefault.IsNull() {
		input.NullableStringKnownDefault = &data.NullableStringKnownDefault.Value
	}

	if !data.NullableStringRandomDefault.IsNull() {
		input.NullableStringRandomDefault = &data.NullableStringRandomDefault.Value
	}

	if !data.FloatRandomDefault.IsNull() {
		input.FloatRandomDefault = &data.FloatRandomDefault.Value
	}

	if !data.NullableFloat.IsNull() {
		input.NullableFloat = &data.NullableFloat.Value
	}

	if !data.NullableFloatEmptyDefault.IsNull() {
		input.NullableFloatEmptyDefault = &data.NullableFloatEmptyDefault.Value
	}

	if !data.NullableFloatKnownDefault.IsNull() {
		input.NullableFloatKnownDefault = &data.NullableFloatKnownDefault.Value
	}

	if !data.NullableFloatRandomDefault.IsNull() {
		input.NullableFloatRandomDefault = &data.NullableFloatRandomDefault.Value
	}

	resp.Diagnostics.Append(data.Nested.As(ctx, &nestedData, types.ObjectAsOptions{})...)

	if resp.Diagnostics.HasError() {
		return
	}

	input.NestedBoolEmptyDefault = nestedData.BoolEmptyDefault.Value

	response, err := updateTeam(context.Background(), *r.client, input, state.Name.Value)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update team, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated a team")

	team := response.UpdateTeamsCollection.Records[0]

	data.BoolEmptyDefault = types.Bool{Value: team.BoolEmptyDefault}
	data.BoolKnownDefault = types.Bool{Value: team.BoolKnownDefault}
	data.StringEmptyDefault = types.String{Value: team.StringEmptyDefault}
	data.StringKnownDefault = types.String{Value: team.StringKnownDefault}
	data.StringRandomDefault = types.String{Value: team.StringRandomDefault}
	data.FloatEmptyDefault = types.Float64{Value: float64(team.FloatEmptyDefault)}
	data.FloatKnownDefault = types.Float64{Value: float64(team.FloatKnownDefault)}
	data.FloatRandomDefault = types.Float64{Value: float64(team.FloatRandomDefault)}

	if team.NullableBool != nil {
		data.NullableBool = types.Bool{Value: *team.NullableBool}
	}

	if team.NullableBoolEmptyDefault != nil {
		data.NullableBoolEmptyDefault = types.Bool{Value: *team.NullableBoolEmptyDefault}
	}

	if team.NullableBoolKnownDefault != nil {
		data.NullableBoolKnownDefault = types.Bool{Value: *team.NullableBoolKnownDefault}
	}

	if team.NullableString != nil {
		data.NullableString = types.String{Value: *team.NullableString}
	}

	if team.NullableStringEmptyDefault != nil {
		data.NullableStringEmptyDefault = types.String{Value: *team.NullableStringEmptyDefault}
	}

	if team.NullableStringKnownDefault != nil {
		data.NullableStringKnownDefault = types.String{Value: *team.NullableStringKnownDefault}
	}

	if team.NullableStringRandomDefault != nil {
		data.NullableStringRandomDefault = types.String{Value: *team.NullableStringRandomDefault}
	}

	if team.NullableFloat != nil {
		data.NullableFloat = types.Float64{Value: float64(*team.NullableFloat)}
	}

	if team.NullableFloatEmptyDefault != nil {
		data.NullableFloatEmptyDefault = types.Float64{Value: float64(*team.NullableFloatEmptyDefault)}
	}

	if team.NullableFloatKnownDefault != nil {
		data.NullableFloatKnownDefault = types.Float64{Value: float64(*team.NullableFloatKnownDefault)}
	}

	if team.NullableFloatRandomDefault != nil {
		data.NullableFloatRandomDefault = types.Float64{Value: float64(*team.NullableFloatRandomDefault)}
	}

	data.Nested = types.Object{
		AttrTypes: map[string]attr.Type{
			"bool_empty_default": types.BoolType,
		},
		Attrs: map[string]attr.Value{
			"bool_empty_default": types.Bool{Value: team.NestedBoolEmptyDefault},
		},
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *TeamResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := deleteTeam(context.Background(), *r.client, data.Name.Value)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete team, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "deleted a team")
}

func (r *TeamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
