// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * !> **WARNING** **aviatrix_app_domain** is part of the Micro-segmentation private preview feature for R2.22.0. If you wish to enable a private preview mode feature, please contact your sales representative or Aviatrix Support.
 * The **aviatrix_app_domain** resource handles the creation and management of App Domains. Available as of Provider R2.22.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix App Domain
 * const testAppDomainIp = new aviatrix.AviatrixAppDomain("test_app_domain_ip", {
 *     selector: {
 *         matchExpressions: [
 *             {
 *                 accountName: "devops",
 *                 region: "us-west-2",
 *                 tags: {
 *                     k3: "v3",
 *                 },
 *                 type: "vm",
 *             },
 *             {
 *                 cidr: "10.0.0.0/16",
 *             },
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * **aviatrix_app_domain** can be imported using the `uuid`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixAppDomain:AviatrixAppDomain test 41984f8b-5a37-4272-89b3-57c79e9ff77c
 * ```
 */
export class AviatrixAppDomain extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAppDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAppDomainState, opts?: pulumi.CustomResourceOptions): AviatrixAppDomain {
        return new AviatrixAppDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAppDomain:AviatrixAppDomain';

    /**
     * Returns true if the given object is an instance of AviatrixAppDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAppDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAppDomain.__pulumiType;
    }

    /**
     * Name of the App Domain.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Block containing match expressions to filter the App Domain.
     */
    public readonly selector!: pulumi.Output<outputs.AviatrixAppDomainSelector>;
    /**
     * UUID of the App Domain.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a AviatrixAppDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAppDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAppDomainArgs | AviatrixAppDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAppDomainState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["selector"] = state ? state.selector : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as AviatrixAppDomainArgs | undefined;
            if ((!args || args.selector === undefined) && !opts.urn) {
                throw new Error("Missing required property 'selector'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["selector"] = args ? args.selector : undefined;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAppDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAppDomain resources.
 */
export interface AviatrixAppDomainState {
    /**
     * Name of the App Domain.
     */
    name?: pulumi.Input<string>;
    /**
     * Block containing match expressions to filter the App Domain.
     */
    selector?: pulumi.Input<inputs.AviatrixAppDomainSelector>;
    /**
     * UUID of the App Domain.
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAppDomain resource.
 */
export interface AviatrixAppDomainArgs {
    /**
     * Name of the App Domain.
     */
    name?: pulumi.Input<string>;
    /**
     * Block containing match expressions to filter the App Domain.
     */
    selector: pulumi.Input<inputs.AviatrixAppDomainSelector>;
}
