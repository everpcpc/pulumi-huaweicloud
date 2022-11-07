// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const kmsAdm = pulumi.output(huaweicloud.Iam.getRole({
 *     displayName: "KMS Administrator",
 * }));
 * ```
 */
export function getRole(args?: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Iam/getRole:getRole", {
        "displayName": args.displayName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleArgs {
    /**
     * Specifies the display name of the role displayed on the console.
     * Required if `name` is empty.
     */
    displayName?: string;
    /**
     * Specifies the name of the role for internal use.
     * Required if `displayName` is empty.
     */
    name?: string;
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    /**
     * The service catalog of the policy.
     */
    readonly catalog: string;
    /**
     * The description of the policy.
     */
    readonly description: string;
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The content of the policy.
     */
    readonly policy: string;
    /**
     * The display mode of the policy.
     */
    readonly type: string;
}

export function getRoleOutput(args?: GetRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRoleResult> {
    return pulumi.output(args).apply(a => getRole(a, opts))
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleOutputArgs {
    /**
     * Specifies the display name of the role displayed on the console.
     * Required if `name` is empty.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies the name of the role for internal use.
     * Required if `displayName` is empty.
     */
    name?: pulumi.Input<string>;
}