// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_device_interface_config** resource allows the configuration of the WAN primary interface and IP for a device, for use in CloudN.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Configure the primary WAN interface and IP for a device.
 * const testDeviceInterfaceConfig = new aviatrix.AviatrixDeviceInterfaceConfig("test_device_interface_config", {
 *     deviceName: "test-device",
 *     wanPrimaryInterface: "eth0",
 *     wanPrimaryInterfacePublicIp: "181.12.43.21",
 * });
 * ```
 *
 * ## Import
 *
 * **device_interface_config** can be imported using the `device_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixDeviceInterfaceConfig:AviatrixDeviceInterfaceConfig test device_name
 * ```
 */
export class AviatrixDeviceInterfaceConfig extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixDeviceInterfaceConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixDeviceInterfaceConfigState, opts?: pulumi.CustomResourceOptions): AviatrixDeviceInterfaceConfig {
        return new AviatrixDeviceInterfaceConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixDeviceInterfaceConfig:AviatrixDeviceInterfaceConfig';

    /**
     * Returns true if the given object is an instance of AviatrixDeviceInterfaceConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixDeviceInterfaceConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixDeviceInterfaceConfig.__pulumiType;
    }

    /**
     * Name of the device.
     */
    public readonly deviceName!: pulumi.Output<string>;
    /**
     * Name of the WAN primary interface.
     */
    public readonly wanPrimaryInterface!: pulumi.Output<string>;
    /**
     * The WAN Primary interface public IP.
     */
    public readonly wanPrimaryInterfacePublicIp!: pulumi.Output<string>;

    /**
     * Create a AviatrixDeviceInterfaceConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixDeviceInterfaceConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixDeviceInterfaceConfigArgs | AviatrixDeviceInterfaceConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixDeviceInterfaceConfigState | undefined;
            resourceInputs["deviceName"] = state ? state.deviceName : undefined;
            resourceInputs["wanPrimaryInterface"] = state ? state.wanPrimaryInterface : undefined;
            resourceInputs["wanPrimaryInterfacePublicIp"] = state ? state.wanPrimaryInterfacePublicIp : undefined;
        } else {
            const args = argsOrState as AviatrixDeviceInterfaceConfigArgs | undefined;
            if ((!args || args.deviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceName'");
            }
            if ((!args || args.wanPrimaryInterface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'wanPrimaryInterface'");
            }
            if ((!args || args.wanPrimaryInterfacePublicIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'wanPrimaryInterfacePublicIp'");
            }
            resourceInputs["deviceName"] = args ? args.deviceName : undefined;
            resourceInputs["wanPrimaryInterface"] = args ? args.wanPrimaryInterface : undefined;
            resourceInputs["wanPrimaryInterfacePublicIp"] = args ? args.wanPrimaryInterfacePublicIp : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixDeviceInterfaceConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixDeviceInterfaceConfig resources.
 */
export interface AviatrixDeviceInterfaceConfigState {
    /**
     * Name of the device.
     */
    deviceName?: pulumi.Input<string>;
    /**
     * Name of the WAN primary interface.
     */
    wanPrimaryInterface?: pulumi.Input<string>;
    /**
     * The WAN Primary interface public IP.
     */
    wanPrimaryInterfacePublicIp?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixDeviceInterfaceConfig resource.
 */
export interface AviatrixDeviceInterfaceConfigArgs {
    /**
     * Name of the device.
     */
    deviceName: pulumi.Input<string>;
    /**
     * Name of the WAN primary interface.
     */
    wanPrimaryInterface: pulumi.Input<string>;
    /**
     * The WAN Primary interface public IP.
     */
    wanPrimaryInterfacePublicIp: pulumi.Input<string>;
}
