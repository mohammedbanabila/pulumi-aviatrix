// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_vgw_conn** resource manages the connection between the Aviatrix transit gateway and AWS VGW for purposes of Transit Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix VGW Connection
 * const testVgwConn = new aviatrix.AviatrixVgwConn("test_vgw_conn", {
 *     bgpLocalAsNum: "65001",
 *     bgpVgwAccount: "dev-account-1",
 *     bgpVgwId: "vgw-abcd1234",
 *     bgpVgwRegion: "us-east-1",
 *     connName: "my-connection-vgw-to-tgw",
 *     gwName: "my-transit-gw",
 *     prependAsPaths: [
 *         "65001",
 *         "65001",
 *     ],
 *     vpcId: "vpc-abcd1234",
 * });
 * ```
 *
 * ## Import
 *
 * **vgw_conn** can be imported using the `conn_name` and `vpc_id`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixVgwConn:AviatrixVgwConn test conn_name~vpc_id
 * ```
 */
export class AviatrixVgwConn extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixVgwConn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixVgwConnState, opts?: pulumi.CustomResourceOptions): AviatrixVgwConn {
        return new AviatrixVgwConn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixVgwConn:AviatrixVgwConn';

    /**
     * Returns true if the given object is an instance of AviatrixVgwConn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixVgwConn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixVgwConn.__pulumiType;
    }

    /**
     * BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
     */
    public readonly bgpLocalAsNum!: pulumi.Output<string>;
    /**
     * Cloud Account used to create the AWS VGW that will be used for this connection. Example: "dev-account-1".
     */
    public readonly bgpVgwAccount!: pulumi.Output<string>;
    /**
     * ID of AWS VGW that will be used for this connection. Example: "vgw-abcd1234".
     */
    public readonly bgpVgwId!: pulumi.Output<string>;
    /**
     * Region of AWS VGW that will be used for this connection. Example: "us-east-1".
     */
    public readonly bgpVgwRegion!: pulumi.Output<string>;
    /**
     * The name of for Transit GW to VGW connection connection which is going to be created. Example: "my-connection-vgw-to-tgw".
     */
    public readonly connName!: pulumi.Output<string>;
    /**
     * Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
     */
    public readonly enableEventTriggeredHa!: pulumi.Output<boolean | undefined>;
    /**
     * Enable learned CIDRs approval for the connection. Requires the transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    public readonly enableLearnedCidrsApproval!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the Transit Gateway. Example: "my-transit-gw".
     */
    public readonly gwName!: pulumi.Output<string>;
    /**
     * Configure manual BGP advertised CIDRs for this connection. Available as of provider version R2.18+.
     */
    public readonly manualBgpAdvertisedCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
     */
    public readonly prependAsPaths!: pulumi.Output<string[] | undefined>;
    /**
     * VPC ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a AviatrixVgwConn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixVgwConnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixVgwConnArgs | AviatrixVgwConnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixVgwConnState | undefined;
            resourceInputs["bgpLocalAsNum"] = state ? state.bgpLocalAsNum : undefined;
            resourceInputs["bgpVgwAccount"] = state ? state.bgpVgwAccount : undefined;
            resourceInputs["bgpVgwId"] = state ? state.bgpVgwId : undefined;
            resourceInputs["bgpVgwRegion"] = state ? state.bgpVgwRegion : undefined;
            resourceInputs["connName"] = state ? state.connName : undefined;
            resourceInputs["enableEventTriggeredHa"] = state ? state.enableEventTriggeredHa : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = state ? state.enableLearnedCidrsApproval : undefined;
            resourceInputs["gwName"] = state ? state.gwName : undefined;
            resourceInputs["manualBgpAdvertisedCidrs"] = state ? state.manualBgpAdvertisedCidrs : undefined;
            resourceInputs["prependAsPaths"] = state ? state.prependAsPaths : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as AviatrixVgwConnArgs | undefined;
            if ((!args || args.bgpLocalAsNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpLocalAsNum'");
            }
            if ((!args || args.bgpVgwAccount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpVgwAccount'");
            }
            if ((!args || args.bgpVgwId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpVgwId'");
            }
            if ((!args || args.bgpVgwRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpVgwRegion'");
            }
            if ((!args || args.connName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connName'");
            }
            if ((!args || args.gwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gwName'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["bgpLocalAsNum"] = args ? args.bgpLocalAsNum : undefined;
            resourceInputs["bgpVgwAccount"] = args ? args.bgpVgwAccount : undefined;
            resourceInputs["bgpVgwId"] = args ? args.bgpVgwId : undefined;
            resourceInputs["bgpVgwRegion"] = args ? args.bgpVgwRegion : undefined;
            resourceInputs["connName"] = args ? args.connName : undefined;
            resourceInputs["enableEventTriggeredHa"] = args ? args.enableEventTriggeredHa : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = args ? args.enableLearnedCidrsApproval : undefined;
            resourceInputs["gwName"] = args ? args.gwName : undefined;
            resourceInputs["manualBgpAdvertisedCidrs"] = args ? args.manualBgpAdvertisedCidrs : undefined;
            resourceInputs["prependAsPaths"] = args ? args.prependAsPaths : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixVgwConn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixVgwConn resources.
 */
export interface AviatrixVgwConnState {
    /**
     * BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
     */
    bgpLocalAsNum?: pulumi.Input<string>;
    /**
     * Cloud Account used to create the AWS VGW that will be used for this connection. Example: "dev-account-1".
     */
    bgpVgwAccount?: pulumi.Input<string>;
    /**
     * ID of AWS VGW that will be used for this connection. Example: "vgw-abcd1234".
     */
    bgpVgwId?: pulumi.Input<string>;
    /**
     * Region of AWS VGW that will be used for this connection. Example: "us-east-1".
     */
    bgpVgwRegion?: pulumi.Input<string>;
    /**
     * The name of for Transit GW to VGW connection connection which is going to be created. Example: "my-connection-vgw-to-tgw".
     */
    connName?: pulumi.Input<string>;
    /**
     * Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
     */
    enableEventTriggeredHa?: pulumi.Input<boolean>;
    /**
     * Enable learned CIDRs approval for the connection. Requires the transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * Name of the Transit Gateway. Example: "my-transit-gw".
     */
    gwName?: pulumi.Input<string>;
    /**
     * Configure manual BGP advertised CIDRs for this connection. Available as of provider version R2.18+.
     */
    manualBgpAdvertisedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
     */
    prependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VPC ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixVgwConn resource.
 */
export interface AviatrixVgwConnArgs {
    /**
     * BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
     */
    bgpLocalAsNum: pulumi.Input<string>;
    /**
     * Cloud Account used to create the AWS VGW that will be used for this connection. Example: "dev-account-1".
     */
    bgpVgwAccount: pulumi.Input<string>;
    /**
     * ID of AWS VGW that will be used for this connection. Example: "vgw-abcd1234".
     */
    bgpVgwId: pulumi.Input<string>;
    /**
     * Region of AWS VGW that will be used for this connection. Example: "us-east-1".
     */
    bgpVgwRegion: pulumi.Input<string>;
    /**
     * The name of for Transit GW to VGW connection connection which is going to be created. Example: "my-connection-vgw-to-tgw".
     */
    connName: pulumi.Input<string>;
    /**
     * Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
     */
    enableEventTriggeredHa?: pulumi.Input<boolean>;
    /**
     * Enable learned CIDRs approval for the connection. Requires the transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * Name of the Transit Gateway. Example: "my-transit-gw".
     */
    gwName: pulumi.Input<string>;
    /**
     * Configure manual BGP advertised CIDRs for this connection. Available as of provider version R2.18+.
     */
    manualBgpAdvertisedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
     */
    prependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VPC ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".
     */
    vpcId: pulumi.Input<string>;
}
