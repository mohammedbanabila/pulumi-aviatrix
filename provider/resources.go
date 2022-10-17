// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aviatrix

import (
	"fmt"
	"path/filepath"

	"github.com/astipkovits/pulumi-aviatrix/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	//"github.com/terraform-providers/terraform-provider-aviatrix/aviatrix"
	"github.com/AviatrixSystems/terraform-provider-aviatrix/v2/aviatrix"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "aviatrix"
	// modules:
	mainMod = "index" // the aviatrix module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(aviatrix.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "aviatrix",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Aviatrix",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Aviatrix",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/astipkovits/pulumi-aviatrix/raw/main/releases/",
		Description:       "A Pulumi package for creating and managing aviatrix cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "aviatrix", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.aviatrix.com",
		Repository: "https://github.com/astipkovits/pulumi-aviatrix",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "AviatrixSystems",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"aviatrix_vpc":                                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixVpc"), Docs: &tfbridge.DocInfo{Source: "github.com/AviatrixSystems/terraform-provider-aviatrix/docs/resources/aviatrix_vpc.md"}},
			"aviatrix_account":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAccount")},
			"aviatrix_account_user":                                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAccountUser")},
			"aviatrix_app_domain":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAppDomain")},
			"aviatrix_arm_peer":                                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixArmPeer")},
			"aviatrix_aws_guard_duty":                                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsGuardDuty")},
			"aviatrix_aws_peer":                                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsPeer")},
			"aviatrix_aws_tgw":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgw")},
			"aviatrix_aws_tgw_connect":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwConnect")},
			"aviatrix_aws_tgw_connect_peer":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwConnectPeer")},
			"aviatrix_aws_tgw_directconnect":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwDirectconnect")},
			"aviatrix_aws_tgw_intra_domain_inspection":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwIntraDomainInspection")},
			"aviatrix_aws_tgw_network_domain":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwNetworkDomain")},
			"aviatrix_aws_tgw_peering":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwPeering")},
			"aviatrix_aws_tgw_peering_domain_conn":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwPeeringDomainConn")},
			"aviatrix_aws_tgw_security_domain":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwSecurityDomain")},
			"aviatrix_aws_tgw_security_domain_connection":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwSecurityDomainConn")},
			"aviatrix_aws_tgw_transit_gateway_attachment":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwTransitGatewayAttachment")},
			"aviatrix_aws_tgw_vpc_attachment":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwVpcAttachment")},
			"aviatrix_aws_tgw_vpn_conn":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAwsTgwVpnConn")},
			"aviatrix_azure_spoke_native_peering":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAzureSpokeNativePeering")},
			"aviatrix_azure_vng_conn":                                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAzureVngConn")},
			"aviatrix_cloudn_registration":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixCloudnRegistration")},
			"aviatrix_cloudn_transit_gateway_attachment":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixCloudnTransitGatewayAttachment")},
			"aviatrix_cloudwatch_agent":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixCloudwatchAgent")},
			"aviatrix_controller_bgp_max_as_limit_config":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerBgpMaxAsLimitConfig")},
			"aviatrix_controller_cert_domain_config":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerCertDomainConfig")},
			"aviatrix_controller_config":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerConfig")},
			"aviatrix_controller_email_config":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerEmailConfig")},
			"aviatrix_controller_email_exception_notification_config": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerEmailExceptionNotificationConfig")},
			"aviatrix_controller_gateway_keepalive_config":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerGatewayKeepaliveConfig")},
			"aviatrix_controller_private_mode_config":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerPrivateModeConfig")},
			"aviatrix_controller_private_oob":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerPrivateOob")},
			"aviatrix_controller_security_group_management_config":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixControllerSecurityGroupManagementConfig")},
			"aviatrix_copilot_association":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixCopilotAssociation")},
			"aviatrix_copilot_security_group_management_config":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixCopilotSecurityGroupManagementConfig")},
			"aviatrix_datadog_agent":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixDatadogAgent")},
			"aviatrix_device_interface_config":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixDeviceInterfaceConfig")},
			"aviatrix_edge_caag":                                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixEdgeCaag")},
			"aviatrix_edge_spoke":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixEdgeSpoke")},
			"aviatrix_edge_spoke_external_device_conn":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixEdgeSpokeExternalDeviceConn")},
			"aviatrix_edge_spoke_transit_attachment":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixEdgeSpokeTransitAttachment")},
			"aviatrix_fqdn":                                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFqdn")},
			"aviatrix_fqdn_pass_through":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFqdnPassThrough")},
			"aviatrix_fqdn_tag_rule":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFqdnTagRule")},
			"aviatrix_gateway":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixGateway")},
			"aviatrix_site2cloud":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSite2Cloud")},
			"aviatrix_spoke_gateway":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSpokeGateway")},
			"aviatrix_spoke_transit_attachment":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSpokeTransitAttachment")},
			"aviatrix_spoke_vpc":                                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSpokeVpc")},
			"aviatrix_trans_peer":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTransPeer")},
			"aviatrix_transit_gateway":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTransitGateway")},
			"aviatrix_transit_gateway_peering":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTransitGatewayPeering")},
			"aviatrix_transit_vpc":                                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTransitVpc")},
			"aviatrix_azure_peer":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixAzurePeer")},
			"aviatrix_filebeat_forwarder":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFilebeatForwarder")},
			"aviatrix_firenet":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFirenet")},
			"aviatrix_firewall":                                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFirewall")},
			"aviatrix_firewall_instance":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFirewallInstance")},
			"aviatrix_firewall_instance_association":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFirewallInstanceAssociation")},
			"aviatrix_firewall_management_access":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFirewallManagementAccess")},
			"aviatrix_firewall_policy":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFirewallPolicy")},
			"aviatrix_firewall_tag":                                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixFirewallTag")},
			"aviatrix_gateway_certificate_config":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixGatewayCertificateConfig")},
			"aviatrix_gateway_dnat":                                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixGatewayDnat")},
			"aviatrix_gateway_snat":                                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixGatewaySnat")},
			"aviatrix_geo_vpn":                                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixGeoVpn")},
			"aviatrix_microseg_policy_list":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixMicrosegPolicyList")},
			"aviatrix_netflow_agent":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixNetflowAgent")},
			"aviatrix_periodic_ping":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixPeriodicPing")},
			"aviatrix_private_mode_lb":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixPrivateModeLb")},
			"aviatrix_private_mode_multicloud_endpoint":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixPrivateModeMulticloudEndpoint")},
			"aviatrix_proxy_config":                                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixProxyConfig")},
			"aviatrix_rbac_group":                                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixRbacGroup")},
			"aviatrix_rbac_group_access_account_attachment":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixRbacGroupAccessAccountAttachment")},
			"aviatrix_rbac_group_permission_attachment":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixRbacGroupPermissionAttachment")},
			"aviatrix_rbac_group_user_attachment":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixRbacGroupUserAttachment")},
			"aviatrix_remote_syslog":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixRemoteSyslog")},
			"aviatrix_saml_endpoint":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSamlEndpoint")},
			"aviatrix_segmentation_network_domain":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSegmentationNetworkDomain")},
			"aviatrix_segmentation_network_domain_association":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSegmentationNetworkDomainAssociation")},
			"aviatrix_segmentation_network_domain_connection_policy":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSegmentationNetworkDomainConnectionPolicy")},
			"aviatrix_segmentation_security_domain":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSegmentationSecurityDomain")},
			"aviatrix_segmentation_security_domain_association":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSegmentationSecurityDomainAssociation")},
			"aviatrix_segmentation_security_domain_connection_policy": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSegmentationSecurityDomainConnectionPolicy")},
			"aviatrix_site2cloud_ca_cert_tag":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSite2CloudCaCertTag")},
			"aviatrix_splunk_logging":                                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSplunkLogging")},
			"aviatrix_spoke_external_device_conn":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "aviatrix_spoke_external_device_conn")},
			"aviatrix_spoke_gateway_subnet_group":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSpokeGatewaySubnetGroup")},
			"aviatrix_sumologic_forwarder":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixSumologicForwarder")},
			"aviatrix_transit_cloudn_conn":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTransitCloudnConn")},
			"aviatrix_transit_external_device_conn":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTransitExternalDeviceConn")},
			"aviatrix_transit_firenet_policy":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTransitFirenetPolicy")},
			"aviatrix_tunnel":                                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixTunnel")},
			"aviatrix_vgw_conn":                                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixVgwConn")},
			"aviatrix_vpn_cert_download":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixVpnCertDownload")},
			"aviatrix_vpn_profile":                                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixVpnProfile")},
			"aviatrix_vpn_user":                                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixVpnUser")},
			"aviatrix_vpn_user_accelerator":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AviatrixVpnUserAccelerator")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function.
			"aviatrix_spoke_gateway":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixSpokeGateway")},
			"aviatrix_vpc":                              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixVpc")},
			"aviatrix_account":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixAccount")},
			"aviatrix_caller_identity":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixCallerIdentity")},
			"aviatrix_device_interfaces":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixDeviceInterfaces")},
			"aviatrix_firenet":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixFirenet")},
			"aviatrix_firenet_firewall_manager":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixFirenetFirewallManager")},
			"aviatrix_firenet_vendor_integration":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixFirenetVendorIntegration")},
			"aviatrix_firewall":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixFirewall")},
			"aviatrix_firewall_instance_images":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixFirewallInstanceImages")},
			"aviatrix_gateway":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixGateway")},
			"aviatrix_gateway_image":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixGatewayImage")},
			"aviatrix_network_domains":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixNetworkDomains")},
			"aviatrix_spoke_gateway_inspection_subnets": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixSpokeGatewayInspectionSubnets")},
			"aviatrix_transit_gateway":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixTransitGateway")},
			"aviatrix_transit_gateways":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixTransitGateways")},
			"aviatrix_vpc_tracker":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAviatrixVpcTracker")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@astipkovits/aviatrix",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/astipkovits/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			//RootNamespace: "astipkovits",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
