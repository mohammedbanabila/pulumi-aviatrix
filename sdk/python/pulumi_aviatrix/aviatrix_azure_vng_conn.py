# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixAzureVngConnArgs', 'AviatrixAzureVngConn']

@pulumi.input_type
class AviatrixAzureVngConnArgs:
    def __init__(__self__, *,
                 connection_name: pulumi.Input[str],
                 primary_gateway_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixAzureVngConn resource.
        :param pulumi.Input[str] connection_name: Connection name
        :param pulumi.Input[str] primary_gateway_name: Primary gateway name
        """
        pulumi.set(__self__, "connection_name", connection_name)
        pulumi.set(__self__, "primary_gateway_name", primary_gateway_name)

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> pulumi.Input[str]:
        """
        Connection name
        """
        return pulumi.get(self, "connection_name")

    @connection_name.setter
    def connection_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_name", value)

    @property
    @pulumi.getter(name="primaryGatewayName")
    def primary_gateway_name(self) -> pulumi.Input[str]:
        """
        Primary gateway name
        """
        return pulumi.get(self, "primary_gateway_name")

    @primary_gateway_name.setter
    def primary_gateway_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "primary_gateway_name", value)


@pulumi.input_type
class _AviatrixAzureVngConnState:
    def __init__(__self__, *,
                 attached: Optional[pulumi.Input[bool]] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 primary_gateway_name: Optional[pulumi.Input[str]] = None,
                 vng_name: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixAzureVngConn resources.
        :param pulumi.Input[bool] attached: VNG attached or not
        :param pulumi.Input[str] connection_name: Connection name
        :param pulumi.Input[str] primary_gateway_name: Primary gateway name
        :param pulumi.Input[str] vng_name: VNG name
        :param pulumi.Input[str] vpc_id: VPC ID
        """
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if connection_name is not None:
            pulumi.set(__self__, "connection_name", connection_name)
        if primary_gateway_name is not None:
            pulumi.set(__self__, "primary_gateway_name", primary_gateway_name)
        if vng_name is not None:
            pulumi.set(__self__, "vng_name", vng_name)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def attached(self) -> Optional[pulumi.Input[bool]]:
        """
        VNG attached or not
        """
        return pulumi.get(self, "attached")

    @attached.setter
    def attached(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "attached", value)

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> Optional[pulumi.Input[str]]:
        """
        Connection name
        """
        return pulumi.get(self, "connection_name")

    @connection_name.setter
    def connection_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_name", value)

    @property
    @pulumi.getter(name="primaryGatewayName")
    def primary_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        Primary gateway name
        """
        return pulumi.get(self, "primary_gateway_name")

    @primary_gateway_name.setter
    def primary_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_gateway_name", value)

    @property
    @pulumi.getter(name="vngName")
    def vng_name(self) -> Optional[pulumi.Input[str]]:
        """
        VNG name
        """
        return pulumi.get(self, "vng_name")

    @vng_name.setter
    def vng_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vng_name", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class AviatrixAzureVngConn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 primary_gateway_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixAzureVngConn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_name: Connection name
        :param pulumi.Input[str] primary_gateway_name: Primary gateway name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixAzureVngConnArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixAzureVngConn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixAzureVngConnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixAzureVngConnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 primary_gateway_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixAzureVngConnArgs.__new__(AviatrixAzureVngConnArgs)

            if connection_name is None and not opts.urn:
                raise TypeError("Missing required property 'connection_name'")
            __props__.__dict__["connection_name"] = connection_name
            if primary_gateway_name is None and not opts.urn:
                raise TypeError("Missing required property 'primary_gateway_name'")
            __props__.__dict__["primary_gateway_name"] = primary_gateway_name
            __props__.__dict__["attached"] = None
            __props__.__dict__["vng_name"] = None
            __props__.__dict__["vpc_id"] = None
        super(AviatrixAzureVngConn, __self__).__init__(
            'aviatrix:index/aviatrixAzureVngConn:AviatrixAzureVngConn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attached: Optional[pulumi.Input[bool]] = None,
            connection_name: Optional[pulumi.Input[str]] = None,
            primary_gateway_name: Optional[pulumi.Input[str]] = None,
            vng_name: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'AviatrixAzureVngConn':
        """
        Get an existing AviatrixAzureVngConn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] attached: VNG attached or not
        :param pulumi.Input[str] connection_name: Connection name
        :param pulumi.Input[str] primary_gateway_name: Primary gateway name
        :param pulumi.Input[str] vng_name: VNG name
        :param pulumi.Input[str] vpc_id: VPC ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixAzureVngConnState.__new__(_AviatrixAzureVngConnState)

        __props__.__dict__["attached"] = attached
        __props__.__dict__["connection_name"] = connection_name
        __props__.__dict__["primary_gateway_name"] = primary_gateway_name
        __props__.__dict__["vng_name"] = vng_name
        __props__.__dict__["vpc_id"] = vpc_id
        return AviatrixAzureVngConn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attached(self) -> pulumi.Output[bool]:
        """
        VNG attached or not
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> pulumi.Output[str]:
        """
        Connection name
        """
        return pulumi.get(self, "connection_name")

    @property
    @pulumi.getter(name="primaryGatewayName")
    def primary_gateway_name(self) -> pulumi.Output[str]:
        """
        Primary gateway name
        """
        return pulumi.get(self, "primary_gateway_name")

    @property
    @pulumi.getter(name="vngName")
    def vng_name(self) -> pulumi.Output[str]:
        """
        VNG name
        """
        return pulumi.get(self, "vng_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        VPC ID
        """
        return pulumi.get(self, "vpc_id")

