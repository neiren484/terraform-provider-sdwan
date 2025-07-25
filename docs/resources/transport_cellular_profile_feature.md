---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_transport_cellular_profile_feature Resource - terraform-provider-sdwan"
subcategory: "Features - Transport"
description: |-
  This resource can manage a Transport Cellular Profile Feature.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_transport_cellular_profile_feature (Resource)

This resource can manage a Transport Cellular Profile Feature.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_transport_cellular_profile_feature" "example" {
  name                     = "Example"
  description              = "My Example"
  feature_profile_id       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  profile_id               = 1
  access_point_name        = "apn1"
  requires_authentication  = true
  authentication_type      = "pap"
  profile_username         = "example"
  profile_password         = "example123!"
  packet_data_network_type = "ipv4"
  no_overwrite             = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the Feature

### Optional

- `access_point_name` (String) Set access point name
- `access_point_name_variable` (String) Variable name
- `authentication_type` (String) Set authentication type, Attribute conditional on `requires_authentication` being equal to `true`
  - Choices: `pap`, `chap`, `pap_chap`
- `authentication_type_variable` (String) Variable name, Attribute conditional on `requires_authentication` being equal to `true`
- `description` (String) The description of the Feature
- `no_overwrite` (Boolean) No Overwrite
- `no_overwrite_variable` (String) Variable name
- `packet_data_network_type` (String) Set packet data network type
  - Choices: `ipv4`, `ipv4v6`, `ipv6`
  - Default value: `ipv4`
- `packet_data_network_type_variable` (String) Variable name
- `profile_id` (Number) Set Profile ID
  - Range: `1`-`16`
- `profile_id_variable` (String) Variable name
- `profile_password` (String) Set the profile password, Attribute conditional on `requires_authentication` being equal to `true`
- `profile_password_variable` (String) Variable name, Attribute conditional on `requires_authentication` being equal to `true`
- `profile_username` (String) Set the profile username, Attribute conditional on `requires_authentication` being equal to `true`
- `profile_username_variable` (String) Variable name, Attribute conditional on `requires_authentication` being equal to `true`
- `requires_authentication` (Boolean) Require authentication type
  - Default value: `false`

### Read-Only

- `id` (String) The id of the Feature
- `version` (Number) The version of the Feature

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Expected import identifier with the format: "transport_cellular_profile_feature_id,feature_profile_id"
terraform import sdwan_transport_cellular_profile_feature.example "f6b2c44c-693c-4763-b010-895aa3d236bd,f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
```
