// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an OMS migration task resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const sourceRegion = config.requireObject("sourceRegion");
 * const sourceBucket = config.requireObject("sourceBucket");
 * const sourceAccessKey = config.requireObject("sourceAccessKey");
 * const sourceSecretKey = config.requireObject("sourceSecretKey");
 * const destRegion = config.requireObject("destRegion");
 * const destBucket = config.requireObject("destBucket");
 * const destAccessKey = config.requireObject("destAccessKey");
 * const destSecretKey = config.requireObject("destSecretKey");
 * const topicUrn = config.requireObject("topicUrn");
 * const test = new huaweicloud.oms.MigrationTask("test", {
 *     sourceObject: {
 *         dataSource: "Aliyun",
 *         region: sourceRegion,
 *         bucket: sourceBucket,
 *         accessKey: sourceAccessKey,
 *         secretKey: sourceSecretKey,
 *         objects: [""],
 *     },
 *     destinationObject: {
 *         region: destRegion,
 *         bucket: destBucket,
 *         accessKey: destAccessKey,
 *         secretKey: destSecretKey,
 *     },
 *     type: "object",
 *     description: "test task",
 *     bandwidthPolicies: [{
 *         maxBandwidth: 2,
 *         start: "15:00",
 *         end: "16:00",
 *     }],
 *     smnConfig: {
 *         topicUrn: topicUrn,
 *         triggerConditions: [
 *             "FAILURE",
 *             "SUCCESS",
 *         ],
 *     },
 * });
 * ```
 */
export class MigrationTask extends pulumi.CustomResource {
    /**
     * Get an existing MigrationTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MigrationTaskState, opts?: pulumi.CustomResourceOptions): MigrationTask {
        return new MigrationTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Oms/migrationTask:MigrationTask';

    /**
     * Returns true if the given object is an instance of MigrationTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MigrationTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MigrationTask.__pulumiType;
    }

    /**
     * Specifies the traffic limit rules. Each element in the array
     * corresponds to the maximum bandwidth of a time segment. A maximum of 5 time segments are allowed, and the time
     * segments must not overlap. The object structure is  documented below.
     */
    public readonly bandwidthPolicies!: pulumi.Output<outputs.Oms.MigrationTaskBandwidthPolicy[] | undefined>;
    /**
     * Specifies the description of the task.
     * Changing this creates a new resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the destination information. The object
     * structure is documented below. Changing this creates a new resource.
     */
    public readonly destinationObject!: pulumi.Output<outputs.Oms.MigrationTaskDestinationObject>;
    /**
     * Specifies whether to record failed objects. If this
     * function is enabled, information about objects that fail to be migrated will be stored in the destination bucket.
     * Default value: **true**. Changing this creates a new resource.
     */
    public readonly enableFailedObjectRecording!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to enable the KMS encryption function.
     * Default value: **false**. Changing this creates a new resource.
     */
    public readonly enableKms!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to automatically restore the archive data. If enabled,
     * archive data is automatically restored and migrated. Default value: **false**. Changing this creates a new resource.
     */
    public readonly enableRestore!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies a time in format **yyyy-MM-dd HH:mm:ss**,
     * e.g. **2006-01-02 15:04:05**. The system migrates only the objects that are modified after the specified time.
     * No time is specified by default. Changing this creates a new resource.
     */
    public readonly migrateSince!: pulumi.Output<string>;
    /**
     * The name of the migration task.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region where the destination bucket is located.
     * Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the SMN message sending configuration.
     * The object structure is  documented below. Changing this creates a new resource.
     */
    public readonly smnConfig!: pulumi.Output<outputs.Oms.MigrationTaskSmnConfig | undefined>;
    /**
     * Specifies the CDN information. If this parameter is contained,
     * using CDN to download source data is supported, the source objects to be migrated are obtained from the CDN domain
     * name during migration. The object structure is documented below.
     * Changing this creates a new resource.
     */
    public readonly sourceCdn!: pulumi.Output<outputs.Oms.MigrationTaskSourceCdn | undefined>;
    /**
     * Specifies the source information. The object
     * structure is documented below. Changing this creates a new resource.
     */
    public readonly sourceObject!: pulumi.Output<outputs.Oms.MigrationTaskSourceObject>;
    /**
     * Specifies whether to start the task. Default value: **true**.
     */
    public readonly startTask!: pulumi.Output<boolean | undefined>;
    /**
     * The status the migration task. The value can be:
     * + **1**: Waiting to migrate.
     * + **2**: Migrating.
     * + **3**: Migration paused.
     * + **4**: Migration failed.
     * + **5**: Migration succeeded.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Specifies the task type. The value can be:
     * + **list**: indicates migrating objects using an object list.
     * + **url_list**: indicates migrating objects using a URL object list.
     * + **object**: indicates migrating selected files or folders.
     * + **prefix**: indicates migrating objects with specified prefixes.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a MigrationTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MigrationTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MigrationTaskArgs | MigrationTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MigrationTaskState | undefined;
            resourceInputs["bandwidthPolicies"] = state ? state.bandwidthPolicies : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationObject"] = state ? state.destinationObject : undefined;
            resourceInputs["enableFailedObjectRecording"] = state ? state.enableFailedObjectRecording : undefined;
            resourceInputs["enableKms"] = state ? state.enableKms : undefined;
            resourceInputs["enableRestore"] = state ? state.enableRestore : undefined;
            resourceInputs["migrateSince"] = state ? state.migrateSince : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["smnConfig"] = state ? state.smnConfig : undefined;
            resourceInputs["sourceCdn"] = state ? state.sourceCdn : undefined;
            resourceInputs["sourceObject"] = state ? state.sourceObject : undefined;
            resourceInputs["startTask"] = state ? state.startTask : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as MigrationTaskArgs | undefined;
            if ((!args || args.destinationObject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationObject'");
            }
            if ((!args || args.sourceObject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceObject'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["bandwidthPolicies"] = args ? args.bandwidthPolicies : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationObject"] = args ? args.destinationObject : undefined;
            resourceInputs["enableFailedObjectRecording"] = args ? args.enableFailedObjectRecording : undefined;
            resourceInputs["enableKms"] = args ? args.enableKms : undefined;
            resourceInputs["enableRestore"] = args ? args.enableRestore : undefined;
            resourceInputs["migrateSince"] = args ? args.migrateSince : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["smnConfig"] = args ? args.smnConfig : undefined;
            resourceInputs["sourceCdn"] = args ? args.sourceCdn : undefined;
            resourceInputs["sourceObject"] = args ? args.sourceObject : undefined;
            resourceInputs["startTask"] = args ? args.startTask : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MigrationTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MigrationTask resources.
 */
export interface MigrationTaskState {
    /**
     * Specifies the traffic limit rules. Each element in the array
     * corresponds to the maximum bandwidth of a time segment. A maximum of 5 time segments are allowed, and the time
     * segments must not overlap. The object structure is  documented below.
     */
    bandwidthPolicies?: pulumi.Input<pulumi.Input<inputs.Oms.MigrationTaskBandwidthPolicy>[]>;
    /**
     * Specifies the description of the task.
     * Changing this creates a new resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the destination information. The object
     * structure is documented below. Changing this creates a new resource.
     */
    destinationObject?: pulumi.Input<inputs.Oms.MigrationTaskDestinationObject>;
    /**
     * Specifies whether to record failed objects. If this
     * function is enabled, information about objects that fail to be migrated will be stored in the destination bucket.
     * Default value: **true**. Changing this creates a new resource.
     */
    enableFailedObjectRecording?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the KMS encryption function.
     * Default value: **false**. Changing this creates a new resource.
     */
    enableKms?: pulumi.Input<boolean>;
    /**
     * Specifies whether to automatically restore the archive data. If enabled,
     * archive data is automatically restored and migrated. Default value: **false**. Changing this creates a new resource.
     */
    enableRestore?: pulumi.Input<boolean>;
    /**
     * Specifies a time in format **yyyy-MM-dd HH:mm:ss**,
     * e.g. **2006-01-02 15:04:05**. The system migrates only the objects that are modified after the specified time.
     * No time is specified by default. Changing this creates a new resource.
     */
    migrateSince?: pulumi.Input<string>;
    /**
     * The name of the migration task.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region where the destination bucket is located.
     * Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the SMN message sending configuration.
     * The object structure is  documented below. Changing this creates a new resource.
     */
    smnConfig?: pulumi.Input<inputs.Oms.MigrationTaskSmnConfig>;
    /**
     * Specifies the CDN information. If this parameter is contained,
     * using CDN to download source data is supported, the source objects to be migrated are obtained from the CDN domain
     * name during migration. The object structure is documented below.
     * Changing this creates a new resource.
     */
    sourceCdn?: pulumi.Input<inputs.Oms.MigrationTaskSourceCdn>;
    /**
     * Specifies the source information. The object
     * structure is documented below. Changing this creates a new resource.
     */
    sourceObject?: pulumi.Input<inputs.Oms.MigrationTaskSourceObject>;
    /**
     * Specifies whether to start the task. Default value: **true**.
     */
    startTask?: pulumi.Input<boolean>;
    /**
     * The status the migration task. The value can be:
     * + **1**: Waiting to migrate.
     * + **2**: Migrating.
     * + **3**: Migration paused.
     * + **4**: Migration failed.
     * + **5**: Migration succeeded.
     */
    status?: pulumi.Input<number>;
    /**
     * Specifies the task type. The value can be:
     * + **list**: indicates migrating objects using an object list.
     * + **url_list**: indicates migrating objects using a URL object list.
     * + **object**: indicates migrating selected files or folders.
     * + **prefix**: indicates migrating objects with specified prefixes.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MigrationTask resource.
 */
export interface MigrationTaskArgs {
    /**
     * Specifies the traffic limit rules. Each element in the array
     * corresponds to the maximum bandwidth of a time segment. A maximum of 5 time segments are allowed, and the time
     * segments must not overlap. The object structure is  documented below.
     */
    bandwidthPolicies?: pulumi.Input<pulumi.Input<inputs.Oms.MigrationTaskBandwidthPolicy>[]>;
    /**
     * Specifies the description of the task.
     * Changing this creates a new resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the destination information. The object
     * structure is documented below. Changing this creates a new resource.
     */
    destinationObject: pulumi.Input<inputs.Oms.MigrationTaskDestinationObject>;
    /**
     * Specifies whether to record failed objects. If this
     * function is enabled, information about objects that fail to be migrated will be stored in the destination bucket.
     * Default value: **true**. Changing this creates a new resource.
     */
    enableFailedObjectRecording?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the KMS encryption function.
     * Default value: **false**. Changing this creates a new resource.
     */
    enableKms?: pulumi.Input<boolean>;
    /**
     * Specifies whether to automatically restore the archive data. If enabled,
     * archive data is automatically restored and migrated. Default value: **false**. Changing this creates a new resource.
     */
    enableRestore?: pulumi.Input<boolean>;
    /**
     * Specifies a time in format **yyyy-MM-dd HH:mm:ss**,
     * e.g. **2006-01-02 15:04:05**. The system migrates only the objects that are modified after the specified time.
     * No time is specified by default. Changing this creates a new resource.
     */
    migrateSince?: pulumi.Input<string>;
    /**
     * Specifies the region where the destination bucket is located.
     * Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the SMN message sending configuration.
     * The object structure is  documented below. Changing this creates a new resource.
     */
    smnConfig?: pulumi.Input<inputs.Oms.MigrationTaskSmnConfig>;
    /**
     * Specifies the CDN information. If this parameter is contained,
     * using CDN to download source data is supported, the source objects to be migrated are obtained from the CDN domain
     * name during migration. The object structure is documented below.
     * Changing this creates a new resource.
     */
    sourceCdn?: pulumi.Input<inputs.Oms.MigrationTaskSourceCdn>;
    /**
     * Specifies the source information. The object
     * structure is documented below. Changing this creates a new resource.
     */
    sourceObject: pulumi.Input<inputs.Oms.MigrationTaskSourceObject>;
    /**
     * Specifies whether to start the task. Default value: **true**.
     */
    startTask?: pulumi.Input<boolean>;
    /**
     * Specifies the task type. The value can be:
     * + **list**: indicates migrating objects using an object list.
     * + **url_list**: indicates migrating objects using a URL object list.
     * + **object**: indicates migrating selected files or folders.
     * + **prefix**: indicates migrating objects with specified prefixes.
     */
    type: pulumi.Input<string>;
}
