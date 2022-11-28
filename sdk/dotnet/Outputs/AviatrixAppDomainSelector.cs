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
    public sealed class AviatrixAppDomainSelector
    {
        /// <summary>
        /// List of match expressions. The App Domain will be a union of all resources matched by each `match_expressions`.`match_expressions` blocks cannot be empty.
        /// </summary>
        public readonly ImmutableArray<Outputs.AviatrixAppDomainSelectorMatchExpression> MatchExpressions;

        [OutputConstructor]
        private AviatrixAppDomainSelector(ImmutableArray<Outputs.AviatrixAppDomainSelectorMatchExpression> matchExpressions)
        {
            MatchExpressions = matchExpressions;
        }
    }
}
