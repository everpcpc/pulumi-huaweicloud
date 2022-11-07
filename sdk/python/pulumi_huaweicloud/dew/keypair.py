# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeypairArgs', 'Keypair']

@pulumi.input_type
class KeypairArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Keypair resource.
        :param pulumi.Input[str] description: Specifies the description of key pair.
        :param pulumi.Input[str] encryption_type: Specifies encryption mode if manages the private key by HuaweiCloud.
               The options are as follows:
               - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
               - **kms**: KMS encryption mode.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] kms_key_name: Specifies the KMS key name to encrypt private keys.
               It's mandatory when the `encryption_type` is `kms`. Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. The name can contain a maximum of 64
               characters, including letters, digits, underscores (_) and hyphens (-).
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] scope: Specifies the scope of key pair. The options are as follows:
               - **account**: Tenant-level, available to all users under the same account.
               - **user**: User-level, only available to that user.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encryption_type is not None:
            pulumi.set(__self__, "encryption_type", encryption_type)
        if key_file is not None:
            pulumi.set(__self__, "key_file", key_file)
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of key pair.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies encryption mode if manages the private key by HuaweiCloud.
        The options are as follows:
        - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
        - **kms**: KMS encryption mode.
        """
        return pulumi.get(self, "encryption_type")

    @encryption_type.setter
    def encryption_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_type", value)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the path of the created private key.
        The private key file (**.pem**) is created only after the resource is created.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "key_file")

    @key_file.setter
    def key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_file", value)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the KMS key name to encrypt private keys.
        It's mandatory when the `encryption_type` is `kms`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "kms_key_name")

    @kms_key_name.setter
    def kms_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the keypair. The name can contain a maximum of 64
        characters, including letters, digits, underscores (_) and hyphens (-).
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the imported OpenSSH-formatted public key.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the keypair resource. If omitted, the
        provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the scope of key pair. The options are as follows:
        - **account**: Tenant-level, available to all users under the same account.
        - **user**: User-level, only available to that user.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class _KeypairState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 is_managed: Optional[pulumi.Input[bool]] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Keypair resources.
        :param pulumi.Input[str] created_at: The key pair creation time.
        :param pulumi.Input[str] description: Specifies the description of key pair.
        :param pulumi.Input[str] encryption_type: Specifies encryption mode if manages the private key by HuaweiCloud.
               The options are as follows:
               - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
               - **kms**: KMS encryption mode.
        :param pulumi.Input[str] fingerprint: Fingerprint information about an key pair.
        :param pulumi.Input[bool] is_managed: Whether the private key is managed by HuaweiCloud.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] kms_key_name: Specifies the KMS key name to encrypt private keys.
               It's mandatory when the `encryption_type` is `kms`. Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. The name can contain a maximum of 64
               characters, including letters, digits, underscores (_) and hyphens (-).
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] scope: Specifies the scope of key pair. The options are as follows:
               - **account**: Tenant-level, available to all users under the same account.
               - **user**: User-level, only available to that user.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encryption_type is not None:
            pulumi.set(__self__, "encryption_type", encryption_type)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if is_managed is not None:
            pulumi.set(__self__, "is_managed", is_managed)
        if key_file is not None:
            pulumi.set(__self__, "key_file", key_file)
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The key pair creation time.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of key pair.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies encryption mode if manages the private key by HuaweiCloud.
        The options are as follows:
        - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
        - **kms**: KMS encryption mode.
        """
        return pulumi.get(self, "encryption_type")

    @encryption_type.setter
    def encryption_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_type", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint information about an key pair.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter(name="isManaged")
    def is_managed(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the private key is managed by HuaweiCloud.
        """
        return pulumi.get(self, "is_managed")

    @is_managed.setter
    def is_managed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_managed", value)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the path of the created private key.
        The private key file (**.pem**) is created only after the resource is created.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "key_file")

    @key_file.setter
    def key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_file", value)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the KMS key name to encrypt private keys.
        It's mandatory when the `encryption_type` is `kms`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "kms_key_name")

    @kms_key_name.setter
    def kms_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the keypair. The name can contain a maximum of 64
        characters, including letters, digits, underscores (_) and hyphens (-).
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the imported OpenSSH-formatted public key.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the keypair resource. If omitted, the
        provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the scope of key pair. The options are as follows:
        - **account**: Tenant-level, available to all users under the same account.
        - **user**: User-level, only available to that user.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


class Keypair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a keypair resource within HuaweiCloud.

        By default, key pairs use the SSH-2 (RSA, 2048) algorithm for encryption and decryption.

        Keys imported support the following cryptographic algorithms:

         * RSA-1024
         * RSA-2048
         * RSA-4096

        ## Example Usage
        ### Create a new keypair which scope is Tenant-level and the private key is managed by HuaweiCloud

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test = huaweicloud.dew.Key("test", key_alias="kms_test")
        test_keypair = huaweicloud.dew.Keypair("test-keypair",
            scope="account",
            encryption_type="kms",
            kms_key_name=test.key_alias)
        ```
        ### Create a new keypair and export private key to current folder

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.dew.Keypair("test-keypair", key_file="private_key.pem")
        ```
        ### Import an existing keypair

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.dew.Keypair("test-keypair", public_key="ssh-rsa AAAAB3NzaC1yc2EAAAlJq5Pu+eizhou7nFFDxXofr2ySF8k/yuA9OnJdVF9Fbf85Z59CWNZBvcAT... root@terra-dev")
        ```

        ## Import

        Keypairs can be imported using the `name`, e.g.

        ```sh
         $ pulumi import huaweicloud:Dew/keypair:Keypair my-keypair test-keypair
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`encryption_type`, and `kms_key_name`. It is generally recommended running `terraform plan` after importing a key pair. You can then decide if changes should be applied to the key pair, or the resource definition should be updated to align with the key pair. Also you can ignore changes as below. resource "huaweicloud_kps_keypair" "test" {

         ...

         lifecycle {

         ignore_changes = [

         encryption_type, kms_key_name,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description of key pair.
        :param pulumi.Input[str] encryption_type: Specifies encryption mode if manages the private key by HuaweiCloud.
               The options are as follows:
               - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
               - **kms**: KMS encryption mode.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] kms_key_name: Specifies the KMS key name to encrypt private keys.
               It's mandatory when the `encryption_type` is `kms`. Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. The name can contain a maximum of 64
               characters, including letters, digits, underscores (_) and hyphens (-).
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] scope: Specifies the scope of key pair. The options are as follows:
               - **account**: Tenant-level, available to all users under the same account.
               - **user**: User-level, only available to that user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[KeypairArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a keypair resource within HuaweiCloud.

        By default, key pairs use the SSH-2 (RSA, 2048) algorithm for encryption and decryption.

        Keys imported support the following cryptographic algorithms:

         * RSA-1024
         * RSA-2048
         * RSA-4096

        ## Example Usage
        ### Create a new keypair which scope is Tenant-level and the private key is managed by HuaweiCloud

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test = huaweicloud.dew.Key("test", key_alias="kms_test")
        test_keypair = huaweicloud.dew.Keypair("test-keypair",
            scope="account",
            encryption_type="kms",
            kms_key_name=test.key_alias)
        ```
        ### Create a new keypair and export private key to current folder

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.dew.Keypair("test-keypair", key_file="private_key.pem")
        ```
        ### Import an existing keypair

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.dew.Keypair("test-keypair", public_key="ssh-rsa AAAAB3NzaC1yc2EAAAlJq5Pu+eizhou7nFFDxXofr2ySF8k/yuA9OnJdVF9Fbf85Z59CWNZBvcAT... root@terra-dev")
        ```

        ## Import

        Keypairs can be imported using the `name`, e.g.

        ```sh
         $ pulumi import huaweicloud:Dew/keypair:Keypair my-keypair test-keypair
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`encryption_type`, and `kms_key_name`. It is generally recommended running `terraform plan` after importing a key pair. You can then decide if changes should be applied to the key pair, or the resource definition should be updated to align with the key pair. Also you can ignore changes as below. resource "huaweicloud_kps_keypair" "test" {

         ...

         lifecycle {

         ignore_changes = [

         encryption_type, kms_key_name,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param KeypairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeypairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeypairArgs.__new__(KeypairArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["encryption_type"] = encryption_type
            __props__.__dict__["key_file"] = key_file
            __props__.__dict__["kms_key_name"] = kms_key_name
            __props__.__dict__["name"] = name
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["region"] = region
            __props__.__dict__["scope"] = scope
            __props__.__dict__["created_at"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["is_managed"] = None
        super(Keypair, __self__).__init__(
            'huaweicloud:Dew/keypair:Keypair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encryption_type: Optional[pulumi.Input[str]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            is_managed: Optional[pulumi.Input[bool]] = None,
            key_file: Optional[pulumi.Input[str]] = None,
            kms_key_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None) -> 'Keypair':
        """
        Get an existing Keypair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The key pair creation time.
        :param pulumi.Input[str] description: Specifies the description of key pair.
        :param pulumi.Input[str] encryption_type: Specifies encryption mode if manages the private key by HuaweiCloud.
               The options are as follows:
               - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
               - **kms**: KMS encryption mode.
        :param pulumi.Input[str] fingerprint: Fingerprint information about an key pair.
        :param pulumi.Input[bool] is_managed: Whether the private key is managed by HuaweiCloud.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] kms_key_name: Specifies the KMS key name to encrypt private keys.
               It's mandatory when the `encryption_type` is `kms`. Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. The name can contain a maximum of 64
               characters, including letters, digits, underscores (_) and hyphens (-).
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] scope: Specifies the scope of key pair. The options are as follows:
               - **account**: Tenant-level, available to all users under the same account.
               - **user**: User-level, only available to that user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeypairState.__new__(_KeypairState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["encryption_type"] = encryption_type
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["is_managed"] = is_managed
        __props__.__dict__["key_file"] = key_file
        __props__.__dict__["kms_key_name"] = kms_key_name
        __props__.__dict__["name"] = name
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["region"] = region
        __props__.__dict__["scope"] = scope
        return Keypair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The key pair creation time.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the description of key pair.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> pulumi.Output[str]:
        """
        Specifies encryption mode if manages the private key by HuaweiCloud.
        The options are as follows:
        - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
        - **kms**: KMS encryption mode.
        """
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint information about an key pair.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="isManaged")
    def is_managed(self) -> pulumi.Output[bool]:
        """
        Whether the private key is managed by HuaweiCloud.
        """
        return pulumi.get(self, "is_managed")

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> pulumi.Output[str]:
        """
        Specifies the path of the created private key.
        The private key file (**.pem**) is created only after the resource is created.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "key_file")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the KMS key name to encrypt private keys.
        It's mandatory when the `encryption_type` is `kms`. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies a unique name for the keypair. The name can contain a maximum of 64
        characters, including letters, digits, underscores (_) and hyphens (-).
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        Specifies the imported OpenSSH-formatted public key.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the keypair resource. If omitted, the
        provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Specifies the scope of key pair. The options are as follows:
        - **account**: Tenant-level, available to all users under the same account.
        - **user**: User-level, only available to that user.
        """
        return pulumi.get(self, "scope")
