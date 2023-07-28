// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the details of a specific IEC VPC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const vpcName = config.requireObject("vpcName");
 * const myVpc = huaweicloud.Iec.getVpc({
 *     name: vpcName,
 * });
 * ```
 */
export function getVpc(args?: GetVpcArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Iec/getVpc:getVpc", {
        "id": args.id,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpc.
 */
export interface GetVpcArgs {
    /**
     * Specifies the ID of the IEC VPC to retrieve.
     */
    id?: string;
    /**
     * Specifies the name of the IEC VPC. The name can contain a maximum of 64 characters. Only
     * letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
     */
    name?: string;
    /**
     * Specifies the region in which to obtain the vpc. If omitted, the provider-level region
     * will be used.
     */
    region?: string;
}

/**
 * A collection of values returned by getVpc.
 */
export interface GetVpcResult {
    /**
     * Indicates the IP address range for the VPC.
     */
    readonly cidr: string;
    readonly id: string;
    /**
     * Indicates the mode of the IEC VPC. Possible values are *SYSTEM* and *CUSTOMER*.
     */
    readonly mode: string;
    readonly name: string;
    readonly region: string;
    /**
     * Indicates the number of subnets.
     */
    readonly subnetNum: number;
}

export function getVpcOutput(args?: GetVpcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcResult> {
    return pulumi.output(args).apply(a => getVpc(a, opts))
}

/**
 * A collection of arguments for invoking getVpc.
 */
export interface GetVpcOutputArgs {
    /**
     * Specifies the ID of the IEC VPC to retrieve.
     */
    id?: pulumi.Input<string>;
    /**
     * Specifies the name of the IEC VPC. The name can contain a maximum of 64 characters. Only
     * letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to obtain the vpc. If omitted, the provider-level region
     * will be used.
     */
    region?: pulumi.Input<string>;
}
