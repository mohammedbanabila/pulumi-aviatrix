// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    public static class GetAviatrixFirenetFirewallManager
    {
        /// <summary>
        /// Use this data source to do 'save' or 'sync' for Aviatrix FireNet firewall manager.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aviatrix = Pulumi.Aviatrix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aviatrix.GetAviatrixFirenetFirewallManager.Invoke(new()
        ///     {
        ///         GatewayName = "transit",
        ///         Password = "password",
        ///         PublicIp = "1.2.3.4",
        ///         RouteTable = "router",
        ///         Save = true,
        ///         Template = "template",
        ///         TemplateStack = "templatestack",
        ///         Username = "admin-api",
        ///         VendorType = "Palo Alto Networks Panorama",
        ///         VpcId = "vpc-abcd123",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAviatrixFirenetFirewallManagerResult> InvokeAsync(GetAviatrixFirenetFirewallManagerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAviatrixFirenetFirewallManagerResult>("aviatrix:index/getAviatrixFirenetFirewallManager:getAviatrixFirenetFirewallManager", args ?? new GetAviatrixFirenetFirewallManagerArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to do 'save' or 'sync' for Aviatrix FireNet firewall manager.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aviatrix = Pulumi.Aviatrix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aviatrix.GetAviatrixFirenetFirewallManager.Invoke(new()
        ///     {
        ///         GatewayName = "transit",
        ///         Password = "password",
        ///         PublicIp = "1.2.3.4",
        ///         RouteTable = "router",
        ///         Save = true,
        ///         Template = "template",
        ///         TemplateStack = "templatestack",
        ///         Username = "admin-api",
        ///         VendorType = "Palo Alto Networks Panorama",
        ///         VpcId = "vpc-abcd123",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAviatrixFirenetFirewallManagerResult> Invoke(GetAviatrixFirenetFirewallManagerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAviatrixFirenetFirewallManagerResult>("aviatrix:index/getAviatrixFirenetFirewallManager:getAviatrixFirenetFirewallManager", args ?? new GetAviatrixFirenetFirewallManagerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAviatrixFirenetFirewallManagerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The FireNet gateway name.
        /// </summary>
        [Input("gatewayName", required: true)]
        public string GatewayName { get; set; } = null!;

        /// <summary>
        /// Number of retries for `save` or `synchronize`. Example: 1. Default value: 0.
        /// </summary>
        [Input("numberOfRetries")]
        public int? NumberOfRetries { get; set; }

        [Input("password")]
        private string? _password;

        /// <summary>
        /// Panorama login password for API calls. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        public string? Password
        {
            get => _password;
            set => _password = value;
        }

        /// <summary>
        /// The public IP address of the Panorama instance. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("publicIp")]
        public string? PublicIp { get; set; }

        /// <summary>
        /// Retry interval in seconds for `save` or `synchronize`. Example: 120. Default value: 300.
        /// </summary>
        [Input("retryInterval")]
        public int? RetryInterval { get; set; }

        /// <summary>
        /// The name of firewall virtual router to program. If left unspecified, the Controller programs the Panorama template’s first router.
        /// </summary>
        [Input("routeTable")]
        public string? RouteTable { get; set; }

        /// <summary>
        /// Switch to save or not.
        /// </summary>
        [Input("save")]
        public bool? Save { get; set; }

        /// <summary>
        /// Switch to sync or not.
        /// </summary>
        [Input("synchronize")]
        public bool? Synchronize { get; set; }

        /// <summary>
        /// Panorama template for each FireNet gateway. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("template")]
        public string? Template { get; set; }

        /// <summary>
        /// Panorama template stack for each FireNet gateway. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("templateStack")]
        public string? TemplateStack { get; set; }

        /// <summary>
        /// Panorama login name for API calls from the Controller. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("username")]
        public string? Username { get; set; }

        /// <summary>
        /// Vendor type. Valid values: "Generic" and "Palo Alto Networks Panorama".
        /// </summary>
        [Input("vendorType", required: true)]
        public string VendorType { get; set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public string VpcId { get; set; } = null!;

        public GetAviatrixFirenetFirewallManagerArgs()
        {
        }
        public static new GetAviatrixFirenetFirewallManagerArgs Empty => new GetAviatrixFirenetFirewallManagerArgs();
    }

    public sealed class GetAviatrixFirenetFirewallManagerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The FireNet gateway name.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// Number of retries for `save` or `synchronize`. Example: 1. Default value: 0.
        /// </summary>
        [Input("numberOfRetries")]
        public Input<int>? NumberOfRetries { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Panorama login password for API calls. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The public IP address of the Panorama instance. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// Retry interval in seconds for `save` or `synchronize`. Example: 120. Default value: 300.
        /// </summary>
        [Input("retryInterval")]
        public Input<int>? RetryInterval { get; set; }

        /// <summary>
        /// The name of firewall virtual router to program. If left unspecified, the Controller programs the Panorama template’s first router.
        /// </summary>
        [Input("routeTable")]
        public Input<string>? RouteTable { get; set; }

        /// <summary>
        /// Switch to save or not.
        /// </summary>
        [Input("save")]
        public Input<bool>? Save { get; set; }

        /// <summary>
        /// Switch to sync or not.
        /// </summary>
        [Input("synchronize")]
        public Input<bool>? Synchronize { get; set; }

        /// <summary>
        /// Panorama template for each FireNet gateway. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// Panorama template stack for each FireNet gateway. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("templateStack")]
        public Input<string>? TemplateStack { get; set; }

        /// <summary>
        /// Panorama login name for API calls from the Controller. Required for vendor type "Palo Alto Networks Panorama".
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Vendor type. Valid values: "Generic" and "Palo Alto Networks Panorama".
        /// </summary>
        [Input("vendorType", required: true)]
        public Input<string> VendorType { get; set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public GetAviatrixFirenetFirewallManagerInvokeArgs()
        {
        }
        public static new GetAviatrixFirenetFirewallManagerInvokeArgs Empty => new GetAviatrixFirenetFirewallManagerInvokeArgs();
    }


    [OutputType]
    public sealed class GetAviatrixFirenetFirewallManagerResult
    {
        public readonly string GatewayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? NumberOfRetries;
        public readonly string? Password;
        public readonly string? PublicIp;
        public readonly int? RetryInterval;
        public readonly string? RouteTable;
        public readonly bool? Save;
        public readonly bool? Synchronize;
        public readonly string? Template;
        public readonly string? TemplateStack;
        public readonly string? Username;
        public readonly string VendorType;
        public readonly string VpcId;

        [OutputConstructor]
        private GetAviatrixFirenetFirewallManagerResult(
            string gatewayName,

            string id,

            int? numberOfRetries,

            string? password,

            string? publicIp,

            int? retryInterval,

            string? routeTable,

            bool? save,

            bool? synchronize,

            string? template,

            string? templateStack,

            string? username,

            string vendorType,

            string vpcId)
        {
            GatewayName = gatewayName;
            Id = id;
            NumberOfRetries = numberOfRetries;
            Password = password;
            PublicIp = publicIp;
            RetryInterval = retryInterval;
            RouteTable = routeTable;
            Save = save;
            Synchronize = synchronize;
            Template = template;
            TemplateStack = templateStack;
            Username = username;
            VendorType = vendorType;
            VpcId = vpcId;
        }
    }
}
