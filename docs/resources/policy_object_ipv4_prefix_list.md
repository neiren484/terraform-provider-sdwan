---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_policy_object_ipv4_prefix_list Resource - terraform-provider-sdwan"
subcategory: "Policy Objects"
description: |-
  This resource can manage a Policy Object IPv4 Prefix List Policy_object.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_policy_object_ipv4_prefix_list (Resource)

This resource can manage a Policy Object IPv4 Prefix List Policy_object.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_policy_object_ipv4_prefix_list" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  entries = [
    {
      ipv4_address       = "10.0.0.0"
      ipv4_prefix_length = 8
      le                 = 24
      ge                 = 16
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entries` (Attributes List) IPv4 Prefix List (see [below for nested schema](#nestedatt--entries))
- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the Policy_object

### Optional

- `description` (String) The description of the Policy_object

### Read-Only

- `id` (String) The id of the Policy_object
- `version` (Number) The version of the Policy_object

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Optional:

- `ge` (Number) IPv4 prefix length with ge range operator
  - Range: `1`-`32`
- `ipv4_address` (String) IPv4 address
- `ipv4_prefix_length` (Number) IPv4 prefix length
  - Range: `0`-`32`
- `le` (Number) IPv4 prefix length with le range operator
  - Range: `1`-`32`

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Expected import identifier with the format: "policy_object_ipv4_prefix_list_id,feature_profile_id"
terraform import sdwan_policy_object_ipv4_prefix_list.example "f6b2c44c-693c-4763-b010-895aa3d236bd,f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
```
