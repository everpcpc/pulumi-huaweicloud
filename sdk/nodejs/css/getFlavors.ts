// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get available flavors of HuaweiCloud CSS node instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.Css.getFlavors({
 *     memory: 32,
 *     type: "ess",
 *     vcpus: 4,
 *     version: "7.9.3",
 * }));
 * ```
 */
export function getFlavors(args?: GetFlavorsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlavorsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Css/getFlavors:getFlavors", {
        "memory": args.memory,
        "name": args.name,
        "region": args.region,
        "type": args.type,
        "vcpus": args.vcpus,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsArgs {
    /**
     * Specifies the memory size(GB) in the CSS flavor.
     */
    memory?: number;
    /**
     * Specifies the name of the CSS flavor.
     */
    name?: string;
    /**
     * Specifies the region in which to obtain the CSS flavors. If omitted, the
     * provider-level region will be used.
     */
    region?: string;
    /**
     * Specifies the node instance type. The options are `ess`, `ess-cold`, `ess-master`
     * and `ess-client`.
     */
    type?: string;
    /**
     * Specifies the number of vCPUs in the CSS flavor.
     */
    vcpus?: number;
    /**
     * Specifies the engine version. The options are `5.5.1`, `6.2.3`, `6.5.4`, `7.1.1`,
     * `7.6.2` and `7.9.3`.
     */
    version?: string;
}

/**
 * A collection of values returned by getFlavors.
 */
export interface GetFlavorsResult {
    /**
     * Indicates the flavors information. Structure is documented below.
     */
    readonly flavors: outputs.Css.GetFlavorsFlavor[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The memory size in GB.
     */
    readonly memory?: number;
    /**
     * The name of the CSS flavor. It is referenced by `node_config.flavor` in `huaweicloud.Css.Cluster`.
     */
    readonly name?: string;
    /**
     * The region where the node resides.
     */
    readonly region: string;
    /**
     * The node instance type.
     */
    readonly type?: string;
    /**
     * The number of vCPUs.
     */
    readonly vcpus?: number;
    /**
     * The engine version.
     */
    readonly version?: string;
}

export function getFlavorsOutput(args?: GetFlavorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlavorsResult> {
    return pulumi.output(args).apply(a => getFlavors(a, opts))
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsOutputArgs {
    /**
     * Specifies the memory size(GB) in the CSS flavor.
     */
    memory?: pulumi.Input<number>;
    /**
     * Specifies the name of the CSS flavor.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to obtain the CSS flavors. If omitted, the
     * provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the node instance type. The options are `ess`, `ess-cold`, `ess-master`
     * and `ess-client`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the number of vCPUs in the CSS flavor.
     */
    vcpus?: pulumi.Input<number>;
    /**
     * Specifies the engine version. The options are `5.5.1`, `6.2.3`, `6.5.4`, `7.1.1`,
     * `7.6.2` and `7.9.3`.
     */
    version?: pulumi.Input<string>;
}