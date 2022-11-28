// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_geo_vpn** resource enables and manages the [Aviatrix Geo VPN feature](https://docs.aviatrix.com/HowTos/GeoVPN.html).
//
// > **NOTE:** If ELBs/gateways are being managed by the Geo VPN, in order to update VPN configurations of the Geo VPN, all the VPN configurations of the ELBs/gateways must be updated simultaneously and share the same values. This can be achieved by managing the VPN configurations through variables and updating their values accordingly.
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
//			_, err := aviatrix.NewAviatrixGeoVpn(ctx, "testGeoVpn", &aviatrix.AviatrixGeoVpnArgs{
//				AccountName: pulumi.String("devops-aws"),
//				CloudType:   pulumi.Int(1),
//				DomainName:  pulumi.String("aviatrix.live"),
//				ElbDnsNames: pulumi.StringArray{
//					pulumi.String("elb-test1-497f5e89.elb.us-west-1.amazonaws.com"),
//					pulumi.String("elb-test2-974f895e.elb.us-east-2.amazonaws.com"),
//				},
//				ServiceName: pulumi.String("vpn"),
//			})
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
// **geo_vpn** can be imported using the `service_name` and `domain_name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn test service_name~domain_name
//
// ```
type AviatrixGeoVpn struct {
	pulumi.CustomResourceState

	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
	CloudType pulumi.IntOutput `pulumi:"cloudType"`
	// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// List of ELB names to attach to this Geo VPN name.
	ElbDnsNames pulumi.StringArrayOutput `pulumi:"elbDnsNames"`
	// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewAviatrixGeoVpn registers a new resource with the given unique name, arguments, and options.
func NewAviatrixGeoVpn(ctx *pulumi.Context,
	name string, args *AviatrixGeoVpnArgs, opts ...pulumi.ResourceOption) (*AviatrixGeoVpn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.CloudType == nil {
		return nil, errors.New("invalid value for required argument 'CloudType'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ElbDnsNames == nil {
		return nil, errors.New("invalid value for required argument 'ElbDnsNames'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixGeoVpn
	err := ctx.RegisterResource("aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixGeoVpn gets an existing AviatrixGeoVpn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixGeoVpn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixGeoVpnState, opts ...pulumi.ResourceOption) (*AviatrixGeoVpn, error) {
	var resource AviatrixGeoVpn
	err := ctx.ReadResource("aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixGeoVpn resources.
type aviatrixGeoVpnState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName *string `pulumi:"accountName"`
	// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
	CloudType *int `pulumi:"cloudType"`
	// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
	DomainName *string `pulumi:"domainName"`
	// List of ELB names to attach to this Geo VPN name.
	ElbDnsNames []string `pulumi:"elbDnsNames"`
	// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
	ServiceName *string `pulumi:"serviceName"`
}

type AviatrixGeoVpnState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringPtrInput
	// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
	CloudType pulumi.IntPtrInput
	// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
	DomainName pulumi.StringPtrInput
	// List of ELB names to attach to this Geo VPN name.
	ElbDnsNames pulumi.StringArrayInput
	// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
	ServiceName pulumi.StringPtrInput
}

func (AviatrixGeoVpnState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixGeoVpnState)(nil)).Elem()
}

type aviatrixGeoVpnArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName string `pulumi:"accountName"`
	// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
	CloudType int `pulumi:"cloudType"`
	// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
	DomainName string `pulumi:"domainName"`
	// List of ELB names to attach to this Geo VPN name.
	ElbDnsNames []string `pulumi:"elbDnsNames"`
	// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a AviatrixGeoVpn resource.
type AviatrixGeoVpnArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringInput
	// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
	CloudType pulumi.IntInput
	// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
	DomainName pulumi.StringInput
	// List of ELB names to attach to this Geo VPN name.
	ElbDnsNames pulumi.StringArrayInput
	// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
	ServiceName pulumi.StringInput
}

func (AviatrixGeoVpnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixGeoVpnArgs)(nil)).Elem()
}

type AviatrixGeoVpnInput interface {
	pulumi.Input

	ToAviatrixGeoVpnOutput() AviatrixGeoVpnOutput
	ToAviatrixGeoVpnOutputWithContext(ctx context.Context) AviatrixGeoVpnOutput
}

func (*AviatrixGeoVpn) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixGeoVpn)(nil)).Elem()
}

func (i *AviatrixGeoVpn) ToAviatrixGeoVpnOutput() AviatrixGeoVpnOutput {
	return i.ToAviatrixGeoVpnOutputWithContext(context.Background())
}

func (i *AviatrixGeoVpn) ToAviatrixGeoVpnOutputWithContext(ctx context.Context) AviatrixGeoVpnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixGeoVpnOutput)
}

// AviatrixGeoVpnArrayInput is an input type that accepts AviatrixGeoVpnArray and AviatrixGeoVpnArrayOutput values.
// You can construct a concrete instance of `AviatrixGeoVpnArrayInput` via:
//
//	AviatrixGeoVpnArray{ AviatrixGeoVpnArgs{...} }
type AviatrixGeoVpnArrayInput interface {
	pulumi.Input

	ToAviatrixGeoVpnArrayOutput() AviatrixGeoVpnArrayOutput
	ToAviatrixGeoVpnArrayOutputWithContext(context.Context) AviatrixGeoVpnArrayOutput
}

type AviatrixGeoVpnArray []AviatrixGeoVpnInput

func (AviatrixGeoVpnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixGeoVpn)(nil)).Elem()
}

func (i AviatrixGeoVpnArray) ToAviatrixGeoVpnArrayOutput() AviatrixGeoVpnArrayOutput {
	return i.ToAviatrixGeoVpnArrayOutputWithContext(context.Background())
}

func (i AviatrixGeoVpnArray) ToAviatrixGeoVpnArrayOutputWithContext(ctx context.Context) AviatrixGeoVpnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixGeoVpnArrayOutput)
}

// AviatrixGeoVpnMapInput is an input type that accepts AviatrixGeoVpnMap and AviatrixGeoVpnMapOutput values.
// You can construct a concrete instance of `AviatrixGeoVpnMapInput` via:
//
//	AviatrixGeoVpnMap{ "key": AviatrixGeoVpnArgs{...} }
type AviatrixGeoVpnMapInput interface {
	pulumi.Input

	ToAviatrixGeoVpnMapOutput() AviatrixGeoVpnMapOutput
	ToAviatrixGeoVpnMapOutputWithContext(context.Context) AviatrixGeoVpnMapOutput
}

type AviatrixGeoVpnMap map[string]AviatrixGeoVpnInput

func (AviatrixGeoVpnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixGeoVpn)(nil)).Elem()
}

func (i AviatrixGeoVpnMap) ToAviatrixGeoVpnMapOutput() AviatrixGeoVpnMapOutput {
	return i.ToAviatrixGeoVpnMapOutputWithContext(context.Background())
}

func (i AviatrixGeoVpnMap) ToAviatrixGeoVpnMapOutputWithContext(ctx context.Context) AviatrixGeoVpnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixGeoVpnMapOutput)
}

type AviatrixGeoVpnOutput struct{ *pulumi.OutputState }

func (AviatrixGeoVpnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixGeoVpn)(nil)).Elem()
}

func (o AviatrixGeoVpnOutput) ToAviatrixGeoVpnOutput() AviatrixGeoVpnOutput {
	return o
}

func (o AviatrixGeoVpnOutput) ToAviatrixGeoVpnOutputWithContext(ctx context.Context) AviatrixGeoVpnOutput {
	return o
}

// This parameter represents the name of a Cloud-Account in Aviatrix controller.
func (o AviatrixGeoVpnOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixGeoVpn) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
func (o AviatrixGeoVpnOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v *AviatrixGeoVpn) pulumi.IntOutput { return v.CloudType }).(pulumi.IntOutput)
}

// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
func (o AviatrixGeoVpnOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixGeoVpn) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// List of ELB names to attach to this Geo VPN name.
func (o AviatrixGeoVpnOutput) ElbDnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixGeoVpn) pulumi.StringArrayOutput { return v.ElbDnsNames }).(pulumi.StringArrayOutput)
}

// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
func (o AviatrixGeoVpnOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixGeoVpn) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type AviatrixGeoVpnArrayOutput struct{ *pulumi.OutputState }

func (AviatrixGeoVpnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixGeoVpn)(nil)).Elem()
}

func (o AviatrixGeoVpnArrayOutput) ToAviatrixGeoVpnArrayOutput() AviatrixGeoVpnArrayOutput {
	return o
}

func (o AviatrixGeoVpnArrayOutput) ToAviatrixGeoVpnArrayOutputWithContext(ctx context.Context) AviatrixGeoVpnArrayOutput {
	return o
}

func (o AviatrixGeoVpnArrayOutput) Index(i pulumi.IntInput) AviatrixGeoVpnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixGeoVpn {
		return vs[0].([]*AviatrixGeoVpn)[vs[1].(int)]
	}).(AviatrixGeoVpnOutput)
}

type AviatrixGeoVpnMapOutput struct{ *pulumi.OutputState }

func (AviatrixGeoVpnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixGeoVpn)(nil)).Elem()
}

func (o AviatrixGeoVpnMapOutput) ToAviatrixGeoVpnMapOutput() AviatrixGeoVpnMapOutput {
	return o
}

func (o AviatrixGeoVpnMapOutput) ToAviatrixGeoVpnMapOutputWithContext(ctx context.Context) AviatrixGeoVpnMapOutput {
	return o
}

func (o AviatrixGeoVpnMapOutput) MapIndex(k pulumi.StringInput) AviatrixGeoVpnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixGeoVpn {
		return vs[0].(map[string]*AviatrixGeoVpn)[vs[1].(string)]
	}).(AviatrixGeoVpnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixGeoVpnInput)(nil)).Elem(), &AviatrixGeoVpn{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixGeoVpnArrayInput)(nil)).Elem(), AviatrixGeoVpnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixGeoVpnMapInput)(nil)).Elem(), AviatrixGeoVpnMap{})
	pulumi.RegisterOutputType(AviatrixGeoVpnOutput{})
	pulumi.RegisterOutputType(AviatrixGeoVpnArrayOutput{})
	pulumi.RegisterOutputType(AviatrixGeoVpnMapOutput{})
}
