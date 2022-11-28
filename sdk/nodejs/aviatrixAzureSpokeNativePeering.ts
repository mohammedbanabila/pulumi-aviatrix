// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_azure_spoke_native_peering** resource allows the creation and management of Aviatrix-created Azure Spoke VNet attachments via Native Peering.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix Azure spoke native peering
 * const test = new aviatrix.AviatrixAzureSpokeNativePeering("test", {
 *     spokeAccountName: "devops-azure",
 *     spokeRegion: "West US",
 *     spokeVpcId: "Foo_VNet:Bar_RG:GUID",
 *     transitGatewayName: "transit-gw-azure",
 * });
 * ```
 *
 * ## Import
 *
 * **azure_spoke_native_peering** can be imported using the `transit_gateway_name`, `spoke_account_name` and `spoke_vpc_id`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering test transit_gateway_name~spoke_account_name~spoke_vpc_id
 * ```
 */
export class AviatrixAzureSpokeNativePeering extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAzureSpokeNativePeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAzureSpokeNativePeeringState, opts?: pulumi.CustomResourceOptions): AviatrixAzureSpokeNativePeering {
        return new AviatrixAzureSpokeNativePeering(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering';

    /**
     * Returns true if the given object is an instance of AviatrixAzureSpokeNativePeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAzureSpokeNativePeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAzureSpokeNativePeering.__pulumiType;
    }

    /**
     * An Aviatrix account that corresponds to a subscription in Azure.
     */
    public readonly spokeAccountName!: pulumi.Output<string>;
    /**
     * Spoke VNet region. Example: "West US".
     */
    public readonly spokeRegion!: pulumi.Output<string>;
    /**
     * Combination of the Spoke's VNet name, resource group and GUID. Example: "Foo_VNet:Bar_RG:GUID".
     */
    public readonly spokeVpcId!: pulumi.Output<string>;
    /**
     * Name of an Transit FireNet-enabled Azure transit gateway.
     */
    public readonly transitGatewayName!: pulumi.Output<string>;

    /**
     * Create a AviatrixAzureSpokeNativePeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAzureSpokeNativePeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAzureSpokeNativePeeringArgs | AviatrixAzureSpokeNativePeeringState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAzureSpokeNativePeeringState | undefined;
            resourceInputs["spokeAccountName"] = state ? state.spokeAccountName : undefined;
            resourceInputs["spokeRegion"] = state ? state.spokeRegion : undefined;
            resourceInputs["spokeVpcId"] = state ? state.spokeVpcId : undefined;
            resourceInputs["transitGatewayName"] = state ? state.transitGatewayName : undefined;
        } else {
            const args = argsOrState as AviatrixAzureSpokeNativePeeringArgs | undefined;
            if ((!args || args.spokeAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spokeAccountName'");
            }
            if ((!args || args.spokeRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spokeRegion'");
            }
            if ((!args || args.spokeVpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spokeVpcId'");
            }
            if ((!args || args.transitGatewayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayName'");
            }
            resourceInputs["spokeAccountName"] = args ? args.spokeAccountName : undefined;
            resourceInputs["spokeRegion"] = args ? args.spokeRegion : undefined;
            resourceInputs["spokeVpcId"] = args ? args.spokeVpcId : undefined;
            resourceInputs["transitGatewayName"] = args ? args.transitGatewayName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAzureSpokeNativePeering.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAzureSpokeNativePeering resources.
 */
export interface AviatrixAzureSpokeNativePeeringState {
    /**
     * An Aviatrix account that corresponds to a subscription in Azure.
     */
    spokeAccountName?: pulumi.Input<string>;
    /**
     * Spoke VNet region. Example: "West US".
     */
    spokeRegion?: pulumi.Input<string>;
    /**
     * Combination of the Spoke's VNet name, resource group and GUID. Example: "Foo_VNet:Bar_RG:GUID".
     */
    spokeVpcId?: pulumi.Input<string>;
    /**
     * Name of an Transit FireNet-enabled Azure transit gateway.
     */
    transitGatewayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAzureSpokeNativePeering resource.
 */
export interface AviatrixAzureSpokeNativePeeringArgs {
    /**
     * An Aviatrix account that corresponds to a subscription in Azure.
     */
    spokeAccountName: pulumi.Input<string>;
    /**
     * Spoke VNet region. Example: "West US".
     */
    spokeRegion: pulumi.Input<string>;
    /**
     * Combination of the Spoke's VNet name, resource group and GUID. Example: "Foo_VNet:Bar_RG:GUID".
     */
    spokeVpcId: pulumi.Input<string>;
    /**
     * Name of an Transit FireNet-enabled Azure transit gateway.
     */
    transitGatewayName: pulumi.Input<string>;
}
