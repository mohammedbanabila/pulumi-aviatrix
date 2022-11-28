// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_private_mode_multicloud_endpoint** resource allows management of a Private Mode multicloud endpoint. This resource is available as of provider version R2.23+.
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
//			_, err := aviatrix.NewAviatrixPrivateModeMulticloudEndpoint(ctx, "test", &aviatrix.AviatrixPrivateModeMulticloudEndpointArgs{
//				AccountName:       pulumi.String("devops"),
//				ControllerLbVpcId: pulumi.String("vpc-abcdefg"),
//				Region:            pulumi.String("us-east-1"),
//				VpcId:             pulumi.String("vpc-abcdef"),
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
// **aviatrix_private_mode_multicloud_endpoint** can be imported using the `vpc_id`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixPrivateModeMulticloudEndpoint:AviatrixPrivateModeMulticloudEndpoint test vpc-1234567
//
// ```
type AviatrixPrivateModeMulticloudEndpoint struct {
	pulumi.CustomResourceState

	// Name of the access account.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// ID of the VPC containing a Private Mode controller load balancer.
	ControllerLbVpcId pulumi.StringOutput `pulumi:"controllerLbVpcId"`
	// DNS entry of the endpoint.
	DnsEntry pulumi.StringOutput `pulumi:"dnsEntry"`
	// Region of the VPC.
	Region pulumi.StringOutput `pulumi:"region"`
	// ID of the VPC to create the endpoint in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewAviatrixPrivateModeMulticloudEndpoint registers a new resource with the given unique name, arguments, and options.
func NewAviatrixPrivateModeMulticloudEndpoint(ctx *pulumi.Context,
	name string, args *AviatrixPrivateModeMulticloudEndpointArgs, opts ...pulumi.ResourceOption) (*AviatrixPrivateModeMulticloudEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ControllerLbVpcId == nil {
		return nil, errors.New("invalid value for required argument 'ControllerLbVpcId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixPrivateModeMulticloudEndpoint
	err := ctx.RegisterResource("aviatrix:index/aviatrixPrivateModeMulticloudEndpoint:AviatrixPrivateModeMulticloudEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixPrivateModeMulticloudEndpoint gets an existing AviatrixPrivateModeMulticloudEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixPrivateModeMulticloudEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixPrivateModeMulticloudEndpointState, opts ...pulumi.ResourceOption) (*AviatrixPrivateModeMulticloudEndpoint, error) {
	var resource AviatrixPrivateModeMulticloudEndpoint
	err := ctx.ReadResource("aviatrix:index/aviatrixPrivateModeMulticloudEndpoint:AviatrixPrivateModeMulticloudEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixPrivateModeMulticloudEndpoint resources.
type aviatrixPrivateModeMulticloudEndpointState struct {
	// Name of the access account.
	AccountName *string `pulumi:"accountName"`
	// ID of the VPC containing a Private Mode controller load balancer.
	ControllerLbVpcId *string `pulumi:"controllerLbVpcId"`
	// DNS entry of the endpoint.
	DnsEntry *string `pulumi:"dnsEntry"`
	// Region of the VPC.
	Region *string `pulumi:"region"`
	// ID of the VPC to create the endpoint in.
	VpcId *string `pulumi:"vpcId"`
}

type AviatrixPrivateModeMulticloudEndpointState struct {
	// Name of the access account.
	AccountName pulumi.StringPtrInput
	// ID of the VPC containing a Private Mode controller load balancer.
	ControllerLbVpcId pulumi.StringPtrInput
	// DNS entry of the endpoint.
	DnsEntry pulumi.StringPtrInput
	// Region of the VPC.
	Region pulumi.StringPtrInput
	// ID of the VPC to create the endpoint in.
	VpcId pulumi.StringPtrInput
}

func (AviatrixPrivateModeMulticloudEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixPrivateModeMulticloudEndpointState)(nil)).Elem()
}

type aviatrixPrivateModeMulticloudEndpointArgs struct {
	// Name of the access account.
	AccountName string `pulumi:"accountName"`
	// ID of the VPC containing a Private Mode controller load balancer.
	ControllerLbVpcId string `pulumi:"controllerLbVpcId"`
	// Region of the VPC.
	Region string `pulumi:"region"`
	// ID of the VPC to create the endpoint in.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a AviatrixPrivateModeMulticloudEndpoint resource.
type AviatrixPrivateModeMulticloudEndpointArgs struct {
	// Name of the access account.
	AccountName pulumi.StringInput
	// ID of the VPC containing a Private Mode controller load balancer.
	ControllerLbVpcId pulumi.StringInput
	// Region of the VPC.
	Region pulumi.StringInput
	// ID of the VPC to create the endpoint in.
	VpcId pulumi.StringInput
}

func (AviatrixPrivateModeMulticloudEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixPrivateModeMulticloudEndpointArgs)(nil)).Elem()
}

type AviatrixPrivateModeMulticloudEndpointInput interface {
	pulumi.Input

	ToAviatrixPrivateModeMulticloudEndpointOutput() AviatrixPrivateModeMulticloudEndpointOutput
	ToAviatrixPrivateModeMulticloudEndpointOutputWithContext(ctx context.Context) AviatrixPrivateModeMulticloudEndpointOutput
}

func (*AviatrixPrivateModeMulticloudEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixPrivateModeMulticloudEndpoint)(nil)).Elem()
}

func (i *AviatrixPrivateModeMulticloudEndpoint) ToAviatrixPrivateModeMulticloudEndpointOutput() AviatrixPrivateModeMulticloudEndpointOutput {
	return i.ToAviatrixPrivateModeMulticloudEndpointOutputWithContext(context.Background())
}

func (i *AviatrixPrivateModeMulticloudEndpoint) ToAviatrixPrivateModeMulticloudEndpointOutputWithContext(ctx context.Context) AviatrixPrivateModeMulticloudEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixPrivateModeMulticloudEndpointOutput)
}

// AviatrixPrivateModeMulticloudEndpointArrayInput is an input type that accepts AviatrixPrivateModeMulticloudEndpointArray and AviatrixPrivateModeMulticloudEndpointArrayOutput values.
// You can construct a concrete instance of `AviatrixPrivateModeMulticloudEndpointArrayInput` via:
//
//	AviatrixPrivateModeMulticloudEndpointArray{ AviatrixPrivateModeMulticloudEndpointArgs{...} }
type AviatrixPrivateModeMulticloudEndpointArrayInput interface {
	pulumi.Input

	ToAviatrixPrivateModeMulticloudEndpointArrayOutput() AviatrixPrivateModeMulticloudEndpointArrayOutput
	ToAviatrixPrivateModeMulticloudEndpointArrayOutputWithContext(context.Context) AviatrixPrivateModeMulticloudEndpointArrayOutput
}

type AviatrixPrivateModeMulticloudEndpointArray []AviatrixPrivateModeMulticloudEndpointInput

func (AviatrixPrivateModeMulticloudEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixPrivateModeMulticloudEndpoint)(nil)).Elem()
}

func (i AviatrixPrivateModeMulticloudEndpointArray) ToAviatrixPrivateModeMulticloudEndpointArrayOutput() AviatrixPrivateModeMulticloudEndpointArrayOutput {
	return i.ToAviatrixPrivateModeMulticloudEndpointArrayOutputWithContext(context.Background())
}

func (i AviatrixPrivateModeMulticloudEndpointArray) ToAviatrixPrivateModeMulticloudEndpointArrayOutputWithContext(ctx context.Context) AviatrixPrivateModeMulticloudEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixPrivateModeMulticloudEndpointArrayOutput)
}

// AviatrixPrivateModeMulticloudEndpointMapInput is an input type that accepts AviatrixPrivateModeMulticloudEndpointMap and AviatrixPrivateModeMulticloudEndpointMapOutput values.
// You can construct a concrete instance of `AviatrixPrivateModeMulticloudEndpointMapInput` via:
//
//	AviatrixPrivateModeMulticloudEndpointMap{ "key": AviatrixPrivateModeMulticloudEndpointArgs{...} }
type AviatrixPrivateModeMulticloudEndpointMapInput interface {
	pulumi.Input

	ToAviatrixPrivateModeMulticloudEndpointMapOutput() AviatrixPrivateModeMulticloudEndpointMapOutput
	ToAviatrixPrivateModeMulticloudEndpointMapOutputWithContext(context.Context) AviatrixPrivateModeMulticloudEndpointMapOutput
}

type AviatrixPrivateModeMulticloudEndpointMap map[string]AviatrixPrivateModeMulticloudEndpointInput

func (AviatrixPrivateModeMulticloudEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixPrivateModeMulticloudEndpoint)(nil)).Elem()
}

func (i AviatrixPrivateModeMulticloudEndpointMap) ToAviatrixPrivateModeMulticloudEndpointMapOutput() AviatrixPrivateModeMulticloudEndpointMapOutput {
	return i.ToAviatrixPrivateModeMulticloudEndpointMapOutputWithContext(context.Background())
}

func (i AviatrixPrivateModeMulticloudEndpointMap) ToAviatrixPrivateModeMulticloudEndpointMapOutputWithContext(ctx context.Context) AviatrixPrivateModeMulticloudEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixPrivateModeMulticloudEndpointMapOutput)
}

type AviatrixPrivateModeMulticloudEndpointOutput struct{ *pulumi.OutputState }

func (AviatrixPrivateModeMulticloudEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixPrivateModeMulticloudEndpoint)(nil)).Elem()
}

func (o AviatrixPrivateModeMulticloudEndpointOutput) ToAviatrixPrivateModeMulticloudEndpointOutput() AviatrixPrivateModeMulticloudEndpointOutput {
	return o
}

func (o AviatrixPrivateModeMulticloudEndpointOutput) ToAviatrixPrivateModeMulticloudEndpointOutputWithContext(ctx context.Context) AviatrixPrivateModeMulticloudEndpointOutput {
	return o
}

// Name of the access account.
func (o AviatrixPrivateModeMulticloudEndpointOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixPrivateModeMulticloudEndpoint) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// ID of the VPC containing a Private Mode controller load balancer.
func (o AviatrixPrivateModeMulticloudEndpointOutput) ControllerLbVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixPrivateModeMulticloudEndpoint) pulumi.StringOutput { return v.ControllerLbVpcId }).(pulumi.StringOutput)
}

// DNS entry of the endpoint.
func (o AviatrixPrivateModeMulticloudEndpointOutput) DnsEntry() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixPrivateModeMulticloudEndpoint) pulumi.StringOutput { return v.DnsEntry }).(pulumi.StringOutput)
}

// Region of the VPC.
func (o AviatrixPrivateModeMulticloudEndpointOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixPrivateModeMulticloudEndpoint) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// ID of the VPC to create the endpoint in.
func (o AviatrixPrivateModeMulticloudEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixPrivateModeMulticloudEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type AviatrixPrivateModeMulticloudEndpointArrayOutput struct{ *pulumi.OutputState }

func (AviatrixPrivateModeMulticloudEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixPrivateModeMulticloudEndpoint)(nil)).Elem()
}

func (o AviatrixPrivateModeMulticloudEndpointArrayOutput) ToAviatrixPrivateModeMulticloudEndpointArrayOutput() AviatrixPrivateModeMulticloudEndpointArrayOutput {
	return o
}

func (o AviatrixPrivateModeMulticloudEndpointArrayOutput) ToAviatrixPrivateModeMulticloudEndpointArrayOutputWithContext(ctx context.Context) AviatrixPrivateModeMulticloudEndpointArrayOutput {
	return o
}

func (o AviatrixPrivateModeMulticloudEndpointArrayOutput) Index(i pulumi.IntInput) AviatrixPrivateModeMulticloudEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixPrivateModeMulticloudEndpoint {
		return vs[0].([]*AviatrixPrivateModeMulticloudEndpoint)[vs[1].(int)]
	}).(AviatrixPrivateModeMulticloudEndpointOutput)
}

type AviatrixPrivateModeMulticloudEndpointMapOutput struct{ *pulumi.OutputState }

func (AviatrixPrivateModeMulticloudEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixPrivateModeMulticloudEndpoint)(nil)).Elem()
}

func (o AviatrixPrivateModeMulticloudEndpointMapOutput) ToAviatrixPrivateModeMulticloudEndpointMapOutput() AviatrixPrivateModeMulticloudEndpointMapOutput {
	return o
}

func (o AviatrixPrivateModeMulticloudEndpointMapOutput) ToAviatrixPrivateModeMulticloudEndpointMapOutputWithContext(ctx context.Context) AviatrixPrivateModeMulticloudEndpointMapOutput {
	return o
}

func (o AviatrixPrivateModeMulticloudEndpointMapOutput) MapIndex(k pulumi.StringInput) AviatrixPrivateModeMulticloudEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixPrivateModeMulticloudEndpoint {
		return vs[0].(map[string]*AviatrixPrivateModeMulticloudEndpoint)[vs[1].(string)]
	}).(AviatrixPrivateModeMulticloudEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixPrivateModeMulticloudEndpointInput)(nil)).Elem(), &AviatrixPrivateModeMulticloudEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixPrivateModeMulticloudEndpointArrayInput)(nil)).Elem(), AviatrixPrivateModeMulticloudEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixPrivateModeMulticloudEndpointMapInput)(nil)).Elem(), AviatrixPrivateModeMulticloudEndpointMap{})
	pulumi.RegisterOutputType(AviatrixPrivateModeMulticloudEndpointOutput{})
	pulumi.RegisterOutputType(AviatrixPrivateModeMulticloudEndpointArrayOutput{})
	pulumi.RegisterOutputType(AviatrixPrivateModeMulticloudEndpointMapOutput{})
}
