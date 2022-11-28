// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Inputs
{

    public sealed class AviatrixPrivateModeLbProxyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID of the proxy.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// VPC ID of the proxy.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public AviatrixPrivateModeLbProxyArgs()
        {
        }
        public static new AviatrixPrivateModeLbProxyArgs Empty => new AviatrixPrivateModeLbProxyArgs();
    }
}
