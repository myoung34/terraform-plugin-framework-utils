# Terraform Plugin Framework Utilities

### Modifiers

Assuming that we have a field of a certain type (`int`, `boolean`, `string` etc..), that field can either be nullable or not and also can have various defaults. We need our modifiers to work with all these scenarios.

| nullable |    null default    |   empty default[^1]   |     known default     |        random default        |
| :------: | :----------------: | :-------------------: | :-------------------: | :--------------------------: |
|    no    |       X[^2]        |   [`DefaultType`][]   |   [`DefaultType`][]   |   [`UseStateForUnknown`][]   |
|   yes    | [`NullableType`][] | [`DefaultType`][][^3] | [`DefaultType`][][^3] | [`UseStateForUnknown`][][^3] |

[^1]: _empty default_ means that the default value of the field on the server is the empty value for that type in golang. e.g. boolean `false`, string `""`, int `0` etc..
[^2]: This scenairo is impossible.
[^3]: End users will not be able to set the value of the field as `null` in the server. This is a limitation on terraform itself.

[`DefaultType`]: #defaulttype
[`NullableType`]: #nullabletype
[`UseStateForUnknown`]: #usestateforunknown

#### DefaultType

Use the appropriate function for type whose name starts with `Default`.

```go
modifiers.DefaultBool(true)

modifiers.DefaultString("")
```

#### NullableType

Use the appropriate function for type whose name starts with `Nullable`.

```go
modifiers.NullableBool()

modifiers.NullableString()
```

#### UseStateForUnknown

Use `UseStateForUnknown` from Terraform plugin framework.

```go
resource.UseStateForUnknown()
```

**NOTE**: Make sure that you omit this property when sending payloads during both creation and updation.

### Validators

#### FloatInSlice

Validates that the number is one of certain values.

```go
validators.FloatInSlice(1, 4, 6)
```

#### StringInSlice

Validates that the string is one of certain values.

```go
validators.StringInSlice(true, "one", "two", "three")

validators.StringInSlice(false, "OnE", "tWo", "tHrEe")
```

#### Match

Validates that the string matches a certain regex.

```go
validators.Match(regexp.MustCompile("^[0-9a-fA-F]{6}$"))
```

#### MinLength

Validates that the string's length is at least of a certain value.

```go
validators.MinLength(1)
```

#### MaxLength

Validates that the string's length is at most of a certain value.

```go
validators.MaxLength(5)
```
