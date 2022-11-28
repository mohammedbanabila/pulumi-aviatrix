// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_vpn_cert_download** resource manages the VPN Certificate Download configuration for SAML Authentication
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
//			_, err := aviatrix.NewAviatrixVpnCertDownload(ctx, "testVpnCertDownload", &aviatrix.AviatrixVpnCertDownloadArgs{
//				DownloadEnabled: pulumi.Bool(true),
//				SamlEndpoints: pulumi.StringArray{
//					pulumi.String("saml_endpoint_name"),
//				},
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
// **vpn_cert_download** can be imported using the default id `vpn_cert_download`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixVpnCertDownload:AviatrixVpnCertDownload test_vpn_cert_download vpn_cert_download
//
// ```
type AviatrixVpnCertDownload struct {
	pulumi.CustomResourceState

	// Whether the VPN Certificate download is enabled. Supported Values: "true", "false".
	DownloadEnabled pulumi.BoolPtrOutput `pulumi:"downloadEnabled"`
	// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is supported. Example: ["samlEndpoint1"].
	SamlEndpoints pulumi.StringArrayOutput `pulumi:"samlEndpoints"`
}

// NewAviatrixVpnCertDownload registers a new resource with the given unique name, arguments, and options.
func NewAviatrixVpnCertDownload(ctx *pulumi.Context,
	name string, args *AviatrixVpnCertDownloadArgs, opts ...pulumi.ResourceOption) (*AviatrixVpnCertDownload, error) {
	if args == nil {
		args = &AviatrixVpnCertDownloadArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixVpnCertDownload
	err := ctx.RegisterResource("aviatrix:index/aviatrixVpnCertDownload:AviatrixVpnCertDownload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixVpnCertDownload gets an existing AviatrixVpnCertDownload resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixVpnCertDownload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixVpnCertDownloadState, opts ...pulumi.ResourceOption) (*AviatrixVpnCertDownload, error) {
	var resource AviatrixVpnCertDownload
	err := ctx.ReadResource("aviatrix:index/aviatrixVpnCertDownload:AviatrixVpnCertDownload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixVpnCertDownload resources.
type aviatrixVpnCertDownloadState struct {
	// Whether the VPN Certificate download is enabled. Supported Values: "true", "false".
	DownloadEnabled *bool `pulumi:"downloadEnabled"`
	// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is supported. Example: ["samlEndpoint1"].
	SamlEndpoints []string `pulumi:"samlEndpoints"`
}

type AviatrixVpnCertDownloadState struct {
	// Whether the VPN Certificate download is enabled. Supported Values: "true", "false".
	DownloadEnabled pulumi.BoolPtrInput
	// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is supported. Example: ["samlEndpoint1"].
	SamlEndpoints pulumi.StringArrayInput
}

func (AviatrixVpnCertDownloadState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixVpnCertDownloadState)(nil)).Elem()
}

type aviatrixVpnCertDownloadArgs struct {
	// Whether the VPN Certificate download is enabled. Supported Values: "true", "false".
	DownloadEnabled *bool `pulumi:"downloadEnabled"`
	// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is supported. Example: ["samlEndpoint1"].
	SamlEndpoints []string `pulumi:"samlEndpoints"`
}

// The set of arguments for constructing a AviatrixVpnCertDownload resource.
type AviatrixVpnCertDownloadArgs struct {
	// Whether the VPN Certificate download is enabled. Supported Values: "true", "false".
	DownloadEnabled pulumi.BoolPtrInput
	// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is supported. Example: ["samlEndpoint1"].
	SamlEndpoints pulumi.StringArrayInput
}

func (AviatrixVpnCertDownloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixVpnCertDownloadArgs)(nil)).Elem()
}

type AviatrixVpnCertDownloadInput interface {
	pulumi.Input

	ToAviatrixVpnCertDownloadOutput() AviatrixVpnCertDownloadOutput
	ToAviatrixVpnCertDownloadOutputWithContext(ctx context.Context) AviatrixVpnCertDownloadOutput
}

func (*AviatrixVpnCertDownload) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixVpnCertDownload)(nil)).Elem()
}

func (i *AviatrixVpnCertDownload) ToAviatrixVpnCertDownloadOutput() AviatrixVpnCertDownloadOutput {
	return i.ToAviatrixVpnCertDownloadOutputWithContext(context.Background())
}

func (i *AviatrixVpnCertDownload) ToAviatrixVpnCertDownloadOutputWithContext(ctx context.Context) AviatrixVpnCertDownloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpnCertDownloadOutput)
}

// AviatrixVpnCertDownloadArrayInput is an input type that accepts AviatrixVpnCertDownloadArray and AviatrixVpnCertDownloadArrayOutput values.
// You can construct a concrete instance of `AviatrixVpnCertDownloadArrayInput` via:
//
//	AviatrixVpnCertDownloadArray{ AviatrixVpnCertDownloadArgs{...} }
type AviatrixVpnCertDownloadArrayInput interface {
	pulumi.Input

	ToAviatrixVpnCertDownloadArrayOutput() AviatrixVpnCertDownloadArrayOutput
	ToAviatrixVpnCertDownloadArrayOutputWithContext(context.Context) AviatrixVpnCertDownloadArrayOutput
}

type AviatrixVpnCertDownloadArray []AviatrixVpnCertDownloadInput

func (AviatrixVpnCertDownloadArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixVpnCertDownload)(nil)).Elem()
}

func (i AviatrixVpnCertDownloadArray) ToAviatrixVpnCertDownloadArrayOutput() AviatrixVpnCertDownloadArrayOutput {
	return i.ToAviatrixVpnCertDownloadArrayOutputWithContext(context.Background())
}

func (i AviatrixVpnCertDownloadArray) ToAviatrixVpnCertDownloadArrayOutputWithContext(ctx context.Context) AviatrixVpnCertDownloadArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpnCertDownloadArrayOutput)
}

// AviatrixVpnCertDownloadMapInput is an input type that accepts AviatrixVpnCertDownloadMap and AviatrixVpnCertDownloadMapOutput values.
// You can construct a concrete instance of `AviatrixVpnCertDownloadMapInput` via:
//
//	AviatrixVpnCertDownloadMap{ "key": AviatrixVpnCertDownloadArgs{...} }
type AviatrixVpnCertDownloadMapInput interface {
	pulumi.Input

	ToAviatrixVpnCertDownloadMapOutput() AviatrixVpnCertDownloadMapOutput
	ToAviatrixVpnCertDownloadMapOutputWithContext(context.Context) AviatrixVpnCertDownloadMapOutput
}

type AviatrixVpnCertDownloadMap map[string]AviatrixVpnCertDownloadInput

func (AviatrixVpnCertDownloadMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixVpnCertDownload)(nil)).Elem()
}

func (i AviatrixVpnCertDownloadMap) ToAviatrixVpnCertDownloadMapOutput() AviatrixVpnCertDownloadMapOutput {
	return i.ToAviatrixVpnCertDownloadMapOutputWithContext(context.Background())
}

func (i AviatrixVpnCertDownloadMap) ToAviatrixVpnCertDownloadMapOutputWithContext(ctx context.Context) AviatrixVpnCertDownloadMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpnCertDownloadMapOutput)
}

type AviatrixVpnCertDownloadOutput struct{ *pulumi.OutputState }

func (AviatrixVpnCertDownloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixVpnCertDownload)(nil)).Elem()
}

func (o AviatrixVpnCertDownloadOutput) ToAviatrixVpnCertDownloadOutput() AviatrixVpnCertDownloadOutput {
	return o
}

func (o AviatrixVpnCertDownloadOutput) ToAviatrixVpnCertDownloadOutputWithContext(ctx context.Context) AviatrixVpnCertDownloadOutput {
	return o
}

// Whether the VPN Certificate download is enabled. Supported Values: "true", "false".
func (o AviatrixVpnCertDownloadOutput) DownloadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixVpnCertDownload) pulumi.BoolPtrOutput { return v.DownloadEnabled }).(pulumi.BoolPtrOutput)
}

// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is supported. Example: ["samlEndpoint1"].
func (o AviatrixVpnCertDownloadOutput) SamlEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixVpnCertDownload) pulumi.StringArrayOutput { return v.SamlEndpoints }).(pulumi.StringArrayOutput)
}

type AviatrixVpnCertDownloadArrayOutput struct{ *pulumi.OutputState }

func (AviatrixVpnCertDownloadArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixVpnCertDownload)(nil)).Elem()
}

func (o AviatrixVpnCertDownloadArrayOutput) ToAviatrixVpnCertDownloadArrayOutput() AviatrixVpnCertDownloadArrayOutput {
	return o
}

func (o AviatrixVpnCertDownloadArrayOutput) ToAviatrixVpnCertDownloadArrayOutputWithContext(ctx context.Context) AviatrixVpnCertDownloadArrayOutput {
	return o
}

func (o AviatrixVpnCertDownloadArrayOutput) Index(i pulumi.IntInput) AviatrixVpnCertDownloadOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixVpnCertDownload {
		return vs[0].([]*AviatrixVpnCertDownload)[vs[1].(int)]
	}).(AviatrixVpnCertDownloadOutput)
}

type AviatrixVpnCertDownloadMapOutput struct{ *pulumi.OutputState }

func (AviatrixVpnCertDownloadMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixVpnCertDownload)(nil)).Elem()
}

func (o AviatrixVpnCertDownloadMapOutput) ToAviatrixVpnCertDownloadMapOutput() AviatrixVpnCertDownloadMapOutput {
	return o
}

func (o AviatrixVpnCertDownloadMapOutput) ToAviatrixVpnCertDownloadMapOutputWithContext(ctx context.Context) AviatrixVpnCertDownloadMapOutput {
	return o
}

func (o AviatrixVpnCertDownloadMapOutput) MapIndex(k pulumi.StringInput) AviatrixVpnCertDownloadOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixVpnCertDownload {
		return vs[0].(map[string]*AviatrixVpnCertDownload)[vs[1].(string)]
	}).(AviatrixVpnCertDownloadOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpnCertDownloadInput)(nil)).Elem(), &AviatrixVpnCertDownload{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpnCertDownloadArrayInput)(nil)).Elem(), AviatrixVpnCertDownloadArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpnCertDownloadMapInput)(nil)).Elem(), AviatrixVpnCertDownloadMap{})
	pulumi.RegisterOutputType(AviatrixVpnCertDownloadOutput{})
	pulumi.RegisterOutputType(AviatrixVpnCertDownloadArrayOutput{})
	pulumi.RegisterOutputType(AviatrixVpnCertDownloadMapOutput{})
}
