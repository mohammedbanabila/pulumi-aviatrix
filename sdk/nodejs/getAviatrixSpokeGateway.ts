// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAviatrixSpokeGateway(args: GetAviatrixSpokeGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetAviatrixSpokeGatewayResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aviatrix:index/getAviatrixSpokeGateway:getAviatrixSpokeGateway", {
        "gwName": args.gwName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAviatrixSpokeGateway.
 */
export interface GetAviatrixSpokeGatewayArgs {
    gwName: string;
}

/**
 * A collection of values returned by getAviatrixSpokeGateway.
 */
export interface GetAviatrixSpokeGatewayResult {
    readonly accountName: string;
    readonly allocateNewEip: boolean;
    readonly approvedLearnedCidrs: string[];
    readonly availabilityDomain: string;
    readonly azureEipNameResourceGroup: string;
    readonly bgpEcmp: boolean;
    readonly bgpHoldTime: number;
    readonly bgpPollingTime: number;
    readonly cloudInstanceId: string;
    readonly cloudType: number;
    readonly customizedSpokeVpcRoutes: string;
    readonly disableRoutePropagation: boolean;
    readonly eip: string;
    readonly enableActiveStandby: boolean;
    readonly enableActiveStandbyPreemptive: boolean;
    readonly enableAutoAdvertiseS2cCidrs: boolean;
    readonly enableBgp: boolean;
    readonly enableEncryptVolume: boolean;
    readonly enableJumboFrame: boolean;
    readonly enableLearnedCidrsApproval: boolean;
    readonly enableMonitorGatewaySubnets: boolean;
    readonly enablePrivateOob: boolean;
    readonly enablePrivateVpcDefaultRoute: boolean;
    readonly enableSkipPublicRouteTableUpdate: boolean;
    readonly enableSpotInstance: boolean;
    readonly enableVpcDnsServer: boolean;
    readonly faultDomain: string;
    readonly filteredSpokeVpcRoutes: string;
    readonly gwName: string;
    readonly gwSize: string;
    readonly haAvailabilityDomain: string;
    readonly haAzureEipNameResourceGroup: string;
    readonly haCloudInstanceId: string;
    readonly haEip: string;
    readonly haFaultDomain: string;
    readonly haGwName: string;
    readonly haGwSize: string;
    readonly haImageVersion: string;
    readonly haInsaneModeAz: string;
    readonly haOobAvailabilityZone: string;
    readonly haOobManagementSubnet: string;
    readonly haPrivateIp: string;
    readonly haPublicIp: string;
    readonly haSecurityGroupId: string;
    readonly haSoftwareVersion: string;
    readonly haSubnet: string;
    readonly haZone: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageVersion: string;
    readonly includedAdvertisedSpokeRoutes: string;
    readonly insaneMode: boolean;
    readonly insaneModeAz: string;
    readonly learnedCidrsApprovalMode: string;
    readonly localAsNumber: string;
    readonly monitorExcludeLists: string[];
    readonly oobAvailabilityZone: string;
    readonly oobManagementSubnet: string;
    readonly prependAsPaths: string[];
    readonly privateIp: string;
    readonly publicIp: string;
    readonly securityGroupId: string;
    readonly singleAzHa: boolean;
    readonly singleIpSnat: boolean;
    readonly softwareVersion: string;
    readonly spokeBgpManualAdvertiseCidrs: string[];
    readonly spotPrice: string;
    readonly subnet: string;
    readonly tagLists: string[];
    readonly tags: {[key: string]: string};
    readonly transitGw: string;
    readonly tunnelDetectionTime: number;
    readonly vpcId: string;
    readonly vpcReg: string;
    readonly zone: string;
}

export function getAviatrixSpokeGatewayOutput(args: GetAviatrixSpokeGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAviatrixSpokeGatewayResult> {
    return pulumi.output(args).apply(a => getAviatrixSpokeGateway(a, opts))
}

/**
 * A collection of arguments for invoking getAviatrixSpokeGateway.
 */
export interface GetAviatrixSpokeGatewayOutputArgs {
    gwName: pulumi.Input<string>;
}
