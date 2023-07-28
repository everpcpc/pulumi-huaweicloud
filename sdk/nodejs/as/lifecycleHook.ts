// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AS Lifecycle Hook resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const hookName = config.requireObject("hookName");
 * const asGroupId = config.requireObject("asGroupId");
 * const smnTopicUrn = config.requireObject("smnTopicUrn");
 * const test = new huaweicloud.as.LifecycleHook("test", {
 *     scalingGroupId: asGroupId,
 *     type: "ADD",
 *     defaultResult: "ABANDON",
 *     notificationTopicUrn: smnTopicUrn,
 *     notificationMessage: "This is a test message",
 * });
 * ```
 *
 * ## Import
 *
 * Lifecycle hooks can be imported using the AS group ID and hook ID separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:As/lifecycleHook:LifecycleHook test &ltAS group ID&gt/&ltLifecycle hook ID&gt
 * ```
 */
export class LifecycleHook extends pulumi.CustomResource {
    /**
     * Get an existing LifecycleHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LifecycleHookState, opts?: pulumi.CustomResourceOptions): LifecycleHook {
        return new LifecycleHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:As/lifecycleHook:LifecycleHook';

    /**
     * Returns true if the given object is an instance of LifecycleHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LifecycleHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LifecycleHook.__pulumiType;
    }

    /**
     * The server time in UTC format when the lifecycle hook is created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Specifies the default lifecycle hook callback operation. This operation is
     * performed when the timeout duration expires. The valid values are *ABANDON* and *CONTINUE*, default to *ABANDON*.
     */
    public readonly defaultResult!: pulumi.Output<string | undefined>;
    /**
     * Specifies the lifecycle hook name. This parameter can contain a maximum of
     * 32 characters, which may consist of letters, digits, underscores (_) and hyphens (-).
     * Changing this creates a new AS lifecycle hook.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies a customized notification. This parameter can contains a maximum
     * of 256 characters, which cannot contain the following characters: <>&'().
     */
    public readonly notificationMessage!: pulumi.Output<string | undefined>;
    /**
     * The topic name in SMN.
     */
    public /*out*/ readonly notificationTopicName!: pulumi.Output<string>;
    /**
     * Specifies a unique topic in SMN.
     */
    public readonly notificationTopicUrn!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the AS lifecycle hook.
     * If omitted, the provider-level region will be used. Changing this creates a new AS lifecycle hook.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the ID of the AS group in UUID format.
     * Changing this creates a new AS lifecycle hook.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * Specifies the lifecycle hook timeout duration, which ranges from 300 to 86400 in the unit
     * of second, default to 3600.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Specifies the lifecycle hook type. The valid values are following strings:
     * + `ADD`: The hook suspends the instance when the instance is started.
     * + `REMOVE`: The hook suspends the instance when the instance is terminated.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a LifecycleHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LifecycleHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LifecycleHookArgs | LifecycleHookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LifecycleHookState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["defaultResult"] = state ? state.defaultResult : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationMessage"] = state ? state.notificationMessage : undefined;
            resourceInputs["notificationTopicName"] = state ? state.notificationTopicName : undefined;
            resourceInputs["notificationTopicUrn"] = state ? state.notificationTopicUrn : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as LifecycleHookArgs | undefined;
            if ((!args || args.notificationTopicUrn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationTopicUrn'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["defaultResult"] = args ? args.defaultResult : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationMessage"] = args ? args.notificationMessage : undefined;
            resourceInputs["notificationTopicUrn"] = args ? args.notificationTopicUrn : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["notificationTopicName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LifecycleHook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LifecycleHook resources.
 */
export interface LifecycleHookState {
    /**
     * The server time in UTC format when the lifecycle hook is created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Specifies the default lifecycle hook callback operation. This operation is
     * performed when the timeout duration expires. The valid values are *ABANDON* and *CONTINUE*, default to *ABANDON*.
     */
    defaultResult?: pulumi.Input<string>;
    /**
     * Specifies the lifecycle hook name. This parameter can contain a maximum of
     * 32 characters, which may consist of letters, digits, underscores (_) and hyphens (-).
     * Changing this creates a new AS lifecycle hook.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies a customized notification. This parameter can contains a maximum
     * of 256 characters, which cannot contain the following characters: <>&'().
     */
    notificationMessage?: pulumi.Input<string>;
    /**
     * The topic name in SMN.
     */
    notificationTopicName?: pulumi.Input<string>;
    /**
     * Specifies a unique topic in SMN.
     */
    notificationTopicUrn?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the AS lifecycle hook.
     * If omitted, the provider-level region will be used. Changing this creates a new AS lifecycle hook.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the ID of the AS group in UUID format.
     * Changing this creates a new AS lifecycle hook.
     */
    scalingGroupId?: pulumi.Input<string>;
    /**
     * Specifies the lifecycle hook timeout duration, which ranges from 300 to 86400 in the unit
     * of second, default to 3600.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the lifecycle hook type. The valid values are following strings:
     * + `ADD`: The hook suspends the instance when the instance is started.
     * + `REMOVE`: The hook suspends the instance when the instance is terminated.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LifecycleHook resource.
 */
export interface LifecycleHookArgs {
    /**
     * Specifies the default lifecycle hook callback operation. This operation is
     * performed when the timeout duration expires. The valid values are *ABANDON* and *CONTINUE*, default to *ABANDON*.
     */
    defaultResult?: pulumi.Input<string>;
    /**
     * Specifies the lifecycle hook name. This parameter can contain a maximum of
     * 32 characters, which may consist of letters, digits, underscores (_) and hyphens (-).
     * Changing this creates a new AS lifecycle hook.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies a customized notification. This parameter can contains a maximum
     * of 256 characters, which cannot contain the following characters: <>&'().
     */
    notificationMessage?: pulumi.Input<string>;
    /**
     * Specifies a unique topic in SMN.
     */
    notificationTopicUrn: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the AS lifecycle hook.
     * If omitted, the provider-level region will be used. Changing this creates a new AS lifecycle hook.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the ID of the AS group in UUID format.
     * Changing this creates a new AS lifecycle hook.
     */
    scalingGroupId: pulumi.Input<string>;
    /**
     * Specifies the lifecycle hook timeout duration, which ranges from 300 to 86400 in the unit
     * of second, default to 3600.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the lifecycle hook type. The valid values are following strings:
     * + `ADD`: The hook suspends the instance when the instance is started.
     * + `REMOVE`: The hook suspends the instance when the instance is terminated.
     */
    type: pulumi.Input<string>;
}
