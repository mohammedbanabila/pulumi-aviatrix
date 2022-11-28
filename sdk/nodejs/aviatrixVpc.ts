// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The **aviatrix_vpc** resource allows the creation and management of Aviatrix-created VPCs of various cloud types.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an AWS VPC
 * const awsVpc = new aviatrix.AviatrixVpc("aws_vpc", {
 *     accountName: "devops",
 *     aviatrixFirenetVpc: false,
 *     aviatrixTransitVpc: false,
 *     cidr: "10.0.0.0/16",
 *     cloudType: 1,
 *     region: "us-west-1",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create a GCP VPC
 * const gcpVpc = new aviatrix.AviatrixVpc("gcp_vpc", {
 *     accountName: "devops",
 *     cloudType: 4,
 *     subnets: [
 *         {
 *             cidr: "10.10.0.0/24",
 *             name: "subnet-1",
 *             region: "us-west1",
 *         },
 *         {
 *             cidr: "10.11.0.0/24",
 *             name: "subnet-2",
 *             region: "us-west2",
 *         },
 *     ],
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Azure VNet
 * const azureVnet = new aviatrix.AviatrixVpc("azure_vnet", {
 *     accountName: "devops",
 *     aviatrixFirenetVpc: false,
 *     cidr: "12.0.0.0/16",
 *     cloudType: 8,
 *     region: "Central US",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an OCI VPC
 * const ociVpc = new aviatrix.AviatrixVpc("oci_vpc", {
 *     accountName: "devops",
 *     cidr: "10.0.0.0/24",
 *     cloudType: 16,
 *     region: "us-ashburn-1",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an AzureGov VNet
 * const azureVnet = new aviatrix.AviatrixVpc("azure_vnet", {
 *     accountName: "devops",
 *     aviatrixFirenetVpc: false,
 *     cidr: "12.0.0.0/16",
 *     cloudType: 32,
 *     region: "USGov Arizona",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an AWSGov VPC
 * const awsgovVnet = new aviatrix.AviatrixVpc("awsgov_vnet", {
 *     accountName: "devops",
 *     aviatrixFirenetVpc: false,
 *     aviatrixTransitVpc: false,
 *     cidr: "12.0.0.0/20",
 *     cloudType: 256,
 *     region: "us-gov-west-1",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an AWS China VPC
 * const awsChinaVnet = new aviatrix.AviatrixVpc("aws_china_vnet", {
 *     accountName: "devops",
 *     aviatrixTransitVpc: false,
 *     cidr: "12.0.0.0/20",
 *     cloudType: 1024,
 *     region: "cn-north-1",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Azure China VNet
 * const azureChinaVnet = new aviatrix.AviatrixVpc("azure_china_vnet", {
 *     accountName: "devops",
 *     cidr: "12.0.0.0/16",
 *     cloudType: 2048,
 *     region: "China North",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Alibaba Cloud VPC
 * const aliyunVpc = new aviatrix.AviatrixVpc("aliyun_vpc", {
 *     accountName: "devops",
 *     cidr: "10.0.0.0/20",
 *     cloudType: 8192,
 *     region: "acs-us-west-1 (Silicon Valley)",
 * });
 * ```
 *
 * ## Import
 *
 * **vpc** can be imported using the VPC's `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixVpc:AviatrixVpc test name
 * ```
 */
export class AviatrixVpc extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixVpc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixVpcState, opts?: pulumi.CustomResourceOptions): AviatrixVpc {
        return new AviatrixVpc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixVpc:AviatrixVpc';

    /**
     * Returns true if the given object is an instance of AviatrixVpc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixVpc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixVpc.__pulumiType;
    }

    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * List of OCI availability domains.
     */
    public /*out*/ readonly availabilityDomains!: pulumi.Output<string[]>;
    /**
     * Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
     */
    public readonly aviatrixFirenetVpc!: pulumi.Output<boolean | undefined>;
    /**
     * Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
     */
    public readonly aviatrixTransitVpc!: pulumi.Output<boolean | undefined>;
    /**
     * Azure VNet resource ID.
     */
    public /*out*/ readonly azureVnetResourceId!: pulumi.Output<string>;
    /**
     * CIDR block.
     */
    public readonly cidr!: pulumi.Output<string | undefined>;
    /**
     * Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
     */
    public readonly cloudType!: pulumi.Output<number>;
    /**
     * Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
     */
    public readonly enableNativeGwlb!: pulumi.Output<boolean | undefined>;
    /**
     * Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    public readonly enablePrivateOobSubnet!: pulumi.Output<boolean | undefined>;
    /**
     * List of OCI fault domains.
     */
    public /*out*/ readonly faultDomains!: pulumi.Output<string[]>;
    /**
     * Name of this subnet.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
     */
    public readonly numOfSubnetPairs!: pulumi.Output<number | undefined>;
    /**
     * Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
     */
    public readonly privateModeSubnets!: pulumi.Output<boolean | undefined>;
    /**
     * List of private subnet of the VPC(AWS, Azure) to be created.
     */
    public /*out*/ readonly privateSubnets!: pulumi.Output<outputs.AviatrixVpcPrivateSubnet[]>;
    /**
     * List of public subnet of the VPC(AWS, Azure) to be created.
     */
    public /*out*/ readonly publicSubnets!: pulumi.Output<outputs.AviatrixVpcPublicSubnet[]>;
    /**
     * Region of this subnet.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
     */
    public readonly resourceGroup!: pulumi.Output<string>;
    /**
     * List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
     */
    public /*out*/ readonly routeTables!: pulumi.Output<string[]>;
    /**
     * Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
     */
    public readonly subnetSize!: pulumi.Output<number | undefined>;
    /**
     * List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
     */
    public readonly subnets!: pulumi.Output<outputs.AviatrixVpcSubnet[]>;
    /**
     * ID of the VPC to be created.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a AviatrixVpc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixVpcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixVpcArgs | AviatrixVpcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixVpcState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["availabilityDomains"] = state ? state.availabilityDomains : undefined;
            resourceInputs["aviatrixFirenetVpc"] = state ? state.aviatrixFirenetVpc : undefined;
            resourceInputs["aviatrixTransitVpc"] = state ? state.aviatrixTransitVpc : undefined;
            resourceInputs["azureVnetResourceId"] = state ? state.azureVnetResourceId : undefined;
            resourceInputs["cidr"] = state ? state.cidr : undefined;
            resourceInputs["cloudType"] = state ? state.cloudType : undefined;
            resourceInputs["enableNativeGwlb"] = state ? state.enableNativeGwlb : undefined;
            resourceInputs["enablePrivateOobSubnet"] = state ? state.enablePrivateOobSubnet : undefined;
            resourceInputs["faultDomains"] = state ? state.faultDomains : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numOfSubnetPairs"] = state ? state.numOfSubnetPairs : undefined;
            resourceInputs["privateModeSubnets"] = state ? state.privateModeSubnets : undefined;
            resourceInputs["privateSubnets"] = state ? state.privateSubnets : undefined;
            resourceInputs["publicSubnets"] = state ? state.publicSubnets : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["resourceGroup"] = state ? state.resourceGroup : undefined;
            resourceInputs["routeTables"] = state ? state.routeTables : undefined;
            resourceInputs["subnetSize"] = state ? state.subnetSize : undefined;
            resourceInputs["subnets"] = state ? state.subnets : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as AviatrixVpcArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.cloudType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudType'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["aviatrixFirenetVpc"] = args ? args.aviatrixFirenetVpc : undefined;
            resourceInputs["aviatrixTransitVpc"] = args ? args.aviatrixTransitVpc : undefined;
            resourceInputs["cidr"] = args ? args.cidr : undefined;
            resourceInputs["cloudType"] = args ? args.cloudType : undefined;
            resourceInputs["enableNativeGwlb"] = args ? args.enableNativeGwlb : undefined;
            resourceInputs["enablePrivateOobSubnet"] = args ? args.enablePrivateOobSubnet : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["numOfSubnetPairs"] = args ? args.numOfSubnetPairs : undefined;
            resourceInputs["privateModeSubnets"] = args ? args.privateModeSubnets : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            resourceInputs["subnetSize"] = args ? args.subnetSize : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["availabilityDomains"] = undefined /*out*/;
            resourceInputs["azureVnetResourceId"] = undefined /*out*/;
            resourceInputs["faultDomains"] = undefined /*out*/;
            resourceInputs["privateSubnets"] = undefined /*out*/;
            resourceInputs["publicSubnets"] = undefined /*out*/;
            resourceInputs["routeTables"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixVpc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixVpc resources.
 */
export interface AviatrixVpcState {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName?: pulumi.Input<string>;
    /**
     * List of OCI availability domains.
     */
    availabilityDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
     */
    aviatrixFirenetVpc?: pulumi.Input<boolean>;
    /**
     * Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
     */
    aviatrixTransitVpc?: pulumi.Input<boolean>;
    /**
     * Azure VNet resource ID.
     */
    azureVnetResourceId?: pulumi.Input<string>;
    /**
     * CIDR block.
     */
    cidr?: pulumi.Input<string>;
    /**
     * Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
     */
    cloudType?: pulumi.Input<number>;
    /**
     * Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
     */
    enableNativeGwlb?: pulumi.Input<boolean>;
    /**
     * Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    enablePrivateOobSubnet?: pulumi.Input<boolean>;
    /**
     * List of OCI fault domains.
     */
    faultDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of this subnet.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
     */
    numOfSubnetPairs?: pulumi.Input<number>;
    /**
     * Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
     */
    privateModeSubnets?: pulumi.Input<boolean>;
    /**
     * List of private subnet of the VPC(AWS, Azure) to be created.
     */
    privateSubnets?: pulumi.Input<pulumi.Input<inputs.AviatrixVpcPrivateSubnet>[]>;
    /**
     * List of public subnet of the VPC(AWS, Azure) to be created.
     */
    publicSubnets?: pulumi.Input<pulumi.Input<inputs.AviatrixVpcPublicSubnet>[]>;
    /**
     * Region of this subnet.
     */
    region?: pulumi.Input<string>;
    /**
     * The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
     */
    resourceGroup?: pulumi.Input<string>;
    /**
     * List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
     */
    routeTables?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
     */
    subnetSize?: pulumi.Input<number>;
    /**
     * List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
     */
    subnets?: pulumi.Input<pulumi.Input<inputs.AviatrixVpcSubnet>[]>;
    /**
     * ID of the VPC to be created.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixVpc resource.
 */
export interface AviatrixVpcArgs {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName: pulumi.Input<string>;
    /**
     * Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
     */
    aviatrixFirenetVpc?: pulumi.Input<boolean>;
    /**
     * Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
     */
    aviatrixTransitVpc?: pulumi.Input<boolean>;
    /**
     * CIDR block.
     */
    cidr?: pulumi.Input<string>;
    /**
     * Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
     */
    cloudType: pulumi.Input<number>;
    /**
     * Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloudType = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
     */
    enableNativeGwlb?: pulumi.Input<boolean>;
    /**
     * Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    enablePrivateOobSubnet?: pulumi.Input<boolean>;
    /**
     * Name of this subnet.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
     */
    numOfSubnetPairs?: pulumi.Input<number>;
    /**
     * Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
     */
    privateModeSubnets?: pulumi.Input<boolean>;
    /**
     * Region of this subnet.
     */
    region?: pulumi.Input<string>;
    /**
     * The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
     */
    resourceGroup?: pulumi.Input<string>;
    /**
     * Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
     */
    subnetSize?: pulumi.Input<number>;
    /**
     * List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
     */
    subnets?: pulumi.Input<pulumi.Input<inputs.AviatrixVpcSubnet>[]>;
}
