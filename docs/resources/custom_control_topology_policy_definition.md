---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_custom_control_topology_policy_definition Resource - terraform-provider-sdwan"
subcategory: "(Classic) Centralized Policies"
description: |-
  This resource can manage a Custom Control Topology Policy Definition .
---

# sdwan_custom_control_topology_policy_definition (Resource)

This resource can manage a Custom Control Topology Policy Definition .

## Example Usage

```terraform
resource "sdwan_custom_control_topology_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  default_action = "reject"
  sequences = [
    {
      id          = 1
      name        = "Region1"
      type        = "route"
      ip_type     = "ipv4"
      base_action = "accept"
      match_entries = [
        {
          type    = "ompTag"
          omp_tag = 100
        }
      ]
      action_entries = [
        {
          type = "set"
          set_parameters = [
            {
              type       = "preference"
              preference = 100
            }
          ]
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the policy definition
- `name` (String) The name of the policy definition

### Optional

- `default_action` (String) Default action, either `accept` or `reject`
  - Choices: `accept`, `reject`
- `sequences` (Attributes List) List of sequences (see [below for nested schema](#nestedatt--sequences))

### Read-Only

- `id` (String) The id of the object
- `type` (String) Type
- `version` (Number) The version of the object

<a id="nestedatt--sequences"></a>
### Nested Schema for `sequences`

Required:

- `id` (Number) Sequence ID
- `name` (String) Sequence name

Optional:

- `action_entries` (Attributes List) List of action entries (see [below for nested schema](#nestedatt--sequences--action_entries))
- `base_action` (String) Base action, either `accept` or `reject`
  - Choices: `accept`, `reject`
- `ip_type` (String) Sequence IP type, either `ipv4`, `ipv6` or `all`
  - Choices: `ipv4`, `ipv6`, `all`
- `match_entries` (Attributes List) List of match entries (see [below for nested schema](#nestedatt--sequences--match_entries))
- `type` (String) Sequence type, either `route` or `tloc`
  - Choices: `route`, `tloc`

<a id="nestedatt--sequences--action_entries"></a>
### Nested Schema for `sequences.action_entries`

Required:

- `type` (String) Type of action entry
  - Choices: `set`, `exportTo`

Optional:

- `export_to_vpn_list_id` (String) Export to VPN list ID, Attribute conditional on `type` being equal to `exportTo`
- `export_to_vpn_list_version` (Number) Export to VPN list version
- `set_parameters` (Attributes List) List of set parameters, Attribute conditional on `type` being equal to `set` (see [below for nested schema](#nestedatt--sequences--action_entries--set_parameters))

<a id="nestedatt--sequences--action_entries--set_parameters"></a>
### Nested Schema for `sequences.action_entries.set_parameters`

Required:

- `type` (String) Type of set parameter
  - Choices: `tlocList`, `tloc`, `tlocAction`, `preference`, `ompTag`, `community`, `communityAdditive`, `service`

Optional:

- `community` (String) Community value, e.g. `1000:10000` or `internet` or `local-AS`, Attribute conditional on `type` being equal to `community`
- `community_additive` (Boolean) Community additive, Attribute conditional on `type` being equal to `communityAdditive`
- `omp_tag` (Number) OMP tag, Attribute conditional on `type` being equal to `ompTag`
  - Range: `0`-`4294967295`
- `preference` (Number) Preference, Attribute conditional on `type` being equal to `preference`
  - Range: `0`-`4294967295`
- `service_tloc_color` (String) Service TLOC color, Attribute conditional on `type` being equal to `service`
- `service_tloc_encapsulation` (String) Service TLOC encapsulation, Attribute conditional on `type` being equal to `service`
  - Choices: `ipsec`, `gre`
- `service_tloc_ip` (String) Service TLOC IP address, Attribute conditional on `type` being equal to `service`
- `service_tloc_list_id` (String) Service TLOC list ID, Attribute conditional on `type` being equal to `service`
- `service_tloc_list_version` (Number) Service TLOC list version
- `service_type` (String) Service type, Attribute conditional on `type` being equal to `service`
  - Choices: `FW`, `IDP`, `IDS`, `netsvc1`, `netsvc2`, `netsvc3`, `netsvc4`, `netsvc5`
- `service_vpn_id` (Number) Service VPN ID, Attribute conditional on `type` being equal to `service`
  - Range: `0`-`65536`
- `tloc_action` (String) TLOC action, Attribute conditional on `type` being equal to `tlocAction`
  - Choices: `strict`, `primary`, `backup`, `ecmp`
- `tloc_color` (String) TLOC color, Attribute conditional on `type` being equal to `tloc`
- `tloc_encapsulation` (String) TLOC encapsulation, Attribute conditional on `type` being equal to `tloc`
  - Choices: `ipsec`, `gre`
- `tloc_ip` (String) TLOC IP address, Attribute conditional on `type` being equal to `tloc`
- `tloc_list_id` (String) TLOC list ID, Attribute conditional on `type` being equal to `tlocList`
- `tloc_list_version` (Number) TLOC list version



<a id="nestedatt--sequences--match_entries"></a>
### Nested Schema for `sequences.match_entries`

Required:

- `type` (String) Type of match entry
  - Choices: `colorList`, `community`, `expandedCommunity`, `ompTag`, `origin`, `originator`, `preference`, `siteList`, `pathType`, `tlocList`, `vpnList`, `prefixList`, `vpn`, `tloc`, `siteId`, `carrier`, `domainId`, `groupId`

Optional:

- `carrier` (String) Carrier, Attribute conditional on `type` being equal to `carrier`
  - Choices: `default`, `carrier1`, `carrier2`, `carrier3`, `carrier4`, `carrier5`, `carrier6`, `carrier7`, `carrier8`
- `color_list_id` (String) Color list ID, Attribute conditional on `type` being equal to `colorList`
- `color_list_version` (Number) Color list version
- `community_list_id` (String) Community list ID, Attribute conditional on `type` being equal to `community`
- `community_list_version` (Number) Community list version
- `domain_id` (Number) Domain ID, Attribute conditional on `type` being equal to `domainId`
  - Range: `0`-`4294967295`
- `expanded_community_list_id` (String) Expanded community list ID, Attribute conditional on `type` being equal to `expandedCommunity`
- `expanded_community_list_version` (Number) Expanded community list version
- `group_id` (Number) Group ID, Attribute conditional on `type` being equal to `groupId`
  - Range: `0`-`4294967295`
- `omp_tag` (Number) OMP tag, Attribute conditional on `type` being equal to `ompTag`
  - Range: `0`-`4294967295`
- `origin` (String) Origin, Attribute conditional on `type` being equal to `origin`
  - Choices: `igp`, `egp`, `incomplete`, `aggregrate`, `bgp`, `bgp-external`, `bgp-internal`, `connected`, `eigrp`, `ospf`, `ospf-inter-area`, `ospf-intra-area`, `ospf-external1`, `ospf-external2`, `rip`, `static`, `eigrp-summary`, `eigrp-internal`, `eigrp-external`, `lisp`, `nat-dia`, `natpool`, `isis`, `isis-level1`, `isis-level2`
- `originator` (String) Originator IP, Attribute conditional on `type` being equal to `originator`
- `path_type` (String) Path type, Attribute conditional on `type` being equal to `pathType`
  - Choices: `hierarchical-path`, `direct-path`, `transport-gateway-path`
- `preference` (Number) Preference, Attribute conditional on `type` being equal to `preference`
  - Range: `0`-`4294967295`
- `prefix_list_id` (String) Prefix list ID, Attribute conditional on `type` being equal to `prefixList`
- `prefix_list_version` (Number) Prefix list version
- `site_id` (Number) Site ID, Attribute conditional on `type` being equal to `siteId`
  - Range: `0`-`4294967295`
- `site_list_id` (String) Site list ID, Attribute conditional on `type` being equal to `siteList`
- `site_list_version` (Number) Site list version
- `tloc_color` (String) TLOC color, Attribute conditional on `type` being equal to `tloc`
- `tloc_encapsulation` (String) TLOC encapsulation, Attribute conditional on `type` being equal to `tloc`
  - Choices: `ipsec`, `gre`
- `tloc_ip` (String) TLOC IP address, Attribute conditional on `type` being equal to `tloc`
- `tloc_list_id` (String) TLOC list ID, Attribute conditional on `type` being equal to `tlocList`
- `tloc_list_version` (Number) TLOC list version
- `vpn_id` (Number) VPN ID, Attribute conditional on `type` being equal to `vpn`
  - Range: `0`-`65536`
- `vpn_list_id` (String) VPN list ID, Attribute conditional on `type` being equal to `vpnList`
- `vpn_list_version` (Number) VPN list version

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import sdwan_custom_control_topology_policy_definition.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
