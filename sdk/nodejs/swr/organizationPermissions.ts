// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages user permissions for the SWR organization resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const organizationName = config.requireObject("organizationName");
 * const user1 = config.requireObject("user1");
 * const user2 = config.requireObject("user2");
 * const test = new huaweicloud.swr.OrganizationPermissions("test", {
 *     organization: organizationName,
 *     users: [
 *         {
 *             userName: user1.name,
 *             userId: user1.id,
 *             permission: "Read",
 *         },
 *         {
 *             userName: user2.name,
 *             userId: user2.id,
 *             permission: "Read",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Organization Permissions can be imported using the `id` (organization name), e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Swr/organizationPermissions:OrganizationPermissions test terraform-test
 * ```
 */
export class OrganizationPermissions extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationPermissionsState, opts?: pulumi.CustomResourceOptions): OrganizationPermissions {
        return new OrganizationPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Swr/organizationPermissions:OrganizationPermissions';

    /**
     * Returns true if the given object is an instance of OrganizationPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationPermissions.__pulumiType;
    }

    /**
     * The creator user name of the organization.
     */
    public /*out*/ readonly creator!: pulumi.Output<string>;
    /**
     * Specifies the name of the organization (namespace) to be accessed.
     * Changing this creates a new resource.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The permission informations of current user.
     */
    public /*out*/ readonly selfPermissions!: pulumi.Output<outputs.Swr.OrganizationPermissionsSelfPermission[]>;
    /**
     * Specifies the users to access to the organization (namespace).
     * Structure is documented below.
     */
    public readonly users!: pulumi.Output<outputs.Swr.OrganizationPermissionsUser[]>;

    /**
     * Create a OrganizationPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationPermissionsArgs | OrganizationPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationPermissionsState | undefined;
            resourceInputs["creator"] = state ? state.creator : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["selfPermissions"] = state ? state.selfPermissions : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as OrganizationPermissionsArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.users === undefined) && !opts.urn) {
                throw new Error("Missing required property 'users'");
            }
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["creator"] = undefined /*out*/;
            resourceInputs["selfPermissions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationPermissions resources.
 */
export interface OrganizationPermissionsState {
    /**
     * The creator user name of the organization.
     */
    creator?: pulumi.Input<string>;
    /**
     * Specifies the name of the organization (namespace) to be accessed.
     * Changing this creates a new resource.
     */
    organization?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The permission informations of current user.
     */
    selfPermissions?: pulumi.Input<pulumi.Input<inputs.Swr.OrganizationPermissionsSelfPermission>[]>;
    /**
     * Specifies the users to access to the organization (namespace).
     * Structure is documented below.
     */
    users?: pulumi.Input<pulumi.Input<inputs.Swr.OrganizationPermissionsUser>[]>;
}

/**
 * The set of arguments for constructing a OrganizationPermissions resource.
 */
export interface OrganizationPermissionsArgs {
    /**
     * Specifies the name of the organization (namespace) to be accessed.
     * Changing this creates a new resource.
     */
    organization: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the users to access to the organization (namespace).
     * Structure is documented below.
     */
    users: pulumi.Input<pulumi.Input<inputs.Swr.OrganizationPermissionsUser>[]>;
}
