// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an AS bandwidth scaling policy resource within HuaweiCloud.
 *
 * > AS cannot scale yearly/monthly bandwidths.
 *
 * ## Example Usage
 * ### AS Recurrence Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const bandwidthId = config.requireObject("bandwidthId");
 * const bwPolicy = new huaweicloud.as.BandwidthPolicy("bwPolicy", {
 *     scalingPolicyName: "bw_policy",
 *     scalingPolicyType: "RECURRENCE",
 *     bandwidthId: bandwidthId,
 *     coolDownTime: 600,
 *     scalingPolicyAction: {
 *         operation: "ADD",
 *         size: 1,
 *     },
 *     scheduledPolicy: {
 *         launchTime: "07:00",
 *         recurrenceType: "Weekly",
 *         recurrenceValue: "1,3,5",
 *         startTime: "2022-09-30T12:00Z",
 *         endTime: "2022-12-30T12:00Z",
 *     },
 * });
 * ```
 * ### AS Scheduled Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const bandwidthId = config.requireObject("bandwidthId");
 * const bwPolicy = new huaweicloud.as.BandwidthPolicy("bwPolicy", {
 *     scalingPolicyName: "bw_policy",
 *     scalingPolicyType: "SCHEDULED",
 *     bandwidthId: bandwidthId,
 *     coolDownTime: 600,
 *     scalingPolicyAction: {
 *         operation: "ADD",
 *         size: 1,
 *     },
 *     scheduledPolicy: {
 *         launchTime: "2022-09-30T12:00Z",
 *     },
 * });
 * ```
 * ### AS Alarm Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const bandwidthId = config.requireObject("bandwidthId");
 * const alarmId = config.requireObject("alarmId");
 * const test = new huaweicloud.as.BandwidthPolicy("test", {
 *     scalingPolicyName: "bw_policy",
 *     scalingPolicyType: "ALARM",
 *     bandwidthId: bandwidthId,
 *     alarmId: alarmId,
 *     scalingPolicyAction: {
 *         operation: "ADD",
 *         size: 1,
 *         limits: 300,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * The bandwidth scaling policies can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:As/bandwidthPolicy:BandwidthPolicy test 0ce123456a00f2591fabc00385ff1234
 * ```
 */
export class BandwidthPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthPolicyState, opts?: pulumi.CustomResourceOptions): BandwidthPolicy {
        return new BandwidthPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:As/bandwidthPolicy:BandwidthPolicy';

    /**
     * Returns true if the given object is an instance of BandwidthPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthPolicy.__pulumiType;
    }

    /**
     * Specifies the alarm rule ID.
     * This parameter is mandatory when `scalingPolicyType` is set to ALARM.
     */
    public readonly alarmId!: pulumi.Output<string>;
    /**
     * Specifies the scaling bandwidth ID.
     */
    public readonly bandwidthId!: pulumi.Output<string>;
    /**
     * Specifies the cooldown period (in seconds).
     * The value ranges from 0 to 86400 and is 300 by default.
     */
    public readonly coolDownTime!: pulumi.Output<number>;
    /**
     * Specifies the description of the AS policy.
     * The value can contain 0 to 256 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the scaling action of the AS policy.
     * The object structure is documented below.
     */
    public readonly scalingPolicyAction!: pulumi.Output<outputs.As.BandwidthPolicyScalingPolicyAction>;
    /**
     * Specifies the AS policy name.
     * The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
     */
    public readonly scalingPolicyName!: pulumi.Output<string>;
    /**
     * Specifies the AS policy type. The options are as follows:
     * - **ALARM** (corresponding to `alarmId`): indicates that the scaling action is triggered by an alarm.
     * - **SCHEDULED** (corresponding to `scheduledPolicy`): indicates that the scaling action is triggered as scheduled.
     * - **RECURRENCE** (corresponding to `scheduledPolicy`): indicates that the scaling action is triggered periodically.
     */
    public readonly scalingPolicyType!: pulumi.Output<string>;
    /**
     * The scaling resource type. The value is fixed to **BANDWIDTH**.
     */
    public /*out*/ readonly scalingResourceType!: pulumi.Output<string>;
    /**
     * Specifies the periodic or scheduled AS policy.
     * This parameter is mandatory when `scalingPolicyType` is set to SCHEDULED or RECURRENCE.
     * The object structure is documented below.
     */
    public readonly scheduledPolicy!: pulumi.Output<outputs.As.BandwidthPolicyScheduledPolicy>;
    /**
     * The AS policy status. The value can be **INSERVICE**, **PAUSED** and **EXECUTING**.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a BandwidthPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPolicyArgs | BandwidthPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BandwidthPolicyState | undefined;
            resourceInputs["alarmId"] = state ? state.alarmId : undefined;
            resourceInputs["bandwidthId"] = state ? state.bandwidthId : undefined;
            resourceInputs["coolDownTime"] = state ? state.coolDownTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["scalingPolicyAction"] = state ? state.scalingPolicyAction : undefined;
            resourceInputs["scalingPolicyName"] = state ? state.scalingPolicyName : undefined;
            resourceInputs["scalingPolicyType"] = state ? state.scalingPolicyType : undefined;
            resourceInputs["scalingResourceType"] = state ? state.scalingResourceType : undefined;
            resourceInputs["scheduledPolicy"] = state ? state.scheduledPolicy : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as BandwidthPolicyArgs | undefined;
            if ((!args || args.bandwidthId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthId'");
            }
            if ((!args || args.scalingPolicyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingPolicyName'");
            }
            if ((!args || args.scalingPolicyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingPolicyType'");
            }
            resourceInputs["alarmId"] = args ? args.alarmId : undefined;
            resourceInputs["bandwidthId"] = args ? args.bandwidthId : undefined;
            resourceInputs["coolDownTime"] = args ? args.coolDownTime : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["scalingPolicyAction"] = args ? args.scalingPolicyAction : undefined;
            resourceInputs["scalingPolicyName"] = args ? args.scalingPolicyName : undefined;
            resourceInputs["scalingPolicyType"] = args ? args.scalingPolicyType : undefined;
            resourceInputs["scheduledPolicy"] = args ? args.scheduledPolicy : undefined;
            resourceInputs["scalingResourceType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BandwidthPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthPolicy resources.
 */
export interface BandwidthPolicyState {
    /**
     * Specifies the alarm rule ID.
     * This parameter is mandatory when `scalingPolicyType` is set to ALARM.
     */
    alarmId?: pulumi.Input<string>;
    /**
     * Specifies the scaling bandwidth ID.
     */
    bandwidthId?: pulumi.Input<string>;
    /**
     * Specifies the cooldown period (in seconds).
     * The value ranges from 0 to 86400 and is 300 by default.
     */
    coolDownTime?: pulumi.Input<number>;
    /**
     * Specifies the description of the AS policy.
     * The value can contain 0 to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the scaling action of the AS policy.
     * The object structure is documented below.
     */
    scalingPolicyAction?: pulumi.Input<inputs.As.BandwidthPolicyScalingPolicyAction>;
    /**
     * Specifies the AS policy name.
     * The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
     */
    scalingPolicyName?: pulumi.Input<string>;
    /**
     * Specifies the AS policy type. The options are as follows:
     * - **ALARM** (corresponding to `alarmId`): indicates that the scaling action is triggered by an alarm.
     * - **SCHEDULED** (corresponding to `scheduledPolicy`): indicates that the scaling action is triggered as scheduled.
     * - **RECURRENCE** (corresponding to `scheduledPolicy`): indicates that the scaling action is triggered periodically.
     */
    scalingPolicyType?: pulumi.Input<string>;
    /**
     * The scaling resource type. The value is fixed to **BANDWIDTH**.
     */
    scalingResourceType?: pulumi.Input<string>;
    /**
     * Specifies the periodic or scheduled AS policy.
     * This parameter is mandatory when `scalingPolicyType` is set to SCHEDULED or RECURRENCE.
     * The object structure is documented below.
     */
    scheduledPolicy?: pulumi.Input<inputs.As.BandwidthPolicyScheduledPolicy>;
    /**
     * The AS policy status. The value can be **INSERVICE**, **PAUSED** and **EXECUTING**.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BandwidthPolicy resource.
 */
export interface BandwidthPolicyArgs {
    /**
     * Specifies the alarm rule ID.
     * This parameter is mandatory when `scalingPolicyType` is set to ALARM.
     */
    alarmId?: pulumi.Input<string>;
    /**
     * Specifies the scaling bandwidth ID.
     */
    bandwidthId: pulumi.Input<string>;
    /**
     * Specifies the cooldown period (in seconds).
     * The value ranges from 0 to 86400 and is 300 by default.
     */
    coolDownTime?: pulumi.Input<number>;
    /**
     * Specifies the description of the AS policy.
     * The value can contain 0 to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the scaling action of the AS policy.
     * The object structure is documented below.
     */
    scalingPolicyAction?: pulumi.Input<inputs.As.BandwidthPolicyScalingPolicyAction>;
    /**
     * Specifies the AS policy name.
     * The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
     */
    scalingPolicyName: pulumi.Input<string>;
    /**
     * Specifies the AS policy type. The options are as follows:
     * - **ALARM** (corresponding to `alarmId`): indicates that the scaling action is triggered by an alarm.
     * - **SCHEDULED** (corresponding to `scheduledPolicy`): indicates that the scaling action is triggered as scheduled.
     * - **RECURRENCE** (corresponding to `scheduledPolicy`): indicates that the scaling action is triggered periodically.
     */
    scalingPolicyType: pulumi.Input<string>;
    /**
     * Specifies the periodic or scheduled AS policy.
     * This parameter is mandatory when `scalingPolicyType` is set to SCHEDULED or RECURRENCE.
     * The object structure is documented below.
     */
    scheduledPolicy?: pulumi.Input<inputs.As.BandwidthPolicyScheduledPolicy>;
}
