# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReferenceTableArgs', 'ReferenceTable']

@pulumi.input_type
class ReferenceTableArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReferenceTable resource.
        :param pulumi.Input[str] type: The type of the reference table, The options are `url`, `user-agent`, `ip`,
               `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] conditions: The conditions of the reference table. The maximum length is 30. The maximum length of
               condition is 2048 characters.
        :param pulumi.Input[str] description: The description of the reference table. The maximum length is 128 characters.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID of WAF reference table.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
               maximum length is 64 characters.
        :param pulumi.Input[str] region: The region in which to create the WAF reference table resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new reference table.
        """
        pulumi.set(__self__, "type", type)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the reference table, The options are `url`, `user-agent`, `ip`,
        `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The conditions of the reference table. The maximum length is 30. The maximum length of
        condition is 2048 characters.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the reference table. The maximum length is 128 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the enterprise project ID of WAF reference table.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
        maximum length is 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the WAF reference table resource. If omitted,
        the provider-level region will be used. Changing this setting will push a new reference table.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ReferenceTableState:
    def __init__(__self__, *,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReferenceTable resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] conditions: The conditions of the reference table. The maximum length is 30. The maximum length of
               condition is 2048 characters.
        :param pulumi.Input[str] creation_time: The server time when reference table was created.
        :param pulumi.Input[str] description: The description of the reference table. The maximum length is 128 characters.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID of WAF reference table.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
               maximum length is 64 characters.
        :param pulumi.Input[str] region: The region in which to create the WAF reference table resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new reference table.
        :param pulumi.Input[str] type: The type of the reference table, The options are `url`, `user-agent`, `ip`,
               `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.
        """
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The conditions of the reference table. The maximum length is 30. The maximum length of
        condition is 2048 characters.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The server time when reference table was created.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the reference table. The maximum length is 128 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the enterprise project ID of WAF reference table.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
        maximum length is 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the WAF reference table resource. If omitted,
        the provider-level region will be used. Changing this setting will push a new reference table.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the reference table, The options are `url`, `user-agent`, `ip`,
        `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class ReferenceTable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a WAF reference table resource within HuaweiCloud.

        > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
        used. The reference table resource can be used in Cloud Mode (professional version), Dedicated Mode and ELB Mode.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        enterprise_project_id = config.require_object("enterpriseProjectId")
        ref_table = huaweicloud.waf.ReferenceTable("refTable",
            type="url",
            enterprise_project_id=enterprise_project_id,
            conditions=[
                "/admin",
                "/manage",
            ])
        ```

        ## Import

        There are two ways to import WAF reference table state. * Using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Waf/referenceTable:ReferenceTable test <id>
        ```

         * Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Waf/referenceTable:ReferenceTable test <id>/<enterprise_project_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] conditions: The conditions of the reference table. The maximum length is 30. The maximum length of
               condition is 2048 characters.
        :param pulumi.Input[str] description: The description of the reference table. The maximum length is 128 characters.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID of WAF reference table.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
               maximum length is 64 characters.
        :param pulumi.Input[str] region: The region in which to create the WAF reference table resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new reference table.
        :param pulumi.Input[str] type: The type of the reference table, The options are `url`, `user-agent`, `ip`,
               `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReferenceTableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a WAF reference table resource within HuaweiCloud.

        > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
        used. The reference table resource can be used in Cloud Mode (professional version), Dedicated Mode and ELB Mode.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        enterprise_project_id = config.require_object("enterpriseProjectId")
        ref_table = huaweicloud.waf.ReferenceTable("refTable",
            type="url",
            enterprise_project_id=enterprise_project_id,
            conditions=[
                "/admin",
                "/manage",
            ])
        ```

        ## Import

        There are two ways to import WAF reference table state. * Using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Waf/referenceTable:ReferenceTable test <id>
        ```

         * Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Waf/referenceTable:ReferenceTable test <id>/<enterprise_project_id>
        ```

        :param str resource_name: The name of the resource.
        :param ReferenceTableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReferenceTableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReferenceTableArgs.__new__(ReferenceTableArgs)

            __props__.__dict__["conditions"] = conditions
            __props__.__dict__["description"] = description
            __props__.__dict__["enterprise_project_id"] = enterprise_project_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["creation_time"] = None
        super(ReferenceTable, __self__).__init__(
            'huaweicloud:Waf/referenceTable:ReferenceTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            conditions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enterprise_project_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ReferenceTable':
        """
        Get an existing ReferenceTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] conditions: The conditions of the reference table. The maximum length is 30. The maximum length of
               condition is 2048 characters.
        :param pulumi.Input[str] creation_time: The server time when reference table was created.
        :param pulumi.Input[str] description: The description of the reference table. The maximum length is 128 characters.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID of WAF reference table.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
               maximum length is 64 characters.
        :param pulumi.Input[str] region: The region in which to create the WAF reference table resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new reference table.
        :param pulumi.Input[str] type: The type of the reference table, The options are `url`, `user-agent`, `ip`,
               `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReferenceTableState.__new__(_ReferenceTableState)

        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["enterprise_project_id"] = enterprise_project_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["type"] = type
        return ReferenceTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The conditions of the reference table. The maximum length is 30. The maximum length of
        condition is 2048 characters.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The server time when reference table was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the reference table. The maximum length is 128 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the enterprise project ID of WAF reference table.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
        maximum length is 64 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the WAF reference table resource. If omitted,
        the provider-level region will be used. Changing this setting will push a new reference table.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the reference table, The options are `url`, `user-agent`, `ip`,
        `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.
        """
        return pulumi.get(self, "type")

