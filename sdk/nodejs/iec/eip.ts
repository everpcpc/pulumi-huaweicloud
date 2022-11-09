// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a eip resource within HuaweiCloud IEC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const iecSites = huaweicloud.Iec.getSites({});
 * const eipTest = new huaweicloud.iec.Eip("eipTest", {siteId: iecSites.then(iecSites => iecSites.sites?[0]?.id)});
 * ```
 *
 * ## Import
 *
 * IEC EIPs can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Iec/eip:Eip eip_test b5ad19d1-57d1-48fd-aab7-1378f9bee169
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
    public static readonly __pulumiType = 'huaweicloud:Iec/eip:Eip';

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

    public /*out*/ readonly bandwidthId!: pulumi.Output<string>;
    public /*out*/ readonly bandwidthName!: pulumi.Output<string>;
    public /*out*/ readonly bandwidthShareType!: pulumi.Output<string>;
    public /*out*/ readonly bandwidthSize!: pulumi.Output<number>;
    /**
     * The version of elastic IP address.
     */
    public readonly ipVersion!: pulumi.Output<number>;
    /**
     * Specifies the line ID of IEC sevice site.
     * Changing this parameter creates a new resource.
     */
    public readonly lineId!: pulumi.Output<string>;
    /**
     * Specifies the port ID which this eip will associate with.
     */
    public readonly portId!: pulumi.Output<string>;
    /**
     * The address of private IP.
     */
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    /**
     * The address of elastic IP.
     */
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the ID of IEC sevice site. Changing this parameter creates a new
     * resource.
     */
    public readonly siteId!: pulumi.Output<string>;
    /**
     * The located information of the IEC site. It contains area, province and city.
     */
    public /*out*/ readonly siteInfo!: pulumi.Output<string>;
    /**
     * The status of elastic IP.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

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
            resourceInputs["bandwidthId"] = state ? state.bandwidthId : undefined;
            resourceInputs["bandwidthName"] = state ? state.bandwidthName : undefined;
            resourceInputs["bandwidthShareType"] = state ? state.bandwidthShareType : undefined;
            resourceInputs["bandwidthSize"] = state ? state.bandwidthSize : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["lineId"] = state ? state.lineId : undefined;
            resourceInputs["portId"] = state ? state.portId : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["siteInfo"] = state ? state.siteInfo : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as EipArgs | undefined;
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["lineId"] = args ? args.lineId : undefined;
            resourceInputs["portId"] = args ? args.portId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["bandwidthId"] = undefined /*out*/;
            resourceInputs["bandwidthName"] = undefined /*out*/;
            resourceInputs["bandwidthShareType"] = undefined /*out*/;
            resourceInputs["bandwidthSize"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["siteInfo"] = undefined /*out*/;
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
    bandwidthId?: pulumi.Input<string>;
    bandwidthName?: pulumi.Input<string>;
    bandwidthShareType?: pulumi.Input<string>;
    bandwidthSize?: pulumi.Input<number>;
    /**
     * The version of elastic IP address.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Specifies the line ID of IEC sevice site.
     * Changing this parameter creates a new resource.
     */
    lineId?: pulumi.Input<string>;
    /**
     * Specifies the port ID which this eip will associate with.
     */
    portId?: pulumi.Input<string>;
    /**
     * The address of private IP.
     */
    privateIp?: pulumi.Input<string>;
    /**
     * The address of elastic IP.
     */
    publicIp?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * Specifies the ID of IEC sevice site. Changing this parameter creates a new
     * resource.
     */
    siteId?: pulumi.Input<string>;
    /**
     * The located information of the IEC site. It contains area, province and city.
     */
    siteInfo?: pulumi.Input<string>;
    /**
     * The status of elastic IP.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Eip resource.
 */
export interface EipArgs {
    /**
     * The version of elastic IP address.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Specifies the line ID of IEC sevice site.
     * Changing this parameter creates a new resource.
     */
    lineId?: pulumi.Input<string>;
    /**
     * Specifies the port ID which this eip will associate with.
     */
    portId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * Specifies the ID of IEC sevice site. Changing this parameter creates a new
     * resource.
     */
    siteId: pulumi.Input<string>;
}
