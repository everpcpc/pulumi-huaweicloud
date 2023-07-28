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
    'FunctionCustomImage',
    'FunctionFuncMount',
    'TriggerApig',
    'TriggerDis',
    'TriggerKafka',
    'TriggerLts',
    'TriggerObs',
    'TriggerSmn',
    'TriggerTimer',
    'GetDependenciesPackageResult',
]

@pulumi.output_type
class FunctionCustomImage(dict):
    def __init__(__self__, *,
                 url: str):
        """
        :param str url: Specifies the URL of SWR image, the URL must start with `swr.`.
        """
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        Specifies the URL of SWR image, the URL must start with `swr.`.
        """
        return pulumi.get(self, "url")


@pulumi.output_type
class FunctionFuncMount(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "localMountPath":
            suggest = "local_mount_path"
        elif key == "mountResource":
            suggest = "mount_resource"
        elif key == "mountSharePath":
            suggest = "mount_share_path"
        elif key == "mountType":
            suggest = "mount_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FunctionFuncMount. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FunctionFuncMount.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FunctionFuncMount.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 local_mount_path: str,
                 mount_resource: str,
                 mount_share_path: str,
                 mount_type: str,
                 status: Optional[str] = None):
        """
        :param str local_mount_path: Specifies the function access path.
        :param str mount_resource: Specifies the ID of the mounted resource (corresponding cloud service).
        :param str mount_share_path: Specifies the remote mount path. Example: 192.168.0.12:/data.
        :param str mount_type: Specifies the mount type. Options: sfs, sfsTurbo, and ecs.
        """
        pulumi.set(__self__, "local_mount_path", local_mount_path)
        pulumi.set(__self__, "mount_resource", mount_resource)
        pulumi.set(__self__, "mount_share_path", mount_share_path)
        pulumi.set(__self__, "mount_type", mount_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="localMountPath")
    def local_mount_path(self) -> str:
        """
        Specifies the function access path.
        """
        return pulumi.get(self, "local_mount_path")

    @property
    @pulumi.getter(name="mountResource")
    def mount_resource(self) -> str:
        """
        Specifies the ID of the mounted resource (corresponding cloud service).
        """
        return pulumi.get(self, "mount_resource")

    @property
    @pulumi.getter(name="mountSharePath")
    def mount_share_path(self) -> str:
        """
        Specifies the remote mount path. Example: 192.168.0.12:/data.
        """
        return pulumi.get(self, "mount_share_path")

    @property
    @pulumi.getter(name="mountType")
    def mount_type(self) -> str:
        """
        Specifies the mount type. Options: sfs, sfsTurbo, and ecs.
        """
        return pulumi.get(self, "mount_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


@pulumi.output_type
class TriggerApig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiName":
            suggest = "api_name"
        elif key == "envName":
            suggest = "env_name"
        elif key == "groupId":
            suggest = "group_id"
        elif key == "instanceId":
            suggest = "instance_id"
        elif key == "requestProtocol":
            suggest = "request_protocol"
        elif key == "securityAuthentication":
            suggest = "security_authentication"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerApig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerApig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerApig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_name: str,
                 env_name: str,
                 group_id: str,
                 instance_id: Optional[str] = None,
                 request_protocol: Optional[str] = None,
                 security_authentication: Optional[str] = None,
                 timeout: Optional[int] = None):
        """
        :param str api_name: Specifies the API name. Changing this will create a new trigger resource.
        :param str env_name: Specifies the API environment name.
               Changing this will create a new trigger resource.
        :param str group_id: Specifies the ID of the APIG group to which the API belongs.
               Changing this will create a new trigger resource.
        :param str instance_id: Specifies the ID of the APIG dedicated instance to which the API belongs.
               Required if the `type` is `DEDICATEDGATEWAY`. Changing this will create a new trigger resource.
        :param str request_protocol: Specifies the request protocol of the API. The valid value are
               **HTTP** and **HTTPS**. Default to **HTTPS**. Changing this will create a new trigger resource.
        :param str security_authentication: Specifies the security authentication mode. The valid values
               are **NONE**, **APP** and **IAM**, default to **IAM**. Changing this will create a new trigger resource.
        :param int timeout: Specifies the timeout for request sending. The valid value is range form
               `1` to `60,000`, default to `5,000`. Changing this will create a new trigger resource.
        """
        pulumi.set(__self__, "api_name", api_name)
        pulumi.set(__self__, "env_name", env_name)
        pulumi.set(__self__, "group_id", group_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if request_protocol is not None:
            pulumi.set(__self__, "request_protocol", request_protocol)
        if security_authentication is not None:
            pulumi.set(__self__, "security_authentication", security_authentication)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> str:
        """
        Specifies the API name. Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "api_name")

    @property
    @pulumi.getter(name="envName")
    def env_name(self) -> str:
        """
        Specifies the API environment name.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "env_name")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        """
        Specifies the ID of the APIG group to which the API belongs.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        Specifies the ID of the APIG dedicated instance to which the API belongs.
        Required if the `type` is `DEDICATEDGATEWAY`. Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="requestProtocol")
    def request_protocol(self) -> Optional[str]:
        """
        Specifies the request protocol of the API. The valid value are
        **HTTP** and **HTTPS**. Default to **HTTPS**. Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "request_protocol")

    @property
    @pulumi.getter(name="securityAuthentication")
    def security_authentication(self) -> Optional[str]:
        """
        Specifies the security authentication mode. The valid values
        are **NONE**, **APP** and **IAM**, default to **IAM**. Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "security_authentication")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        """
        Specifies the timeout for request sending. The valid value is range form
        `1` to `60,000`, default to `5,000`. Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "timeout")


@pulumi.output_type
class TriggerDis(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxFetchBytes":
            suggest = "max_fetch_bytes"
        elif key == "pullPeriod":
            suggest = "pull_period"
        elif key == "serialEnable":
            suggest = "serial_enable"
        elif key == "startingPosition":
            suggest = "starting_position"
        elif key == "streamName":
            suggest = "stream_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerDis. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerDis.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerDis.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_fetch_bytes: int,
                 pull_period: int,
                 serial_enable: bool,
                 starting_position: str,
                 stream_name: str):
        """
        :param int max_fetch_bytes: Specifies the maximum volume of data that can be obtained for a single
               request, in Byte. Only the records with a size smaller than this value can be obtained.
               The valid value is range from `1,024` to `4,194,304`.
               Changing this will create a new trigger resource.
        :param int pull_period: Specifies the interval at which data is pulled from the specified stream.
               The valid value is range from `2` to `60,000`.
               Changing this will create a new trigger resource.
        :param bool serial_enable: Specifies the determines whether to pull data only after the data pulled
               in the last period has been processed.
               Changing this will create a new trigger resource.
        :param str starting_position: Specifies the type of starting position for DIS queue.
               The valid values are as follows:
               + **TRIM_HORIZON**: Starts reading from the earliest data stored in the partitions.
               + **LATEST**: Starts reading from the latest data stored in the partitions.
               Changing this will create a new trigger resource.
        :param str stream_name: Specifies the name of the DIS stream resource.
               Changing this will create a new trigger resource.
        """
        pulumi.set(__self__, "max_fetch_bytes", max_fetch_bytes)
        pulumi.set(__self__, "pull_period", pull_period)
        pulumi.set(__self__, "serial_enable", serial_enable)
        pulumi.set(__self__, "starting_position", starting_position)
        pulumi.set(__self__, "stream_name", stream_name)

    @property
    @pulumi.getter(name="maxFetchBytes")
    def max_fetch_bytes(self) -> int:
        """
        Specifies the maximum volume of data that can be obtained for a single
        request, in Byte. Only the records with a size smaller than this value can be obtained.
        The valid value is range from `1,024` to `4,194,304`.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "max_fetch_bytes")

    @property
    @pulumi.getter(name="pullPeriod")
    def pull_period(self) -> int:
        """
        Specifies the interval at which data is pulled from the specified stream.
        The valid value is range from `2` to `60,000`.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "pull_period")

    @property
    @pulumi.getter(name="serialEnable")
    def serial_enable(self) -> bool:
        """
        Specifies the determines whether to pull data only after the data pulled
        in the last period has been processed.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "serial_enable")

    @property
    @pulumi.getter(name="startingPosition")
    def starting_position(self) -> str:
        """
        Specifies the type of starting position for DIS queue.
        The valid values are as follows:
        + **TRIM_HORIZON**: Starts reading from the earliest data stored in the partitions.
        + **LATEST**: Starts reading from the latest data stored in the partitions.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "starting_position")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> str:
        """
        Specifies the name of the DIS stream resource.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "stream_name")


@pulumi.output_type
class TriggerKafka(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "instanceId":
            suggest = "instance_id"
        elif key == "topicIds":
            suggest = "topic_ids"
        elif key == "batchSize":
            suggest = "batch_size"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerKafka. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerKafka.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerKafka.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 instance_id: str,
                 topic_ids: Sequence[str],
                 batch_size: Optional[int] = None):
        """
        :param str instance_id: Specifies the ID of the APIG dedicated instance to which the API belongs.
               Required if the `type` is `DEDICATEDGATEWAY`. Changing this will create a new trigger resource.
        :param Sequence[str] topic_ids: Specifies one or more topic IDs of DMS kafka instance.
               Changing this will create a new trigger resource.
        :param int batch_size: Specifies the The number of messages consumed from the topic each time.
               The valid value is range from `1` to `1,000`. Defaults to `100`.
               Changing this will create a new trigger resource.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "topic_ids", topic_ids)
        if batch_size is not None:
            pulumi.set(__self__, "batch_size", batch_size)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        Specifies the ID of the APIG dedicated instance to which the API belongs.
        Required if the `type` is `DEDICATEDGATEWAY`. Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="topicIds")
    def topic_ids(self) -> Sequence[str]:
        """
        Specifies one or more topic IDs of DMS kafka instance.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "topic_ids")

    @property
    @pulumi.getter(name="batchSize")
    def batch_size(self) -> Optional[int]:
        """
        Specifies the The number of messages consumed from the topic each time.
        The valid value is range from `1` to `1,000`. Defaults to `100`.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "batch_size")


@pulumi.output_type
class TriggerLts(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "logGroupId":
            suggest = "log_group_id"
        elif key == "logTopicId":
            suggest = "log_topic_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerLts. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerLts.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerLts.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 log_group_id: str,
                 log_topic_id: str):
        """
        :param str log_group_id: Specifies the log group ID.
               Changing this will create a new trigger resource.
        :param str log_topic_id: Specifies the log stream ID.
               Changing this will create a new trigger resource.
        """
        pulumi.set(__self__, "log_group_id", log_group_id)
        pulumi.set(__self__, "log_topic_id", log_topic_id)

    @property
    @pulumi.getter(name="logGroupId")
    def log_group_id(self) -> str:
        """
        Specifies the log group ID.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "log_group_id")

    @property
    @pulumi.getter(name="logTopicId")
    def log_topic_id(self) -> str:
        """
        Specifies the log stream ID.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "log_topic_id")


@pulumi.output_type
class TriggerObs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bucketName":
            suggest = "bucket_name"
        elif key == "eventNotificationName":
            suggest = "event_notification_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerObs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerObs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerObs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bucket_name: str,
                 event_notification_name: str,
                 events: Sequence[str],
                 prefix: Optional[str] = None,
                 suffix: Optional[str] = None):
        """
        :param str bucket_name: Specifies the OBS bucket name.
               Changing this will create a new trigger resource.
        :param str event_notification_name: Specifies the event notification name.
               Changing this will create a new trigger resource.
        :param Sequence[str] events: Specifies the events that can trigger functions.
               Changing this will create a new trigger resource.
               The valid values are as follows:
               + **ObjectCreated**, **Put**, **Post**, **Copy** and **CompleteMultipartUpload**.
               + **ObjectRemoved**, **Delete** and **DeleteMarkerCreated**.
        :param str prefix: Specifies the prefix to limit notifications to objects beginning with this keyword.
               Changing this will create a new trigger resource.
        :param str suffix: Specifies the suffix to limit notifications to objects ending with this keyword.
               Changing this will create a new trigger resource.
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "event_notification_name", event_notification_name)
        pulumi.set(__self__, "events", events)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if suffix is not None:
            pulumi.set(__self__, "suffix", suffix)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        """
        Specifies the OBS bucket name.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="eventNotificationName")
    def event_notification_name(self) -> str:
        """
        Specifies the event notification name.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "event_notification_name")

    @property
    @pulumi.getter
    def events(self) -> Sequence[str]:
        """
        Specifies the events that can trigger functions.
        Changing this will create a new trigger resource.
        The valid values are as follows:
        + **ObjectCreated**, **Put**, **Post**, **Copy** and **CompleteMultipartUpload**.
        + **ObjectRemoved**, **Delete** and **DeleteMarkerCreated**.
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter
    def prefix(self) -> Optional[str]:
        """
        Specifies the prefix to limit notifications to objects beginning with this keyword.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def suffix(self) -> Optional[str]:
        """
        Specifies the suffix to limit notifications to objects ending with this keyword.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "suffix")


@pulumi.output_type
class TriggerSmn(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "topicUrn":
            suggest = "topic_urn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerSmn. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerSmn.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerSmn.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 topic_urn: str):
        """
        :param str topic_urn: Specifies the Uniform Resource Name (URN) for SMN topic.
               Changing this will create a new trigger resource.
        """
        pulumi.set(__self__, "topic_urn", topic_urn)

    @property
    @pulumi.getter(name="topicUrn")
    def topic_urn(self) -> str:
        """
        Specifies the Uniform Resource Name (URN) for SMN topic.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "topic_urn")


@pulumi.output_type
class TriggerTimer(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "scheduleType":
            suggest = "schedule_type"
        elif key == "additionalInformation":
            suggest = "additional_information"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerTimer. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerTimer.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerTimer.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 schedule: str,
                 schedule_type: str,
                 additional_information: Optional[str] = None):
        """
        :param str name: Specifies the trigger name, which can contains of 1 to 64 characters.
               The name must start with a letter, only letters, digits, hyphens (-) and underscores (_) are allowed.
               Changing this will create a new trigger resource.
        :param str schedule: Specifies the time schedule.
               For the rate type, schedule is composed of time and time unit.
               The time unit supports minutes (m), hours (h) and days (d).
               For the corn expression, please refer to the HuaweiCloud
               [document](https://support.huaweicloud.com/en-us/usermanual-functiongraph/functiongraph_01_0908.html).
               Changing this will create a new trigger resource.
        :param str schedule_type: Specifies the type of the time schedule.
               The valid values are **Rate** and **Cron**.
               Changing this will create a new trigger resource.
        :param str additional_information: Specifies the event used by the timer to trigger the function.
               Changing this will create a new trigger resource.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "schedule_type", schedule_type)
        if additional_information is not None:
            pulumi.set(__self__, "additional_information", additional_information)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the trigger name, which can contains of 1 to 64 characters.
        The name must start with a letter, only letters, digits, hyphens (-) and underscores (_) are allowed.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        Specifies the time schedule.
        For the rate type, schedule is composed of time and time unit.
        The time unit supports minutes (m), hours (h) and days (d).
        For the corn expression, please refer to the HuaweiCloud
        [document](https://support.huaweicloud.com/en-us/usermanual-functiongraph/functiongraph_01_0908.html).
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="scheduleType")
    def schedule_type(self) -> str:
        """
        Specifies the type of the time schedule.
        The valid values are **Rate** and **Cron**.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "schedule_type")

    @property
    @pulumi.getter(name="additionalInformation")
    def additional_information(self) -> Optional[str]:
        """
        Specifies the event used by the timer to trigger the function.
        Changing this will create a new trigger resource.
        """
        return pulumi.get(self, "additional_information")


@pulumi.output_type
class GetDependenciesPackageResult(dict):
    def __init__(__self__, *,
                 etag: str,
                 file_name: str,
                 id: str,
                 link: str,
                 name: str,
                 owner: str,
                 runtime: str,
                 size: int):
        """
        :param str etag: Unique ID of the dependent package.
        :param str file_name: File name of the Dependent package.
        :param str id: Dependent package ID.
        :param str link: URL of the dependent package in the OBS console.
        :param str name: Specifies the dependent package runtime to match.
        :param str owner: Dependent package owner.
        :param str runtime: Specifies the dependent package runtime to match. Valid values: **Java8**,
               **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**, **Python3.6**, **Go1.8**,
               **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and **PHP7.3**.
        :param int size: Dependent package size.
        """
        pulumi.set(__self__, "etag", etag)
        pulumi.set(__self__, "file_name", file_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "link", link)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "runtime", runtime)
        pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Unique ID of the dependent package.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> str:
        """
        File name of the Dependent package.
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Dependent package ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def link(self) -> str:
        """
        URL of the dependent package in the OBS console.
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the dependent package runtime to match.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        Dependent package owner.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def runtime(self) -> str:
        """
        Specifies the dependent package runtime to match. Valid values: **Java8**,
        **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**, **Python3.6**, **Go1.8**,
        **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and **PHP7.3**.
        """
        return pulumi.get(self, "runtime")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        Dependent package size.
        """
        return pulumi.get(self, "size")


