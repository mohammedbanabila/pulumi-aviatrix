// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_aws_tgw_connect_peer** resource allows the creation and management of AWS TGW Connect peers. This
// resource is available as of provider version R2.18.1+.
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
//			_, err := aviatrix.NewAviatrixAwsTgwConnectPeer(ctx, "test", &aviatrix.AviatrixAwsTgwConnectPeerArgs{
//				TgwName:             pulumi.Any(aviatrix_aws_tgw.Test_aws_tgw.Tgw_name),
//				ConnectionName:      pulumi.Any(aviatrix_aws_tgw_connect.Test_aws_tgw_connect.Connection_name),
//				ConnectPeerName:     pulumi.String("connect-peer-test"),
//				ConnectAttachmentId: pulumi.Any(aviatrix_aws_tgw_connect.Test_aws_tgw_connect.Connect_attachment_id),
//				PeerAsNumber:        pulumi.String("65001"),
//				PeerGreAddress:      pulumi.String("172.31.1.11"),
//				BgpInsideCidrs: pulumi.StringArray{
//					pulumi.String("169.254.6.0/29"),
//				},
//				TgwGreAddress: pulumi.String("10.0.0.32"),
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
// **aws_tgw_connect_peer** can be imported using the `tgw_name`, `connection_name` and `connect_peer_name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixAwsTgwConnectPeer:AviatrixAwsTgwConnectPeer test tgw_name~~connection_name~~connect_peer_name
//
// ```
type AviatrixAwsTgwConnectPeer struct {
	pulumi.CustomResourceState

	// Set of BGP Inside CIDR Block(s).
	BgpInsideCidrs pulumi.StringArrayOutput `pulumi:"bgpInsideCidrs"`
	// Connect Attachment ID.
	ConnectAttachmentId pulumi.StringOutput `pulumi:"connectAttachmentId"`
	// Connect Peer ID.
	ConnectPeerId pulumi.StringOutput `pulumi:"connectPeerId"`
	// TGW Connect peer name.
	ConnectPeerName pulumi.StringOutput `pulumi:"connectPeerName"`
	// TGW Connect connection name.
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// Peer AS Number.
	PeerAsNumber pulumi.StringOutput `pulumi:"peerAsNumber"`
	// Peer GRE IP Address.
	PeerGreAddress pulumi.StringOutput `pulumi:"peerGreAddress"`
	// AWS TGW GRE IP Address.
	TgwGreAddress pulumi.StringPtrOutput `pulumi:"tgwGreAddress"`
	// AWS TGW name.
	TgwName pulumi.StringOutput `pulumi:"tgwName"`
}

// NewAviatrixAwsTgwConnectPeer registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAwsTgwConnectPeer(ctx *pulumi.Context,
	name string, args *AviatrixAwsTgwConnectPeerArgs, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwConnectPeer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BgpInsideCidrs == nil {
		return nil, errors.New("invalid value for required argument 'BgpInsideCidrs'")
	}
	if args.ConnectAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectAttachmentId'")
	}
	if args.ConnectPeerName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectPeerName'")
	}
	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.PeerAsNumber == nil {
		return nil, errors.New("invalid value for required argument 'PeerAsNumber'")
	}
	if args.PeerGreAddress == nil {
		return nil, errors.New("invalid value for required argument 'PeerGreAddress'")
	}
	if args.TgwName == nil {
		return nil, errors.New("invalid value for required argument 'TgwName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAwsTgwConnectPeer
	err := ctx.RegisterResource("aviatrix:index/aviatrixAwsTgwConnectPeer:AviatrixAwsTgwConnectPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAwsTgwConnectPeer gets an existing AviatrixAwsTgwConnectPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAwsTgwConnectPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAwsTgwConnectPeerState, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwConnectPeer, error) {
	var resource AviatrixAwsTgwConnectPeer
	err := ctx.ReadResource("aviatrix:index/aviatrixAwsTgwConnectPeer:AviatrixAwsTgwConnectPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAwsTgwConnectPeer resources.
type aviatrixAwsTgwConnectPeerState struct {
	// Set of BGP Inside CIDR Block(s).
	BgpInsideCidrs []string `pulumi:"bgpInsideCidrs"`
	// Connect Attachment ID.
	ConnectAttachmentId *string `pulumi:"connectAttachmentId"`
	// Connect Peer ID.
	ConnectPeerId *string `pulumi:"connectPeerId"`
	// TGW Connect peer name.
	ConnectPeerName *string `pulumi:"connectPeerName"`
	// TGW Connect connection name.
	ConnectionName *string `pulumi:"connectionName"`
	// Peer AS Number.
	PeerAsNumber *string `pulumi:"peerAsNumber"`
	// Peer GRE IP Address.
	PeerGreAddress *string `pulumi:"peerGreAddress"`
	// AWS TGW GRE IP Address.
	TgwGreAddress *string `pulumi:"tgwGreAddress"`
	// AWS TGW name.
	TgwName *string `pulumi:"tgwName"`
}

type AviatrixAwsTgwConnectPeerState struct {
	// Set of BGP Inside CIDR Block(s).
	BgpInsideCidrs pulumi.StringArrayInput
	// Connect Attachment ID.
	ConnectAttachmentId pulumi.StringPtrInput
	// Connect Peer ID.
	ConnectPeerId pulumi.StringPtrInput
	// TGW Connect peer name.
	ConnectPeerName pulumi.StringPtrInput
	// TGW Connect connection name.
	ConnectionName pulumi.StringPtrInput
	// Peer AS Number.
	PeerAsNumber pulumi.StringPtrInput
	// Peer GRE IP Address.
	PeerGreAddress pulumi.StringPtrInput
	// AWS TGW GRE IP Address.
	TgwGreAddress pulumi.StringPtrInput
	// AWS TGW name.
	TgwName pulumi.StringPtrInput
}

func (AviatrixAwsTgwConnectPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwConnectPeerState)(nil)).Elem()
}

type aviatrixAwsTgwConnectPeerArgs struct {
	// Set of BGP Inside CIDR Block(s).
	BgpInsideCidrs []string `pulumi:"bgpInsideCidrs"`
	// Connect Attachment ID.
	ConnectAttachmentId string `pulumi:"connectAttachmentId"`
	// TGW Connect peer name.
	ConnectPeerName string `pulumi:"connectPeerName"`
	// TGW Connect connection name.
	ConnectionName string `pulumi:"connectionName"`
	// Peer AS Number.
	PeerAsNumber string `pulumi:"peerAsNumber"`
	// Peer GRE IP Address.
	PeerGreAddress string `pulumi:"peerGreAddress"`
	// AWS TGW GRE IP Address.
	TgwGreAddress *string `pulumi:"tgwGreAddress"`
	// AWS TGW name.
	TgwName string `pulumi:"tgwName"`
}

// The set of arguments for constructing a AviatrixAwsTgwConnectPeer resource.
type AviatrixAwsTgwConnectPeerArgs struct {
	// Set of BGP Inside CIDR Block(s).
	BgpInsideCidrs pulumi.StringArrayInput
	// Connect Attachment ID.
	ConnectAttachmentId pulumi.StringInput
	// TGW Connect peer name.
	ConnectPeerName pulumi.StringInput
	// TGW Connect connection name.
	ConnectionName pulumi.StringInput
	// Peer AS Number.
	PeerAsNumber pulumi.StringInput
	// Peer GRE IP Address.
	PeerGreAddress pulumi.StringInput
	// AWS TGW GRE IP Address.
	TgwGreAddress pulumi.StringPtrInput
	// AWS TGW name.
	TgwName pulumi.StringInput
}

func (AviatrixAwsTgwConnectPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwConnectPeerArgs)(nil)).Elem()
}

type AviatrixAwsTgwConnectPeerInput interface {
	pulumi.Input

	ToAviatrixAwsTgwConnectPeerOutput() AviatrixAwsTgwConnectPeerOutput
	ToAviatrixAwsTgwConnectPeerOutputWithContext(ctx context.Context) AviatrixAwsTgwConnectPeerOutput
}

func (*AviatrixAwsTgwConnectPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwConnectPeer)(nil)).Elem()
}

func (i *AviatrixAwsTgwConnectPeer) ToAviatrixAwsTgwConnectPeerOutput() AviatrixAwsTgwConnectPeerOutput {
	return i.ToAviatrixAwsTgwConnectPeerOutputWithContext(context.Background())
}

func (i *AviatrixAwsTgwConnectPeer) ToAviatrixAwsTgwConnectPeerOutputWithContext(ctx context.Context) AviatrixAwsTgwConnectPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwConnectPeerOutput)
}

// AviatrixAwsTgwConnectPeerArrayInput is an input type that accepts AviatrixAwsTgwConnectPeerArray and AviatrixAwsTgwConnectPeerArrayOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwConnectPeerArrayInput` via:
//
//	AviatrixAwsTgwConnectPeerArray{ AviatrixAwsTgwConnectPeerArgs{...} }
type AviatrixAwsTgwConnectPeerArrayInput interface {
	pulumi.Input

	ToAviatrixAwsTgwConnectPeerArrayOutput() AviatrixAwsTgwConnectPeerArrayOutput
	ToAviatrixAwsTgwConnectPeerArrayOutputWithContext(context.Context) AviatrixAwsTgwConnectPeerArrayOutput
}

type AviatrixAwsTgwConnectPeerArray []AviatrixAwsTgwConnectPeerInput

func (AviatrixAwsTgwConnectPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwConnectPeer)(nil)).Elem()
}

func (i AviatrixAwsTgwConnectPeerArray) ToAviatrixAwsTgwConnectPeerArrayOutput() AviatrixAwsTgwConnectPeerArrayOutput {
	return i.ToAviatrixAwsTgwConnectPeerArrayOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwConnectPeerArray) ToAviatrixAwsTgwConnectPeerArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwConnectPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwConnectPeerArrayOutput)
}

// AviatrixAwsTgwConnectPeerMapInput is an input type that accepts AviatrixAwsTgwConnectPeerMap and AviatrixAwsTgwConnectPeerMapOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwConnectPeerMapInput` via:
//
//	AviatrixAwsTgwConnectPeerMap{ "key": AviatrixAwsTgwConnectPeerArgs{...} }
type AviatrixAwsTgwConnectPeerMapInput interface {
	pulumi.Input

	ToAviatrixAwsTgwConnectPeerMapOutput() AviatrixAwsTgwConnectPeerMapOutput
	ToAviatrixAwsTgwConnectPeerMapOutputWithContext(context.Context) AviatrixAwsTgwConnectPeerMapOutput
}

type AviatrixAwsTgwConnectPeerMap map[string]AviatrixAwsTgwConnectPeerInput

func (AviatrixAwsTgwConnectPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwConnectPeer)(nil)).Elem()
}

func (i AviatrixAwsTgwConnectPeerMap) ToAviatrixAwsTgwConnectPeerMapOutput() AviatrixAwsTgwConnectPeerMapOutput {
	return i.ToAviatrixAwsTgwConnectPeerMapOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwConnectPeerMap) ToAviatrixAwsTgwConnectPeerMapOutputWithContext(ctx context.Context) AviatrixAwsTgwConnectPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwConnectPeerMapOutput)
}

type AviatrixAwsTgwConnectPeerOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwConnectPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwConnectPeer)(nil)).Elem()
}

func (o AviatrixAwsTgwConnectPeerOutput) ToAviatrixAwsTgwConnectPeerOutput() AviatrixAwsTgwConnectPeerOutput {
	return o
}

func (o AviatrixAwsTgwConnectPeerOutput) ToAviatrixAwsTgwConnectPeerOutputWithContext(ctx context.Context) AviatrixAwsTgwConnectPeerOutput {
	return o
}

// Set of BGP Inside CIDR Block(s).
func (o AviatrixAwsTgwConnectPeerOutput) BgpInsideCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringArrayOutput { return v.BgpInsideCidrs }).(pulumi.StringArrayOutput)
}

// Connect Attachment ID.
func (o AviatrixAwsTgwConnectPeerOutput) ConnectAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringOutput { return v.ConnectAttachmentId }).(pulumi.StringOutput)
}

// Connect Peer ID.
func (o AviatrixAwsTgwConnectPeerOutput) ConnectPeerId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringOutput { return v.ConnectPeerId }).(pulumi.StringOutput)
}

// TGW Connect peer name.
func (o AviatrixAwsTgwConnectPeerOutput) ConnectPeerName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringOutput { return v.ConnectPeerName }).(pulumi.StringOutput)
}

// TGW Connect connection name.
func (o AviatrixAwsTgwConnectPeerOutput) ConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringOutput { return v.ConnectionName }).(pulumi.StringOutput)
}

// Peer AS Number.
func (o AviatrixAwsTgwConnectPeerOutput) PeerAsNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringOutput { return v.PeerAsNumber }).(pulumi.StringOutput)
}

// Peer GRE IP Address.
func (o AviatrixAwsTgwConnectPeerOutput) PeerGreAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringOutput { return v.PeerGreAddress }).(pulumi.StringOutput)
}

// AWS TGW GRE IP Address.
func (o AviatrixAwsTgwConnectPeerOutput) TgwGreAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringPtrOutput { return v.TgwGreAddress }).(pulumi.StringPtrOutput)
}

// AWS TGW name.
func (o AviatrixAwsTgwConnectPeerOutput) TgwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwConnectPeer) pulumi.StringOutput { return v.TgwName }).(pulumi.StringOutput)
}

type AviatrixAwsTgwConnectPeerArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwConnectPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwConnectPeer)(nil)).Elem()
}

func (o AviatrixAwsTgwConnectPeerArrayOutput) ToAviatrixAwsTgwConnectPeerArrayOutput() AviatrixAwsTgwConnectPeerArrayOutput {
	return o
}

func (o AviatrixAwsTgwConnectPeerArrayOutput) ToAviatrixAwsTgwConnectPeerArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwConnectPeerArrayOutput {
	return o
}

func (o AviatrixAwsTgwConnectPeerArrayOutput) Index(i pulumi.IntInput) AviatrixAwsTgwConnectPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAwsTgwConnectPeer {
		return vs[0].([]*AviatrixAwsTgwConnectPeer)[vs[1].(int)]
	}).(AviatrixAwsTgwConnectPeerOutput)
}

type AviatrixAwsTgwConnectPeerMapOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwConnectPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwConnectPeer)(nil)).Elem()
}

func (o AviatrixAwsTgwConnectPeerMapOutput) ToAviatrixAwsTgwConnectPeerMapOutput() AviatrixAwsTgwConnectPeerMapOutput {
	return o
}

func (o AviatrixAwsTgwConnectPeerMapOutput) ToAviatrixAwsTgwConnectPeerMapOutputWithContext(ctx context.Context) AviatrixAwsTgwConnectPeerMapOutput {
	return o
}

func (o AviatrixAwsTgwConnectPeerMapOutput) MapIndex(k pulumi.StringInput) AviatrixAwsTgwConnectPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAwsTgwConnectPeer {
		return vs[0].(map[string]*AviatrixAwsTgwConnectPeer)[vs[1].(string)]
	}).(AviatrixAwsTgwConnectPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwConnectPeerInput)(nil)).Elem(), &AviatrixAwsTgwConnectPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwConnectPeerArrayInput)(nil)).Elem(), AviatrixAwsTgwConnectPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwConnectPeerMapInput)(nil)).Elem(), AviatrixAwsTgwConnectPeerMap{})
	pulumi.RegisterOutputType(AviatrixAwsTgwConnectPeerOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwConnectPeerArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwConnectPeerMapOutput{})
}
