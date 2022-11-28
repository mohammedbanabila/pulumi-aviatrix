// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_controller_email_exception_notification_config** resource allows management of an Aviatrix Controller's email exception notification config. This resource is available as of provider version R2.19+.
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
//			_, err := aviatrix.NewAviatrixControllerEmailExceptionNotificationConfig(ctx, "test", &aviatrix.AviatrixControllerEmailExceptionNotificationConfigArgs{
//				EnableEmailExceptionNotification: pulumi.Bool(false),
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
// **aviatrix_controller_email_exception_notification_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixControllerEmailExceptionNotificationConfig:AviatrixControllerEmailExceptionNotificationConfig test 10-11-12-13
//
// ```
type AviatrixControllerEmailExceptionNotificationConfig struct {
	pulumi.CustomResourceState

	// Enable exception email notification. When set to true, exception email will be sent to "exception@aviatrix.com", when set to false, exception email will be sent to controller's admin email. Valid values: true, false. Default value: true.
	EnableEmailExceptionNotification pulumi.BoolPtrOutput `pulumi:"enableEmailExceptionNotification"`
}

// NewAviatrixControllerEmailExceptionNotificationConfig registers a new resource with the given unique name, arguments, and options.
func NewAviatrixControllerEmailExceptionNotificationConfig(ctx *pulumi.Context,
	name string, args *AviatrixControllerEmailExceptionNotificationConfigArgs, opts ...pulumi.ResourceOption) (*AviatrixControllerEmailExceptionNotificationConfig, error) {
	if args == nil {
		args = &AviatrixControllerEmailExceptionNotificationConfigArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixControllerEmailExceptionNotificationConfig
	err := ctx.RegisterResource("aviatrix:index/aviatrixControllerEmailExceptionNotificationConfig:AviatrixControllerEmailExceptionNotificationConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixControllerEmailExceptionNotificationConfig gets an existing AviatrixControllerEmailExceptionNotificationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixControllerEmailExceptionNotificationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixControllerEmailExceptionNotificationConfigState, opts ...pulumi.ResourceOption) (*AviatrixControllerEmailExceptionNotificationConfig, error) {
	var resource AviatrixControllerEmailExceptionNotificationConfig
	err := ctx.ReadResource("aviatrix:index/aviatrixControllerEmailExceptionNotificationConfig:AviatrixControllerEmailExceptionNotificationConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixControllerEmailExceptionNotificationConfig resources.
type aviatrixControllerEmailExceptionNotificationConfigState struct {
	// Enable exception email notification. When set to true, exception email will be sent to "exception@aviatrix.com", when set to false, exception email will be sent to controller's admin email. Valid values: true, false. Default value: true.
	EnableEmailExceptionNotification *bool `pulumi:"enableEmailExceptionNotification"`
}

type AviatrixControllerEmailExceptionNotificationConfigState struct {
	// Enable exception email notification. When set to true, exception email will be sent to "exception@aviatrix.com", when set to false, exception email will be sent to controller's admin email. Valid values: true, false. Default value: true.
	EnableEmailExceptionNotification pulumi.BoolPtrInput
}

func (AviatrixControllerEmailExceptionNotificationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerEmailExceptionNotificationConfigState)(nil)).Elem()
}

type aviatrixControllerEmailExceptionNotificationConfigArgs struct {
	// Enable exception email notification. When set to true, exception email will be sent to "exception@aviatrix.com", when set to false, exception email will be sent to controller's admin email. Valid values: true, false. Default value: true.
	EnableEmailExceptionNotification *bool `pulumi:"enableEmailExceptionNotification"`
}

// The set of arguments for constructing a AviatrixControllerEmailExceptionNotificationConfig resource.
type AviatrixControllerEmailExceptionNotificationConfigArgs struct {
	// Enable exception email notification. When set to true, exception email will be sent to "exception@aviatrix.com", when set to false, exception email will be sent to controller's admin email. Valid values: true, false. Default value: true.
	EnableEmailExceptionNotification pulumi.BoolPtrInput
}

func (AviatrixControllerEmailExceptionNotificationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerEmailExceptionNotificationConfigArgs)(nil)).Elem()
}

type AviatrixControllerEmailExceptionNotificationConfigInput interface {
	pulumi.Input

	ToAviatrixControllerEmailExceptionNotificationConfigOutput() AviatrixControllerEmailExceptionNotificationConfigOutput
	ToAviatrixControllerEmailExceptionNotificationConfigOutputWithContext(ctx context.Context) AviatrixControllerEmailExceptionNotificationConfigOutput
}

func (*AviatrixControllerEmailExceptionNotificationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerEmailExceptionNotificationConfig)(nil)).Elem()
}

func (i *AviatrixControllerEmailExceptionNotificationConfig) ToAviatrixControllerEmailExceptionNotificationConfigOutput() AviatrixControllerEmailExceptionNotificationConfigOutput {
	return i.ToAviatrixControllerEmailExceptionNotificationConfigOutputWithContext(context.Background())
}

func (i *AviatrixControllerEmailExceptionNotificationConfig) ToAviatrixControllerEmailExceptionNotificationConfigOutputWithContext(ctx context.Context) AviatrixControllerEmailExceptionNotificationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerEmailExceptionNotificationConfigOutput)
}

// AviatrixControllerEmailExceptionNotificationConfigArrayInput is an input type that accepts AviatrixControllerEmailExceptionNotificationConfigArray and AviatrixControllerEmailExceptionNotificationConfigArrayOutput values.
// You can construct a concrete instance of `AviatrixControllerEmailExceptionNotificationConfigArrayInput` via:
//
//	AviatrixControllerEmailExceptionNotificationConfigArray{ AviatrixControllerEmailExceptionNotificationConfigArgs{...} }
type AviatrixControllerEmailExceptionNotificationConfigArrayInput interface {
	pulumi.Input

	ToAviatrixControllerEmailExceptionNotificationConfigArrayOutput() AviatrixControllerEmailExceptionNotificationConfigArrayOutput
	ToAviatrixControllerEmailExceptionNotificationConfigArrayOutputWithContext(context.Context) AviatrixControllerEmailExceptionNotificationConfigArrayOutput
}

type AviatrixControllerEmailExceptionNotificationConfigArray []AviatrixControllerEmailExceptionNotificationConfigInput

func (AviatrixControllerEmailExceptionNotificationConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerEmailExceptionNotificationConfig)(nil)).Elem()
}

func (i AviatrixControllerEmailExceptionNotificationConfigArray) ToAviatrixControllerEmailExceptionNotificationConfigArrayOutput() AviatrixControllerEmailExceptionNotificationConfigArrayOutput {
	return i.ToAviatrixControllerEmailExceptionNotificationConfigArrayOutputWithContext(context.Background())
}

func (i AviatrixControllerEmailExceptionNotificationConfigArray) ToAviatrixControllerEmailExceptionNotificationConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerEmailExceptionNotificationConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerEmailExceptionNotificationConfigArrayOutput)
}

// AviatrixControllerEmailExceptionNotificationConfigMapInput is an input type that accepts AviatrixControllerEmailExceptionNotificationConfigMap and AviatrixControllerEmailExceptionNotificationConfigMapOutput values.
// You can construct a concrete instance of `AviatrixControllerEmailExceptionNotificationConfigMapInput` via:
//
//	AviatrixControllerEmailExceptionNotificationConfigMap{ "key": AviatrixControllerEmailExceptionNotificationConfigArgs{...} }
type AviatrixControllerEmailExceptionNotificationConfigMapInput interface {
	pulumi.Input

	ToAviatrixControllerEmailExceptionNotificationConfigMapOutput() AviatrixControllerEmailExceptionNotificationConfigMapOutput
	ToAviatrixControllerEmailExceptionNotificationConfigMapOutputWithContext(context.Context) AviatrixControllerEmailExceptionNotificationConfigMapOutput
}

type AviatrixControllerEmailExceptionNotificationConfigMap map[string]AviatrixControllerEmailExceptionNotificationConfigInput

func (AviatrixControllerEmailExceptionNotificationConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerEmailExceptionNotificationConfig)(nil)).Elem()
}

func (i AviatrixControllerEmailExceptionNotificationConfigMap) ToAviatrixControllerEmailExceptionNotificationConfigMapOutput() AviatrixControllerEmailExceptionNotificationConfigMapOutput {
	return i.ToAviatrixControllerEmailExceptionNotificationConfigMapOutputWithContext(context.Background())
}

func (i AviatrixControllerEmailExceptionNotificationConfigMap) ToAviatrixControllerEmailExceptionNotificationConfigMapOutputWithContext(ctx context.Context) AviatrixControllerEmailExceptionNotificationConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerEmailExceptionNotificationConfigMapOutput)
}

type AviatrixControllerEmailExceptionNotificationConfigOutput struct{ *pulumi.OutputState }

func (AviatrixControllerEmailExceptionNotificationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerEmailExceptionNotificationConfig)(nil)).Elem()
}

func (o AviatrixControllerEmailExceptionNotificationConfigOutput) ToAviatrixControllerEmailExceptionNotificationConfigOutput() AviatrixControllerEmailExceptionNotificationConfigOutput {
	return o
}

func (o AviatrixControllerEmailExceptionNotificationConfigOutput) ToAviatrixControllerEmailExceptionNotificationConfigOutputWithContext(ctx context.Context) AviatrixControllerEmailExceptionNotificationConfigOutput {
	return o
}

// Enable exception email notification. When set to true, exception email will be sent to "exception@aviatrix.com", when set to false, exception email will be sent to controller's admin email. Valid values: true, false. Default value: true.
func (o AviatrixControllerEmailExceptionNotificationConfigOutput) EnableEmailExceptionNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixControllerEmailExceptionNotificationConfig) pulumi.BoolPtrOutput {
		return v.EnableEmailExceptionNotification
	}).(pulumi.BoolPtrOutput)
}

type AviatrixControllerEmailExceptionNotificationConfigArrayOutput struct{ *pulumi.OutputState }

func (AviatrixControllerEmailExceptionNotificationConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerEmailExceptionNotificationConfig)(nil)).Elem()
}

func (o AviatrixControllerEmailExceptionNotificationConfigArrayOutput) ToAviatrixControllerEmailExceptionNotificationConfigArrayOutput() AviatrixControllerEmailExceptionNotificationConfigArrayOutput {
	return o
}

func (o AviatrixControllerEmailExceptionNotificationConfigArrayOutput) ToAviatrixControllerEmailExceptionNotificationConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerEmailExceptionNotificationConfigArrayOutput {
	return o
}

func (o AviatrixControllerEmailExceptionNotificationConfigArrayOutput) Index(i pulumi.IntInput) AviatrixControllerEmailExceptionNotificationConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixControllerEmailExceptionNotificationConfig {
		return vs[0].([]*AviatrixControllerEmailExceptionNotificationConfig)[vs[1].(int)]
	}).(AviatrixControllerEmailExceptionNotificationConfigOutput)
}

type AviatrixControllerEmailExceptionNotificationConfigMapOutput struct{ *pulumi.OutputState }

func (AviatrixControllerEmailExceptionNotificationConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerEmailExceptionNotificationConfig)(nil)).Elem()
}

func (o AviatrixControllerEmailExceptionNotificationConfigMapOutput) ToAviatrixControllerEmailExceptionNotificationConfigMapOutput() AviatrixControllerEmailExceptionNotificationConfigMapOutput {
	return o
}

func (o AviatrixControllerEmailExceptionNotificationConfigMapOutput) ToAviatrixControllerEmailExceptionNotificationConfigMapOutputWithContext(ctx context.Context) AviatrixControllerEmailExceptionNotificationConfigMapOutput {
	return o
}

func (o AviatrixControllerEmailExceptionNotificationConfigMapOutput) MapIndex(k pulumi.StringInput) AviatrixControllerEmailExceptionNotificationConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixControllerEmailExceptionNotificationConfig {
		return vs[0].(map[string]*AviatrixControllerEmailExceptionNotificationConfig)[vs[1].(string)]
	}).(AviatrixControllerEmailExceptionNotificationConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerEmailExceptionNotificationConfigInput)(nil)).Elem(), &AviatrixControllerEmailExceptionNotificationConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerEmailExceptionNotificationConfigArrayInput)(nil)).Elem(), AviatrixControllerEmailExceptionNotificationConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerEmailExceptionNotificationConfigMapInput)(nil)).Elem(), AviatrixControllerEmailExceptionNotificationConfigMap{})
	pulumi.RegisterOutputType(AviatrixControllerEmailExceptionNotificationConfigOutput{})
	pulumi.RegisterOutputType(AviatrixControllerEmailExceptionNotificationConfigArrayOutput{})
	pulumi.RegisterOutputType(AviatrixControllerEmailExceptionNotificationConfigMapOutput{})
}
