// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_controller_gateway_keepalive_config** resource allows management of an Aviatrix Controller's gateway keepalive template configuration. This resource is available as of provider version R2.19.2+.
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
//			_, err := aviatrix.NewAviatrixControllerGatewayKeepaliveConfig(ctx, "testGatewayKeepalive", &aviatrix.AviatrixControllerGatewayKeepaliveConfigArgs{
//				KeepaliveSpeed: pulumi.String("medium"),
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
// **aviatrix_controller_gateway_keepalive_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixControllerGatewayKeepaliveConfig:AviatrixControllerGatewayKeepaliveConfig test_gateway_keepalive 10-11-12-13
//
// ```
type AviatrixControllerGatewayKeepaliveConfig struct {
	pulumi.CustomResourceState

	// The gateway keepalive template name. Must be one of "slow", "medium" or "fast". Visit [here](https://docs.aviatrix.com/HowTos/gateway.html#gateway-keepalives) for the complete documentation about the gateway keepalive configuration.
	KeepaliveSpeed pulumi.StringOutput `pulumi:"keepaliveSpeed"`
}

// NewAviatrixControllerGatewayKeepaliveConfig registers a new resource with the given unique name, arguments, and options.
func NewAviatrixControllerGatewayKeepaliveConfig(ctx *pulumi.Context,
	name string, args *AviatrixControllerGatewayKeepaliveConfigArgs, opts ...pulumi.ResourceOption) (*AviatrixControllerGatewayKeepaliveConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeepaliveSpeed == nil {
		return nil, errors.New("invalid value for required argument 'KeepaliveSpeed'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixControllerGatewayKeepaliveConfig
	err := ctx.RegisterResource("aviatrix:index/aviatrixControllerGatewayKeepaliveConfig:AviatrixControllerGatewayKeepaliveConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixControllerGatewayKeepaliveConfig gets an existing AviatrixControllerGatewayKeepaliveConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixControllerGatewayKeepaliveConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixControllerGatewayKeepaliveConfigState, opts ...pulumi.ResourceOption) (*AviatrixControllerGatewayKeepaliveConfig, error) {
	var resource AviatrixControllerGatewayKeepaliveConfig
	err := ctx.ReadResource("aviatrix:index/aviatrixControllerGatewayKeepaliveConfig:AviatrixControllerGatewayKeepaliveConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixControllerGatewayKeepaliveConfig resources.
type aviatrixControllerGatewayKeepaliveConfigState struct {
	// The gateway keepalive template name. Must be one of "slow", "medium" or "fast". Visit [here](https://docs.aviatrix.com/HowTos/gateway.html#gateway-keepalives) for the complete documentation about the gateway keepalive configuration.
	KeepaliveSpeed *string `pulumi:"keepaliveSpeed"`
}

type AviatrixControllerGatewayKeepaliveConfigState struct {
	// The gateway keepalive template name. Must be one of "slow", "medium" or "fast". Visit [here](https://docs.aviatrix.com/HowTos/gateway.html#gateway-keepalives) for the complete documentation about the gateway keepalive configuration.
	KeepaliveSpeed pulumi.StringPtrInput
}

func (AviatrixControllerGatewayKeepaliveConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerGatewayKeepaliveConfigState)(nil)).Elem()
}

type aviatrixControllerGatewayKeepaliveConfigArgs struct {
	// The gateway keepalive template name. Must be one of "slow", "medium" or "fast". Visit [here](https://docs.aviatrix.com/HowTos/gateway.html#gateway-keepalives) for the complete documentation about the gateway keepalive configuration.
	KeepaliveSpeed string `pulumi:"keepaliveSpeed"`
}

// The set of arguments for constructing a AviatrixControllerGatewayKeepaliveConfig resource.
type AviatrixControllerGatewayKeepaliveConfigArgs struct {
	// The gateway keepalive template name. Must be one of "slow", "medium" or "fast". Visit [here](https://docs.aviatrix.com/HowTos/gateway.html#gateway-keepalives) for the complete documentation about the gateway keepalive configuration.
	KeepaliveSpeed pulumi.StringInput
}

func (AviatrixControllerGatewayKeepaliveConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerGatewayKeepaliveConfigArgs)(nil)).Elem()
}

type AviatrixControllerGatewayKeepaliveConfigInput interface {
	pulumi.Input

	ToAviatrixControllerGatewayKeepaliveConfigOutput() AviatrixControllerGatewayKeepaliveConfigOutput
	ToAviatrixControllerGatewayKeepaliveConfigOutputWithContext(ctx context.Context) AviatrixControllerGatewayKeepaliveConfigOutput
}

func (*AviatrixControllerGatewayKeepaliveConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerGatewayKeepaliveConfig)(nil)).Elem()
}

func (i *AviatrixControllerGatewayKeepaliveConfig) ToAviatrixControllerGatewayKeepaliveConfigOutput() AviatrixControllerGatewayKeepaliveConfigOutput {
	return i.ToAviatrixControllerGatewayKeepaliveConfigOutputWithContext(context.Background())
}

func (i *AviatrixControllerGatewayKeepaliveConfig) ToAviatrixControllerGatewayKeepaliveConfigOutputWithContext(ctx context.Context) AviatrixControllerGatewayKeepaliveConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerGatewayKeepaliveConfigOutput)
}

// AviatrixControllerGatewayKeepaliveConfigArrayInput is an input type that accepts AviatrixControllerGatewayKeepaliveConfigArray and AviatrixControllerGatewayKeepaliveConfigArrayOutput values.
// You can construct a concrete instance of `AviatrixControllerGatewayKeepaliveConfigArrayInput` via:
//
//	AviatrixControllerGatewayKeepaliveConfigArray{ AviatrixControllerGatewayKeepaliveConfigArgs{...} }
type AviatrixControllerGatewayKeepaliveConfigArrayInput interface {
	pulumi.Input

	ToAviatrixControllerGatewayKeepaliveConfigArrayOutput() AviatrixControllerGatewayKeepaliveConfigArrayOutput
	ToAviatrixControllerGatewayKeepaliveConfigArrayOutputWithContext(context.Context) AviatrixControllerGatewayKeepaliveConfigArrayOutput
}

type AviatrixControllerGatewayKeepaliveConfigArray []AviatrixControllerGatewayKeepaliveConfigInput

func (AviatrixControllerGatewayKeepaliveConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerGatewayKeepaliveConfig)(nil)).Elem()
}

func (i AviatrixControllerGatewayKeepaliveConfigArray) ToAviatrixControllerGatewayKeepaliveConfigArrayOutput() AviatrixControllerGatewayKeepaliveConfigArrayOutput {
	return i.ToAviatrixControllerGatewayKeepaliveConfigArrayOutputWithContext(context.Background())
}

func (i AviatrixControllerGatewayKeepaliveConfigArray) ToAviatrixControllerGatewayKeepaliveConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerGatewayKeepaliveConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerGatewayKeepaliveConfigArrayOutput)
}

// AviatrixControllerGatewayKeepaliveConfigMapInput is an input type that accepts AviatrixControllerGatewayKeepaliveConfigMap and AviatrixControllerGatewayKeepaliveConfigMapOutput values.
// You can construct a concrete instance of `AviatrixControllerGatewayKeepaliveConfigMapInput` via:
//
//	AviatrixControllerGatewayKeepaliveConfigMap{ "key": AviatrixControllerGatewayKeepaliveConfigArgs{...} }
type AviatrixControllerGatewayKeepaliveConfigMapInput interface {
	pulumi.Input

	ToAviatrixControllerGatewayKeepaliveConfigMapOutput() AviatrixControllerGatewayKeepaliveConfigMapOutput
	ToAviatrixControllerGatewayKeepaliveConfigMapOutputWithContext(context.Context) AviatrixControllerGatewayKeepaliveConfigMapOutput
}

type AviatrixControllerGatewayKeepaliveConfigMap map[string]AviatrixControllerGatewayKeepaliveConfigInput

func (AviatrixControllerGatewayKeepaliveConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerGatewayKeepaliveConfig)(nil)).Elem()
}

func (i AviatrixControllerGatewayKeepaliveConfigMap) ToAviatrixControllerGatewayKeepaliveConfigMapOutput() AviatrixControllerGatewayKeepaliveConfigMapOutput {
	return i.ToAviatrixControllerGatewayKeepaliveConfigMapOutputWithContext(context.Background())
}

func (i AviatrixControllerGatewayKeepaliveConfigMap) ToAviatrixControllerGatewayKeepaliveConfigMapOutputWithContext(ctx context.Context) AviatrixControllerGatewayKeepaliveConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerGatewayKeepaliveConfigMapOutput)
}

type AviatrixControllerGatewayKeepaliveConfigOutput struct{ *pulumi.OutputState }

func (AviatrixControllerGatewayKeepaliveConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerGatewayKeepaliveConfig)(nil)).Elem()
}

func (o AviatrixControllerGatewayKeepaliveConfigOutput) ToAviatrixControllerGatewayKeepaliveConfigOutput() AviatrixControllerGatewayKeepaliveConfigOutput {
	return o
}

func (o AviatrixControllerGatewayKeepaliveConfigOutput) ToAviatrixControllerGatewayKeepaliveConfigOutputWithContext(ctx context.Context) AviatrixControllerGatewayKeepaliveConfigOutput {
	return o
}

// The gateway keepalive template name. Must be one of "slow", "medium" or "fast". Visit [here](https://docs.aviatrix.com/HowTos/gateway.html#gateway-keepalives) for the complete documentation about the gateway keepalive configuration.
func (o AviatrixControllerGatewayKeepaliveConfigOutput) KeepaliveSpeed() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixControllerGatewayKeepaliveConfig) pulumi.StringOutput { return v.KeepaliveSpeed }).(pulumi.StringOutput)
}

type AviatrixControllerGatewayKeepaliveConfigArrayOutput struct{ *pulumi.OutputState }

func (AviatrixControllerGatewayKeepaliveConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerGatewayKeepaliveConfig)(nil)).Elem()
}

func (o AviatrixControllerGatewayKeepaliveConfigArrayOutput) ToAviatrixControllerGatewayKeepaliveConfigArrayOutput() AviatrixControllerGatewayKeepaliveConfigArrayOutput {
	return o
}

func (o AviatrixControllerGatewayKeepaliveConfigArrayOutput) ToAviatrixControllerGatewayKeepaliveConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerGatewayKeepaliveConfigArrayOutput {
	return o
}

func (o AviatrixControllerGatewayKeepaliveConfigArrayOutput) Index(i pulumi.IntInput) AviatrixControllerGatewayKeepaliveConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixControllerGatewayKeepaliveConfig {
		return vs[0].([]*AviatrixControllerGatewayKeepaliveConfig)[vs[1].(int)]
	}).(AviatrixControllerGatewayKeepaliveConfigOutput)
}

type AviatrixControllerGatewayKeepaliveConfigMapOutput struct{ *pulumi.OutputState }

func (AviatrixControllerGatewayKeepaliveConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerGatewayKeepaliveConfig)(nil)).Elem()
}

func (o AviatrixControllerGatewayKeepaliveConfigMapOutput) ToAviatrixControllerGatewayKeepaliveConfigMapOutput() AviatrixControllerGatewayKeepaliveConfigMapOutput {
	return o
}

func (o AviatrixControllerGatewayKeepaliveConfigMapOutput) ToAviatrixControllerGatewayKeepaliveConfigMapOutputWithContext(ctx context.Context) AviatrixControllerGatewayKeepaliveConfigMapOutput {
	return o
}

func (o AviatrixControllerGatewayKeepaliveConfigMapOutput) MapIndex(k pulumi.StringInput) AviatrixControllerGatewayKeepaliveConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixControllerGatewayKeepaliveConfig {
		return vs[0].(map[string]*AviatrixControllerGatewayKeepaliveConfig)[vs[1].(string)]
	}).(AviatrixControllerGatewayKeepaliveConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerGatewayKeepaliveConfigInput)(nil)).Elem(), &AviatrixControllerGatewayKeepaliveConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerGatewayKeepaliveConfigArrayInput)(nil)).Elem(), AviatrixControllerGatewayKeepaliveConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerGatewayKeepaliveConfigMapInput)(nil)).Elem(), AviatrixControllerGatewayKeepaliveConfigMap{})
	pulumi.RegisterOutputType(AviatrixControllerGatewayKeepaliveConfigOutput{})
	pulumi.RegisterOutputType(AviatrixControllerGatewayKeepaliveConfigArrayOutput{})
	pulumi.RegisterOutputType(AviatrixControllerGatewayKeepaliveConfigMapOutput{})
}
