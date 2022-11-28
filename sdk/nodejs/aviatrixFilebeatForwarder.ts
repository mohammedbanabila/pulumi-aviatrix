// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_filebeat_forwarder** resource allows the enabling and disabling of filebeat forwarder.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@astipkovits/aviatrix";
 * import * as fs from "fs";
 *
 * // Enable filebeat forwarder
 * const testFilebeatForwarder = new aviatrix.AviatrixFilebeatForwarder("testFilebeatForwarder", {
 *     server: "1.2.3.4",
 *     port: 10,
 *     trustedCaFile: fs.readFileSync("/path/to/ca.pem"),
 *     configFile: fs.readFileSync("/path/to/config.txt"),
 *     excludedGateways: [
 *         "a",
 *         "b",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * **filebeat_forwarder** can be imported using "filebeat_forwarder", e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixFilebeatForwarder:AviatrixFilebeatForwarder test filebeat_forwarder
 * ```
 */
export class AviatrixFilebeatForwarder extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixFilebeatForwarder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixFilebeatForwarderState, opts?: pulumi.CustomResourceOptions): AviatrixFilebeatForwarder {
        return new AviatrixFilebeatForwarder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixFilebeatForwarder:AviatrixFilebeatForwarder';

    /**
     * Returns true if the given object is an instance of AviatrixFilebeatForwarder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixFilebeatForwarder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixFilebeatForwarder.__pulumiType;
    }

    /**
     * The config file. Use the `file` function to read from a file.
     */
    public readonly configFile!: pulumi.Output<string | undefined>;
    /**
     * List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
     */
    public readonly excludedGateways!: pulumi.Output<string[] | undefined>;
    /**
     * Port number.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Server IP.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * The status of filebeat forwarder.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The trusted CA file. Use the `file` function to read from a file.
     */
    public readonly trustedCaFile!: pulumi.Output<string | undefined>;

    /**
     * Create a AviatrixFilebeatForwarder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixFilebeatForwarderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixFilebeatForwarderArgs | AviatrixFilebeatForwarderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixFilebeatForwarderState | undefined;
            resourceInputs["configFile"] = state ? state.configFile : undefined;
            resourceInputs["excludedGateways"] = state ? state.excludedGateways : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trustedCaFile"] = state ? state.trustedCaFile : undefined;
        } else {
            const args = argsOrState as AviatrixFilebeatForwarderArgs | undefined;
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.server === undefined) && !opts.urn) {
                throw new Error("Missing required property 'server'");
            }
            resourceInputs["configFile"] = args ? args.configFile : undefined;
            resourceInputs["excludedGateways"] = args ? args.excludedGateways : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["trustedCaFile"] = args ? args.trustedCaFile : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixFilebeatForwarder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixFilebeatForwarder resources.
 */
export interface AviatrixFilebeatForwarderState {
    /**
     * The config file. Use the `file` function to read from a file.
     */
    configFile?: pulumi.Input<string>;
    /**
     * List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
     */
    excludedGateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Port number.
     */
    port?: pulumi.Input<number>;
    /**
     * Server IP.
     */
    server?: pulumi.Input<string>;
    /**
     * The status of filebeat forwarder.
     */
    status?: pulumi.Input<string>;
    /**
     * The trusted CA file. Use the `file` function to read from a file.
     */
    trustedCaFile?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixFilebeatForwarder resource.
 */
export interface AviatrixFilebeatForwarderArgs {
    /**
     * The config file. Use the `file` function to read from a file.
     */
    configFile?: pulumi.Input<string>;
    /**
     * List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
     */
    excludedGateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Port number.
     */
    port: pulumi.Input<number>;
    /**
     * Server IP.
     */
    server: pulumi.Input<string>;
    /**
     * The trusted CA file. Use the `file` function to read from a file.
     */
    trustedCaFile?: pulumi.Input<string>;
}
