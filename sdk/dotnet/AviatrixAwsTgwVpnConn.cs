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
    /// The **aviatrix_aws_tgw_vpn_conn** resource allows the creation and management of Aviatrix AWS TGW VPN connections in their selected Security Domain.
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
    ///     // Create an Aviatrix AWS TGW VPN Connection (dynamic)
    ///     var testAwsTgwVpnConn = new Aviatrix.AviatrixAwsTgwVpnConn("testAwsTgwVpnConn", new()
    ///     {
    ///         ConnectionName = "my-conn1",
    ///         ConnectionType = "dynamic",
    ///         PublicIp = "40.0.0.0",
    ///         RemoteAsNumber = "12",
    ///         RouteDomainName = "Default_Domain",
    ///         TgwName = "test-tgw1",
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
    ///     // Create an Aviatrix AWS TGW VPN Connection (static)
    ///     var testAwsTgwVpnConn = new Aviatrix.AviatrixAwsTgwVpnConn("testAwsTgwVpnConn", new()
    ///     {
    ///         ConnectionName = "my-conn1",
    ///         ConnectionType = "static",
    ///         PublicIp = "40.0.0.0",
    ///         RemoteCidr = "16.0.0.0/16,16.1.0.0/16",
    ///         RouteDomainName = "Default_Domain",
    ///         TgwName = "test-tgw1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **aws_tgw_vpn_conn** can be imported using the `tgw_name` and `vpn_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn test tgw_name~vpn_id
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn")]
    public partial class AviatrixAwsTgwVpnConn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unique name of the connection.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a static VPN connection. Default value: 'dynamic'.
        /// </summary>
        [Output("connectionType")]
        public Output<string?> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// Enable Global Acceleration. Type: Boolean. Default: false.
        /// </summary>
        [Output("enableGlobalAcceleration")]
        public Output<bool?> EnableGlobalAcceleration { get; private set; } = null!;

        /// <summary>
        /// Switch to enable/disable [encrypted transit approval](https://docs.aviatrix.com/HowTos/tgw_approval.html) for AWS TGW VPN connection. Valid values: true, false. Default value: false.
        /// </summary>
        [Output("enableLearnedCidrsApproval")]
        public Output<bool?> EnableLearnedCidrsApproval { get; private set; } = null!;

        /// <summary>
        /// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
        /// </summary>
        [Output("insideIpCidrTun1")]
        public Output<string?> InsideIpCidrTun1 { get; private set; } = null!;

        /// <summary>
        /// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
        /// </summary>
        [Output("insideIpCidrTun2")]
        public Output<string?> InsideIpCidrTun2 { get; private set; } = null!;

        /// <summary>
        /// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.
        /// </summary>
        [Output("preSharedKeyTun1")]
        public Output<string?> PreSharedKeyTun1 { get; private set; } = null!;

        /// <summary>
        /// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.
        /// </summary>
        [Output("preSharedKeyTun2")]
        public Output<string?> PreSharedKeyTun2 { get; private set; } = null!;

        /// <summary>
        /// Public IP address. Example: "40.0.0.0".
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// AWS side as a number. Integer between 1-4294967294. Example: "12". **Required for a dynamic VPN connection.**
        /// </summary>
        [Output("remoteAsNumber")]
        public Output<string?> RemoteAsNumber { get; private set; } = null!;

        /// <summary>
        /// Remote CIDRs separated by ",". Example: AWS: "16.0.0.0/16,16.1.0.0/16". **Required for a static VPN connection.**
        /// </summary>
        [Output("remoteCidr")]
        public Output<string?> RemoteCidr { get; private set; } = null!;

        /// <summary>
        /// The name of a route domain, to which the vpn will be attached.
        /// </summary>
        [Output("routeDomainName")]
        public Output<string> RouteDomainName { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the name of an AWS TGW.
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;

        /// <summary>
        /// ID of the VPN generated by creation of the connection.
        /// </summary>
        [Output("vpnId")]
        public Output<string> VpnId { get; private set; } = null!;

        /// <summary>
        /// AWS TGW VPN tunnel data.
        /// </summary>
        [Output("vpnTunnelDatas")]
        public Output<ImmutableArray<Outputs.AviatrixAwsTgwVpnConnVpnTunnelData>> VpnTunnelDatas { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwVpnConn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwVpnConn(string name, AviatrixAwsTgwVpnConnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn", name, args ?? new AviatrixAwsTgwVpnConnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwVpnConn(string name, Input<string> id, AviatrixAwsTgwVpnConnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/astipkovits",
                AdditionalSecretOutputs =
                {
                    "preSharedKeyTun1",
                    "preSharedKeyTun2",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixAwsTgwVpnConn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwVpnConn Get(string name, Input<string> id, AviatrixAwsTgwVpnConnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwVpnConn(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwVpnConnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name of the connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a static VPN connection. Default value: 'dynamic'.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Enable Global Acceleration. Type: Boolean. Default: false.
        /// </summary>
        [Input("enableGlobalAcceleration")]
        public Input<bool>? EnableGlobalAcceleration { get; set; }

        /// <summary>
        /// Switch to enable/disable [encrypted transit approval](https://docs.aviatrix.com/HowTos/tgw_approval.html) for AWS TGW VPN connection. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
        /// </summary>
        [Input("insideIpCidrTun1")]
        public Input<string>? InsideIpCidrTun1 { get; set; }

        /// <summary>
        /// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
        /// </summary>
        [Input("insideIpCidrTun2")]
        public Input<string>? InsideIpCidrTun2 { get; set; }

        [Input("preSharedKeyTun1")]
        private Input<string>? _preSharedKeyTun1;

        /// <summary>
        /// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.
        /// </summary>
        public Input<string>? PreSharedKeyTun1
        {
            get => _preSharedKeyTun1;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _preSharedKeyTun1 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("preSharedKeyTun2")]
        private Input<string>? _preSharedKeyTun2;

        /// <summary>
        /// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.
        /// </summary>
        public Input<string>? PreSharedKeyTun2
        {
            get => _preSharedKeyTun2;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _preSharedKeyTun2 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Public IP address. Example: "40.0.0.0".
        /// </summary>
        [Input("publicIp", required: true)]
        public Input<string> PublicIp { get; set; } = null!;

        /// <summary>
        /// AWS side as a number. Integer between 1-4294967294. Example: "12". **Required for a dynamic VPN connection.**
        /// </summary>
        [Input("remoteAsNumber")]
        public Input<string>? RemoteAsNumber { get; set; }

        /// <summary>
        /// Remote CIDRs separated by ",". Example: AWS: "16.0.0.0/16,16.1.0.0/16". **Required for a static VPN connection.**
        /// </summary>
        [Input("remoteCidr")]
        public Input<string>? RemoteCidr { get; set; }

        /// <summary>
        /// The name of a route domain, to which the vpn will be attached.
        /// </summary>
        [Input("routeDomainName", required: true)]
        public Input<string> RouteDomainName { get; set; } = null!;

        /// <summary>
        /// This parameter represents the name of an AWS TGW.
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        public AviatrixAwsTgwVpnConnArgs()
        {
        }
        public static new AviatrixAwsTgwVpnConnArgs Empty => new AviatrixAwsTgwVpnConnArgs();
    }

    public sealed class AviatrixAwsTgwVpnConnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name of the connection.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a static VPN connection. Default value: 'dynamic'.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Enable Global Acceleration. Type: Boolean. Default: false.
        /// </summary>
        [Input("enableGlobalAcceleration")]
        public Input<bool>? EnableGlobalAcceleration { get; set; }

        /// <summary>
        /// Switch to enable/disable [encrypted transit approval](https://docs.aviatrix.com/HowTos/tgw_approval.html) for AWS TGW VPN connection. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
        /// </summary>
        [Input("insideIpCidrTun1")]
        public Input<string>? InsideIpCidrTun1 { get; set; }

        /// <summary>
        /// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
        /// </summary>
        [Input("insideIpCidrTun2")]
        public Input<string>? InsideIpCidrTun2 { get; set; }

        [Input("preSharedKeyTun1")]
        private Input<string>? _preSharedKeyTun1;

        /// <summary>
        /// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.
        /// </summary>
        public Input<string>? PreSharedKeyTun1
        {
            get => _preSharedKeyTun1;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _preSharedKeyTun1 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("preSharedKeyTun2")]
        private Input<string>? _preSharedKeyTun2;

        /// <summary>
        /// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.
        /// </summary>
        public Input<string>? PreSharedKeyTun2
        {
            get => _preSharedKeyTun2;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _preSharedKeyTun2 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Public IP address. Example: "40.0.0.0".
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// AWS side as a number. Integer between 1-4294967294. Example: "12". **Required for a dynamic VPN connection.**
        /// </summary>
        [Input("remoteAsNumber")]
        public Input<string>? RemoteAsNumber { get; set; }

        /// <summary>
        /// Remote CIDRs separated by ",". Example: AWS: "16.0.0.0/16,16.1.0.0/16". **Required for a static VPN connection.**
        /// </summary>
        [Input("remoteCidr")]
        public Input<string>? RemoteCidr { get; set; }

        /// <summary>
        /// The name of a route domain, to which the vpn will be attached.
        /// </summary>
        [Input("routeDomainName")]
        public Input<string>? RouteDomainName { get; set; }

        /// <summary>
        /// This parameter represents the name of an AWS TGW.
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        /// <summary>
        /// ID of the VPN generated by creation of the connection.
        /// </summary>
        [Input("vpnId")]
        public Input<string>? VpnId { get; set; }

        [Input("vpnTunnelDatas")]
        private InputList<Inputs.AviatrixAwsTgwVpnConnVpnTunnelDataGetArgs>? _vpnTunnelDatas;

        /// <summary>
        /// AWS TGW VPN tunnel data.
        /// </summary>
        public InputList<Inputs.AviatrixAwsTgwVpnConnVpnTunnelDataGetArgs> VpnTunnelDatas
        {
            get => _vpnTunnelDatas ?? (_vpnTunnelDatas = new InputList<Inputs.AviatrixAwsTgwVpnConnVpnTunnelDataGetArgs>());
            set => _vpnTunnelDatas = value;
        }

        public AviatrixAwsTgwVpnConnState()
        {
        }
        public static new AviatrixAwsTgwVpnConnState Empty => new AviatrixAwsTgwVpnConnState();
    }
}
