// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * **transit_external_device_conn** can be imported using the `connection_name` and `vpc_id`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixTransitExternalDeviceConn:AviatrixTransitExternalDeviceConn test connection_name~vpc_id
 * ```
 */
export class AviatrixTransitExternalDeviceConn extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixTransitExternalDeviceConn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixTransitExternalDeviceConnState, opts?: pulumi.CustomResourceOptions): AviatrixTransitExternalDeviceConn {
        return new AviatrixTransitExternalDeviceConn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixTransitExternalDeviceConn:AviatrixTransitExternalDeviceConn';

    /**
     * Returns true if the given object is an instance of AviatrixTransitExternalDeviceConn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixTransitExternalDeviceConn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixTransitExternalDeviceConn.__pulumiType;
    }

    /**
     * Set of approved CIDRs. Requires `enableLearnedCidrsApproval` to be true. Type: Set(String).
     */
    public readonly approvedCidrs!: pulumi.Output<string[]>;
    /**
     * Backup BGP MD5 Authentication Key. Valid with HA enabled for connection. Example: 'avx03,avx04'. For BGP LAN ActiveMesh mode disabled, example: 'avx03'.
     */
    public readonly backupBgpMd5Key!: pulumi.Output<string | undefined>;
    /**
     * Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required if HA enabled for 'bgp' connection.
     */
    public readonly backupBgpRemoteAsNum!: pulumi.Output<string | undefined>;
    /**
     * Backup direct connect for backup external device.
     */
    public readonly backupDirectConnect!: pulumi.Output<boolean | undefined>;
    /**
     * Backup Local LAN IP. Required for GCP HA BGP over LAN connection.
     */
    public readonly backupLocalLanIp!: pulumi.Output<string>;
    /**
     * Source CIDR for the tunnel from the backup Aviatrix transit gateway.
     */
    public readonly backupLocalTunnelCidr!: pulumi.Output<string>;
    /**
     * Backup Pre-Shared Key.
     */
    public readonly backupPreSharedKey!: pulumi.Output<string | undefined>;
    /**
     * Backup remote gateway IP.
     */
    public readonly backupRemoteGatewayIp!: pulumi.Output<string | undefined>;
    /**
     * Backup Remote LAN IP. Required for HA BGP over LAN connection.
     */
    public readonly backupRemoteLanIp!: pulumi.Output<string | undefined>;
    /**
     * Destination CIDR for the tunnel to the backup external device.
     */
    public readonly backupRemoteTunnelCidr!: pulumi.Output<string>;
    /**
     * BGP local ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.
     */
    public readonly bgpLocalAsNum!: pulumi.Output<string | undefined>;
    /**
     * BGP MD5 Authentication Key. Example: 'avx01,avx02'. For BGP LAN ActiveMesh mode disabled, example: 'avx01'.
     */
    public readonly bgpMd5Key!: pulumi.Output<string | undefined>;
    /**
     * BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.
     */
    public readonly bgpRemoteAsNum!: pulumi.Output<string | undefined>;
    /**
     * Transit external device connection name.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'.
     */
    public readonly connectionType!: pulumi.Output<string | undefined>;
    /**
     * Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption. Valid values: true, false. **NOTE: Please see notes here for more information.**
     */
    public readonly customAlgorithms!: pulumi.Output<boolean | undefined>;
    /**
     * Set true for private network infrastructure.
     */
    public readonly directConnect!: pulumi.Output<boolean | undefined>;
    /**
     * Switch to enable BGP LAN ActiveMesh mode. Only valid for GCP with Remote Gateway HA enabled. Default: false. Available as of provider version R2.21+.
     */
    public readonly enableBgpLanActivemesh!: pulumi.Output<boolean | undefined>;
    /**
     * Switch to allow this connection to communicate with a Network Domain via Connection Policy.
     */
    public readonly enableEdgeSegmentation!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
     */
    public readonly enableEventTriggeredHa!: pulumi.Output<boolean | undefined>;
    /**
     * Set as true to enable IKEv2 protocol.
     */
    public readonly enableIkev2!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Jumbo Frame for the transit external device connection. Only valid with 'GRE' tunnels under 'bgp' connection. Requires transit to be jumbo frame and insane mode enabled. Valid values: true, false. Default value: false. Available as of provider version R2.22.2+.
     */
    public readonly enableJumboFrame!: pulumi.Output<boolean | undefined>;
    /**
     * Enable learned CIDRs approval for the connection. Only valid with `connectionType` = 'bgp'. Requires the transit_gateway's `learnedCidrsApprovalMode` attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    public readonly enableLearnedCidrsApproval!: pulumi.Output<boolean | undefined>;
    /**
     * Aviatrix transit gateway name.
     */
    public readonly gwName!: pulumi.Output<string>;
    /**
     * Set as true if there are two external devices.
     * * `backupRemoteGatewayIp ` - (Optional) Backup remote gateway IP. Required if HA enabled.
     */
    public readonly haEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Local LAN IP. Required for GCP BGP over LAN connection.
     */
    public readonly localLanIp!: pulumi.Output<string>;
    /**
     * Source CIDR for the tunnel from the Aviatrix transit gateway.
     */
    public readonly localTunnelCidr!: pulumi.Output<string>;
    /**
     * Configure manual BGP advertised CIDRs for this connection. Only valid with `connectionType`= 'bgp'. Available as of provider version R2.18+.
     */
    public readonly manualBgpAdvertisedCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'. Default value: 'SHA-256'.
     */
    public readonly phase1Authentication!: pulumi.Output<string | undefined>;
    /**
     * Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.
     */
    public readonly phase1DhGroups!: pulumi.Output<string | undefined>;
    /**
     * Phase one Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", and "AES-256-GCM-128". Default value: "AES-256-CBC".
     */
    public readonly phase1Encryption!: pulumi.Output<string | undefined>;
    /**
     * Phase 1 remote identifier of the IPsec tunnel. This can be configured to be either the public IP address or the private IP address of the peer terminating the IPsec tunnel. Example: ["1.2.3.4"] when HA is disabled, ["1.2.3.4", "5.6.7.8"] when HA is enabled. Available as of provider version R2.19+.
     */
    public readonly phase1RemoteIdentifiers!: pulumi.Output<string[] | undefined>;
    /**
     * Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'. Default value: 'HMAC-SHA-256'.
     */
    public readonly phase2Authentication!: pulumi.Output<string | undefined>;
    /**
     * Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.
     */
    public readonly phase2DhGroups!: pulumi.Output<string | undefined>;
    /**
     * Phase two Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", "AES-256-GCM-128" and "NULL-ENCR". Default value: "AES-256-CBC".
     */
    public readonly phase2Encryption!: pulumi.Output<string | undefined>;
    /**
     * Pre-Shared Key.
     */
    public readonly preSharedKey!: pulumi.Output<string | undefined>;
    /**
     * Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
     */
    public readonly prependAsPaths!: pulumi.Output<string[] | undefined>;
    /**
     * Remote gateway IP. Required when `tunnelProtocol` != 'LAN'.
     */
    public readonly remoteGatewayIp!: pulumi.Output<string | undefined>;
    /**
     * Remote LAN IP. Required for BGP over LAN connection.
     */
    public readonly remoteLanIp!: pulumi.Output<string | undefined>;
    /**
     * Remote CIDRs joined as a string with ','. Required for a 'static' type connection.
     */
    public readonly remoteSubnet!: pulumi.Output<string | undefined>;
    /**
     * Destination CIDR for the tunnel to the external device.
     */
    public readonly remoteTunnelCidr!: pulumi.Output<string>;
    /**
     * Name of the remote VPC for a LAN BGP connection with an Azure Transit Gateway. Required when `connectionType` = 'bgp' and `tunnelProtocol` = 'LAN' with an Azure transit gateway. Must be in the format "<vnet-name>:<resource-group-name>:<subscription-id>". Available as of provider version R2.18+.
     */
    public readonly remoteVpcName!: pulumi.Output<string | undefined>;
    /**
     * Switch to HA Standby Transit Gateway connection. Only valid with Transit Gateway that has [Active-Standby Mode](https://docs.aviatrix.com/HowTos/transit_advanced.html#active-standby) enabled and for non-HA external device. Valid values: true, false. Default: false. Available in provider version R2.17.1+.
     */
    public readonly switchToHaStandbyGateway!: pulumi.Output<boolean | undefined>;
    /**
     * Tunnel protocol, only valid with `connectionType` = 'bgp'. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive. Available as of provider version R2.18+.
     */
    public readonly tunnelProtocol!: pulumi.Output<string | undefined>;
    /**
     * VPC ID of the Aviatrix transit gateway. For GCP BGP over LAN connection, it is in the format of "vpc_name~-~project_name".
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a AviatrixTransitExternalDeviceConn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixTransitExternalDeviceConnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixTransitExternalDeviceConnArgs | AviatrixTransitExternalDeviceConnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixTransitExternalDeviceConnState | undefined;
            resourceInputs["approvedCidrs"] = state ? state.approvedCidrs : undefined;
            resourceInputs["backupBgpMd5Key"] = state ? state.backupBgpMd5Key : undefined;
            resourceInputs["backupBgpRemoteAsNum"] = state ? state.backupBgpRemoteAsNum : undefined;
            resourceInputs["backupDirectConnect"] = state ? state.backupDirectConnect : undefined;
            resourceInputs["backupLocalLanIp"] = state ? state.backupLocalLanIp : undefined;
            resourceInputs["backupLocalTunnelCidr"] = state ? state.backupLocalTunnelCidr : undefined;
            resourceInputs["backupPreSharedKey"] = state ? state.backupPreSharedKey : undefined;
            resourceInputs["backupRemoteGatewayIp"] = state ? state.backupRemoteGatewayIp : undefined;
            resourceInputs["backupRemoteLanIp"] = state ? state.backupRemoteLanIp : undefined;
            resourceInputs["backupRemoteTunnelCidr"] = state ? state.backupRemoteTunnelCidr : undefined;
            resourceInputs["bgpLocalAsNum"] = state ? state.bgpLocalAsNum : undefined;
            resourceInputs["bgpMd5Key"] = state ? state.bgpMd5Key : undefined;
            resourceInputs["bgpRemoteAsNum"] = state ? state.bgpRemoteAsNum : undefined;
            resourceInputs["connectionName"] = state ? state.connectionName : undefined;
            resourceInputs["connectionType"] = state ? state.connectionType : undefined;
            resourceInputs["customAlgorithms"] = state ? state.customAlgorithms : undefined;
            resourceInputs["directConnect"] = state ? state.directConnect : undefined;
            resourceInputs["enableBgpLanActivemesh"] = state ? state.enableBgpLanActivemesh : undefined;
            resourceInputs["enableEdgeSegmentation"] = state ? state.enableEdgeSegmentation : undefined;
            resourceInputs["enableEventTriggeredHa"] = state ? state.enableEventTriggeredHa : undefined;
            resourceInputs["enableIkev2"] = state ? state.enableIkev2 : undefined;
            resourceInputs["enableJumboFrame"] = state ? state.enableJumboFrame : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = state ? state.enableLearnedCidrsApproval : undefined;
            resourceInputs["gwName"] = state ? state.gwName : undefined;
            resourceInputs["haEnabled"] = state ? state.haEnabled : undefined;
            resourceInputs["localLanIp"] = state ? state.localLanIp : undefined;
            resourceInputs["localTunnelCidr"] = state ? state.localTunnelCidr : undefined;
            resourceInputs["manualBgpAdvertisedCidrs"] = state ? state.manualBgpAdvertisedCidrs : undefined;
            resourceInputs["phase1Authentication"] = state ? state.phase1Authentication : undefined;
            resourceInputs["phase1DhGroups"] = state ? state.phase1DhGroups : undefined;
            resourceInputs["phase1Encryption"] = state ? state.phase1Encryption : undefined;
            resourceInputs["phase1RemoteIdentifiers"] = state ? state.phase1RemoteIdentifiers : undefined;
            resourceInputs["phase2Authentication"] = state ? state.phase2Authentication : undefined;
            resourceInputs["phase2DhGroups"] = state ? state.phase2DhGroups : undefined;
            resourceInputs["phase2Encryption"] = state ? state.phase2Encryption : undefined;
            resourceInputs["preSharedKey"] = state ? state.preSharedKey : undefined;
            resourceInputs["prependAsPaths"] = state ? state.prependAsPaths : undefined;
            resourceInputs["remoteGatewayIp"] = state ? state.remoteGatewayIp : undefined;
            resourceInputs["remoteLanIp"] = state ? state.remoteLanIp : undefined;
            resourceInputs["remoteSubnet"] = state ? state.remoteSubnet : undefined;
            resourceInputs["remoteTunnelCidr"] = state ? state.remoteTunnelCidr : undefined;
            resourceInputs["remoteVpcName"] = state ? state.remoteVpcName : undefined;
            resourceInputs["switchToHaStandbyGateway"] = state ? state.switchToHaStandbyGateway : undefined;
            resourceInputs["tunnelProtocol"] = state ? state.tunnelProtocol : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as AviatrixTransitExternalDeviceConnArgs | undefined;
            if ((!args || args.connectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.gwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gwName'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["approvedCidrs"] = args ? args.approvedCidrs : undefined;
            resourceInputs["backupBgpMd5Key"] = args?.backupBgpMd5Key ? pulumi.secret(args.backupBgpMd5Key) : undefined;
            resourceInputs["backupBgpRemoteAsNum"] = args ? args.backupBgpRemoteAsNum : undefined;
            resourceInputs["backupDirectConnect"] = args ? args.backupDirectConnect : undefined;
            resourceInputs["backupLocalLanIp"] = args ? args.backupLocalLanIp : undefined;
            resourceInputs["backupLocalTunnelCidr"] = args ? args.backupLocalTunnelCidr : undefined;
            resourceInputs["backupPreSharedKey"] = args?.backupPreSharedKey ? pulumi.secret(args.backupPreSharedKey) : undefined;
            resourceInputs["backupRemoteGatewayIp"] = args ? args.backupRemoteGatewayIp : undefined;
            resourceInputs["backupRemoteLanIp"] = args ? args.backupRemoteLanIp : undefined;
            resourceInputs["backupRemoteTunnelCidr"] = args ? args.backupRemoteTunnelCidr : undefined;
            resourceInputs["bgpLocalAsNum"] = args ? args.bgpLocalAsNum : undefined;
            resourceInputs["bgpMd5Key"] = args?.bgpMd5Key ? pulumi.secret(args.bgpMd5Key) : undefined;
            resourceInputs["bgpRemoteAsNum"] = args ? args.bgpRemoteAsNum : undefined;
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["connectionType"] = args ? args.connectionType : undefined;
            resourceInputs["customAlgorithms"] = args ? args.customAlgorithms : undefined;
            resourceInputs["directConnect"] = args ? args.directConnect : undefined;
            resourceInputs["enableBgpLanActivemesh"] = args ? args.enableBgpLanActivemesh : undefined;
            resourceInputs["enableEdgeSegmentation"] = args ? args.enableEdgeSegmentation : undefined;
            resourceInputs["enableEventTriggeredHa"] = args ? args.enableEventTriggeredHa : undefined;
            resourceInputs["enableIkev2"] = args ? args.enableIkev2 : undefined;
            resourceInputs["enableJumboFrame"] = args ? args.enableJumboFrame : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = args ? args.enableLearnedCidrsApproval : undefined;
            resourceInputs["gwName"] = args ? args.gwName : undefined;
            resourceInputs["haEnabled"] = args ? args.haEnabled : undefined;
            resourceInputs["localLanIp"] = args ? args.localLanIp : undefined;
            resourceInputs["localTunnelCidr"] = args ? args.localTunnelCidr : undefined;
            resourceInputs["manualBgpAdvertisedCidrs"] = args ? args.manualBgpAdvertisedCidrs : undefined;
            resourceInputs["phase1Authentication"] = args ? args.phase1Authentication : undefined;
            resourceInputs["phase1DhGroups"] = args ? args.phase1DhGroups : undefined;
            resourceInputs["phase1Encryption"] = args ? args.phase1Encryption : undefined;
            resourceInputs["phase1RemoteIdentifiers"] = args ? args.phase1RemoteIdentifiers : undefined;
            resourceInputs["phase2Authentication"] = args ? args.phase2Authentication : undefined;
            resourceInputs["phase2DhGroups"] = args ? args.phase2DhGroups : undefined;
            resourceInputs["phase2Encryption"] = args ? args.phase2Encryption : undefined;
            resourceInputs["preSharedKey"] = args?.preSharedKey ? pulumi.secret(args.preSharedKey) : undefined;
            resourceInputs["prependAsPaths"] = args ? args.prependAsPaths : undefined;
            resourceInputs["remoteGatewayIp"] = args ? args.remoteGatewayIp : undefined;
            resourceInputs["remoteLanIp"] = args ? args.remoteLanIp : undefined;
            resourceInputs["remoteSubnet"] = args ? args.remoteSubnet : undefined;
            resourceInputs["remoteTunnelCidr"] = args ? args.remoteTunnelCidr : undefined;
            resourceInputs["remoteVpcName"] = args ? args.remoteVpcName : undefined;
            resourceInputs["switchToHaStandbyGateway"] = args ? args.switchToHaStandbyGateway : undefined;
            resourceInputs["tunnelProtocol"] = args ? args.tunnelProtocol : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["backupBgpMd5Key", "backupPreSharedKey", "bgpMd5Key", "preSharedKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AviatrixTransitExternalDeviceConn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixTransitExternalDeviceConn resources.
 */
export interface AviatrixTransitExternalDeviceConnState {
    /**
     * Set of approved CIDRs. Requires `enableLearnedCidrsApproval` to be true. Type: Set(String).
     */
    approvedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup BGP MD5 Authentication Key. Valid with HA enabled for connection. Example: 'avx03,avx04'. For BGP LAN ActiveMesh mode disabled, example: 'avx03'.
     */
    backupBgpMd5Key?: pulumi.Input<string>;
    /**
     * Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required if HA enabled for 'bgp' connection.
     */
    backupBgpRemoteAsNum?: pulumi.Input<string>;
    /**
     * Backup direct connect for backup external device.
     */
    backupDirectConnect?: pulumi.Input<boolean>;
    /**
     * Backup Local LAN IP. Required for GCP HA BGP over LAN connection.
     */
    backupLocalLanIp?: pulumi.Input<string>;
    /**
     * Source CIDR for the tunnel from the backup Aviatrix transit gateway.
     */
    backupLocalTunnelCidr?: pulumi.Input<string>;
    /**
     * Backup Pre-Shared Key.
     */
    backupPreSharedKey?: pulumi.Input<string>;
    /**
     * Backup remote gateway IP.
     */
    backupRemoteGatewayIp?: pulumi.Input<string>;
    /**
     * Backup Remote LAN IP. Required for HA BGP over LAN connection.
     */
    backupRemoteLanIp?: pulumi.Input<string>;
    /**
     * Destination CIDR for the tunnel to the backup external device.
     */
    backupRemoteTunnelCidr?: pulumi.Input<string>;
    /**
     * BGP local ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.
     */
    bgpLocalAsNum?: pulumi.Input<string>;
    /**
     * BGP MD5 Authentication Key. Example: 'avx01,avx02'. For BGP LAN ActiveMesh mode disabled, example: 'avx01'.
     */
    bgpMd5Key?: pulumi.Input<string>;
    /**
     * BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.
     */
    bgpRemoteAsNum?: pulumi.Input<string>;
    /**
     * Transit external device connection name.
     */
    connectionName?: pulumi.Input<string>;
    /**
     * Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'.
     */
    connectionType?: pulumi.Input<string>;
    /**
     * Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption. Valid values: true, false. **NOTE: Please see notes here for more information.**
     */
    customAlgorithms?: pulumi.Input<boolean>;
    /**
     * Set true for private network infrastructure.
     */
    directConnect?: pulumi.Input<boolean>;
    /**
     * Switch to enable BGP LAN ActiveMesh mode. Only valid for GCP with Remote Gateway HA enabled. Default: false. Available as of provider version R2.21+.
     */
    enableBgpLanActivemesh?: pulumi.Input<boolean>;
    /**
     * Switch to allow this connection to communicate with a Network Domain via Connection Policy.
     */
    enableEdgeSegmentation?: pulumi.Input<boolean>;
    /**
     * Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
     */
    enableEventTriggeredHa?: pulumi.Input<boolean>;
    /**
     * Set as true to enable IKEv2 protocol.
     */
    enableIkev2?: pulumi.Input<boolean>;
    /**
     * Enable Jumbo Frame for the transit external device connection. Only valid with 'GRE' tunnels under 'bgp' connection. Requires transit to be jumbo frame and insane mode enabled. Valid values: true, false. Default value: false. Available as of provider version R2.22.2+.
     */
    enableJumboFrame?: pulumi.Input<boolean>;
    /**
     * Enable learned CIDRs approval for the connection. Only valid with `connectionType` = 'bgp'. Requires the transit_gateway's `learnedCidrsApprovalMode` attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * Aviatrix transit gateway name.
     */
    gwName?: pulumi.Input<string>;
    /**
     * Set as true if there are two external devices.
     * * `backupRemoteGatewayIp ` - (Optional) Backup remote gateway IP. Required if HA enabled.
     */
    haEnabled?: pulumi.Input<boolean>;
    /**
     * Local LAN IP. Required for GCP BGP over LAN connection.
     */
    localLanIp?: pulumi.Input<string>;
    /**
     * Source CIDR for the tunnel from the Aviatrix transit gateway.
     */
    localTunnelCidr?: pulumi.Input<string>;
    /**
     * Configure manual BGP advertised CIDRs for this connection. Only valid with `connectionType`= 'bgp'. Available as of provider version R2.18+.
     */
    manualBgpAdvertisedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'. Default value: 'SHA-256'.
     */
    phase1Authentication?: pulumi.Input<string>;
    /**
     * Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.
     */
    phase1DhGroups?: pulumi.Input<string>;
    /**
     * Phase one Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", and "AES-256-GCM-128". Default value: "AES-256-CBC".
     */
    phase1Encryption?: pulumi.Input<string>;
    /**
     * Phase 1 remote identifier of the IPsec tunnel. This can be configured to be either the public IP address or the private IP address of the peer terminating the IPsec tunnel. Example: ["1.2.3.4"] when HA is disabled, ["1.2.3.4", "5.6.7.8"] when HA is enabled. Available as of provider version R2.19+.
     */
    phase1RemoteIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'. Default value: 'HMAC-SHA-256'.
     */
    phase2Authentication?: pulumi.Input<string>;
    /**
     * Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.
     */
    phase2DhGroups?: pulumi.Input<string>;
    /**
     * Phase two Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", "AES-256-GCM-128" and "NULL-ENCR". Default value: "AES-256-CBC".
     */
    phase2Encryption?: pulumi.Input<string>;
    /**
     * Pre-Shared Key.
     */
    preSharedKey?: pulumi.Input<string>;
    /**
     * Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
     */
    prependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Remote gateway IP. Required when `tunnelProtocol` != 'LAN'.
     */
    remoteGatewayIp?: pulumi.Input<string>;
    /**
     * Remote LAN IP. Required for BGP over LAN connection.
     */
    remoteLanIp?: pulumi.Input<string>;
    /**
     * Remote CIDRs joined as a string with ','. Required for a 'static' type connection.
     */
    remoteSubnet?: pulumi.Input<string>;
    /**
     * Destination CIDR for the tunnel to the external device.
     */
    remoteTunnelCidr?: pulumi.Input<string>;
    /**
     * Name of the remote VPC for a LAN BGP connection with an Azure Transit Gateway. Required when `connectionType` = 'bgp' and `tunnelProtocol` = 'LAN' with an Azure transit gateway. Must be in the format "<vnet-name>:<resource-group-name>:<subscription-id>". Available as of provider version R2.18+.
     */
    remoteVpcName?: pulumi.Input<string>;
    /**
     * Switch to HA Standby Transit Gateway connection. Only valid with Transit Gateway that has [Active-Standby Mode](https://docs.aviatrix.com/HowTos/transit_advanced.html#active-standby) enabled and for non-HA external device. Valid values: true, false. Default: false. Available in provider version R2.17.1+.
     */
    switchToHaStandbyGateway?: pulumi.Input<boolean>;
    /**
     * Tunnel protocol, only valid with `connectionType` = 'bgp'. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive. Available as of provider version R2.18+.
     */
    tunnelProtocol?: pulumi.Input<string>;
    /**
     * VPC ID of the Aviatrix transit gateway. For GCP BGP over LAN connection, it is in the format of "vpc_name~-~project_name".
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixTransitExternalDeviceConn resource.
 */
export interface AviatrixTransitExternalDeviceConnArgs {
    /**
     * Set of approved CIDRs. Requires `enableLearnedCidrsApproval` to be true. Type: Set(String).
     */
    approvedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup BGP MD5 Authentication Key. Valid with HA enabled for connection. Example: 'avx03,avx04'. For BGP LAN ActiveMesh mode disabled, example: 'avx03'.
     */
    backupBgpMd5Key?: pulumi.Input<string>;
    /**
     * Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required if HA enabled for 'bgp' connection.
     */
    backupBgpRemoteAsNum?: pulumi.Input<string>;
    /**
     * Backup direct connect for backup external device.
     */
    backupDirectConnect?: pulumi.Input<boolean>;
    /**
     * Backup Local LAN IP. Required for GCP HA BGP over LAN connection.
     */
    backupLocalLanIp?: pulumi.Input<string>;
    /**
     * Source CIDR for the tunnel from the backup Aviatrix transit gateway.
     */
    backupLocalTunnelCidr?: pulumi.Input<string>;
    /**
     * Backup Pre-Shared Key.
     */
    backupPreSharedKey?: pulumi.Input<string>;
    /**
     * Backup remote gateway IP.
     */
    backupRemoteGatewayIp?: pulumi.Input<string>;
    /**
     * Backup Remote LAN IP. Required for HA BGP over LAN connection.
     */
    backupRemoteLanIp?: pulumi.Input<string>;
    /**
     * Destination CIDR for the tunnel to the backup external device.
     */
    backupRemoteTunnelCidr?: pulumi.Input<string>;
    /**
     * BGP local ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.
     */
    bgpLocalAsNum?: pulumi.Input<string>;
    /**
     * BGP MD5 Authentication Key. Example: 'avx01,avx02'. For BGP LAN ActiveMesh mode disabled, example: 'avx01'.
     */
    bgpMd5Key?: pulumi.Input<string>;
    /**
     * BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.
     */
    bgpRemoteAsNum?: pulumi.Input<string>;
    /**
     * Transit external device connection name.
     */
    connectionName: pulumi.Input<string>;
    /**
     * Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'.
     */
    connectionType?: pulumi.Input<string>;
    /**
     * Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption. Valid values: true, false. **NOTE: Please see notes here for more information.**
     */
    customAlgorithms?: pulumi.Input<boolean>;
    /**
     * Set true for private network infrastructure.
     */
    directConnect?: pulumi.Input<boolean>;
    /**
     * Switch to enable BGP LAN ActiveMesh mode. Only valid for GCP with Remote Gateway HA enabled. Default: false. Available as of provider version R2.21+.
     */
    enableBgpLanActivemesh?: pulumi.Input<boolean>;
    /**
     * Switch to allow this connection to communicate with a Network Domain via Connection Policy.
     */
    enableEdgeSegmentation?: pulumi.Input<boolean>;
    /**
     * Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
     */
    enableEventTriggeredHa?: pulumi.Input<boolean>;
    /**
     * Set as true to enable IKEv2 protocol.
     */
    enableIkev2?: pulumi.Input<boolean>;
    /**
     * Enable Jumbo Frame for the transit external device connection. Only valid with 'GRE' tunnels under 'bgp' connection. Requires transit to be jumbo frame and insane mode enabled. Valid values: true, false. Default value: false. Available as of provider version R2.22.2+.
     */
    enableJumboFrame?: pulumi.Input<boolean>;
    /**
     * Enable learned CIDRs approval for the connection. Only valid with `connectionType` = 'bgp'. Requires the transit_gateway's `learnedCidrsApprovalMode` attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * Aviatrix transit gateway name.
     */
    gwName: pulumi.Input<string>;
    /**
     * Set as true if there are two external devices.
     * * `backupRemoteGatewayIp ` - (Optional) Backup remote gateway IP. Required if HA enabled.
     */
    haEnabled?: pulumi.Input<boolean>;
    /**
     * Local LAN IP. Required for GCP BGP over LAN connection.
     */
    localLanIp?: pulumi.Input<string>;
    /**
     * Source CIDR for the tunnel from the Aviatrix transit gateway.
     */
    localTunnelCidr?: pulumi.Input<string>;
    /**
     * Configure manual BGP advertised CIDRs for this connection. Only valid with `connectionType`= 'bgp'. Available as of provider version R2.18+.
     */
    manualBgpAdvertisedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'. Default value: 'SHA-256'.
     */
    phase1Authentication?: pulumi.Input<string>;
    /**
     * Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.
     */
    phase1DhGroups?: pulumi.Input<string>;
    /**
     * Phase one Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", and "AES-256-GCM-128". Default value: "AES-256-CBC".
     */
    phase1Encryption?: pulumi.Input<string>;
    /**
     * Phase 1 remote identifier of the IPsec tunnel. This can be configured to be either the public IP address or the private IP address of the peer terminating the IPsec tunnel. Example: ["1.2.3.4"] when HA is disabled, ["1.2.3.4", "5.6.7.8"] when HA is enabled. Available as of provider version R2.19+.
     */
    phase1RemoteIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'. Default value: 'HMAC-SHA-256'.
     */
    phase2Authentication?: pulumi.Input<string>;
    /**
     * Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.
     */
    phase2DhGroups?: pulumi.Input<string>;
    /**
     * Phase two Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", "AES-256-GCM-128" and "NULL-ENCR". Default value: "AES-256-CBC".
     */
    phase2Encryption?: pulumi.Input<string>;
    /**
     * Pre-Shared Key.
     */
    preSharedKey?: pulumi.Input<string>;
    /**
     * Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
     */
    prependAsPaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Remote gateway IP. Required when `tunnelProtocol` != 'LAN'.
     */
    remoteGatewayIp?: pulumi.Input<string>;
    /**
     * Remote LAN IP. Required for BGP over LAN connection.
     */
    remoteLanIp?: pulumi.Input<string>;
    /**
     * Remote CIDRs joined as a string with ','. Required for a 'static' type connection.
     */
    remoteSubnet?: pulumi.Input<string>;
    /**
     * Destination CIDR for the tunnel to the external device.
     */
    remoteTunnelCidr?: pulumi.Input<string>;
    /**
     * Name of the remote VPC for a LAN BGP connection with an Azure Transit Gateway. Required when `connectionType` = 'bgp' and `tunnelProtocol` = 'LAN' with an Azure transit gateway. Must be in the format "<vnet-name>:<resource-group-name>:<subscription-id>". Available as of provider version R2.18+.
     */
    remoteVpcName?: pulumi.Input<string>;
    /**
     * Switch to HA Standby Transit Gateway connection. Only valid with Transit Gateway that has [Active-Standby Mode](https://docs.aviatrix.com/HowTos/transit_advanced.html#active-standby) enabled and for non-HA external device. Valid values: true, false. Default: false. Available in provider version R2.17.1+.
     */
    switchToHaStandbyGateway?: pulumi.Input<boolean>;
    /**
     * Tunnel protocol, only valid with `connectionType` = 'bgp'. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive. Available as of provider version R2.18+.
     */
    tunnelProtocol?: pulumi.Input<string>;
    /**
     * VPC ID of the Aviatrix transit gateway. For GCP BGP over LAN connection, it is in the format of "vpc_name~-~project_name".
     */
    vpcId: pulumi.Input<string>;
}
