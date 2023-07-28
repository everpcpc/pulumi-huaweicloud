// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Image resource within HuaweiCloud IMS.
 *
 * ## Example Usage
 * ### Creating an image from OBS bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const imsTestFile = new huaweicloud.Ims.Image("ims_test_file", {
 *     description: "Create an image from the OBS bucket.",
 *     imageUrl: "ims-image:centos70.qcow2",
 *     minDisk: 40,
 *     tags: {
 *         foo: "bar1",
 *         key: "value",
 *     },
 * });
 * ```
 * ### Creating a whole image from an existing ECS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vaultId = config.requireObject("vaultId");
 * const instanceId = config.requireObject("instanceId");
 * const test = new huaweicloud.ims.Image("test", {
 *     instanceId: instanceId,
 *     vaultId: vaultId,
 *     tags: {
 *         foo: "bar2",
 *         key: "value",
 *     },
 * });
 * ```
 * ### Creating a whole image from CBR backup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const backupId = config.requireObject("backupId");
 * const test = new huaweicloud.ims.Image("test", {
 *     backupId: backupId,
 *     tags: {
 *         foo: "bar1",
 *         key: "value",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Images can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Ims/image:Image my_image <id>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response. The missing attributes include`vault_id`. It is generally recommended running `terraform plan` after importing the image. You can then decide if changes should be applied to the image, or the resource definition should be updated to align with the image. Also you can ignore changes as below. resource "huaweicloud_images_image" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  vault_id,
 *
 *  ]
 *
 *  } }
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Ims/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * The ID of the CBR backup that needs to be converted into an image. This
     * parameter is mandatory when you create a private whole image from a CBR backup.
     */
    public readonly backupId!: pulumi.Output<string>;
    /**
     * The checksum of the data associated with the image.
     */
    public /*out*/ readonly checksum!: pulumi.Output<string>;
    /**
     * The master key used for encrypting an image.
     */
    public readonly cmkId!: pulumi.Output<string | undefined>;
    /**
     * The image resource. The pattern can be 'instance,*instance_id*', 'file,*image_url*'
     * or 'server_backup,*backup_id*'.
     */
    public /*out*/ readonly dataOrigin!: pulumi.Output<string>;
    /**
     * A description of the image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The image file format. The value can be `vhd`, `zvhd`, `raw`, `zvhd2`, or `qcow2`.
     */
    public /*out*/ readonly diskFormat!: pulumi.Output<string>;
    /**
     * The enterprise project id of the image. Changing this creates a
     * new image.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The size(bytes) of the image file format.
     */
    public /*out*/ readonly imageSize!: pulumi.Output<string>;
    /**
     * The URL of the external image file in the OBS bucket. This parameter is
     * mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
     * name:Image file name*.
     */
    public readonly imageUrl!: pulumi.Output<string | undefined>;
    /**
     * The ID of the ECS that needs to be converted into an image. This
     * parameter is mandatory when you create a private image or a private whole image from an ECS.
     * If the value of `vaultId` is not empty, then a whole image will be created.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * If automatic configuration is required, set the value to true. Otherwise, set
     * the value to false.
     */
    public readonly isConfig!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum memory of the image in the unit of MB.
     */
    public readonly maxRam!: pulumi.Output<number>;
    /**
     * The minimum size of the system disk in the unit of GB. This parameter is
     * mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
     * to 1024 GB.
     */
    public readonly minDisk!: pulumi.Output<number | undefined>;
    /**
     * The minimum memory of the image in the unit of MB. The default value is 0,
     * indicating that the memory is not restricted.
     */
    public readonly minRam!: pulumi.Output<number>;
    /**
     * The name of the image.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The OS version. This parameter is valid when you create a private image
     * from an external file uploaded to an OBS bucket.
     */
    public readonly osVersion!: pulumi.Output<string>;
    /**
     * The status of the image.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags of the image.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The ID of the vault to which an ECS is to be added or has been added.
     * This parameter is mandatory when you create a private whole image from an ECS.
     */
    public readonly vaultId!: pulumi.Output<string | undefined>;
    /**
     * Whether the image is visible to other tenants.
     */
    public /*out*/ readonly visibility!: pulumi.Output<string>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs | ImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageState | undefined;
            resourceInputs["backupId"] = state ? state.backupId : undefined;
            resourceInputs["checksum"] = state ? state.checksum : undefined;
            resourceInputs["cmkId"] = state ? state.cmkId : undefined;
            resourceInputs["dataOrigin"] = state ? state.dataOrigin : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["diskFormat"] = state ? state.diskFormat : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["imageSize"] = state ? state.imageSize : undefined;
            resourceInputs["imageUrl"] = state ? state.imageUrl : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isConfig"] = state ? state.isConfig : undefined;
            resourceInputs["maxRam"] = state ? state.maxRam : undefined;
            resourceInputs["minDisk"] = state ? state.minDisk : undefined;
            resourceInputs["minRam"] = state ? state.minRam : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["osVersion"] = state ? state.osVersion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vaultId"] = state ? state.vaultId : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ImageArgs | undefined;
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["cmkId"] = args ? args.cmkId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["imageUrl"] = args ? args.imageUrl : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isConfig"] = args ? args.isConfig : undefined;
            resourceInputs["maxRam"] = args ? args.maxRam : undefined;
            resourceInputs["minDisk"] = args ? args.minDisk : undefined;
            resourceInputs["minRam"] = args ? args.minRam : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["osVersion"] = args ? args.osVersion : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vaultId"] = args ? args.vaultId : undefined;
            resourceInputs["checksum"] = undefined /*out*/;
            resourceInputs["dataOrigin"] = undefined /*out*/;
            resourceInputs["diskFormat"] = undefined /*out*/;
            resourceInputs["imageSize"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["visibility"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Image resources.
 */
export interface ImageState {
    /**
     * The ID of the CBR backup that needs to be converted into an image. This
     * parameter is mandatory when you create a private whole image from a CBR backup.
     */
    backupId?: pulumi.Input<string>;
    /**
     * The checksum of the data associated with the image.
     */
    checksum?: pulumi.Input<string>;
    /**
     * The master key used for encrypting an image.
     */
    cmkId?: pulumi.Input<string>;
    /**
     * The image resource. The pattern can be 'instance,*instance_id*', 'file,*image_url*'
     * or 'server_backup,*backup_id*'.
     */
    dataOrigin?: pulumi.Input<string>;
    /**
     * A description of the image.
     */
    description?: pulumi.Input<string>;
    /**
     * The image file format. The value can be `vhd`, `zvhd`, `raw`, `zvhd2`, or `qcow2`.
     */
    diskFormat?: pulumi.Input<string>;
    /**
     * The enterprise project id of the image. Changing this creates a
     * new image.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The size(bytes) of the image file format.
     */
    imageSize?: pulumi.Input<string>;
    /**
     * The URL of the external image file in the OBS bucket. This parameter is
     * mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
     * name:Image file name*.
     */
    imageUrl?: pulumi.Input<string>;
    /**
     * The ID of the ECS that needs to be converted into an image. This
     * parameter is mandatory when you create a private image or a private whole image from an ECS.
     * If the value of `vaultId` is not empty, then a whole image will be created.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * If automatic configuration is required, set the value to true. Otherwise, set
     * the value to false.
     */
    isConfig?: pulumi.Input<boolean>;
    /**
     * The maximum memory of the image in the unit of MB.
     */
    maxRam?: pulumi.Input<number>;
    /**
     * The minimum size of the system disk in the unit of GB. This parameter is
     * mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
     * to 1024 GB.
     */
    minDisk?: pulumi.Input<number>;
    /**
     * The minimum memory of the image in the unit of MB. The default value is 0,
     * indicating that the memory is not restricted.
     */
    minRam?: pulumi.Input<number>;
    /**
     * The name of the image.
     */
    name?: pulumi.Input<string>;
    /**
     * The OS version. This parameter is valid when you create a private image
     * from an external file uploaded to an OBS bucket.
     */
    osVersion?: pulumi.Input<string>;
    /**
     * The status of the image.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags of the image.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
     */
    type?: pulumi.Input<string>;
    /**
     * The ID of the vault to which an ECS is to be added or has been added.
     * This parameter is mandatory when you create a private whole image from an ECS.
     */
    vaultId?: pulumi.Input<string>;
    /**
     * Whether the image is visible to other tenants.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * The ID of the CBR backup that needs to be converted into an image. This
     * parameter is mandatory when you create a private whole image from a CBR backup.
     */
    backupId?: pulumi.Input<string>;
    /**
     * The master key used for encrypting an image.
     */
    cmkId?: pulumi.Input<string>;
    /**
     * A description of the image.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the image. Changing this creates a
     * new image.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The URL of the external image file in the OBS bucket. This parameter is
     * mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
     * name:Image file name*.
     */
    imageUrl?: pulumi.Input<string>;
    /**
     * The ID of the ECS that needs to be converted into an image. This
     * parameter is mandatory when you create a private image or a private whole image from an ECS.
     * If the value of `vaultId` is not empty, then a whole image will be created.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * If automatic configuration is required, set the value to true. Otherwise, set
     * the value to false.
     */
    isConfig?: pulumi.Input<boolean>;
    /**
     * The maximum memory of the image in the unit of MB.
     */
    maxRam?: pulumi.Input<number>;
    /**
     * The minimum size of the system disk in the unit of GB. This parameter is
     * mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
     * to 1024 GB.
     */
    minDisk?: pulumi.Input<number>;
    /**
     * The minimum memory of the image in the unit of MB. The default value is 0,
     * indicating that the memory is not restricted.
     */
    minRam?: pulumi.Input<number>;
    /**
     * The name of the image.
     */
    name?: pulumi.Input<string>;
    /**
     * The OS version. This parameter is valid when you create a private image
     * from an external file uploaded to an OBS bucket.
     */
    osVersion?: pulumi.Input<string>;
    /**
     * The tags of the image.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
     */
    type?: pulumi.Input<string>;
    /**
     * The ID of the vault to which an ECS is to be added or has been added.
     * This parameter is mandatory when you create a private whole image from an ECS.
     */
    vaultId?: pulumi.Input<string>;
}
