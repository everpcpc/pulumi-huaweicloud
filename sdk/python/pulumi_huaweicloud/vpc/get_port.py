# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPortResult',
    'AwaitableGetPortResult',
    'get_port',
    'get_port_output',
]

@pulumi.output_type
class GetPortResult:
    """
    A collection of values returned by getPort.
    """
    def __init__(__self__, admin_state_up=None, all_allowed_ips=None, all_fixed_ips=None, all_security_group_ids=None, device_id=None, device_owner=None, fixed_ip=None, id=None, mac_address=None, name=None, network_id=None, port_id=None, project_id=None, region=None, security_group_ids=None, status=None, tenant_id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        if admin_state_up is not None:
            warnings.warn("""this field is deprecated""", DeprecationWarning)
            pulumi.log.warn("""admin_state_up is deprecated: this field is deprecated""")

        pulumi.set(__self__, "admin_state_up", admin_state_up)
        if all_allowed_ips and not isinstance(all_allowed_ips, list):
            raise TypeError("Expected argument 'all_allowed_ips' to be a list")
        pulumi.set(__self__, "all_allowed_ips", all_allowed_ips)
        if all_fixed_ips and not isinstance(all_fixed_ips, list):
            raise TypeError("Expected argument 'all_fixed_ips' to be a list")
        pulumi.set(__self__, "all_fixed_ips", all_fixed_ips)
        if all_security_group_ids and not isinstance(all_security_group_ids, list):
            raise TypeError("Expected argument 'all_security_group_ids' to be a list")
        pulumi.set(__self__, "all_security_group_ids", all_security_group_ids)
        if device_id and not isinstance(device_id, str):
            raise TypeError("Expected argument 'device_id' to be a str")
        pulumi.set(__self__, "device_id", device_id)
        if device_owner and not isinstance(device_owner, str):
            raise TypeError("Expected argument 'device_owner' to be a str")
        pulumi.set(__self__, "device_owner", device_owner)
        if fixed_ip and not isinstance(fixed_ip, str):
            raise TypeError("Expected argument 'fixed_ip' to be a str")
        pulumi.set(__self__, "fixed_ip", fixed_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        pulumi.set(__self__, "port_id", port_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        if project_id is not None:
            warnings.warn("""this field is deprecated""", DeprecationWarning)
            pulumi.log.warn("""project_id is deprecated: this field is deprecated""")

        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        if tenant_id is not None:
            warnings.warn("""this field is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: this field is deprecated""")

        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> bool:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="allAllowedIps")
    def all_allowed_ips(self) -> Sequence[str]:
        """
        The collection of allowed IP addresses on the port.
        """
        return pulumi.get(self, "all_allowed_ips")

    @property
    @pulumi.getter(name="allFixedIps")
    def all_fixed_ips(self) -> Sequence[str]:
        """
        The collection of Fixed IP addresses on the port.
        """
        return pulumi.get(self, "all_fixed_ips")

    @property
    @pulumi.getter(name="allSecurityGroupIds")
    def all_security_group_ids(self) -> Sequence[str]:
        """
        The collection of security group IDs applied on the port.
        """
        return pulumi.get(self, "all_security_group_ids")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> str:
        """
        The ID of the device the port belongs to.
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="deviceOwner")
    def device_owner(self) -> str:
        """
        The device owner of the port.
        """
        return pulumi.get(self, "device_owner")

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[str]:
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the port.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> str:
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")


class AwaitableGetPortResult(GetPortResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortResult(
            admin_state_up=self.admin_state_up,
            all_allowed_ips=self.all_allowed_ips,
            all_fixed_ips=self.all_fixed_ips,
            all_security_group_ids=self.all_security_group_ids,
            device_id=self.device_id,
            device_owner=self.device_owner,
            fixed_ip=self.fixed_ip,
            id=self.id,
            mac_address=self.mac_address,
            name=self.name,
            network_id=self.network_id,
            port_id=self.port_id,
            project_id=self.project_id,
            region=self.region,
            security_group_ids=self.security_group_ids,
            status=self.status,
            tenant_id=self.tenant_id)


def get_port(admin_state_up: Optional[bool] = None,
             device_id: Optional[str] = None,
             device_owner: Optional[str] = None,
             fixed_ip: Optional[str] = None,
             mac_address: Optional[str] = None,
             name: Optional[str] = None,
             network_id: Optional[str] = None,
             port_id: Optional[str] = None,
             project_id: Optional[str] = None,
             region: Optional[str] = None,
             security_group_ids: Optional[Sequence[str]] = None,
             status: Optional[str] = None,
             tenant_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortResult:
    """
    Use this data source to get the ID of an available HuaweiCloud port.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    port1 = huaweicloud.Vpc.get_port(network_id=var["network_id"],
        fixed_ip="192.168.0.100")
    ```


    :param str device_id: The ID of the device the port belongs to.
    :param str device_owner: The device owner of the port.
    :param str fixed_ip: Specifies the port IP address filter.
    :param str mac_address: Specifies the MAC address of the port.
    :param str name: The name of the port.
    :param str network_id: Specifies the ID of the network the port belongs to.
    :param str port_id: Specifies the ID of the port.
    :param str region: Specifies the region in which to obtain the port. If omitted, the provider-level region
           will be used.
    :param Sequence[str] security_group_ids: The list of port security group IDs to filter.
    :param str status: Specifies the status of the port.
    """
    __args__ = dict()
    __args__['adminStateUp'] = admin_state_up
    __args__['deviceId'] = device_id
    __args__['deviceOwner'] = device_owner
    __args__['fixedIp'] = fixed_ip
    __args__['macAddress'] = mac_address
    __args__['name'] = name
    __args__['networkId'] = network_id
    __args__['portId'] = port_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['securityGroupIds'] = security_group_ids
    __args__['status'] = status
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Vpc/getPort:getPort', __args__, opts=opts, typ=GetPortResult).value

    return AwaitableGetPortResult(
        admin_state_up=__ret__.admin_state_up,
        all_allowed_ips=__ret__.all_allowed_ips,
        all_fixed_ips=__ret__.all_fixed_ips,
        all_security_group_ids=__ret__.all_security_group_ids,
        device_id=__ret__.device_id,
        device_owner=__ret__.device_owner,
        fixed_ip=__ret__.fixed_ip,
        id=__ret__.id,
        mac_address=__ret__.mac_address,
        name=__ret__.name,
        network_id=__ret__.network_id,
        port_id=__ret__.port_id,
        project_id=__ret__.project_id,
        region=__ret__.region,
        security_group_ids=__ret__.security_group_ids,
        status=__ret__.status,
        tenant_id=__ret__.tenant_id)


@_utilities.lift_output_func(get_port)
def get_port_output(admin_state_up: Optional[pulumi.Input[Optional[bool]]] = None,
                    device_id: Optional[pulumi.Input[Optional[str]]] = None,
                    device_owner: Optional[pulumi.Input[Optional[str]]] = None,
                    fixed_ip: Optional[pulumi.Input[Optional[str]]] = None,
                    mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    network_id: Optional[pulumi.Input[Optional[str]]] = None,
                    port_id: Optional[pulumi.Input[Optional[str]]] = None,
                    project_id: Optional[pulumi.Input[Optional[str]]] = None,
                    region: Optional[pulumi.Input[Optional[str]]] = None,
                    security_group_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    status: Optional[pulumi.Input[Optional[str]]] = None,
                    tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortResult]:
    """
    Use this data source to get the ID of an available HuaweiCloud port.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    port1 = huaweicloud.Vpc.get_port(network_id=var["network_id"],
        fixed_ip="192.168.0.100")
    ```


    :param str device_id: The ID of the device the port belongs to.
    :param str device_owner: The device owner of the port.
    :param str fixed_ip: Specifies the port IP address filter.
    :param str mac_address: Specifies the MAC address of the port.
    :param str name: The name of the port.
    :param str network_id: Specifies the ID of the network the port belongs to.
    :param str port_id: Specifies the ID of the port.
    :param str region: Specifies the region in which to obtain the port. If omitted, the provider-level region
           will be used.
    :param Sequence[str] security_group_ids: The list of port security group IDs to filter.
    :param str status: Specifies the status of the port.
    """
    ...
