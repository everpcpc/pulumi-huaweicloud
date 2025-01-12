// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IoTDA device within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create a directly connected device and an indirectly connected device
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const spaceId = config.requireObject("spaceId");
 * const productId = config.requireObject("productId");
 * const secret = config.requireObject("secret");
 * const device = new huaweicloud.iotda.Device("device", {
 *     nodeId: "device_SN_1",
 *     spaceId: spaceId,
 *     productId: productId,
 *     secret: secret,
 *     tags: {
 *         foo: "bar",
 *         key: "value",
 *     },
 * });
 * const subDevice = new huaweicloud.iotda.Device("subDevice", {
 *     nodeId: "device_SN_2",
 *     spaceId: spaceId,
 *     productId: productId,
 *     gatewayId: device.id,
 * });
 * ```
 *
 * ## Import
 *
 * Devices can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:IoTDA/device:Device test 10022532f4f94f26b01daa1e424853e1
 * ```
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceState, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:IoTDA/device:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * The authentication type of device. The options are as follows:
     * + **SECRET**: Use a secret for identity authentication.
     * + **CERTIFICATES**: Use an x.509 certificate for identity authentication.
     */
    public /*out*/ readonly authType!: pulumi.Output<string>;
    /**
     * Specifies the description of device. The description contains a maximum of 2048
     * characters. Only letters, Chinese characters, digits, hyphens (-), underscore (_) and the following special characters
     * are allowed: `?'#().,&%@!`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the device ID, which contains 4 to 256 characters.
     * Only letters, digits, hyphens (-) and underscore (_) are allowed. If omitted, the platform will automatically allocate
     * a device ID. Changing this parameter will create a new resource.
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * Specifies a fingerprint of X.509 certificate for identity authentication,
     * which is a 40-digit or 64-digit hexadecimal string. For more detail, please see
     * [Registering a Device Authenticated by an X.509 Certificate](https://support.huaweicloud.com/en-us/usermanual-iothub/iot_01_0055.html).
     */
    public readonly fingerprint!: pulumi.Output<string>;
    /**
     * Specifies whether to freeze the device. Defaults to `false`.
     */
    public readonly frozen!: pulumi.Output<boolean>;
    /**
     * Specifies the gateway ID which is the device ID of the parent device.
     * The child device is not directly connected to the platform. If omitted, it means to create a device directly connected
     * to the platform, the `deviceId` of the device is the same as the `gatewayId`.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * Specifies the device name, which contains 4 to 256 characters. Only letters,
     * Chinese characters, digits, hyphens (-), underscore (_) and the following special characters are allowed: `?'#().,&%@!`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the node ID, which contains 4 to 256 characters.
     * The node ID can be IMEI, MAC address, or serial number. Changing this parameter will create a new resource.
     */
    public readonly nodeId!: pulumi.Output<string>;
    /**
     * The node type of device. The options are as follows:
     * + **GATEWAY**: Directly connected device.
     * + **ENDPOINT**: Indirectly connected device.
     * + **UNKNOWN**: Unknown type.
     */
    public /*out*/ readonly nodeType!: pulumi.Output<string>;
    /**
     * Specifies the product ID which the device belongs to.
     * Changing this parameter will create a new resource.
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the IoTDA device resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies a secret for identity authentication, which contains 8 to 32 characters.
     * Only letters, digits, hyphens (-) and underscore (_) are allowed.
     */
    public readonly secret!: pulumi.Output<string>;
    /**
     * Specifies the resource space ID which the device belongs to.
     * Changing this parameter will create a new resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * The status of device. The valid values are **INACTIVE**, **ONLINE**, **OFFLINE**, **FROZEN**, **ABNORMAL**.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the device.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceArgs | DeviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceState | undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["frozen"] = state ? state.frozen : undefined;
            resourceInputs["gatewayId"] = state ? state.gatewayId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeId"] = state ? state.nodeId : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["productId"] = state ? state.productId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DeviceArgs | undefined;
            if ((!args || args.nodeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeId'");
            }
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            if ((!args || args.spaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spaceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["fingerprint"] = args ? args.fingerprint : undefined;
            resourceInputs["frozen"] = args ? args.frozen : undefined;
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeId"] = args ? args.nodeId : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secret"] = args ? args.secret : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["authType"] = undefined /*out*/;
            resourceInputs["nodeType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Device.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Device resources.
 */
export interface DeviceState {
    /**
     * The authentication type of device. The options are as follows:
     * + **SECRET**: Use a secret for identity authentication.
     * + **CERTIFICATES**: Use an x.509 certificate for identity authentication.
     */
    authType?: pulumi.Input<string>;
    /**
     * Specifies the description of device. The description contains a maximum of 2048
     * characters. Only letters, Chinese characters, digits, hyphens (-), underscore (_) and the following special characters
     * are allowed: `?'#().,&%@!`.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the device ID, which contains 4 to 256 characters.
     * Only letters, digits, hyphens (-) and underscore (_) are allowed. If omitted, the platform will automatically allocate
     * a device ID. Changing this parameter will create a new resource.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * Specifies a fingerprint of X.509 certificate for identity authentication,
     * which is a 40-digit or 64-digit hexadecimal string. For more detail, please see
     * [Registering a Device Authenticated by an X.509 Certificate](https://support.huaweicloud.com/en-us/usermanual-iothub/iot_01_0055.html).
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * Specifies whether to freeze the device. Defaults to `false`.
     */
    frozen?: pulumi.Input<boolean>;
    /**
     * Specifies the gateway ID which is the device ID of the parent device.
     * The child device is not directly connected to the platform. If omitted, it means to create a device directly connected
     * to the platform, the `deviceId` of the device is the same as the `gatewayId`.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * Specifies the device name, which contains 4 to 256 characters. Only letters,
     * Chinese characters, digits, hyphens (-), underscore (_) and the following special characters are allowed: `?'#().,&%@!`.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the node ID, which contains 4 to 256 characters.
     * The node ID can be IMEI, MAC address, or serial number. Changing this parameter will create a new resource.
     */
    nodeId?: pulumi.Input<string>;
    /**
     * The node type of device. The options are as follows:
     * + **GATEWAY**: Directly connected device.
     * + **ENDPOINT**: Indirectly connected device.
     * + **UNKNOWN**: Unknown type.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Specifies the product ID which the device belongs to.
     * Changing this parameter will create a new resource.
     */
    productId?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the IoTDA device resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies a secret for identity authentication, which contains 8 to 32 characters.
     * Only letters, digits, hyphens (-) and underscore (_) are allowed.
     */
    secret?: pulumi.Input<string>;
    /**
     * Specifies the resource space ID which the device belongs to.
     * Changing this parameter will create a new resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * The status of device. The valid values are **INACTIVE**, **ONLINE**, **OFFLINE**, **FROZEN**, **ABNORMAL**.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the device.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * Specifies the description of device. The description contains a maximum of 2048
     * characters. Only letters, Chinese characters, digits, hyphens (-), underscore (_) and the following special characters
     * are allowed: `?'#().,&%@!`.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the device ID, which contains 4 to 256 characters.
     * Only letters, digits, hyphens (-) and underscore (_) are allowed. If omitted, the platform will automatically allocate
     * a device ID. Changing this parameter will create a new resource.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * Specifies a fingerprint of X.509 certificate for identity authentication,
     * which is a 40-digit or 64-digit hexadecimal string. For more detail, please see
     * [Registering a Device Authenticated by an X.509 Certificate](https://support.huaweicloud.com/en-us/usermanual-iothub/iot_01_0055.html).
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * Specifies whether to freeze the device. Defaults to `false`.
     */
    frozen?: pulumi.Input<boolean>;
    /**
     * Specifies the gateway ID which is the device ID of the parent device.
     * The child device is not directly connected to the platform. If omitted, it means to create a device directly connected
     * to the platform, the `deviceId` of the device is the same as the `gatewayId`.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * Specifies the device name, which contains 4 to 256 characters. Only letters,
     * Chinese characters, digits, hyphens (-), underscore (_) and the following special characters are allowed: `?'#().,&%@!`.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the node ID, which contains 4 to 256 characters.
     * The node ID can be IMEI, MAC address, or serial number. Changing this parameter will create a new resource.
     */
    nodeId: pulumi.Input<string>;
    /**
     * Specifies the product ID which the device belongs to.
     * Changing this parameter will create a new resource.
     */
    productId: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the IoTDA device resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies a secret for identity authentication, which contains 8 to 32 characters.
     * Only letters, digits, hyphens (-) and underscore (_) are allowed.
     */
    secret?: pulumi.Input<string>;
    /**
     * Specifies the resource space ID which the device belongs to.
     * Changing this parameter will create a new resource.
     */
    spaceId: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the device.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
