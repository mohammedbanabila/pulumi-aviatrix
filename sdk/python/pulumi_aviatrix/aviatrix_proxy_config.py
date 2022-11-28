# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixProxyConfigArgs', 'AviatrixProxyConfig']

@pulumi.input_type
class AviatrixProxyConfigArgs:
    def __init__(__self__, *,
                 http_proxy: pulumi.Input[str],
                 https_proxy: pulumi.Input[str],
                 proxy_ca_certificate: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixProxyConfig resource.
        :param pulumi.Input[str] http_proxy: Http proxy URL.
        :param pulumi.Input[str] https_proxy: Https proxy URL.
        :param pulumi.Input[str] proxy_ca_certificate: Server CA Certificate file. Use the `file` function to read from a file.
        """
        pulumi.set(__self__, "http_proxy", http_proxy)
        pulumi.set(__self__, "https_proxy", https_proxy)
        if proxy_ca_certificate is not None:
            pulumi.set(__self__, "proxy_ca_certificate", proxy_ca_certificate)

    @property
    @pulumi.getter(name="httpProxy")
    def http_proxy(self) -> pulumi.Input[str]:
        """
        Http proxy URL.
        """
        return pulumi.get(self, "http_proxy")

    @http_proxy.setter
    def http_proxy(self, value: pulumi.Input[str]):
        pulumi.set(self, "http_proxy", value)

    @property
    @pulumi.getter(name="httpsProxy")
    def https_proxy(self) -> pulumi.Input[str]:
        """
        Https proxy URL.
        """
        return pulumi.get(self, "https_proxy")

    @https_proxy.setter
    def https_proxy(self, value: pulumi.Input[str]):
        pulumi.set(self, "https_proxy", value)

    @property
    @pulumi.getter(name="proxyCaCertificate")
    def proxy_ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Server CA Certificate file. Use the `file` function to read from a file.
        """
        return pulumi.get(self, "proxy_ca_certificate")

    @proxy_ca_certificate.setter
    def proxy_ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_ca_certificate", value)


@pulumi.input_type
class _AviatrixProxyConfigState:
    def __init__(__self__, *,
                 http_proxy: Optional[pulumi.Input[str]] = None,
                 https_proxy: Optional[pulumi.Input[str]] = None,
                 proxy_ca_certificate: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixProxyConfig resources.
        :param pulumi.Input[str] http_proxy: Http proxy URL.
        :param pulumi.Input[str] https_proxy: Https proxy URL.
        :param pulumi.Input[str] proxy_ca_certificate: Server CA Certificate file. Use the `file` function to read from a file.
        """
        if http_proxy is not None:
            pulumi.set(__self__, "http_proxy", http_proxy)
        if https_proxy is not None:
            pulumi.set(__self__, "https_proxy", https_proxy)
        if proxy_ca_certificate is not None:
            pulumi.set(__self__, "proxy_ca_certificate", proxy_ca_certificate)

    @property
    @pulumi.getter(name="httpProxy")
    def http_proxy(self) -> Optional[pulumi.Input[str]]:
        """
        Http proxy URL.
        """
        return pulumi.get(self, "http_proxy")

    @http_proxy.setter
    def http_proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_proxy", value)

    @property
    @pulumi.getter(name="httpsProxy")
    def https_proxy(self) -> Optional[pulumi.Input[str]]:
        """
        Https proxy URL.
        """
        return pulumi.get(self, "https_proxy")

    @https_proxy.setter
    def https_proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "https_proxy", value)

    @property
    @pulumi.getter(name="proxyCaCertificate")
    def proxy_ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Server CA Certificate file. Use the `file` function to read from a file.
        """
        return pulumi.get(self, "proxy_ca_certificate")

    @proxy_ca_certificate.setter
    def proxy_ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_ca_certificate", value)


class AviatrixProxyConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_proxy: Optional[pulumi.Input[str]] = None,
                 https_proxy: Optional[pulumi.Input[str]] = None,
                 proxy_ca_certificate: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **aviatrix_proxy_config** resource allows management of an Aviatrix Controller's proxy configurations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Controller Proxy Config
        test_proxy_config = aviatrix.AviatrixProxyConfig("testProxyConfig",
            http_proxy="172.31.52.145:3127",
            https_proxy="172.31.52.145:3129",
            proxy_ca_certificate=(lambda path: open(path).read())("/path/to/ca.pem"))
        ```

        ## Import

        **controller_proxy_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13

        ```sh
         $ pulumi import aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig test 10-11-12-13
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_proxy: Http proxy URL.
        :param pulumi.Input[str] https_proxy: Https proxy URL.
        :param pulumi.Input[str] proxy_ca_certificate: Server CA Certificate file. Use the `file` function to read from a file.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixProxyConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **aviatrix_proxy_config** resource allows management of an Aviatrix Controller's proxy configurations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Controller Proxy Config
        test_proxy_config = aviatrix.AviatrixProxyConfig("testProxyConfig",
            http_proxy="172.31.52.145:3127",
            https_proxy="172.31.52.145:3129",
            proxy_ca_certificate=(lambda path: open(path).read())("/path/to/ca.pem"))
        ```

        ## Import

        **controller_proxy_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13

        ```sh
         $ pulumi import aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig test 10-11-12-13
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixProxyConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixProxyConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_proxy: Optional[pulumi.Input[str]] = None,
                 https_proxy: Optional[pulumi.Input[str]] = None,
                 proxy_ca_certificate: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixProxyConfigArgs.__new__(AviatrixProxyConfigArgs)

            if http_proxy is None and not opts.urn:
                raise TypeError("Missing required property 'http_proxy'")
            __props__.__dict__["http_proxy"] = http_proxy
            if https_proxy is None and not opts.urn:
                raise TypeError("Missing required property 'https_proxy'")
            __props__.__dict__["https_proxy"] = https_proxy
            __props__.__dict__["proxy_ca_certificate"] = proxy_ca_certificate
        super(AviatrixProxyConfig, __self__).__init__(
            'aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            http_proxy: Optional[pulumi.Input[str]] = None,
            https_proxy: Optional[pulumi.Input[str]] = None,
            proxy_ca_certificate: Optional[pulumi.Input[str]] = None) -> 'AviatrixProxyConfig':
        """
        Get an existing AviatrixProxyConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_proxy: Http proxy URL.
        :param pulumi.Input[str] https_proxy: Https proxy URL.
        :param pulumi.Input[str] proxy_ca_certificate: Server CA Certificate file. Use the `file` function to read from a file.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixProxyConfigState.__new__(_AviatrixProxyConfigState)

        __props__.__dict__["http_proxy"] = http_proxy
        __props__.__dict__["https_proxy"] = https_proxy
        __props__.__dict__["proxy_ca_certificate"] = proxy_ca_certificate
        return AviatrixProxyConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="httpProxy")
    def http_proxy(self) -> pulumi.Output[str]:
        """
        Http proxy URL.
        """
        return pulumi.get(self, "http_proxy")

    @property
    @pulumi.getter(name="httpsProxy")
    def https_proxy(self) -> pulumi.Output[str]:
        """
        Https proxy URL.
        """
        return pulumi.get(self, "https_proxy")

    @property
    @pulumi.getter(name="proxyCaCertificate")
    def proxy_ca_certificate(self) -> pulumi.Output[Optional[str]]:
        """
        Server CA Certificate file. Use the `file` function to read from a file.
        """
        return pulumi.get(self, "proxy_ca_certificate")

