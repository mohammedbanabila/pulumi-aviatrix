// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    public static class GetAviatrixTransitGateways
    {
        /// <summary>
        /// The **aviatrix_transit_gateways** data source provides details about all transit gateways created by the Aviatrix Controller.
        /// </summary>
        public static Task<GetAviatrixTransitGatewaysResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAviatrixTransitGatewaysResult>("aviatrix:index/getAviatrixTransitGateways:getAviatrixTransitGateways", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAviatrixTransitGatewaysResult
    {
        /// <summary>
        /// The list of all transit gateways
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAviatrixTransitGatewaysGatewayListResult> GatewayLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAviatrixTransitGatewaysResult(
            ImmutableArray<Outputs.GetAviatrixTransitGatewaysGatewayListResult> gatewayLists,

            string id)
        {
            GatewayLists = gatewayLists;
            Id = id;
        }
    }
}
