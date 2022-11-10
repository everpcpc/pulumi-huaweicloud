// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages CSMS(Cloud Secret Management Service) secrets within HuaweiCloud.
 *
 * ## Example Usage
 * ### Encrypt Plaintext
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test1 = new huaweicloud.Dew.Secret("test1", {
 *     secretText: "this is a password",
 * });
 * ```
 * ### Encrypt JSON Data
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const test2 = new huaweicloud.dew.Secret("test2", {secretText: JSON.stringify({
 *     username: "admin",
 *     password: "123456",
 * })});
 * ```
 *
 * ## Import
 *
 * CSMS secret can be imported using the ID and the name of secret, separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dew/secret:Secret test 93cba7f5-550b-45dc-912e-277b3296fb27/test_secret
 * ```
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Dew/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Secret.__pulumiType;
    }

    /**
     * Time when the CSMS secrets created, in UTC format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of a secret.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * The latest version id.
     */
    public /*out*/ readonly latestVersion!: pulumi.Output<string>;
    /**
     * The secret name. The maximum length is 64 characters.
     * Only digits, letters, underscores(_), hyphens(-) and dots(.) are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the CSMS secrets.
     * If omitted, the provider-level region will be used. Changing this setting will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret ID in UUID format.
     */
    public /*out*/ readonly secretId!: pulumi.Output<string>;
    /**
     * The plaintext of a secret in text format. The maximum size is 32 KB.
     */
    public readonly secretText!: pulumi.Output<string>;
    /**
     * The CSMS secret status. Values can be: ENABLED, DISABLED, PENDING_DELETE and FROZEN.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags of a CSMS secrets, key/value pair format.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["latestVersion"] = state ? state.latestVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretId"] = state ? state.secretId : undefined;
            resourceInputs["secretText"] = state ? state.secretText : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            if ((!args || args.secretText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretText'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretText"] = args ? args.secretText : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["latestVersion"] = undefined /*out*/;
            resourceInputs["secretId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Secret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * Time when the CSMS secrets created, in UTC format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of a secret.
     */
    description?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The latest version id.
     */
    latestVersion?: pulumi.Input<string>;
    /**
     * The secret name. The maximum length is 64 characters.
     * Only digits, letters, underscores(_), hyphens(-) and dots(.) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the CSMS secrets.
     * If omitted, the provider-level region will be used. Changing this setting will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret ID in UUID format.
     */
    secretId?: pulumi.Input<string>;
    /**
     * The plaintext of a secret in text format. The maximum size is 32 KB.
     */
    secretText?: pulumi.Input<string>;
    /**
     * The CSMS secret status. Values can be: ENABLED, DISABLED, PENDING_DELETE and FROZEN.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags of a CSMS secrets, key/value pair format.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * The description of a secret.
     */
    description?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The secret name. The maximum length is 64 characters.
     * Only digits, letters, underscores(_), hyphens(-) and dots(.) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the CSMS secrets.
     * If omitted, the provider-level region will be used. Changing this setting will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The plaintext of a secret in text format. The maximum size is 32 KB.
     */
    secretText: pulumi.Input<string>;
    /**
     * The tags of a CSMS secrets, key/value pair format.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
