// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a CCE add-on resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.requireObject("clusterId");
 * const addonTest = new huaweicloud.cce.Addon("addonTest", {
 *     clusterId: clusterId,
 *     templateName: "metrics-server",
 *     version: "1.0.0",
 * });
 * ```
 *
 * ## Import
 *
 * CCE add-on can be imported using the cluster ID and add-on ID separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cce/addon:Addon my_addon bb6923e4-b16e-11eb-b0cd-0255ac101da1/c7ecb230-b16f-11eb-b3b6-0255ac1015a3
 * ```
 */
export class Addon extends pulumi.CustomResource {
    /**
     * Get an existing Addon resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddonState, opts?: pulumi.CustomResourceOptions): Addon {
        return new Addon(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cce/addon:Addon';

    /**
     * Returns true if the given object is an instance of Addon.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Addon {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Addon.__pulumiType;
    }

    /**
     * Specifies the cluster ID.
     * Changing this parameter will create a new resource.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Description of add-on instance.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the CCE add-on resource.
     * If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Add-on status information.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the name of the add-on template.
     * Changing this parameter will create a new resource.
     */
    public readonly templateName!: pulumi.Output<string>;
    /**
     * Specifies the add-on template installation parameters.
     * These parameters vary depending on the add-on. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    public readonly values!: pulumi.Output<outputs.Cce.AddonValues | undefined>;
    /**
     * Specifies the version of the add-on.
     * Changing this parameter will create a new resource.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Addon resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AddonArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddonArgs | AddonState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AddonState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
            resourceInputs["values"] = state ? state.values : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AddonArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["values"] = args ? args.values : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Addon.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Addon resources.
 */
export interface AddonState {
    /**
     * Specifies the cluster ID.
     * Changing this parameter will create a new resource.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Description of add-on instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the CCE add-on resource.
     * If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Add-on status information.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the name of the add-on template.
     * Changing this parameter will create a new resource.
     */
    templateName?: pulumi.Input<string>;
    /**
     * Specifies the add-on template installation parameters.
     * These parameters vary depending on the add-on. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    values?: pulumi.Input<inputs.Cce.AddonValues>;
    /**
     * Specifies the version of the add-on.
     * Changing this parameter will create a new resource.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Addon resource.
 */
export interface AddonArgs {
    /**
     * Specifies the cluster ID.
     * Changing this parameter will create a new resource.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the CCE add-on resource.
     * If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the name of the add-on template.
     * Changing this parameter will create a new resource.
     */
    templateName: pulumi.Input<string>;
    /**
     * Specifies the add-on template installation parameters.
     * These parameters vary depending on the add-on. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    values?: pulumi.Input<inputs.Cce.AddonValues>;
    /**
     * Specifies the version of the add-on.
     * Changing this parameter will create a new resource.
     */
    version: pulumi.Input<string>;
}