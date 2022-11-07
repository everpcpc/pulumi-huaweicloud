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
    'FileSystemAccessRule',
    'GetTurbosTurboResult',
]

@pulumi.output_type
class FileSystemAccessRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessLevel":
            suggest = "access_level"
        elif key == "accessRuleId":
            suggest = "access_rule_id"
        elif key == "accessTo":
            suggest = "access_to"
        elif key == "accessType":
            suggest = "access_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FileSystemAccessRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FileSystemAccessRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FileSystemAccessRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_level: Optional[str] = None,
                 access_rule_id: Optional[str] = None,
                 access_to: Optional[str] = None,
                 access_type: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str access_level: Specifies the access level of the shared file system. Possible values are *ro* (
               read-only)
               and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
        :param str access_rule_id: The UUID of the share access rule.
        :param str access_to: Specifies the value that defines the access rule. The value contains 1 to 255
               characters. Changing this will create a new access rule. The value varies according to the scenario:
               + Set the VPC ID in VPC authorization scenarios.
               + Set this parameter in IP address authorization scenario:
               - For an NFS shared file system, the value in the format of *VPC_ID#IP_address#priority#user_permission*.
               For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
               - For a CIFS shared file system, the value in the format of *VPC_ID#IP_address#priority*.
               For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#0.
        :param str access_type: Specifies the type of the share access rule. The default value is *cert*. Changing
               this will create a new access rule.
        :param str status: The status of the share access rule.
        """
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if access_rule_id is not None:
            pulumi.set(__self__, "access_rule_id", access_rule_id)
        if access_to is not None:
            pulumi.set(__self__, "access_to", access_to)
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[str]:
        """
        Specifies the access level of the shared file system. Possible values are *ro* (
        read-only)
        and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="accessRuleId")
    def access_rule_id(self) -> Optional[str]:
        """
        The UUID of the share access rule.
        """
        return pulumi.get(self, "access_rule_id")

    @property
    @pulumi.getter(name="accessTo")
    def access_to(self) -> Optional[str]:
        """
        Specifies the value that defines the access rule. The value contains 1 to 255
        characters. Changing this will create a new access rule. The value varies according to the scenario:
        + Set the VPC ID in VPC authorization scenarios.
        + Set this parameter in IP address authorization scenario:
        - For an NFS shared file system, the value in the format of *VPC_ID#IP_address#priority#user_permission*.
        For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
        - For a CIFS shared file system, the value in the format of *VPC_ID#IP_address#priority*.
        For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#0.
        """
        return pulumi.get(self, "access_to")

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[str]:
        """
        Specifies the type of the share access rule. The default value is *cert*. Changing
        this will create a new access rule.
        """
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the share access rule.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetTurbosTurboResult(dict):
    def __init__(__self__, *,
                 availability_zone: str,
                 available_capacity: str,
                 crypt_key_id: str,
                 enhanced: bool,
                 enterprise_project_id: str,
                 export_location: str,
                 id: str,
                 name: str,
                 security_group_id: str,
                 share_proto: str,
                 share_type: str,
                 size: int,
                 subnet_id: str,
                 version: str,
                 vpc_id: str):
        """
        :param str availability_zone: The availability zone where the SFS turbo file system is located.
        :param str available_capacity: The available capacity of the SFS turbo file system, in GB.
        :param str crypt_key_id: The ID of a KMS key to encrypt the SFS turbo file system.
        :param bool enhanced: Whether the SFS turbo file system is enhanced.
        :param str enterprise_project_id: The enterprise project ID of the SFS turbo file system.
        :param str export_location: The mount point of the SFS turbo file system.
        :param str id: The resource ID of the SFS turbo file system.
        :param str name: Specifies the name of the SFS turbo file system.
        :param str security_group_id: The ID of the security group to which the SFS turbo belongs.
        :param str share_proto: Specifies the protocol of the SFS turbo file system. The valid value is **NFS**.
        :param str share_type: Specifies the type of the SFS turbo file system.
               The valid values are **STANDARD** and **PERFORMANCE**.
        :param int size: Specifies the capacity of the SFS turbo file system, in GB.
               The value ranges from `500` to `32,768`, and must be large than `10,240` for an enhanced file system.
        :param str subnet_id: The **network ID** of the subnet to which the SFS turbo belongs.
        :param str version: The version of the SFS turbo file system.
        :param str vpc_id: The ID of the VPC to which the SFS turbo belongs.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "available_capacity", available_capacity)
        pulumi.set(__self__, "crypt_key_id", crypt_key_id)
        pulumi.set(__self__, "enhanced", enhanced)
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        pulumi.set(__self__, "export_location", export_location)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "share_proto", share_proto)
        pulumi.set(__self__, "share_type", share_type)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "version", version)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The availability zone where the SFS turbo file system is located.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="availableCapacity")
    def available_capacity(self) -> str:
        """
        The available capacity of the SFS turbo file system, in GB.
        """
        return pulumi.get(self, "available_capacity")

    @property
    @pulumi.getter(name="cryptKeyId")
    def crypt_key_id(self) -> str:
        """
        The ID of a KMS key to encrypt the SFS turbo file system.
        """
        return pulumi.get(self, "crypt_key_id")

    @property
    @pulumi.getter
    def enhanced(self) -> bool:
        """
        Whether the SFS turbo file system is enhanced.
        """
        return pulumi.get(self, "enhanced")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> str:
        """
        The enterprise project ID of the SFS turbo file system.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="exportLocation")
    def export_location(self) -> str:
        """
        The mount point of the SFS turbo file system.
        """
        return pulumi.get(self, "export_location")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource ID of the SFS turbo file system.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the SFS turbo file system.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        """
        The ID of the security group to which the SFS turbo belongs.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="shareProto")
    def share_proto(self) -> str:
        """
        Specifies the protocol of the SFS turbo file system. The valid value is **NFS**.
        """
        return pulumi.get(self, "share_proto")

    @property
    @pulumi.getter(name="shareType")
    def share_type(self) -> str:
        """
        Specifies the type of the SFS turbo file system.
        The valid values are **STANDARD** and **PERFORMANCE**.
        """
        return pulumi.get(self, "share_type")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        Specifies the capacity of the SFS turbo file system, in GB.
        The value ranges from `500` to `32,768`, and must be large than `10,240` for an enhanced file system.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The **network ID** of the subnet to which the SFS turbo belongs.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of the SFS turbo file system.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the VPC to which the SFS turbo belongs.
        """
        return pulumi.get(self, "vpc_id")

