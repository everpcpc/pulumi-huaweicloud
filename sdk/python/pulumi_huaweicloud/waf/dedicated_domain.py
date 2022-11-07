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

__all__ = ['DedicatedDomainArgs', 'DedicatedDomain']

@pulumi.input_type
class DedicatedDomainArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 servers: pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]],
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 keep_policy: Optional[pulumi.Input[bool]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 protect_status: Optional[pulumi.Input[int]] = None,
                 proxy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DedicatedDomain resource.
        :param pulumi.Input[str] domain: Specifies the domain name to be protected. For example, www.example.com or
               *.example.com. Changing this creates a new domain.
        :param pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]] servers: The server configuration list of the domain. A maximum of 80 can be configured.
               The object structure is documented below.
        :param pulumi.Input[str] certificate_id: Specifies the certificate ID. This parameter is mandatory when `client_protocol`
               is set to HTTPS.
        :param pulumi.Input[bool] keep_policy: Specifies whether to retain the policy when deleting a domain name.
               Defaults to `true`.
        :param pulumi.Input[str] policy_id: Specifies the policy ID associated with the domain. If not specified, a new policy
               will be created automatically.
        :param pulumi.Input[int] protect_status: The protection status of domain, `0`: suspended, `1`: enabled.
               Default value is `1`.
        :param pulumi.Input[bool] proxy: Specifies whether a proxy is configured. Default value is `false`.
        :param pulumi.Input[str] region: The region in which to create the dedicated mode domain resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new domain.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "servers", servers)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if keep_policy is not None:
            pulumi.set(__self__, "keep_policy", keep_policy)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protect_status is not None:
            pulumi.set(__self__, "protect_status", protect_status)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Specifies the domain name to be protected. For example, www.example.com or
        *.example.com. Changing this creates a new domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]]:
        """
        The server configuration list of the domain. A maximum of 80 can be configured.
        The object structure is documented below.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]]):
        pulumi.set(self, "servers", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the certificate ID. This parameter is mandatory when `client_protocol`
        is set to HTTPS.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="keepPolicy")
    def keep_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to retain the policy when deleting a domain name.
        Defaults to `true`.
        """
        return pulumi.get(self, "keep_policy")

    @keep_policy.setter
    def keep_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_policy", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the policy ID associated with the domain. If not specified, a new policy
        will be created automatically.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="protectStatus")
    def protect_status(self) -> Optional[pulumi.Input[int]]:
        """
        The protection status of domain, `0`: suspended, `1`: enabled.
        Default value is `1`.
        """
        return pulumi.get(self, "protect_status")

    @protect_status.setter
    def protect_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "protect_status", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether a proxy is configured. Default value is `false`.
        """
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the dedicated mode domain resource. If omitted,
        the provider-level region will be used. Changing this setting will push a new domain.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _DedicatedDomainState:
    def __init__(__self__, *,
                 access_status: Optional[pulumi.Input[int]] = None,
                 alarm_page: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 cipher: Optional[pulumi.Input[str]] = None,
                 compliance_certification: Optional[pulumi.Input[Mapping[str, pulumi.Input[bool]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 keep_policy: Optional[pulumi.Input[bool]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 protect_status: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]]] = None,
                 tls: Optional[pulumi.Input[str]] = None,
                 traffic_identifier: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DedicatedDomain resources.
        :param pulumi.Input[int] access_status: Whether a domain name is connected to WAF. Valid values are:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] alarm_page: The alarm page of domain. Valid values are:
        :param pulumi.Input[str] certificate_id: Specifies the certificate ID. This parameter is mandatory when `client_protocol`
               is set to HTTPS.
        :param pulumi.Input[str] certificate_name: The name of the certificate used by the domain name.
        :param pulumi.Input[Mapping[str, pulumi.Input[bool]]] compliance_certification: The compliance certifications of the domain, values are:
        :param pulumi.Input[str] domain: Specifies the domain name to be protected. For example, www.example.com or
               *.example.com. Changing this creates a new domain.
        :param pulumi.Input[bool] keep_policy: Specifies whether to retain the policy when deleting a domain name.
               Defaults to `true`.
        :param pulumi.Input[str] policy_id: Specifies the policy ID associated with the domain. If not specified, a new policy
               will be created automatically.
        :param pulumi.Input[int] protect_status: The protection status of domain, `0`: suspended, `1`: enabled.
               Default value is `1`.
        :param pulumi.Input[str] protocol: The protocol type of the client. The options are `HTTP` and `HTTPS`.
        :param pulumi.Input[bool] proxy: Specifies whether a proxy is configured. Default value is `false`.
        :param pulumi.Input[str] region: The region in which to create the dedicated mode domain resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new domain.
        :param pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]] servers: The server configuration list of the domain. A maximum of 80 can be configured.
               The object structure is documented below.
        :param pulumi.Input[str] tls: The TLS configuration of domain.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] traffic_identifier: The traffic identifier of domain. Valid values are:
        """
        if access_status is not None:
            pulumi.set(__self__, "access_status", access_status)
        if alarm_page is not None:
            pulumi.set(__self__, "alarm_page", alarm_page)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if cipher is not None:
            pulumi.set(__self__, "cipher", cipher)
        if compliance_certification is not None:
            pulumi.set(__self__, "compliance_certification", compliance_certification)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if keep_policy is not None:
            pulumi.set(__self__, "keep_policy", keep_policy)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protect_status is not None:
            pulumi.set(__self__, "protect_status", protect_status)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)
        if traffic_identifier is not None:
            pulumi.set(__self__, "traffic_identifier", traffic_identifier)

    @property
    @pulumi.getter(name="accessStatus")
    def access_status(self) -> Optional[pulumi.Input[int]]:
        """
        Whether a domain name is connected to WAF. Valid values are:
        """
        return pulumi.get(self, "access_status")

    @access_status.setter
    def access_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "access_status", value)

    @property
    @pulumi.getter(name="alarmPage")
    def alarm_page(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The alarm page of domain. Valid values are:
        """
        return pulumi.get(self, "alarm_page")

    @alarm_page.setter
    def alarm_page(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "alarm_page", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the certificate ID. This parameter is mandatory when `client_protocol`
        is set to HTTPS.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the certificate used by the domain name.
        """
        return pulumi.get(self, "certificate_name")

    @certificate_name.setter
    def certificate_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_name", value)

    @property
    @pulumi.getter
    def cipher(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cipher")

    @cipher.setter
    def cipher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cipher", value)

    @property
    @pulumi.getter(name="complianceCertification")
    def compliance_certification(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[bool]]]]:
        """
        The compliance certifications of the domain, values are:
        """
        return pulumi.get(self, "compliance_certification")

    @compliance_certification.setter
    def compliance_certification(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[bool]]]]):
        pulumi.set(self, "compliance_certification", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the domain name to be protected. For example, www.example.com or
        *.example.com. Changing this creates a new domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="keepPolicy")
    def keep_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to retain the policy when deleting a domain name.
        Defaults to `true`.
        """
        return pulumi.get(self, "keep_policy")

    @keep_policy.setter
    def keep_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_policy", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the policy ID associated with the domain. If not specified, a new policy
        will be created automatically.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="protectStatus")
    def protect_status(self) -> Optional[pulumi.Input[int]]:
        """
        The protection status of domain, `0`: suspended, `1`: enabled.
        Default value is `1`.
        """
        return pulumi.get(self, "protect_status")

    @protect_status.setter
    def protect_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "protect_status", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol type of the client. The options are `HTTP` and `HTTPS`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether a proxy is configured. Default value is `false`.
        """
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the dedicated mode domain resource. If omitted,
        the provider-level region will be used. Changing this setting will push a new domain.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]]]:
        """
        The server configuration list of the domain. A maximum of 80 can be configured.
        The object structure is documented below.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DedicatedDomainServerArgs']]]]):
        pulumi.set(self, "servers", value)

    @property
    @pulumi.getter
    def tls(self) -> Optional[pulumi.Input[str]]:
        """
        The TLS configuration of domain.
        """
        return pulumi.get(self, "tls")

    @tls.setter
    def tls(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls", value)

    @property
    @pulumi.getter(name="trafficIdentifier")
    def traffic_identifier(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The traffic identifier of domain. Valid values are:
        """
        return pulumi.get(self, "traffic_identifier")

    @traffic_identifier.setter
    def traffic_identifier(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "traffic_identifier", value)


class DedicatedDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 keep_policy: Optional[pulumi.Input[bool]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 protect_status: Optional[pulumi.Input[int]] = None,
                 proxy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DedicatedDomainServerArgs']]]]] = None,
                 __props__=None):
        """
        Manages a dedicated mode domain resource within HuaweiCloud.

        > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
        used. The dedicated mode domain name resource can be used in Dedicated Mode and ELB Mode.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        certificated_id = config.require_object("certificatedId")
        vpc_id = config.require_object("vpcId")
        dedicated_engine_id = config.require_object("dedicatedEngineId")
        domain1 = huaweicloud.waf.DedicatedDomain("domain1",
            domain="www.example.com",
            certificate_id=huaweicloud_waf_certificate["certificate_1"]["id"],
            servers=[huaweicloud.waf.DedicatedDomainServerArgs(
                client_protocol="HTTPS",
                server_protocol="HTTP",
                address="192.168.1.100",
                port=8080,
                type="ipv4",
                vpc_id=vpc_id,
            )])
        ```

        ## Import

        Dedicated mode domain can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:Waf/dedicatedDomain:DedicatedDomain domain_1 69e9a86becb4424298cc6bdeacbf69d5
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: Specifies the certificate ID. This parameter is mandatory when `client_protocol`
               is set to HTTPS.
        :param pulumi.Input[str] domain: Specifies the domain name to be protected. For example, www.example.com or
               *.example.com. Changing this creates a new domain.
        :param pulumi.Input[bool] keep_policy: Specifies whether to retain the policy when deleting a domain name.
               Defaults to `true`.
        :param pulumi.Input[str] policy_id: Specifies the policy ID associated with the domain. If not specified, a new policy
               will be created automatically.
        :param pulumi.Input[int] protect_status: The protection status of domain, `0`: suspended, `1`: enabled.
               Default value is `1`.
        :param pulumi.Input[bool] proxy: Specifies whether a proxy is configured. Default value is `false`.
        :param pulumi.Input[str] region: The region in which to create the dedicated mode domain resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new domain.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DedicatedDomainServerArgs']]]] servers: The server configuration list of the domain. A maximum of 80 can be configured.
               The object structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a dedicated mode domain resource within HuaweiCloud.

        > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
        used. The dedicated mode domain name resource can be used in Dedicated Mode and ELB Mode.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        certificated_id = config.require_object("certificatedId")
        vpc_id = config.require_object("vpcId")
        dedicated_engine_id = config.require_object("dedicatedEngineId")
        domain1 = huaweicloud.waf.DedicatedDomain("domain1",
            domain="www.example.com",
            certificate_id=huaweicloud_waf_certificate["certificate_1"]["id"],
            servers=[huaweicloud.waf.DedicatedDomainServerArgs(
                client_protocol="HTTPS",
                server_protocol="HTTP",
                address="192.168.1.100",
                port=8080,
                type="ipv4",
                vpc_id=vpc_id,
            )])
        ```

        ## Import

        Dedicated mode domain can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:Waf/dedicatedDomain:DedicatedDomain domain_1 69e9a86becb4424298cc6bdeacbf69d5
        ```

        :param str resource_name: The name of the resource.
        :param DedicatedDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 keep_policy: Optional[pulumi.Input[bool]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 protect_status: Optional[pulumi.Input[int]] = None,
                 proxy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DedicatedDomainServerArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedDomainArgs.__new__(DedicatedDomainArgs)

            __props__.__dict__["certificate_id"] = certificate_id
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["keep_policy"] = keep_policy
            __props__.__dict__["policy_id"] = policy_id
            __props__.__dict__["protect_status"] = protect_status
            __props__.__dict__["proxy"] = proxy
            __props__.__dict__["region"] = region
            if servers is None and not opts.urn:
                raise TypeError("Missing required property 'servers'")
            __props__.__dict__["servers"] = servers
            __props__.__dict__["access_status"] = None
            __props__.__dict__["alarm_page"] = None
            __props__.__dict__["certificate_name"] = None
            __props__.__dict__["cipher"] = None
            __props__.__dict__["compliance_certification"] = None
            __props__.__dict__["protocol"] = None
            __props__.__dict__["tls"] = None
            __props__.__dict__["traffic_identifier"] = None
        super(DedicatedDomain, __self__).__init__(
            'huaweicloud:Waf/dedicatedDomain:DedicatedDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_status: Optional[pulumi.Input[int]] = None,
            alarm_page: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            certificate_id: Optional[pulumi.Input[str]] = None,
            certificate_name: Optional[pulumi.Input[str]] = None,
            cipher: Optional[pulumi.Input[str]] = None,
            compliance_certification: Optional[pulumi.Input[Mapping[str, pulumi.Input[bool]]]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            keep_policy: Optional[pulumi.Input[bool]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            protect_status: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            proxy: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DedicatedDomainServerArgs']]]]] = None,
            tls: Optional[pulumi.Input[str]] = None,
            traffic_identifier: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'DedicatedDomain':
        """
        Get an existing DedicatedDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] access_status: Whether a domain name is connected to WAF. Valid values are:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] alarm_page: The alarm page of domain. Valid values are:
        :param pulumi.Input[str] certificate_id: Specifies the certificate ID. This parameter is mandatory when `client_protocol`
               is set to HTTPS.
        :param pulumi.Input[str] certificate_name: The name of the certificate used by the domain name.
        :param pulumi.Input[Mapping[str, pulumi.Input[bool]]] compliance_certification: The compliance certifications of the domain, values are:
        :param pulumi.Input[str] domain: Specifies the domain name to be protected. For example, www.example.com or
               *.example.com. Changing this creates a new domain.
        :param pulumi.Input[bool] keep_policy: Specifies whether to retain the policy when deleting a domain name.
               Defaults to `true`.
        :param pulumi.Input[str] policy_id: Specifies the policy ID associated with the domain. If not specified, a new policy
               will be created automatically.
        :param pulumi.Input[int] protect_status: The protection status of domain, `0`: suspended, `1`: enabled.
               Default value is `1`.
        :param pulumi.Input[str] protocol: The protocol type of the client. The options are `HTTP` and `HTTPS`.
        :param pulumi.Input[bool] proxy: Specifies whether a proxy is configured. Default value is `false`.
        :param pulumi.Input[str] region: The region in which to create the dedicated mode domain resource. If omitted,
               the provider-level region will be used. Changing this setting will push a new domain.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DedicatedDomainServerArgs']]]] servers: The server configuration list of the domain. A maximum of 80 can be configured.
               The object structure is documented below.
        :param pulumi.Input[str] tls: The TLS configuration of domain.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] traffic_identifier: The traffic identifier of domain. Valid values are:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DedicatedDomainState.__new__(_DedicatedDomainState)

        __props__.__dict__["access_status"] = access_status
        __props__.__dict__["alarm_page"] = alarm_page
        __props__.__dict__["certificate_id"] = certificate_id
        __props__.__dict__["certificate_name"] = certificate_name
        __props__.__dict__["cipher"] = cipher
        __props__.__dict__["compliance_certification"] = compliance_certification
        __props__.__dict__["domain"] = domain
        __props__.__dict__["keep_policy"] = keep_policy
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["protect_status"] = protect_status
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["proxy"] = proxy
        __props__.__dict__["region"] = region
        __props__.__dict__["servers"] = servers
        __props__.__dict__["tls"] = tls
        __props__.__dict__["traffic_identifier"] = traffic_identifier
        return DedicatedDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessStatus")
    def access_status(self) -> pulumi.Output[int]:
        """
        Whether a domain name is connected to WAF. Valid values are:
        """
        return pulumi.get(self, "access_status")

    @property
    @pulumi.getter(name="alarmPage")
    def alarm_page(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The alarm page of domain. Valid values are:
        """
        return pulumi.get(self, "alarm_page")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the certificate ID. This parameter is mandatory when `client_protocol`
        is set to HTTPS.
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> pulumi.Output[str]:
        """
        The name of the certificate used by the domain name.
        """
        return pulumi.get(self, "certificate_name")

    @property
    @pulumi.getter
    def cipher(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cipher")

    @property
    @pulumi.getter(name="complianceCertification")
    def compliance_certification(self) -> pulumi.Output[Mapping[str, bool]]:
        """
        The compliance certifications of the domain, values are:
        """
        return pulumi.get(self, "compliance_certification")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Specifies the domain name to be protected. For example, www.example.com or
        *.example.com. Changing this creates a new domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="keepPolicy")
    def keep_policy(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to retain the policy when deleting a domain name.
        Defaults to `true`.
        """
        return pulumi.get(self, "keep_policy")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        Specifies the policy ID associated with the domain. If not specified, a new policy
        will be created automatically.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="protectStatus")
    def protect_status(self) -> pulumi.Output[int]:
        """
        The protection status of domain, `0`: suspended, `1`: enabled.
        Default value is `1`.
        """
        return pulumi.get(self, "protect_status")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol type of the client. The options are `HTTP` and `HTTPS`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def proxy(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether a proxy is configured. Default value is `false`.
        """
        return pulumi.get(self, "proxy")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the dedicated mode domain resource. If omitted,
        the provider-level region will be used. Changing this setting will push a new domain.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Output[Sequence['outputs.DedicatedDomainServer']]:
        """
        The server configuration list of the domain. A maximum of 80 can be configured.
        The object structure is documented below.
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter
    def tls(self) -> pulumi.Output[str]:
        """
        The TLS configuration of domain.
        """
        return pulumi.get(self, "tls")

    @property
    @pulumi.getter(name="trafficIdentifier")
    def traffic_identifier(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The traffic identifier of domain. Valid values are:
        """
        return pulumi.get(self, "traffic_identifier")
