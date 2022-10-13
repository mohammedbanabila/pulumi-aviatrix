// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixTransPeer struct {
	pulumi.CustomResourceState

	// Name of nexthop gateway.
	Nexthop pulumi.StringOutput `pulumi:"nexthop"`
	// Destination CIDR.
	ReachableCidr pulumi.StringOutput `pulumi:"reachableCidr"`
	// Name of Source gateway.
	Source pulumi.StringOutput `pulumi:"source"`
}

// NewAviatrixTransPeer registers a new resource with the given unique name, arguments, and options.
func NewAviatrixTransPeer(ctx *pulumi.Context,
	name string, args *AviatrixTransPeerArgs, opts ...pulumi.ResourceOption) (*AviatrixTransPeer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Nexthop == nil {
		return nil, errors.New("invalid value for required argument 'Nexthop'")
	}
	if args.ReachableCidr == nil {
		return nil, errors.New("invalid value for required argument 'ReachableCidr'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixTransPeer
	err := ctx.RegisterResource("aviatrix:index/aviatrixTransPeer:AviatrixTransPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixTransPeer gets an existing AviatrixTransPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixTransPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixTransPeerState, opts ...pulumi.ResourceOption) (*AviatrixTransPeer, error) {
	var resource AviatrixTransPeer
	err := ctx.ReadResource("aviatrix:index/aviatrixTransPeer:AviatrixTransPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixTransPeer resources.
type aviatrixTransPeerState struct {
	// Name of nexthop gateway.
	Nexthop *string `pulumi:"nexthop"`
	// Destination CIDR.
	ReachableCidr *string `pulumi:"reachableCidr"`
	// Name of Source gateway.
	Source *string `pulumi:"source"`
}

type AviatrixTransPeerState struct {
	// Name of nexthop gateway.
	Nexthop pulumi.StringPtrInput
	// Destination CIDR.
	ReachableCidr pulumi.StringPtrInput
	// Name of Source gateway.
	Source pulumi.StringPtrInput
}

func (AviatrixTransPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixTransPeerState)(nil)).Elem()
}

type aviatrixTransPeerArgs struct {
	// Name of nexthop gateway.
	Nexthop string `pulumi:"nexthop"`
	// Destination CIDR.
	ReachableCidr string `pulumi:"reachableCidr"`
	// Name of Source gateway.
	Source string `pulumi:"source"`
}

// The set of arguments for constructing a AviatrixTransPeer resource.
type AviatrixTransPeerArgs struct {
	// Name of nexthop gateway.
	Nexthop pulumi.StringInput
	// Destination CIDR.
	ReachableCidr pulumi.StringInput
	// Name of Source gateway.
	Source pulumi.StringInput
}

func (AviatrixTransPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixTransPeerArgs)(nil)).Elem()
}

type AviatrixTransPeerInput interface {
	pulumi.Input

	ToAviatrixTransPeerOutput() AviatrixTransPeerOutput
	ToAviatrixTransPeerOutputWithContext(ctx context.Context) AviatrixTransPeerOutput
}

func (*AviatrixTransPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixTransPeer)(nil)).Elem()
}

func (i *AviatrixTransPeer) ToAviatrixTransPeerOutput() AviatrixTransPeerOutput {
	return i.ToAviatrixTransPeerOutputWithContext(context.Background())
}

func (i *AviatrixTransPeer) ToAviatrixTransPeerOutputWithContext(ctx context.Context) AviatrixTransPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTransPeerOutput)
}

// AviatrixTransPeerArrayInput is an input type that accepts AviatrixTransPeerArray and AviatrixTransPeerArrayOutput values.
// You can construct a concrete instance of `AviatrixTransPeerArrayInput` via:
//
//	AviatrixTransPeerArray{ AviatrixTransPeerArgs{...} }
type AviatrixTransPeerArrayInput interface {
	pulumi.Input

	ToAviatrixTransPeerArrayOutput() AviatrixTransPeerArrayOutput
	ToAviatrixTransPeerArrayOutputWithContext(context.Context) AviatrixTransPeerArrayOutput
}

type AviatrixTransPeerArray []AviatrixTransPeerInput

func (AviatrixTransPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixTransPeer)(nil)).Elem()
}

func (i AviatrixTransPeerArray) ToAviatrixTransPeerArrayOutput() AviatrixTransPeerArrayOutput {
	return i.ToAviatrixTransPeerArrayOutputWithContext(context.Background())
}

func (i AviatrixTransPeerArray) ToAviatrixTransPeerArrayOutputWithContext(ctx context.Context) AviatrixTransPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTransPeerArrayOutput)
}

// AviatrixTransPeerMapInput is an input type that accepts AviatrixTransPeerMap and AviatrixTransPeerMapOutput values.
// You can construct a concrete instance of `AviatrixTransPeerMapInput` via:
//
//	AviatrixTransPeerMap{ "key": AviatrixTransPeerArgs{...} }
type AviatrixTransPeerMapInput interface {
	pulumi.Input

	ToAviatrixTransPeerMapOutput() AviatrixTransPeerMapOutput
	ToAviatrixTransPeerMapOutputWithContext(context.Context) AviatrixTransPeerMapOutput
}

type AviatrixTransPeerMap map[string]AviatrixTransPeerInput

func (AviatrixTransPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixTransPeer)(nil)).Elem()
}

func (i AviatrixTransPeerMap) ToAviatrixTransPeerMapOutput() AviatrixTransPeerMapOutput {
	return i.ToAviatrixTransPeerMapOutputWithContext(context.Background())
}

func (i AviatrixTransPeerMap) ToAviatrixTransPeerMapOutputWithContext(ctx context.Context) AviatrixTransPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTransPeerMapOutput)
}

type AviatrixTransPeerOutput struct{ *pulumi.OutputState }

func (AviatrixTransPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixTransPeer)(nil)).Elem()
}

func (o AviatrixTransPeerOutput) ToAviatrixTransPeerOutput() AviatrixTransPeerOutput {
	return o
}

func (o AviatrixTransPeerOutput) ToAviatrixTransPeerOutputWithContext(ctx context.Context) AviatrixTransPeerOutput {
	return o
}

// Name of nexthop gateway.
func (o AviatrixTransPeerOutput) Nexthop() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransPeer) pulumi.StringOutput { return v.Nexthop }).(pulumi.StringOutput)
}

// Destination CIDR.
func (o AviatrixTransPeerOutput) ReachableCidr() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransPeer) pulumi.StringOutput { return v.ReachableCidr }).(pulumi.StringOutput)
}

// Name of Source gateway.
func (o AviatrixTransPeerOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransPeer) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

type AviatrixTransPeerArrayOutput struct{ *pulumi.OutputState }

func (AviatrixTransPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixTransPeer)(nil)).Elem()
}

func (o AviatrixTransPeerArrayOutput) ToAviatrixTransPeerArrayOutput() AviatrixTransPeerArrayOutput {
	return o
}

func (o AviatrixTransPeerArrayOutput) ToAviatrixTransPeerArrayOutputWithContext(ctx context.Context) AviatrixTransPeerArrayOutput {
	return o
}

func (o AviatrixTransPeerArrayOutput) Index(i pulumi.IntInput) AviatrixTransPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixTransPeer {
		return vs[0].([]*AviatrixTransPeer)[vs[1].(int)]
	}).(AviatrixTransPeerOutput)
}

type AviatrixTransPeerMapOutput struct{ *pulumi.OutputState }

func (AviatrixTransPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixTransPeer)(nil)).Elem()
}

func (o AviatrixTransPeerMapOutput) ToAviatrixTransPeerMapOutput() AviatrixTransPeerMapOutput {
	return o
}

func (o AviatrixTransPeerMapOutput) ToAviatrixTransPeerMapOutputWithContext(ctx context.Context) AviatrixTransPeerMapOutput {
	return o
}

func (o AviatrixTransPeerMapOutput) MapIndex(k pulumi.StringInput) AviatrixTransPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixTransPeer {
		return vs[0].(map[string]*AviatrixTransPeer)[vs[1].(string)]
	}).(AviatrixTransPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTransPeerInput)(nil)).Elem(), &AviatrixTransPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTransPeerArrayInput)(nil)).Elem(), AviatrixTransPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTransPeerMapInput)(nil)).Elem(), AviatrixTransPeerMap{})
	pulumi.RegisterOutputType(AviatrixTransPeerOutput{})
	pulumi.RegisterOutputType(AviatrixTransPeerArrayOutput{})
	pulumi.RegisterOutputType(AviatrixTransPeerMapOutput{})
}
