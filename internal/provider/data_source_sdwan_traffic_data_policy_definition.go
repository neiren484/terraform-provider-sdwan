// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TrafficDataPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &TrafficDataPolicyDefinitionDataSource{}
)

func NewTrafficDataPolicyDefinitionDataSource() datasource.DataSource {
	return &TrafficDataPolicyDefinitionDataSource{}
}

type TrafficDataPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *TrafficDataPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traffic_data_policy_definition"
}

func (d *TrafficDataPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Traffic Data Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Computed:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Default action, either `accept` or `drop`",
				Computed:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "List of sequences",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Sequence ID",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence name",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Sequence type",
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: "Sequence IP type, either `ipv4`, `ipv6` or `all`",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base action, either `accept` or `drop`",
							Computed:            true,
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of match entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of match entry",
										Computed:            true,
									},
									"application_list_id": schema.StringAttribute{
										MarkdownDescription: "Application list ID",
										Computed:            true,
									},
									"application_list_version": schema.Int64Attribute{
										MarkdownDescription: "Application list version",
										Computed:            true,
									},
									"dns_application_list_id": schema.StringAttribute{
										MarkdownDescription: "DNS Application list ID",
										Computed:            true,
									},
									"dns_application_list_version": schema.Int64Attribute{
										MarkdownDescription: "DNS Application list version",
										Computed:            true,
									},
									"icmp_message": schema.StringAttribute{
										MarkdownDescription: "ICMP Message",
										Computed:            true,
									},
									"dns": schema.StringAttribute{
										MarkdownDescription: "DNS request or response",
										Computed:            true,
									},
									"dscp": schema.StringAttribute{
										MarkdownDescription: "DSCP value",
										Computed:            true,
									},
									"packet_length": schema.Int64Attribute{
										MarkdownDescription: "Packet length",
										Computed:            true,
									},
									"plp": schema.StringAttribute{
										MarkdownDescription: "PLP",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "IP Protocol, 0-255 (Single value or multiple values separated by spaces)",
										Computed:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Source Data Prefix list ID",
										Computed:            true,
									},
									"source_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Source Data Prefix list version",
										Computed:            true,
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: "Source IP",
										Computed:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: "Source port, 0-65535 (Single value, range or multiple values separated by spaces)",
										Computed:            true,
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Destination Data Prefix list ID",
										Computed:            true,
									},
									"destination_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Destination Data Prefix list version",
										Computed:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: "Destination IP",
										Computed:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: "Destination port, 0-65535 (Single value, range or multiple values separated by spaces)",
										Computed:            true,
									},
									"destination_region": schema.StringAttribute{
										MarkdownDescription: "Destination region",
										Computed:            true,
									},
									"tcp": schema.StringAttribute{
										MarkdownDescription: "TCP flags",
										Computed:            true,
									},
									"traffic_to": schema.StringAttribute{
										MarkdownDescription: "Traffic to",
										Computed:            true,
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of action entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of action entry",
										Computed:            true,
									},
									"cflowd": schema.BoolAttribute{
										MarkdownDescription: "Enable cflowd",
										Computed:            true,
									},
									"counter": schema.StringAttribute{
										MarkdownDescription: "Counter name",
										Computed:            true,
									},
									"dre_optimization": schema.BoolAttribute{
										MarkdownDescription: "Enable DRE optimization",
										Computed:            true,
									},
									"fallback_to_routing": schema.BoolAttribute{
										MarkdownDescription: "Enable fallback to routing",
										Computed:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: "Enable logging",
										Computed:            true,
									},
									"loss_correction": schema.StringAttribute{
										MarkdownDescription: "Loss correction",
										Computed:            true,
									},
									"loss_correction_fec": schema.StringAttribute{
										MarkdownDescription: "Loss correction FEC",
										Computed:            true,
									},
									"loss_correction_packet_duplication": schema.StringAttribute{
										MarkdownDescription: "Loss correction packet duplication",
										Computed:            true,
									},
									"loss_correction_fec_threshold": schema.StringAttribute{
										MarkdownDescription: "Loss correction FEC threshold",
										Computed:            true,
									},
									"nat_pool": schema.StringAttribute{
										MarkdownDescription: "NAT pool",
										Computed:            true,
									},
									"nat_pool_id": schema.Int64Attribute{
										MarkdownDescription: "NAT pool ID",
										Computed:            true,
									},
									"redirect_dns": schema.StringAttribute{
										MarkdownDescription: "Redirect DNS",
										Computed:            true,
									},
									"redirect_dns_type": schema.StringAttribute{
										MarkdownDescription: "Redirect DNS type",
										Computed:            true,
									},
									"redirect_dns_address": schema.StringAttribute{
										MarkdownDescription: "Redirect DNS IP address",
										Computed:            true,
									},
									"service_node_group": schema.StringAttribute{
										MarkdownDescription: "Service node group",
										Computed:            true,
									},
									"secure_internet_gateway": schema.BoolAttribute{
										MarkdownDescription: "Enable secure internet gateway",
										Computed:            true,
									},
									"tcp_optimization": schema.BoolAttribute{
										MarkdownDescription: "Enable TCP optimization",
										Computed:            true,
									},
									"set_parameters": schema.ListNestedAttribute{
										MarkdownDescription: "List of set parameters",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "Type of set parameter",
													Computed:            true,
												},
												"dscp": schema.Int64Attribute{
													MarkdownDescription: "DSCP",
													Computed:            true,
												},
												"forwarding_class": schema.StringAttribute{
													MarkdownDescription: "Forwarding class",
													Computed:            true,
												},
												"next_hop": schema.StringAttribute{
													MarkdownDescription: "Next hop IP",
													Computed:            true,
												},
												"local_tloc_list_color": schema.StringAttribute{
													MarkdownDescription: "Local TLOC list color. Space separated list of colors.",
													Computed:            true,
												},
												"local_tloc_list_encap": schema.StringAttribute{
													MarkdownDescription: "Local TLOC list encapsulation.",
													Computed:            true,
												},
												"local_tloc_list_restrict": schema.BoolAttribute{
													MarkdownDescription: "Local TLOC list restrict",
													Computed:            true,
												},
												"next_hop_loose": schema.BoolAttribute{
													MarkdownDescription: "Use routing table entry to forward the packet in case Next-hop is not available",
													Computed:            true,
												},
												"policer_list_id": schema.StringAttribute{
													MarkdownDescription: "Policer list ID",
													Computed:            true,
												},
												"policer_list_version": schema.Int64Attribute{
													MarkdownDescription: "Policer list version",
													Computed:            true,
												},
												"preferred_color_group_list": schema.StringAttribute{
													MarkdownDescription: "Preferred color group list ID",
													Computed:            true,
												},
												"preferred_color_group_list_version": schema.Int64Attribute{
													MarkdownDescription: "Preferred color group list version",
													Computed:            true,
												},
												"tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "TLOC list ID",
													Computed:            true,
												},
												"tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: "TLOC list version",
													Computed:            true,
												},
												"tloc_ip": schema.StringAttribute{
													MarkdownDescription: "TLOC IP address",
													Computed:            true,
												},
												"tloc_color": schema.StringAttribute{
													MarkdownDescription: "TLOC color",
													Computed:            true,
												},
												"tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "TLOC encapsulation",
													Computed:            true,
												},
												"service_type": schema.StringAttribute{
													MarkdownDescription: "Service type",
													Computed:            true,
												},
												"service_vpn_id": schema.Int64Attribute{
													MarkdownDescription: "Service VPN ID",
													Computed:            true,
												},
												"service_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "Service TLOC list ID",
													Computed:            true,
												},
												"service_tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: "Service TLOC list version",
													Computed:            true,
												},
												"service_tloc_ip": schema.StringAttribute{
													MarkdownDescription: "Service TLOC IP address",
													Computed:            true,
												},
												"service_tloc_local": schema.BoolAttribute{
													MarkdownDescription: "Service TLOC Local",
													Computed:            true,
												},
												"service_tloc_restrict": schema.BoolAttribute{
													MarkdownDescription: "Service TLOC Restrict",
													Computed:            true,
												},
												"service_tloc_color": schema.StringAttribute{
													MarkdownDescription: "Service TLOC color",
													Computed:            true,
												},
												"service_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "Service TLOC encapsulation",
													Computed:            true,
												},
												"vpn_id": schema.Int64Attribute{
													MarkdownDescription: "DSCP",
													Computed:            true,
												},
											},
										},
									},
									"nat_parameters": schema.ListNestedAttribute{
										MarkdownDescription: "List of NAT parameters",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "Type of NAT parameter",
													Computed:            true,
												},
												"vpn_id": schema.Int64Attribute{
													MarkdownDescription: "DSCP",
													Computed:            true,
												},
												"fallback": schema.BoolAttribute{
													MarkdownDescription: "Fallback",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *TrafficDataPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TrafficDataPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TrafficDataPolicyDefinition

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)
	config.Type = types.StringValue("data")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
