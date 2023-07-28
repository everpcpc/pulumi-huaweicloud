// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * !> **WARNING:** It has been deprecated.
 *
 * Manages a VPC channel resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const channelName = config.requireObject("channelName");
 * const ecsId1 = config.requireObject("ecsId1");
 * const ecsId2 = config.requireObject("ecsId2");
 * const test = new huaweicloud.dedicatedapig.VpcChannel("test", {
 *     instanceId: instanceId,
 *     port: 8080,
 *     protocol: "HTTPS",
 *     path: "/",
 *     httpCode: "201,202,203",
 *     members: [
 *         {
 *             id: ecsId1,
 *             weight: 30,
 *         },
 *         {
 *             id: ecsId2,
 *             weight: 70,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * VPC Channels can be imported using their `name` and the ID of the related dedicated instance, separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedApig/vpcChannel:VpcChannel test <instance_id>/<name>
 * ```
 */
export class VpcChannel extends pulumi.CustomResource {
    /**
     * Get an existing VpcChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcChannelState, opts?: pulumi.CustomResourceOptions): VpcChannel {
        return new VpcChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedApig/vpcChannel:VpcChannel';

    /**
     * Returns true if the given object is an instance of VpcChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcChannel.__pulumiType;
    }

    /**
     * Specifies the distribution algorithm.  
     * The valid types are **WRR**, **WLC**, **SH** and **URI hashing**, defaults to **WRR**.
     */
    public readonly algorithm!: pulumi.Output<string | undefined>;
    /**
     * The time when the VPC channel was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Specifies the healthy threshold, which refers to the number of consecutive
     * successful checks required for a backend server to be considered healthy.
     * The valid value ranges from `2` to `10`, defaults to `2`.
     */
    public readonly healthyThreshold!: pulumi.Output<number | undefined>;
    /**
     * Specifies the response codes for determining a successful HTTP response.  
     * The valid value ranges from `100` to `599` and the valid formats are as follows:
     * + The multiple values, for example, **200,201,202**.
     * + The range, for example, **200-299**.
     * + Both multiple values and ranges, for example, **201,202,210-299**.
     */
    public readonly httpCode!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the dedicated instance to which the VPC channel
     * belongs.
     * Changing this will create a new resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the interval between consecutive checks, in second.  
     * The valid value ranges from `5` to `300`, defaults to `10`.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * Specifies the member type of the VPC channel.  
     * The valid types are **ECS** and **EIP**, defaults to **ECS**.
     */
    public readonly memberType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the configuration of the backend servers that bind the VPC channel.  
     * The object structure is documented below.
     */
    public readonly members!: pulumi.Output<outputs.DedicatedApig.VpcChannelMember[]>;
    /**
     * Specifies the name of the VPC channel.  
     * The valid length is limited from `3` to `64`, only chinese and english letters, digits, hyphens (-), underscores (_)
     * and dots (.) are allowed.
     * The name must start with a chinese or english letter, and the Chinese characters must be in **UTF-8** or **Unicode**
     * format.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the destination path for health checks.  
     * Required if the `protocol` is **HTTP** or **HTTPS**.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * Specifies the host port of the VPC channel.  
     * The valid value ranges from `1` to `65,535`.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Specifies the protocol for performing health checks on backend servers in the VPC
     * channel.
     * The valid values are **TCP**, **HTTP** and **HTTPS**, defaults to **TCP**.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies the region where the VPC channel is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The current status of the VPC channel, supports **Normal** and **Abnormal**.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the timeout for determining whether a health check fails, in second.  
     * The value must be less than the value of the time `interval`.
     * The valid value ranges from `2` to `30`, defaults to `5`.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Specifies the unhealthy threshold, which refers to the number of consecutive
     * failed checks required for a backend server to be considered unhealthy.
     * The valid value ranges from `2` to `10`, defaults to `5`.
     */
    public readonly unhealthyThreshold!: pulumi.Output<number | undefined>;

    /**
     * Create a VpcChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcChannelArgs | VpcChannelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcChannelState | undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["healthyThreshold"] = state ? state.healthyThreshold : undefined;
            resourceInputs["httpCode"] = state ? state.httpCode : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["memberType"] = state ? state.memberType : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["unhealthyThreshold"] = state ? state.unhealthyThreshold : undefined;
        } else {
            const args = argsOrState as VpcChannelArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            resourceInputs["httpCode"] = args ? args.httpCode : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["memberType"] = args ? args.memberType : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcChannel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcChannel resources.
 */
export interface VpcChannelState {
    /**
     * Specifies the distribution algorithm.  
     * The valid types are **WRR**, **WLC**, **SH** and **URI hashing**, defaults to **WRR**.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * The time when the VPC channel was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Specifies the healthy threshold, which refers to the number of consecutive
     * successful checks required for a backend server to be considered healthy.
     * The valid value ranges from `2` to `10`, defaults to `2`.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * Specifies the response codes for determining a successful HTTP response.  
     * The valid value ranges from `100` to `599` and the valid formats are as follows:
     * + The multiple values, for example, **200,201,202**.
     * + The range, for example, **200-299**.
     * + Both multiple values and ranges, for example, **201,202,210-299**.
     */
    httpCode?: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated instance to which the VPC channel
     * belongs.
     * Changing this will create a new resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the interval between consecutive checks, in second.  
     * The valid value ranges from `5` to `300`, defaults to `10`.
     */
    interval?: pulumi.Input<number>;
    /**
     * Specifies the member type of the VPC channel.  
     * The valid types are **ECS** and **EIP**, defaults to **ECS**.
     */
    memberType?: pulumi.Input<string>;
    /**
     * Specifies the configuration of the backend servers that bind the VPC channel.  
     * The object structure is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.VpcChannelMember>[]>;
    /**
     * Specifies the name of the VPC channel.  
     * The valid length is limited from `3` to `64`, only chinese and english letters, digits, hyphens (-), underscores (_)
     * and dots (.) are allowed.
     * The name must start with a chinese or english letter, and the Chinese characters must be in **UTF-8** or **Unicode**
     * format.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the destination path for health checks.  
     * Required if the `protocol` is **HTTP** or **HTTPS**.
     */
    path?: pulumi.Input<string>;
    /**
     * Specifies the host port of the VPC channel.  
     * The valid value ranges from `1` to `65,535`.
     */
    port?: pulumi.Input<number>;
    /**
     * Specifies the protocol for performing health checks on backend servers in the VPC
     * channel.
     * The valid values are **TCP**, **HTTP** and **HTTPS**, defaults to **TCP**.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Specifies the region where the VPC channel is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The current status of the VPC channel, supports **Normal** and **Abnormal**.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the timeout for determining whether a health check fails, in second.  
     * The value must be less than the value of the time `interval`.
     * The valid value ranges from `2` to `30`, defaults to `5`.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the unhealthy threshold, which refers to the number of consecutive
     * failed checks required for a backend server to be considered unhealthy.
     * The valid value ranges from `2` to `10`, defaults to `5`.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VpcChannel resource.
 */
export interface VpcChannelArgs {
    /**
     * Specifies the distribution algorithm.  
     * The valid types are **WRR**, **WLC**, **SH** and **URI hashing**, defaults to **WRR**.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * Specifies the healthy threshold, which refers to the number of consecutive
     * successful checks required for a backend server to be considered healthy.
     * The valid value ranges from `2` to `10`, defaults to `2`.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * Specifies the response codes for determining a successful HTTP response.  
     * The valid value ranges from `100` to `599` and the valid formats are as follows:
     * + The multiple values, for example, **200,201,202**.
     * + The range, for example, **200-299**.
     * + Both multiple values and ranges, for example, **201,202,210-299**.
     */
    httpCode?: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated instance to which the VPC channel
     * belongs.
     * Changing this will create a new resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the interval between consecutive checks, in second.  
     * The valid value ranges from `5` to `300`, defaults to `10`.
     */
    interval?: pulumi.Input<number>;
    /**
     * Specifies the member type of the VPC channel.  
     * The valid types are **ECS** and **EIP**, defaults to **ECS**.
     */
    memberType?: pulumi.Input<string>;
    /**
     * Specifies the configuration of the backend servers that bind the VPC channel.  
     * The object structure is documented below.
     */
    members: pulumi.Input<pulumi.Input<inputs.DedicatedApig.VpcChannelMember>[]>;
    /**
     * Specifies the name of the VPC channel.  
     * The valid length is limited from `3` to `64`, only chinese and english letters, digits, hyphens (-), underscores (_)
     * and dots (.) are allowed.
     * The name must start with a chinese or english letter, and the Chinese characters must be in **UTF-8** or **Unicode**
     * format.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the destination path for health checks.  
     * Required if the `protocol` is **HTTP** or **HTTPS**.
     */
    path?: pulumi.Input<string>;
    /**
     * Specifies the host port of the VPC channel.  
     * The valid value ranges from `1` to `65,535`.
     */
    port: pulumi.Input<number>;
    /**
     * Specifies the protocol for performing health checks on backend servers in the VPC
     * channel.
     * The valid values are **TCP**, **HTTP** and **HTTPS**, defaults to **TCP**.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Specifies the region where the VPC channel is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the timeout for determining whether a health check fails, in second.  
     * The value must be less than the value of the time `interval`.
     * The valid value ranges from `2` to `30`, defaults to `5`.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the unhealthy threshold, which refers to the number of consecutive
     * failed checks required for a backend server to be considered unhealthy.
     * The valid value ranges from `2` to `10`, defaults to `5`.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}
