// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    public static class GetAviatrixAccount
    {
        /// <summary>
        /// The **aviatrix_account** data source provides details about a specific cloud account created by the Aviatrix Controller.
        /// 
        /// This data source can prove useful when a module accepts an account's detail as an input variable.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aviatrix = Pulumi.Aviatrix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aviatrix.GetAviatrixAccount.Invoke(new()
        ///     {
        ///         AccountName = "username",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAviatrixAccountResult> InvokeAsync(GetAviatrixAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAviatrixAccountResult>("aviatrix:index/getAviatrixAccount:getAviatrixAccount", args ?? new GetAviatrixAccountArgs(), options.WithDefaults());

        /// <summary>
        /// The **aviatrix_account** data source provides details about a specific cloud account created by the Aviatrix Controller.
        /// 
        /// This data source can prove useful when a module accepts an account's detail as an input variable.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aviatrix = Pulumi.Aviatrix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aviatrix.GetAviatrixAccount.Invoke(new()
        ///     {
        ///         AccountName = "username",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAviatrixAccountResult> Invoke(GetAviatrixAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAviatrixAccountResult>("aviatrix:index/getAviatrixAccount:getAviatrixAccount", args ?? new GetAviatrixAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAviatrixAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Account name. This can be used for logging in to CloudN console or UserConnect controller.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        public GetAviatrixAccountArgs()
        {
        }
        public static new GetAviatrixAccountArgs Empty => new GetAviatrixAccountArgs();
    }

    public sealed class GetAviatrixAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Account name. This can be used for logging in to CloudN console or UserConnect controller.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        public GetAviatrixAccountInvokeArgs()
        {
        }
        public static new GetAviatrixAccountInvokeArgs Empty => new GetAviatrixAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetAviatrixAccountResult
    {
        public readonly string AccountName;
        /// <summary>
        /// Alibaba Cloud Account ID.
        /// </summary>
        public readonly string AlicloudAccountId;
        /// <summary>
        /// Azure ARM Subscription ID.
        /// </summary>
        public readonly string ArmSubscriptionId;
        /// <summary>
        /// AWS Account number.
        /// </summary>
        public readonly string AwsAccountNumber;
        /// <summary>
        /// AWS Top Secret Region or Secret Region Custom Certificate Authority file name on the controller. Available as of provider R2.19.5+.
        /// </summary>
        public readonly string AwsCaCertPath;
        /// <summary>
        /// A separate AWS App role ARN to assign to gateways created by the controller. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AwsGatewayRoleApp;
        /// <summary>
        /// A separate AWS EC2 role ARN to assign to gateways created by the controller. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AwsGatewayRoleEc2;
        /// <summary>
        /// AWS App role ARN.
        /// </summary>
        public readonly string AwsRoleArn;
        /// <summary>
        /// AWS EC2 role ARN.
        /// </summary>
        public readonly string AwsRoleEc2;
        /// <summary>
        /// AWSChina Account number. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AwschinaAccountNumber;
        /// <summary>
        /// If enabled, `awschina_role_app` and `awschina_role_ec2` will be set. Available as of provider version R2.19+.
        /// </summary>
        public readonly bool AwschinaIam;
        /// <summary>
        /// AWSChina App role ARN. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AwschinaRoleApp;
        /// <summary>
        /// AWSChina EC2 role ARN. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AwschinaRoleEc2;
        /// <summary>
        /// AWSGov Account number.
        /// </summary>
        public readonly string AwsgovAccountNumber;
        /// <summary>
        /// If enabled, `awsgov_role_app` and `awschina_role_ec2` will be set. Available as of provider version R2.19+.
        /// </summary>
        public readonly bool AwsgovIam;
        /// <summary>
        /// AWSGov App role ARN. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AwsgovRoleApp;
        /// <summary>
        /// AWSGov EC2 role ARN. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AwsgovRoleEc2;
        /// <summary>
        /// AWS Secret Region Account Number. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwssAccountNumber;
        /// <summary>
        /// AWS Secret Region Account Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwssCapAccountName;
        /// <summary>
        /// AWS Secret Region CAP Agency. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwssCapAgency;
        /// <summary>
        /// AWS Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.
        /// </summary>
        public readonly string AwssCapCertKeyPath;
        /// <summary>
        /// AWS Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.
        /// </summary>
        public readonly string AwssCapCertPath;
        /// <summary>
        /// AWS Secret Region Role Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwssCapRoleName;
        /// <summary>
        /// AWS Secret Region CAP Url. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwssCapUrl;
        /// <summary>
        /// AWS Top Secret Region Account Number. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwstsAccountNumber;
        /// <summary>
        /// AWS Top Secret Region CAP Agency. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwstsCapAgency;
        /// <summary>
        /// AWS Top Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.
        /// </summary>
        public readonly string AwstsCapCertKeyPath;
        public readonly string AwstsCapCertPath;
        /// <summary>
        /// AWS Top Secret Region Mission. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwstsCapMission;
        /// <summary>
        /// AWS Top Secret Region Role Name. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
        /// `awsts_cap_cert_path` - AWS Top Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.
        /// </summary>
        public readonly string AwstsCapRoleName;
        /// <summary>
        /// AWS Top Secret Region CAP Url. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.
        /// </summary>
        public readonly string AwstsCapUrl;
        /// <summary>
        /// AzureChina ARM Subscription ID. Available as of provider version R2.19+.
        /// </summary>
        public readonly string AzurechinaSubscriptionId;
        /// <summary>
        /// AzureGov ARM Subscription ID.
        /// </summary>
        public readonly string AzuregovSubscriptionId;
        /// <summary>
        /// Type of cloud service provider.
        /// </summary>
        public readonly int CloudType;
        /// <summary>
        /// GCloud Project ID.
        /// </summary>
        public readonly string GcloudProjectId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAviatrixAccountResult(
            string accountName,

            string alicloudAccountId,

            string armSubscriptionId,

            string awsAccountNumber,

            string awsCaCertPath,

            string awsGatewayRoleApp,

            string awsGatewayRoleEc2,

            string awsRoleArn,

            string awsRoleEc2,

            string awschinaAccountNumber,

            bool awschinaIam,

            string awschinaRoleApp,

            string awschinaRoleEc2,

            string awsgovAccountNumber,

            bool awsgovIam,

            string awsgovRoleApp,

            string awsgovRoleEc2,

            string awssAccountNumber,

            string awssCapAccountName,

            string awssCapAgency,

            string awssCapCertKeyPath,

            string awssCapCertPath,

            string awssCapRoleName,

            string awssCapUrl,

            string awstsAccountNumber,

            string awstsCapAgency,

            string awstsCapCertKeyPath,

            string awstsCapCertPath,

            string awstsCapMission,

            string awstsCapRoleName,

            string awstsCapUrl,

            string azurechinaSubscriptionId,

            string azuregovSubscriptionId,

            int cloudType,

            string gcloudProjectId,

            string id)
        {
            AccountName = accountName;
            AlicloudAccountId = alicloudAccountId;
            ArmSubscriptionId = armSubscriptionId;
            AwsAccountNumber = awsAccountNumber;
            AwsCaCertPath = awsCaCertPath;
            AwsGatewayRoleApp = awsGatewayRoleApp;
            AwsGatewayRoleEc2 = awsGatewayRoleEc2;
            AwsRoleArn = awsRoleArn;
            AwsRoleEc2 = awsRoleEc2;
            AwschinaAccountNumber = awschinaAccountNumber;
            AwschinaIam = awschinaIam;
            AwschinaRoleApp = awschinaRoleApp;
            AwschinaRoleEc2 = awschinaRoleEc2;
            AwsgovAccountNumber = awsgovAccountNumber;
            AwsgovIam = awsgovIam;
            AwsgovRoleApp = awsgovRoleApp;
            AwsgovRoleEc2 = awsgovRoleEc2;
            AwssAccountNumber = awssAccountNumber;
            AwssCapAccountName = awssCapAccountName;
            AwssCapAgency = awssCapAgency;
            AwssCapCertKeyPath = awssCapCertKeyPath;
            AwssCapCertPath = awssCapCertPath;
            AwssCapRoleName = awssCapRoleName;
            AwssCapUrl = awssCapUrl;
            AwstsAccountNumber = awstsAccountNumber;
            AwstsCapAgency = awstsCapAgency;
            AwstsCapCertKeyPath = awstsCapCertKeyPath;
            AwstsCapCertPath = awstsCapCertPath;
            AwstsCapMission = awstsCapMission;
            AwstsCapRoleName = awstsCapRoleName;
            AwstsCapUrl = awstsCapUrl;
            AzurechinaSubscriptionId = azurechinaSubscriptionId;
            AzuregovSubscriptionId = azuregovSubscriptionId;
            CloudType = cloudType;
            GcloudProjectId = gcloudProjectId;
            Id = id;
        }
    }
}
