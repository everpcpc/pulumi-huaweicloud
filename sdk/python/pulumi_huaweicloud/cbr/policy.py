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
from ._inputs import *

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 backup_cycle: pulumi.Input['PolicyBackupCycleArgs'],
                 type: pulumi.Input[str],
                 backup_quantity: Optional[pulumi.Input[int]] = None,
                 destination_project_id: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 enable_acceleration: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 long_term_retention: Optional[pulumi.Input['PolicyLongTermRetentionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 time_period: Optional[pulumi.Input[int]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input['PolicyBackupCycleArgs'] backup_cycle: Specifies the scheduling rule for the policy backup execution.
               The object structure is documented below.
        :param pulumi.Input[str] type: Specifies the protection type of the policy.
               Valid values are **backup** and **replication**.
               Changing this will create a new policy.
        :param pulumi.Input[int] backup_quantity: Specifies the maximum number of retained backups. The value ranges from `2` to
               `99,999`. This parameter and `time_period` are alternative.
        :param pulumi.Input[str] destination_project_id: Specifies the ID of the replication destination project, which is
               mandatory for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[str] destination_region: Specifies the name of the replication destination region, which is mandatory
               for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[bool] enable_acceleration: Specifies whether to enable the acceleration function to shorten
               the replication time for cross-region.
               Changing this will create a new policy.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the policy. Default to **true**.
        :param pulumi.Input['PolicyLongTermRetentionArgs'] long_term_retention: Specifies the long-term retention rules, which is an advanced options of
               the `backup_quantity`. The object structure is documented below.
        :param pulumi.Input[str] name: Specifies the policy name.  
               This parameter can contain a maximum of 64
               characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
        :param pulumi.Input[str] region: Specifies the region where the policy is located. If omitted, the
               provider-level region will be used. Changing this will create a new policy.
        :param pulumi.Input[int] time_period: Specifies the duration (in days) for retained backups. The value ranges from `2` to
               `99,999`.
        :param pulumi.Input[str] time_zone: Specifies the UTC time zone, e.g. `UTC+08:00`.
               Only available if `long_term_retention` is set.
        """
        pulumi.set(__self__, "backup_cycle", backup_cycle)
        pulumi.set(__self__, "type", type)
        if backup_quantity is not None:
            pulumi.set(__self__, "backup_quantity", backup_quantity)
        if destination_project_id is not None:
            pulumi.set(__self__, "destination_project_id", destination_project_id)
        if destination_region is not None:
            pulumi.set(__self__, "destination_region", destination_region)
        if enable_acceleration is not None:
            pulumi.set(__self__, "enable_acceleration", enable_acceleration)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if long_term_retention is not None:
            pulumi.set(__self__, "long_term_retention", long_term_retention)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter(name="backupCycle")
    def backup_cycle(self) -> pulumi.Input['PolicyBackupCycleArgs']:
        """
        Specifies the scheduling rule for the policy backup execution.
        The object structure is documented below.
        """
        return pulumi.get(self, "backup_cycle")

    @backup_cycle.setter
    def backup_cycle(self, value: pulumi.Input['PolicyBackupCycleArgs']):
        pulumi.set(self, "backup_cycle", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the protection type of the policy.
        Valid values are **backup** and **replication**.
        Changing this will create a new policy.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="backupQuantity")
    def backup_quantity(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of retained backups. The value ranges from `2` to
        `99,999`. This parameter and `time_period` are alternative.
        """
        return pulumi.get(self, "backup_quantity")

    @backup_quantity.setter
    def backup_quantity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "backup_quantity", value)

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the replication destination project, which is
        mandatory for cross-region replication. Required if `protection_type` is **replication**.
        """
        return pulumi.get(self, "destination_project_id")

    @destination_project_id.setter
    def destination_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_project_id", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the replication destination region, which is mandatory
        for cross-region replication. Required if `protection_type` is **replication**.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="enableAcceleration")
    def enable_acceleration(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the acceleration function to shorten
        the replication time for cross-region.
        Changing this will create a new policy.
        """
        return pulumi.get(self, "enable_acceleration")

    @enable_acceleration.setter
    def enable_acceleration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_acceleration", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the policy. Default to **true**.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="longTermRetention")
    def long_term_retention(self) -> Optional[pulumi.Input['PolicyLongTermRetentionArgs']]:
        """
        Specifies the long-term retention rules, which is an advanced options of
        the `backup_quantity`. The object structure is documented below.
        """
        return pulumi.get(self, "long_term_retention")

    @long_term_retention.setter
    def long_term_retention(self, value: Optional[pulumi.Input['PolicyLongTermRetentionArgs']]):
        pulumi.set(self, "long_term_retention", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the policy name.  
        This parameter can contain a maximum of 64
        characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the policy is located. If omitted, the
        provider-level region will be used. Changing this will create a new policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the duration (in days) for retained backups. The value ranges from `2` to
        `99,999`.
        """
        return pulumi.get(self, "time_period")

    @time_period.setter
    def time_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_period", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the UTC time zone, e.g. `UTC+08:00`.
        Only available if `long_term_retention` is set.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_zone", value)


@pulumi.input_type
class _PolicyState:
    def __init__(__self__, *,
                 backup_cycle: Optional[pulumi.Input['PolicyBackupCycleArgs']] = None,
                 backup_quantity: Optional[pulumi.Input[int]] = None,
                 destination_project_id: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 enable_acceleration: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 long_term_retention: Optional[pulumi.Input['PolicyLongTermRetentionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 time_period: Optional[pulumi.Input[int]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Policy resources.
        :param pulumi.Input['PolicyBackupCycleArgs'] backup_cycle: Specifies the scheduling rule for the policy backup execution.
               The object structure is documented below.
        :param pulumi.Input[int] backup_quantity: Specifies the maximum number of retained backups. The value ranges from `2` to
               `99,999`. This parameter and `time_period` are alternative.
        :param pulumi.Input[str] destination_project_id: Specifies the ID of the replication destination project, which is
               mandatory for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[str] destination_region: Specifies the name of the replication destination region, which is mandatory
               for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[bool] enable_acceleration: Specifies whether to enable the acceleration function to shorten
               the replication time for cross-region.
               Changing this will create a new policy.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the policy. Default to **true**.
        :param pulumi.Input['PolicyLongTermRetentionArgs'] long_term_retention: Specifies the long-term retention rules, which is an advanced options of
               the `backup_quantity`. The object structure is documented below.
        :param pulumi.Input[str] name: Specifies the policy name.  
               This parameter can contain a maximum of 64
               characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
        :param pulumi.Input[str] region: Specifies the region where the policy is located. If omitted, the
               provider-level region will be used. Changing this will create a new policy.
        :param pulumi.Input[int] time_period: Specifies the duration (in days) for retained backups. The value ranges from `2` to
               `99,999`.
        :param pulumi.Input[str] time_zone: Specifies the UTC time zone, e.g. `UTC+08:00`.
               Only available if `long_term_retention` is set.
        :param pulumi.Input[str] type: Specifies the protection type of the policy.
               Valid values are **backup** and **replication**.
               Changing this will create a new policy.
        """
        if backup_cycle is not None:
            pulumi.set(__self__, "backup_cycle", backup_cycle)
        if backup_quantity is not None:
            pulumi.set(__self__, "backup_quantity", backup_quantity)
        if destination_project_id is not None:
            pulumi.set(__self__, "destination_project_id", destination_project_id)
        if destination_region is not None:
            pulumi.set(__self__, "destination_region", destination_region)
        if enable_acceleration is not None:
            pulumi.set(__self__, "enable_acceleration", enable_acceleration)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if long_term_retention is not None:
            pulumi.set(__self__, "long_term_retention", long_term_retention)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="backupCycle")
    def backup_cycle(self) -> Optional[pulumi.Input['PolicyBackupCycleArgs']]:
        """
        Specifies the scheduling rule for the policy backup execution.
        The object structure is documented below.
        """
        return pulumi.get(self, "backup_cycle")

    @backup_cycle.setter
    def backup_cycle(self, value: Optional[pulumi.Input['PolicyBackupCycleArgs']]):
        pulumi.set(self, "backup_cycle", value)

    @property
    @pulumi.getter(name="backupQuantity")
    def backup_quantity(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of retained backups. The value ranges from `2` to
        `99,999`. This parameter and `time_period` are alternative.
        """
        return pulumi.get(self, "backup_quantity")

    @backup_quantity.setter
    def backup_quantity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "backup_quantity", value)

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the replication destination project, which is
        mandatory for cross-region replication. Required if `protection_type` is **replication**.
        """
        return pulumi.get(self, "destination_project_id")

    @destination_project_id.setter
    def destination_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_project_id", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the replication destination region, which is mandatory
        for cross-region replication. Required if `protection_type` is **replication**.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="enableAcceleration")
    def enable_acceleration(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the acceleration function to shorten
        the replication time for cross-region.
        Changing this will create a new policy.
        """
        return pulumi.get(self, "enable_acceleration")

    @enable_acceleration.setter
    def enable_acceleration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_acceleration", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the policy. Default to **true**.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="longTermRetention")
    def long_term_retention(self) -> Optional[pulumi.Input['PolicyLongTermRetentionArgs']]:
        """
        Specifies the long-term retention rules, which is an advanced options of
        the `backup_quantity`. The object structure is documented below.
        """
        return pulumi.get(self, "long_term_retention")

    @long_term_retention.setter
    def long_term_retention(self, value: Optional[pulumi.Input['PolicyLongTermRetentionArgs']]):
        pulumi.set(self, "long_term_retention", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the policy name.  
        This parameter can contain a maximum of 64
        characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the policy is located. If omitted, the
        provider-level region will be used. Changing this will create a new policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the duration (in days) for retained backups. The value ranges from `2` to
        `99,999`.
        """
        return pulumi.get(self, "time_period")

    @time_period.setter
    def time_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_period", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the UTC time zone, e.g. `UTC+08:00`.
        Only available if `long_term_retention` is set.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_zone", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the protection type of the policy.
        Valid values are **backup** and **replication**.
        Changing this will create a new policy.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_cycle: Optional[pulumi.Input[pulumi.InputType['PolicyBackupCycleArgs']]] = None,
                 backup_quantity: Optional[pulumi.Input[int]] = None,
                 destination_project_id: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 enable_acceleration: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 long_term_retention: Optional[pulumi.Input[pulumi.InputType['PolicyLongTermRetentionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 time_period: Optional[pulumi.Input[int]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a CBR Policy resource within Huaweicloud.

        ## Example Usage

        ## Import

        Policies can be imported by their `id`. For example,

        ```sh
         $ pulumi import huaweicloud:Cbr/policy:Policy test 4d2c2939-774f-42ef-ab15-e5b126b11ace
        ```

         Note that the imported state may not be identical to your resource definition, due to the attribute missing from the API response. The missing attribute is`enable_acceleration`. It is generally recommended running `terraform plan` after importing a policy. You can then decide if changes should be applied to the policy, or the resource definition should be updated to align with the policy. Also you can ignore changes as below. resource "huaweicloud_cbr_policy" "test" {

         ...

         lifecycle {

         ignore_changes = [

         enable_acceleration,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyBackupCycleArgs']] backup_cycle: Specifies the scheduling rule for the policy backup execution.
               The object structure is documented below.
        :param pulumi.Input[int] backup_quantity: Specifies the maximum number of retained backups. The value ranges from `2` to
               `99,999`. This parameter and `time_period` are alternative.
        :param pulumi.Input[str] destination_project_id: Specifies the ID of the replication destination project, which is
               mandatory for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[str] destination_region: Specifies the name of the replication destination region, which is mandatory
               for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[bool] enable_acceleration: Specifies whether to enable the acceleration function to shorten
               the replication time for cross-region.
               Changing this will create a new policy.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the policy. Default to **true**.
        :param pulumi.Input[pulumi.InputType['PolicyLongTermRetentionArgs']] long_term_retention: Specifies the long-term retention rules, which is an advanced options of
               the `backup_quantity`. The object structure is documented below.
        :param pulumi.Input[str] name: Specifies the policy name.  
               This parameter can contain a maximum of 64
               characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
        :param pulumi.Input[str] region: Specifies the region where the policy is located. If omitted, the
               provider-level region will be used. Changing this will create a new policy.
        :param pulumi.Input[int] time_period: Specifies the duration (in days) for retained backups. The value ranges from `2` to
               `99,999`.
        :param pulumi.Input[str] time_zone: Specifies the UTC time zone, e.g. `UTC+08:00`.
               Only available if `long_term_retention` is set.
        :param pulumi.Input[str] type: Specifies the protection type of the policy.
               Valid values are **backup** and **replication**.
               Changing this will create a new policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a CBR Policy resource within Huaweicloud.

        ## Example Usage

        ## Import

        Policies can be imported by their `id`. For example,

        ```sh
         $ pulumi import huaweicloud:Cbr/policy:Policy test 4d2c2939-774f-42ef-ab15-e5b126b11ace
        ```

         Note that the imported state may not be identical to your resource definition, due to the attribute missing from the API response. The missing attribute is`enable_acceleration`. It is generally recommended running `terraform plan` after importing a policy. You can then decide if changes should be applied to the policy, or the resource definition should be updated to align with the policy. Also you can ignore changes as below. resource "huaweicloud_cbr_policy" "test" {

         ...

         lifecycle {

         ignore_changes = [

         enable_acceleration,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_cycle: Optional[pulumi.Input[pulumi.InputType['PolicyBackupCycleArgs']]] = None,
                 backup_quantity: Optional[pulumi.Input[int]] = None,
                 destination_project_id: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 enable_acceleration: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 long_term_retention: Optional[pulumi.Input[pulumi.InputType['PolicyLongTermRetentionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 time_period: Optional[pulumi.Input[int]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyArgs.__new__(PolicyArgs)

            if backup_cycle is None and not opts.urn:
                raise TypeError("Missing required property 'backup_cycle'")
            __props__.__dict__["backup_cycle"] = backup_cycle
            __props__.__dict__["backup_quantity"] = backup_quantity
            __props__.__dict__["destination_project_id"] = destination_project_id
            __props__.__dict__["destination_region"] = destination_region
            __props__.__dict__["enable_acceleration"] = enable_acceleration
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["long_term_retention"] = long_term_retention
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["time_period"] = time_period
            __props__.__dict__["time_zone"] = time_zone
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Policy, __self__).__init__(
            'huaweicloud:Cbr/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_cycle: Optional[pulumi.Input[pulumi.InputType['PolicyBackupCycleArgs']]] = None,
            backup_quantity: Optional[pulumi.Input[int]] = None,
            destination_project_id: Optional[pulumi.Input[str]] = None,
            destination_region: Optional[pulumi.Input[str]] = None,
            enable_acceleration: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            long_term_retention: Optional[pulumi.Input[pulumi.InputType['PolicyLongTermRetentionArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            time_period: Optional[pulumi.Input[int]] = None,
            time_zone: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyBackupCycleArgs']] backup_cycle: Specifies the scheduling rule for the policy backup execution.
               The object structure is documented below.
        :param pulumi.Input[int] backup_quantity: Specifies the maximum number of retained backups. The value ranges from `2` to
               `99,999`. This parameter and `time_period` are alternative.
        :param pulumi.Input[str] destination_project_id: Specifies the ID of the replication destination project, which is
               mandatory for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[str] destination_region: Specifies the name of the replication destination region, which is mandatory
               for cross-region replication. Required if `protection_type` is **replication**.
        :param pulumi.Input[bool] enable_acceleration: Specifies whether to enable the acceleration function to shorten
               the replication time for cross-region.
               Changing this will create a new policy.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the policy. Default to **true**.
        :param pulumi.Input[pulumi.InputType['PolicyLongTermRetentionArgs']] long_term_retention: Specifies the long-term retention rules, which is an advanced options of
               the `backup_quantity`. The object structure is documented below.
        :param pulumi.Input[str] name: Specifies the policy name.  
               This parameter can contain a maximum of 64
               characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
        :param pulumi.Input[str] region: Specifies the region where the policy is located. If omitted, the
               provider-level region will be used. Changing this will create a new policy.
        :param pulumi.Input[int] time_period: Specifies the duration (in days) for retained backups. The value ranges from `2` to
               `99,999`.
        :param pulumi.Input[str] time_zone: Specifies the UTC time zone, e.g. `UTC+08:00`.
               Only available if `long_term_retention` is set.
        :param pulumi.Input[str] type: Specifies the protection type of the policy.
               Valid values are **backup** and **replication**.
               Changing this will create a new policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyState.__new__(_PolicyState)

        __props__.__dict__["backup_cycle"] = backup_cycle
        __props__.__dict__["backup_quantity"] = backup_quantity
        __props__.__dict__["destination_project_id"] = destination_project_id
        __props__.__dict__["destination_region"] = destination_region
        __props__.__dict__["enable_acceleration"] = enable_acceleration
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["long_term_retention"] = long_term_retention
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["time_period"] = time_period
        __props__.__dict__["time_zone"] = time_zone
        __props__.__dict__["type"] = type
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupCycle")
    def backup_cycle(self) -> pulumi.Output['outputs.PolicyBackupCycle']:
        """
        Specifies the scheduling rule for the policy backup execution.
        The object structure is documented below.
        """
        return pulumi.get(self, "backup_cycle")

    @property
    @pulumi.getter(name="backupQuantity")
    def backup_quantity(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the maximum number of retained backups. The value ranges from `2` to
        `99,999`. This parameter and `time_period` are alternative.
        """
        return pulumi.get(self, "backup_quantity")

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID of the replication destination project, which is
        mandatory for cross-region replication. Required if `protection_type` is **replication**.
        """
        return pulumi.get(self, "destination_project_id")

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the name of the replication destination region, which is mandatory
        for cross-region replication. Required if `protection_type` is **replication**.
        """
        return pulumi.get(self, "destination_region")

    @property
    @pulumi.getter(name="enableAcceleration")
    def enable_acceleration(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable the acceleration function to shorten
        the replication time for cross-region.
        Changing this will create a new policy.
        """
        return pulumi.get(self, "enable_acceleration")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable the policy. Default to **true**.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="longTermRetention")
    def long_term_retention(self) -> pulumi.Output[Optional['outputs.PolicyLongTermRetention']]:
        """
        Specifies the long-term retention rules, which is an advanced options of
        the `backup_quantity`. The object structure is documented below.
        """
        return pulumi.get(self, "long_term_retention")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the policy name.  
        This parameter can contain a maximum of 64
        characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region where the policy is located. If omitted, the
        provider-level region will be used. Changing this will create a new policy.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the duration (in days) for retained backups. The value ranges from `2` to
        `99,999`.
        """
        return pulumi.get(self, "time_period")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Output[str]:
        """
        Specifies the UTC time zone, e.g. `UTC+08:00`.
        Only available if `long_term_retention` is set.
        """
        return pulumi.get(self, "time_zone")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the protection type of the policy.
        Valid values are **backup** and **replication**.
        Changing this will create a new policy.
        """
        return pulumi.get(self, "type")

