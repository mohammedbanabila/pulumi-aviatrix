// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Inputs
{

    public sealed class AviatrixFqdnDomainNameArgs : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("fqdn", required: true)]
        public Input<string> Fqdn { get; set; } = null!;

        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        [Input("proto", required: true)]
        public Input<string> Proto { get; set; } = null!;

        public AviatrixFqdnDomainNameArgs()
        {
        }
        public static new AviatrixFqdnDomainNameArgs Empty => new AviatrixFqdnDomainNameArgs();
    }
}
