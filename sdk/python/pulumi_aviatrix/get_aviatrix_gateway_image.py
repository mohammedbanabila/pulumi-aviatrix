# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAviatrixGatewayImageResult',
    'AwaitableGetAviatrixGatewayImageResult',
    'get_aviatrix_gateway_image',
    'get_aviatrix_gateway_image_output',
]

@pulumi.output_type
class GetAviatrixGatewayImageResult:
    """
    A collection of values returned by getAviatrixGatewayImage.
    """
    def __init__(__self__, cloud_type=None, id=None, image_version=None, software_version=None):
        if cloud_type and not isinstance(cloud_type, int):
            raise TypeError("Expected argument 'cloud_type' to be a int")
        pulumi.set(__self__, "cloud_type", cloud_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_version and not isinstance(image_version, str):
            raise TypeError("Expected argument 'image_version' to be a str")
        pulumi.set(__self__, "image_version", image_version)
        if software_version and not isinstance(software_version, str):
            raise TypeError("Expected argument 'software_version' to be a str")
        pulumi.set(__self__, "software_version", software_version)

    @property
    @pulumi.getter(name="cloudType")
    def cloud_type(self) -> int:
        return pulumi.get(self, "cloud_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageVersion")
    def image_version(self) -> str:
        """
        Image version that is compatible with the given cloud_type and software_version.
        """
        return pulumi.get(self, "image_version")

    @property
    @pulumi.getter(name="softwareVersion")
    def software_version(self) -> str:
        return pulumi.get(self, "software_version")


class AwaitableGetAviatrixGatewayImageResult(GetAviatrixGatewayImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAviatrixGatewayImageResult(
            cloud_type=self.cloud_type,
            id=self.id,
            image_version=self.image_version,
            software_version=self.software_version)


def get_aviatrix_gateway_image(cloud_type: Optional[int] = None,
                               software_version: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAviatrixGatewayImageResult:
    """
    The **aviatrix_gateway_image** data source provides the current image version that pairs with the given software version
    and cloud type.

    This data source is useful for getting the correct image_version for a gateway when upgrading the software_version of
    the gateway.


    :param int cloud_type: Cloud type. Type: Integer. Example: 1 (AWS)
    :param str software_version: Software version. Type: String. Example: "6.4.2487"
    """
    __args__ = dict()
    __args__['cloudType'] = cloud_type
    __args__['softwareVersion'] = software_version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aviatrix:index/getAviatrixGatewayImage:getAviatrixGatewayImage', __args__, opts=opts, typ=GetAviatrixGatewayImageResult).value

    return AwaitableGetAviatrixGatewayImageResult(
        cloud_type=__ret__.cloud_type,
        id=__ret__.id,
        image_version=__ret__.image_version,
        software_version=__ret__.software_version)


@_utilities.lift_output_func(get_aviatrix_gateway_image)
def get_aviatrix_gateway_image_output(cloud_type: Optional[pulumi.Input[int]] = None,
                                      software_version: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAviatrixGatewayImageResult]:
    """
    The **aviatrix_gateway_image** data source provides the current image version that pairs with the given software version
    and cloud type.

    This data source is useful for getting the correct image_version for a gateway when upgrading the software_version of
    the gateway.


    :param int cloud_type: Cloud type. Type: Integer. Example: 1 (AWS)
    :param str software_version: Software version. Type: String. Example: "6.4.2487"
    """
    ...
