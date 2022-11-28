// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix Firewall
 * const statefulFirewall1 = new aviatrix.AviatrixFirewall("stateful_firewall_1", {
 *     baseLogEnabled: true,
 *     basePolicy: "allow-all",
 *     gwName: "gateway-1",
 *     manageFirewallPolicies: false,
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@astipkovits/aviatrix";
 *
 * // Create an Aviatrix Firewall with in-line rules
 * const statefulFirewall1 = new aviatrix.AviatrixFirewall("statefulFirewall1", {
 *     gwName: "gateway-1",
 *     basePolicy: "allow-all",
 *     baseLogEnabled: true,
 *     policies: [
 *         {
 *             protocol: "all",
 *             srcIp: "10.17.0.224/32",
 *             logEnabled: true,
 *             dstIp: "10.12.0.172/32",
 *             action: "force-drop",
 *             port: "0:65535",
 *             description: "first_policy",
 *         },
 *         {
 *             protocol: "tcp",
 *             srcIp: "10.16.0.224/32",
 *             logEnabled: false,
 *             dstIp: "10.12.1.172/32",
 *             action: "force-drop",
 *             port: "325",
 *             description: "second_policy",
 *         },
 *         {
 *             protocol: "udp",
 *             srcIp: "10.14.0.225/32",
 *             logEnabled: false,
 *             dstIp: "10.13.1.173/32",
 *             action: "deny",
 *             port: "325",
 *             description: "third_policy",
 *         },
 *         {
 *             protocol: "tcp",
 *             srcIp: aviatrix_firewall_tag.test.firewall_tag,
 *             logEnabled: false,
 *             dstIp: "10.13.1.173/32",
 *             action: "deny",
 *             port: "325",
 *             description: "fourth_policy",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * **firewall** can be imported using the `gw_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixFirewall:AviatrixFirewall test gw_name
 * ```
 */
export class AviatrixFirewall extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixFirewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixFirewallState, opts?: pulumi.CustomResourceOptions): AviatrixFirewall {
        return new AviatrixFirewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixFirewall:AviatrixFirewall';

    /**
     * Returns true if the given object is an instance of AviatrixFirewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixFirewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixFirewall.__pulumiType;
    }

    /**
     * Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
     */
    public readonly baseLogEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
     */
    public readonly basePolicy!: pulumi.Output<string | undefined>;
    /**
     * Gateway name to attach firewall policy to.
     */
    public readonly gwName!: pulumi.Output<string>;
    /**
     * Enable to manage firewall policies via in-line rules. If false, policies must be managed using `aviatrix.AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
     */
    public readonly manageFirewallPolicies!: pulumi.Output<boolean | undefined>;
    /**
     * New access policy for the gateway. Seven fields are required for each policy item: `srcIp`, `dstIp`, `protocol`, `port`, `action`, `logEnabled` and `description`. No duplicate rules (with same `srcIp`, `dstIp`, `protocol` and `port`) are allowed.
     */
    public readonly policies!: pulumi.Output<outputs.AviatrixFirewallPolicy[] | undefined>;

    /**
     * Create a AviatrixFirewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixFirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixFirewallArgs | AviatrixFirewallState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixFirewallState | undefined;
            resourceInputs["baseLogEnabled"] = state ? state.baseLogEnabled : undefined;
            resourceInputs["basePolicy"] = state ? state.basePolicy : undefined;
            resourceInputs["gwName"] = state ? state.gwName : undefined;
            resourceInputs["manageFirewallPolicies"] = state ? state.manageFirewallPolicies : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as AviatrixFirewallArgs | undefined;
            if ((!args || args.gwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gwName'");
            }
            resourceInputs["baseLogEnabled"] = args ? args.baseLogEnabled : undefined;
            resourceInputs["basePolicy"] = args ? args.basePolicy : undefined;
            resourceInputs["gwName"] = args ? args.gwName : undefined;
            resourceInputs["manageFirewallPolicies"] = args ? args.manageFirewallPolicies : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixFirewall.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixFirewall resources.
 */
export interface AviatrixFirewallState {
    /**
     * Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
     */
    baseLogEnabled?: pulumi.Input<boolean>;
    /**
     * New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
     */
    basePolicy?: pulumi.Input<string>;
    /**
     * Gateway name to attach firewall policy to.
     */
    gwName?: pulumi.Input<string>;
    /**
     * Enable to manage firewall policies via in-line rules. If false, policies must be managed using `aviatrix.AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
     */
    manageFirewallPolicies?: pulumi.Input<boolean>;
    /**
     * New access policy for the gateway. Seven fields are required for each policy item: `srcIp`, `dstIp`, `protocol`, `port`, `action`, `logEnabled` and `description`. No duplicate rules (with same `srcIp`, `dstIp`, `protocol` and `port`) are allowed.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.AviatrixFirewallPolicy>[]>;
}

/**
 * The set of arguments for constructing a AviatrixFirewall resource.
 */
export interface AviatrixFirewallArgs {
    /**
     * Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
     */
    baseLogEnabled?: pulumi.Input<boolean>;
    /**
     * New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
     */
    basePolicy?: pulumi.Input<string>;
    /**
     * Gateway name to attach firewall policy to.
     */
    gwName: pulumi.Input<string>;
    /**
     * Enable to manage firewall policies via in-line rules. If false, policies must be managed using `aviatrix.AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
     */
    manageFirewallPolicies?: pulumi.Input<boolean>;
    /**
     * New access policy for the gateway. Seven fields are required for each policy item: `srcIp`, `dstIp`, `protocol`, `port`, `action`, `logEnabled` and `description`. No duplicate rules (with same `srcIp`, `dstIp`, `protocol` and `port`) are allowed.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.AviatrixFirewallPolicy>[]>;
}
