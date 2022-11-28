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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an Aviatrix Stateful Firewall Policy
    ///     var testFirewallPolicy = new Aviatrix.AviatrixFirewallPolicy("testFirewallPolicy", new()
    ///     {
    ///         GwName = aviatrix_firewall.Test_firewall.Gw_name,
    ///         SrcIp = "10.15.0.224/32",
    ///         DstIp = "10.12.0.172/32",
    ///         Protocol = "tcp",
    ///         Port = "0:65535",
    ///         Action = "allow",
    ///         LogEnabled = true,
    ///         Description = "Test policy.",
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
    ///     // Create an Aviatrix Stateful Firewall Policy and insert it to a specific position
    ///     var testFirewallPolicy = new Aviatrix.AviatrixFirewallPolicy("testFirewallPolicy", new()
    ///     {
    ///         GwName = aviatrix_firewall.Test_firewall.Gw_name,
    ///         SrcIp = "10.15.0.225/32",
    ///         DstIp = "10.12.0.173/32",
    ///         Protocol = "tcp",
    ///         Port = "0:65535",
    ///         Action = "allow",
    ///         LogEnabled = true,
    ///         Description = "Test policy.",
    ///         Position = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **firewall_policy** can be imported using the `gw_name`, `src_ip`, `dst_ip`, `protocol`, `port` and `action` separated by `~`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixFirewallPolicy:AviatrixFirewallPolicy test "gw_name~src_ip~dst_ip~protocol~port~action"
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixFirewallPolicy:AviatrixFirewallPolicy")]
    public partial class AviatrixFirewallPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Valid values: "allow", "deny" and "force-drop" (in stateful firewall rule to allow immediate packet dropping on established sessions).
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Description of the policy. Example: "This is policy no.1".
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        [Output("dstIp")]
        public Output<string> DstIp { get; private set; } = null!;

        /// <summary>
        /// Gateway name to attach firewall policy to.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Valid values: true, false. Default value: false.
        /// </summary>
        [Output("logEnabled")]
        public Output<bool?> LogEnabled { get; private set; } = null!;

        /// <summary>
        /// A single port or a range of port numbers. Example: "25", "25:1024".
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// Position in the policy list, where the firewall policy will be inserted to. Valid values: any positive integer. Example: 2. If it is larger than the size of policy list, the policy will be inserted to the end.
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// : "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp".
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        [Output("srcIp")]
        public Output<string> SrcIp { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixFirewallPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixFirewallPolicy(string name, AviatrixFirewallPolicyArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewallPolicy:AviatrixFirewallPolicy", name, args ?? new AviatrixFirewallPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixFirewallPolicy(string name, Input<string> id, AviatrixFirewallPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewallPolicy:AviatrixFirewallPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixFirewallPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixFirewallPolicy Get(string name, Input<string> id, AviatrixFirewallPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixFirewallPolicy(name, id, state, options);
        }
    }

    public sealed class AviatrixFirewallPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid values: "allow", "deny" and "force-drop" (in stateful firewall rule to allow immediate packet dropping on established sessions).
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// Description of the policy. Example: "This is policy no.1".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        [Input("dstIp", required: true)]
        public Input<string> DstIp { get; set; } = null!;

        /// <summary>
        /// Gateway name to attach firewall policy to.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// Valid values: true, false. Default value: false.
        /// </summary>
        [Input("logEnabled")]
        public Input<bool>? LogEnabled { get; set; }

        /// <summary>
        /// A single port or a range of port numbers. Example: "25", "25:1024".
        /// </summary>
        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        /// <summary>
        /// Position in the policy list, where the firewall policy will be inserted to. Valid values: any positive integer. Example: 2. If it is larger than the size of policy list, the policy will be inserted to the end.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// : "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp".
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        [Input("srcIp", required: true)]
        public Input<string> SrcIp { get; set; } = null!;

        public AviatrixFirewallPolicyArgs()
        {
        }
        public static new AviatrixFirewallPolicyArgs Empty => new AviatrixFirewallPolicyArgs();
    }

    public sealed class AviatrixFirewallPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid values: "allow", "deny" and "force-drop" (in stateful firewall rule to allow immediate packet dropping on established sessions).
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Description of the policy. Example: "This is policy no.1".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        [Input("dstIp")]
        public Input<string>? DstIp { get; set; }

        /// <summary>
        /// Gateway name to attach firewall policy to.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// Valid values: true, false. Default value: false.
        /// </summary>
        [Input("logEnabled")]
        public Input<bool>? LogEnabled { get; set; }

        /// <summary>
        /// A single port or a range of port numbers. Example: "25", "25:1024".
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Position in the policy list, where the firewall policy will be inserted to. Valid values: any positive integer. Example: 2. If it is larger than the size of policy list, the policy will be inserted to the end.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// : "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp".
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        [Input("srcIp")]
        public Input<string>? SrcIp { get; set; }

        public AviatrixFirewallPolicyState()
        {
        }
        public static new AviatrixFirewallPolicyState Empty => new AviatrixFirewallPolicyState();
    }
}
