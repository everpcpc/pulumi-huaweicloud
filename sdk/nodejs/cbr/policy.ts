// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a CBR Policy resource within Huaweicloud.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Policies can be imported by their `id`. For example,
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cbr/policy:Policy test 4d2c2939-774f-42ef-ab15-e5b126b11ace
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to the attribute missing from the API response. The missing attribute is`enable_acceleration`. It is generally recommended running `terraform plan` after importing a policy. You can then decide if changes should be applied to the policy, or the resource definition should be updated to align with the policy. Also you can ignore changes as below. resource "huaweicloud_cbr_policy" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  enable_acceleration,
 *
 *  ]
 *
 *  } }
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cbr/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Specifies the scheduling rule for the policy backup execution.
     * The object structure is documented below.
     */
    public readonly backupCycle!: pulumi.Output<outputs.Cbr.PolicyBackupCycle>;
    /**
     * Specifies the maximum number of retained backups. The value ranges from `2` to
     * `99,999`. This parameter and `timePeriod` are alternative.
     */
    public readonly backupQuantity!: pulumi.Output<number | undefined>;
    /**
     * Specifies the ID of the replication destination project, which is
     * mandatory for cross-region replication. Required if `protectionType` is **replication**.
     */
    public readonly destinationProjectId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the replication destination region, which is mandatory
     * for cross-region replication. Required if `protectionType` is **replication**.
     */
    public readonly destinationRegion!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to enable the acceleration function to shorten
     * the replication time for cross-region.
     * Changing this will create a new policy.
     */
    public readonly enableAcceleration!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to enable the policy. Default to **true**.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the long-term retention rules, which is an advanced options of
     * the `backupQuantity`. The object structure is documented below.
     */
    public readonly longTermRetention!: pulumi.Output<outputs.Cbr.PolicyLongTermRetention | undefined>;
    /**
     * Specifies the policy name.  
     * This parameter can contain a maximum of 64
     * characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region where the policy is located. If omitted, the
     * provider-level region will be used. Changing this will create a new policy.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the duration (in days) for retained backups. The value ranges from `2` to
     * `99,999`.
     */
    public readonly timePeriod!: pulumi.Output<number | undefined>;
    /**
     * Specifies the UTC time zone, e.g. `UTC+08:00`.
     * Only available if `longTermRetention` is set.
     */
    public readonly timeZone!: pulumi.Output<string>;
    /**
     * Specifies the protection type of the policy.
     * Valid values are **backup** and **replication**.
     * Changing this will create a new policy.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            resourceInputs["backupCycle"] = state ? state.backupCycle : undefined;
            resourceInputs["backupQuantity"] = state ? state.backupQuantity : undefined;
            resourceInputs["destinationProjectId"] = state ? state.destinationProjectId : undefined;
            resourceInputs["destinationRegion"] = state ? state.destinationRegion : undefined;
            resourceInputs["enableAcceleration"] = state ? state.enableAcceleration : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["longTermRetention"] = state ? state.longTermRetention : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["timePeriod"] = state ? state.timePeriod : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.backupCycle === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupCycle'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["backupCycle"] = args ? args.backupCycle : undefined;
            resourceInputs["backupQuantity"] = args ? args.backupQuantity : undefined;
            resourceInputs["destinationProjectId"] = args ? args.destinationProjectId : undefined;
            resourceInputs["destinationRegion"] = args ? args.destinationRegion : undefined;
            resourceInputs["enableAcceleration"] = args ? args.enableAcceleration : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["longTermRetention"] = args ? args.longTermRetention : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["timePeriod"] = args ? args.timePeriod : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Specifies the scheduling rule for the policy backup execution.
     * The object structure is documented below.
     */
    backupCycle?: pulumi.Input<inputs.Cbr.PolicyBackupCycle>;
    /**
     * Specifies the maximum number of retained backups. The value ranges from `2` to
     * `99,999`. This parameter and `timePeriod` are alternative.
     */
    backupQuantity?: pulumi.Input<number>;
    /**
     * Specifies the ID of the replication destination project, which is
     * mandatory for cross-region replication. Required if `protectionType` is **replication**.
     */
    destinationProjectId?: pulumi.Input<string>;
    /**
     * Specifies the name of the replication destination region, which is mandatory
     * for cross-region replication. Required if `protectionType` is **replication**.
     */
    destinationRegion?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the acceleration function to shorten
     * the replication time for cross-region.
     * Changing this will create a new policy.
     */
    enableAcceleration?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the policy. Default to **true**.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the long-term retention rules, which is an advanced options of
     * the `backupQuantity`. The object structure is documented below.
     */
    longTermRetention?: pulumi.Input<inputs.Cbr.PolicyLongTermRetention>;
    /**
     * Specifies the policy name.  
     * This parameter can contain a maximum of 64
     * characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region where the policy is located. If omitted, the
     * provider-level region will be used. Changing this will create a new policy.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the duration (in days) for retained backups. The value ranges from `2` to
     * `99,999`.
     */
    timePeriod?: pulumi.Input<number>;
    /**
     * Specifies the UTC time zone, e.g. `UTC+08:00`.
     * Only available if `longTermRetention` is set.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * Specifies the protection type of the policy.
     * Valid values are **backup** and **replication**.
     * Changing this will create a new policy.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Specifies the scheduling rule for the policy backup execution.
     * The object structure is documented below.
     */
    backupCycle: pulumi.Input<inputs.Cbr.PolicyBackupCycle>;
    /**
     * Specifies the maximum number of retained backups. The value ranges from `2` to
     * `99,999`. This parameter and `timePeriod` are alternative.
     */
    backupQuantity?: pulumi.Input<number>;
    /**
     * Specifies the ID of the replication destination project, which is
     * mandatory for cross-region replication. Required if `protectionType` is **replication**.
     */
    destinationProjectId?: pulumi.Input<string>;
    /**
     * Specifies the name of the replication destination region, which is mandatory
     * for cross-region replication. Required if `protectionType` is **replication**.
     */
    destinationRegion?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the acceleration function to shorten
     * the replication time for cross-region.
     * Changing this will create a new policy.
     */
    enableAcceleration?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the policy. Default to **true**.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the long-term retention rules, which is an advanced options of
     * the `backupQuantity`. The object structure is documented below.
     */
    longTermRetention?: pulumi.Input<inputs.Cbr.PolicyLongTermRetention>;
    /**
     * Specifies the policy name.  
     * This parameter can contain a maximum of 64
     * characters, which may consist of chinese characters, letters, digits, underscores(_) and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region where the policy is located. If omitted, the
     * provider-level region will be used. Changing this will create a new policy.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the duration (in days) for retained backups. The value ranges from `2` to
     * `99,999`.
     */
    timePeriod?: pulumi.Input<number>;
    /**
     * Specifies the UTC time zone, e.g. `UTC+08:00`.
     * Only available if `longTermRetention` is set.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * Specifies the protection type of the policy.
     * Valid values are **backup** and **replication**.
     * Changing this will create a new policy.
     */
    type: pulumi.Input<string>;
}
