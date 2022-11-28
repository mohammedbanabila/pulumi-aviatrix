// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Outputs
{

    [OutputType]
    public sealed class AviatrixVpcPrivateSubnet
    {
        /// <summary>
        /// CIDR block.
        /// </summary>
        public readonly string? Cidr;
        /// <summary>
        /// Name of this subnet.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// ID of this subnet.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private AviatrixVpcPrivateSubnet(
            string? cidr,

            string? name,

            string? subnetId)
        {
            Cidr = cidr;
            Name = name;
            SubnetId = subnetId;
        }
    }
}
