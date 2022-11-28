// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix Spoke Transit Attachment
 * const testAttachment = new aviatrix.AviatrixSpokeTransitAttachment("test_attachment", {
 *     routeTables: [
 *         "rtb-737d540c",
 *         "rtb-626d045c",
 *     ],
 *     spokeGwName: "spoke-gw",
 *     transitGwName: "transit-gw",
 * });
 * ```
 *
 * ## Import
 *
 * **spoke_transit_attachment** can be imported using the `spoke_gw_name` and `transit_gw_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixSpokeTransitAttachment:AviatrixSpokeTransitAttachment test spoke_gw_name~transit_gw_name
 * ```
 */
export class AviatrixSpokeTransitAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixSpokeTransitAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixSpokeTransitAttachmentState, opts?: pulumi.CustomResourceOptions): AviatrixSpokeTransitAttachment {
        return new AviatrixSpokeTransitAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixSpokeTransitAttachment:AviatrixSpokeTransitAttachment';

    /**
     * Returns true if the given object is an instance of AviatrixSpokeTransitAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixSpokeTransitAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixSpokeTransitAttachment.__pulumiType;
    }

    /**
     * Indicates whether the maximum amount of HPE tunnels will be created. Only valid when transit and spoke gateways are each launched in Insane Mode and in the same cloud type. Default value: true. Available as of provider version R2.22.2+.
     */
    public readonly enableMaxPerformance!: pulumi.Output<boolean | undefined>;
    /**
     * Learned routes will be propagated to these route tables. Example: ["rtb-212ff547","rtb-04539787"].
     */
    public readonly routeTables!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates whether the spoke gateway is BGP enabled or not.
     */
    public /*out*/ readonly spokeBgpEnabled!: pulumi.Output<boolean>;
    /**
     * Name of the spoke gateway to attach to transit network.
     */
    public readonly spokeGwName!: pulumi.Output<string>;
    /**
     * Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on spoke_gateway_name. Available as of provider version R2.23+.
     */
    public readonly spokePrependAsPaths!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the transit gateway to attach the spoke gateway to.
     */
    public readonly transitGwName!: pulumi.Output<string>;
    /**
     * Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on transit_gateway_name. Available as of provider version R2.23+.
     */
    public readonly transitPrependAsPaths!: pulumi.Output<string[] | undefined>;

    /**
     * Create a AviatrixSpokeTransitAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixSpokeTransitAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixSpokeTransitAttachmentArgs | AviatrixSpokeTransitAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixSpokeTransitAttachmentState | undefined;
            resourceInputs["enableMaxPerformance"] = state ? state.enableMaxPerformance : undefined;
            resourceInputs["routeTables"] = state ? state.routeTables : undefined;
            resourceInputs["spokeBgpEnabled"] = state ? state.spokeBgpEnabled : undefined;
            resourceInputs["spokeGwName"] = state ? state.spokeGwName : undefined;
            resourceInputs["spokePrependAsPaths"] = state ? state.spokePrependAsPaths : undefined;
            resourceInputs["transitGwName"] = state ? state.transitGwName : undefined;
            resourceInputs["transitPrependAsPaths"] = state ? state.transitPrependAsPaths : undefined;
        } else {
            const args = argsOrState as AviatrixSpokeTransitAttachmentArgs | undefined;
            if ((!args || args.spokeGwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spokeGwName'");
            }
            if ((!args || args.transitGwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGwName'");
            }
            resourceInputs["enableMaxPerformance"] = args ? args.enableMaxPerformance : undefined;
            resourceInputs["routeTables"] = args ? args.routeTables : undefined;
            resourceInputs["spokeGwName"] = args ? args.spokeGwName : undefined;
            resourceInputs["spokePrependAsPaths"] = args ? args.spokePrependAsPaths : undefined;
            resourceInputs["transitGwName"] = args ? args.transitGwName : undefined;
            resourceInputs["transitPrependAsPaths"] = args ? args.transitPrependAsPaths : undefined;
            resourceInputs["spokeBgpEnabled"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixSpokeTransitAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixSpokeTransitAttachment resources.
 */
export interface AviatrixSpokeTransitAttachmentState {
    /**
     * Indicates whether the maximum amount of HPE tunnels will be created. Only valid when transit and spoke gateways are each launched in Insane Mode and in the same cloud type. Default value: true. Available as of provider version R2.22.2+.
     */
    enableMaxPerformance?: pulumi.Input<boolean>;
    /**
     * Learned routes will be propagated to these route tables. Example: ["rtb-212ff547","rtb-04539787"].
     */
    routeTables?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether the spoke gateway is BGP enabled or not.
     */
    spokeBgpEnabled?: pulumi.Input<boolean>;
    /**
     * Name of the spoke gateway to attach to transit network.
     */
    spokeGwName?: pulumi.Input<string>;
    /**
     * Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on spoke_gateway_name. Available as of provider version R2.23+.
     */
    spokePrependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the transit gateway to attach the spoke gateway to.
     */
    transitGwName?: pulumi.Input<string>;
    /**
     * Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on transit_gateway_name. Available as of provider version R2.23+.
     */
    transitPrependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AviatrixSpokeTransitAttachment resource.
 */
export interface AviatrixSpokeTransitAttachmentArgs {
    /**
     * Indicates whether the maximum amount of HPE tunnels will be created. Only valid when transit and spoke gateways are each launched in Insane Mode and in the same cloud type. Default value: true. Available as of provider version R2.22.2+.
     */
    enableMaxPerformance?: pulumi.Input<boolean>;
    /**
     * Learned routes will be propagated to these route tables. Example: ["rtb-212ff547","rtb-04539787"].
     */
    routeTables?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the spoke gateway to attach to transit network.
     */
    spokeGwName: pulumi.Input<string>;
    /**
     * Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on spoke_gateway_name. Available as of provider version R2.23+.
     */
    spokePrependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the transit gateway to attach the spoke gateway to.
     */
    transitGwName: pulumi.Input<string>;
    /**
     * Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on transit_gateway_name. Available as of provider version R2.23+.
     */
    transitPrependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
}
