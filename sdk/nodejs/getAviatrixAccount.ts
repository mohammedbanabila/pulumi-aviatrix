// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_account** data source provides details about a specific cloud account created by the Aviatrix Controller.
 *
 * This data source can prove useful when a module accepts an account's detail as an input variable.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Aviatrix Account Data Source
 * const foo = pulumi.output(aviatrix.getAviatrixAccount({
 *     accountName: "username",
 * }));
 * ```
 */
export function getAviatrixAccount(args: GetAviatrixAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAviatrixAccountResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aviatrix:index/getAviatrixAccount:getAviatrixAccount", {
        "accountName": args.accountName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAviatrixAccount.
 */
export interface GetAviatrixAccountArgs {
    /**
     * Account name. This can be used for logging in to CloudN console or UserConnect controller.
     */
    accountName: string;
}

/**
 * A collection of values returned by getAviatrixAccount.
 */
export interface GetAviatrixAccountResult {
    readonly accountName: string;
    /**
     * Alibaba Cloud Account ID.
     */
    readonly alicloudAccountId: string;
    /**
     * Azure ARM Subscription ID.
     */
    readonly armSubscriptionId: string;
    /**
     * AWS Account number.
     */
    readonly awsAccountNumber: string;
    /**
     * AWS Top Secret Region or Secret Region Custom Certificate Authority file name on the controller. Available as of provider R2.19.5+.
     */
    readonly awsCaCertPath: string;
    /**
     * A separate AWS App role ARN to assign to gateways created by the controller. Available as of provider version R2.19+.
     */
    readonly awsGatewayRoleApp: string;
    /**
     * A separate AWS EC2 role ARN to assign to gateways created by the controller. Available as of provider version R2.19+.
     */
    readonly awsGatewayRoleEc2: string;
    /**
     * AWS App role ARN.
     */
    readonly awsRoleArn: string;
    /**
     * AWS EC2 role ARN.
     */
    readonly awsRoleEc2: string;
    /**
     * AWSChina Account number. Available as of provider version R2.19+.
     */
    readonly awschinaAccountNumber: string;
    /**
     * If enabled, `awschinaRoleApp` and `awschinaRoleEc2` will be set. Available as of provider version R2.19+.
     */
    readonly awschinaIam: boolean;
    /**
     * AWSChina App role ARN. Available as of provider version R2.19+.
     */
    readonly awschinaRoleApp: string;
    /**
     * AWSChina EC2 role ARN. Available as of provider version R2.19+.
     */
    readonly awschinaRoleEc2: string;
    /**
     * AWSGov Account number.
     */
    readonly awsgovAccountNumber: string;
    /**
     * If enabled, `awsgovRoleApp` and `awschinaRoleEc2` will be set. Available as of provider version R2.19+.
     */
    readonly awsgovIam: boolean;
    /**
     * AWSGov App role ARN. Available as of provider version R2.19+.
     */
    readonly awsgovRoleApp: string;
    /**
     * AWSGov EC2 role ARN. Available as of provider version R2.19+.
     */
    readonly awsgovRoleEc2: string;
    /**
     * AWS Secret Region Account Number. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awssAccountNumber: string;
    /**
     * AWS Secret Region Account Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awssCapAccountName: string;
    /**
     * AWS Secret Region CAP Agency. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awssCapAgency: string;
    /**
     * AWS Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.
     */
    readonly awssCapCertKeyPath: string;
    /**
     * AWS Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.
     */
    readonly awssCapCertPath: string;
    /**
     * AWS Secret Region Role Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awssCapRoleName: string;
    /**
     * AWS Secret Region CAP Url. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awssCapUrl: string;
    /**
     * AWS Top Secret Region Account Number. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awstsAccountNumber: string;
    /**
     * AWS Top Secret Region CAP Agency. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awstsCapAgency: string;
    /**
     * AWS Top Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.
     */
    readonly awstsCapCertKeyPath: string;
    readonly awstsCapCertPath: string;
    /**
     * AWS Top Secret Region Mission. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awstsCapMission: string;
    /**
     * AWS Top Secret Region Role Name. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
     * `awstsCapCertPath` - AWS Top Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.
     */
    readonly awstsCapRoleName: string;
    /**
     * AWS Top Secret Region CAP Url. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
     */
    readonly awstsCapUrl: string;
    /**
     * AzureChina ARM Subscription ID. Available as of provider version R2.19+.
     */
    readonly azurechinaSubscriptionId: string;
    /**
     * AzureGov ARM Subscription ID.
     */
    readonly azuregovSubscriptionId: string;
    /**
     * Type of cloud service provider.
     */
    readonly cloudType: number;
    /**
     * GCloud Project ID.
     */
    readonly gcloudProjectId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

export function getAviatrixAccountOutput(args: GetAviatrixAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAviatrixAccountResult> {
    return pulumi.output(args).apply(a => getAviatrixAccount(a, opts))
}

/**
 * A collection of arguments for invoking getAviatrixAccount.
 */
export interface GetAviatrixAccountOutputArgs {
    /**
     * Account name. This can be used for logging in to CloudN console or UserConnect controller.
     */
    accountName: pulumi.Input<string>;
}
