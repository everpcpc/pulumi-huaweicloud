# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetServerResult',
    'AwaitableGetServerResult',
    'get_server',
    'get_server_output',
]

@pulumi.output_type
class GetServerResult:
    """
    A collection of values returned by getServer.
    """
    def __init__(__self__, coverage_sites=None, edgecloud_id=None, edgecloud_name=None, flavor_id=None, flavor_name=None, id=None, image_id=None, image_name=None, key_pair=None, name=None, nics=None, public_ip=None, security_groups=None, status=None, system_disk_id=None, user_data=None, volume_attacheds=None, vpc_id=None):
        if coverage_sites and not isinstance(coverage_sites, list):
            raise TypeError("Expected argument 'coverage_sites' to be a list")
        pulumi.set(__self__, "coverage_sites", coverage_sites)
        if edgecloud_id and not isinstance(edgecloud_id, str):
            raise TypeError("Expected argument 'edgecloud_id' to be a str")
        pulumi.set(__self__, "edgecloud_id", edgecloud_id)
        if edgecloud_name and not isinstance(edgecloud_name, str):
            raise TypeError("Expected argument 'edgecloud_name' to be a str")
        pulumi.set(__self__, "edgecloud_name", edgecloud_name)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if flavor_name and not isinstance(flavor_name, str):
            raise TypeError("Expected argument 'flavor_name' to be a str")
        pulumi.set(__self__, "flavor_name", flavor_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if image_name and not isinstance(image_name, str):
            raise TypeError("Expected argument 'image_name' to be a str")
        pulumi.set(__self__, "image_name", image_name)
        if key_pair and not isinstance(key_pair, str):
            raise TypeError("Expected argument 'key_pair' to be a str")
        pulumi.set(__self__, "key_pair", key_pair)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nics and not isinstance(nics, list):
            raise TypeError("Expected argument 'nics' to be a list")
        pulumi.set(__self__, "nics", nics)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_disk_id and not isinstance(system_disk_id, str):
            raise TypeError("Expected argument 'system_disk_id' to be a str")
        pulumi.set(__self__, "system_disk_id", system_disk_id)
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        pulumi.set(__self__, "user_data", user_data)
        if volume_attacheds and not isinstance(volume_attacheds, list):
            raise TypeError("Expected argument 'volume_attacheds' to be a list")
        pulumi.set(__self__, "volume_attacheds", volume_attacheds)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="coverageSites")
    def coverage_sites(self) -> Sequence['outputs.GetServerCoverageSiteResult']:
        """
        An array of site ID and operator for the IEC server. The object structure is documented below.
        """
        return pulumi.get(self, "coverage_sites")

    @property
    @pulumi.getter(name="edgecloudId")
    def edgecloud_id(self) -> str:
        return pulumi.get(self, "edgecloud_id")

    @property
    @pulumi.getter(name="edgecloudName")
    def edgecloud_name(self) -> str:
        """
        The Name of the edgecloud service.
        """
        return pulumi.get(self, "edgecloud_name")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> str:
        """
        The flavor ID of the IEC server.
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="flavorName")
    def flavor_name(self) -> str:
        """
        The flavor name of the IEC server.
        """
        return pulumi.get(self, "flavor_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        The image ID of the IEC server.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> str:
        """
        The image name of the IEC server.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> str:
        """
        The name of a key pair to put on the IEC server.
        """
        return pulumi.get(self, "key_pair")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nics(self) -> Sequence['outputs.GetServerNicResult']:
        """
        An array of one or more networks to attach to the IEC server. The object structure is documented below.
        """
        return pulumi.get(self, "nics")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        """
        The EIP address that is associted to the IEC server.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[str]:
        """
        An array of one or more security group IDs to associate with the IEC server.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemDiskId")
    def system_disk_id(self) -> str:
        """
        The system disk voume ID.
        """
        return pulumi.get(self, "system_disk_id")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> str:
        """
        The user data (information after encoding) configured during IEC server creation.
        """
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter(name="volumeAttacheds")
    def volume_attacheds(self) -> Sequence['outputs.GetServerVolumeAttachedResult']:
        """
        An array of one or more disks to attach to the IEC server. The object structure is documented
        below.
        """
        return pulumi.get(self, "volume_attacheds")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of vpc for the IEC server.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetServerResult(GetServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerResult(
            coverage_sites=self.coverage_sites,
            edgecloud_id=self.edgecloud_id,
            edgecloud_name=self.edgecloud_name,
            flavor_id=self.flavor_id,
            flavor_name=self.flavor_name,
            id=self.id,
            image_id=self.image_id,
            image_name=self.image_name,
            key_pair=self.key_pair,
            name=self.name,
            nics=self.nics,
            public_ip=self.public_ip,
            security_groups=self.security_groups,
            status=self.status,
            system_disk_id=self.system_disk_id,
            user_data=self.user_data,
            volume_attacheds=self.volume_attacheds,
            vpc_id=self.vpc_id)


def get_server(edgecloud_id: Optional[str] = None,
               name: Optional[str] = None,
               status: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerResult:
    """
    Use this data source to get the details of a specified IEC server.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    server_name = config.require_object("serverName")
    demo = huaweicloud.Iec.get_server(name=server_name)
    ```


    :param str edgecloud_id: Specifies the ID of the edgecloud service.
    :param str name: Specifies the IEC server name, which can be queried with a regular expression.
    :param str status: Specifies the status of IEC server.
    """
    __args__ = dict()
    __args__['edgecloudId'] = edgecloud_id
    __args__['name'] = name
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Iec/getServer:getServer', __args__, opts=opts, typ=GetServerResult).value

    return AwaitableGetServerResult(
        coverage_sites=__ret__.coverage_sites,
        edgecloud_id=__ret__.edgecloud_id,
        edgecloud_name=__ret__.edgecloud_name,
        flavor_id=__ret__.flavor_id,
        flavor_name=__ret__.flavor_name,
        id=__ret__.id,
        image_id=__ret__.image_id,
        image_name=__ret__.image_name,
        key_pair=__ret__.key_pair,
        name=__ret__.name,
        nics=__ret__.nics,
        public_ip=__ret__.public_ip,
        security_groups=__ret__.security_groups,
        status=__ret__.status,
        system_disk_id=__ret__.system_disk_id,
        user_data=__ret__.user_data,
        volume_attacheds=__ret__.volume_attacheds,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_server)
def get_server_output(edgecloud_id: Optional[pulumi.Input[Optional[str]]] = None,
                      name: Optional[pulumi.Input[str]] = None,
                      status: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerResult]:
    """
    Use this data source to get the details of a specified IEC server.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    server_name = config.require_object("serverName")
    demo = huaweicloud.Iec.get_server(name=server_name)
    ```


    :param str edgecloud_id: Specifies the ID of the edgecloud service.
    :param str name: Specifies the IEC server name, which can be queried with a regular expression.
    :param str status: Specifies the status of IEC server.
    """
    ...
