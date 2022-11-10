// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a VPC route resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Add route to the default route table
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const nexthop = config.requireObject("nexthop");
 * const vpcRoute = new huaweicloud.vpc.Route("vpcRoute", {
 *     vpcId: vpcId,
 *     destination: "192.168.0.0/16",
 *     type: "peering",
 *     nexthop: nexthop,
 * });
 * ```
 * ### Add route to a custom route table
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const nexthop = config.requireObject("nexthop");
 * const rtb = huaweicloud.Vpc.getRouteTable({
 *     vpcId: vpcId,
 *     name: "demo",
 * });
 * const vpcRoute = new huaweicloud.vpc.Route("vpcRoute", {
 *     vpcId: vpcId,
 *     routeTableId: rtb.then(rtb => rtb.id),
 *     destination: "172.16.8.0/24",
 *     type: "ecs",
 *     nexthop: nexthop,
 * });
 * ```
 *
 * ## Import
 *
 * VPC routes can be imported using the route table ID and their `destination` separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Vpc/route:Route test <route_table_id>/<destination>
 * ```
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteState, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Vpc/route:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * - Specifies the supplementary information about the route.
     * The value is a string of no more than 255 characters and cannot contain angle brackets (< or >).
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * - Specifies the destination address in the CIDR notation format,
     * for example, 192.168.200.0/24. The destination of each route must be unique and cannot overlap with any
     * subnet in the VPC. Changing this creates a new resource.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * - Specifies the next hop.
     * + If the route type is **ecs**, the value is an ECS instance ID in the VPC.
     * + If the route type is **eni**, the value is the extension NIC of an ECS in the VPC.
     * + If the route type is **vip**, the value is a virtual IP address.
     * + If the route type is **nat**, the value is a VPN gateway ID.
     * + If the route type is **peering**, the value is a VPC peering connection ID.
     * + If the route type is **vpn**, the value is a VPN gateway ID.
     * + If the route type is **dc**, the value is a Direct Connect gateway ID.
     * + If the route type is **cc**, the value is a Cloud Connection ID.
     */
    public readonly nexthop!: pulumi.Output<string>;
    /**
     * The region in which to create the VPC route. If omitted, the provider-level
     * region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * - Specifies the route table ID for which a route is to be added.
     * If the value is not set, the route will be added to the *default* route table.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * The name of route table.
     */
    public /*out*/ readonly routeTableName!: pulumi.Output<string>;
    /**
     * - Specifies the route type. Currently, the value can be:
     * **ecs**, **eni**, **vip**, **nat**, **peering**, **vpn**, **dc** and **cc**.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * - Specifies the VPC for which a route is to be added. Changing this creates a
     * new resource.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteArgs | RouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["nexthop"] = state ? state.nexthop : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
            resourceInputs["routeTableName"] = state ? state.routeTableName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as RouteArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.nexthop === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nexthop'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["nexthop"] = args ? args.nexthop : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["routeTableName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Route.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Route resources.
 */
export interface RouteState {
    /**
     * - Specifies the supplementary information about the route.
     * The value is a string of no more than 255 characters and cannot contain angle brackets (< or >).
     */
    description?: pulumi.Input<string>;
    /**
     * - Specifies the destination address in the CIDR notation format,
     * for example, 192.168.200.0/24. The destination of each route must be unique and cannot overlap with any
     * subnet in the VPC. Changing this creates a new resource.
     */
    destination?: pulumi.Input<string>;
    /**
     * - Specifies the next hop.
     * + If the route type is **ecs**, the value is an ECS instance ID in the VPC.
     * + If the route type is **eni**, the value is the extension NIC of an ECS in the VPC.
     * + If the route type is **vip**, the value is a virtual IP address.
     * + If the route type is **nat**, the value is a VPN gateway ID.
     * + If the route type is **peering**, the value is a VPC peering connection ID.
     * + If the route type is **vpn**, the value is a VPN gateway ID.
     * + If the route type is **dc**, the value is a Direct Connect gateway ID.
     * + If the route type is **cc**, the value is a Cloud Connection ID.
     */
    nexthop?: pulumi.Input<string>;
    /**
     * The region in which to create the VPC route. If omitted, the provider-level
     * region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * - Specifies the route table ID for which a route is to be added.
     * If the value is not set, the route will be added to the *default* route table.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The name of route table.
     */
    routeTableName?: pulumi.Input<string>;
    /**
     * - Specifies the route type. Currently, the value can be:
     * **ecs**, **eni**, **vip**, **nat**, **peering**, **vpn**, **dc** and **cc**.
     */
    type?: pulumi.Input<string>;
    /**
     * - Specifies the VPC for which a route is to be added. Changing this creates a
     * new resource.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * - Specifies the supplementary information about the route.
     * The value is a string of no more than 255 characters and cannot contain angle brackets (< or >).
     */
    description?: pulumi.Input<string>;
    /**
     * - Specifies the destination address in the CIDR notation format,
     * for example, 192.168.200.0/24. The destination of each route must be unique and cannot overlap with any
     * subnet in the VPC. Changing this creates a new resource.
     */
    destination: pulumi.Input<string>;
    /**
     * - Specifies the next hop.
     * + If the route type is **ecs**, the value is an ECS instance ID in the VPC.
     * + If the route type is **eni**, the value is the extension NIC of an ECS in the VPC.
     * + If the route type is **vip**, the value is a virtual IP address.
     * + If the route type is **nat**, the value is a VPN gateway ID.
     * + If the route type is **peering**, the value is a VPC peering connection ID.
     * + If the route type is **vpn**, the value is a VPN gateway ID.
     * + If the route type is **dc**, the value is a Direct Connect gateway ID.
     * + If the route type is **cc**, the value is a Cloud Connection ID.
     */
    nexthop: pulumi.Input<string>;
    /**
     * The region in which to create the VPC route. If omitted, the provider-level
     * region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * - Specifies the route table ID for which a route is to be added.
     * If the value is not set, the route will be added to the *default* route table.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * - Specifies the route type. Currently, the value can be:
     * **ecs**, **eni**, **vip**, **nat**, **peering**, **vpn**, **dc** and **cc**.
     */
    type: pulumi.Input<string>;
    /**
     * - Specifies the VPC for which a route is to be added. Changing this creates a
     * new resource.
     */
    vpcId: pulumi.Input<string>;
}
