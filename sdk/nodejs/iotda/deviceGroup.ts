// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IoTDA device group within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const spaceId = config.requireObject("spaceId");
 * const deviceId = config.requireObject("deviceId");
 * const group = new huaweicloud.iotda.DeviceGroup("group", {
 *     spaceId: spaceId,
 *     deviceIds: [deviceId],
 * });
 * ```
 *
 * ## Import
 *
 * Groups can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:IoTDA/deviceGroup:DeviceGroup test 10022532f4f94f26b01daa1e424853e1
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`space_id`. It is generally recommended running `terraform plan` after importing the resource. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the group. Also you can ignore changes as below. resource "huaweicloud_iotda_device_group" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  space_id,
 *
 *  ]
 *
 *  } }
 */
export class DeviceGroup extends pulumi.CustomResource {
    /**
     * Get an existing DeviceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceGroupState, opts?: pulumi.CustomResourceOptions): DeviceGroup {
        return new DeviceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:IoTDA/deviceGroup:DeviceGroup';

    /**
     * Returns true if the given object is an instance of DeviceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceGroup.__pulumiType;
    }

    /**
     * Specifies the description of device group. The description contains a maximum of 64
     * characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters
     * are allowed: `?'#().,&%@!`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the list of device IDs bound to the group.
     */
    public readonly deviceIds!: pulumi.Output<string[]>;
    /**
     * Specifies the name of device group. The name contains a maximum of 64 characters.
     * Only letters, digits, hyphens (-) and underscores (_) are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the parent group id.
     * Changing this parameter will create a new resource.
     */
    public readonly parentGroupId!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the IoTDA device group resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the resource space ID to which the device group belongs.
     * Changing this parameter will create a new resource.
     */
    public readonly spaceId!: pulumi.Output<string>;

    /**
     * Create a DeviceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceGroupArgs | DeviceGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceIds"] = state ? state.deviceIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentGroupId"] = state ? state.parentGroupId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
        } else {
            const args = argsOrState as DeviceGroupArgs | undefined;
            if ((!args || args.spaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spaceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceIds"] = args ? args.deviceIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentGroupId"] = args ? args.parentGroupId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceGroup resources.
 */
export interface DeviceGroupState {
    /**
     * Specifies the description of device group. The description contains a maximum of 64
     * characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters
     * are allowed: `?'#().,&%@!`.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the list of device IDs bound to the group.
     */
    deviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of device group. The name contains a maximum of 64 characters.
     * Only letters, digits, hyphens (-) and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the parent group id.
     * Changing this parameter will create a new resource.
     */
    parentGroupId?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the IoTDA device group resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the resource space ID to which the device group belongs.
     * Changing this parameter will create a new resource.
     */
    spaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeviceGroup resource.
 */
export interface DeviceGroupArgs {
    /**
     * Specifies the description of device group. The description contains a maximum of 64
     * characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters
     * are allowed: `?'#().,&%@!`.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the list of device IDs bound to the group.
     */
    deviceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of device group. The name contains a maximum of 64 characters.
     * Only letters, digits, hyphens (-) and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the parent group id.
     * Changing this parameter will create a new resource.
     */
    parentGroupId?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the IoTDA device group resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the resource space ID to which the device group belongs.
     * Changing this parameter will create a new resource.
     */
    spaceId: pulumi.Input<string>;
}