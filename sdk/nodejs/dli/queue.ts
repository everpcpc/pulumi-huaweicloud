// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages DLI Queue resource within HuaweiCloud
 *
 * ## Example Usage
 * ### Create a queue
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const queue = new huaweicloud.Dli.Queue("queue", {
 *     cuCount: 16,
 *     tags: {
 *         foo: "bar",
 *         key: "value",
 *     },
 * });
 * ```
 * ### Create a queue with CIDR Block
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const queue = new huaweicloud.Dli.Queue("queue", {
 *     cuCount: 16,
 *     resourceMode: 1,
 *     tags: {
 *         foo: "bar",
 *         key: "value",
 *     },
 *     vpcCidr: "172.16.0.0/14",
 * });
 * ```
 *
 * ## Import
 *
 * DLI queue can be imported by
 *
 * `id`. For example,
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dli/queue:Queue example abc123
 * ```
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Dli/queue:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * Time when a queue is created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * Minimum number of CUs that are bound to a queue. Initial value can be `16`,
     * `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of 16
     */
    public readonly cuCount!: pulumi.Output<number>;
    /**
     * Description of a queue. Changing this parameter will create a new
     * resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Enterprise project ID. The value 0 indicates the default
     * enterprise project. Changing this parameter will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Indicates the queue feature. Changing this parameter will create a new
     * resource. The options are as follows:
     * + basic: basic type (default value)
     * + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
     */
    public readonly feature!: pulumi.Output<string | undefined>;
    /**
     * @deprecated management_subnet_cidr is Deprecated
     */
    public readonly managementSubnetCidr!: pulumi.Output<string | undefined>;
    /**
     * Name of a queue. Name of a newly created resource queue. The name can contain
     * only digits, letters, and underscores (\_), but cannot contain only digits or start with an underscore (_). Length
     * range: 1 to 128 characters. Changing this parameter will create a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * CPU architecture of queue compute resources. Changing this parameter will
     * create a new resource. The options are as follows:
     * + x8664 : default value
     * + aarch64
     */
    public readonly platform!: pulumi.Output<string | undefined>;
    /**
     * Indicates the queue type. Changing this parameter will create a new
     * resource. The options are as follows:
     * + sql
     * + general
     */
    public readonly queueType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the region in which to create the dli queue resource. If omitted,
     * the provider-level region will be used. Changing this will create a new VPC channel resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Queue resource mode. Changing this parameter will create a new
     * resource. The options are as follows:
     * + 0: indicates the shared resource mode.
     * + 1: indicates the exclusive resource mode.
     */
    public readonly resourceMode!: pulumi.Output<number | undefined>;
    /**
     * @deprecated subnet_cidr is Deprecated
     */
    public readonly subnetCidr!: pulumi.Output<string | undefined>;
    /**
     * Label of a queue. Changing this parameter will create a new resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The CIDR block of a queue. If use DLI enhanced datasource connections, the CIDR block
     * cannot be the same as that of the data source.
     * The CIDR blocks supported by different CU specifications:
     */
    public readonly vpcCidr!: pulumi.Output<string>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueArgs | QueueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QueueState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["cuCount"] = state ? state.cuCount : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["feature"] = state ? state.feature : undefined;
            resourceInputs["managementSubnetCidr"] = state ? state.managementSubnetCidr : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platform"] = state ? state.platform : undefined;
            resourceInputs["queueType"] = state ? state.queueType : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["resourceMode"] = state ? state.resourceMode : undefined;
            resourceInputs["subnetCidr"] = state ? state.subnetCidr : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcCidr"] = state ? state.vpcCidr : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            if ((!args || args.cuCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cuCount'");
            }
            resourceInputs["cuCount"] = args ? args.cuCount : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["feature"] = args ? args.feature : undefined;
            resourceInputs["managementSubnetCidr"] = args ? args.managementSubnetCidr : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["platform"] = args ? args.platform : undefined;
            resourceInputs["queueType"] = args ? args.queueType : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["resourceMode"] = args ? args.resourceMode : undefined;
            resourceInputs["subnetCidr"] = args ? args.subnetCidr : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcCidr"] = args ? args.vpcCidr : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Queue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * Time when a queue is created.
     */
    createTime?: pulumi.Input<number>;
    /**
     * Minimum number of CUs that are bound to a queue. Initial value can be `16`,
     * `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of 16
     */
    cuCount?: pulumi.Input<number>;
    /**
     * Description of a queue. Changing this parameter will create a new
     * resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Enterprise project ID. The value 0 indicates the default
     * enterprise project. Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Indicates the queue feature. Changing this parameter will create a new
     * resource. The options are as follows:
     * + basic: basic type (default value)
     * + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
     */
    feature?: pulumi.Input<string>;
    /**
     * @deprecated management_subnet_cidr is Deprecated
     */
    managementSubnetCidr?: pulumi.Input<string>;
    /**
     * Name of a queue. Name of a newly created resource queue. The name can contain
     * only digits, letters, and underscores (\_), but cannot contain only digits or start with an underscore (_). Length
     * range: 1 to 128 characters. Changing this parameter will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * CPU architecture of queue compute resources. Changing this parameter will
     * create a new resource. The options are as follows:
     * + x8664 : default value
     * + aarch64
     */
    platform?: pulumi.Input<string>;
    /**
     * Indicates the queue type. Changing this parameter will create a new
     * resource. The options are as follows:
     * + sql
     * + general
     */
    queueType?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the dli queue resource. If omitted,
     * the provider-level region will be used. Changing this will create a new VPC channel resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Queue resource mode. Changing this parameter will create a new
     * resource. The options are as follows:
     * + 0: indicates the shared resource mode.
     * + 1: indicates the exclusive resource mode.
     */
    resourceMode?: pulumi.Input<number>;
    /**
     * @deprecated subnet_cidr is Deprecated
     */
    subnetCidr?: pulumi.Input<string>;
    /**
     * Label of a queue. Changing this parameter will create a new resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The CIDR block of a queue. If use DLI enhanced datasource connections, the CIDR block
     * cannot be the same as that of the data source.
     * The CIDR blocks supported by different CU specifications:
     */
    vpcCidr?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * Minimum number of CUs that are bound to a queue. Initial value can be `16`,
     * `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of 16
     */
    cuCount: pulumi.Input<number>;
    /**
     * Description of a queue. Changing this parameter will create a new
     * resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Enterprise project ID. The value 0 indicates the default
     * enterprise project. Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Indicates the queue feature. Changing this parameter will create a new
     * resource. The options are as follows:
     * + basic: basic type (default value)
     * + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
     */
    feature?: pulumi.Input<string>;
    /**
     * @deprecated management_subnet_cidr is Deprecated
     */
    managementSubnetCidr?: pulumi.Input<string>;
    /**
     * Name of a queue. Name of a newly created resource queue. The name can contain
     * only digits, letters, and underscores (\_), but cannot contain only digits or start with an underscore (_). Length
     * range: 1 to 128 characters. Changing this parameter will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * CPU architecture of queue compute resources. Changing this parameter will
     * create a new resource. The options are as follows:
     * + x8664 : default value
     * + aarch64
     */
    platform?: pulumi.Input<string>;
    /**
     * Indicates the queue type. Changing this parameter will create a new
     * resource. The options are as follows:
     * + sql
     * + general
     */
    queueType?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the dli queue resource. If omitted,
     * the provider-level region will be used. Changing this will create a new VPC channel resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Queue resource mode. Changing this parameter will create a new
     * resource. The options are as follows:
     * + 0: indicates the shared resource mode.
     * + 1: indicates the exclusive resource mode.
     */
    resourceMode?: pulumi.Input<number>;
    /**
     * @deprecated subnet_cidr is Deprecated
     */
    subnetCidr?: pulumi.Input<string>;
    /**
     * Label of a queue. Changing this parameter will create a new resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The CIDR block of a queue. If use DLI enhanced datasource connections, the CIDR block
     * cannot be the same as that of the data source.
     * The CIDR blocks supported by different CU specifications:
     */
    vpcCidr?: pulumi.Input<string>;
}