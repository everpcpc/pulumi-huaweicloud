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
    'JobDestinationDb',
    'JobLimitSpeed',
    'JobSourceDb',
]

@pulumi.output_type
class JobDestinationDb(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "engineType":
            suggest = "engine_type"
        elif key == "instanceId":
            suggest = "instance_id"
        elif key == "sslCertCheckSum":
            suggest = "ssl_cert_check_sum"
        elif key == "sslCertKey":
            suggest = "ssl_cert_key"
        elif key == "sslCertName":
            suggest = "ssl_cert_name"
        elif key == "sslCertPassword":
            suggest = "ssl_cert_password"
        elif key == "sslEnabled":
            suggest = "ssl_enabled"
        elif key == "subnetId":
            suggest = "subnet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobDestinationDb. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobDestinationDb.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobDestinationDb.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 engine_type: str,
                 ip: str,
                 password: str,
                 port: int,
                 user: str,
                 instance_id: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 ssl_cert_check_sum: Optional[str] = None,
                 ssl_cert_key: Optional[str] = None,
                 ssl_cert_name: Optional[str] = None,
                 ssl_cert_password: Optional[str] = None,
                 ssl_enabled: Optional[bool] = None,
                 subnet_id: Optional[str] = None):
        """
        :param str engine_type: Specifies the engine type of database. Changing this parameter will
               create a new resource. The options are as follows: `mysql`, `mongodb`, `gaussdbv5`.
        :param str ip: Specifies the IP of database. Changing this parameter will create a new resource.
        :param str password: Specifies the password of database.
               Changing this parameter will create a new resource.
        :param int port: Specifies the port of database. Changing this parameter will create a new resource.
        :param str user: Specifies the user name of database.
               Changing this parameter will create a new resource.
        :param str instance_id: Specifies the instance id of database when it is a RDS database.
               Changing this parameter will create a new resource.
        :param str name: Specifies the name of database.
               Changing this parameter will create a new resource.
        :param str region: Specifies the region which the database belongs when it is a RDS database.
               Changing this parameter will create a new resource.
        :param str ssl_cert_check_sum: Specifies the checksum of SSL certificate content.
               It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        :param str ssl_cert_key: Specifies the SSL certificate content, encrypted with base64.
               It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        :param str ssl_cert_name: Specifies SSL certificate name.
               It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        :param str ssl_cert_password: Specifies SSL certificate password. It is mandatory when
               `ssl_enabled` is `true` and the certificate file suffix is `.p12`. Changing this parameter will create a new resource.
        :param bool ssl_enabled: Specifies whether to enable SSL connection.
               Changing this parameter will create a new resource.
        :param str subnet_id: Specifies subnet ID of database when it is a RDS database.
               It is mandatory when `direction` is `down`. Changing this parameter will create a new resource.
        """
        pulumi.set(__self__, "engine_type", engine_type)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "user", user)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if ssl_cert_check_sum is not None:
            pulumi.set(__self__, "ssl_cert_check_sum", ssl_cert_check_sum)
        if ssl_cert_key is not None:
            pulumi.set(__self__, "ssl_cert_key", ssl_cert_key)
        if ssl_cert_name is not None:
            pulumi.set(__self__, "ssl_cert_name", ssl_cert_name)
        if ssl_cert_password is not None:
            pulumi.set(__self__, "ssl_cert_password", ssl_cert_password)
        if ssl_enabled is not None:
            pulumi.set(__self__, "ssl_enabled", ssl_enabled)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> str:
        """
        Specifies the engine type of database. Changing this parameter will
        create a new resource. The options are as follows: `mysql`, `mongodb`, `gaussdbv5`.
        """
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        Specifies the IP of database. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Specifies the password of database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Specifies the port of database. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def user(self) -> str:
        """
        Specifies the user name of database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        Specifies the instance id of database when it is a RDS database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Specifies the name of database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        Specifies the region which the database belongs when it is a RDS database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sslCertCheckSum")
    def ssl_cert_check_sum(self) -> Optional[str]:
        """
        Specifies the checksum of SSL certificate content.
        It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_check_sum")

    @property
    @pulumi.getter(name="sslCertKey")
    def ssl_cert_key(self) -> Optional[str]:
        """
        Specifies the SSL certificate content, encrypted with base64.
        It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_key")

    @property
    @pulumi.getter(name="sslCertName")
    def ssl_cert_name(self) -> Optional[str]:
        """
        Specifies SSL certificate name.
        It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_name")

    @property
    @pulumi.getter(name="sslCertPassword")
    def ssl_cert_password(self) -> Optional[str]:
        """
        Specifies SSL certificate password. It is mandatory when
        `ssl_enabled` is `true` and the certificate file suffix is `.p12`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_password")

    @property
    @pulumi.getter(name="sslEnabled")
    def ssl_enabled(self) -> Optional[bool]:
        """
        Specifies whether to enable SSL connection.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_enabled")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        Specifies subnet ID of database when it is a RDS database.
        It is mandatory when `direction` is `down`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "subnet_id")


@pulumi.output_type
class JobLimitSpeed(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endTime":
            suggest = "end_time"
        elif key == "startTime":
            suggest = "start_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobLimitSpeed. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobLimitSpeed.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobLimitSpeed.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 end_time: str,
                 speed: str,
                 start_time: str):
        """
        :param str end_time: Specifies the time to end speed limit, this time is UTC time. The input must
               end at 59 minutes, the format is `hh:mm`, for example: 15:59. Changing this parameter will create a new resource.
        :param str speed: Specifies the transmission speed, the value range is 1 to 9999, unit: `MB/s`.
               Changing this parameter will create a new resource.
        :param str start_time: Specifies the time to start speed limit, this time is UTC time. The start
               time is the whole hour, if there is a minute, it will be ignored, the format is `hh:mm`, and the hour number
               is two digits, for example: 01:00. Changing this parameter will create a new resource.
        """
        pulumi.set(__self__, "end_time", end_time)
        pulumi.set(__self__, "speed", speed)
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        Specifies the time to end speed limit, this time is UTC time. The input must
        end at 59 minutes, the format is `hh:mm`, for example: 15:59. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def speed(self) -> str:
        """
        Specifies the transmission speed, the value range is 1 to 9999, unit: `MB/s`.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "speed")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Specifies the time to start speed limit, this time is UTC time. The start
        time is the whole hour, if there is a minute, it will be ignored, the format is `hh:mm`, and the hour number
        is two digits, for example: 01:00. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "start_time")


@pulumi.output_type
class JobSourceDb(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "engineType":
            suggest = "engine_type"
        elif key == "instanceId":
            suggest = "instance_id"
        elif key == "sslCertCheckSum":
            suggest = "ssl_cert_check_sum"
        elif key == "sslCertKey":
            suggest = "ssl_cert_key"
        elif key == "sslCertName":
            suggest = "ssl_cert_name"
        elif key == "sslCertPassword":
            suggest = "ssl_cert_password"
        elif key == "sslEnabled":
            suggest = "ssl_enabled"
        elif key == "subnetId":
            suggest = "subnet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobSourceDb. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobSourceDb.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobSourceDb.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 engine_type: str,
                 ip: str,
                 password: str,
                 port: int,
                 user: str,
                 instance_id: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 ssl_cert_check_sum: Optional[str] = None,
                 ssl_cert_key: Optional[str] = None,
                 ssl_cert_name: Optional[str] = None,
                 ssl_cert_password: Optional[str] = None,
                 ssl_enabled: Optional[bool] = None,
                 subnet_id: Optional[str] = None):
        """
        :param str engine_type: Specifies the engine type of database. Changing this parameter will
               create a new resource. The options are as follows: `mysql`, `mongodb`, `gaussdbv5`.
        :param str ip: Specifies the IP of database. Changing this parameter will create a new resource.
        :param str password: Specifies the password of database.
               Changing this parameter will create a new resource.
        :param int port: Specifies the port of database. Changing this parameter will create a new resource.
        :param str user: Specifies the user name of database.
               Changing this parameter will create a new resource.
        :param str instance_id: Specifies the instance id of database when it is a RDS database.
               Changing this parameter will create a new resource.
        :param str name: Specifies the name of database.
               Changing this parameter will create a new resource.
        :param str region: Specifies the region which the database belongs when it is a RDS database.
               Changing this parameter will create a new resource.
        :param str ssl_cert_check_sum: Specifies the checksum of SSL certificate content.
               It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        :param str ssl_cert_key: Specifies the SSL certificate content, encrypted with base64.
               It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        :param str ssl_cert_name: Specifies SSL certificate name.
               It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        :param str ssl_cert_password: Specifies SSL certificate password. It is mandatory when
               `ssl_enabled` is `true` and the certificate file suffix is `.p12`. Changing this parameter will create a new resource.
        :param bool ssl_enabled: Specifies whether to enable SSL connection.
               Changing this parameter will create a new resource.
        :param str subnet_id: Specifies subnet ID of database when it is a RDS database.
               It is mandatory when `direction` is `down`. Changing this parameter will create a new resource.
        """
        pulumi.set(__self__, "engine_type", engine_type)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "user", user)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if ssl_cert_check_sum is not None:
            pulumi.set(__self__, "ssl_cert_check_sum", ssl_cert_check_sum)
        if ssl_cert_key is not None:
            pulumi.set(__self__, "ssl_cert_key", ssl_cert_key)
        if ssl_cert_name is not None:
            pulumi.set(__self__, "ssl_cert_name", ssl_cert_name)
        if ssl_cert_password is not None:
            pulumi.set(__self__, "ssl_cert_password", ssl_cert_password)
        if ssl_enabled is not None:
            pulumi.set(__self__, "ssl_enabled", ssl_enabled)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> str:
        """
        Specifies the engine type of database. Changing this parameter will
        create a new resource. The options are as follows: `mysql`, `mongodb`, `gaussdbv5`.
        """
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        Specifies the IP of database. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Specifies the password of database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Specifies the port of database. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def user(self) -> str:
        """
        Specifies the user name of database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        Specifies the instance id of database when it is a RDS database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Specifies the name of database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        Specifies the region which the database belongs when it is a RDS database.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sslCertCheckSum")
    def ssl_cert_check_sum(self) -> Optional[str]:
        """
        Specifies the checksum of SSL certificate content.
        It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_check_sum")

    @property
    @pulumi.getter(name="sslCertKey")
    def ssl_cert_key(self) -> Optional[str]:
        """
        Specifies the SSL certificate content, encrypted with base64.
        It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_key")

    @property
    @pulumi.getter(name="sslCertName")
    def ssl_cert_name(self) -> Optional[str]:
        """
        Specifies SSL certificate name.
        It is mandatory when `ssl_enabled` is `true`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_name")

    @property
    @pulumi.getter(name="sslCertPassword")
    def ssl_cert_password(self) -> Optional[str]:
        """
        Specifies SSL certificate password. It is mandatory when
        `ssl_enabled` is `true` and the certificate file suffix is `.p12`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_cert_password")

    @property
    @pulumi.getter(name="sslEnabled")
    def ssl_enabled(self) -> Optional[bool]:
        """
        Specifies whether to enable SSL connection.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "ssl_enabled")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        Specifies subnet ID of database when it is a RDS database.
        It is mandatory when `direction` is `down`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "subnet_id")


