// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of EIPs.
 *
 * ## Example Usage
 *
 * An example filter by name and tag
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const publicIp = config.requireObject("publicIp");
 * const eip = huaweicloud.Eip.getEips({
 *     publicIps: [publicIp],
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * export const eipIds = [eip.then(eip => eip.eips)].map(__item => __item?.id);
 * ```
 */
export function getEips(args?: GetEipsArgs, opts?: pulumi.InvokeOptions): Promise<GetEipsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Eip/getEips:getEips", {
        "enterpriseProjectId": args.enterpriseProjectId,
        "ids": args.ids,
        "ipVersion": args.ipVersion,
        "portIds": args.portIds,
        "publicIps": args.publicIps,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getEips.
 */
export interface GetEipsArgs {
    /**
     * Specifies the enterprise project ID which the desired EIP belongs to.
     */
    enterpriseProjectId?: string;
    /**
     * Specifies an array of one or more IDs of the desired EIP.
     */
    ids?: string[];
    /**
     * Specifies ip version of the desired EIP. The options are:
     * + `4`: IPv4.
     * + `6`: IPv6.
     */
    ipVersion?: number;
    /**
     * Specifies an array of one or more port ids which bound to the desired EIP.
     */
    portIds?: string[];
    /**
     * Specifies an array of one or more public ip addresses of the desired EIP.
     */
    publicIps?: string[];
    /**
     * Specifies the region in which to obtain the EIP. If omitted, the provider-level region
     * will be used.
     */
    region?: string;
    /**
     * Specifies the included key/value pairs which associated with the desired EIP.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getEips.
 */
export interface GetEipsResult {
    /**
     * Indicates a list of all EIPs found. Structure is documented below.
     */
    readonly eips: outputs.Eip.GetEipsEip[];
    /**
     * The the enterprise project ID of the EIP.
     */
    readonly enterpriseProjectId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The ip version of the EIP.
     */
    readonly ipVersion?: number;
    readonly portIds?: string[];
    readonly publicIps?: string[];
    readonly region: string;
    /**
     * The key/value pairs which associated with the EIP.
     */
    readonly tags?: {[key: string]: string};
}

export function getEipsOutput(args?: GetEipsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEipsResult> {
    return pulumi.output(args).apply(a => getEips(a, opts))
}

/**
 * A collection of arguments for invoking getEips.
 */
export interface GetEipsOutputArgs {
    /**
     * Specifies the enterprise project ID which the desired EIP belongs to.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more IDs of the desired EIP.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies ip version of the desired EIP. The options are:
     * + `4`: IPv4.
     * + `6`: IPv6.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Specifies an array of one or more port ids which bound to the desired EIP.
     */
    portIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies an array of one or more public ip addresses of the desired EIP.
     */
    publicIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the region in which to obtain the EIP. If omitted, the provider-level region
     * will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the included key/value pairs which associated with the desired EIP.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
