# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixControllerCertDomainConfigArgs', 'AviatrixControllerCertDomainConfig']

@pulumi.input_type
class AviatrixControllerCertDomainConfigArgs:
    def __init__(__self__, *,
                 cert_domain: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixControllerCertDomainConfig resource.
        :param pulumi.Input[str] cert_domain: Domain name that is used in FQDN for generating cert.
        """
        if cert_domain is not None:
            pulumi.set(__self__, "cert_domain", cert_domain)

    @property
    @pulumi.getter(name="certDomain")
    def cert_domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name that is used in FQDN for generating cert.
        """
        return pulumi.get(self, "cert_domain")

    @cert_domain.setter
    def cert_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_domain", value)


@pulumi.input_type
class _AviatrixControllerCertDomainConfigState:
    def __init__(__self__, *,
                 cert_domain: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixControllerCertDomainConfig resources.
        :param pulumi.Input[str] cert_domain: Domain name that is used in FQDN for generating cert.
        """
        if cert_domain is not None:
            pulumi.set(__self__, "cert_domain", cert_domain)

    @property
    @pulumi.getter(name="certDomain")
    def cert_domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name that is used in FQDN for generating cert.
        """
        return pulumi.get(self, "cert_domain")

    @cert_domain.setter
    def cert_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_domain", value)


class AviatrixControllerCertDomainConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert_domain: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixControllerCertDomainConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_domain: Domain name that is used in FQDN for generating cert.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AviatrixControllerCertDomainConfigArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixControllerCertDomainConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixControllerCertDomainConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixControllerCertDomainConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert_domain: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixControllerCertDomainConfigArgs.__new__(AviatrixControllerCertDomainConfigArgs)

            __props__.__dict__["cert_domain"] = cert_domain
        super(AviatrixControllerCertDomainConfig, __self__).__init__(
            'aviatrix:index/aviatrixControllerCertDomainConfig:AviatrixControllerCertDomainConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cert_domain: Optional[pulumi.Input[str]] = None) -> 'AviatrixControllerCertDomainConfig':
        """
        Get an existing AviatrixControllerCertDomainConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_domain: Domain name that is used in FQDN for generating cert.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixControllerCertDomainConfigState.__new__(_AviatrixControllerCertDomainConfigState)

        __props__.__dict__["cert_domain"] = cert_domain
        return AviatrixControllerCertDomainConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certDomain")
    def cert_domain(self) -> pulumi.Output[Optional[str]]:
        """
        Domain name that is used in FQDN for generating cert.
        """
        return pulumi.get(self, "cert_domain")

