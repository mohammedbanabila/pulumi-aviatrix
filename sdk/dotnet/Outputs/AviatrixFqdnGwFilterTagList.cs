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
    public sealed class AviatrixFqdnGwFilterTagList
    {
        public readonly string GwName;
        public readonly ImmutableArray<string> SourceIpLists;

        [OutputConstructor]
        private AviatrixFqdnGwFilterTagList(
            string gwName,

            ImmutableArray<string> sourceIpLists)
        {
            GwName = gwName;
            SourceIpLists = sourceIpLists;
        }
    }
}
