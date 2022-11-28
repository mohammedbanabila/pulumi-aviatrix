// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_firenet** resource allows the creation and management of [Aviatrix Firewall Networks](https://docs.aviatrix.com/HowTos/firewall_network_faq.html).
//
// > **NOTE:** This resource is used in conjunction with multiple other resources that may include, and are not limited to: **firewall_instance**, **firewall_instance_association**, **aws_tgw**, and **transit_gateway** resources or even **aviatrix_fqdn**, under the Aviatrix FireNet solution. Explicit dependencies may be set using `dependsOn`. For more information on proper FireNet configuration, please see the workflow [here](https://docs.aviatrix.com/HowTos/firewall_network_workflow.html).
//
// ## Import
//
// **firenet** can be imported using the `vpc_id`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixFirenet:AviatrixFirenet test vpc_id
//
// ```
type AviatrixFirenet struct {
	pulumi.CustomResourceState

	// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
	EastWestInspectionExcludedCidrs pulumi.StringArrayOutput `pulumi:"eastWestInspectionExcludedCidrs"`
	// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
	EgressEnabled pulumi.BoolPtrOutput `pulumi:"egressEnabled"`
	// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
	EgressStaticCidrs pulumi.StringArrayOutput `pulumi:"egressStaticCidrs"`
	// Dynamic block of firewall instance(s) to be associated with the FireNet.
	//
	// Deprecated: Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
	FirewallInstanceAssociations AviatrixFirenetFirewallInstanceAssociationArrayOutput `pulumi:"firewallInstanceAssociations"`
	// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
	HashingAlgorithm pulumi.StringPtrOutput `pulumi:"hashingAlgorithm"`
	// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
	InspectionEnabled pulumi.BoolPtrOutput `pulumi:"inspectionEnabled"`
	// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
	KeepAliveViaLanInterfaceEnabled pulumi.BoolPtrOutput `pulumi:"keepAliveViaLanInterfaceEnabled"`
	// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewallInstanceAssociation` blocks can be used. If set to false, all firewall associations must be managed via standalone `AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
	ManageFirewallInstanceAssociation pulumi.BoolPtrOutput `pulumi:"manageFirewallInstanceAssociation"`
	// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
	TgwSegmentationForEgressEnabled pulumi.BoolPtrOutput `pulumi:"tgwSegmentationForEgressEnabled"`
	// VPC ID of the Security VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewAviatrixFirenet registers a new resource with the given unique name, arguments, and options.
func NewAviatrixFirenet(ctx *pulumi.Context,
	name string, args *AviatrixFirenetArgs, opts ...pulumi.ResourceOption) (*AviatrixFirenet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixFirenet
	err := ctx.RegisterResource("aviatrix:index/aviatrixFirenet:AviatrixFirenet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixFirenet gets an existing AviatrixFirenet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixFirenet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixFirenetState, opts ...pulumi.ResourceOption) (*AviatrixFirenet, error) {
	var resource AviatrixFirenet
	err := ctx.ReadResource("aviatrix:index/aviatrixFirenet:AviatrixFirenet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixFirenet resources.
type aviatrixFirenetState struct {
	// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
	EastWestInspectionExcludedCidrs []string `pulumi:"eastWestInspectionExcludedCidrs"`
	// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
	EgressEnabled *bool `pulumi:"egressEnabled"`
	// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
	EgressStaticCidrs []string `pulumi:"egressStaticCidrs"`
	// Dynamic block of firewall instance(s) to be associated with the FireNet.
	//
	// Deprecated: Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
	FirewallInstanceAssociations []AviatrixFirenetFirewallInstanceAssociation `pulumi:"firewallInstanceAssociations"`
	// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
	HashingAlgorithm *string `pulumi:"hashingAlgorithm"`
	// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
	InspectionEnabled *bool `pulumi:"inspectionEnabled"`
	// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
	KeepAliveViaLanInterfaceEnabled *bool `pulumi:"keepAliveViaLanInterfaceEnabled"`
	// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewallInstanceAssociation` blocks can be used. If set to false, all firewall associations must be managed via standalone `AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
	ManageFirewallInstanceAssociation *bool `pulumi:"manageFirewallInstanceAssociation"`
	// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
	TgwSegmentationForEgressEnabled *bool `pulumi:"tgwSegmentationForEgressEnabled"`
	// VPC ID of the Security VPC.
	VpcId *string `pulumi:"vpcId"`
}

type AviatrixFirenetState struct {
	// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
	EastWestInspectionExcludedCidrs pulumi.StringArrayInput
	// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
	EgressEnabled pulumi.BoolPtrInput
	// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
	EgressStaticCidrs pulumi.StringArrayInput
	// Dynamic block of firewall instance(s) to be associated with the FireNet.
	//
	// Deprecated: Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
	FirewallInstanceAssociations AviatrixFirenetFirewallInstanceAssociationArrayInput
	// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
	HashingAlgorithm pulumi.StringPtrInput
	// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
	InspectionEnabled pulumi.BoolPtrInput
	// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
	KeepAliveViaLanInterfaceEnabled pulumi.BoolPtrInput
	// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewallInstanceAssociation` blocks can be used. If set to false, all firewall associations must be managed via standalone `AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
	ManageFirewallInstanceAssociation pulumi.BoolPtrInput
	// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
	TgwSegmentationForEgressEnabled pulumi.BoolPtrInput
	// VPC ID of the Security VPC.
	VpcId pulumi.StringPtrInput
}

func (AviatrixFirenetState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixFirenetState)(nil)).Elem()
}

type aviatrixFirenetArgs struct {
	// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
	EastWestInspectionExcludedCidrs []string `pulumi:"eastWestInspectionExcludedCidrs"`
	// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
	EgressEnabled *bool `pulumi:"egressEnabled"`
	// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
	EgressStaticCidrs []string `pulumi:"egressStaticCidrs"`
	// Dynamic block of firewall instance(s) to be associated with the FireNet.
	//
	// Deprecated: Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
	FirewallInstanceAssociations []AviatrixFirenetFirewallInstanceAssociation `pulumi:"firewallInstanceAssociations"`
	// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
	HashingAlgorithm *string `pulumi:"hashingAlgorithm"`
	// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
	InspectionEnabled *bool `pulumi:"inspectionEnabled"`
	// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
	KeepAliveViaLanInterfaceEnabled *bool `pulumi:"keepAliveViaLanInterfaceEnabled"`
	// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewallInstanceAssociation` blocks can be used. If set to false, all firewall associations must be managed via standalone `AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
	ManageFirewallInstanceAssociation *bool `pulumi:"manageFirewallInstanceAssociation"`
	// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
	TgwSegmentationForEgressEnabled *bool `pulumi:"tgwSegmentationForEgressEnabled"`
	// VPC ID of the Security VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a AviatrixFirenet resource.
type AviatrixFirenetArgs struct {
	// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
	EastWestInspectionExcludedCidrs pulumi.StringArrayInput
	// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
	EgressEnabled pulumi.BoolPtrInput
	// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
	EgressStaticCidrs pulumi.StringArrayInput
	// Dynamic block of firewall instance(s) to be associated with the FireNet.
	//
	// Deprecated: Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
	FirewallInstanceAssociations AviatrixFirenetFirewallInstanceAssociationArrayInput
	// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
	HashingAlgorithm pulumi.StringPtrInput
	// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
	InspectionEnabled pulumi.BoolPtrInput
	// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
	KeepAliveViaLanInterfaceEnabled pulumi.BoolPtrInput
	// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewallInstanceAssociation` blocks can be used. If set to false, all firewall associations must be managed via standalone `AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
	ManageFirewallInstanceAssociation pulumi.BoolPtrInput
	// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
	TgwSegmentationForEgressEnabled pulumi.BoolPtrInput
	// VPC ID of the Security VPC.
	VpcId pulumi.StringInput
}

func (AviatrixFirenetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixFirenetArgs)(nil)).Elem()
}

type AviatrixFirenetInput interface {
	pulumi.Input

	ToAviatrixFirenetOutput() AviatrixFirenetOutput
	ToAviatrixFirenetOutputWithContext(ctx context.Context) AviatrixFirenetOutput
}

func (*AviatrixFirenet) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixFirenet)(nil)).Elem()
}

func (i *AviatrixFirenet) ToAviatrixFirenetOutput() AviatrixFirenetOutput {
	return i.ToAviatrixFirenetOutputWithContext(context.Background())
}

func (i *AviatrixFirenet) ToAviatrixFirenetOutputWithContext(ctx context.Context) AviatrixFirenetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirenetOutput)
}

// AviatrixFirenetArrayInput is an input type that accepts AviatrixFirenetArray and AviatrixFirenetArrayOutput values.
// You can construct a concrete instance of `AviatrixFirenetArrayInput` via:
//
//	AviatrixFirenetArray{ AviatrixFirenetArgs{...} }
type AviatrixFirenetArrayInput interface {
	pulumi.Input

	ToAviatrixFirenetArrayOutput() AviatrixFirenetArrayOutput
	ToAviatrixFirenetArrayOutputWithContext(context.Context) AviatrixFirenetArrayOutput
}

type AviatrixFirenetArray []AviatrixFirenetInput

func (AviatrixFirenetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixFirenet)(nil)).Elem()
}

func (i AviatrixFirenetArray) ToAviatrixFirenetArrayOutput() AviatrixFirenetArrayOutput {
	return i.ToAviatrixFirenetArrayOutputWithContext(context.Background())
}

func (i AviatrixFirenetArray) ToAviatrixFirenetArrayOutputWithContext(ctx context.Context) AviatrixFirenetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirenetArrayOutput)
}

// AviatrixFirenetMapInput is an input type that accepts AviatrixFirenetMap and AviatrixFirenetMapOutput values.
// You can construct a concrete instance of `AviatrixFirenetMapInput` via:
//
//	AviatrixFirenetMap{ "key": AviatrixFirenetArgs{...} }
type AviatrixFirenetMapInput interface {
	pulumi.Input

	ToAviatrixFirenetMapOutput() AviatrixFirenetMapOutput
	ToAviatrixFirenetMapOutputWithContext(context.Context) AviatrixFirenetMapOutput
}

type AviatrixFirenetMap map[string]AviatrixFirenetInput

func (AviatrixFirenetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixFirenet)(nil)).Elem()
}

func (i AviatrixFirenetMap) ToAviatrixFirenetMapOutput() AviatrixFirenetMapOutput {
	return i.ToAviatrixFirenetMapOutputWithContext(context.Background())
}

func (i AviatrixFirenetMap) ToAviatrixFirenetMapOutputWithContext(ctx context.Context) AviatrixFirenetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirenetMapOutput)
}

type AviatrixFirenetOutput struct{ *pulumi.OutputState }

func (AviatrixFirenetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixFirenet)(nil)).Elem()
}

func (o AviatrixFirenetOutput) ToAviatrixFirenetOutput() AviatrixFirenetOutput {
	return o
}

func (o AviatrixFirenetOutput) ToAviatrixFirenetOutputWithContext(ctx context.Context) AviatrixFirenetOutput {
	return o
}

// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
func (o AviatrixFirenetOutput) EastWestInspectionExcludedCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.StringArrayOutput { return v.EastWestInspectionExcludedCidrs }).(pulumi.StringArrayOutput)
}

// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
func (o AviatrixFirenetOutput) EgressEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.BoolPtrOutput { return v.EgressEnabled }).(pulumi.BoolPtrOutput)
}

// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
func (o AviatrixFirenetOutput) EgressStaticCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.StringArrayOutput { return v.EgressStaticCidrs }).(pulumi.StringArrayOutput)
}

// Dynamic block of firewall instance(s) to be associated with the FireNet.
//
// Deprecated: Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
func (o AviatrixFirenetOutput) FirewallInstanceAssociations() AviatrixFirenetFirewallInstanceAssociationArrayOutput {
	return o.ApplyT(func(v *AviatrixFirenet) AviatrixFirenetFirewallInstanceAssociationArrayOutput {
		return v.FirewallInstanceAssociations
	}).(AviatrixFirenetFirewallInstanceAssociationArrayOutput)
}

// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
func (o AviatrixFirenetOutput) HashingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.StringPtrOutput { return v.HashingAlgorithm }).(pulumi.StringPtrOutput)
}

// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
func (o AviatrixFirenetOutput) InspectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.BoolPtrOutput { return v.InspectionEnabled }).(pulumi.BoolPtrOutput)
}

// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
func (o AviatrixFirenetOutput) KeepAliveViaLanInterfaceEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.BoolPtrOutput { return v.KeepAliveViaLanInterfaceEnabled }).(pulumi.BoolPtrOutput)
}

// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewallInstanceAssociation` blocks can be used. If set to false, all firewall associations must be managed via standalone `AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
func (o AviatrixFirenetOutput) ManageFirewallInstanceAssociation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.BoolPtrOutput { return v.ManageFirewallInstanceAssociation }).(pulumi.BoolPtrOutput)
}

// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
func (o AviatrixFirenetOutput) TgwSegmentationForEgressEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.BoolPtrOutput { return v.TgwSegmentationForEgressEnabled }).(pulumi.BoolPtrOutput)
}

// VPC ID of the Security VPC.
func (o AviatrixFirenetOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirenet) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type AviatrixFirenetArrayOutput struct{ *pulumi.OutputState }

func (AviatrixFirenetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixFirenet)(nil)).Elem()
}

func (o AviatrixFirenetArrayOutput) ToAviatrixFirenetArrayOutput() AviatrixFirenetArrayOutput {
	return o
}

func (o AviatrixFirenetArrayOutput) ToAviatrixFirenetArrayOutputWithContext(ctx context.Context) AviatrixFirenetArrayOutput {
	return o
}

func (o AviatrixFirenetArrayOutput) Index(i pulumi.IntInput) AviatrixFirenetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixFirenet {
		return vs[0].([]*AviatrixFirenet)[vs[1].(int)]
	}).(AviatrixFirenetOutput)
}

type AviatrixFirenetMapOutput struct{ *pulumi.OutputState }

func (AviatrixFirenetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixFirenet)(nil)).Elem()
}

func (o AviatrixFirenetMapOutput) ToAviatrixFirenetMapOutput() AviatrixFirenetMapOutput {
	return o
}

func (o AviatrixFirenetMapOutput) ToAviatrixFirenetMapOutputWithContext(ctx context.Context) AviatrixFirenetMapOutput {
	return o
}

func (o AviatrixFirenetMapOutput) MapIndex(k pulumi.StringInput) AviatrixFirenetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixFirenet {
		return vs[0].(map[string]*AviatrixFirenet)[vs[1].(string)]
	}).(AviatrixFirenetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirenetInput)(nil)).Elem(), &AviatrixFirenet{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirenetArrayInput)(nil)).Elem(), AviatrixFirenetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirenetMapInput)(nil)).Elem(), AviatrixFirenetMap{})
	pulumi.RegisterOutputType(AviatrixFirenetOutput{})
	pulumi.RegisterOutputType(AviatrixFirenetArrayOutput{})
	pulumi.RegisterOutputType(AviatrixFirenetMapOutput{})
}
