// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an ELB monitor resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### TCP Health Check
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const monitorTcp = new huaweicloud.elb.Monitor("monitorTcp", {
 *     poolId: _var.pool_id,
 *     type: "TCP",
 *     delay: 5,
 *     timeout: 3,
 *     maxRetries: 3,
 * });
 * ```
 * ### UDP Health Check
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const monitorUdp = new huaweicloud.elb.Monitor("monitorUdp", {
 *     poolId: _var.pool_id,
 *     type: "UDP_CONNECT",
 *     delay: 5,
 *     timeout: 3,
 *     maxRetries: 3,
 * });
 * ```
 * ### HTTP Health Check
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const monitorHttp = new huaweicloud.elb.Monitor("monitorHttp", {
 *     poolId: _var.pool_id,
 *     type: "HTTP",
 *     delay: 5,
 *     timeout: 3,
 *     maxRetries: 3,
 *     urlPath: "/test",
 *     expectedCodes: "200-202",
 * });
 * ```
 *
 * ## Import
 *
 * ELB monitor can be imported using the monitor ID, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Elb/monitor:Monitor monitor_1 5c20fdad-7288-11eb-b817-0255ac10158b
 * ```
 */
export class Monitor extends pulumi.CustomResource {
    /**
     * Get an existing Monitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitorState, opts?: pulumi.CustomResourceOptions): Monitor {
        return new Monitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Elb/monitor:Monitor';

    /**
     * Returns true if the given object is an instance of Monitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Monitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Monitor.__pulumiType;
    }

    /**
     * The administrative state of the monitor.
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the maximum time between health checks in the unit of second. The value ranges
     * from 1 to 50.
     */
    public readonly delay!: pulumi.Output<number>;
    /**
     * Specifies the expected HTTP status code. Required for HTTP type.
     * You can either specify a single status like "200", or a range like "200-202".
     */
    public readonly expectedCodes!: pulumi.Output<string>;
    /**
     * Specifies the HTTP request method. Required for HTTP type.
     * The default value is *GET*.
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * Specifies the maximum number of consecutive health checks after which the backend
     * servers are declared *healthy*. The value ranges from 1 to 10.
     */
    public readonly maxRetries!: pulumi.Output<number>;
    /**
     * Specifies the health check name. The value contains a maximum of 255 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the id of the pool that this monitor will be assigned to. Changing
     * this creates a new monitor.
     */
    public readonly poolId!: pulumi.Output<string>;
    /**
     * Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
     * the port of the backend server will be used as the health check port.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The region in which to create the ELB monitor resource. If omitted, the
     * provider-level region will be used. Changing this creates a new monitor.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Specifies the health check timeout duration in the unit of second.
     * The value ranges from 1 to 50 and must be less than the `delay` value.
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Specifies the monitor protocol.
     * The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
     * If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the HTTP request path for the health check. Required for HTTP type.
     * The value starts with a slash (/) and contains a maximum of 255 characters.
     */
    public readonly urlPath!: pulumi.Output<string>;

    /**
     * Create a Monitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitorArgs | MonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitorState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["delay"] = state ? state.delay : undefined;
            resourceInputs["expectedCodes"] = state ? state.expectedCodes : undefined;
            resourceInputs["httpMethod"] = state ? state.httpMethod : undefined;
            resourceInputs["maxRetries"] = state ? state.maxRetries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["poolId"] = state ? state.poolId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["urlPath"] = state ? state.urlPath : undefined;
        } else {
            const args = argsOrState as MonitorArgs | undefined;
            if ((!args || args.delay === undefined) && !opts.urn) {
                throw new Error("Missing required property 'delay'");
            }
            if ((!args || args.maxRetries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxRetries'");
            }
            if ((!args || args.poolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolId'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["delay"] = args ? args.delay : undefined;
            resourceInputs["expectedCodes"] = args ? args.expectedCodes : undefined;
            resourceInputs["httpMethod"] = args ? args.httpMethod : undefined;
            resourceInputs["maxRetries"] = args ? args.maxRetries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["poolId"] = args ? args.poolId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["urlPath"] = args ? args.urlPath : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Monitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Monitor resources.
 */
export interface MonitorState {
    /**
     * The administrative state of the monitor.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Specifies the maximum time between health checks in the unit of second. The value ranges
     * from 1 to 50.
     */
    delay?: pulumi.Input<number>;
    /**
     * Specifies the expected HTTP status code. Required for HTTP type.
     * You can either specify a single status like "200", or a range like "200-202".
     */
    expectedCodes?: pulumi.Input<string>;
    /**
     * Specifies the HTTP request method. Required for HTTP type.
     * The default value is *GET*.
     */
    httpMethod?: pulumi.Input<string>;
    /**
     * Specifies the maximum number of consecutive health checks after which the backend
     * servers are declared *healthy*. The value ranges from 1 to 10.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Specifies the health check name. The value contains a maximum of 255 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the id of the pool that this monitor will be assigned to. Changing
     * this creates a new monitor.
     */
    poolId?: pulumi.Input<string>;
    /**
     * Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
     * the port of the backend server will be used as the health check port.
     */
    port?: pulumi.Input<number>;
    /**
     * The region in which to create the ELB monitor resource. If omitted, the
     * provider-level region will be used. Changing this creates a new monitor.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Specifies the health check timeout duration in the unit of second.
     * The value ranges from 1 to 50 and must be less than the `delay` value.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the monitor protocol.
     * The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
     * If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the HTTP request path for the health check. Required for HTTP type.
     * The value starts with a slash (/) and contains a maximum of 255 characters.
     */
    urlPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Monitor resource.
 */
export interface MonitorArgs {
    /**
     * The administrative state of the monitor.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Specifies the maximum time between health checks in the unit of second. The value ranges
     * from 1 to 50.
     */
    delay: pulumi.Input<number>;
    /**
     * Specifies the expected HTTP status code. Required for HTTP type.
     * You can either specify a single status like "200", or a range like "200-202".
     */
    expectedCodes?: pulumi.Input<string>;
    /**
     * Specifies the HTTP request method. Required for HTTP type.
     * The default value is *GET*.
     */
    httpMethod?: pulumi.Input<string>;
    /**
     * Specifies the maximum number of consecutive health checks after which the backend
     * servers are declared *healthy*. The value ranges from 1 to 10.
     */
    maxRetries: pulumi.Input<number>;
    /**
     * Specifies the health check name. The value contains a maximum of 255 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the id of the pool that this monitor will be assigned to. Changing
     * this creates a new monitor.
     */
    poolId: pulumi.Input<string>;
    /**
     * Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
     * the port of the backend server will be used as the health check port.
     */
    port?: pulumi.Input<number>;
    /**
     * The region in which to create the ELB monitor resource. If omitted, the
     * provider-level region will be used. Changing this creates a new monitor.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Specifies the health check timeout duration in the unit of second.
     * The value ranges from 1 to 50 and must be less than the `delay` value.
     */
    timeout: pulumi.Input<number>;
    /**
     * Specifies the monitor protocol.
     * The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
     * If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
     */
    type: pulumi.Input<string>;
    /**
     * Specifies the HTTP request path for the health check. Required for HTTP type.
     * The value starts with a slash (/) and contains a maximum of 255 characters.
     */
    urlPath?: pulumi.Input<string>;
}
