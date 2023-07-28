// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the details of a specific IEC network ACL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const aclName = config.requireObject("aclName");
 * const firewall = huaweicloud.Iec.getNetwork_acl({
 *     name: aclName,
 * });
 * ```
 */
export function getNetwork_acl(args?: GetNetwork_aclArgs, opts?: pulumi.InvokeOptions): Promise<GetNetwork_aclResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Iec/getNetwork_acl:getNetwork_acl", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork_acl.
 */
export interface GetNetwork_aclArgs {
    /**
     * Specifies the ID of the IEC network ACL to retrieve.
     */
    id?: string;
    /**
     * Specifies the name if the IEC network ACL with a maximum of 64 characters.
     */
    name?: string;
}

/**
 * A collection of values returned by getNetwork_acl.
 */
export interface GetNetwork_aclResult {
    /**
     * The description of the IEC network ACL.
     */
    readonly description: string;
    readonly id: string;
    /**
     * A list of the IDs of ingress rules associated with the IEC network ACL.
     */
    readonly inboundRules: string[];
    readonly name: string;
    readonly networks: outputs.Iec.GetNetwork_aclNetwork[];
    /**
     * A list of the IDs of egress rules associated with the IEC network ACL.
     */
    readonly outboundRules: string[];
    /**
     * The status of the IEC network ACL.
     */
    readonly status: string;
}

export function getNetwork_aclOutput(args?: GetNetwork_aclOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetwork_aclResult> {
    return pulumi.output(args).apply(a => getNetwork_acl(a, opts))
}

/**
 * A collection of arguments for invoking getNetwork_acl.
 */
export interface GetNetwork_aclOutputArgs {
    /**
     * Specifies the ID of the IEC network ACL to retrieve.
     */
    id?: pulumi.Input<string>;
    /**
     * Specifies the name if the IEC network ACL with a maximum of 64 characters.
     */
    name?: pulumi.Input<string>;
}
