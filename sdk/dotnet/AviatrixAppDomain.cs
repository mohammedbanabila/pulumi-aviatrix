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
    /// !&gt; **WARNING** **aviatrix_app_domain** is part of the Micro-segmentation private preview feature for R2.22.0. If you wish to enable a private preview mode feature, please contact your sales representative or Aviatrix Support.
    /// The **aviatrix_app_domain** resource handles the creation and management of App Domains. Available as of Provider R2.22.0+.
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
    ///     // Create an Aviatrix App Domain
    ///     var testAppDomainIp = new Aviatrix.AviatrixAppDomain("testAppDomainIp", new()
    ///     {
    ///         Selector = new Aviatrix.Inputs.AviatrixAppDomainSelectorArgs
    ///         {
    ///             MatchExpressions = new[]
    ///             {
    ///                 new Aviatrix.Inputs.AviatrixAppDomainSelectorMatchExpressionArgs
    ///                 {
    ///                     AccountName = "devops",
    ///                     Region = "us-west-2",
    ///                     Tags = 
    ///                     {
    ///                         { "k3", "v3" },
    ///                     },
    ///                     Type = "vm",
    ///                 },
    ///                 new Aviatrix.Inputs.AviatrixAppDomainSelectorMatchExpressionArgs
    ///                 {
    ///                     Cidr = "10.0.0.0/16",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **aviatrix_app_domain** can be imported using the `uuid`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixAppDomain:AviatrixAppDomain test 41984f8b-5a37-4272-89b3-57c79e9ff77c
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixAppDomain:AviatrixAppDomain")]
    public partial class AviatrixAppDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the App Domain.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Block containing match expressions to filter the App Domain.
        /// </summary>
        [Output("selector")]
        public Output<Outputs.AviatrixAppDomainSelector> Selector { get; private set; } = null!;

        /// <summary>
        /// UUID of the App Domain.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAppDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAppDomain(string name, AviatrixAppDomainArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAppDomain:AviatrixAppDomain", name, args ?? new AviatrixAppDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAppDomain(string name, Input<string> id, AviatrixAppDomainState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAppDomain:AviatrixAppDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAppDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAppDomain Get(string name, Input<string> id, AviatrixAppDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAppDomain(name, id, state, options);
        }
    }

    public sealed class AviatrixAppDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the App Domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Block containing match expressions to filter the App Domain.
        /// </summary>
        [Input("selector", required: true)]
        public Input<Inputs.AviatrixAppDomainSelectorArgs> Selector { get; set; } = null!;

        public AviatrixAppDomainArgs()
        {
        }
        public static new AviatrixAppDomainArgs Empty => new AviatrixAppDomainArgs();
    }

    public sealed class AviatrixAppDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the App Domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Block containing match expressions to filter the App Domain.
        /// </summary>
        [Input("selector")]
        public Input<Inputs.AviatrixAppDomainSelectorGetArgs>? Selector { get; set; }

        /// <summary>
        /// UUID of the App Domain.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public AviatrixAppDomainState()
        {
        }
        public static new AviatrixAppDomainState Empty => new AviatrixAppDomainState();
    }
}
