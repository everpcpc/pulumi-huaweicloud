// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an SMN subscription resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const topic1 = new huaweicloud.smn.Topic("topic1", {displayName: "The display name of topic_1"});
 * const subscription1 = new huaweicloud.smn.Subscription("subscription1", {
 *     topicUrn: topic1.id,
 *     endpoint: "mailtest@gmail.com",
 *     protocol: "email",
 *     remark: "O&M",
 * });
 * const subscription2 = new huaweicloud.smn.Subscription("subscription2", {
 *     topicUrn: topic1.id,
 *     endpoint: "13600000000",
 *     protocol: "sms",
 *     remark: "O&M",
 * });
 * ```
 *
 * ## Import
 *
 * SMN subscription can be imported using the `id` (subscription urn), e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Smn/subscription:Subscription subscription_1 urn:smn:cn-north-4:0970dd7a1300f5672ff2c003c60ae115:topic_1:a2aa5a1f66df494184f4e108398de1a6
 * ```
 */
export class Subscription extends pulumi.CustomResource {
    /**
     * Get an existing Subscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionState, opts?: pulumi.CustomResourceOptions): Subscription {
        return new Subscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Smn/subscription:Subscription';

    /**
     * Returns true if the given object is an instance of Subscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subscription.__pulumiType;
    }

    /**
     * Message endpoint. Changing this parameter will create a new resource.
     * + **For an HTTP subscription**, the endpoint starts with http://.
     * + **For an HTTPS subscription**, the endpoint starts with https://.
     * + **For an email subscription**, the endpoint is an mail address.
     * + **For an SMS message subscription**, the endpoint is a phone number,
     * the format is \[+\]\[country code\]\[phone number\], e.g. +86185xxxx0000.
     * + **For a functionstage subscription**, the endpoint is a function urn.
     * + **For a functiongraph subscription**, the endpoint is a workflow ID.
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * Project ID of the topic creator.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * Specifies the protocol of the message endpoint. Currently, email, sms, http,
     * https, functionstage and functiongraph are supported. Changing this parameter will create a new resource.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The region in which to create the SMN subscription resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Remark information. The remarks must be a UTF-8-coded character string
     * containing 128 bytes.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Subscription status. 0 indicates that the subscription is not confirmed. 1 indicates that the subscription
     * is confirmed. 3 indicates that the subscription is canceled.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Resource identifier of a subscription, which is unique.
     */
    public /*out*/ readonly subscriptionUrn!: pulumi.Output<string>;
    /**
     * Specifies the resource identifier of a topic, which is unique.
     * Changing this parameter will create a new resource.
     */
    public readonly topicUrn!: pulumi.Output<string>;

    /**
     * Create a Subscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionArgs | SubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionState | undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subscriptionUrn"] = state ? state.subscriptionUrn : undefined;
            resourceInputs["topicUrn"] = state ? state.topicUrn : undefined;
        } else {
            const args = argsOrState as SubscriptionArgs | undefined;
            if ((!args || args.endpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoint'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.topicUrn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicUrn'");
            }
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["topicUrn"] = args ? args.topicUrn : undefined;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subscriptionUrn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subscription resources.
 */
export interface SubscriptionState {
    /**
     * Message endpoint. Changing this parameter will create a new resource.
     * + **For an HTTP subscription**, the endpoint starts with http://.
     * + **For an HTTPS subscription**, the endpoint starts with https://.
     * + **For an email subscription**, the endpoint is an mail address.
     * + **For an SMS message subscription**, the endpoint is a phone number,
     * the format is \[+\]\[country code\]\[phone number\], e.g. +86185xxxx0000.
     * + **For a functionstage subscription**, the endpoint is a function urn.
     * + **For a functiongraph subscription**, the endpoint is a workflow ID.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Project ID of the topic creator.
     */
    owner?: pulumi.Input<string>;
    /**
     * Specifies the protocol of the message endpoint. Currently, email, sms, http,
     * https, functionstage and functiongraph are supported. Changing this parameter will create a new resource.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to create the SMN subscription resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Remark information. The remarks must be a UTF-8-coded character string
     * containing 128 bytes.
     */
    remark?: pulumi.Input<string>;
    /**
     * Subscription status. 0 indicates that the subscription is not confirmed. 1 indicates that the subscription
     * is confirmed. 3 indicates that the subscription is canceled.
     */
    status?: pulumi.Input<number>;
    /**
     * Resource identifier of a subscription, which is unique.
     */
    subscriptionUrn?: pulumi.Input<string>;
    /**
     * Specifies the resource identifier of a topic, which is unique.
     * Changing this parameter will create a new resource.
     */
    topicUrn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subscription resource.
 */
export interface SubscriptionArgs {
    /**
     * Message endpoint. Changing this parameter will create a new resource.
     * + **For an HTTP subscription**, the endpoint starts with http://.
     * + **For an HTTPS subscription**, the endpoint starts with https://.
     * + **For an email subscription**, the endpoint is an mail address.
     * + **For an SMS message subscription**, the endpoint is a phone number,
     * the format is \[+\]\[country code\]\[phone number\], e.g. +86185xxxx0000.
     * + **For a functionstage subscription**, the endpoint is a function urn.
     * + **For a functiongraph subscription**, the endpoint is a workflow ID.
     */
    endpoint: pulumi.Input<string>;
    /**
     * Specifies the protocol of the message endpoint. Currently, email, sms, http,
     * https, functionstage and functiongraph are supported. Changing this parameter will create a new resource.
     */
    protocol: pulumi.Input<string>;
    /**
     * The region in which to create the SMN subscription resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Remark information. The remarks must be a UTF-8-coded character string
     * containing 128 bytes.
     */
    remark?: pulumi.Input<string>;
    /**
     * Specifies the resource identifier of a topic, which is unique.
     * Changing this parameter will create a new resource.
     */
    topicUrn: pulumi.Input<string>;
}
