// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages CDM cluster resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### create a cdm cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.requireObject("name");
 * const flavorId = config.requireObject("flavorId");
 * const availabilityZone = config.requireObject("availabilityZone");
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const secgroupId = config.requireObject("secgroupId");
 * const test = huaweicloud.Cdm.getFlavors({});
 * const cluster = new huaweicloud.cdm.Cluster("cluster", {
 *     availabilityZone: availabilityZone,
 *     flavorId: test.then(test => test.flavors?[0]?.id),
 *     subnetId: subnetId,
 *     vpcId: vpcId,
 *     securityGroupId: secgroupId,
 * });
 * ```
 *
 * ## Import
 *
 * Clusters can be imported by `id`. For example,
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cdm/cluster:Cluster test b11b407c-e604-4e8d-8bc4-92398320b847
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`email` and `phone_num`.
 *
 * It is generally recommended running `terraform plan` after importing a cluster.
 *
 * You can then decide if changes should be applied to the cluster, or the resource definition should be updated to align with the cluster. Also you can ignore changes as below. resource "huaweicloud_cdm_cluster" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  email, phone_num,
 *
 *  ]
 *
 *  } }
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cdm/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Specifies available zone.
     * Changing this parameter will create a new resource.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Create time. The format is: `YYYY-MM-DDThh:mm:ss`.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Specifies email address for receiving notifications when a table/file migration
     * job fails or an EIP exception occurs. The max number is 5. Changing this parameter will create a new resource.
     */
    public readonly emails!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the enterprise project id.
     * Changing this parameter will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string | undefined>;
    /**
     * Specifies flavor id. Changing this parameter will create a new resource.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * Instance list. Structure is documented below.
     */
    public /*out*/ readonly instances!: pulumi.Output<outputs.Cdm.ClusterInstance[]>;
    /**
     * Specifies Whether to enable auto shutdown. The auto shutdown and scheduled
     * startup/shutdown functions cannot be enabled at the same time. When auto shutdown is enabled, if no job is running in
     * the cluster and no scheduled job is created, a cluster will be automatically shut down 15 minutes after it starts
     * running to reduce costs. The default value is `false`. Changing this parameter will create a new resource.
     */
    public readonly isAutoOff!: pulumi.Output<boolean>;
    /**
     * Specifies cluster name. Changing this parameter will create a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies phone number for receiving notifications when a table/file
     * migration job fails or an EIP exception occurs. The max number is 5. Changing this parameter will create a new resource.
     */
    public readonly phoneNums!: pulumi.Output<string[] | undefined>;
    /**
     * EIP bound to the cluster.
     */
    public /*out*/ readonly publicEndpoint!: pulumi.Output<string>;
    /**
     * Public IP.
     */
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    /**
     * The region in which to create the cluster resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies time for scheduled startup of a CDM cluster.
     * The CDM cluster starts at this time every day. The scheduled startup/shutdown and auto shutdown function cannot be
     * enabled at the same time. The time format is `hh:mm:ss`. Changing this parameter will create a new resource.
     */
    public readonly scheduleBootTime!: pulumi.Output<string>;
    /**
     * Specifies time for scheduled shutdown of a CDM cluster.
     * The system shuts down directly at this time every day without waiting for unfinished jobs to complete.
     * The scheduled startup/shutdown and auto shutdown function cannot be enabled at the same time.
     * The time format is `hh:mm:ss`. Changing this parameter will create a new resource.
     */
    public readonly scheduleOffTime!: pulumi.Output<string>;
    /**
     * Specifies security group ID.
     * Changing this parameter will create a new resource.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies subnet ID. Changing this parameter will create a new resource.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Specifies cluster version. Default value is `2.8.6.2`.
     * Changing this parameter will create a new resource.
     */
    public readonly version!: pulumi.Output<string | undefined>;
    /**
     * Specifies VPC ID. Changing this parameter will create a new resource.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["emails"] = state ? state.emails : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["instances"] = state ? state.instances : undefined;
            resourceInputs["isAutoOff"] = state ? state.isAutoOff : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["phoneNums"] = state ? state.phoneNums : undefined;
            resourceInputs["publicEndpoint"] = state ? state.publicEndpoint : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["scheduleBootTime"] = state ? state.scheduleBootTime : undefined;
            resourceInputs["scheduleOffTime"] = state ? state.scheduleOffTime : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.flavorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavorId'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["emails"] = args ? args.emails : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["isAutoOff"] = args ? args.isAutoOff : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["phoneNums"] = args ? args.phoneNums : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["scheduleBootTime"] = args ? args.scheduleBootTime : undefined;
            resourceInputs["scheduleOffTime"] = args ? args.scheduleOffTime : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["instances"] = undefined /*out*/;
            resourceInputs["publicEndpoint"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Specifies available zone.
     * Changing this parameter will create a new resource.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Create time. The format is: `YYYY-MM-DDThh:mm:ss`.
     */
    created?: pulumi.Input<string>;
    /**
     * Specifies email address for receiving notifications when a table/file migration
     * job fails or an EIP exception occurs. The max number is 5. Changing this parameter will create a new resource.
     */
    emails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the enterprise project id.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies flavor id. Changing this parameter will create a new resource.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * Instance list. Structure is documented below.
     */
    instances?: pulumi.Input<pulumi.Input<inputs.Cdm.ClusterInstance>[]>;
    /**
     * Specifies Whether to enable auto shutdown. The auto shutdown and scheduled
     * startup/shutdown functions cannot be enabled at the same time. When auto shutdown is enabled, if no job is running in
     * the cluster and no scheduled job is created, a cluster will be automatically shut down 15 minutes after it starts
     * running to reduce costs. The default value is `false`. Changing this parameter will create a new resource.
     */
    isAutoOff?: pulumi.Input<boolean>;
    /**
     * Specifies cluster name. Changing this parameter will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies phone number for receiving notifications when a table/file
     * migration job fails or an EIP exception occurs. The max number is 5. Changing this parameter will create a new resource.
     */
    phoneNums?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * EIP bound to the cluster.
     */
    publicEndpoint?: pulumi.Input<string>;
    /**
     * Public IP.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * The region in which to create the cluster resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies time for scheduled startup of a CDM cluster.
     * The CDM cluster starts at this time every day. The scheduled startup/shutdown and auto shutdown function cannot be
     * enabled at the same time. The time format is `hh:mm:ss`. Changing this parameter will create a new resource.
     */
    scheduleBootTime?: pulumi.Input<string>;
    /**
     * Specifies time for scheduled shutdown of a CDM cluster.
     * The system shuts down directly at this time every day without waiting for unfinished jobs to complete.
     * The scheduled startup/shutdown and auto shutdown function cannot be enabled at the same time.
     * The time format is `hh:mm:ss`. Changing this parameter will create a new resource.
     */
    scheduleOffTime?: pulumi.Input<string>;
    /**
     * Specifies security group ID.
     * Changing this parameter will create a new resource.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Status.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies subnet ID. Changing this parameter will create a new resource.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Specifies cluster version. Default value is `2.8.6.2`.
     * Changing this parameter will create a new resource.
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies VPC ID. Changing this parameter will create a new resource.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Specifies available zone.
     * Changing this parameter will create a new resource.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Specifies email address for receiving notifications when a table/file migration
     * job fails or an EIP exception occurs. The max number is 5. Changing this parameter will create a new resource.
     */
    emails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the enterprise project id.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies flavor id. Changing this parameter will create a new resource.
     */
    flavorId: pulumi.Input<string>;
    /**
     * Specifies Whether to enable auto shutdown. The auto shutdown and scheduled
     * startup/shutdown functions cannot be enabled at the same time. When auto shutdown is enabled, if no job is running in
     * the cluster and no scheduled job is created, a cluster will be automatically shut down 15 minutes after it starts
     * running to reduce costs. The default value is `false`. Changing this parameter will create a new resource.
     */
    isAutoOff?: pulumi.Input<boolean>;
    /**
     * Specifies cluster name. Changing this parameter will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies phone number for receiving notifications when a table/file
     * migration job fails or an EIP exception occurs. The max number is 5. Changing this parameter will create a new resource.
     */
    phoneNums?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to create the cluster resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies time for scheduled startup of a CDM cluster.
     * The CDM cluster starts at this time every day. The scheduled startup/shutdown and auto shutdown function cannot be
     * enabled at the same time. The time format is `hh:mm:ss`. Changing this parameter will create a new resource.
     */
    scheduleBootTime?: pulumi.Input<string>;
    /**
     * Specifies time for scheduled shutdown of a CDM cluster.
     * The system shuts down directly at this time every day without waiting for unfinished jobs to complete.
     * The scheduled startup/shutdown and auto shutdown function cannot be enabled at the same time.
     * The time format is `hh:mm:ss`. Changing this parameter will create a new resource.
     */
    scheduleOffTime?: pulumi.Input<string>;
    /**
     * Specifies security group ID.
     * Changing this parameter will create a new resource.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * Specifies subnet ID. Changing this parameter will create a new resource.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Specifies cluster version. Default value is `2.8.6.2`.
     * Changing this parameter will create a new resource.
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies VPC ID. Changing this parameter will create a new resource.
     */
    vpcId: pulumi.Input<string>;
}