// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the details of a specified compute instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const ecsName = config.requireObject("ecsName");
 * const demo = huaweicloud.Ecs.getInstance({
 *     name: ecsName,
 * });
 * ```
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Ecs/getInstance:getInstance", {
        "enterpriseProjectId": args.enterpriseProjectId,
        "fixedIpV4": args.fixedIpV4,
        "flavorId": args.flavorId,
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * Specifies the enterprise project id.
     */
    enterpriseProjectId?: string;
    /**
     * Specifies the IPv4 addresses of the ECS.
     */
    fixedIpV4?: string;
    /**
     * Specifies the flavor ID.
     */
    flavorId?: string;
    /**
     * Specifies the ECS ID.
     */
    instanceId?: string;
    /**
     * Specifies the ECS name, which can be queried with a regular expression.
     */
    name?: string;
    /**
     * The region in which to obtain the instance.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    /**
     * The availability zone where the instance is located.
     */
    readonly availabilityZone: string;
    readonly enterpriseProjectId: string;
    /**
     * The fixed IPv4 address of the instance on this network.
     */
    readonly fixedIpV4?: string;
    readonly flavorId: string;
    /**
     * The flavor name of the instance.
     */
    readonly flavorName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The image ID of the instance.
     */
    readonly imageId: string;
    /**
     * The image name of the instance.
     */
    readonly imageName: string;
    readonly instanceId?: string;
    /**
     * The key pair that is used to authenticate the instance.
     */
    readonly keyPair: string;
    readonly name: string;
    /**
     * An array of one or more networks to attach to the instance.
     * The network object structure is documented below.
     */
    readonly networks: outputs.Ecs.GetInstanceNetwork[];
    /**
     * The EIP address that is associted to the instance.
     */
    readonly publicIp: string;
    readonly region: string;
    /**
     * The scheduler with hints on how the instance should be launched.
     * The scheduler hints structure is documented below.
     */
    readonly schedulerHints: outputs.Ecs.GetInstanceSchedulerHint[];
    /**
     * An array of one or more security group IDs to associate with the instance.
     */
    readonly securityGroupIds: string[];
    /**
     * An array of one or more security groups to associate with the instance.
     */
    readonly securityGroups: string[];
    /**
     * The status of the instance.
     */
    readonly status: string;
    /**
     * The system disk voume ID.
     */
    readonly systemDiskId: string;
    /**
     * The key/value pairs to associate with the instance.
     */
    readonly tags: {[key: string]: string};
    /**
     * The user data (information after encoding) configured during instance creation.
     */
    readonly userData: string;
    /**
     * An array of one or more disks to attach to the instance.
     * The volume attached object structure is documented below.
     */
    readonly volumeAttacheds: outputs.Ecs.GetInstanceVolumeAttached[];
}

export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * Specifies the enterprise project id.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the IPv4 addresses of the ECS.
     */
    fixedIpV4?: pulumi.Input<string>;
    /**
     * Specifies the flavor ID.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * Specifies the ECS ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the ECS name, which can be queried with a regular expression.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the instance.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
}
