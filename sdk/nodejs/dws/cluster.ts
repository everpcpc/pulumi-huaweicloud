// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages Cluster in the Data Warehouse Service.
 *
 * ## Example Usage
 * ### Dws Cluster Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.requireObject("availabilityZone");
 * const networkId = config.requireObject("networkId");
 * const vpcId = config.requireObject("vpcId");
 * const userName = config.requireObject("userName");
 * const userPwd = config.requireObject("userPwd");
 * const dwsClusterName = config.requireObject("dwsClusterName");
 * const secgroup = new huaweicloud.vpc.Secgroup("secgroup", {description: "terraform security group"});
 * const cluster = new huaweicloud.dws.Cluster("cluster", {
 *     nodeType: "dws.m3.xlarge",
 *     numberOfNode: 3,
 *     networkId: networkId,
 *     vpcId: vpcId,
 *     securityGroupId: secgroup.id,
 *     availabilityZone: availabilityZone,
 *     userName: userName,
 *     userPwd: userPwd,
 * });
 * ```
 *
 * ## Import
 *
 * Cluster can be imported using the following format
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dws/cluster:Cluster test 47ad727e-9dcc-4833-bde0-bb298607c719
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`user_pwd`, `number_of_cn`. It is generally recommended running `terraform plan` after importing a cluster. You can then decide if changes should be applied to the cluster, or the resource definition should be updated to align with the cluster. Also you can ignore changes as below. resource "huaweicloud_dws_cluster" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  user_pwd, number_of_cn,
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
    public static readonly __pulumiType = 'huaweicloud:Dws/cluster:Cluster';

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
     * AZ in a cluster.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * View the private network connection information about the cluster. Structure is documented below.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.Dws.ClusterEndpoint[]>;
    /**
     * Specifies the enterprise project id of the dws cluster,
     * Value 0 indicates the default enterprise project.
     * Changing this parameter will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Cluster name, which must be unique and contains 4 to 64 characters, which
     * consist of letters, digits, hyphens(-), or underscores(_) only and must start with a letter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network ID, which is used for configuring cluster network.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Node type.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * Specifies the number of CN. If you use a large-scale cluster, deploy
     * multiple CNs.
     */
    public readonly numberOfCn!: pulumi.Output<number | undefined>;
    /**
     * Number of nodes in a cluster. The value ranges from 3 to 32. When expanding,
     * add at least 3 nodes.
     */
    public readonly numberOfNode!: pulumi.Output<number>;
    /**
     * Service port of a cluster (8000 to 10000). The default value is 8000.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * List of private network IP address.
     */
    public /*out*/ readonly privateIps!: pulumi.Output<string[]>;
    /**
     * Public network connection information about the cluster. If the value is not specified, the
     * public network connection information is not used by default Structure is documented below.
     */
    public /*out*/ readonly publicEndpoints!: pulumi.Output<outputs.Dws.ClusterPublicEndpoint[]>;
    /**
     * A nested object resource Structure is documented below.
     */
    public readonly publicIp!: pulumi.Output<outputs.Dws.ClusterPublicIp>;
    /**
     * The recent event number.
     */
    public /*out*/ readonly recentEvent!: pulumi.Output<number>;
    /**
     * The region in which to create the cluster resource. If omitted, the
     * provider-level region will be used. Changing this creates a new cluster resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * ID of a security group. The ID is used for configuring cluster
     * network.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Cluster status, which can be one of the following:  CREATING AVAILABLE UNAVAILABLE CREATION FAILED.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Sub-status of clusters in the AVAILABLE state. The value can be one of the following:  NORMAL READONLY
     * REDISTRIBUTING REDISTRIBUTION-FAILURE UNBALANCED UNBALANCED | READONLY DEGRADED DEGRADED | READONLY DEGRADED |
     * UNBALANCED UNBALANCED | REDISTRIBUTING UNBALANCED | REDISTRIBUTION-FAILURE READONLY | REDISTRIBUTION-FAILURE
     * UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED |
     * REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY
     */
    public /*out*/ readonly subStatus!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Cluster management task. The value can be one of the following:
     * RESTORING SNAPSHOTTING GROWING REBOOTING SETTING_CONFIGURATION CONFIGURING_EXT_DATASOURCE DELETING_EXT_DATASOURCE
     * REBOOT_FAILURE RESIZE_FAILURE
     */
    public /*out*/ readonly taskStatus!: pulumi.Output<string>;
    /**
     * Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;
    /**
     * Administrator username for logging in to a data warehouse cluster The
     * administrator username must:  Consist of lowercase letters, digits, or underscores. Start with a lowercase letter or
     * an underscore. Contain 1 to 63 characters. Cannot be a keyword of the DWS database.
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Administrator password for logging in to a data warehouse cluster A password
     * must conform to the following rules:  Contains 8 to 32 characters. Cannot be the same as the username or the username
     * written in reverse order. Contains three types of the following:
     * Lowercase letters Uppercase letters Digits Special characters
     * ~!@#%^&*()-_=+|[{}];:,<.>/?
     */
    public readonly userPwd!: pulumi.Output<string>;
    /**
     * Data warehouse version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * VPC ID, which is used for configuring cluster network.
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
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["numberOfCn"] = state ? state.numberOfCn : undefined;
            resourceInputs["numberOfNode"] = state ? state.numberOfNode : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privateIps"] = state ? state.privateIps : undefined;
            resourceInputs["publicEndpoints"] = state ? state.publicEndpoints : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["recentEvent"] = state ? state.recentEvent : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subStatus"] = state ? state.subStatus : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["taskStatus"] = state ? state.taskStatus : undefined;
            resourceInputs["updated"] = state ? state.updated : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["userPwd"] = state ? state.userPwd : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            if ((!args || args.numberOfNode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'numberOfNode'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            if ((!args || args.userPwd === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPwd'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["numberOfCn"] = args ? args.numberOfCn : undefined;
            resourceInputs["numberOfNode"] = args ? args.numberOfNode : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["publicIp"] = args ? args.publicIp : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["userPwd"] = args ? args.userPwd : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["privateIps"] = undefined /*out*/;
            resourceInputs["publicEndpoints"] = undefined /*out*/;
            resourceInputs["recentEvent"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subStatus"] = undefined /*out*/;
            resourceInputs["taskStatus"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
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
     * AZ in a cluster.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ
     */
    created?: pulumi.Input<string>;
    /**
     * View the private network connection information about the cluster. Structure is documented below.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.Dws.ClusterEndpoint>[]>;
    /**
     * Specifies the enterprise project id of the dws cluster,
     * Value 0 indicates the default enterprise project.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Cluster name, which must be unique and contains 4 to 64 characters, which
     * consist of letters, digits, hyphens(-), or underscores(_) only and must start with a letter.
     */
    name?: pulumi.Input<string>;
    /**
     * Network ID, which is used for configuring cluster network.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Node type.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Specifies the number of CN. If you use a large-scale cluster, deploy
     * multiple CNs.
     */
    numberOfCn?: pulumi.Input<number>;
    /**
     * Number of nodes in a cluster. The value ranges from 3 to 32. When expanding,
     * add at least 3 nodes.
     */
    numberOfNode?: pulumi.Input<number>;
    /**
     * Service port of a cluster (8000 to 10000). The default value is 8000.
     */
    port?: pulumi.Input<number>;
    /**
     * List of private network IP address.
     */
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Public network connection information about the cluster. If the value is not specified, the
     * public network connection information is not used by default Structure is documented below.
     */
    publicEndpoints?: pulumi.Input<pulumi.Input<inputs.Dws.ClusterPublicEndpoint>[]>;
    /**
     * A nested object resource Structure is documented below.
     */
    publicIp?: pulumi.Input<inputs.Dws.ClusterPublicIp>;
    /**
     * The recent event number.
     */
    recentEvent?: pulumi.Input<number>;
    /**
     * The region in which to create the cluster resource. If omitted, the
     * provider-level region will be used. Changing this creates a new cluster resource.
     */
    region?: pulumi.Input<string>;
    /**
     * ID of a security group. The ID is used for configuring cluster
     * network.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Cluster status, which can be one of the following:  CREATING AVAILABLE UNAVAILABLE CREATION FAILED.
     */
    status?: pulumi.Input<string>;
    /**
     * Sub-status of clusters in the AVAILABLE state. The value can be one of the following:  NORMAL READONLY
     * REDISTRIBUTING REDISTRIBUTION-FAILURE UNBALANCED UNBALANCED | READONLY DEGRADED DEGRADED | READONLY DEGRADED |
     * UNBALANCED UNBALANCED | REDISTRIBUTING UNBALANCED | REDISTRIBUTION-FAILURE READONLY | REDISTRIBUTION-FAILURE
     * UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED |
     * REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY
     */
    subStatus?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the cluster.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Cluster management task. The value can be one of the following:
     * RESTORING SNAPSHOTTING GROWING REBOOTING SETTING_CONFIGURATION CONFIGURING_EXT_DATASOURCE DELETING_EXT_DATASOURCE
     * REBOOT_FAILURE RESIZE_FAILURE
     */
    taskStatus?: pulumi.Input<string>;
    /**
     * Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ
     */
    updated?: pulumi.Input<string>;
    /**
     * Administrator username for logging in to a data warehouse cluster The
     * administrator username must:  Consist of lowercase letters, digits, or underscores. Start with a lowercase letter or
     * an underscore. Contain 1 to 63 characters. Cannot be a keyword of the DWS database.
     */
    userName?: pulumi.Input<string>;
    /**
     * Administrator password for logging in to a data warehouse cluster A password
     * must conform to the following rules:  Contains 8 to 32 characters. Cannot be the same as the username or the username
     * written in reverse order. Contains three types of the following:
     * Lowercase letters Uppercase letters Digits Special characters
     * ~!@#%^&*()-_=+|[{}];:,<.>/?
     */
    userPwd?: pulumi.Input<string>;
    /**
     * Data warehouse version.
     */
    version?: pulumi.Input<string>;
    /**
     * VPC ID, which is used for configuring cluster network.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * AZ in a cluster.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project id of the dws cluster,
     * Value 0 indicates the default enterprise project.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Cluster name, which must be unique and contains 4 to 64 characters, which
     * consist of letters, digits, hyphens(-), or underscores(_) only and must start with a letter.
     */
    name?: pulumi.Input<string>;
    /**
     * Network ID, which is used for configuring cluster network.
     */
    networkId: pulumi.Input<string>;
    /**
     * Node type.
     */
    nodeType: pulumi.Input<string>;
    /**
     * Specifies the number of CN. If you use a large-scale cluster, deploy
     * multiple CNs.
     */
    numberOfCn?: pulumi.Input<number>;
    /**
     * Number of nodes in a cluster. The value ranges from 3 to 32. When expanding,
     * add at least 3 nodes.
     */
    numberOfNode: pulumi.Input<number>;
    /**
     * Service port of a cluster (8000 to 10000). The default value is 8000.
     */
    port?: pulumi.Input<number>;
    /**
     * A nested object resource Structure is documented below.
     */
    publicIp?: pulumi.Input<inputs.Dws.ClusterPublicIp>;
    /**
     * The region in which to create the cluster resource. If omitted, the
     * provider-level region will be used. Changing this creates a new cluster resource.
     */
    region?: pulumi.Input<string>;
    /**
     * ID of a security group. The ID is used for configuring cluster
     * network.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the cluster.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Administrator username for logging in to a data warehouse cluster The
     * administrator username must:  Consist of lowercase letters, digits, or underscores. Start with a lowercase letter or
     * an underscore. Contain 1 to 63 characters. Cannot be a keyword of the DWS database.
     */
    userName: pulumi.Input<string>;
    /**
     * Administrator password for logging in to a data warehouse cluster A password
     * must conform to the following rules:  Contains 8 to 32 characters. Cannot be the same as the username or the username
     * written in reverse order. Contains three types of the following:
     * Lowercase letters Uppercase letters Digits Special characters
     * ~!@#%^&*()-_=+|[{}];:,<.>/?
     */
    userPwd: pulumi.Input<string>;
    /**
     * VPC ID, which is used for configuring cluster network.
     */
    vpcId: pulumi.Input<string>;
}
