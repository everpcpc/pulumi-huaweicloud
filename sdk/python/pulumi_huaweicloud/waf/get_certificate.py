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
    'GetCertificateResult',
    'AwaitableGetCertificateResult',
    'get_certificate',
    'get_certificate_output',
]

@pulumi.output_type
class GetCertificateResult:
    """
    A collection of values returned by getCertificate.
    """
    def __init__(__self__, expiration=None, expire_status=None, id=None, name=None, region=None):
        if expiration and not isinstance(expiration, str):
            raise TypeError("Expected argument 'expiration' to be a str")
        pulumi.set(__self__, "expiration", expiration)
        if expire_status and not isinstance(expire_status, int):
            raise TypeError("Expected argument 'expire_status' to be a int")
        pulumi.set(__self__, "expire_status", expire_status)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def expiration(self) -> str:
        """
        Indicates the time when the certificate expires.
        """
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter(name="expireStatus")
    def expire_status(self) -> Optional[int]:
        return pulumi.get(self, "expire_status")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetCertificateResult(GetCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateResult(
            expiration=self.expiration,
            expire_status=self.expire_status,
            id=self.id,
            name=self.name,
            region=self.region)


def get_certificate(expire_status: Optional[int] = None,
                    name: Optional[str] = None,
                    region: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateResult:
    """
    Get the certificate in the WAF, including the one pushed from SCM.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    certificate1 = huaweicloud.Waf.get_certificate(name="certificate name")
    domain1 = huaweicloud.waf.Domain("domain1",
        domain="www.domainname.com",
        certificate_id=certificate1.id,
        certificate_name=certificate1.name,
        keep_policy=False,
        proxy=False,
        servers=[huaweicloud.waf.DomainServerArgs(
            client_protocol="HTTPS",
            server_protocol="HTTP",
            address="192.168.10.1",
            port=8080,
        )])
    ```


    :param int expire_status: The expire status of certificate. Defaults is `0`. The value can be:
           + `0`: not expire
           + `1`: has expired
           + `2`: wil expired soon
    :param str name: The name of certificate. The value is case sensitive and supports fuzzy matching.
    :param str region: The region in which to obtain the WAF. If omitted, the provider-level region will be
           used.
    """
    __args__ = dict()
    __args__['expireStatus'] = expire_status
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Waf/getCertificate:getCertificate', __args__, opts=opts, typ=GetCertificateResult).value

    return AwaitableGetCertificateResult(
        expiration=__ret__.expiration,
        expire_status=__ret__.expire_status,
        id=__ret__.id,
        name=__ret__.name,
        region=__ret__.region)


@_utilities.lift_output_func(get_certificate)
def get_certificate_output(expire_status: Optional[pulumi.Input[Optional[int]]] = None,
                           name: Optional[pulumi.Input[str]] = None,
                           region: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCertificateResult]:
    """
    Get the certificate in the WAF, including the one pushed from SCM.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    certificate1 = huaweicloud.Waf.get_certificate(name="certificate name")
    domain1 = huaweicloud.waf.Domain("domain1",
        domain="www.domainname.com",
        certificate_id=certificate1.id,
        certificate_name=certificate1.name,
        keep_policy=False,
        proxy=False,
        servers=[huaweicloud.waf.DomainServerArgs(
            client_protocol="HTTPS",
            server_protocol="HTTP",
            address="192.168.10.1",
            port=8080,
        )])
    ```


    :param int expire_status: The expire status of certificate. Defaults is `0`. The value can be:
           + `0`: not expire
           + `1`: has expired
           + `2`: wil expired soon
    :param str name: The name of certificate. The value is case sensitive and supports fuzzy matching.
    :param str region: The region in which to obtain the WAF. If omitted, the provider-level region will be
           used.
    """
    ...