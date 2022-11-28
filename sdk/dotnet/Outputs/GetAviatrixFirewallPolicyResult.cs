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
    public sealed class GetAviatrixFirewallPolicyResult
    {
        public readonly string Action;
        public readonly string Description;
        /// <summary>
        /// CIDRs separated by a comma or tag names such 'HR' or 'marketing' etc.
        /// </summary>
        public readonly string DstIp;
        /// <summary>
        /// Indicates whether logging is enabled or not.
        /// * `description`- Policy description.
        /// </summary>
        public readonly bool LogEnabled;
        /// <summary>
        /// A single port or a range of port numbers.
        /// * `action`- `allow`, `deny` or `force-drop`(allow immediate packet dropping on established sessions).
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// `all`, `tcp`, `udp`, `icmp`, `sctp`, `rdp` or `dccp`.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// CIDRs separated by a comma or tag names such 'HR' or 'marketing' etc.
        /// </summary>
        public readonly string SrcIp;

        [OutputConstructor]
        private GetAviatrixFirewallPolicyResult(
            string action,

            string description,

            string dstIp,

            bool logEnabled,

            string port,

            string protocol,

            string srcIp)
        {
            Action = action;
            Description = description;
            DstIp = dstIp;
            LogEnabled = logEnabled;
            Port = port;
            Protocol = protocol;
            SrcIp = srcIp;
        }
    }
}
