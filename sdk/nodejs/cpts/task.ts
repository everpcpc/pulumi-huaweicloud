// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a pressure Test Task resource within HuaweiCloud CPTS.
 * The task resource only supports serial execution mode.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const testProject = new huaweicloud.cpts.Project("testProject", {});
 * const testTask = new huaweicloud.cpts.Task("testTask", {projectId: testProject.id});
 * ```
 *
 * ## Import
 *
 * Tasks can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cpts/task:Task test 1090
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`operation`. It is generally recommended running `terraform plan` after importing an instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. resource "huaweicloud_cpts_task" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  operation,
 *
 *  ]
 *
 *  } }
 */
export class Task extends pulumi.CustomResource {
    /**
     * Get an existing Task resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TaskState, opts?: pulumi.CustomResourceOptions): Task {
        return new Task(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cpts/task:Task';

    /**
     * Returns true if the given object is an instance of Task.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Task {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Task.__pulumiType;
    }

    /**
     * Specifies benchmark concurrency of the task, the value range is 0 to
     * 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
     * `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
     */
    public readonly benchmarkConcurrency!: pulumi.Output<number | undefined>;
    /**
     * Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
     * is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
     */
    public readonly clusterId!: pulumi.Output<number | undefined>;
    /**
     * Specifies the name of the task, which can contain a maximum of 42 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies whether to enable the task or stop the task. The options are as follows:
     * + **enable**: Starting the pressure test task.
     * + **stop**: Stop the pressure test tasks.
     */
    public readonly operation!: pulumi.Output<string | undefined>;
    /**
     * Specifies the CPTS project ID which the task belongs to.
     * Changing this parameter will create a new resource.
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Specifies the region in which to create the task resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The status of the task. The options are as follows:
     * + **0**: The task is running.
     * + **2**: The task is idle.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;

    /**
     * Create a Task resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TaskArgs | TaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TaskState | undefined;
            resourceInputs["benchmarkConcurrency"] = state ? state.benchmarkConcurrency : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operation"] = state ? state.operation : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as TaskArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["benchmarkConcurrency"] = args ? args.benchmarkConcurrency : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operation"] = args ? args.operation : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Task.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Task resources.
 */
export interface TaskState {
    /**
     * Specifies benchmark concurrency of the task, the value range is 0 to
     * 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
     * `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
     */
    benchmarkConcurrency?: pulumi.Input<number>;
    /**
     * Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
     * is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
     */
    clusterId?: pulumi.Input<number>;
    /**
     * Specifies the name of the task, which can contain a maximum of 42 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the task or stop the task. The options are as follows:
     * + **enable**: Starting the pressure test task.
     * + **stop**: Stop the pressure test tasks.
     */
    operation?: pulumi.Input<string>;
    /**
     * Specifies the CPTS project ID which the task belongs to.
     * Changing this parameter will create a new resource.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Specifies the region in which to create the task resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the task. The options are as follows:
     * + **0**: The task is running.
     * + **2**: The task is idle.
     */
    status?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Task resource.
 */
export interface TaskArgs {
    /**
     * Specifies benchmark concurrency of the task, the value range is 0 to
     * 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
     * `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
     */
    benchmarkConcurrency?: pulumi.Input<number>;
    /**
     * Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
     * is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
     */
    clusterId?: pulumi.Input<number>;
    /**
     * Specifies the name of the task, which can contain a maximum of 42 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the task or stop the task. The options are as follows:
     * + **enable**: Starting the pressure test task.
     * + **stop**: Stop the pressure test tasks.
     */
    operation?: pulumi.Input<string>;
    /**
     * Specifies the CPTS project ID which the task belongs to.
     * Changing this parameter will create a new resource.
     */
    projectId: pulumi.Input<number>;
    /**
     * Specifies the region in which to create the task resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}
