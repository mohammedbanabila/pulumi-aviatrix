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
    /// The **aviatrix_azure_spoke_native_peering** resource allows the creation and management of Aviatrix-created Azure Spoke VNet attachments via Native Peering.
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
    ///     // Create an Aviatrix Azure spoke native peering
    ///     var test = new Aviatrix.AviatrixAzureSpokeNativePeering("test", new()
    ///     {
    ///         SpokeAccountName = "devops-azure",
    ///         SpokeRegion = "West US",
    ///         SpokeVpcId = "Foo_VNet:Bar_RG:GUID",
    ///         TransitGatewayName = "transit-gw-azure",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **azure_spoke_native_peering** can be imported using the `transit_gateway_name`, `spoke_account_name` and `spoke_vpc_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering test transit_gateway_name~spoke_account_name~spoke_vpc_id
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering")]
    public partial class AviatrixAzureSpokeNativePeering : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An Aviatrix account that corresponds to a subscription in Azure.
        /// </summary>
        [Output("spokeAccountName")]
        public Output<string> SpokeAccountName { get; private set; } = null!;

        /// <summary>
        /// Spoke VNet region. Example: "West US".
        /// </summary>
        [Output("spokeRegion")]
        public Output<string> SpokeRegion { get; private set; } = null!;

        /// <summary>
        /// Combination of the Spoke's VNet name, resource group and GUID. Example: "Foo_VNet:Bar_RG:GUID".
        /// </summary>
        [Output("spokeVpcId")]
        public Output<string> SpokeVpcId { get; private set; } = null!;

        /// <summary>
        /// Name of an Transit FireNet-enabled Azure transit gateway.
        /// </summary>
        [Output("transitGatewayName")]
        public Output<string> TransitGatewayName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAzureSpokeNativePeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAzureSpokeNativePeering(string name, AviatrixAzureSpokeNativePeeringArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering", name, args ?? new AviatrixAzureSpokeNativePeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAzureSpokeNativePeering(string name, Input<string> id, AviatrixAzureSpokeNativePeeringState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAzureSpokeNativePeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAzureSpokeNativePeering Get(string name, Input<string> id, AviatrixAzureSpokeNativePeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAzureSpokeNativePeering(name, id, state, options);
        }
    }

    public sealed class AviatrixAzureSpokeNativePeeringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An Aviatrix account that corresponds to a subscription in Azure.
        /// </summary>
        [Input("spokeAccountName", required: true)]
        public Input<string> SpokeAccountName { get; set; } = null!;

        /// <summary>
        /// Spoke VNet region. Example: "West US".
        /// </summary>
        [Input("spokeRegion", required: true)]
        public Input<string> SpokeRegion { get; set; } = null!;

        /// <summary>
        /// Combination of the Spoke's VNet name, resource group and GUID. Example: "Foo_VNet:Bar_RG:GUID".
        /// </summary>
        [Input("spokeVpcId", required: true)]
        public Input<string> SpokeVpcId { get; set; } = null!;

        /// <summary>
        /// Name of an Transit FireNet-enabled Azure transit gateway.
        /// </summary>
        [Input("transitGatewayName", required: true)]
        public Input<string> TransitGatewayName { get; set; } = null!;

        public AviatrixAzureSpokeNativePeeringArgs()
        {
        }
        public static new AviatrixAzureSpokeNativePeeringArgs Empty => new AviatrixAzureSpokeNativePeeringArgs();
    }

    public sealed class AviatrixAzureSpokeNativePeeringState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An Aviatrix account that corresponds to a subscription in Azure.
        /// </summary>
        [Input("spokeAccountName")]
        public Input<string>? SpokeAccountName { get; set; }

        /// <summary>
        /// Spoke VNet region. Example: "West US".
        /// </summary>
        [Input("spokeRegion")]
        public Input<string>? SpokeRegion { get; set; }

        /// <summary>
        /// Combination of the Spoke's VNet name, resource group and GUID. Example: "Foo_VNet:Bar_RG:GUID".
        /// </summary>
        [Input("spokeVpcId")]
        public Input<string>? SpokeVpcId { get; set; }

        /// <summary>
        /// Name of an Transit FireNet-enabled Azure transit gateway.
        /// </summary>
        [Input("transitGatewayName")]
        public Input<string>? TransitGatewayName { get; set; }

        public AviatrixAzureSpokeNativePeeringState()
        {
        }
        public static new AviatrixAzureSpokeNativePeeringState Empty => new AviatrixAzureSpokeNativePeeringState();
    }
}
