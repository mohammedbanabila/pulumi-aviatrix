// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_aws_peer** resource allows the creation and management of Aviatrix-created native AWS intra and inter-region VPC peerings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix AWS Peering
 * const testAwspeer = new aviatrix.AviatrixAwsPeer("test_awspeer", {
 *     accountName1: "test1-account",
 *     accountName2: "test2-account",
 *     rtbList1s: ["rtb-abcd1234"],
 *     rtbList2s: ["rtb-wxyz5678"],
 *     vpcId1: "vpc-abcd1234",
 *     vpcId2: "vpc-rdef3333",
 *     vpcReg1: "us-east-1",
 *     vpcReg2: "us-west-1",
 * });
 * ```
 *
 * ## Import
 *
 * **aws_peer** can be imported using the `vpc_id1` and `vpc_id2`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixAwsPeer:AviatrixAwsPeer test vpc_id1~vpc_id2
 * ```
 */
export class AviatrixAwsPeer extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsPeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsPeerState, opts?: pulumi.CustomResourceOptions): AviatrixAwsPeer {
        return new AviatrixAwsPeer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsPeer:AviatrixAwsPeer';

    /**
     * Returns true if the given object is an instance of AviatrixAwsPeer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsPeer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsPeer.__pulumiType;
    }

    /**
     * This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
     */
    public readonly accountName1!: pulumi.Output<string>;
    /**
     * This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
     */
    public readonly accountName2!: pulumi.Output<string>;
    /**
     * List of route table ID of vpc_id1.
     */
    public /*out*/ readonly rtbList1Outputs!: pulumi.Output<string[]>;
    /**
     * List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"].
     */
    public readonly rtbList1s!: pulumi.Output<string[] | undefined>;
    /**
     * List of route table ID of vpc_id2.
     */
    public /*out*/ readonly rtbList2Outputs!: pulumi.Output<string[]>;
    /**
     * List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"].
     */
    public readonly rtbList2s!: pulumi.Output<string[] | undefined>;
    /**
     * VPC ID of AWS cloud. Example: AWS: "vpc-abcd1234".
     */
    public readonly vpcId1!: pulumi.Output<string>;
    /**
     * VPC ID of AWS cloud. Example: AWS: "vpc-abcd1234".
     */
    public readonly vpcId2!: pulumi.Output<string>;
    /**
     * Region of AWS cloud. Example: AWS: "us-east-1".
     */
    public readonly vpcReg1!: pulumi.Output<string>;
    /**
     * Region of AWS cloud. Example: AWS: "us-east-1".
     */
    public readonly vpcReg2!: pulumi.Output<string>;

    /**
     * Create a AviatrixAwsPeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsPeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsPeerArgs | AviatrixAwsPeerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsPeerState | undefined;
            resourceInputs["accountName1"] = state ? state.accountName1 : undefined;
            resourceInputs["accountName2"] = state ? state.accountName2 : undefined;
            resourceInputs["rtbList1Outputs"] = state ? state.rtbList1Outputs : undefined;
            resourceInputs["rtbList1s"] = state ? state.rtbList1s : undefined;
            resourceInputs["rtbList2Outputs"] = state ? state.rtbList2Outputs : undefined;
            resourceInputs["rtbList2s"] = state ? state.rtbList2s : undefined;
            resourceInputs["vpcId1"] = state ? state.vpcId1 : undefined;
            resourceInputs["vpcId2"] = state ? state.vpcId2 : undefined;
            resourceInputs["vpcReg1"] = state ? state.vpcReg1 : undefined;
            resourceInputs["vpcReg2"] = state ? state.vpcReg2 : undefined;
        } else {
            const args = argsOrState as AviatrixAwsPeerArgs | undefined;
            if ((!args || args.accountName1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName1'");
            }
            if ((!args || args.accountName2 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName2'");
            }
            if ((!args || args.vpcId1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId1'");
            }
            if ((!args || args.vpcId2 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId2'");
            }
            if ((!args || args.vpcReg1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcReg1'");
            }
            if ((!args || args.vpcReg2 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcReg2'");
            }
            resourceInputs["accountName1"] = args ? args.accountName1 : undefined;
            resourceInputs["accountName2"] = args ? args.accountName2 : undefined;
            resourceInputs["rtbList1s"] = args ? args.rtbList1s : undefined;
            resourceInputs["rtbList2s"] = args ? args.rtbList2s : undefined;
            resourceInputs["vpcId1"] = args ? args.vpcId1 : undefined;
            resourceInputs["vpcId2"] = args ? args.vpcId2 : undefined;
            resourceInputs["vpcReg1"] = args ? args.vpcReg1 : undefined;
            resourceInputs["vpcReg2"] = args ? args.vpcReg2 : undefined;
            resourceInputs["rtbList1Outputs"] = undefined /*out*/;
            resourceInputs["rtbList2Outputs"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsPeer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsPeer resources.
 */
export interface AviatrixAwsPeerState {
    /**
     * This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
     */
    accountName1?: pulumi.Input<string>;
    /**
     * This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
     */
    accountName2?: pulumi.Input<string>;
    /**
     * List of route table ID of vpc_id1.
     */
    rtbList1Outputs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"].
     */
    rtbList1s?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of route table ID of vpc_id2.
     */
    rtbList2Outputs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"].
     */
    rtbList2s?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VPC ID of AWS cloud. Example: AWS: "vpc-abcd1234".
     */
    vpcId1?: pulumi.Input<string>;
    /**
     * VPC ID of AWS cloud. Example: AWS: "vpc-abcd1234".
     */
    vpcId2?: pulumi.Input<string>;
    /**
     * Region of AWS cloud. Example: AWS: "us-east-1".
     */
    vpcReg1?: pulumi.Input<string>;
    /**
     * Region of AWS cloud. Example: AWS: "us-east-1".
     */
    vpcReg2?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAwsPeer resource.
 */
export interface AviatrixAwsPeerArgs {
    /**
     * This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
     */
    accountName1: pulumi.Input<string>;
    /**
     * This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
     */
    accountName2: pulumi.Input<string>;
    /**
     * List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"].
     */
    rtbList1s?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"].
     */
    rtbList2s?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VPC ID of AWS cloud. Example: AWS: "vpc-abcd1234".
     */
    vpcId1: pulumi.Input<string>;
    /**
     * VPC ID of AWS cloud. Example: AWS: "vpc-abcd1234".
     */
    vpcId2: pulumi.Input<string>;
    /**
     * Region of AWS cloud. Example: AWS: "us-east-1".
     */
    vpcReg1: pulumi.Input<string>;
    /**
     * Region of AWS cloud. Example: AWS: "us-east-1".
     */
    vpcReg2: pulumi.Input<string>;
}
