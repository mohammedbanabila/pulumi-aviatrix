// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_aws_tgw_network_domain** resource allows the creation and management of Aviatrix network domains.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/astipkovits/pulumi-aviatrix/sdk/go/aviatrix"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testAwsTgw, err := aviatrix.NewAviatrixAwsTgw(ctx, "testAwsTgw", &aviatrix.AviatrixAwsTgwArgs{
//				AccountName:                    pulumi.String("devops"),
//				AwsSideAsNumber:                pulumi.String("64512"),
//				Region:                         pulumi.String("us-east-1"),
//				TgwName:                        pulumi.String("test-AWS-TGW"),
//				ManageSecurityDomain:           pulumi.Bool(false),
//				ManageVpcAttachment:            pulumi.Bool(false),
//				ManageTransitGatewayAttachment: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			defaultDomain, err := aviatrix.NewAviatrixAwsTgwNetworkDomain(ctx, "defaultDomain", &aviatrix.AviatrixAwsTgwNetworkDomainArgs{
//				TgwName: testAwsTgw.TgwName,
//			})
//			if err != nil {
//				return err
//			}
//			sharedServiceDomain, err := aviatrix.NewAviatrixAwsTgwNetworkDomain(ctx, "sharedServiceDomain", &aviatrix.AviatrixAwsTgwNetworkDomainArgs{
//				TgwName: testAwsTgw.TgwName,
//			})
//			if err != nil {
//				return err
//			}
//			aviatrixEdgeDomain, err := aviatrix.NewAviatrixAwsTgwNetworkDomain(ctx, "aviatrixEdgeDomain", &aviatrix.AviatrixAwsTgwNetworkDomainArgs{
//				TgwName: testAwsTgw.TgwName,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aviatrix.NewAviatrixAwsTgwSecurityDomainConn(ctx, "defaultSdConn1", &aviatrix.AviatrixAwsTgwSecurityDomainConnArgs{
//				TgwName:     testAwsTgw.TgwName,
//				DomainName1: aviatrixEdgeDomain.Name,
//				DomainName2: defaultDomain.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aviatrix.NewAviatrixAwsTgwSecurityDomainConn(ctx, "defaultSdConn2", &aviatrix.AviatrixAwsTgwSecurityDomainConnArgs{
//				TgwName:     testAwsTgw.TgwName,
//				DomainName1: aviatrixEdgeDomain.Name,
//				DomainName2: sharedServiceDomain.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aviatrix.NewAviatrixAwsTgwSecurityDomainConn(ctx, "defaultSdConn3", &aviatrix.AviatrixAwsTgwSecurityDomainConnArgs{
//				TgwName:     testAwsTgw.TgwName,
//				DomainName1: defaultDomain.Name,
//				DomainName2: sharedServiceDomain.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aviatrix.NewAviatrixAwsTgwNetworkDomain(ctx, "test", &aviatrix.AviatrixAwsTgwNetworkDomainArgs{
//				TgwName: testAwsTgw.TgwName,
//			}, pulumi.DependsOn([]pulumi.Resource{
//				defaultDomain,
//				sharedServiceDomain,
//				aviatrixEdgeDomain,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// **aws_tgw_network_domain** can be imported using the `name` and `tgw_name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain test tgw_name~name
//
// ```
type AviatrixAwsTgwNetworkDomain struct {
	pulumi.CustomResourceState

	// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
	AviatrixFirewall pulumi.BoolPtrOutput `pulumi:"aviatrixFirewall"`
	// The name of the network domain.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
	NativeEgress pulumi.BoolPtrOutput `pulumi:"nativeEgress"`
	// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
	NativeFirewall pulumi.BoolPtrOutput `pulumi:"nativeFirewall"`
	// The AWS TGW name of the network domain.
	TgwName pulumi.StringOutput `pulumi:"tgwName"`
}

// NewAviatrixAwsTgwNetworkDomain registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAwsTgwNetworkDomain(ctx *pulumi.Context,
	name string, args *AviatrixAwsTgwNetworkDomainArgs, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwNetworkDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TgwName == nil {
		return nil, errors.New("invalid value for required argument 'TgwName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAwsTgwNetworkDomain
	err := ctx.RegisterResource("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAwsTgwNetworkDomain gets an existing AviatrixAwsTgwNetworkDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAwsTgwNetworkDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAwsTgwNetworkDomainState, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwNetworkDomain, error) {
	var resource AviatrixAwsTgwNetworkDomain
	err := ctx.ReadResource("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAwsTgwNetworkDomain resources.
type aviatrixAwsTgwNetworkDomainState struct {
	// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
	AviatrixFirewall *bool `pulumi:"aviatrixFirewall"`
	// The name of the network domain.
	Name *string `pulumi:"name"`
	// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
	NativeEgress *bool `pulumi:"nativeEgress"`
	// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
	NativeFirewall *bool `pulumi:"nativeFirewall"`
	// The AWS TGW name of the network domain.
	TgwName *string `pulumi:"tgwName"`
}

type AviatrixAwsTgwNetworkDomainState struct {
	// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
	AviatrixFirewall pulumi.BoolPtrInput
	// The name of the network domain.
	Name pulumi.StringPtrInput
	// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
	NativeEgress pulumi.BoolPtrInput
	// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
	NativeFirewall pulumi.BoolPtrInput
	// The AWS TGW name of the network domain.
	TgwName pulumi.StringPtrInput
}

func (AviatrixAwsTgwNetworkDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwNetworkDomainState)(nil)).Elem()
}

type aviatrixAwsTgwNetworkDomainArgs struct {
	// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
	AviatrixFirewall *bool `pulumi:"aviatrixFirewall"`
	// The name of the network domain.
	Name *string `pulumi:"name"`
	// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
	NativeEgress *bool `pulumi:"nativeEgress"`
	// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
	NativeFirewall *bool `pulumi:"nativeFirewall"`
	// The AWS TGW name of the network domain.
	TgwName string `pulumi:"tgwName"`
}

// The set of arguments for constructing a AviatrixAwsTgwNetworkDomain resource.
type AviatrixAwsTgwNetworkDomainArgs struct {
	// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
	AviatrixFirewall pulumi.BoolPtrInput
	// The name of the network domain.
	Name pulumi.StringPtrInput
	// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
	NativeEgress pulumi.BoolPtrInput
	// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
	NativeFirewall pulumi.BoolPtrInput
	// The AWS TGW name of the network domain.
	TgwName pulumi.StringInput
}

func (AviatrixAwsTgwNetworkDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwNetworkDomainArgs)(nil)).Elem()
}

type AviatrixAwsTgwNetworkDomainInput interface {
	pulumi.Input

	ToAviatrixAwsTgwNetworkDomainOutput() AviatrixAwsTgwNetworkDomainOutput
	ToAviatrixAwsTgwNetworkDomainOutputWithContext(ctx context.Context) AviatrixAwsTgwNetworkDomainOutput
}

func (*AviatrixAwsTgwNetworkDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwNetworkDomain)(nil)).Elem()
}

func (i *AviatrixAwsTgwNetworkDomain) ToAviatrixAwsTgwNetworkDomainOutput() AviatrixAwsTgwNetworkDomainOutput {
	return i.ToAviatrixAwsTgwNetworkDomainOutputWithContext(context.Background())
}

func (i *AviatrixAwsTgwNetworkDomain) ToAviatrixAwsTgwNetworkDomainOutputWithContext(ctx context.Context) AviatrixAwsTgwNetworkDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwNetworkDomainOutput)
}

// AviatrixAwsTgwNetworkDomainArrayInput is an input type that accepts AviatrixAwsTgwNetworkDomainArray and AviatrixAwsTgwNetworkDomainArrayOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwNetworkDomainArrayInput` via:
//
//	AviatrixAwsTgwNetworkDomainArray{ AviatrixAwsTgwNetworkDomainArgs{...} }
type AviatrixAwsTgwNetworkDomainArrayInput interface {
	pulumi.Input

	ToAviatrixAwsTgwNetworkDomainArrayOutput() AviatrixAwsTgwNetworkDomainArrayOutput
	ToAviatrixAwsTgwNetworkDomainArrayOutputWithContext(context.Context) AviatrixAwsTgwNetworkDomainArrayOutput
}

type AviatrixAwsTgwNetworkDomainArray []AviatrixAwsTgwNetworkDomainInput

func (AviatrixAwsTgwNetworkDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwNetworkDomain)(nil)).Elem()
}

func (i AviatrixAwsTgwNetworkDomainArray) ToAviatrixAwsTgwNetworkDomainArrayOutput() AviatrixAwsTgwNetworkDomainArrayOutput {
	return i.ToAviatrixAwsTgwNetworkDomainArrayOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwNetworkDomainArray) ToAviatrixAwsTgwNetworkDomainArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwNetworkDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwNetworkDomainArrayOutput)
}

// AviatrixAwsTgwNetworkDomainMapInput is an input type that accepts AviatrixAwsTgwNetworkDomainMap and AviatrixAwsTgwNetworkDomainMapOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwNetworkDomainMapInput` via:
//
//	AviatrixAwsTgwNetworkDomainMap{ "key": AviatrixAwsTgwNetworkDomainArgs{...} }
type AviatrixAwsTgwNetworkDomainMapInput interface {
	pulumi.Input

	ToAviatrixAwsTgwNetworkDomainMapOutput() AviatrixAwsTgwNetworkDomainMapOutput
	ToAviatrixAwsTgwNetworkDomainMapOutputWithContext(context.Context) AviatrixAwsTgwNetworkDomainMapOutput
}

type AviatrixAwsTgwNetworkDomainMap map[string]AviatrixAwsTgwNetworkDomainInput

func (AviatrixAwsTgwNetworkDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwNetworkDomain)(nil)).Elem()
}

func (i AviatrixAwsTgwNetworkDomainMap) ToAviatrixAwsTgwNetworkDomainMapOutput() AviatrixAwsTgwNetworkDomainMapOutput {
	return i.ToAviatrixAwsTgwNetworkDomainMapOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwNetworkDomainMap) ToAviatrixAwsTgwNetworkDomainMapOutputWithContext(ctx context.Context) AviatrixAwsTgwNetworkDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwNetworkDomainMapOutput)
}

type AviatrixAwsTgwNetworkDomainOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwNetworkDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwNetworkDomain)(nil)).Elem()
}

func (o AviatrixAwsTgwNetworkDomainOutput) ToAviatrixAwsTgwNetworkDomainOutput() AviatrixAwsTgwNetworkDomainOutput {
	return o
}

func (o AviatrixAwsTgwNetworkDomainOutput) ToAviatrixAwsTgwNetworkDomainOutputWithContext(ctx context.Context) AviatrixAwsTgwNetworkDomainOutput {
	return o
}

// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
func (o AviatrixAwsTgwNetworkDomainOutput) AviatrixFirewall() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwNetworkDomain) pulumi.BoolPtrOutput { return v.AviatrixFirewall }).(pulumi.BoolPtrOutput)
}

// The name of the network domain.
func (o AviatrixAwsTgwNetworkDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwNetworkDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
func (o AviatrixAwsTgwNetworkDomainOutput) NativeEgress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwNetworkDomain) pulumi.BoolPtrOutput { return v.NativeEgress }).(pulumi.BoolPtrOutput)
}

// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
func (o AviatrixAwsTgwNetworkDomainOutput) NativeFirewall() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwNetworkDomain) pulumi.BoolPtrOutput { return v.NativeFirewall }).(pulumi.BoolPtrOutput)
}

// The AWS TGW name of the network domain.
func (o AviatrixAwsTgwNetworkDomainOutput) TgwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwNetworkDomain) pulumi.StringOutput { return v.TgwName }).(pulumi.StringOutput)
}

type AviatrixAwsTgwNetworkDomainArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwNetworkDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwNetworkDomain)(nil)).Elem()
}

func (o AviatrixAwsTgwNetworkDomainArrayOutput) ToAviatrixAwsTgwNetworkDomainArrayOutput() AviatrixAwsTgwNetworkDomainArrayOutput {
	return o
}

func (o AviatrixAwsTgwNetworkDomainArrayOutput) ToAviatrixAwsTgwNetworkDomainArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwNetworkDomainArrayOutput {
	return o
}

func (o AviatrixAwsTgwNetworkDomainArrayOutput) Index(i pulumi.IntInput) AviatrixAwsTgwNetworkDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAwsTgwNetworkDomain {
		return vs[0].([]*AviatrixAwsTgwNetworkDomain)[vs[1].(int)]
	}).(AviatrixAwsTgwNetworkDomainOutput)
}

type AviatrixAwsTgwNetworkDomainMapOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwNetworkDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwNetworkDomain)(nil)).Elem()
}

func (o AviatrixAwsTgwNetworkDomainMapOutput) ToAviatrixAwsTgwNetworkDomainMapOutput() AviatrixAwsTgwNetworkDomainMapOutput {
	return o
}

func (o AviatrixAwsTgwNetworkDomainMapOutput) ToAviatrixAwsTgwNetworkDomainMapOutputWithContext(ctx context.Context) AviatrixAwsTgwNetworkDomainMapOutput {
	return o
}

func (o AviatrixAwsTgwNetworkDomainMapOutput) MapIndex(k pulumi.StringInput) AviatrixAwsTgwNetworkDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAwsTgwNetworkDomain {
		return vs[0].(map[string]*AviatrixAwsTgwNetworkDomain)[vs[1].(string)]
	}).(AviatrixAwsTgwNetworkDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwNetworkDomainInput)(nil)).Elem(), &AviatrixAwsTgwNetworkDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwNetworkDomainArrayInput)(nil)).Elem(), AviatrixAwsTgwNetworkDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwNetworkDomainMapInput)(nil)).Elem(), AviatrixAwsTgwNetworkDomainMap{})
	pulumi.RegisterOutputType(AviatrixAwsTgwNetworkDomainOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwNetworkDomainArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwNetworkDomainMapOutput{})
}
