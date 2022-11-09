// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an EIP resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### EIP with Dedicated Bandwidth
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const eip1 = new huaweicloud.Vpc.Eip("eip_1", {
 *     bandwidth: {
 *         chargeMode: "traffic",
 *         name: "test",
 *         shareType: "PER",
 *         size: 10,
 *     },
 *     publicip: {
 *         type: "5_bgp",
 *     },
 * });
 * ```
 * ### EIP with Shared Bandwidth
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const bandwidth1 = new huaweicloud.vpc.Bandwidth("bandwidth1", {size: 5});
 * const eip1 = new huaweicloud.vpc.Eip("eip1", {
 *     publicip: {
 *         type: "5_bgp",
 *     },
 *     bandwidth: {
 *         shareType: "WHOLE",
 *         id: bandwidth1.id,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * EIPs can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Vpc/eip:Eip eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1
 * ```
 */
export class Eip extends pulumi.CustomResource {
    /**
     * Get an existing Eip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EipState, opts?: pulumi.CustomResourceOptions): Eip {
        return new Eip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Vpc/eip:Eip';

    /**
     * Returns true if the given object is an instance of Eip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Eip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Eip.__pulumiType;
    }

    /**
     * The IPv4 address of the EIP.
     */
    public /*out*/ readonly address!: pulumi.Output<string>;
    public readonly autoPay!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether auto renew is enabled.
     * Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * The bandwidth object.
     */
    public readonly bandwidth!: pulumi.Output<outputs.Vpc.EipBandwidth>;
    /**
     * Specifies the charging mode of the elastic IP. Valid values are
     * *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
     */
    public readonly chargingMode!: pulumi.Output<string>;
    /**
     * The enterprise project id of the elastic IP. Changing this
     * creates a new eip.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The IPv6 address of the EIP.
     */
    public /*out*/ readonly ipv6Address!: pulumi.Output<string>;
    /**
     * The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
     * underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the charging period of the elastic IP. If `periodUnit` is set to
     * *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
     * is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Specifies the charging period unit of the elastic IP. Valid values are
     * *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
     * eip.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The port ID which the EIP associated with.
     */
    public /*out*/ readonly portId!: pulumi.Output<string>;
    /**
     * The private IP address bound to the EIP.
     */
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    /**
     * The elastic IP address object.
     */
    public readonly publicip!: pulumi.Output<outputs.Vpc.EipPublicip>;
    /**
     * The region in which to create the EIP resource. If omitted, the provider-level
     * region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The status of EIP.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the elastic IP.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Eip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EipArgs | EipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EipState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["autoPay"] = state ? state.autoPay : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["portId"] = state ? state.portId : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["publicip"] = state ? state.publicip : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EipArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.publicip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicip'");
            }
            resourceInputs["autoPay"] = args ? args.autoPay : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["publicip"] = args ? args.publicip : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["address"] = undefined /*out*/;
            resourceInputs["ipv6Address"] = undefined /*out*/;
            resourceInputs["portId"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Eip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Eip resources.
 */
export interface EipState {
    /**
     * The IPv4 address of the EIP.
     */
    address?: pulumi.Input<string>;
    autoPay?: pulumi.Input<string>;
    /**
     * Specifies whether auto renew is enabled.
     * Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * The bandwidth object.
     */
    bandwidth?: pulumi.Input<inputs.Vpc.EipBandwidth>;
    /**
     * Specifies the charging mode of the elastic IP. Valid values are
     * *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * The enterprise project id of the elastic IP. Changing this
     * creates a new eip.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The IPv6 address of the EIP.
     */
    ipv6Address?: pulumi.Input<string>;
    /**
     * The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
     * underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the elastic IP. If `periodUnit` is set to
     * *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
     * is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the elastic IP. Valid values are
     * *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
     * eip.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The port ID which the EIP associated with.
     */
    portId?: pulumi.Input<string>;
    /**
     * The private IP address bound to the EIP.
     */
    privateIp?: pulumi.Input<string>;
    /**
     * The elastic IP address object.
     */
    publicip?: pulumi.Input<inputs.Vpc.EipPublicip>;
    /**
     * The region in which to create the EIP resource. If omitted, the provider-level
     * region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of EIP.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the elastic IP.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Eip resource.
 */
export interface EipArgs {
    autoPay?: pulumi.Input<string>;
    /**
     * Specifies whether auto renew is enabled.
     * Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * The bandwidth object.
     */
    bandwidth: pulumi.Input<inputs.Vpc.EipBandwidth>;
    /**
     * Specifies the charging mode of the elastic IP. Valid values are
     * *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * The enterprise project id of the elastic IP. Changing this
     * creates a new eip.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
     * underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the elastic IP. If `periodUnit` is set to
     * *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
     * is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the elastic IP. Valid values are
     * *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
     * eip.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The elastic IP address object.
     */
    publicip: pulumi.Input<inputs.Vpc.EipPublicip>;
    /**
     * The region in which to create the EIP resource. If omitted, the provider-level
     * region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the elastic IP.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
