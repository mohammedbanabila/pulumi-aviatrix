// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixCopilotSecurityGroupManagementConfig:AviatrixCopilotSecurityGroupManagementConfig")]
    public partial class AviatrixCopilotSecurityGroupManagementConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access account name.
        /// </summary>
        [Output("accountName")]
        public Output<string?> AccountName { get; private set; } = null!;

        /// <summary>
        /// Cloud type.
        /// </summary>
        [Output("cloudType")]
        public Output<int?> CloudType { get; private set; } = null!;

        /// <summary>
        /// Switch to enable copilot security group management.
        /// </summary>
        [Output("enableCopilotSecurityGroupManagement")]
        public Output<bool> EnableCopilotSecurityGroupManagement { get; private set; } = null!;

        /// <summary>
        /// Copilot instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string?> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Copilot region. Valid for AWS and Azure.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Copilot zone. Valid for GCP.
        /// </summary>
        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixCopilotSecurityGroupManagementConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixCopilotSecurityGroupManagementConfig(string name, AviatrixCopilotSecurityGroupManagementConfigArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixCopilotSecurityGroupManagementConfig:AviatrixCopilotSecurityGroupManagementConfig", name, args ?? new AviatrixCopilotSecurityGroupManagementConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixCopilotSecurityGroupManagementConfig(string name, Input<string> id, AviatrixCopilotSecurityGroupManagementConfigState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixCopilotSecurityGroupManagementConfig:AviatrixCopilotSecurityGroupManagementConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/astipkovits/pulumi-aviatrix/raw/main/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixCopilotSecurityGroupManagementConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixCopilotSecurityGroupManagementConfig Get(string name, Input<string> id, AviatrixCopilotSecurityGroupManagementConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixCopilotSecurityGroupManagementConfig(name, id, state, options);
        }
    }

    public sealed class AviatrixCopilotSecurityGroupManagementConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access account name.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Cloud type.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Switch to enable copilot security group management.
        /// </summary>
        [Input("enableCopilotSecurityGroupManagement", required: true)]
        public Input<bool> EnableCopilotSecurityGroupManagement { get; set; } = null!;

        /// <summary>
        /// Copilot instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Copilot region. Valid for AWS and Azure.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Copilot zone. Valid for GCP.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AviatrixCopilotSecurityGroupManagementConfigArgs()
        {
        }
        public static new AviatrixCopilotSecurityGroupManagementConfigArgs Empty => new AviatrixCopilotSecurityGroupManagementConfigArgs();
    }

    public sealed class AviatrixCopilotSecurityGroupManagementConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access account name.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Cloud type.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Switch to enable copilot security group management.
        /// </summary>
        [Input("enableCopilotSecurityGroupManagement")]
        public Input<bool>? EnableCopilotSecurityGroupManagement { get; set; }

        /// <summary>
        /// Copilot instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Copilot region. Valid for AWS and Azure.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Copilot zone. Valid for GCP.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AviatrixCopilotSecurityGroupManagementConfigState()
        {
        }
        public static new AviatrixCopilotSecurityGroupManagementConfigState Empty => new AviatrixCopilotSecurityGroupManagementConfigState();
    }
}
