// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an agency resource within huawei cloud.
 *
 * ## Example Usage
 * ### Delegate another HUAWEI CLOUD account to perform operations on your resources
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const agency = new huaweicloud.Iam.Agency("agency", {
 *     allResourcesRoles: ["Server Administrator"],
 *     delegatedDomainName: "***",
 *     description: "test agency",
 *     domainRoles: ["Anti-DDoS Administrator"],
 *     projectRoles: [{
 *         project: "cn-north-1",
 *         roles: ["Tenant Administrator"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Agencies can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Iam/agency:Agency agency 0b97661f9900f23f4fc2c00971ea4dc0
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to `all_resources_roles` field is missing from the API response. It is generally recommended running `terraform plan` after importing an agency. You can then decide if changes should be applied to the agency, or the resource definition should be updated to align with the agency. Also you can ignore changes as below. hcl resource "huaweicloud_identity_agency" "agency" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [all_resources_roles]
 *
 *  } }
 */
export class Agency extends pulumi.CustomResource {
    /**
     * Get an existing Agency resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AgencyState, opts?: pulumi.CustomResourceOptions): Agency {
        return new Agency(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Iam/agency:Agency';

    /**
     * Returns true if the given object is an instance of Agency.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Agency {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Agency.__pulumiType;
    }

    /**
     * Specifies an array of one or more role names which stand for the permissions
     * to be granted to agency on all resources, including those in enterprise projects, region-specific projects,
     * and global services under your account.
     */
    public readonly allResourcesRoles!: pulumi.Output<string[] | undefined>;
    /**
     * The time when the agency was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Specifies the name of delegated user domain.
     */
    public readonly delegatedDomainName!: pulumi.Output<string | undefined>;
    /**
     * schema: Internal
     */
    public readonly delegatedServiceName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the supplementary information about the agency. The value is a string of
     * 0 to 255 characters, excluding these characters: '**@#$%^&*<>\\**'.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more role names which stand for the permissions to be
     * granted to agency on domain.
     */
    public readonly domainRoles!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the validity period of an agency. The valid value are *FOREVER*, *ONEDAY*
     * or the specific days, for example, "20". The default value is *FOREVER*.
     */
    public readonly duration!: pulumi.Output<string | undefined>;
    /**
     * The expiration time of agency.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Specifies the name of agency. The name is a string of 1 to 64 characters.
     * Changing this will create a new agency.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more roles and projects which are used to grant
     * permissions to agency on project. The structure is documented below.
     */
    public readonly projectRoles!: pulumi.Output<outputs.Iam.AgencyProjectRole[] | undefined>;

    /**
     * Create a Agency resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AgencyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AgencyArgs | AgencyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AgencyState | undefined;
            resourceInputs["allResourcesRoles"] = state ? state.allResourcesRoles : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["delegatedDomainName"] = state ? state.delegatedDomainName : undefined;
            resourceInputs["delegatedServiceName"] = state ? state.delegatedServiceName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainRoles"] = state ? state.domainRoles : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectRoles"] = state ? state.projectRoles : undefined;
        } else {
            const args = argsOrState as AgencyArgs | undefined;
            resourceInputs["allResourcesRoles"] = args ? args.allResourcesRoles : undefined;
            resourceInputs["delegatedDomainName"] = args ? args.delegatedDomainName : undefined;
            resourceInputs["delegatedServiceName"] = args ? args.delegatedServiceName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domainRoles"] = args ? args.domainRoles : undefined;
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectRoles"] = args ? args.projectRoles : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Agency.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Agency resources.
 */
export interface AgencyState {
    /**
     * Specifies an array of one or more role names which stand for the permissions
     * to be granted to agency on all resources, including those in enterprise projects, region-specific projects,
     * and global services under your account.
     */
    allResourcesRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time when the agency was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Specifies the name of delegated user domain.
     */
    delegatedDomainName?: pulumi.Input<string>;
    /**
     * schema: Internal
     */
    delegatedServiceName?: pulumi.Input<string>;
    /**
     * Specifies the supplementary information about the agency. The value is a string of
     * 0 to 255 characters, excluding these characters: '**@#$%^&*<>\\**'.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more role names which stand for the permissions to be
     * granted to agency on domain.
     */
    domainRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the validity period of an agency. The valid value are *FOREVER*, *ONEDAY*
     * or the specific days, for example, "20". The default value is *FOREVER*.
     */
    duration?: pulumi.Input<string>;
    /**
     * The expiration time of agency.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Specifies the name of agency. The name is a string of 1 to 64 characters.
     * Changing this will create a new agency.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more roles and projects which are used to grant
     * permissions to agency on project. The structure is documented below.
     */
    projectRoles?: pulumi.Input<pulumi.Input<inputs.Iam.AgencyProjectRole>[]>;
}

/**
 * The set of arguments for constructing a Agency resource.
 */
export interface AgencyArgs {
    /**
     * Specifies an array of one or more role names which stand for the permissions
     * to be granted to agency on all resources, including those in enterprise projects, region-specific projects,
     * and global services under your account.
     */
    allResourcesRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of delegated user domain.
     */
    delegatedDomainName?: pulumi.Input<string>;
    /**
     * schema: Internal
     */
    delegatedServiceName?: pulumi.Input<string>;
    /**
     * Specifies the supplementary information about the agency. The value is a string of
     * 0 to 255 characters, excluding these characters: '**@#$%^&*<>\\**'.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more role names which stand for the permissions to be
     * granted to agency on domain.
     */
    domainRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the validity period of an agency. The valid value are *FOREVER*, *ONEDAY*
     * or the specific days, for example, "20". The default value is *FOREVER*.
     */
    duration?: pulumi.Input<string>;
    /**
     * Specifies the name of agency. The name is a string of 1 to 64 characters.
     * Changing this will create a new agency.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more roles and projects which are used to grant
     * permissions to agency on project. The structure is documented below.
     */
    projectRoles?: pulumi.Input<pulumi.Input<inputs.Iam.AgencyProjectRole>[]>;
}
