# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 app_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_action: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the application
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] app_codes: Specifies an array of one or more application codes that the application has.  
               Up to five application codes can be created.
               The valid length of each application code is limited from can contain `64` to `180`.
               The application code must start with a letter, digit, plus sign (+) or slash (/).
               Only letters, digits and following special special characters are allowed: `!@#$%+-_/=`.
        :param pulumi.Input[str] description: Specifies the application description.  
               The description contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
        :param pulumi.Input[str] name: Specifies the application name.  
               The valid length is limited from can contain `3` to `64`, only Chinese and English letters, digits and hyphens (-)
               are allowed.
               The name must start with a Chinese or English letter.
        :param pulumi.Input[str] region: Specifies the region where the application is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[str] secret_action: Specifies the secret action to be done for the application.  
               The valid action is **RESET**.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if app_codes is not None:
            pulumi.set(__self__, "app_codes", app_codes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_action is not None:
            pulumi.set(__self__, "secret_action", secret_action)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the dedicated instance to which the application
        belongs.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="appCodes")
    def app_codes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more application codes that the application has.  
        Up to five application codes can be created.
        The valid length of each application code is limited from can contain `64` to `180`.
        The application code must start with a letter, digit, plus sign (+) or slash (/).
        Only letters, digits and following special special characters are allowed: `!@#$%+-_/=`.
        """
        return pulumi.get(self, "app_codes")

    @app_codes.setter
    def app_codes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "app_codes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the application description.  
        The description contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the application name.  
        The valid length is limited from can contain `3` to `64`, only Chinese and English letters, digits and hyphens (-)
        are allowed.
        The name must start with a Chinese or English letter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the application is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretAction")
    def secret_action(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the secret action to be done for the application.  
        The valid action is **RESET**.
        """
        return pulumi.get(self, "secret_action")

    @secret_action.setter
    def secret_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_action", value)


@pulumi.input_type
class _ApplicationState:
    def __init__(__self__, *,
                 app_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 app_key: Optional[pulumi.Input[str]] = None,
                 app_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 registration_time: Optional[pulumi.Input[str]] = None,
                 secret_action: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Application resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] app_codes: Specifies an array of one or more application codes that the application has.  
               Up to five application codes can be created.
               The valid length of each application code is limited from can contain `64` to `180`.
               The application code must start with a letter, digit, plus sign (+) or slash (/).
               Only letters, digits and following special special characters are allowed: `!@#$%+-_/=`.
        :param pulumi.Input[str] app_key: App key.
        :param pulumi.Input[str] app_secret: App secret.
        :param pulumi.Input[str] description: Specifies the application description.  
               The description contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the application
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the application name.  
               The valid length is limited from can contain `3` to `64`, only Chinese and English letters, digits and hyphens (-)
               are allowed.
               The name must start with a Chinese or English letter.
        :param pulumi.Input[str] region: Specifies the region where the application is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[str] registration_time: the registration time.
        :param pulumi.Input[str] secret_action: Specifies the secret action to be done for the application.  
               The valid action is **RESET**.
        :param pulumi.Input[str] updated_at: The latest update time of the application.
        """
        if app_codes is not None:
            pulumi.set(__self__, "app_codes", app_codes)
        if app_key is not None:
            pulumi.set(__self__, "app_key", app_key)
        if app_secret is not None:
            pulumi.set(__self__, "app_secret", app_secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if registration_time is not None:
            pulumi.set(__self__, "registration_time", registration_time)
        if secret_action is not None:
            pulumi.set(__self__, "secret_action", secret_action)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="appCodes")
    def app_codes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more application codes that the application has.  
        Up to five application codes can be created.
        The valid length of each application code is limited from can contain `64` to `180`.
        The application code must start with a letter, digit, plus sign (+) or slash (/).
        Only letters, digits and following special special characters are allowed: `!@#$%+-_/=`.
        """
        return pulumi.get(self, "app_codes")

    @app_codes.setter
    def app_codes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "app_codes", value)

    @property
    @pulumi.getter(name="appKey")
    def app_key(self) -> Optional[pulumi.Input[str]]:
        """
        App key.
        """
        return pulumi.get(self, "app_key")

    @app_key.setter
    def app_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_key", value)

    @property
    @pulumi.getter(name="appSecret")
    def app_secret(self) -> Optional[pulumi.Input[str]]:
        """
        App secret.
        """
        return pulumi.get(self, "app_secret")

    @app_secret.setter
    def app_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_secret", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the application description.  
        The description contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the dedicated instance to which the application
        belongs.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the application name.  
        The valid length is limited from can contain `3` to `64`, only Chinese and English letters, digits and hyphens (-)
        are allowed.
        The name must start with a Chinese or English letter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the application is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="registrationTime")
    def registration_time(self) -> Optional[pulumi.Input[str]]:
        """
        the registration time.
        """
        return pulumi.get(self, "registration_time")

    @registration_time.setter
    def registration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registration_time", value)

    @property
    @pulumi.getter(name="secretAction")
    def secret_action(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the secret action to be done for the application.  
        The valid action is **RESET**.
        """
        return pulumi.get(self, "secret_action")

    @secret_action.setter
    def secret_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_action", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The latest update time of the application.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_action: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an APIG application resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        app_name = config.require_object("appName")
        app_code = config.require_object("appCode")
        test = huaweicloud.dedicated_apig.Application("test",
            instance_id=instance_id,
            description="Created by script",
            app_codes=[app_code])
        ```

        ## Import

        Applications can be imported using their `id` and the ID of the related dedicated instance, separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/application:Application test <instance_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] app_codes: Specifies an array of one or more application codes that the application has.  
               Up to five application codes can be created.
               The valid length of each application code is limited from can contain `64` to `180`.
               The application code must start with a letter, digit, plus sign (+) or slash (/).
               Only letters, digits and following special special characters are allowed: `!@#$%+-_/=`.
        :param pulumi.Input[str] description: Specifies the application description.  
               The description contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the application
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the application name.  
               The valid length is limited from can contain `3` to `64`, only Chinese and English letters, digits and hyphens (-)
               are allowed.
               The name must start with a Chinese or English letter.
        :param pulumi.Input[str] region: Specifies the region where the application is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[str] secret_action: Specifies the secret action to be done for the application.  
               The valid action is **RESET**.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an APIG application resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        app_name = config.require_object("appName")
        app_code = config.require_object("appCode")
        test = huaweicloud.dedicated_apig.Application("test",
            instance_id=instance_id,
            description="Created by script",
            app_codes=[app_code])
        ```

        ## Import

        Applications can be imported using their `id` and the ID of the related dedicated instance, separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/application:Application test <instance_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_action: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["app_codes"] = app_codes
            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["secret_action"] = secret_action
            __props__.__dict__["app_key"] = None
            __props__.__dict__["app_secret"] = None
            __props__.__dict__["registration_time"] = None
            __props__.__dict__["updated_at"] = None
        super(Application, __self__).__init__(
            'huaweicloud:DedicatedApig/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            app_key: Optional[pulumi.Input[str]] = None,
            app_secret: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            registration_time: Optional[pulumi.Input[str]] = None,
            secret_action: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] app_codes: Specifies an array of one or more application codes that the application has.  
               Up to five application codes can be created.
               The valid length of each application code is limited from can contain `64` to `180`.
               The application code must start with a letter, digit, plus sign (+) or slash (/).
               Only letters, digits and following special special characters are allowed: `!@#$%+-_/=`.
        :param pulumi.Input[str] app_key: App key.
        :param pulumi.Input[str] app_secret: App secret.
        :param pulumi.Input[str] description: Specifies the application description.  
               The description contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the application
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the application name.  
               The valid length is limited from can contain `3` to `64`, only Chinese and English letters, digits and hyphens (-)
               are allowed.
               The name must start with a Chinese or English letter.
        :param pulumi.Input[str] region: Specifies the region where the application is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[str] registration_time: the registration time.
        :param pulumi.Input[str] secret_action: Specifies the secret action to be done for the application.  
               The valid action is **RESET**.
        :param pulumi.Input[str] updated_at: The latest update time of the application.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationState.__new__(_ApplicationState)

        __props__.__dict__["app_codes"] = app_codes
        __props__.__dict__["app_key"] = app_key
        __props__.__dict__["app_secret"] = app_secret
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["registration_time"] = registration_time
        __props__.__dict__["secret_action"] = secret_action
        __props__.__dict__["updated_at"] = updated_at
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appCodes")
    def app_codes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies an array of one or more application codes that the application has.  
        Up to five application codes can be created.
        The valid length of each application code is limited from can contain `64` to `180`.
        The application code must start with a letter, digit, plus sign (+) or slash (/).
        Only letters, digits and following special special characters are allowed: `!@#$%+-_/=`.
        """
        return pulumi.get(self, "app_codes")

    @property
    @pulumi.getter(name="appKey")
    def app_key(self) -> pulumi.Output[str]:
        """
        App key.
        """
        return pulumi.get(self, "app_key")

    @property
    @pulumi.getter(name="appSecret")
    def app_secret(self) -> pulumi.Output[str]:
        """
        App secret.
        """
        return pulumi.get(self, "app_secret")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Specifies the application description.  
        The description contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the dedicated instance to which the application
        belongs.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the application name.  
        The valid length is limited from can contain `3` to `64`, only Chinese and English letters, digits and hyphens (-)
        are allowed.
        The name must start with a Chinese or English letter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region where the application is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="registrationTime")
    def registration_time(self) -> pulumi.Output[str]:
        """
        the registration time.
        """
        return pulumi.get(self, "registration_time")

    @property
    @pulumi.getter(name="secretAction")
    def secret_action(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the secret action to be done for the application.  
        The valid action is **RESET**.
        """
        return pulumi.get(self, "secret_action")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The latest update time of the application.
        """
        return pulumi.get(self, "updated_at")

