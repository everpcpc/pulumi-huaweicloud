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
    'GetCassandraInstanceResult',
    'AwaitableGetCassandraInstanceResult',
    'get_cassandra_instance',
    'get_cassandra_instance_output',
]

@pulumi.output_type
class GetCassandraInstanceResult:
    """
    A collection of values returned by getCassandraInstance.
    """
    def __init__(__self__, availability_zone=None, backup_strategies=None, datastores=None, db_user_name=None, enterprise_project_id=None, flavor=None, id=None, mode=None, name=None, node_num=None, nodes=None, port=None, private_ips=None, region=None, security_group_id=None, status=None, subnet_id=None, tags=None, volume_size=None, vpc_id=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if backup_strategies and not isinstance(backup_strategies, list):
            raise TypeError("Expected argument 'backup_strategies' to be a list")
        pulumi.set(__self__, "backup_strategies", backup_strategies)
        if datastores and not isinstance(datastores, list):
            raise TypeError("Expected argument 'datastores' to be a list")
        pulumi.set(__self__, "datastores", datastores)
        if db_user_name and not isinstance(db_user_name, str):
            raise TypeError("Expected argument 'db_user_name' to be a str")
        pulumi.set(__self__, "db_user_name", db_user_name)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if flavor and not isinstance(flavor, str):
            raise TypeError("Expected argument 'flavor' to be a str")
        pulumi.set(__self__, "flavor", flavor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_num and not isinstance(node_num, int):
            raise TypeError("Expected argument 'node_num' to be a int")
        pulumi.set(__self__, "node_num", node_num)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        pulumi.set(__self__, "private_ips", private_ips)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if volume_size and not isinstance(volume_size, int):
            raise TypeError("Expected argument 'volume_size' to be a int")
        pulumi.set(__self__, "volume_size", volume_size)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        Indicates the availability zone where the node resides.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="backupStrategies")
    def backup_strategies(self) -> Sequence['outputs.GetCassandraInstanceBackupStrategyResult']:
        """
        Indicates the advanced backup policy. Structure is documented below.
        """
        return pulumi.get(self, "backup_strategies")

    @property
    @pulumi.getter
    def datastores(self) -> Sequence['outputs.GetCassandraInstanceDatastoreResult']:
        """
        Indicates the database information. Structure is documented below.
        """
        return pulumi.get(self, "datastores")

    @property
    @pulumi.getter(name="dbUserName")
    def db_user_name(self) -> str:
        """
        Indicates the default username.
        """
        return pulumi.get(self, "db_user_name")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> str:
        """
        Indicates the enterprise project id.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def flavor(self) -> str:
        """
        Indicates the instance specifications.
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        Indicates the instance mode.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Indicates the node name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> int:
        """
        Indicates the count of the nodes.
        """
        return pulumi.get(self, "node_num")

    @property
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.GetCassandraInstanceNodeResult']:
        """
        Indicates the instance nodes information. Structure is documented below.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Indicates the database port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> Sequence[str]:
        """
        Indicates the list of private IP address of the nodes.
        """
        return pulumi.get(self, "private_ips")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        """
        Indicates the security group ID.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Indicates the node status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Indicates the key/value tags of the instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> int:
        """
        Indicates the size of the volume.
        """
        return pulumi.get(self, "volume_size")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetCassandraInstanceResult(GetCassandraInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCassandraInstanceResult(
            availability_zone=self.availability_zone,
            backup_strategies=self.backup_strategies,
            datastores=self.datastores,
            db_user_name=self.db_user_name,
            enterprise_project_id=self.enterprise_project_id,
            flavor=self.flavor,
            id=self.id,
            mode=self.mode,
            name=self.name,
            node_num=self.node_num,
            nodes=self.nodes,
            port=self.port,
            private_ips=self.private_ips,
            region=self.region,
            security_group_id=self.security_group_id,
            status=self.status,
            subnet_id=self.subnet_id,
            tags=self.tags,
            volume_size=self.volume_size,
            vpc_id=self.vpc_id)


def get_cassandra_instance(name: Optional[str] = None,
                           region: Optional[str] = None,
                           subnet_id: Optional[str] = None,
                           vpc_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCassandraInstanceResult:
    """
    Use this data source to get available HuaweiCloud gaussdb cassandra instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDBforNoSQL.get_cassandra_instance(name="gaussdb-instance")
    ```


    :param str name: Specifies the name of the instance.
    :param str region: The region in which to obtain the instance. If omitted, the provider-level region will
           be used.
    :param str subnet_id: Specifies the network ID of a subnet.
    :param str vpc_id: Specifies the VPC ID.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    __args__['subnetId'] = subnet_id
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:GaussDBforNoSQL/getCassandraInstance:getCassandraInstance', __args__, opts=opts, typ=GetCassandraInstanceResult).value

    return AwaitableGetCassandraInstanceResult(
        availability_zone=__ret__.availability_zone,
        backup_strategies=__ret__.backup_strategies,
        datastores=__ret__.datastores,
        db_user_name=__ret__.db_user_name,
        enterprise_project_id=__ret__.enterprise_project_id,
        flavor=__ret__.flavor,
        id=__ret__.id,
        mode=__ret__.mode,
        name=__ret__.name,
        node_num=__ret__.node_num,
        nodes=__ret__.nodes,
        port=__ret__.port,
        private_ips=__ret__.private_ips,
        region=__ret__.region,
        security_group_id=__ret__.security_group_id,
        status=__ret__.status,
        subnet_id=__ret__.subnet_id,
        tags=__ret__.tags,
        volume_size=__ret__.volume_size,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_cassandra_instance)
def get_cassandra_instance_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                  region: Optional[pulumi.Input[Optional[str]]] = None,
                                  subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCassandraInstanceResult]:
    """
    Use this data source to get available HuaweiCloud gaussdb cassandra instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDBforNoSQL.get_cassandra_instance(name="gaussdb-instance")
    ```


    :param str name: Specifies the name of the instance.
    :param str region: The region in which to obtain the instance. If omitted, the provider-level region will
           be used.
    :param str subnet_id: Specifies the network ID of a subnet.
    :param str vpc_id: Specifies the VPC ID.
    """
    ...