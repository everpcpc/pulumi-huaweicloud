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
    'KafkaInstanceCrossVpcAccess',
    'KafkaPermissionsPolicy',
    'GetFlavorsFlavorResult',
    'GetFlavorsFlavorIoResult',
    'GetFlavorsFlavorPropertyResult',
    'GetFlavorsFlavorSupportFeatureResult',
    'GetFlavorsFlavorSupportFeaturePropertyResult',
    'GetInstancesInstanceResult',
    'GetInstancesInstanceCrossVpcAccessResult',
]

@pulumi.output_type
class KafkaInstanceCrossVpcAccess(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "advertisedIp":
            suggest = "advertised_ip"
        elif key == "lisenterIp":
            suggest = "lisenter_ip"
        elif key == "portId":
            suggest = "port_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KafkaInstanceCrossVpcAccess. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KafkaInstanceCrossVpcAccess.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KafkaInstanceCrossVpcAccess.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 advertised_ip: Optional[str] = None,
                 lisenter_ip: Optional[str] = None,
                 port: Optional[int] = None,
                 port_id: Optional[str] = None):
        """
        :param str advertised_ip: -(Optional, String) The advertised IP Address or domain name.
        :param str lisenter_ip: The listener IP address.
        :param int port: The port number.
        :param str port_id: The port ID associated with the address.
        """
        if advertised_ip is not None:
            pulumi.set(__self__, "advertised_ip", advertised_ip)
        if lisenter_ip is not None:
            pulumi.set(__self__, "lisenter_ip", lisenter_ip)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)

    @property
    @pulumi.getter(name="advertisedIp")
    def advertised_ip(self) -> Optional[str]:
        """
        -(Optional, String) The advertised IP Address or domain name.
        """
        return pulumi.get(self, "advertised_ip")

    @property
    @pulumi.getter(name="lisenterIp")
    def lisenter_ip(self) -> Optional[str]:
        """
        The listener IP address.
        """
        return pulumi.get(self, "lisenter_ip")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        The port number.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[str]:
        """
        The port ID associated with the address.
        """
        return pulumi.get(self, "port_id")


@pulumi.output_type
class KafkaPermissionsPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessPolicy":
            suggest = "access_policy"
        elif key == "userName":
            suggest = "user_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KafkaPermissionsPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KafkaPermissionsPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KafkaPermissionsPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_policy: str,
                 user_name: str):
        """
        :param str access_policy: -(Required, String) Specifies the permissions type. The value can be:
               + **all**: publish and subscribe permissions.
               + **pub**: publish permissions.
               + **sub**: subscribe permissions.
        :param str user_name: -(Required, String) Specifies the user name.
        """
        pulumi.set(__self__, "access_policy", access_policy)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessPolicy")
    def access_policy(self) -> str:
        """
        -(Required, String) Specifies the permissions type. The value can be:
        + **all**: publish and subscribe permissions.
        + **pub**: publish permissions.
        + **sub**: subscribe permissions.
        """
        return pulumi.get(self, "access_policy")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        -(Required, String) Specifies the user name.
        """
        return pulumi.get(self, "user_name")


@pulumi.output_type
class GetFlavorsFlavorResult(dict):
    def __init__(__self__, *,
                 arch_types: Sequence[str],
                 charging_modes: Sequence[str],
                 id: str,
                 ios: Sequence['outputs.GetFlavorsFlavorIoResult'],
                 properties: Sequence['outputs.GetFlavorsFlavorPropertyResult'],
                 support_features: Sequence['outputs.GetFlavorsFlavorSupportFeatureResult'],
                 type: str,
                 vm_specification: str):
        """
        :param Sequence[str] arch_types: The list of supported CPU architectures.
        :param Sequence[str] charging_modes: The list of supported billing modes.
        :param str id: The flavor ID.
        :param Sequence['GetFlavorsFlavorIoArgs'] ios: The list of supported disk IO types.
               The object structure is documented below.
        :param Sequence['GetFlavorsFlavorPropertyArgs'] properties: The function property details.
               The object structure is documented below.
        :param Sequence['GetFlavorsFlavorSupportFeatureArgs'] support_features: The list of features supported by the current specification.
               The object structure is documented below.
        :param str type: Specifies flavor type. The valid values are **single** and **cluster**.
        :param str vm_specification: The underlying VM specification.
        """
        pulumi.set(__self__, "arch_types", arch_types)
        pulumi.set(__self__, "charging_modes", charging_modes)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ios", ios)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "support_features", support_features)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "vm_specification", vm_specification)

    @property
    @pulumi.getter(name="archTypes")
    def arch_types(self) -> Sequence[str]:
        """
        The list of supported CPU architectures.
        """
        return pulumi.get(self, "arch_types")

    @property
    @pulumi.getter(name="chargingModes")
    def charging_modes(self) -> Sequence[str]:
        """
        The list of supported billing modes.
        """
        return pulumi.get(self, "charging_modes")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The flavor ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ios(self) -> Sequence['outputs.GetFlavorsFlavorIoResult']:
        """
        The list of supported disk IO types.
        The object structure is documented below.
        """
        return pulumi.get(self, "ios")

    @property
    @pulumi.getter
    def properties(self) -> Sequence['outputs.GetFlavorsFlavorPropertyResult']:
        """
        The function property details.
        The object structure is documented below.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="supportFeatures")
    def support_features(self) -> Sequence['outputs.GetFlavorsFlavorSupportFeatureResult']:
        """
        The list of features supported by the current specification.
        The object structure is documented below.
        """
        return pulumi.get(self, "support_features")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies flavor type. The valid values are **single** and **cluster**.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmSpecification")
    def vm_specification(self) -> str:
        """
        The underlying VM specification.
        """
        return pulumi.get(self, "vm_specification")


@pulumi.output_type
class GetFlavorsFlavorIoResult(dict):
    def __init__(__self__, *,
                 availability_zones: Sequence[str],
                 storage_spec_code: str,
                 type: str,
                 unavailability_zones: Sequence[str]):
        """
        :param Sequence[str] availability_zones: Specifies the list of availability zones with available resources.
        :param str storage_spec_code: Specifies the disk IO encoding.
               + **dms.physical.storage.high.v2**: Type of the disk that uses high I/O.
               + **dms.physical.storage.ultra.v2**: Type of the disk that uses ultra-high I/O.
        :param str type: Specifies flavor type. The valid values are **single** and **cluster**.
        :param Sequence[str] unavailability_zones: The list of unavailability zones with available resources.
        """
        pulumi.set(__self__, "availability_zones", availability_zones)
        pulumi.set(__self__, "storage_spec_code", storage_spec_code)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "unavailability_zones", unavailability_zones)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Sequence[str]:
        """
        Specifies the list of availability zones with available resources.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="storageSpecCode")
    def storage_spec_code(self) -> str:
        """
        Specifies the disk IO encoding.
        + **dms.physical.storage.high.v2**: Type of the disk that uses high I/O.
        + **dms.physical.storage.ultra.v2**: Type of the disk that uses ultra-high I/O.
        """
        return pulumi.get(self, "storage_spec_code")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies flavor type. The valid values are **single** and **cluster**.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="unavailabilityZones")
    def unavailability_zones(self) -> Sequence[str]:
        """
        The list of unavailability zones with available resources.
        """
        return pulumi.get(self, "unavailability_zones")


@pulumi.output_type
class GetFlavorsFlavorPropertyResult(dict):
    def __init__(__self__, *,
                 flavor_alias: str,
                 max_bandwidth_per_broker: int,
                 max_broker: int,
                 max_consumer_per_broker: int,
                 max_partition_per_broker: int,
                 max_storage_per_node: int,
                 max_tps_per_broker: int,
                 min_broker: int,
                 min_storage_per_node: int):
        """
        :param str flavor_alias: The flavor ID alias.
        :param int max_bandwidth_per_broker: The maximum bandwidth per broker.
        :param int max_broker: The maximum number of brokers.
        :param int max_consumer_per_broker: The maximum number of consumers per broker.
        :param int max_partition_per_broker: The maximum number of partitions per broker.
        :param int max_storage_per_node: The maximum storage per node. The unit is GB.
        :param int max_tps_per_broker: The maximum TPS per broker.
        :param int min_broker: The minimum number of brokers.
        :param int min_storage_per_node: The minimum storage per node. The unit is GB.
        """
        pulumi.set(__self__, "flavor_alias", flavor_alias)
        pulumi.set(__self__, "max_bandwidth_per_broker", max_bandwidth_per_broker)
        pulumi.set(__self__, "max_broker", max_broker)
        pulumi.set(__self__, "max_consumer_per_broker", max_consumer_per_broker)
        pulumi.set(__self__, "max_partition_per_broker", max_partition_per_broker)
        pulumi.set(__self__, "max_storage_per_node", max_storage_per_node)
        pulumi.set(__self__, "max_tps_per_broker", max_tps_per_broker)
        pulumi.set(__self__, "min_broker", min_broker)
        pulumi.set(__self__, "min_storage_per_node", min_storage_per_node)

    @property
    @pulumi.getter(name="flavorAlias")
    def flavor_alias(self) -> str:
        """
        The flavor ID alias.
        """
        return pulumi.get(self, "flavor_alias")

    @property
    @pulumi.getter(name="maxBandwidthPerBroker")
    def max_bandwidth_per_broker(self) -> int:
        """
        The maximum bandwidth per broker.
        """
        return pulumi.get(self, "max_bandwidth_per_broker")

    @property
    @pulumi.getter(name="maxBroker")
    def max_broker(self) -> int:
        """
        The maximum number of brokers.
        """
        return pulumi.get(self, "max_broker")

    @property
    @pulumi.getter(name="maxConsumerPerBroker")
    def max_consumer_per_broker(self) -> int:
        """
        The maximum number of consumers per broker.
        """
        return pulumi.get(self, "max_consumer_per_broker")

    @property
    @pulumi.getter(name="maxPartitionPerBroker")
    def max_partition_per_broker(self) -> int:
        """
        The maximum number of partitions per broker.
        """
        return pulumi.get(self, "max_partition_per_broker")

    @property
    @pulumi.getter(name="maxStoragePerNode")
    def max_storage_per_node(self) -> int:
        """
        The maximum storage per node. The unit is GB.
        """
        return pulumi.get(self, "max_storage_per_node")

    @property
    @pulumi.getter(name="maxTpsPerBroker")
    def max_tps_per_broker(self) -> int:
        """
        The maximum TPS per broker.
        """
        return pulumi.get(self, "max_tps_per_broker")

    @property
    @pulumi.getter(name="minBroker")
    def min_broker(self) -> int:
        """
        The minimum number of brokers.
        """
        return pulumi.get(self, "min_broker")

    @property
    @pulumi.getter(name="minStoragePerNode")
    def min_storage_per_node(self) -> int:
        """
        The minimum storage per node. The unit is GB.
        """
        return pulumi.get(self, "min_storage_per_node")


@pulumi.output_type
class GetFlavorsFlavorSupportFeatureResult(dict):
    def __init__(__self__, *,
                 name: str,
                 properties: Sequence['outputs.GetFlavorsFlavorSupportFeaturePropertyResult']):
        """
        :param str name: The function name, e.g. **connector_obs**.
        :param Sequence['GetFlavorsFlavorSupportFeaturePropertyArgs'] properties: The function property details.
               The object structure is documented below.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The function name, e.g. **connector_obs**.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> Sequence['outputs.GetFlavorsFlavorSupportFeaturePropertyResult']:
        """
        The function property details.
        The object structure is documented below.
        """
        return pulumi.get(self, "properties")


@pulumi.output_type
class GetFlavorsFlavorSupportFeaturePropertyResult(dict):
    def __init__(__self__, *,
                 max_node: int,
                 max_task: int,
                 min_node: int,
                 min_task: int):
        """
        :param int max_node: The maximum number of nodes for the dump function.
        :param int max_task: The maximum number of tasks for the dump function.
        :param int min_node: The minimum number of nodes for the dump function.
        :param int min_task: The minimum number of tasks for the dump function.
        """
        pulumi.set(__self__, "max_node", max_node)
        pulumi.set(__self__, "max_task", max_task)
        pulumi.set(__self__, "min_node", min_node)
        pulumi.set(__self__, "min_task", min_task)

    @property
    @pulumi.getter(name="maxNode")
    def max_node(self) -> int:
        """
        The maximum number of nodes for the dump function.
        """
        return pulumi.get(self, "max_node")

    @property
    @pulumi.getter(name="maxTask")
    def max_task(self) -> int:
        """
        The maximum number of tasks for the dump function.
        """
        return pulumi.get(self, "max_task")

    @property
    @pulumi.getter(name="minNode")
    def min_node(self) -> int:
        """
        The minimum number of nodes for the dump function.
        """
        return pulumi.get(self, "min_node")

    @property
    @pulumi.getter(name="minTask")
    def min_task(self) -> int:
        """
        The minimum number of tasks for the dump function.
        """
        return pulumi.get(self, "min_task")


@pulumi.output_type
class GetInstancesInstanceResult(dict):
    def __init__(__self__, *,
                 access_user: str,
                 availability_zones: Sequence[str],
                 connect_address: str,
                 cross_vpc_accesses: Sequence['outputs.GetInstancesInstanceCrossVpcAccessResult'],
                 description: str,
                 dumping: bool,
                 enable_auto_topic: bool,
                 enable_public_ip: bool,
                 engine_version: str,
                 enterprise_project_id: str,
                 id: str,
                 maintain_begin: str,
                 maintain_end: str,
                 manager_user: str,
                 manegement_connect_address: str,
                 name: str,
                 network_id: str,
                 partition_num: int,
                 port: int,
                 product_id: str,
                 public_conn_addresses: str,
                 public_ip_ids: Sequence[str],
                 resource_spec_code: str,
                 retention_policy: str,
                 security_group_id: str,
                 ssl_enable: bool,
                 status: str,
                 storage_space: int,
                 storage_spec_code: str,
                 tags: Mapping[str, str],
                 type: str,
                 used_storage_space: int,
                 user_id: str,
                 user_name: str,
                 vpc_id: str):
        """
        :param str access_user: The access username.
        :param Sequence[str] availability_zones: The list of AZ names.
        :param str connect_address: The IP address for instance connection.
        :param Sequence['GetInstancesInstanceCrossVpcAccessArgs'] cross_vpc_accesses: Indicates the Access information of cross-VPC. The structure is documented below.
        :param str description: The instance description.
        :param bool dumping: Whether to dumping is enabled.
        :param bool enable_auto_topic: Whether to enable automatic topic creation.
        :param bool enable_public_ip: Whether public access to the instance is enabled.
        :param str engine_version: The kafka engine version.
        :param str enterprise_project_id: Specifies the enterprise project ID to which all instances of the list
               belong.
        :param str id: The instance ID.
        :param str maintain_begin: The time at which a maintenance time window starts, the format is `HH:mm`.
        :param str maintain_end: The time at which a maintenance time window ends, the format is `HH:mm`.
        :param str manager_user: The username for logging in to the Kafka Manager.
        :param str manegement_connect_address: The connection address of the Kafka manager of an instance.
        :param str name: Specifies the kafka instance name for data-source queries.
        :param str network_id: The subnet ID to which the instance belongs.
        :param int partition_num: The maximum number of topics in the DMS kafka instance.
        :param int port: The port number.
        :param str product_id: The product ID used by the instance.
        :param str public_conn_addresses: The instance public access address.
               The format of each connection address is `{IP address}:{port}`.
        :param Sequence[str] public_ip_ids: The IDs of the elastic IP address (EIP).
        :param str resource_spec_code: The resource specifications identifier.
        :param str security_group_id: The security group ID associated with the instance.
        :param bool ssl_enable: Whether the Kafka SASL_SSL is enabled.
        :param str status: Specifies the kafka instance status for data-source queries.
        :param int storage_space: The message storage capacity, in GB unit.
        :param str storage_spec_code: The storage I/O specification.
        :param Mapping[str, str] tags: The key/value pairs to associate with the instance.
        :param str type: The instance type.
        :param int used_storage_space: The used message storage space, in GB unit.
        :param str user_id: The user ID who created the instance.
        :param str user_name: The username who created the instance.
        :param str vpc_id: The VPC ID to which the instance belongs.
        """
        pulumi.set(__self__, "access_user", access_user)
        pulumi.set(__self__, "availability_zones", availability_zones)
        pulumi.set(__self__, "connect_address", connect_address)
        pulumi.set(__self__, "cross_vpc_accesses", cross_vpc_accesses)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "dumping", dumping)
        pulumi.set(__self__, "enable_auto_topic", enable_auto_topic)
        pulumi.set(__self__, "enable_public_ip", enable_public_ip)
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "maintain_begin", maintain_begin)
        pulumi.set(__self__, "maintain_end", maintain_end)
        pulumi.set(__self__, "manager_user", manager_user)
        pulumi.set(__self__, "manegement_connect_address", manegement_connect_address)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "partition_num", partition_num)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "product_id", product_id)
        pulumi.set(__self__, "public_conn_addresses", public_conn_addresses)
        pulumi.set(__self__, "public_ip_ids", public_ip_ids)
        pulumi.set(__self__, "resource_spec_code", resource_spec_code)
        pulumi.set(__self__, "retention_policy", retention_policy)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "ssl_enable", ssl_enable)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "storage_space", storage_space)
        pulumi.set(__self__, "storage_spec_code", storage_spec_code)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "used_storage_space", used_storage_space)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "user_name", user_name)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="accessUser")
    def access_user(self) -> str:
        """
        The access username.
        """
        return pulumi.get(self, "access_user")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Sequence[str]:
        """
        The list of AZ names.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="connectAddress")
    def connect_address(self) -> str:
        """
        The IP address for instance connection.
        """
        return pulumi.get(self, "connect_address")

    @property
    @pulumi.getter(name="crossVpcAccesses")
    def cross_vpc_accesses(self) -> Sequence['outputs.GetInstancesInstanceCrossVpcAccessResult']:
        """
        Indicates the Access information of cross-VPC. The structure is documented below.
        """
        return pulumi.get(self, "cross_vpc_accesses")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The instance description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def dumping(self) -> bool:
        """
        Whether to dumping is enabled.
        """
        return pulumi.get(self, "dumping")

    @property
    @pulumi.getter(name="enableAutoTopic")
    def enable_auto_topic(self) -> bool:
        """
        Whether to enable automatic topic creation.
        """
        return pulumi.get(self, "enable_auto_topic")

    @property
    @pulumi.getter(name="enablePublicIp")
    def enable_public_ip(self) -> bool:
        """
        Whether public access to the instance is enabled.
        """
        return pulumi.get(self, "enable_public_ip")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        """
        The kafka engine version.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> str:
        """
        Specifies the enterprise project ID to which all instances of the list
        belong.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The instance ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maintainBegin")
    def maintain_begin(self) -> str:
        """
        The time at which a maintenance time window starts, the format is `HH:mm`.
        """
        return pulumi.get(self, "maintain_begin")

    @property
    @pulumi.getter(name="maintainEnd")
    def maintain_end(self) -> str:
        """
        The time at which a maintenance time window ends, the format is `HH:mm`.
        """
        return pulumi.get(self, "maintain_end")

    @property
    @pulumi.getter(name="managerUser")
    def manager_user(self) -> str:
        """
        The username for logging in to the Kafka Manager.
        """
        return pulumi.get(self, "manager_user")

    @property
    @pulumi.getter(name="manegementConnectAddress")
    def manegement_connect_address(self) -> str:
        """
        The connection address of the Kafka manager of an instance.
        """
        return pulumi.get(self, "manegement_connect_address")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the kafka instance name for data-source queries.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        """
        The subnet ID to which the instance belongs.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="partitionNum")
    def partition_num(self) -> int:
        """
        The maximum number of topics in the DMS kafka instance.
        """
        return pulumi.get(self, "partition_num")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port number.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> str:
        """
        The product ID used by the instance.
        """
        return pulumi.get(self, "product_id")

    @property
    @pulumi.getter(name="publicConnAddresses")
    def public_conn_addresses(self) -> str:
        """
        The instance public access address.
        The format of each connection address is `{IP address}:{port}`.
        """
        return pulumi.get(self, "public_conn_addresses")

    @property
    @pulumi.getter(name="publicIpIds")
    def public_ip_ids(self) -> Sequence[str]:
        """
        The IDs of the elastic IP address (EIP).
        """
        return pulumi.get(self, "public_ip_ids")

    @property
    @pulumi.getter(name="resourceSpecCode")
    def resource_spec_code(self) -> str:
        """
        The resource specifications identifier.
        """
        return pulumi.get(self, "resource_spec_code")

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> str:
        return pulumi.get(self, "retention_policy")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        """
        The security group ID associated with the instance.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="sslEnable")
    def ssl_enable(self) -> bool:
        """
        Whether the Kafka SASL_SSL is enabled.
        """
        return pulumi.get(self, "ssl_enable")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Specifies the kafka instance status for data-source queries.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageSpace")
    def storage_space(self) -> int:
        """
        The message storage capacity, in GB unit.
        """
        return pulumi.get(self, "storage_space")

    @property
    @pulumi.getter(name="storageSpecCode")
    def storage_spec_code(self) -> str:
        """
        The storage I/O specification.
        """
        return pulumi.get(self, "storage_spec_code")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        The key/value pairs to associate with the instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The instance type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usedStorageSpace")
    def used_storage_space(self) -> int:
        """
        The used message storage space, in GB unit.
        """
        return pulumi.get(self, "used_storage_space")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The user ID who created the instance.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The username who created the instance.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The VPC ID to which the instance belongs.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GetInstancesInstanceCrossVpcAccessResult(dict):
    def __init__(__self__, *,
                 advertised_ip: str,
                 lisenter_ip: str,
                 port: int,
                 port_id: str):
        """
        :param str advertised_ip: The advertised IP Address.
        :param str lisenter_ip: The listener IP address.
        :param int port: The port number.
        :param str port_id: The port ID associated with the address.
        """
        pulumi.set(__self__, "advertised_ip", advertised_ip)
        pulumi.set(__self__, "lisenter_ip", lisenter_ip)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "port_id", port_id)

    @property
    @pulumi.getter(name="advertisedIp")
    def advertised_ip(self) -> str:
        """
        The advertised IP Address.
        """
        return pulumi.get(self, "advertised_ip")

    @property
    @pulumi.getter(name="lisenterIp")
    def lisenter_ip(self) -> str:
        """
        The listener IP address.
        """
        return pulumi.get(self, "lisenter_ip")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port number.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> str:
        """
        The port ID associated with the address.
        """
        return pulumi.get(self, "port_id")

