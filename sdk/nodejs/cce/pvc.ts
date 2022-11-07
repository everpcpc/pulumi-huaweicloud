// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a CCE Persistent Volume Claim resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create PVC with EVS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.requireObject("clusterId");
 * const namespace = config.requireObject("namespace");
 * const pvcName = config.requireObject("pvcName");
 * const test = new huaweicloud.cce.Pvc("test", {
 *     clusterId: clusterId,
 *     namespace: namespace,
 *     annotations: {
 *         "everest.io/disk-volume-type": "SSD",
 *     },
 *     storageClassName: "csi-disk",
 *     accessModes: ["ReadWriteOnce"],
 *     storage: "10Gi",
 * });
 * ```
 * ### Create PVC with OBS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.requireObject("clusterId");
 * const namespace = config.requireObject("namespace");
 * const pvcName = config.requireObject("pvcName");
 * const test = new huaweicloud.cce.Pvc("test", {
 *     clusterId: clusterId,
 *     namespace: namespace,
 *     annotations: {
 *         "everest.io/obs-volume-type": "STANDARD",
 *         "csi.storage.k8s.io/fstype": "obsfs",
 *     },
 *     storageClassName: "csi-obs",
 *     accessModes: ["ReadWriteMany"],
 *     storage: "1Gi",
 * });
 * ```
 * ### Create PVC with SFS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.requireObject("clusterId");
 * const namespace = config.requireObject("namespace");
 * const pvcName = config.requireObject("pvcName");
 * const test = new huaweicloud.cce.Pvc("test", {
 *     clusterId: clusterId,
 *     namespace: namespace,
 *     storageClassName: "csi-nas",
 *     accessModes: ["ReadWriteMany"],
 *     storage: "10Gi",
 * });
 * ```
 *
 * ## Import
 *
 * CCE PVC can be imported using the cluster ID, namespace and name separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cce/pvc:Pvc test 5c20fdad-7288-11eb-b817-0255ac10158b/default/pvc_name
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`annotations`. It is generally recommended running `terraform plan` after importing a PVC. You can then decide if changes should be applied to the PVC, or the resource definition should be updated to align with the PVC. Also you can ignore changes as below. resource "huaweicloud_cce_pvc" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  annotations,
 *
 *  ]
 *
 *  } }
 */
export class Pvc extends pulumi.CustomResource {
    /**
     * Get an existing Pvc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PvcState, opts?: pulumi.CustomResourceOptions): Pvc {
        return new Pvc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cce/pvc:Pvc';

    /**
     * Returns true if the given object is an instance of Pvc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pvc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pvc.__pulumiType;
    }

    /**
     * Specifies the desired access modes the volume should have.
     * The valid values are as follows:
     * + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
     * + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
     * + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
     */
    public readonly accessModes!: pulumi.Output<string[]>;
    /**
     * Specifies the unstructured key value map for external parameters.
     * Changing this will create a new PVC resource.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the cluster ID to which the CCE PVC belongs.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The server time when PVC was created.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Specifies the map of string keys and values for labels.
     * Changing this will create a new PVC resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the unique name of the PVC resource. This parameter can contain a
     * maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
     * lowercase letters and digits. Changing this will create a new PVC resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the namespace to logically divide your containers into different
     * group. Changing this will create a new PVC resource.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the PVC resource.
     * If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The current phase of the PVC.
     * + **Pending**: Not yet bound.
     * + **Bound**: Already bound.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the minimum amount of storage resources required.
     * Changing this creates a new PVC resource.
     */
    public readonly storage!: pulumi.Output<string>;
    /**
     * Specifies the type of the storage bound to the CCE PVC.
     * The valid values are as follows:
     * + **csi-disk**: EVS.
     * + **csi-obs**: OBS.
     * + **csi-nas**: SFS.
     */
    public readonly storageClassName!: pulumi.Output<string>;

    /**
     * Create a Pvc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PvcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PvcArgs | PvcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PvcState | undefined;
            resourceInputs["accessModes"] = state ? state.accessModes : undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["storageClassName"] = state ? state.storageClassName : undefined;
        } else {
            const args = argsOrState as PvcArgs | undefined;
            if ((!args || args.accessModes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessModes'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            if ((!args || args.storageClassName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageClassName'");
            }
            resourceInputs["accessModes"] = args ? args.accessModes : undefined;
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["storageClassName"] = args ? args.storageClassName : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pvc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pvc resources.
 */
export interface PvcState {
    /**
     * Specifies the desired access modes the volume should have.
     * The valid values are as follows:
     * + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
     * + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
     * + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
     */
    accessModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the unstructured key value map for external parameters.
     * Changing this will create a new PVC resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the cluster ID to which the CCE PVC belongs.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The server time when PVC was created.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * Specifies the map of string keys and values for labels.
     * Changing this will create a new PVC resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the unique name of the PVC resource. This parameter can contain a
     * maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
     * lowercase letters and digits. Changing this will create a new PVC resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the namespace to logically divide your containers into different
     * group. Changing this will create a new PVC resource.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the PVC resource.
     * If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The current phase of the PVC.
     * + **Pending**: Not yet bound.
     * + **Bound**: Already bound.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the minimum amount of storage resources required.
     * Changing this creates a new PVC resource.
     */
    storage?: pulumi.Input<string>;
    /**
     * Specifies the type of the storage bound to the CCE PVC.
     * The valid values are as follows:
     * + **csi-disk**: EVS.
     * + **csi-obs**: OBS.
     * + **csi-nas**: SFS.
     */
    storageClassName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pvc resource.
 */
export interface PvcArgs {
    /**
     * Specifies the desired access modes the volume should have.
     * The valid values are as follows:
     * + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
     * + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
     * + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
     */
    accessModes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the unstructured key value map for external parameters.
     * Changing this will create a new PVC resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the cluster ID to which the CCE PVC belongs.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Specifies the map of string keys and values for labels.
     * Changing this will create a new PVC resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the unique name of the PVC resource. This parameter can contain a
     * maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
     * lowercase letters and digits. Changing this will create a new PVC resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the namespace to logically divide your containers into different
     * group. Changing this will create a new PVC resource.
     */
    namespace: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the PVC resource.
     * If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the minimum amount of storage resources required.
     * Changing this creates a new PVC resource.
     */
    storage: pulumi.Input<string>;
    /**
     * Specifies the type of the storage bound to the CCE PVC.
     * The valid values are as follows:
     * + **csi-disk**: EVS.
     * + **csi-obs**: OBS.
     * + **csi-nas**: SFS.
     */
    storageClassName: pulumi.Input<string>;
}