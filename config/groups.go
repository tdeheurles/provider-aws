// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/name"
)

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		// "aws_route53_resolver_rule": "route53resolver" -> (route53resolver, Rule)
		words := strings.Split(strings.TrimPrefix(resource, "aws_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupMap contains all overrides we'd like to make to the default group search.
// It's written with data from TF Provider AWS repo service grouping in here:
// https://github.com/hashicorp/terraform-provider-aws/tree/main/internal/service
//
// At the end, all of them are based on grouping of the AWS Go SDK.
// The initial grouping is calculated based on folder grouping of AWS TF Provider
// which is based on Go SDK. Here is the script used to fetch that list:
// https://gist.github.com/muvaf/8d61365ffc1df7757864422ba16d7819
var GroupMap = map[string]GroupKindCalculator{
	"aws_alb_listener_certificate":                             ReplaceGroupWords("elbv2", 0),
	"aws_alb_listener_rule":                                    ReplaceGroupWords("elbv2", 0),
	"aws_alb_listener":                                         ReplaceGroupWords("elbv2", 0),
	"aws_alb_target_group_attachment":                          ReplaceGroupWords("elbv2", 0),
	"aws_alb_target_group":                                     ReplaceGroupWords("elbv2", 0),
	"aws_alb":                                                  ReplaceGroupWords("elbv2", 0),
	"aws_ami_copy":                                             ReplaceGroupWords("ec2", 0),
	"aws_ami_from_instance":                                    ReplaceGroupWords("ec2", 0),
	"aws_ami_launch_permission":                                ReplaceGroupWords("ec2", 0),
	"aws_ami":                                                  ReplaceGroupWords("ec2", 0),
	"aws_api_gateway_account":                                  ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_api_key":                                  ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_authorizer":                               ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_base_path_mapping":                        ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_client_certificate":                       ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_deployment":                               ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_documentation_part":                       ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_documentation_version":                    ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_domain_name":                              ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_gateway_response":                         ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_integration_response":                     ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_integration":                              ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_method_response":                          ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_method_settings":                          ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_method":                                   ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_model":                                    ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_request_validator":                        ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_resource":                                 ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_rest_api_policy":                          ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_rest_api":                                 ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_stage":                                    ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_usage_plan_key":                           ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_usage_plan":                               ReplaceGroupWords("apigateway", 2),
	"aws_api_gateway_vpc_link":                                 ReplaceGroupWords("apigateway", 2),
	"aws_app_cookie_stickiness_policy":                         ReplaceGroupWords("elb", 0),
	"aws_cloudcontrolapi_resource":                             ReplaceGroupWords("cloudcontrol", 1),
	"aws_cloudhsm_v2_cluster":                                  ReplaceGroupWords("cloudhsmv2", 2),
	"aws_cloudhsm_v2_hsm":                                      ReplaceGroupWords("cloudhsmv2", 2),
	"aws_cloudtrail":                                           ReplaceGroupWords("cloudtrail", 0),
	"aws_cloudwatch_event_api_destination":                     ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_event_archive":                             ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_event_bus_policy":                          ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_event_bus":                                 ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_event_connection":                          ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_event_permission":                          ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_event_rule":                                ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_event_target":                              ReplaceGroupWords("cloudwatchevents", 2),
	"aws_cloudwatch_log_destination_policy":                    ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_cloudwatch_log_destination":                           ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_cloudwatch_log_group":                                 ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_cloudwatch_log_metric_filter":                         ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_cloudwatch_log_resource_policy":                       ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_cloudwatch_log_stream":                                ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_cloudwatch_log_subscription_filter":                   ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_cloudwatch_query_definition":                          ReplaceGroupWords("cloudwatchlogs", 2),
	"aws_codedeploy_app":                                       ReplaceGroupWords("deploy", 1),
	"aws_codedeploy_deployment_config":                         ReplaceGroupWords("deploy", 1),
	"aws_codedeploy_deployment_group":                          ReplaceGroupWords("deploy", 1),
	"aws_codepipeline":                                         ReplaceGroupWords("codepipeline", 0),
	"aws_cognito_identity_pool_provider_principal_tag":         ReplaceGroupWords("cognitoidentity", 0),
	"aws_cognito_identity_pool_roles_attachment":               ReplaceGroupWords("cognitoidentity", 2),
	"aws_cognito_identity_pool":                                ReplaceGroupWords("cognitoidentity", 2),
	"aws_cognito_identity_provider":                            ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_resource_server":                              ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_risk_configuration":                           ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_user_group":                                   ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_user_in_group":                                ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_user_pool_client":                             ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_user_pool_domain":                             ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_user_pool_ui_customization":                   ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_user_pool":                                    ReplaceGroupWords("cognitoidp", 1),
	"aws_cognito_user":                                         ReplaceGroupWords("cognitoidp", 1),
	"aws_config_aggregate_authorization":                       ReplaceGroupWords("configservice", 1),
	"aws_config_config_rule":                                   ReplaceGroupWords("configservice", 1),
	"aws_config_configuration_aggregator":                      ReplaceGroupWords("configservice", 1),
	"aws_config_configuration_recorder_status":                 ReplaceGroupWords("configservice", 1),
	"aws_config_configuration_recorder":                        ReplaceGroupWords("configservice", 1),
	"aws_config_conformance_pack":                              ReplaceGroupWords("configservice", 1),
	"aws_config_delivery_channel":                              ReplaceGroupWords("configservice", 1),
	"aws_config_organization_conformance_pack":                 ReplaceGroupWords("configservice", 1),
	"aws_config_organization_custom_rule":                      ReplaceGroupWords("configservice", 1),
	"aws_config_organization_managed_rule":                     ReplaceGroupWords("configservice", 1),
	"aws_config_remediation_configuration":                     ReplaceGroupWords("configservice", 1),
	"aws_customer_gateway":                                     ReplaceGroupWords("ec2", 0),
	"aws_db_cluster_snapshot":                                  ReplaceGroupWords("rds", 1),
	"aws_db_event_subscription":                                ReplaceGroupWords("rds", 1),
	"aws_db_instance_automated_backups_replication":            ReplaceGroupWords("rds", 0),
	"aws_db_instance_role_association":                         ReplaceGroupWords("rds", 1),
	"aws_db_instance":                                          ReplaceGroupWords("rds", 1),
	"aws_db_option_group":                                      ReplaceGroupWords("rds", 1),
	"aws_db_parameter_group":                                   ReplaceGroupWords("rds", 1),
	"aws_db_proxy_default_target_group":                        ReplaceGroupWords("rds", 1),
	"aws_db_proxy_endpoint":                                    ReplaceGroupWords("rds", 1),
	"aws_db_proxy_target":                                      ReplaceGroupWords("rds", 1),
	"aws_db_proxy":                                             ReplaceGroupWords("rds", 1),
	"aws_db_snapshot_copy":                                     ReplaceGroupWords("rds", 0),
	"aws_db_snapshot":                                          ReplaceGroupWords("rds", 1),
	"aws_db_subnet_group":                                      ReplaceGroupWords("rds", 1),
	"aws_default_network_acl":                                  ReplaceGroupWords("ec2", 0),
	"aws_default_route_table":                                  ReplaceGroupWords("ec2", 0),
	"aws_default_security_group":                               ReplaceGroupWords("ec2", 0),
	"aws_default_subnet":                                       ReplaceGroupWords("ec2", 0),
	"aws_default_vpc_dhcp_options":                             ReplaceGroupWords("ec2", 0),
	"aws_default_vpc":                                          ReplaceGroupWords("ec2", 0),
	"aws_directory_service_conditional_forwarder":              ReplaceGroupWords("ds", 2),
	"aws_directory_service_directory":                          ReplaceGroupWords("ds", 2),
	"aws_directory_service_log_subscription":                   ReplaceGroupWords("ds", 2),
	"aws_directory_service_shared_directory_accepter":          ReplaceGroupWords("ds", 2),
	"aws_directory_service_shared_directory":                   ReplaceGroupWords("ds", 2),
	"aws_dx_bgp_peer":                                          ReplaceGroupWords("directconnect", 1),
	"aws_dx_connection_association":                            ReplaceGroupWords("directconnect", 1),
	"aws_dx_connection_confirmation":                           ReplaceGroupWords("directconnect", 1),
	"aws_dx_connection":                                        ReplaceGroupWords("directconnect", 1),
	"aws_dx_gateway_association_proposal":                      ReplaceGroupWords("directconnect", 1),
	"aws_dx_gateway_association":                               ReplaceGroupWords("directconnect", 1),
	"aws_dx_gateway":                                           ReplaceGroupWords("directconnect", 1),
	"aws_dx_hosted_connection":                                 ReplaceGroupWords("directconnect", 1),
	"aws_dx_hosted_private_virtual_interface_accepter":         ReplaceGroupWords("directconnect", 1),
	"aws_dx_hosted_private_virtual_interface":                  ReplaceGroupWords("directconnect", 1),
	"aws_dx_hosted_public_virtual_interface_accepter":          ReplaceGroupWords("directconnect", 1),
	"aws_dx_hosted_public_virtual_interface":                   ReplaceGroupWords("directconnect", 1),
	"aws_dx_hosted_transit_virtual_interface_accepter":         ReplaceGroupWords("directconnect", 1),
	"aws_dx_hosted_transit_virtual_interface":                  ReplaceGroupWords("directconnect", 1),
	"aws_dx_lag":                                               ReplaceGroupWords("directconnect", 1),
	"aws_dx_private_virtual_interface":                         ReplaceGroupWords("directconnect", 1),
	"aws_dx_public_virtual_interface":                          ReplaceGroupWords("directconnect", 1),
	"aws_dx_transit_virtual_interface":                         ReplaceGroupWords("directconnect", 1),
	"aws_ebs_default_kms_key":                                  ReplaceGroupWords("ec2", 0),
	"aws_ebs_encryption_by_default":                            ReplaceGroupWords("ec2", 0),
	"aws_ebs_snapshot_copy":                                    ReplaceGroupWords("ec2", 0),
	"aws_ebs_snapshot_import":                                  ReplaceGroupWords("ec2", 0),
	"aws_ebs_snapshot":                                         ReplaceGroupWords("ec2", 0),
	"aws_ebs_volume":                                           ReplaceGroupWords("ec2", 0),
	"aws_egress_only_internet_gateway":                         ReplaceGroupWords("ec2", 0),
	"aws_eip_association":                                      ReplaceGroupWords("ec2", 0),
	"aws_eip":                                                  ReplaceGroupWords("ec2", 0),
	"aws_elastic_beanstalk_application_version":                ReplaceGroupWords("elasticbeanstalk", 2),
	"aws_elastic_beanstalk_application":                        ReplaceGroupWords("elasticbeanstalk", 2),
	"aws_elastic_beanstalk_configuration_template":             ReplaceGroupWords("elasticbeanstalk", 2),
	"aws_elastic_beanstalk_environment":                        ReplaceGroupWords("elasticbeanstalk", 2),
	"aws_elb":                                                  ReplaceGroupWords("elb", 0),
	"aws_flow_log":                                             ReplaceGroupWords("ec2", 0),
	"aws_instance":                                             ReplaceGroupWords("ec2", 0),
	"aws_internet_gateway_attachment":                          ReplaceGroupWords("ec2", 0),
	"aws_internet_gateway":                                     ReplaceGroupWords("ec2", 0),
	"aws_key_pair":                                             ReplaceGroupWords("ec2", 0),
	"aws_kinesis_analytics_application":                        ReplaceGroupWords("kinesisanalytics", 2),
	"aws_kinesis_firehose_delivery_stream":                     ReplaceGroupWords("firehose", 2),
	"aws_kinesis_video_stream":                                 ReplaceGroupWords("kinesisvideo", 2),
	"aws_launch_configuration":                                 ReplaceGroupWords("autoscaling", 0),
	"aws_launch_template":                                      ReplaceGroupWords("ec2", 0),
	"aws_lb_cookie_stickiness_policy":                          ReplaceGroupWords("elb", 0),
	"aws_lb_listener_certificate":                              ReplaceGroupWords("elbv2", 0),
	"aws_lb_listener_rule":                                     ReplaceGroupWords("elbv2", 0),
	"aws_lb_listener":                                          ReplaceGroupWords("elbv2", 0),
	"aws_lb_ssl_negotiation_policy":                            ReplaceGroupWords("elb", 0),
	"aws_lb_target_group_attachment":                           ReplaceGroupWords("elbv2", 0),
	"aws_lb_target_group":                                      ReplaceGroupWords("elbv2", 0),
	"aws_lb":                                                   ReplaceGroupWords("elbv2", 0),
	"aws_lex_bot_alias":                                        ReplaceGroupWords("lexmodels", 1),
	"aws_lex_bot":                                              ReplaceGroupWords("lexmodels", 1),
	"aws_lex_intent":                                           ReplaceGroupWords("lexmodels", 1),
	"aws_lex_slot_type":                                        ReplaceGroupWords("lexmodels", 1),
	"aws_load_balancer_backend_server_policy":                  ReplaceGroupWords("elb", 2),
	"aws_load_balancer_listener_policy":                        ReplaceGroupWords("elb", 2),
	"aws_load_balancer_policy":                                 ReplaceGroupWords("elb", 2),
	"aws_main_route_table_association":                         ReplaceGroupWords("ec2", 0),
	"aws_media_convert_queue":                                  ReplaceGroupWords("mediaconvert", 2),
	"aws_media_package_channel":                                ReplaceGroupWords("mediapackage", 2),
	"aws_media_store_container_policy":                         ReplaceGroupWords("mediastore", 2),
	"aws_media_store_container":                                ReplaceGroupWords("mediastore", 2),
	"aws_msk_cluster":                                          ReplaceGroupWords("kafka", 1),
	"aws_msk_configuration":                                    ReplaceGroupWords("kafka", 1),
	"aws_msk_replicator":                                       ReplaceGroupWords("kafka", 1),
	"aws_msk_scram_secret_association":                         ReplaceGroupWords("kafka", 1),
	"aws_msk_serverless_cluster":                               ReplaceGroupWords("kafka", 1),
	"aws_msk_single_scram_secret_association":                  ReplaceGroupWords("kafka", 1),
	"aws_mskconnect_connector":                                 ReplaceGroupWords("kafkaconnect", 1),
	"aws_mskconnect_custom_plugin":                             ReplaceGroupWords("kafkaconnect", 1),
	"aws_mskconnect_worker_configuration":                      ReplaceGroupWords("kafkaconnect", 1),
	"aws_nat_gateway":                                          ReplaceGroupWords("ec2", 0),
	"aws_network_acl_association":                              ReplaceGroupWords("ec2", 0),
	"aws_network_acl_rule":                                     ReplaceGroupWords("ec2", 0),
	"aws_network_acl":                                          ReplaceGroupWords("ec2", 0),
	"aws_network_interface_attachment":                         ReplaceGroupWords("ec2", 0),
	"aws_network_interface_sg_attachment":                      ReplaceGroupWords("ec2", 0),
	"aws_network_interface":                                    ReplaceGroupWords("ec2", 0),
	"aws_placement_group":                                      ReplaceGroupWords("ec2", 0),
	"aws_prometheus_alert_manager_definition":                  ReplaceGroupWords("amp", 1),
	"aws_prometheus_rule_group_namespace":                      ReplaceGroupWords("amp", 1),
	"aws_prometheus_workspace":                                 ReplaceGroupWords("amp", 1),
	"aws_proxy_protocol_policy":                                ReplaceGroupWords("elb", 0),
	"aws_route_table_association":                              ReplaceGroupWords("ec2", 0),
	"aws_route_table":                                          ReplaceGroupWords("ec2", 0),
	"aws_route":                                                ReplaceGroupWords("ec2", 0),
	"aws_route53_resolver_dnssec_config":                       ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_endpoint":                            ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_firewall_config":                     ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_firewall_domain_list":                ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_firewall_rule_group_association":     ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_firewall_rule_group":                 ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_firewall_rule":                       ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_query_log_config_association":        ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_query_log_config":                    ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_rule_association":                    ReplaceGroupWords("route53resolver", 2),
	"aws_route53_resolver_rule":                                ReplaceGroupWords("route53resolver", 2),
	"aws_s3_access_point":                                      ReplaceGroupWords("s3control", 1),
	"aws_s3_account_public_access_block":                       ReplaceGroupWords("s3control", 1),
	"aws_security_group_rule":                                  ReplaceGroupWords("ec2", 0),
	"aws_security_group":                                       ReplaceGroupWords("ec2", 0),
	"aws_serverlessapplicationrepository_cloudformation_stack": ReplaceGroupWords("serverlessrepo", 1),
	"aws_service_discovery_http_namespace":                     ReplaceGroupWords("servicediscovery", 2),
	"aws_service_discovery_instance":                           ReplaceGroupWords("servicediscovery", 2),
	"aws_service_discovery_private_dns_namespace":              ReplaceGroupWords("servicediscovery", 2),
	"aws_service_discovery_public_dns_namespace":               ReplaceGroupWords("servicediscovery", 2),
	"aws_service_discovery_service":                            ReplaceGroupWords("servicediscovery", 2),
	"aws_snapshot_create_volume_permission":                    ReplaceGroupWords("ec2", 0),
	"aws_spot_datafeed_subscription":                           ReplaceGroupWords("ec2", 0),
	"aws_spot_fleet_request":                                   ReplaceGroupWords("ec2", 0),
	"aws_spot_instance_request":                                ReplaceGroupWords("ec2", 0),
	"aws_subnet":                                               ReplaceGroupWords("ec2", 0),
	"aws_volume_attachment":                                    ReplaceGroupWords("ec2", 0),
	"aws_vpc_dhcp_options_association":                         ReplaceGroupWords("ec2", 0),
	"aws_vpc_dhcp_options":                                     ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_connection_accepter":                     ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_connection_notification":                 ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_policy":                                  ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_route_table_association":                 ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_security_group_association":              ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_service_allowed_principal":               ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_service":                                 ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint_subnet_association":                      ReplaceGroupWords("ec2", 0),
	"aws_vpc_endpoint":                                         ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipam_organization_admin_account":                  ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipam_pool_cidr_allocation":                        ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipam_pool_cidr":                                   ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipam_pool":                                        ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipam_preview_next_cidr":                           ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipam_scope":                                       ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipam":                                             ReplaceGroupWords("ec2", 0),
	"aws_vpc_security_group_egress_rule":                       ReplaceGroupWords("ec2", 0),
	"aws_vpc_security_group_ingress_rule":                      ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipv4_cidr_block_association":                      ReplaceGroupWords("ec2", 0),
	"aws_vpc_ipv6_cidr_block_association":                      ReplaceGroupWords("ec2", 0),
	"aws_vpc_peering_connection_accepter":                      ReplaceGroupWords("ec2", 0),
	"aws_vpc_peering_connection_options":                       ReplaceGroupWords("ec2", 0),
	"aws_vpc_peering_connection":                               ReplaceGroupWords("ec2", 0),
	"aws_vpc":                                                  ReplaceGroupWords("ec2", 0),
	"aws_vpn_connection_route":                                 ReplaceGroupWords("ec2", 0),
	"aws_vpn_connection":                                       ReplaceGroupWords("ec2", 0),
	"aws_vpn_gateway_attachment":                               ReplaceGroupWords("ec2", 0),
	"aws_vpn_gateway_route_propagation":                        ReplaceGroupWords("ec2", 0),
	"aws_vpn_gateway":                                          ReplaceGroupWords("ec2", 0),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"aws_autoscaling_group":                    "AutoscalingGroup",
	"aws_cloudformation_type":                  "CloudFormationType",
	"aws_cloudtrail":                           "Trail",
	"aws_config_configuration_recorder_status": "AWSConfigurationRecorderStatus",
}
