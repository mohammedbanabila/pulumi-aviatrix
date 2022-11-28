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
    /// The aviatrix.AviatrixSpokeVpc resource allows to create and manage Aviatrix Spoke Gateways.
    /// 
    /// !&gt; **WARNING:** The `aviatrix.AviatrixSpokeVpc` resource is deprecated as of **Release 2.0**. It is currently kept for backward-compatibility and will be removed in the future. Please use the spoke gateway resource instead. If this is already in the state, please remove it from the state file and import as `aviatrix.AviatrixSpokeGateway`.
    /// 
    /// ## Import
    /// 
    /// Instance spoke_vpc can be imported using the gw_name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc test gw_name
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc")]
    public partial class AviatrixSpokeVpc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Cloud instance ID.
        /// </summary>
        [Output("cloudInstanceId")]
        public Output<string> CloudInstanceId { get; private set; } = null!;

        /// <summary>
        /// Type of cloud service provider. AWS=1, GCP=4, ARM=8.
        /// </summary>
        [Output("cloudType")]
        public Output<int> CloudType { get; private set; } = null!;

        /// <summary>
        /// Specify whether enabling NAT feature on the gateway or not.
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
        /// HA Subnet. Required for enabling HA for AWS/ARM gateways. Setting to empty/unset will disable HA. Setting to a valid subnet (Example: 10.12.0.0/24) will create an HA gateway on the subnet.
        /// </summary>
        [Output("haSubnet")]
        public Output<string?> HaSubnet { get; private set; } = null!;

        /// <summary>
        /// HA Zone. Required for enabling HA for GCP gateway. Setting to empty/unset will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".
        /// </summary>
        [Output("haZone")]
        public Output<string?> HaZone { get; private set; } = null!;

        /// <summary>
        /// Set to "enabled" if this feature is desired.
        /// </summary>
        [Output("singleAzHa")]
        public Output<string?> SingleAzHa { get; private set; } = null!;

        /// <summary>
        /// Public Subnet Info. Example: AWS: "CIDRZONESubnetName", etc...
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Instance tag of cloud provider. Example: key1:value1,key002:value002, etc... Only AWS (cloud_type is "1") is supported
        /// </summary>
        [Output("tagLists")]
        public Output<ImmutableArray<string>> TagLists { get; private set; } = null!;

        /// <summary>
        /// Specify the transit Gateway.
        /// </summary>
        [Output("transitGw")]
        public Output<string?> TransitGw { get; private set; } = null!;

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider. Required if cloud_type is "1" or "4". Example: AWS: "vpc-abcd1234", etc...
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west1-b", ARM: "East US 2", etc...
        /// </summary>
        [Output("vpcReg")]
        public Output<string> VpcReg { get; private set; } = null!;

        /// <summary>
        /// Size of the gateway instance. Example: AWS: "t2.large", GCP: "f1.micro", ARM: "StandardD2", etc...
        /// </summary>
        [Output("vpcSize")]
        public Output<string> VpcSize { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixSpokeVpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixSpokeVpc(string name, AviatrixSpokeVpcArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc", name, args ?? new AviatrixSpokeVpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixSpokeVpc(string name, Input<string> id, AviatrixSpokeVpcState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixSpokeVpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixSpokeVpc Get(string name, Input<string> id, AviatrixSpokeVpcState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixSpokeVpc(name, id, state, options);
        }
    }

    public sealed class AviatrixSpokeVpcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Type of cloud service provider. AWS=1, GCP=4, ARM=8.
        /// </summary>
        [Input("cloudType", required: true)]
        public Input<int> CloudType { get; set; } = null!;

        /// <summary>
        /// Specify whether enabling NAT feature on the gateway or not.
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
        /// HA Subnet. Required for enabling HA for AWS/ARM gateways. Setting to empty/unset will disable HA. Setting to a valid subnet (Example: 10.12.0.0/24) will create an HA gateway on the subnet.
        /// </summary>
        [Input("haSubnet")]
        public Input<string>? HaSubnet { get; set; }

        /// <summary>
        /// HA Zone. Required for enabling HA for GCP gateway. Setting to empty/unset will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".
        /// </summary>
        [Input("haZone")]
        public Input<string>? HaZone { get; set; }

        /// <summary>
        /// Set to "enabled" if this feature is desired.
        /// </summary>
        [Input("singleAzHa")]
        public Input<string>? SingleAzHa { get; set; }

        /// <summary>
        /// Public Subnet Info. Example: AWS: "CIDRZONESubnetName", etc...
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// Instance tag of cloud provider. Example: key1:value1,key002:value002, etc... Only AWS (cloud_type is "1") is supported
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// Specify the transit Gateway.
        /// </summary>
        [Input("transitGw")]
        public Input<string>? TransitGw { get; set; }

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider. Required if cloud_type is "1" or "4". Example: AWS: "vpc-abcd1234", etc...
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west1-b", ARM: "East US 2", etc...
        /// </summary>
        [Input("vpcReg", required: true)]
        public Input<string> VpcReg { get; set; } = null!;

        /// <summary>
        /// Size of the gateway instance. Example: AWS: "t2.large", GCP: "f1.micro", ARM: "StandardD2", etc...
        /// </summary>
        [Input("vpcSize", required: true)]
        public Input<string> VpcSize { get; set; } = null!;

        public AviatrixSpokeVpcArgs()
        {
        }
        public static new AviatrixSpokeVpcArgs Empty => new AviatrixSpokeVpcArgs();
    }

    public sealed class AviatrixSpokeVpcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Cloud instance ID.
        /// </summary>
        [Input("cloudInstanceId")]
        public Input<string>? CloudInstanceId { get; set; }

        /// <summary>
        /// Type of cloud service provider. AWS=1, GCP=4, ARM=8.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Specify whether enabling NAT feature on the gateway or not.
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
        /// HA Subnet. Required for enabling HA for AWS/ARM gateways. Setting to empty/unset will disable HA. Setting to a valid subnet (Example: 10.12.0.0/24) will create an HA gateway on the subnet.
        /// </summary>
        [Input("haSubnet")]
        public Input<string>? HaSubnet { get; set; }

        /// <summary>
        /// HA Zone. Required for enabling HA for GCP gateway. Setting to empty/unset will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".
        /// </summary>
        [Input("haZone")]
        public Input<string>? HaZone { get; set; }

        /// <summary>
        /// Set to "enabled" if this feature is desired.
        /// </summary>
        [Input("singleAzHa")]
        public Input<string>? SingleAzHa { get; set; }

        /// <summary>
        /// Public Subnet Info. Example: AWS: "CIDRZONESubnetName", etc...
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// Instance tag of cloud provider. Example: key1:value1,key002:value002, etc... Only AWS (cloud_type is "1") is supported
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// Specify the transit Gateway.
        /// </summary>
        [Input("transitGw")]
        public Input<string>? TransitGw { get; set; }

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider. Required if cloud_type is "1" or "4". Example: AWS: "vpc-abcd1234", etc...
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west1-b", ARM: "East US 2", etc...
        /// </summary>
        [Input("vpcReg")]
        public Input<string>? VpcReg { get; set; }

        /// <summary>
        /// Size of the gateway instance. Example: AWS: "t2.large", GCP: "f1.micro", ARM: "StandardD2", etc...
        /// </summary>
        [Input("vpcSize")]
        public Input<string>? VpcSize { get; set; }

        public AviatrixSpokeVpcState()
        {
        }
        public static new AviatrixSpokeVpcState Empty => new AviatrixSpokeVpcState();
    }
}
