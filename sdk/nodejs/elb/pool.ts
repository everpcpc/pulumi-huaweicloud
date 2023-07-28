// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an ELB pool resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const pool1 = new huaweicloud.Elb.Pool("pool_1", {
 *     lbMethod: "ROUND_ROBIN",
 *     listenerId: "d9415786-5f1a-428b-b35f-2f1523e146d2",
 *     persistences: [{
 *         cookieName: "testCookie",
 *         type: "HTTP_COOKIE",
 *     }],
 *     protocol: "HTTP",
 * });
 * ```
 *
 * ## Import
 *
 * ELB pool can be imported using the pool ID, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Elb/pool:Pool pool_1 5c20fdad-7288-11eb-b817-0255ac10158b
 * ```
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Elb/pool:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Human-readable description for the pool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The load balancing algorithm to distribute traffic to the pool's members. Must be one
     * of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
     */
    public readonly lbMethod!: pulumi.Output<string>;
    /**
     * The Listener on which the members of the pool will be associated with.
     * Changing this creates a new pool. Note:  At least one of LoadbalancerID or ListenerID must be provided.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * The load balancer on which to provision this pool. Changing this
     * creates a new pool. Note:  At least one of LoadbalancerID or ListenerID must be provided.
     */
    public readonly loadbalancerId!: pulumi.Output<string>;
    /**
     * Human-readable name for the pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Omit this field to prevent session persistence. Indicates whether
     * connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
     */
    public readonly persistences!: pulumi.Output<outputs.Elb.PoolPersistence[] | undefined>;
    /**
     * The protocol - can either be TCP, UDP or HTTP.
     * + When the protocol used by the listener is UDP, the protocol of the backend pool must be UDP.
     * + When the protocol used by the listener is TCP, the protocol of the backend pool must be TCP.
     * + When the protocol used by the listener is HTTP or TERMINATED_HTTPS, the protocol of the backend pool must be HTTP.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The region in which to create the ELB pool resource. If omitted, the the
     * provider-level region will be used. Changing this creates a new pool.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PoolArgs | PoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PoolState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lbMethod"] = state ? state.lbMethod : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["loadbalancerId"] = state ? state.loadbalancerId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["persistences"] = state ? state.persistences : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            if ((!args || args.lbMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbMethod'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["lbMethod"] = args ? args.lbMethod : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["loadbalancerId"] = args ? args.loadbalancerId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["persistences"] = args ? args.persistences : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the pool.
     */
    description?: pulumi.Input<string>;
    /**
     * The load balancing algorithm to distribute traffic to the pool's members. Must be one
     * of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
     */
    lbMethod?: pulumi.Input<string>;
    /**
     * The Listener on which the members of the pool will be associated with.
     * Changing this creates a new pool. Note:  At least one of LoadbalancerID or ListenerID must be provided.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this pool. Changing this
     * creates a new pool. Note:  At least one of LoadbalancerID or ListenerID must be provided.
     */
    loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the pool.
     */
    name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence. Indicates whether
     * connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
     */
    persistences?: pulumi.Input<pulumi.Input<inputs.Elb.PoolPersistence>[]>;
    /**
     * The protocol - can either be TCP, UDP or HTTP.
     * + When the protocol used by the listener is UDP, the protocol of the backend pool must be UDP.
     * + When the protocol used by the listener is TCP, the protocol of the backend pool must be TCP.
     * + When the protocol used by the listener is HTTP or TERMINATED_HTTPS, the protocol of the backend pool must be HTTP.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to create the ELB pool resource. If omitted, the the
     * provider-level region will be used. Changing this creates a new pool.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the pool.
     */
    description?: pulumi.Input<string>;
    /**
     * The load balancing algorithm to distribute traffic to the pool's members. Must be one
     * of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
     */
    lbMethod: pulumi.Input<string>;
    /**
     * The Listener on which the members of the pool will be associated with.
     * Changing this creates a new pool. Note:  At least one of LoadbalancerID or ListenerID must be provided.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this pool. Changing this
     * creates a new pool. Note:  At least one of LoadbalancerID or ListenerID must be provided.
     */
    loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the pool.
     */
    name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence. Indicates whether
     * connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
     */
    persistences?: pulumi.Input<pulumi.Input<inputs.Elb.PoolPersistence>[]>;
    /**
     * The protocol - can either be TCP, UDP or HTTP.
     * + When the protocol used by the listener is UDP, the protocol of the backend pool must be UDP.
     * + When the protocol used by the listener is TCP, the protocol of the backend pool must be TCP.
     * + When the protocol used by the listener is HTTP or TERMINATED_HTTPS, the protocol of the backend pool must be HTTP.
     */
    protocol: pulumi.Input<string>;
    /**
     * The region in which to create the ELB pool resource. If omitted, the the
     * provider-level region will be used. Changing this creates a new pool.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
}
