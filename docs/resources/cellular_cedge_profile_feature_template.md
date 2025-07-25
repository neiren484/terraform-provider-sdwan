---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cellular_cedge_profile_feature_template Resource - terraform-provider-sdwan"
subcategory: "(Classic) Feature Templates"
description: |-
  This resource can manage a Cellular cEdge Profile feature template.
  Minimum SD-WAN Manager version: 15.0.0
---

# sdwan_cellular_cedge_profile_feature_template (Resource)

This resource can manage a Cellular cEdge Profile feature template.
  - Minimum SD-WAN Manager version: `15.0.0`

## Example Usage

```terraform
resource "sdwan_cellular_cedge_profile_feature_template" "example" {
  name                     = "Example"
  description              = "My Example"
  device_types             = ["vedge-C8000V"]
  profile_id               = 1
  access_point_name        = "APN1"
  authentication_type      = "chap"
  packet_data_network_type = "ipv4"
  profile_username         = "MyUsername"
  profile_password         = "MyPassword"
  no_overwrite             = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the feature template
- `device_types` (Set of String) List of supported device types
  - Choices: `vedge-C8000V`, `vedge-C8300-1N1S-4T2X`, `vedge-C8300-1N1S-6T`, `vedge-C8300-2N2S-6T`, `vedge-C8300-2N2S-4T2X`, `vedge-C8500-12X4QC`, `vedge-C8500-12X`, `vedge-C8500-20X6C`, `vedge-C8500L-8S4X`, `vedge-C8200-1N-4T`, `vedge-C8200L-1N-4T`
- `name` (String) The name of the feature template

### Optional

- `access_point_name` (String) Set access point name
- `access_point_name_variable` (String) Variable name
- `authentication_type` (String) Set authentication type
  - Choices: `none`, `pap`, `chap`, `pap_chap`
  - Default value: `none`
- `authentication_type_variable` (String) Variable name
- `no_overwrite` (Boolean) No Overwrite
- `no_overwrite_variable` (String) Variable name
- `packet_data_network_type` (String) Set packet data network type
  - Choices: `ipv4`, `ipv4v6`, `ipv6`
  - Default value: `ipv4`
- `packet_data_network_type_variable` (String) Variable name
- `profile_id` (Number) Set Profile ID
  - Range: `1`-`16`
- `profile_id_variable` (String) Variable name
- `profile_password` (String) Set the profile password
- `profile_password_variable` (String) Variable name
- `profile_username` (String) Set the profile username
- `profile_username_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the feature template
- `template_type` (String) The template type
- `version` (Number) The version of the feature template

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import sdwan_cellular_cedge_profile_feature_template.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
