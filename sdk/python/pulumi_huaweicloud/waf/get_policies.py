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
    'GetPoliciesResult',
    'AwaitableGetPoliciesResult',
    'get_policies',
    'get_policies_output',
]

@pulumi.output_type
class GetPoliciesResult:
    """
    A collection of values returned by getPolicies.
    """
    def __init__(__self__, enterprise_project_id=None, id=None, name=None, policies=None, region=None):
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The WAF policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GetPoliciesPolicyResult']:
        """
        A list of WAF policies.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetPoliciesResult(GetPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoliciesResult(
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            name=self.name,
            policies=self.policies,
            region=self.region)


def get_policies(enterprise_project_id: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPoliciesResult:
    """
    Use this data source to get a list of WAF policies.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    policy_name = config.require_object("policyName")
    enterprise_project_id = config.require_object("enterpriseProjectId")
    policies = huaweicloud.Waf.get_policies(name=policy_name,
        enterprise_project_id=enterprise_project_id)
    ```


    :param str enterprise_project_id: Specifies the enterprise project ID of WAF policies.
    :param str name: Policy name used for matching. The value is case sensitive and supports fuzzy matching.
    :param str region: The region in which to obtain the WAF policies. If omitted, the provider-level region
           will be used.
    """
    __args__ = dict()
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Waf/getPolicies:getPolicies', __args__, opts=opts, typ=GetPoliciesResult).value

    return AwaitableGetPoliciesResult(
        enterprise_project_id=__ret__.enterprise_project_id,
        id=__ret__.id,
        name=__ret__.name,
        policies=__ret__.policies,
        region=__ret__.region)


@_utilities.lift_output_func(get_policies)
def get_policies_output(enterprise_project_id: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPoliciesResult]:
    """
    Use this data source to get a list of WAF policies.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    policy_name = config.require_object("policyName")
    enterprise_project_id = config.require_object("enterpriseProjectId")
    policies = huaweicloud.Waf.get_policies(name=policy_name,
        enterprise_project_id=enterprise_project_id)
    ```


    :param str enterprise_project_id: Specifies the enterprise project ID of WAF policies.
    :param str name: Policy name used for matching. The value is case sensitive and supports fuzzy matching.
    :param str region: The region in which to obtain the WAF policies. If omitted, the provider-level region
           will be used.
    """
    ...
