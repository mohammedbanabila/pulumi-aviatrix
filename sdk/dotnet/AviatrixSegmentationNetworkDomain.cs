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
    /// The **aviatrix_segmentation_network_domain** resource handles creation of [Transit Segmentation](https://docs.aviatrix.com/HowTos/transit_segmentation_faq.html) Network Domains.
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
    ///     // Create an Aviatrix Segmentation Network Domain
    ///     var testSegmentationNetworkDomain = new Aviatrix.AviatrixSegmentationNetworkDomain("testSegmentationNetworkDomain", new()
    ///     {
    ///         DomainName = "domain-a",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **aviatrix_segmentation_network_domain** can be imported using the `domain_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixSegmentationNetworkDomain:AviatrixSegmentationNetworkDomain test domain_name
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixSegmentationNetworkDomain:AviatrixSegmentationNetworkDomain")]
    public partial class AviatrixSegmentationNetworkDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Network Domain.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixSegmentationNetworkDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixSegmentationNetworkDomain(string name, AviatrixSegmentationNetworkDomainArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSegmentationNetworkDomain:AviatrixSegmentationNetworkDomain", name, args ?? new AviatrixSegmentationNetworkDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixSegmentationNetworkDomain(string name, Input<string> id, AviatrixSegmentationNetworkDomainState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSegmentationNetworkDomain:AviatrixSegmentationNetworkDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixSegmentationNetworkDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixSegmentationNetworkDomain Get(string name, Input<string> id, AviatrixSegmentationNetworkDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixSegmentationNetworkDomain(name, id, state, options);
        }
    }

    public sealed class AviatrixSegmentationNetworkDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Network Domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        public AviatrixSegmentationNetworkDomainArgs()
        {
        }
        public static new AviatrixSegmentationNetworkDomainArgs Empty => new AviatrixSegmentationNetworkDomainArgs();
    }

    public sealed class AviatrixSegmentationNetworkDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Network Domain.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        public AviatrixSegmentationNetworkDomainState()
        {
        }
        public static new AviatrixSegmentationNetworkDomainState Empty => new AviatrixSegmentationNetworkDomainState();
    }
}
