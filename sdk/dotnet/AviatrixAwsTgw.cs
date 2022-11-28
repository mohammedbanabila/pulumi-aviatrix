// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    /// <summary>
    /// The **aviatrix_aws_tgw** resource allows the creation and management of Aviatrix-created AWS TGWs.
    /// 
    /// &gt; **NOTE:** If you are planning to attach VPCs to the **aviatrix_aws_tgw** resource and anticipate updating it often and/or using advanced options such as customized route advertisement, we highly recommend managing those VPCs outside this resource by setting `manage_vpc_attachment` to false and using the **aviatrix_aws_tgw_vpc_attachment** resource instead of the in-line `attached_vpc {}` block.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an Aviatrix AWS TGW
    ///     var testAwsTgw = new Aviatrix.AviatrixAwsTgw("testAwsTgw", new()
    ///     {
    ///         AccountName = "devops",
    ///         AwsSideAsNumber = "64512",
    ///         ManageTransitGatewayAttachment = false,
    ///         ManageVpcAttachment = false,
    ///         Region = "us-east-1",
    ///         SecurityDomains = new[]
    ///         {
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Default_Domain",
    ///                     "Shared_Service_Domain",
    ///                     "mysdn1",
    ///                 },
    ///                 SecurityDomainName = "Aviatrix_Edge_Domain",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Aviatrix_Edge_Domain",
    ///                     "Shared_Service_Domain",
    ///                 },
    ///                 SecurityDomainName = "Default_Domain",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Aviatrix_Edge_Domain",
    ///                     "Default_Domain",
    ///                 },
    ///                 SecurityDomainName = "Shared_Service_Domain",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Aviatrix_Edge_Domain",
    ///                 },
    ///                 SecurityDomainName = "SDN1",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 SecurityDomainName = "mysdn2",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 AviatrixFirewall = true,
    ///                 SecurityDomainName = "firewall-domain",
    ///             },
    ///         },
    ///         TgwName = "test-AWS-TGW",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an Aviatrix AWSGov TGW
    ///     var testAwsGovTgw = new Aviatrix.AviatrixAwsTgw("testAwsGovTgw", new()
    ///     {
    ///         AccountName = "devops",
    ///         AwsSideAsNumber = "64512",
    ///         CloudType = 256,
    ///         ManageTransitGatewayAttachment = false,
    ///         ManageVpcAttachment = false,
    ///         Region = "us-gov-east-1",
    ///         SecurityDomains = new[]
    ///         {
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Default_Domain",
    ///                     "Shared_Service_Domain",
    ///                     "mysdn1",
    ///                 },
    ///                 SecurityDomainName = "Aviatrix_Edge_Domain",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Aviatrix_Edge_Domain",
    ///                     "Shared_Service_Domain",
    ///                 },
    ///                 SecurityDomainName = "Default_Domain",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Aviatrix_Edge_Domain",
    ///                     "Default_Domain",
    ///                 },
    ///                 SecurityDomainName = "Shared_Service_Domain",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 ConnectedDomains = new[]
    ///                 {
    ///                     "Aviatrix_Edge_Domain",
    ///                 },
    ///                 SecurityDomainName = "SDN1",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 SecurityDomainName = "mysdn2",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixAwsTgwSecurityDomainArgs
    ///             {
    ///                 AviatrixFirewall = true,
    ///                 SecurityDomainName = "firewall-domain",
    ///             },
    ///         },
    ///         TgwName = "test-AWSGov-TGW",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **aws_tgw** can be imported using the `tgw_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw test tgw_name
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw")]
    public partial class AviatrixAwsTgw : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the cloud account in the Aviatrix controller.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
        /// </summary>
        [Output("attachedAviatrixTransitGateways")]
        public Output<ImmutableArray<string>> AttachedAviatrixTransitGateways { get; private set; } = null!;

        /// <summary>
        /// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
        /// </summary>
        [Output("awsSideAsNumber")]
        public Output<string> AwsSideAsNumber { get; private set; } = null!;

        /// <summary>
        /// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
        /// </summary>
        [Output("cidrs")]
        public Output<ImmutableArray<string>> Cidrs { get; private set; } = null!;

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
        /// </summary>
        [Output("cloudType")]
        public Output<int?> CloudType { get; private set; } = null!;

        /// <summary>
        /// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
        /// </summary>
        [Output("enableMulticast")]
        public Output<bool?> EnableMulticast { get; private set; } = null!;

        /// <summary>
        /// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
        /// </summary>
        [Output("inspectionMode")]
        public Output<string?> InspectionMode { get; private set; } = null!;

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Output("manageSecurityDomain")]
        public Output<bool?> ManageSecurityDomain { get; private set; } = null!;

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Output("manageTransitGatewayAttachment")]
        public Output<bool?> ManageTransitGatewayAttachment { get; private set; } = null!;

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Output("manageVpcAttachment")]
        public Output<bool?> ManageVpcAttachment { get; private set; } = null!;

        /// <summary>
        /// AWS region of AWS TGW to be created in
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
        /// </summary>
        [Output("securityDomains")]
        public Output<ImmutableArray<Outputs.AviatrixAwsTgwSecurityDomain>> SecurityDomains { get; private set; } = null!;

        /// <summary>
        /// TGW ID. Available as of provider version R2.19+.
        /// </summary>
        [Output("tgwId")]
        public Output<string> TgwId { get; private set; } = null!;

        /// <summary>
        /// Name of the AWS TGW to be created
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgw resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgw(string name, AviatrixAwsTgwArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw", name, args ?? new AviatrixAwsTgwArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgw(string name, Input<string> id, AviatrixAwsTgwState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/astipkovits",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixAwsTgw resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgw Get(string name, Input<string> id, AviatrixAwsTgwState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgw(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the cloud account in the Aviatrix controller.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("attachedAviatrixTransitGateways")]
        private InputList<string>? _attachedAviatrixTransitGateways;

        /// <summary>
        /// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
        /// </summary>
        [Obsolete(@"Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.")]
        public InputList<string> AttachedAviatrixTransitGateways
        {
            get => _attachedAviatrixTransitGateways ?? (_attachedAviatrixTransitGateways = new InputList<string>());
            set => _attachedAviatrixTransitGateways = value;
        }

        /// <summary>
        /// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
        /// </summary>
        [Input("awsSideAsNumber", required: true)]
        public Input<string> AwsSideAsNumber { get; set; } = null!;

        [Input("cidrs")]
        private InputList<string>? _cidrs;

        /// <summary>
        /// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
        /// </summary>
        public InputList<string> Cidrs
        {
            get => _cidrs ?? (_cidrs = new InputList<string>());
            set => _cidrs = value;
        }

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
        /// </summary>
        [Input("enableMulticast")]
        public Input<bool>? EnableMulticast { get; set; }

        /// <summary>
        /// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
        /// </summary>
        [Input("inspectionMode")]
        public Input<string>? InspectionMode { get; set; }

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("manageSecurityDomain")]
        public Input<bool>? ManageSecurityDomain { get; set; }

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("manageTransitGatewayAttachment")]
        public Input<bool>? ManageTransitGatewayAttachment { get; set; }

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("manageVpcAttachment")]
        public Input<bool>? ManageVpcAttachment { get; set; }

        /// <summary>
        /// AWS region of AWS TGW to be created in
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("securityDomains")]
        private InputList<Inputs.AviatrixAwsTgwSecurityDomainArgs>? _securityDomains;

        /// <summary>
        /// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
        /// </summary>
        [Obsolete(@"Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.")]
        public InputList<Inputs.AviatrixAwsTgwSecurityDomainArgs> SecurityDomains
        {
            get => _securityDomains ?? (_securityDomains = new InputList<Inputs.AviatrixAwsTgwSecurityDomainArgs>());
            set => _securityDomains = value;
        }

        /// <summary>
        /// Name of the AWS TGW to be created
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        public AviatrixAwsTgwArgs()
        {
        }
        public static new AviatrixAwsTgwArgs Empty => new AviatrixAwsTgwArgs();
    }

    public sealed class AviatrixAwsTgwState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the cloud account in the Aviatrix controller.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("attachedAviatrixTransitGateways")]
        private InputList<string>? _attachedAviatrixTransitGateways;

        /// <summary>
        /// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
        /// </summary>
        [Obsolete(@"Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.")]
        public InputList<string> AttachedAviatrixTransitGateways
        {
            get => _attachedAviatrixTransitGateways ?? (_attachedAviatrixTransitGateways = new InputList<string>());
            set => _attachedAviatrixTransitGateways = value;
        }

        /// <summary>
        /// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
        /// </summary>
        [Input("awsSideAsNumber")]
        public Input<string>? AwsSideAsNumber { get; set; }

        [Input("cidrs")]
        private InputList<string>? _cidrs;

        /// <summary>
        /// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
        /// </summary>
        public InputList<string> Cidrs
        {
            get => _cidrs ?? (_cidrs = new InputList<string>());
            set => _cidrs = value;
        }

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
        /// </summary>
        [Input("enableMulticast")]
        public Input<bool>? EnableMulticast { get; set; }

        /// <summary>
        /// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
        /// </summary>
        [Input("inspectionMode")]
        public Input<string>? InspectionMode { get; set; }

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("manageSecurityDomain")]
        public Input<bool>? ManageSecurityDomain { get; set; }

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("manageTransitGatewayAttachment")]
        public Input<bool>? ManageTransitGatewayAttachment { get; set; }

        /// <summary>
        /// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("manageVpcAttachment")]
        public Input<bool>? ManageVpcAttachment { get; set; }

        /// <summary>
        /// AWS region of AWS TGW to be created in
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityDomains")]
        private InputList<Inputs.AviatrixAwsTgwSecurityDomainGetArgs>? _securityDomains;

        /// <summary>
        /// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
        /// </summary>
        [Obsolete(@"Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.")]
        public InputList<Inputs.AviatrixAwsTgwSecurityDomainGetArgs> SecurityDomains
        {
            get => _securityDomains ?? (_securityDomains = new InputList<Inputs.AviatrixAwsTgwSecurityDomainGetArgs>());
            set => _securityDomains = value;
        }

        /// <summary>
        /// TGW ID. Available as of provider version R2.19+.
        /// </summary>
        [Input("tgwId")]
        public Input<string>? TgwId { get; set; }

        /// <summary>
        /// Name of the AWS TGW to be created
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        public AviatrixAwsTgwState()
        {
        }
        public static new AviatrixAwsTgwState Empty => new AviatrixAwsTgwState();
    }
}
