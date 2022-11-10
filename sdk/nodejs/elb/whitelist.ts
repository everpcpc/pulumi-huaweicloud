// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an ELB whitelist resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const listener1 = new huaweicloud.elb.Listener("listener1", {
 *     protocol: "HTTP",
 *     protocolPort: 8080,
 *     loadbalancerId: _var.loadbalancer_id,
 * });
 * const whitelist1 = new huaweicloud.elb.Whitelist("whitelist1", {
 *     enableWhitelist: true,
 *     whitelist: "192.168.11.1,192.168.0.1/24,192.168.201.18/8",
 *     listenerId: listener1.id,
 * });
 * ```
 *
 * ## Import
 *
 * ELB whitelist can be imported using the whitelist ID, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Elb/whitelist:Whitelist whitelist_1 5c20fdad-7288-11eb-b817-0255ac10158b
 * ```
 */
export class Whitelist extends pulumi.CustomResource {
    /**
     * Get an existing Whitelist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WhitelistState, opts?: pulumi.CustomResourceOptions): Whitelist {
        return new Whitelist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Elb/whitelist:Whitelist';

    /**
     * Returns true if the given object is an instance of Whitelist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Whitelist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Whitelist.__pulumiType;
    }

    /**
     * Specify whether to enable access control.
     */
    public readonly enableWhitelist!: pulumi.Output<boolean | undefined>;
    /**
     * The Listener ID that the whitelist will be associated with. Changing this
     * creates a new whitelist.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * The region in which to create the ELB whitelist resource. If omitted, the
     * provider-level region will be used. Changing this creates a new whitelist.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Specifies the IP addresses in the whitelist. Use commas(,) to separate the multiple
     * IP addresses.
     */
    public readonly whitelist!: pulumi.Output<string | undefined>;

    /**
     * Create a Whitelist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WhitelistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WhitelistArgs | WhitelistState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WhitelistState | undefined;
            resourceInputs["enableWhitelist"] = state ? state.enableWhitelist : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["whitelist"] = state ? state.whitelist : undefined;
        } else {
            const args = argsOrState as WhitelistArgs | undefined;
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            resourceInputs["enableWhitelist"] = args ? args.enableWhitelist : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["whitelist"] = args ? args.whitelist : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Whitelist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Whitelist resources.
 */
export interface WhitelistState {
    /**
     * Specify whether to enable access control.
     */
    enableWhitelist?: pulumi.Input<boolean>;
    /**
     * The Listener ID that the whitelist will be associated with. Changing this
     * creates a new whitelist.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The region in which to create the ELB whitelist resource. If omitted, the
     * provider-level region will be used. Changing this creates a new whitelist.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Specifies the IP addresses in the whitelist. Use commas(,) to separate the multiple
     * IP addresses.
     */
    whitelist?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Whitelist resource.
 */
export interface WhitelistArgs {
    /**
     * Specify whether to enable access control.
     */
    enableWhitelist?: pulumi.Input<boolean>;
    /**
     * The Listener ID that the whitelist will be associated with. Changing this
     * creates a new whitelist.
     */
    listenerId: pulumi.Input<string>;
    /**
     * The region in which to create the ELB whitelist resource. If omitted, the
     * provider-level region will be used. Changing this creates a new whitelist.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Specifies the IP addresses in the whitelist. Use commas(,) to separate the multiple
     * IP addresses.
     */
    whitelist?: pulumi.Input<string>;
}
