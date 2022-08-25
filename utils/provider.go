package utils

// `ConvertProviderType` is a helper function for `NewResource` and `NewDataSource`
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: `in.(*provider)`), however using this can prevent
// potential panics.
// func ConvertProviderType[T provider.Provider](in provider.Provider) (T, diag.Diagnostics) {
// 	var diags diag.Diagnostics

// 	p, ok := in.(*T)

// 	if !ok {
// 		diags.AddError(
// 			"Unexpected Provider Instance Type",
// 			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
// 		)
// 		return T{}, diags
// 	}

// 	if p == nil {
// 		diags.AddError(
// 			"Unexpected Provider Instance Type",
// 			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
// 		)
// 		return T{}, diags
// 	}

// 	return *p, diags
// }
