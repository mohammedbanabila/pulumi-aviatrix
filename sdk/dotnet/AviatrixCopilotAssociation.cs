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
    /// The **aviatrix_copilot_association** resource allows management of controller CoPilot Association. This resource is available as of provider version R2.19+.
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
    ///     // Create a CoPilot Association
    ///     var testCopilotAssociation = new Aviatrix.AviatrixCopilotAssociation("testCopilotAssociation", new()
    ///     {
    ///         CopilotAddress = "copilot.aviatrix.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **aviatrix_copilot_association** can be imported using controller IP, e.g. controller IP is 10.11.12.13
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixCopilotAssociation:AviatrixCopilotAssociation test_copilot_association 10-11-12-13
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixCopilotAssociation:AviatrixCopilotAssociation")]
    public partial class AviatrixCopilotAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// CoPilot instance IP Address or Hostname.
        /// </summary>
        [Output("copilotAddress")]
        public Output<string> CopilotAddress { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixCopilotAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixCopilotAssociation(string name, AviatrixCopilotAssociationArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixCopilotAssociation:AviatrixCopilotAssociation", name, args ?? new AviatrixCopilotAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixCopilotAssociation(string name, Input<string> id, AviatrixCopilotAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixCopilotAssociation:AviatrixCopilotAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixCopilotAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixCopilotAssociation Get(string name, Input<string> id, AviatrixCopilotAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixCopilotAssociation(name, id, state, options);
        }
    }

    public sealed class AviatrixCopilotAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CoPilot instance IP Address or Hostname.
        /// </summary>
        [Input("copilotAddress", required: true)]
        public Input<string> CopilotAddress { get; set; } = null!;

        public AviatrixCopilotAssociationArgs()
        {
        }
        public static new AviatrixCopilotAssociationArgs Empty => new AviatrixCopilotAssociationArgs();
    }

    public sealed class AviatrixCopilotAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CoPilot instance IP Address or Hostname.
        /// </summary>
        [Input("copilotAddress")]
        public Input<string>? CopilotAddress { get; set; }

        public AviatrixCopilotAssociationState()
        {
        }
        public static new AviatrixCopilotAssociationState Empty => new AviatrixCopilotAssociationState();
    }
}
