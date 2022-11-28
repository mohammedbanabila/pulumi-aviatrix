// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@astipkovits/aviatrix";
 *
 * // Create an Aviatrix Stateful Firewall Policy
 * const testFirewallPolicy = new aviatrix.AviatrixFirewallPolicy("testFirewallPolicy", {
 *     gwName: aviatrix_firewall.test_firewall.gw_name,
 *     srcIp: "10.15.0.224/32",
 *     dstIp: "10.12.0.172/32",
 *     protocol: "tcp",
 *     port: "0:65535",
 *     action: "allow",
 *     logEnabled: true,
 *     description: "Test policy.",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@astipkovits/aviatrix";
 *
 * // Create an Aviatrix Stateful Firewall Policy and insert it to a specific position
 * const testFirewallPolicy = new aviatrix.AviatrixFirewallPolicy("testFirewallPolicy", {
 *     gwName: aviatrix_firewall.test_firewall.gw_name,
 *     srcIp: "10.15.0.225/32",
 *     dstIp: "10.12.0.173/32",
 *     protocol: "tcp",
 *     port: "0:65535",
 *     action: "allow",
 *     logEnabled: true,
 *     description: "Test policy.",
 *     position: 2,
 * });
 * ```
 *
 * ## Import
 *
 * **firewall_policy** can be imported using the `gw_name`, `src_ip`, `dst_ip`, `protocol`, `port` and `action` separated by `~`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixFirewallPolicy:AviatrixFirewallPolicy test "gw_name~src_ip~dst_ip~protocol~port~action"
 * ```
 */
export class AviatrixFirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixFirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixFirewallPolicyState, opts?: pulumi.CustomResourceOptions): AviatrixFirewallPolicy {
        return new AviatrixFirewallPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixFirewallPolicy:AviatrixFirewallPolicy';

    /**
     * Returns true if the given object is an instance of AviatrixFirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixFirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixFirewallPolicy.__pulumiType;
    }

    /**
     * Valid values: "allow", "deny" and "force-drop" (in stateful firewall rule to allow immediate packet dropping on established sessions).
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Description of the policy. Example: "This is policy no.1".
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
     */
    public readonly dstIp!: pulumi.Output<string>;
    /**
     * Gateway name to attach firewall policy to.
     */
    public readonly gwName!: pulumi.Output<string>;
    /**
     * Valid values: true, false. Default value: false.
     */
    public readonly logEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A single port or a range of port numbers. Example: "25", "25:1024".
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Position in the policy list, where the firewall policy will be inserted to. Valid values: any positive integer. Example: 2. If it is larger than the size of policy list, the policy will be inserted to the end.
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * : "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp".
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
     */
    public readonly srcIp!: pulumi.Output<string>;

    /**
     * Create a AviatrixFirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixFirewallPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixFirewallPolicyArgs | AviatrixFirewallPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixFirewallPolicyState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dstIp"] = state ? state.dstIp : undefined;
            resourceInputs["gwName"] = state ? state.gwName : undefined;
            resourceInputs["logEnabled"] = state ? state.logEnabled : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["srcIp"] = state ? state.srcIp : undefined;
        } else {
            const args = argsOrState as AviatrixFirewallPolicyArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.dstIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstIp'");
            }
            if ((!args || args.gwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gwName'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.srcIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcIp'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dstIp"] = args ? args.dstIp : undefined;
            resourceInputs["gwName"] = args ? args.gwName : undefined;
            resourceInputs["logEnabled"] = args ? args.logEnabled : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["srcIp"] = args ? args.srcIp : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixFirewallPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixFirewallPolicy resources.
 */
export interface AviatrixFirewallPolicyState {
    /**
     * Valid values: "allow", "deny" and "force-drop" (in stateful firewall rule to allow immediate packet dropping on established sessions).
     */
    action?: pulumi.Input<string>;
    /**
     * Description of the policy. Example: "This is policy no.1".
     */
    description?: pulumi.Input<string>;
    /**
     * CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
     */
    dstIp?: pulumi.Input<string>;
    /**
     * Gateway name to attach firewall policy to.
     */
    gwName?: pulumi.Input<string>;
    /**
     * Valid values: true, false. Default value: false.
     */
    logEnabled?: pulumi.Input<boolean>;
    /**
     * A single port or a range of port numbers. Example: "25", "25:1024".
     */
    port?: pulumi.Input<string>;
    /**
     * Position in the policy list, where the firewall policy will be inserted to. Valid values: any positive integer. Example: 2. If it is larger than the size of policy list, the policy will be inserted to the end.
     */
    position?: pulumi.Input<number>;
    /**
     * : "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp".
     */
    protocol?: pulumi.Input<string>;
    /**
     * CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
     */
    srcIp?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixFirewallPolicy resource.
 */
export interface AviatrixFirewallPolicyArgs {
    /**
     * Valid values: "allow", "deny" and "force-drop" (in stateful firewall rule to allow immediate packet dropping on established sessions).
     */
    action: pulumi.Input<string>;
    /**
     * Description of the policy. Example: "This is policy no.1".
     */
    description?: pulumi.Input<string>;
    /**
     * CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
     */
    dstIp: pulumi.Input<string>;
    /**
     * Gateway name to attach firewall policy to.
     */
    gwName: pulumi.Input<string>;
    /**
     * Valid values: true, false. Default value: false.
     */
    logEnabled?: pulumi.Input<boolean>;
    /**
     * A single port or a range of port numbers. Example: "25", "25:1024".
     */
    port: pulumi.Input<string>;
    /**
     * Position in the policy list, where the firewall policy will be inserted to. Valid values: any positive integer. Example: 2. If it is larger than the size of policy list, the policy will be inserted to the end.
     */
    position?: pulumi.Input<number>;
    /**
     * : "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp".
     */
    protocol?: pulumi.Input<string>;
    /**
     * CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
     */
    srcIp: pulumi.Input<string>;
}
