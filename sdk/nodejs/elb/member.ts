// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an ELB member resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const member1 = new huaweicloud.elb.Member("member1", {
 *     address: "192.168.199.23",
 *     protocolPort: 8080,
 *     poolId: _var.pool_id,
 *     subnetId: _var.subnet_id,
 * });
 * ```
 *
 * ## Import
 *
 * ELB member can be imported using the pool ID and member ID separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Elb/member:Member member_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
 * ```
 */
export class Member extends pulumi.CustomResource {
    /**
     * Get an existing Member resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberState, opts?: pulumi.CustomResourceOptions): Member {
        return new Member(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Elb/member:Member';

    /**
     * Returns true if the given object is an instance of Member.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Member {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Member.__pulumiType;
    }

    /**
     * The IP address of the member to receive traffic from the load balancer.
     * Changing this creates a new member.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Human-readable name for the member.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the pool that this member will be assigned to.
     */
    public readonly poolId!: pulumi.Output<string>;
    /**
     * The port on which to listen for client traffic. Changing this creates a
     * new member.
     */
    public readonly protocolPort!: pulumi.Output<number>;
    /**
     * The region in which to create the ELB member resource. If omitted, the the
     * provider-level region will be used. Changing this creates a new member.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The subnet in which to access the member
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * A positive integer value that indicates the relative portion of traffic that this member
     * should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a
     * member with a weight of 2.
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a Member resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MemberArgs | MemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MemberState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["poolId"] = state ? state.poolId : undefined;
            resourceInputs["protocolPort"] = state ? state.protocolPort : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as MemberArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.poolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolId'");
            }
            if ((!args || args.protocolPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocolPort'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["poolId"] = args ? args.poolId : undefined;
            resourceInputs["protocolPort"] = args ? args.protocolPort : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Member.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Member resources.
 */
export interface MemberState {
    /**
     * The IP address of the member to receive traffic from the load balancer.
     * Changing this creates a new member.
     */
    address?: pulumi.Input<string>;
    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable name for the member.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the pool that this member will be assigned to.
     */
    poolId?: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic. Changing this creates a
     * new member.
     */
    protocolPort?: pulumi.Input<number>;
    /**
     * The region in which to create the ELB member resource. If omitted, the the
     * provider-level region will be used. Changing this creates a new member.
     */
    region?: pulumi.Input<string>;
    /**
     * The subnet in which to access the member
     */
    subnetId?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * A positive integer value that indicates the relative portion of traffic that this member
     * should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a
     * member with a weight of 2.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Member resource.
 */
export interface MemberArgs {
    /**
     * The IP address of the member to receive traffic from the load balancer.
     * Changing this creates a new member.
     */
    address: pulumi.Input<string>;
    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable name for the member.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the pool that this member will be assigned to.
     */
    poolId: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic. Changing this creates a
     * new member.
     */
    protocolPort: pulumi.Input<number>;
    /**
     * The region in which to create the ELB member resource. If omitted, the the
     * provider-level region will be used. Changing this creates a new member.
     */
    region?: pulumi.Input<string>;
    /**
     * The subnet in which to access the member
     */
    subnetId: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * A positive integer value that indicates the relative portion of traffic that this member
     * should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a
     * member with a weight of 2.
     */
    weight?: pulumi.Input<number>;
}
