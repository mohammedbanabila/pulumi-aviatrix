// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_vpc** resource allows the creation and management of Aviatrix-created VPCs of various cloud types.
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "awsVpc", &aviatrix.AviatrixVpcArgs{
//				AccountName:        pulumi.String("devops"),
//				AviatrixFirenetVpc: pulumi.Bool(false),
//				AviatrixTransitVpc: pulumi.Bool(false),
//				Cidr:               pulumi.String("10.0.0.0/16"),
//				CloudType:          pulumi.Int(1),
//				Region:             pulumi.String("us-west-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "gcpVpc", &aviatrix.AviatrixVpcArgs{
//				AccountName: pulumi.String("devops"),
//				CloudType:   pulumi.Int(4),
//				Subnets: AviatrixVpcSubnetArray{
//					&AviatrixVpcSubnetArgs{
//						Cidr:   pulumi.String("10.10.0.0/24"),
//						Name:   pulumi.String("subnet-1"),
//						Region: pulumi.String("us-west1"),
//					},
//					&AviatrixVpcSubnetArgs{
//						Cidr:   pulumi.String("10.11.0.0/24"),
//						Name:   pulumi.String("subnet-2"),
//						Region: pulumi.String("us-west2"),
//					},
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "azureVnet", &aviatrix.AviatrixVpcArgs{
//				AccountName:        pulumi.String("devops"),
//				AviatrixFirenetVpc: pulumi.Bool(false),
//				Cidr:               pulumi.String("12.0.0.0/16"),
//				CloudType:          pulumi.Int(8),
//				Region:             pulumi.String("Central US"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "ociVpc", &aviatrix.AviatrixVpcArgs{
//				AccountName: pulumi.String("devops"),
//				Cidr:        pulumi.String("10.0.0.0/24"),
//				CloudType:   pulumi.Int(16),
//				Region:      pulumi.String("us-ashburn-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "azureVnet", &aviatrix.AviatrixVpcArgs{
//				AccountName:        pulumi.String("devops"),
//				AviatrixFirenetVpc: pulumi.Bool(false),
//				Cidr:               pulumi.String("12.0.0.0/16"),
//				CloudType:          pulumi.Int(32),
//				Region:             pulumi.String("USGov Arizona"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "awsgovVnet", &aviatrix.AviatrixVpcArgs{
//				AccountName:        pulumi.String("devops"),
//				AviatrixFirenetVpc: pulumi.Bool(false),
//				AviatrixTransitVpc: pulumi.Bool(false),
//				Cidr:               pulumi.String("12.0.0.0/20"),
//				CloudType:          pulumi.Int(256),
//				Region:             pulumi.String("us-gov-west-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "awsChinaVnet", &aviatrix.AviatrixVpcArgs{
//				AccountName:        pulumi.String("devops"),
//				AviatrixTransitVpc: pulumi.Bool(false),
//				Cidr:               pulumi.String("12.0.0.0/20"),
//				CloudType:          pulumi.Int(1024),
//				Region:             pulumi.String("cn-north-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "azureChinaVnet", &aviatrix.AviatrixVpcArgs{
//				AccountName: pulumi.String("devops"),
//				Cidr:        pulumi.String("12.0.0.0/16"),
//				CloudType:   pulumi.Int(2048),
//				Region:      pulumi.String("China North"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixVpc(ctx, "aliyunVpc", &aviatrix.AviatrixVpcArgs{
//				AccountName: pulumi.String("devops"),
//				Cidr:        pulumi.String("10.0.0.0/20"),
//				CloudType:   pulumi.Int(8192),
//				Region:      pulumi.String("acs-us-west-1 (Silicon Valley)"),
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
// **vpc** can be imported using the VPC's `name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixVpc:AviatrixVpc test name
//
// ```
type AviatrixVpc struct {
	pulumi.CustomResourceState

	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// List of OCI availability domains.
	AvailabilityDomains pulumi.StringArrayOutput `pulumi:"availabilityDomains"`
	// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixFirenetVpc pulumi.BoolPtrOutput `pulumi:"aviatrixFirenetVpc"`
	// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixTransitVpc pulumi.BoolPtrOutput `pulumi:"aviatrixTransitVpc"`
	// Azure VNet resource ID.
	AzureVnetResourceId pulumi.StringOutput `pulumi:"azureVnetResourceId"`
	// CIDR block.
	Cidr pulumi.StringPtrOutput `pulumi:"cidr"`
	// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
	CloudType pulumi.IntOutput `pulumi:"cloudType"`
	// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	EnableNativeGwlb pulumi.BoolPtrOutput `pulumi:"enableNativeGwlb"`
	// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
	EnablePrivateOobSubnet pulumi.BoolPtrOutput `pulumi:"enablePrivateOobSubnet"`
	// List of OCI fault domains.
	FaultDomains pulumi.StringArrayOutput `pulumi:"faultDomains"`
	// Name of this subnet.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
	NumOfSubnetPairs pulumi.IntPtrOutput `pulumi:"numOfSubnetPairs"`
	// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
	PrivateModeSubnets pulumi.BoolPtrOutput `pulumi:"privateModeSubnets"`
	// List of private subnet of the VPC(AWS, Azure) to be created.
	PrivateSubnets AviatrixVpcPrivateSubnetArrayOutput `pulumi:"privateSubnets"`
	// List of public subnet of the VPC(AWS, Azure) to be created.
	PublicSubnets AviatrixVpcPublicSubnetArrayOutput `pulumi:"publicSubnets"`
	// Region of this subnet.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
	RouteTables pulumi.StringArrayOutput `pulumi:"routeTables"`
	// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
	SubnetSize pulumi.IntPtrOutput `pulumi:"subnetSize"`
	// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
	Subnets AviatrixVpcSubnetArrayOutput `pulumi:"subnets"`
	// ID of the VPC to be created.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewAviatrixVpc registers a new resource with the given unique name, arguments, and options.
func NewAviatrixVpc(ctx *pulumi.Context,
	name string, args *AviatrixVpcArgs, opts ...pulumi.ResourceOption) (*AviatrixVpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.CloudType == nil {
		return nil, errors.New("invalid value for required argument 'CloudType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixVpc
	err := ctx.RegisterResource("aviatrix:index/aviatrixVpc:AviatrixVpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixVpc gets an existing AviatrixVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixVpcState, opts ...pulumi.ResourceOption) (*AviatrixVpc, error) {
	var resource AviatrixVpc
	err := ctx.ReadResource("aviatrix:index/aviatrixVpc:AviatrixVpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixVpc resources.
type aviatrixVpcState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName *string `pulumi:"accountName"`
	// List of OCI availability domains.
	AvailabilityDomains []string `pulumi:"availabilityDomains"`
	// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixFirenetVpc *bool `pulumi:"aviatrixFirenetVpc"`
	// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixTransitVpc *bool `pulumi:"aviatrixTransitVpc"`
	// Azure VNet resource ID.
	AzureVnetResourceId *string `pulumi:"azureVnetResourceId"`
	// CIDR block.
	Cidr *string `pulumi:"cidr"`
	// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
	CloudType *int `pulumi:"cloudType"`
	// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	EnableNativeGwlb *bool `pulumi:"enableNativeGwlb"`
	// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
	EnablePrivateOobSubnet *bool `pulumi:"enablePrivateOobSubnet"`
	// List of OCI fault domains.
	FaultDomains []string `pulumi:"faultDomains"`
	// Name of this subnet.
	Name *string `pulumi:"name"`
	// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
	NumOfSubnetPairs *int `pulumi:"numOfSubnetPairs"`
	// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
	PrivateModeSubnets *bool `pulumi:"privateModeSubnets"`
	// List of private subnet of the VPC(AWS, Azure) to be created.
	PrivateSubnets []AviatrixVpcPrivateSubnet `pulumi:"privateSubnets"`
	// List of public subnet of the VPC(AWS, Azure) to be created.
	PublicSubnets []AviatrixVpcPublicSubnet `pulumi:"publicSubnets"`
	// Region of this subnet.
	Region *string `pulumi:"region"`
	// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
	RouteTables []string `pulumi:"routeTables"`
	// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
	SubnetSize *int `pulumi:"subnetSize"`
	// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
	Subnets []AviatrixVpcSubnet `pulumi:"subnets"`
	// ID of the VPC to be created.
	VpcId *string `pulumi:"vpcId"`
}

type AviatrixVpcState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringPtrInput
	// List of OCI availability domains.
	AvailabilityDomains pulumi.StringArrayInput
	// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixFirenetVpc pulumi.BoolPtrInput
	// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixTransitVpc pulumi.BoolPtrInput
	// Azure VNet resource ID.
	AzureVnetResourceId pulumi.StringPtrInput
	// CIDR block.
	Cidr pulumi.StringPtrInput
	// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
	CloudType pulumi.IntPtrInput
	// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	EnableNativeGwlb pulumi.BoolPtrInput
	// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
	EnablePrivateOobSubnet pulumi.BoolPtrInput
	// List of OCI fault domains.
	FaultDomains pulumi.StringArrayInput
	// Name of this subnet.
	Name pulumi.StringPtrInput
	// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
	NumOfSubnetPairs pulumi.IntPtrInput
	// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
	PrivateModeSubnets pulumi.BoolPtrInput
	// List of private subnet of the VPC(AWS, Azure) to be created.
	PrivateSubnets AviatrixVpcPrivateSubnetArrayInput
	// List of public subnet of the VPC(AWS, Azure) to be created.
	PublicSubnets AviatrixVpcPublicSubnetArrayInput
	// Region of this subnet.
	Region pulumi.StringPtrInput
	// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
	ResourceGroup pulumi.StringPtrInput
	// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
	RouteTables pulumi.StringArrayInput
	// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
	SubnetSize pulumi.IntPtrInput
	// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
	Subnets AviatrixVpcSubnetArrayInput
	// ID of the VPC to be created.
	VpcId pulumi.StringPtrInput
}

func (AviatrixVpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixVpcState)(nil)).Elem()
}

type aviatrixVpcArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName string `pulumi:"accountName"`
	// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixFirenetVpc *bool `pulumi:"aviatrixFirenetVpc"`
	// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixTransitVpc *bool `pulumi:"aviatrixTransitVpc"`
	// CIDR block.
	Cidr *string `pulumi:"cidr"`
	// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
	CloudType int `pulumi:"cloudType"`
	// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	EnableNativeGwlb *bool `pulumi:"enableNativeGwlb"`
	// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
	EnablePrivateOobSubnet *bool `pulumi:"enablePrivateOobSubnet"`
	// Name of this subnet.
	Name *string `pulumi:"name"`
	// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
	NumOfSubnetPairs *int `pulumi:"numOfSubnetPairs"`
	// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
	PrivateModeSubnets *bool `pulumi:"privateModeSubnets"`
	// Region of this subnet.
	Region *string `pulumi:"region"`
	// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
	SubnetSize *int `pulumi:"subnetSize"`
	// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
	Subnets []AviatrixVpcSubnet `pulumi:"subnets"`
}

// The set of arguments for constructing a AviatrixVpc resource.
type AviatrixVpcArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringInput
	// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixFirenetVpc pulumi.BoolPtrInput
	// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
	AviatrixTransitVpc pulumi.BoolPtrInput
	// CIDR block.
	Cidr pulumi.StringPtrInput
	// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
	CloudType pulumi.IntInput
	// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	EnableNativeGwlb pulumi.BoolPtrInput
	// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
	EnablePrivateOobSubnet pulumi.BoolPtrInput
	// Name of this subnet.
	Name pulumi.StringPtrInput
	// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
	NumOfSubnetPairs pulumi.IntPtrInput
	// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
	PrivateModeSubnets pulumi.BoolPtrInput
	// Region of this subnet.
	Region pulumi.StringPtrInput
	// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
	ResourceGroup pulumi.StringPtrInput
	// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
	SubnetSize pulumi.IntPtrInput
	// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
	Subnets AviatrixVpcSubnetArrayInput
}

func (AviatrixVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixVpcArgs)(nil)).Elem()
}

type AviatrixVpcInput interface {
	pulumi.Input

	ToAviatrixVpcOutput() AviatrixVpcOutput
	ToAviatrixVpcOutputWithContext(ctx context.Context) AviatrixVpcOutput
}

func (*AviatrixVpc) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixVpc)(nil)).Elem()
}

func (i *AviatrixVpc) ToAviatrixVpcOutput() AviatrixVpcOutput {
	return i.ToAviatrixVpcOutputWithContext(context.Background())
}

func (i *AviatrixVpc) ToAviatrixVpcOutputWithContext(ctx context.Context) AviatrixVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpcOutput)
}

// AviatrixVpcArrayInput is an input type that accepts AviatrixVpcArray and AviatrixVpcArrayOutput values.
// You can construct a concrete instance of `AviatrixVpcArrayInput` via:
//
//	AviatrixVpcArray{ AviatrixVpcArgs{...} }
type AviatrixVpcArrayInput interface {
	pulumi.Input

	ToAviatrixVpcArrayOutput() AviatrixVpcArrayOutput
	ToAviatrixVpcArrayOutputWithContext(context.Context) AviatrixVpcArrayOutput
}

type AviatrixVpcArray []AviatrixVpcInput

func (AviatrixVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixVpc)(nil)).Elem()
}

func (i AviatrixVpcArray) ToAviatrixVpcArrayOutput() AviatrixVpcArrayOutput {
	return i.ToAviatrixVpcArrayOutputWithContext(context.Background())
}

func (i AviatrixVpcArray) ToAviatrixVpcArrayOutputWithContext(ctx context.Context) AviatrixVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpcArrayOutput)
}

// AviatrixVpcMapInput is an input type that accepts AviatrixVpcMap and AviatrixVpcMapOutput values.
// You can construct a concrete instance of `AviatrixVpcMapInput` via:
//
//	AviatrixVpcMap{ "key": AviatrixVpcArgs{...} }
type AviatrixVpcMapInput interface {
	pulumi.Input

	ToAviatrixVpcMapOutput() AviatrixVpcMapOutput
	ToAviatrixVpcMapOutputWithContext(context.Context) AviatrixVpcMapOutput
}

type AviatrixVpcMap map[string]AviatrixVpcInput

func (AviatrixVpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixVpc)(nil)).Elem()
}

func (i AviatrixVpcMap) ToAviatrixVpcMapOutput() AviatrixVpcMapOutput {
	return i.ToAviatrixVpcMapOutputWithContext(context.Background())
}

func (i AviatrixVpcMap) ToAviatrixVpcMapOutputWithContext(ctx context.Context) AviatrixVpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpcMapOutput)
}

type AviatrixVpcOutput struct{ *pulumi.OutputState }

func (AviatrixVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixVpc)(nil)).Elem()
}

func (o AviatrixVpcOutput) ToAviatrixVpcOutput() AviatrixVpcOutput {
	return o
}

func (o AviatrixVpcOutput) ToAviatrixVpcOutputWithContext(ctx context.Context) AviatrixVpcOutput {
	return o
}

// This parameter represents the name of a Cloud-Account in Aviatrix controller.
func (o AviatrixVpcOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// List of OCI availability domains.
func (o AviatrixVpcOutput) AvailabilityDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringArrayOutput { return v.AvailabilityDomains }).(pulumi.StringArrayOutput)
}

// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
func (o AviatrixVpcOutput) AviatrixFirenetVpc() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.BoolPtrOutput { return v.AviatrixFirenetVpc }).(pulumi.BoolPtrOutput)
}

// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
func (o AviatrixVpcOutput) AviatrixTransitVpc() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.BoolPtrOutput { return v.AviatrixTransitVpc }).(pulumi.BoolPtrOutput)
}

// Azure VNet resource ID.
func (o AviatrixVpcOutput) AzureVnetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringOutput { return v.AzureVnetResourceId }).(pulumi.StringOutput)
}

// CIDR block.
func (o AviatrixVpcOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringPtrOutput { return v.Cidr }).(pulumi.StringPtrOutput)
}

// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
func (o AviatrixVpcOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.IntOutput { return v.CloudType }).(pulumi.IntOutput)
}

// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
func (o AviatrixVpcOutput) EnableNativeGwlb() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.BoolPtrOutput { return v.EnableNativeGwlb }).(pulumi.BoolPtrOutput)
}

// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
func (o AviatrixVpcOutput) EnablePrivateOobSubnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.BoolPtrOutput { return v.EnablePrivateOobSubnet }).(pulumi.BoolPtrOutput)
}

// List of OCI fault domains.
func (o AviatrixVpcOutput) FaultDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringArrayOutput { return v.FaultDomains }).(pulumi.StringArrayOutput)
}

// Name of this subnet.
func (o AviatrixVpcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
func (o AviatrixVpcOutput) NumOfSubnetPairs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.IntPtrOutput { return v.NumOfSubnetPairs }).(pulumi.IntPtrOutput)
}

// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
func (o AviatrixVpcOutput) PrivateModeSubnets() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.BoolPtrOutput { return v.PrivateModeSubnets }).(pulumi.BoolPtrOutput)
}

// List of private subnet of the VPC(AWS, Azure) to be created.
func (o AviatrixVpcOutput) PrivateSubnets() AviatrixVpcPrivateSubnetArrayOutput {
	return o.ApplyT(func(v *AviatrixVpc) AviatrixVpcPrivateSubnetArrayOutput { return v.PrivateSubnets }).(AviatrixVpcPrivateSubnetArrayOutput)
}

// List of public subnet of the VPC(AWS, Azure) to be created.
func (o AviatrixVpcOutput) PublicSubnets() AviatrixVpcPublicSubnetArrayOutput {
	return o.ApplyT(func(v *AviatrixVpc) AviatrixVpcPublicSubnetArrayOutput { return v.PublicSubnets }).(AviatrixVpcPublicSubnetArrayOutput)
}

// Region of this subnet.
func (o AviatrixVpcOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
func (o AviatrixVpcOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
func (o AviatrixVpcOutput) RouteTables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringArrayOutput { return v.RouteTables }).(pulumi.StringArrayOutput)
}

// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
func (o AviatrixVpcOutput) SubnetSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.IntPtrOutput { return v.SubnetSize }).(pulumi.IntPtrOutput)
}

// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
func (o AviatrixVpcOutput) Subnets() AviatrixVpcSubnetArrayOutput {
	return o.ApplyT(func(v *AviatrixVpc) AviatrixVpcSubnetArrayOutput { return v.Subnets }).(AviatrixVpcSubnetArrayOutput)
}

// ID of the VPC to be created.
func (o AviatrixVpcOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixVpc) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type AviatrixVpcArrayOutput struct{ *pulumi.OutputState }

func (AviatrixVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixVpc)(nil)).Elem()
}

func (o AviatrixVpcArrayOutput) ToAviatrixVpcArrayOutput() AviatrixVpcArrayOutput {
	return o
}

func (o AviatrixVpcArrayOutput) ToAviatrixVpcArrayOutputWithContext(ctx context.Context) AviatrixVpcArrayOutput {
	return o
}

func (o AviatrixVpcArrayOutput) Index(i pulumi.IntInput) AviatrixVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixVpc {
		return vs[0].([]*AviatrixVpc)[vs[1].(int)]
	}).(AviatrixVpcOutput)
}

type AviatrixVpcMapOutput struct{ *pulumi.OutputState }

func (AviatrixVpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixVpc)(nil)).Elem()
}

func (o AviatrixVpcMapOutput) ToAviatrixVpcMapOutput() AviatrixVpcMapOutput {
	return o
}

func (o AviatrixVpcMapOutput) ToAviatrixVpcMapOutputWithContext(ctx context.Context) AviatrixVpcMapOutput {
	return o
}

func (o AviatrixVpcMapOutput) MapIndex(k pulumi.StringInput) AviatrixVpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixVpc {
		return vs[0].(map[string]*AviatrixVpc)[vs[1].(string)]
	}).(AviatrixVpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpcInput)(nil)).Elem(), &AviatrixVpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpcArrayInput)(nil)).Elem(), AviatrixVpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpcMapInput)(nil)).Elem(), AviatrixVpcMap{})
	pulumi.RegisterOutputType(AviatrixVpcOutput{})
	pulumi.RegisterOutputType(AviatrixVpcArrayOutput{})
	pulumi.RegisterOutputType(AviatrixVpcMapOutput{})
}
