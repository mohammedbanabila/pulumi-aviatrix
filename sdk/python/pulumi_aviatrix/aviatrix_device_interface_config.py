# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixDeviceInterfaceConfigArgs', 'AviatrixDeviceInterfaceConfig']

@pulumi.input_type
class AviatrixDeviceInterfaceConfigArgs:
    def __init__(__self__, *,
                 device_name: pulumi.Input[str],
                 wan_primary_interface: pulumi.Input[str],
                 wan_primary_interface_public_ip: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixDeviceInterfaceConfig resource.
        :param pulumi.Input[str] device_name: Name of the device.
        :param pulumi.Input[str] wan_primary_interface: Name of the WAN primary interface.
        :param pulumi.Input[str] wan_primary_interface_public_ip: The WAN Primary interface public IP.
        """
        pulumi.set(__self__, "device_name", device_name)
        pulumi.set(__self__, "wan_primary_interface", wan_primary_interface)
        pulumi.set(__self__, "wan_primary_interface_public_ip", wan_primary_interface_public_ip)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Input[str]:
        """
        Name of the device.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter(name="wanPrimaryInterface")
    def wan_primary_interface(self) -> pulumi.Input[str]:
        """
        Name of the WAN primary interface.
        """
        return pulumi.get(self, "wan_primary_interface")

    @wan_primary_interface.setter
    def wan_primary_interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "wan_primary_interface", value)

    @property
    @pulumi.getter(name="wanPrimaryInterfacePublicIp")
    def wan_primary_interface_public_ip(self) -> pulumi.Input[str]:
        """
        The WAN Primary interface public IP.
        """
        return pulumi.get(self, "wan_primary_interface_public_ip")

    @wan_primary_interface_public_ip.setter
    def wan_primary_interface_public_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "wan_primary_interface_public_ip", value)


@pulumi.input_type
class _AviatrixDeviceInterfaceConfigState:
    def __init__(__self__, *,
                 device_name: Optional[pulumi.Input[str]] = None,
                 wan_primary_interface: Optional[pulumi.Input[str]] = None,
                 wan_primary_interface_public_ip: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixDeviceInterfaceConfig resources.
        :param pulumi.Input[str] device_name: Name of the device.
        :param pulumi.Input[str] wan_primary_interface: Name of the WAN primary interface.
        :param pulumi.Input[str] wan_primary_interface_public_ip: The WAN Primary interface public IP.
        """
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if wan_primary_interface is not None:
            pulumi.set(__self__, "wan_primary_interface", wan_primary_interface)
        if wan_primary_interface_public_ip is not None:
            pulumi.set(__self__, "wan_primary_interface_public_ip", wan_primary_interface_public_ip)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the device.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter(name="wanPrimaryInterface")
    def wan_primary_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the WAN primary interface.
        """
        return pulumi.get(self, "wan_primary_interface")

    @wan_primary_interface.setter
    def wan_primary_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wan_primary_interface", value)

    @property
    @pulumi.getter(name="wanPrimaryInterfacePublicIp")
    def wan_primary_interface_public_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The WAN Primary interface public IP.
        """
        return pulumi.get(self, "wan_primary_interface_public_ip")

    @wan_primary_interface_public_ip.setter
    def wan_primary_interface_public_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wan_primary_interface_public_ip", value)


class AviatrixDeviceInterfaceConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 wan_primary_interface: Optional[pulumi.Input[str]] = None,
                 wan_primary_interface_public_ip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **aviatrix_device_interface_config** resource allows the configuration of the WAN primary interface and IP for a device, for use in CloudN.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Configure the primary WAN interface and IP for a device.
        test_device_interface_config = aviatrix.AviatrixDeviceInterfaceConfig("testDeviceInterfaceConfig",
            device_name="test-device",
            wan_primary_interface="eth0",
            wan_primary_interface_public_ip="181.12.43.21")
        ```

        ## Import

        **device_interface_config** can be imported using the `device_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixDeviceInterfaceConfig:AviatrixDeviceInterfaceConfig test device_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: Name of the device.
        :param pulumi.Input[str] wan_primary_interface: Name of the WAN primary interface.
        :param pulumi.Input[str] wan_primary_interface_public_ip: The WAN Primary interface public IP.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixDeviceInterfaceConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **aviatrix_device_interface_config** resource allows the configuration of the WAN primary interface and IP for a device, for use in CloudN.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Configure the primary WAN interface and IP for a device.
        test_device_interface_config = aviatrix.AviatrixDeviceInterfaceConfig("testDeviceInterfaceConfig",
            device_name="test-device",
            wan_primary_interface="eth0",
            wan_primary_interface_public_ip="181.12.43.21")
        ```

        ## Import

        **device_interface_config** can be imported using the `device_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixDeviceInterfaceConfig:AviatrixDeviceInterfaceConfig test device_name
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixDeviceInterfaceConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixDeviceInterfaceConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 wan_primary_interface: Optional[pulumi.Input[str]] = None,
                 wan_primary_interface_public_ip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixDeviceInterfaceConfigArgs.__new__(AviatrixDeviceInterfaceConfigArgs)

            if device_name is None and not opts.urn:
                raise TypeError("Missing required property 'device_name'")
            __props__.__dict__["device_name"] = device_name
            if wan_primary_interface is None and not opts.urn:
                raise TypeError("Missing required property 'wan_primary_interface'")
            __props__.__dict__["wan_primary_interface"] = wan_primary_interface
            if wan_primary_interface_public_ip is None and not opts.urn:
                raise TypeError("Missing required property 'wan_primary_interface_public_ip'")
            __props__.__dict__["wan_primary_interface_public_ip"] = wan_primary_interface_public_ip
        super(AviatrixDeviceInterfaceConfig, __self__).__init__(
            'aviatrix:index/aviatrixDeviceInterfaceConfig:AviatrixDeviceInterfaceConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_name: Optional[pulumi.Input[str]] = None,
            wan_primary_interface: Optional[pulumi.Input[str]] = None,
            wan_primary_interface_public_ip: Optional[pulumi.Input[str]] = None) -> 'AviatrixDeviceInterfaceConfig':
        """
        Get an existing AviatrixDeviceInterfaceConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: Name of the device.
        :param pulumi.Input[str] wan_primary_interface: Name of the WAN primary interface.
        :param pulumi.Input[str] wan_primary_interface_public_ip: The WAN Primary interface public IP.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixDeviceInterfaceConfigState.__new__(_AviatrixDeviceInterfaceConfigState)

        __props__.__dict__["device_name"] = device_name
        __props__.__dict__["wan_primary_interface"] = wan_primary_interface
        __props__.__dict__["wan_primary_interface_public_ip"] = wan_primary_interface_public_ip
        return AviatrixDeviceInterfaceConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Output[str]:
        """
        Name of the device.
        """
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter(name="wanPrimaryInterface")
    def wan_primary_interface(self) -> pulumi.Output[str]:
        """
        Name of the WAN primary interface.
        """
        return pulumi.get(self, "wan_primary_interface")

    @property
    @pulumi.getter(name="wanPrimaryInterfacePublicIp")
    def wan_primary_interface_public_ip(self) -> pulumi.Output[str]:
        """
        The WAN Primary interface public IP.
        """
        return pulumi.get(self, "wan_primary_interface_public_ip")

