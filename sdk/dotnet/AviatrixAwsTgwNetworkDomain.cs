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
    /// The **aviatrix_aws_tgw_network_domain** resource allows the creation and management of Aviatrix network domains.
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
    ///     // Create an Aviatrix AWS TGW network domain
    ///     var testAwsTgw = new Aviatrix.AviatrixAwsTgw("testAwsTgw", new()
    ///     {
    ///         AccountName = "devops",
    ///         AwsSideAsNumber = "64512",
    ///         Region = "us-east-1",
    ///         TgwName = "test-AWS-TGW",
    ///         ManageSecurityDomain = false,
    ///         ManageVpcAttachment = false,
    ///         ManageTransitGatewayAttachment = false,
    ///     });
    /// 
    ///     var defaultDomain = new Aviatrix.AviatrixAwsTgwNetworkDomain("defaultDomain", new()
    ///     {
    ///         TgwName = testAwsTgw.TgwName,
    ///     });
    /// 
    ///     var sharedServiceDomain = new Aviatrix.AviatrixAwsTgwNetworkDomain("sharedServiceDomain", new()
    ///     {
    ///         TgwName = testAwsTgw.TgwName,
    ///     });
    /// 
    ///     var aviatrixEdgeDomain = new Aviatrix.AviatrixAwsTgwNetworkDomain("aviatrixEdgeDomain", new()
    ///     {
    ///         TgwName = testAwsTgw.TgwName,
    ///     });
    /// 
    ///     var defaultSdConn1 = new Aviatrix.AviatrixAwsTgwSecurityDomainConn("defaultSdConn1", new()
    ///     {
    ///         TgwName = testAwsTgw.TgwName,
    ///         DomainName1 = aviatrixEdgeDomain.Name,
    ///         DomainName2 = defaultDomain.Name,
    ///     });
    /// 
    ///     var defaultSdConn2 = new Aviatrix.AviatrixAwsTgwSecurityDomainConn("defaultSdConn2", new()
    ///     {
    ///         TgwName = testAwsTgw.TgwName,
    ///         DomainName1 = aviatrixEdgeDomain.Name,
    ///         DomainName2 = sharedServiceDomain.Name,
    ///     });
    /// 
    ///     var defaultSdConn3 = new Aviatrix.AviatrixAwsTgwSecurityDomainConn("defaultSdConn3", new()
    ///     {
    ///         TgwName = testAwsTgw.TgwName,
    ///         DomainName1 = defaultDomain.Name,
    ///         DomainName2 = sharedServiceDomain.Name,
    ///     });
    /// 
    ///     var test = new Aviatrix.AviatrixAwsTgwNetworkDomain("test", new()
    ///     {
    ///         TgwName = testAwsTgw.TgwName,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             defaultDomain,
    ///             sharedServiceDomain,
    ///             aviatrixEdgeDomain,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **aws_tgw_network_domain** can be imported using the `name` and `tgw_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain test tgw_name~name
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain")]
    public partial class AviatrixAwsTgwNetworkDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
        /// </summary>
        [Output("aviatrixFirewall")]
        public Output<bool?> AviatrixFirewall { get; private set; } = null!;

        /// <summary>
        /// The name of the network domain.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
        /// </summary>
        [Output("nativeEgress")]
        public Output<bool?> NativeEgress { get; private set; } = null!;

        /// <summary>
        /// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
        /// </summary>
        [Output("nativeFirewall")]
        public Output<bool?> NativeFirewall { get; private set; } = null!;

        /// <summary>
        /// The AWS TGW name of the network domain.
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwNetworkDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwNetworkDomain(string name, AviatrixAwsTgwNetworkDomainArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain", name, args ?? new AviatrixAwsTgwNetworkDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwNetworkDomain(string name, Input<string> id, AviatrixAwsTgwNetworkDomainState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAwsTgwNetworkDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwNetworkDomain Get(string name, Input<string> id, AviatrixAwsTgwNetworkDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwNetworkDomain(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwNetworkDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("aviatrixFirewall")]
        public Input<bool>? AviatrixFirewall { get; set; }

        /// <summary>
        /// The name of the network domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
        /// </summary>
        [Input("nativeEgress")]
        public Input<bool>? NativeEgress { get; set; }

        /// <summary>
        /// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
        /// </summary>
        [Input("nativeFirewall")]
        public Input<bool>? NativeFirewall { get; set; }

        /// <summary>
        /// The AWS TGW name of the network domain.
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        public AviatrixAwsTgwNetworkDomainArgs()
        {
        }
        public static new AviatrixAwsTgwNetworkDomainArgs Empty => new AviatrixAwsTgwNetworkDomainArgs();
    }

    public sealed class AviatrixAwsTgwNetworkDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("aviatrixFirewall")]
        public Input<bool>? AviatrixFirewall { get; set; }

        /// <summary>
        /// The name of the network domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.
        /// </summary>
        [Input("nativeEgress")]
        public Input<bool>? NativeEgress { get; set; }

        /// <summary>
        /// Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.
        /// </summary>
        [Input("nativeFirewall")]
        public Input<bool>? NativeFirewall { get; set; }

        /// <summary>
        /// The AWS TGW name of the network domain.
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        public AviatrixAwsTgwNetworkDomainState()
        {
        }
        public static new AviatrixAwsTgwNetworkDomainState Empty => new AviatrixAwsTgwNetworkDomainState();
    }
}
