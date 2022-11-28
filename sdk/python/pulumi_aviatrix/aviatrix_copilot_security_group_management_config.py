# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixCopilotSecurityGroupManagementConfigArgs', 'AviatrixCopilotSecurityGroupManagementConfig']

@pulumi.input_type
class AviatrixCopilotSecurityGroupManagementConfigArgs:
    def __init__(__self__, *,
                 enable_copilot_security_group_management: pulumi.Input[bool],
                 account_name: Optional[pulumi.Input[str]] = None,
                 cloud_type: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixCopilotSecurityGroupManagementConfig resource.
        :param pulumi.Input[bool] enable_copilot_security_group_management: Switch to enable copilot security group management. Valid values: true, false.
        :param pulumi.Input[str] account_name: Aviatrix access account name. Required to enable copilot security group management.
        :param pulumi.Input[int] cloud_type: Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.
        :param pulumi.Input[str] instance_id: CoPilot instance ID. Required to enable copilot security group management.
        :param pulumi.Input[str] region: Region where CoPilot is deployed. Required and valid for AWS and Azure.
        :param pulumi.Input[str] vpc_id: VPC ID. Required to enable copilot security group management.
        :param pulumi.Input[str] zone: Zone where CoPilot is deployed. Required and valid for GCP.
        """
        pulumi.set(__self__, "enable_copilot_security_group_management", enable_copilot_security_group_management)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if cloud_type is not None:
            pulumi.set(__self__, "cloud_type", cloud_type)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="enableCopilotSecurityGroupManagement")
    def enable_copilot_security_group_management(self) -> pulumi.Input[bool]:
        """
        Switch to enable copilot security group management. Valid values: true, false.
        """
        return pulumi.get(self, "enable_copilot_security_group_management")

    @enable_copilot_security_group_management.setter
    def enable_copilot_security_group_management(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable_copilot_security_group_management", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Aviatrix access account name. Required to enable copilot security group management.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="cloudType")
    def cloud_type(self) -> Optional[pulumi.Input[int]]:
        """
        Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.
        """
        return pulumi.get(self, "cloud_type")

    @cloud_type.setter
    def cloud_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cloud_type", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        CoPilot instance ID. Required to enable copilot security group management.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region where CoPilot is deployed. Required and valid for AWS and Azure.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID. Required to enable copilot security group management.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Zone where CoPilot is deployed. Required and valid for GCP.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _AviatrixCopilotSecurityGroupManagementConfigState:
    def __init__(__self__, *,
                 account_name: Optional[pulumi.Input[str]] = None,
                 cloud_type: Optional[pulumi.Input[int]] = None,
                 enable_copilot_security_group_management: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixCopilotSecurityGroupManagementConfig resources.
        :param pulumi.Input[str] account_name: Aviatrix access account name. Required to enable copilot security group management.
        :param pulumi.Input[int] cloud_type: Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.
        :param pulumi.Input[bool] enable_copilot_security_group_management: Switch to enable copilot security group management. Valid values: true, false.
        :param pulumi.Input[str] instance_id: CoPilot instance ID. Required to enable copilot security group management.
        :param pulumi.Input[str] region: Region where CoPilot is deployed. Required and valid for AWS and Azure.
        :param pulumi.Input[str] vpc_id: VPC ID. Required to enable copilot security group management.
        :param pulumi.Input[str] zone: Zone where CoPilot is deployed. Required and valid for GCP.
        """
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if cloud_type is not None:
            pulumi.set(__self__, "cloud_type", cloud_type)
        if enable_copilot_security_group_management is not None:
            pulumi.set(__self__, "enable_copilot_security_group_management", enable_copilot_security_group_management)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Aviatrix access account name. Required to enable copilot security group management.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="cloudType")
    def cloud_type(self) -> Optional[pulumi.Input[int]]:
        """
        Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.
        """
        return pulumi.get(self, "cloud_type")

    @cloud_type.setter
    def cloud_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cloud_type", value)

    @property
    @pulumi.getter(name="enableCopilotSecurityGroupManagement")
    def enable_copilot_security_group_management(self) -> Optional[pulumi.Input[bool]]:
        """
        Switch to enable copilot security group management. Valid values: true, false.
        """
        return pulumi.get(self, "enable_copilot_security_group_management")

    @enable_copilot_security_group_management.setter
    def enable_copilot_security_group_management(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_copilot_security_group_management", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        CoPilot instance ID. Required to enable copilot security group management.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region where CoPilot is deployed. Required and valid for AWS and Azure.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID. Required to enable copilot security group management.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Zone where CoPilot is deployed. Required and valid for GCP.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class AviatrixCopilotSecurityGroupManagementConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 cloud_type: Optional[pulumi.Input[int]] = None,
                 enable_copilot_security_group_management: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **aviatrix_copilot_security_group_management_config** resource allows management of controller CoPilot security group management configuration. This resource is available as of provider version R2.23+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Enable the CoPilot Security Group Management
        test = aviatrix.AviatrixCopilotSecurityGroupManagementConfig("test",
            account_name="aws-account",
            cloud_type=1,
            enable_copilot_security_group_management=True,
            instance_id="i-1234567890",
            region="us-east-1",
            vpc_id="vpc-1234567890")
        ```

        ## Import

        **aviatrix_copilot_security_group_management_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13

        ```sh
         $ pulumi import aviatrix:index/aviatrixCopilotSecurityGroupManagementConfig:AviatrixCopilotSecurityGroupManagementConfig test 10-11-12-13
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Aviatrix access account name. Required to enable copilot security group management.
        :param pulumi.Input[int] cloud_type: Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.
        :param pulumi.Input[bool] enable_copilot_security_group_management: Switch to enable copilot security group management. Valid values: true, false.
        :param pulumi.Input[str] instance_id: CoPilot instance ID. Required to enable copilot security group management.
        :param pulumi.Input[str] region: Region where CoPilot is deployed. Required and valid for AWS and Azure.
        :param pulumi.Input[str] vpc_id: VPC ID. Required to enable copilot security group management.
        :param pulumi.Input[str] zone: Zone where CoPilot is deployed. Required and valid for GCP.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixCopilotSecurityGroupManagementConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **aviatrix_copilot_security_group_management_config** resource allows management of controller CoPilot security group management configuration. This resource is available as of provider version R2.23+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Enable the CoPilot Security Group Management
        test = aviatrix.AviatrixCopilotSecurityGroupManagementConfig("test",
            account_name="aws-account",
            cloud_type=1,
            enable_copilot_security_group_management=True,
            instance_id="i-1234567890",
            region="us-east-1",
            vpc_id="vpc-1234567890")
        ```

        ## Import

        **aviatrix_copilot_security_group_management_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13

        ```sh
         $ pulumi import aviatrix:index/aviatrixCopilotSecurityGroupManagementConfig:AviatrixCopilotSecurityGroupManagementConfig test 10-11-12-13
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixCopilotSecurityGroupManagementConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixCopilotSecurityGroupManagementConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 cloud_type: Optional[pulumi.Input[int]] = None,
                 enable_copilot_security_group_management: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixCopilotSecurityGroupManagementConfigArgs.__new__(AviatrixCopilotSecurityGroupManagementConfigArgs)

            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["cloud_type"] = cloud_type
            if enable_copilot_security_group_management is None and not opts.urn:
                raise TypeError("Missing required property 'enable_copilot_security_group_management'")
            __props__.__dict__["enable_copilot_security_group_management"] = enable_copilot_security_group_management
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region"] = region
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["zone"] = zone
        super(AviatrixCopilotSecurityGroupManagementConfig, __self__).__init__(
            'aviatrix:index/aviatrixCopilotSecurityGroupManagementConfig:AviatrixCopilotSecurityGroupManagementConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            cloud_type: Optional[pulumi.Input[int]] = None,
            enable_copilot_security_group_management: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'AviatrixCopilotSecurityGroupManagementConfig':
        """
        Get an existing AviatrixCopilotSecurityGroupManagementConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Aviatrix access account name. Required to enable copilot security group management.
        :param pulumi.Input[int] cloud_type: Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.
        :param pulumi.Input[bool] enable_copilot_security_group_management: Switch to enable copilot security group management. Valid values: true, false.
        :param pulumi.Input[str] instance_id: CoPilot instance ID. Required to enable copilot security group management.
        :param pulumi.Input[str] region: Region where CoPilot is deployed. Required and valid for AWS and Azure.
        :param pulumi.Input[str] vpc_id: VPC ID. Required to enable copilot security group management.
        :param pulumi.Input[str] zone: Zone where CoPilot is deployed. Required and valid for GCP.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixCopilotSecurityGroupManagementConfigState.__new__(_AviatrixCopilotSecurityGroupManagementConfigState)

        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["cloud_type"] = cloud_type
        __props__.__dict__["enable_copilot_security_group_management"] = enable_copilot_security_group_management
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["region"] = region
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["zone"] = zone
        return AviatrixCopilotSecurityGroupManagementConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[Optional[str]]:
        """
        Aviatrix access account name. Required to enable copilot security group management.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="cloudType")
    def cloud_type(self) -> pulumi.Output[Optional[int]]:
        """
        Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.
        """
        return pulumi.get(self, "cloud_type")

    @property
    @pulumi.getter(name="enableCopilotSecurityGroupManagement")
    def enable_copilot_security_group_management(self) -> pulumi.Output[bool]:
        """
        Switch to enable copilot security group management. Valid values: true, false.
        """
        return pulumi.get(self, "enable_copilot_security_group_management")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        CoPilot instance ID. Required to enable copilot security group management.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        Region where CoPilot is deployed. Required and valid for AWS and Azure.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        VPC ID. Required to enable copilot security group management.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[Optional[str]]:
        """
        Zone where CoPilot is deployed. Required and valid for GCP.
        """
        return pulumi.get(self, "zone")

