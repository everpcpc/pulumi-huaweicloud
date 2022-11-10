// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an API gateway API resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const tfApigwGroup = new huaweicloud.sharedapig.Group("tfApigwGroup", {description: "your descpiption"});
 * const tfApigwApi = new huaweicloud.sharedapig.Api("tfApigwApi", {
 *     groupId: tfApigwGroup.id,
 *     description: "your descpiption",
 *     tags: [
 *         "tag1",
 *         "tag2",
 *     ],
 *     visibility: 2,
 *     authType: "IAM",
 *     backendType: "HTTP",
 *     requestProtocol: "HTTPS",
 *     requestMethod: "GET",
 *     requestUri: "/test/path1",
 *     exampleSuccessResponse: "example response",
 *     httpBackend: {
 *         protocol: "HTTPS",
 *         method: "GET",
 *         uri: "/web/openapi",
 *         urlDomain: "myhuaweicloud.com",
 *         timeout: 10000,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * API can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:SharedApig/api:Api api "774438a28a574ac8a496325d1bf51807"
 * ```
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiState, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:SharedApig/api:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
     * NONE'.
     */
    public readonly authType!: pulumi.Output<string>;
    /**
     * the backend parameter list (documented below).
     */
    public readonly backendParameters!: pulumi.Output<outputs.SharedApig.ApiBackendParameter[] | undefined>;
    /**
     * Specifies the service backend type. The value can be:
     * + 'HTTP': the web service backend
     * + 'FUNCTION': the FunctionGraph service backend
     * + 'MOCK': the Mock service backend
     */
    public readonly backendType!: pulumi.Output<string>;
    /**
     * Specifies whether CORS is supported or not.
     */
    public readonly cors!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the description of the API. The description cannot exceed 255 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the example response for a failed request The length cannot
     * exceed 20,480 characters.
     */
    public readonly exampleFailureResponse!: pulumi.Output<string | undefined>;
    /**
     * Specifies the example response for a successful request. The length
     * cannot exceed 20,480 characters.
     */
    public readonly exampleSuccessResponse!: pulumi.Output<string>;
    /**
     * Specifies the configuration when backendType selected 'FUNCTION' (documented
     * below).
     */
    public readonly functionBackend!: pulumi.Output<outputs.SharedApig.ApiFunctionBackend | undefined>;
    /**
     * Specifies the ID of the API group. Changing this creates a new resource.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The name of the API group to which the API belongs.
     */
    public /*out*/ readonly groupName!: pulumi.Output<string>;
    /**
     * Specifies the configuration when backendType selected 'HTTP' (documented below).
     */
    public readonly httpBackend!: pulumi.Output<outputs.SharedApig.ApiHttpBackend | undefined>;
    /**
     * Specifies the configuration when backendType selected 'MOCK' (documented below).
     */
    public readonly mockBackend!: pulumi.Output<outputs.SharedApig.ApiMockBackend | undefined>;
    /**
     * Specifies the name of the API. An API name consists of 3–64 characters, starting with a
     * letter. Only letters, digits, and underscores (_) are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the API resource. If omitted, the provider-level
     * region will be used. Changing this creates a new API resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the request method, including 'GET','POST','PUT' and etc..
     */
    public readonly requestMethod!: pulumi.Output<string>;
    /**
     * the request parameter list (documented below).
     */
    public readonly requestParameters!: pulumi.Output<outputs.SharedApig.ApiRequestParameter[] | undefined>;
    /**
     * Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
     * which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
     */
    public readonly requestProtocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies the request path of the API. The value must comply with URI
     * specifications.
     */
    public readonly requestUri!: pulumi.Output<string>;
    /**
     * the tags of API in format of string list.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the version of the API. A maximum of 16 characters are allowed.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Specifies whether the API is available to the public. The value can be 1 (public) and
     * 2 (private). Defaults to 2.
     */
    public readonly visibility!: pulumi.Output<number | undefined>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiArgs | ApiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiState | undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["backendParameters"] = state ? state.backendParameters : undefined;
            resourceInputs["backendType"] = state ? state.backendType : undefined;
            resourceInputs["cors"] = state ? state.cors : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["exampleFailureResponse"] = state ? state.exampleFailureResponse : undefined;
            resourceInputs["exampleSuccessResponse"] = state ? state.exampleSuccessResponse : undefined;
            resourceInputs["functionBackend"] = state ? state.functionBackend : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["httpBackend"] = state ? state.httpBackend : undefined;
            resourceInputs["mockBackend"] = state ? state.mockBackend : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["requestMethod"] = state ? state.requestMethod : undefined;
            resourceInputs["requestParameters"] = state ? state.requestParameters : undefined;
            resourceInputs["requestProtocol"] = state ? state.requestProtocol : undefined;
            resourceInputs["requestUri"] = state ? state.requestUri : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ApiArgs | undefined;
            if ((!args || args.authType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authType'");
            }
            if ((!args || args.backendType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendType'");
            }
            if ((!args || args.exampleSuccessResponse === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exampleSuccessResponse'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.requestMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestMethod'");
            }
            if ((!args || args.requestUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestUri'");
            }
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["backendParameters"] = args ? args.backendParameters : undefined;
            resourceInputs["backendType"] = args ? args.backendType : undefined;
            resourceInputs["cors"] = args ? args.cors : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["exampleFailureResponse"] = args ? args.exampleFailureResponse : undefined;
            resourceInputs["exampleSuccessResponse"] = args ? args.exampleSuccessResponse : undefined;
            resourceInputs["functionBackend"] = args ? args.functionBackend : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["httpBackend"] = args ? args.httpBackend : undefined;
            resourceInputs["mockBackend"] = args ? args.mockBackend : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestMethod"] = args ? args.requestMethod : undefined;
            resourceInputs["requestParameters"] = args ? args.requestParameters : undefined;
            resourceInputs["requestProtocol"] = args ? args.requestProtocol : undefined;
            resourceInputs["requestUri"] = args ? args.requestUri : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["groupName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Api.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Api resources.
 */
export interface ApiState {
    /**
     * Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
     * NONE'.
     */
    authType?: pulumi.Input<string>;
    /**
     * the backend parameter list (documented below).
     */
    backendParameters?: pulumi.Input<pulumi.Input<inputs.SharedApig.ApiBackendParameter>[]>;
    /**
     * Specifies the service backend type. The value can be:
     * + 'HTTP': the web service backend
     * + 'FUNCTION': the FunctionGraph service backend
     * + 'MOCK': the Mock service backend
     */
    backendType?: pulumi.Input<string>;
    /**
     * Specifies whether CORS is supported or not.
     */
    cors?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the API. The description cannot exceed 255 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the example response for a failed request The length cannot
     * exceed 20,480 characters.
     */
    exampleFailureResponse?: pulumi.Input<string>;
    /**
     * Specifies the example response for a successful request. The length
     * cannot exceed 20,480 characters.
     */
    exampleSuccessResponse?: pulumi.Input<string>;
    /**
     * Specifies the configuration when backendType selected 'FUNCTION' (documented
     * below).
     */
    functionBackend?: pulumi.Input<inputs.SharedApig.ApiFunctionBackend>;
    /**
     * Specifies the ID of the API group. Changing this creates a new resource.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The name of the API group to which the API belongs.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Specifies the configuration when backendType selected 'HTTP' (documented below).
     */
    httpBackend?: pulumi.Input<inputs.SharedApig.ApiHttpBackend>;
    /**
     * Specifies the configuration when backendType selected 'MOCK' (documented below).
     */
    mockBackend?: pulumi.Input<inputs.SharedApig.ApiMockBackend>;
    /**
     * Specifies the name of the API. An API name consists of 3–64 characters, starting with a
     * letter. Only letters, digits, and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the API resource. If omitted, the provider-level
     * region will be used. Changing this creates a new API resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the request method, including 'GET','POST','PUT' and etc..
     */
    requestMethod?: pulumi.Input<string>;
    /**
     * the request parameter list (documented below).
     */
    requestParameters?: pulumi.Input<pulumi.Input<inputs.SharedApig.ApiRequestParameter>[]>;
    /**
     * Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
     * which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
     */
    requestProtocol?: pulumi.Input<string>;
    /**
     * Specifies the request path of the API. The value must comply with URI
     * specifications.
     */
    requestUri?: pulumi.Input<string>;
    /**
     * the tags of API in format of string list.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the version of the API. A maximum of 16 characters are allowed.
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies whether the API is available to the public. The value can be 1 (public) and
     * 2 (private). Defaults to 2.
     */
    visibility?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    /**
     * Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
     * NONE'.
     */
    authType: pulumi.Input<string>;
    /**
     * the backend parameter list (documented below).
     */
    backendParameters?: pulumi.Input<pulumi.Input<inputs.SharedApig.ApiBackendParameter>[]>;
    /**
     * Specifies the service backend type. The value can be:
     * + 'HTTP': the web service backend
     * + 'FUNCTION': the FunctionGraph service backend
     * + 'MOCK': the Mock service backend
     */
    backendType: pulumi.Input<string>;
    /**
     * Specifies whether CORS is supported or not.
     */
    cors?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the API. The description cannot exceed 255 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the example response for a failed request The length cannot
     * exceed 20,480 characters.
     */
    exampleFailureResponse?: pulumi.Input<string>;
    /**
     * Specifies the example response for a successful request. The length
     * cannot exceed 20,480 characters.
     */
    exampleSuccessResponse: pulumi.Input<string>;
    /**
     * Specifies the configuration when backendType selected 'FUNCTION' (documented
     * below).
     */
    functionBackend?: pulumi.Input<inputs.SharedApig.ApiFunctionBackend>;
    /**
     * Specifies the ID of the API group. Changing this creates a new resource.
     */
    groupId: pulumi.Input<string>;
    /**
     * Specifies the configuration when backendType selected 'HTTP' (documented below).
     */
    httpBackend?: pulumi.Input<inputs.SharedApig.ApiHttpBackend>;
    /**
     * Specifies the configuration when backendType selected 'MOCK' (documented below).
     */
    mockBackend?: pulumi.Input<inputs.SharedApig.ApiMockBackend>;
    /**
     * Specifies the name of the API. An API name consists of 3–64 characters, starting with a
     * letter. Only letters, digits, and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the API resource. If omitted, the provider-level
     * region will be used. Changing this creates a new API resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the request method, including 'GET','POST','PUT' and etc..
     */
    requestMethod: pulumi.Input<string>;
    /**
     * the request parameter list (documented below).
     */
    requestParameters?: pulumi.Input<pulumi.Input<inputs.SharedApig.ApiRequestParameter>[]>;
    /**
     * Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
     * which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
     */
    requestProtocol?: pulumi.Input<string>;
    /**
     * Specifies the request path of the API. The value must comply with URI
     * specifications.
     */
    requestUri: pulumi.Input<string>;
    /**
     * the tags of API in format of string list.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the version of the API. A maximum of 16 characters are allowed.
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies whether the API is available to the public. The value can be 1 (public) and
     * 2 (private). Defaults to 2.
     */
    visibility?: pulumi.Input<number>;
}
