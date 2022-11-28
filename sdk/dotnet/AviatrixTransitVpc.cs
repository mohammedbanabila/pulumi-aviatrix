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
    /// The aviatrix.AviatrixTransitVpc resource creates and manages the Aviatrix Transit Network Gateways.
    /// 
    /// !&gt; **WARNING:** The `aviatrix.AviatrixTransitVpc` resource is deprecated as of **Release 2.0**. It is currently kept for backward-compatibility and will be removed in the future. Please use the transit gateway resource instead. If this is already in the state, please remove it from state file and import as `aviatrix.AviatrixTransitGateway`.
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
    ///     // Manage Aviatrix Transit Network Gateways in aws
    ///     var testTransitGwAws = new Aviatrix.AviatrixTransitVpc("testTransitGwAws", new()
    ///     {
    ///         AccountName = "devops_aws",
    ///         CloudType = 1,
    ///         ConnectedTransit = "yes",
    ///         EnableHybridConnection = true,
    ///         GwName = "transit",
    ///         HaGwSize = "t2.micro",
    ///         HaSubnet = "10.1.0.0/24",
    ///         Subnet = "10.1.0.0/24",
    ///         TagLists = new[]
    ///         {
    ///             "name:value",
    ///             "name1:value1",
    ///             "name2:value2",
    ///         },
    ///         VpcId = "vpc-abcd1234",
    ///         VpcReg = "us-east-1",
    ///         VpcSize = "t2.micro",
    ///     });
    /// 
    ///     // Manage Aviatrix Transit Network Gateways in azure
    ///     var testTransitGwAzure = new Aviatrix.AviatrixTransitVpc("testTransitGwAzure", new()
    ///     {
    ///         AccountName = "devops_azure",
    ///         CloudType = 8,
    ///         ConnectedTransit = "yes",
    ///         GwName = "transit",
    ///         HaGwSize = "Standard_B1s",
    ///         HaSubnet = "10.30.0.0/24",
    ///         Subnet = "10.30.0.0/24",
    ///         VpcId = "vnet1:hello",
    ///         VpcReg = "West US",
    ///         VpcSize = "Standard_B1s",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance transit_vpc can be imported using the gw_name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixTransitVpc:AviatrixTransitVpc test gw_name
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixTransitVpc:AviatrixTransitVpc")]
    public partial class AviatrixTransitVpc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Use 1 for AWS.
        /// </summary>
        [Output("cloudType")]
        public Output<int> CloudType { get; private set; } = null!;

        /// <summary>
        /// Specify Connected Transit status. Supported values: true, false.
        /// </summary>
        [Output("connectedTransit")]
        public Output<string?> ConnectedTransit { get; private set; } = null!;

        /// <summary>
        /// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
        /// </summary>
        [Output("enableFirenetInterfaces")]
        public Output<bool?> EnableFirenetInterfaces { get; private set; } = null!;

        /// <summary>
        /// Sign of readiness for TGW connection. Only supported for aws. Example: false.
        /// </summary>
        [Output("enableHybridConnection")]
        public Output<bool?> EnableHybridConnection { get; private set; } = null!;

        /// <summary>
        /// Enable NAT for this container.
        /// </summary>
        [Output("enableNat")]
        public Output<string?> EnableNat { get; private set; } = null!;

        /// <summary>
        /// Name of the gateway which is going to be created.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
        /// </summary>
        [Output("haGwSize")]
        public Output<string?> HaGwSize { get; private set; } = null!;

        /// <summary>
        /// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insane_mode is enabled and ha_subnet is set.
        /// </summary>
        [Output("haInsaneModeAz")]
        public Output<string?> HaInsaneModeAz { get; private set; } = null!;

        /// <summary>
        /// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
        /// </summary>
        [Output("haSubnet")]
        public Output<string?> HaSubnet { get; private set; } = null!;

        /// <summary>
        /// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
        /// </summary>
        [Output("insaneMode")]
        public Output<bool?> InsaneMode { get; private set; } = null!;

        /// <summary>
        /// AZ of subnet being created for Insane Mode Transit Gateway. Required if insane_mode is enabled.
        /// </summary>
        [Output("insaneModeAz")]
        public Output<string?> InsaneModeAz { get; private set; } = null!;

        /// <summary>
        /// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
        /// </summary>
        [Output("tagLists")]
        public Output<ImmutableArray<string>> TagLists { get; private set; } = null!;

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
        /// </summary>
        [Output("vpcReg")]
        public Output<string> VpcReg { get; private set; } = null!;

        /// <summary>
        /// Size of the gateway instance. Example: AWS: "t2.large", etc...
        /// </summary>
        [Output("vpcSize")]
        public Output<string> VpcSize { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixTransitVpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixTransitVpc(string name, AviatrixTransitVpcArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixTransitVpc:AviatrixTransitVpc", name, args ?? new AviatrixTransitVpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixTransitVpc(string name, Input<string> id, AviatrixTransitVpcState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixTransitVpc:AviatrixTransitVpc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixTransitVpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixTransitVpc Get(string name, Input<string> id, AviatrixTransitVpcState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixTransitVpc(name, id, state, options);
        }
    }

    public sealed class AviatrixTransitVpcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Use 1 for AWS.
        /// </summary>
        [Input("cloudType", required: true)]
        public Input<int> CloudType { get; set; } = null!;

        /// <summary>
        /// Specify Connected Transit status. Supported values: true, false.
        /// </summary>
        [Input("connectedTransit")]
        public Input<string>? ConnectedTransit { get; set; }

        /// <summary>
        /// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
        /// </summary>
        [Input("enableFirenetInterfaces")]
        public Input<bool>? EnableFirenetInterfaces { get; set; }

        /// <summary>
        /// Sign of readiness for TGW connection. Only supported for aws. Example: false.
        /// </summary>
        [Input("enableHybridConnection")]
        public Input<bool>? EnableHybridConnection { get; set; }

        /// <summary>
        /// Enable NAT for this container.
        /// </summary>
        [Input("enableNat")]
        public Input<string>? EnableNat { get; set; }

        /// <summary>
        /// Name of the gateway which is going to be created.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
        /// </summary>
        [Input("haGwSize")]
        public Input<string>? HaGwSize { get; set; }

        /// <summary>
        /// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insane_mode is enabled and ha_subnet is set.
        /// </summary>
        [Input("haInsaneModeAz")]
        public Input<string>? HaInsaneModeAz { get; set; }

        /// <summary>
        /// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
        /// </summary>
        [Input("haSubnet")]
        public Input<string>? HaSubnet { get; set; }

        /// <summary>
        /// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
        /// </summary>
        [Input("insaneMode")]
        public Input<bool>? InsaneMode { get; set; }

        /// <summary>
        /// AZ of subnet being created for Insane Mode Transit Gateway. Required if insane_mode is enabled.
        /// </summary>
        [Input("insaneModeAz")]
        public Input<string>? InsaneModeAz { get; set; }

        /// <summary>
        /// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
        /// </summary>
        [Input("vpcReg", required: true)]
        public Input<string> VpcReg { get; set; } = null!;

        /// <summary>
        /// Size of the gateway instance. Example: AWS: "t2.large", etc...
        /// </summary>
        [Input("vpcSize", required: true)]
        public Input<string> VpcSize { get; set; } = null!;

        public AviatrixTransitVpcArgs()
        {
        }
        public static new AviatrixTransitVpcArgs Empty => new AviatrixTransitVpcArgs();
    }

    public sealed class AviatrixTransitVpcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Use 1 for AWS.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Specify Connected Transit status. Supported values: true, false.
        /// </summary>
        [Input("connectedTransit")]
        public Input<string>? ConnectedTransit { get; set; }

        /// <summary>
        /// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
        /// </summary>
        [Input("enableFirenetInterfaces")]
        public Input<bool>? EnableFirenetInterfaces { get; set; }

        /// <summary>
        /// Sign of readiness for TGW connection. Only supported for aws. Example: false.
        /// </summary>
        [Input("enableHybridConnection")]
        public Input<bool>? EnableHybridConnection { get; set; }

        /// <summary>
        /// Enable NAT for this container.
        /// </summary>
        [Input("enableNat")]
        public Input<string>? EnableNat { get; set; }

        /// <summary>
        /// Name of the gateway which is going to be created.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
        /// </summary>
        [Input("haGwSize")]
        public Input<string>? HaGwSize { get; set; }

        /// <summary>
        /// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insane_mode is enabled and ha_subnet is set.
        /// </summary>
        [Input("haInsaneModeAz")]
        public Input<string>? HaInsaneModeAz { get; set; }

        /// <summary>
        /// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
        /// </summary>
        [Input("haSubnet")]
        public Input<string>? HaSubnet { get; set; }

        /// <summary>
        /// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
        /// </summary>
        [Input("insaneMode")]
        public Input<bool>? InsaneMode { get; set; }

        /// <summary>
        /// AZ of subnet being created for Insane Mode Transit Gateway. Required if insane_mode is enabled.
        /// </summary>
        [Input("insaneModeAz")]
        public Input<string>? InsaneModeAz { get; set; }

        /// <summary>
        /// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
        /// </summary>
        [Input("vpcReg")]
        public Input<string>? VpcReg { get; set; }

        /// <summary>
        /// Size of the gateway instance. Example: AWS: "t2.large", etc...
        /// </summary>
        [Input("vpcSize")]
        public Input<string>? VpcSize { get; set; }

        public AviatrixTransitVpcState()
        {
        }
        public static new AviatrixTransitVpcState Empty => new AviatrixTransitVpcState();
    }
}
