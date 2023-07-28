// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an ELB loadbalancer resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Basic Loadbalancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const ipv4SubnetId = config.requireObject("ipv4SubnetId");
 * const lb1 = new huaweicloud.elb.Loadbalancer("lb1", {
 *     vipSubnetId: ipv4SubnetId,
 *     tags: {
 *         key: "value",
 *     },
 * });
 * ```
 * ### Loadbalancer With EIP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const ipv4SubnetId = config.requireObject("ipv4SubnetId");
 * const lb1 = new huaweicloud.elb.Loadbalancer("lb1", {vipSubnetId: ipv4_subnet_id});
 * const eip1 = new huaweicloud.vpc.EipAssociate("eip1", {
 *     publicIp: "1.2.3.4",
 *     portId: lb1.vipPortId,
 * });
 * ```
 *
 * ## Import
 *
 * Load balancers can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Elb/loadbalancer:Loadbalancer test 3e3632db-36c6-4b28-a92e-e72e6562daa6
 * ```
 */
export class Loadbalancer extends pulumi.CustomResource {
    /**
     * Get an existing Loadbalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadbalancerState, opts?: pulumi.CustomResourceOptions): Loadbalancer {
        return new Loadbalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Elb/loadbalancer:Loadbalancer';

    /**
     * Returns true if the given object is an instance of Loadbalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Loadbalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Loadbalancer.__pulumiType;
    }

    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Human-readable description for the loadbalancer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The enterprise project id of the loadbalancer. Changing this
     * creates a new loadbalancer.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * schema: Deprecated
     */
    public readonly flavor!: pulumi.Output<string | undefined>;
    /**
     * schema: Deprecated
     */
    public readonly loadbalancerProvider!: pulumi.Output<string>;
    /**
     * Human-readable name for the loadbalancer. Does not have to be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The EIP address that is associated to the Load Balancer instance.
     */
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    /**
     * The region in which to create the loadbalancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new loadbalancer.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * schema: Deprecated
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The key/value pairs to associate with the loadbalancer.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * @deprecated tenant_id is deprecated
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The ip address of the load balancer. Changing this creates a new
     * loadbalancer.
     */
    public readonly vipAddress!: pulumi.Output<string>;
    /**
     * The Port ID of the Load Balancer IP.
     */
    public /*out*/ readonly vipPortId!: pulumi.Output<string>;
    /**
     * The **IPv4 subnet ID** of the subnet where the load balancer works.
     * Changing this creates a new loadbalancer.
     */
    public readonly vipSubnetId!: pulumi.Output<string>;

    /**
     * Create a Loadbalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadbalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadbalancerArgs | LoadbalancerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadbalancerState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["flavor"] = state ? state.flavor : undefined;
            resourceInputs["loadbalancerProvider"] = state ? state.loadbalancerProvider : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["vipAddress"] = state ? state.vipAddress : undefined;
            resourceInputs["vipPortId"] = state ? state.vipPortId : undefined;
            resourceInputs["vipSubnetId"] = state ? state.vipSubnetId : undefined;
        } else {
            const args = argsOrState as LoadbalancerArgs | undefined;
            if ((!args || args.vipSubnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vipSubnetId'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["flavor"] = args ? args.flavor : undefined;
            resourceInputs["loadbalancerProvider"] = args ? args.loadbalancerProvider : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["vipAddress"] = args ? args.vipAddress : undefined;
            resourceInputs["vipSubnetId"] = args ? args.vipSubnetId : undefined;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["vipPortId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Loadbalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Loadbalancer resources.
 */
export interface LoadbalancerState {
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the loadbalancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the loadbalancer. Changing this
     * creates a new loadbalancer.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    flavor?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    loadbalancerProvider?: pulumi.Input<string>;
    /**
     * Human-readable name for the loadbalancer. Does not have to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The EIP address that is associated to the Load Balancer instance.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * The region in which to create the loadbalancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new loadbalancer.
     */
    region?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The key/value pairs to associate with the loadbalancer.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The ip address of the load balancer. Changing this creates a new
     * loadbalancer.
     */
    vipAddress?: pulumi.Input<string>;
    /**
     * The Port ID of the Load Balancer IP.
     */
    vipPortId?: pulumi.Input<string>;
    /**
     * The **IPv4 subnet ID** of the subnet where the load balancer works.
     * Changing this creates a new loadbalancer.
     */
    vipSubnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Loadbalancer resource.
 */
export interface LoadbalancerArgs {
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the loadbalancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the loadbalancer. Changing this
     * creates a new loadbalancer.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    flavor?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    loadbalancerProvider?: pulumi.Input<string>;
    /**
     * Human-readable name for the loadbalancer. Does not have to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the loadbalancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new loadbalancer.
     */
    region?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The key/value pairs to associate with the loadbalancer.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The ip address of the load balancer. Changing this creates a new
     * loadbalancer.
     */
    vipAddress?: pulumi.Input<string>;
    /**
     * The **IPv4 subnet ID** of the subnet where the load balancer works.
     * Changing this creates a new loadbalancer.
     */
    vipSubnetId: pulumi.Input<string>;
}
