// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages CTS key event notification resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Complete Notification
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const topicUrn = config.requireObject("topicUrn");
 * const notify = new huaweicloud.cts.Notification("notify", {
 *     operationType: "complete",
 *     smnTopic: topicUrn,
 * });
 * ```
 * ### Customized Notification
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const topicUrn = config.requireObject("topicUrn");
 * const notify = new huaweicloud.cts.Notification("notify", {
 *     operationType: "customized",
 *     smnTopic: topicUrn,
 *     operations: [{
 *         service: "ECS",
 *         resource: "ecs",
 *         traceNames: [
 *             "createServer",
 *             "deleteServer",
 *         ],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * CTS notifications can be imported using `name`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cts/notification:Notification tracker your_notification
 * ```
 */
export class Notification extends pulumi.CustomResource {
    /**
     * Get an existing Notification resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationState, opts?: pulumi.CustomResourceOptions): Notification {
        return new Notification(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cts/notification:Notification';

    /**
     * Returns true if the given object is an instance of Notification.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Notification {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Notification.__pulumiType;
    }

    /**
     * Specifies whether notification is enabled, defaults to true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the notification name. The value contains a maximum of 64 characters,
     * and only letters, digits, underscores(_), and Chinese characters are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The notification ID in UUID format.
     */
    public /*out*/ readonly notificationId!: pulumi.Output<string>;
    /**
     * Specifies the operation type, possible options include **complete** and
     * **customized**.
     */
    public readonly operationType!: pulumi.Output<string>;
    /**
     * Specifies an array of users. Notifications will be sent when specified users
     * perform specified operations. All users are selected by default.
     * The object structure is documented below.
     */
    public readonly operationUsers!: pulumi.Output<outputs.Cts.NotificationOperationUser[] | undefined>;
    /**
     * Specifies an array of operations that will trigger notifications.
     * For details, see [Supported Services and Operations](https://support.huaweicloud.com/intl/en-us/usermanual-cts/cts_03_0022.html).
     * The object structure is documented below.
     */
    public readonly operations!: pulumi.Output<outputs.Cts.NotificationOperation[] | undefined>;
    /**
     * Specifies the region in which to manage the CTS notification resource.
     * If omitted, the provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the URN of a topic.
     */
    public readonly smnTopic!: pulumi.Output<string | undefined>;
    /**
     * The notification status, the value can be **enabled** or **disabled**.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Notification resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationArgs | NotificationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationId"] = state ? state.notificationId : undefined;
            resourceInputs["operationType"] = state ? state.operationType : undefined;
            resourceInputs["operationUsers"] = state ? state.operationUsers : undefined;
            resourceInputs["operations"] = state ? state.operations : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["smnTopic"] = state ? state.smnTopic : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as NotificationArgs | undefined;
            if ((!args || args.operationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operationType'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operationType"] = args ? args.operationType : undefined;
            resourceInputs["operationUsers"] = args ? args.operationUsers : undefined;
            resourceInputs["operations"] = args ? args.operations : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["smnTopic"] = args ? args.smnTopic : undefined;
            resourceInputs["notificationId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Notification.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Notification resources.
 */
export interface NotificationState {
    /**
     * Specifies whether notification is enabled, defaults to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the notification name. The value contains a maximum of 64 characters,
     * and only letters, digits, underscores(_), and Chinese characters are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * The notification ID in UUID format.
     */
    notificationId?: pulumi.Input<string>;
    /**
     * Specifies the operation type, possible options include **complete** and
     * **customized**.
     */
    operationType?: pulumi.Input<string>;
    /**
     * Specifies an array of users. Notifications will be sent when specified users
     * perform specified operations. All users are selected by default.
     * The object structure is documented below.
     */
    operationUsers?: pulumi.Input<pulumi.Input<inputs.Cts.NotificationOperationUser>[]>;
    /**
     * Specifies an array of operations that will trigger notifications.
     * For details, see [Supported Services and Operations](https://support.huaweicloud.com/intl/en-us/usermanual-cts/cts_03_0022.html).
     * The object structure is documented below.
     */
    operations?: pulumi.Input<pulumi.Input<inputs.Cts.NotificationOperation>[]>;
    /**
     * Specifies the region in which to manage the CTS notification resource.
     * If omitted, the provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the URN of a topic.
     */
    smnTopic?: pulumi.Input<string>;
    /**
     * The notification status, the value can be **enabled** or **disabled**.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Notification resource.
 */
export interface NotificationArgs {
    /**
     * Specifies whether notification is enabled, defaults to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the notification name. The value contains a maximum of 64 characters,
     * and only letters, digits, underscores(_), and Chinese characters are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the operation type, possible options include **complete** and
     * **customized**.
     */
    operationType: pulumi.Input<string>;
    /**
     * Specifies an array of users. Notifications will be sent when specified users
     * perform specified operations. All users are selected by default.
     * The object structure is documented below.
     */
    operationUsers?: pulumi.Input<pulumi.Input<inputs.Cts.NotificationOperationUser>[]>;
    /**
     * Specifies an array of operations that will trigger notifications.
     * For details, see [Supported Services and Operations](https://support.huaweicloud.com/intl/en-us/usermanual-cts/cts_03_0022.html).
     * The object structure is documented below.
     */
    operations?: pulumi.Input<pulumi.Input<inputs.Cts.NotificationOperation>[]>;
    /**
     * Specifies the region in which to manage the CTS notification resource.
     * If omitted, the provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the URN of a topic.
     */
    smnTopic?: pulumi.Input<string>;
}
