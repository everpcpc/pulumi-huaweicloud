// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a network VIP resource within HuaweiCloud VPC.
 *
 * ## Import
 *
 * Network VIPs can be imported using their `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Vpc/vip:Vip test ce595799-da26-4015-8db5-7733c6db292e
 * ```
 */
export class Vip extends pulumi.CustomResource {
    /**
     * Get an existing Vip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VipState, opts?: pulumi.CustomResourceOptions): Vip {
        return new Vip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Vpc/vip:Vip';

    /**
     * Returns true if the given object is an instance of Vip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vip.__pulumiType;
    }

    /**
     * The device owner of the VIP.
     */
    public /*out*/ readonly deviceOwner!: pulumi.Output<string>;
    /**
     * Specifies the IP address desired in the subnet for this VIP.
     * Changing this will create a new VIP resource.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * Specifies the IP version, either `4` (default) or `6`.
     * Changing this will create a new VIP resource.
     */
    public readonly ipVersion!: pulumi.Output<number>;
    /**
     * The MAC address of the VIP.
     */
    public /*out*/ readonly macAddress!: pulumi.Output<string>;
    /**
     * Specifies a unique name for the VIP.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the network ID of the VPC subnet to which the VIP belongs.
     * Changing this will create a new VIP resource.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the VIP.
     * If omitted, the provider-level region will be used. Changing this will create a new VIP resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The VIP status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The subnet ID in which to allocate IP address for this VIP.
     *
     * @deprecated use ip_version instead
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a Vip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VipArgs | VipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VipState | undefined;
            resourceInputs["deviceOwner"] = state ? state.deviceOwner : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["macAddress"] = state ? state.macAddress : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as VipArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["deviceOwner"] = undefined /*out*/;
            resourceInputs["macAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vip resources.
 */
export interface VipState {
    /**
     * The device owner of the VIP.
     */
    deviceOwner?: pulumi.Input<string>;
    /**
     * Specifies the IP address desired in the subnet for this VIP.
     * Changing this will create a new VIP resource.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Specifies the IP version, either `4` (default) or `6`.
     * Changing this will create a new VIP resource.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * The MAC address of the VIP.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * Specifies a unique name for the VIP.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the network ID of the VPC subnet to which the VIP belongs.
     * Changing this will create a new VIP resource.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the VIP.
     * If omitted, the provider-level region will be used. Changing this will create a new VIP resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The VIP status.
     */
    status?: pulumi.Input<string>;
    /**
     * The subnet ID in which to allocate IP address for this VIP.
     *
     * @deprecated use ip_version instead
     */
    subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vip resource.
 */
export interface VipArgs {
    /**
     * Specifies the IP address desired in the subnet for this VIP.
     * Changing this will create a new VIP resource.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Specifies the IP version, either `4` (default) or `6`.
     * Changing this will create a new VIP resource.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Specifies a unique name for the VIP.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the network ID of the VPC subnet to which the VIP belongs.
     * Changing this will create a new VIP resource.
     */
    networkId: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the VIP.
     * If omitted, the provider-level region will be used. Changing this will create a new VIP resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The subnet ID in which to allocate IP address for this VIP.
     *
     * @deprecated use ip_version instead
     */
    subnetId?: pulumi.Input<string>;
}