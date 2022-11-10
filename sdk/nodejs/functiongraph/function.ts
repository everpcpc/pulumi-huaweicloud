// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Function resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### With base64 func code
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const f1 = new huaweicloud.FunctionGraph.Function("f_1", {
 *     agency: "test",
 *     app: "default",
 *     codeType: "inline",
 *     description: "fuction test",
 *     funcCode: "aW1wb3J0IGpzb24KZGVmIGhhbmRsZXIgKGV2ZW50LCBjb250ZXh0KToKICAgIG91dHB1dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganNvbi5kdW1wcyhldmVudCkKICAgIHJldHVybiBvdXRwdXQ=",
 *     handler: "test.handler",
 *     memorySize: 128,
 *     runtime: "Python2.7",
 *     timeout: 3,
 * });
 * ```
 * ### With text code
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const f1 = new huaweicloud.FunctionGraph.Function("f_1", {
 *     agency: "test",
 *     app: "default",
 *     codeType: "inline",
 *     description: "fuction test",
 *     funcCode: `# -*- coding:utf-8 -*-
 * import json
 * def handler (event, context):
 *     return {
 *         "statusCode": 200,
 *         "isBase64Encoded": False,
 *         "body": json.dumps(event),
 *         "headers": {
 *             "Content-Type": "application/json"
 *         }
 *     }
 * `,
 *     handler: "test.handler",
 *     memorySize: 128,
 *     runtime: "Python2.7",
 *     timeout: 3,
 * });
 * ```
 * ### Create function using SWR image
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const functionName = config.requireObject("functionName");
 * const agencyName = config.requireObject("agencyName");
 * const imageUrl = config.requireObject("imageUrl");
 * const bySwrImage = new huaweicloud.functiongraph.Function("bySwrImage", {
 *     agency: agencyName,
 *     app: "default",
 *     runtime: "Custom Image",
 *     memorySize: 128,
 *     timeout: 3,
 *     customImage: {
 *         url: imageUrl,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Functions can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:FunctionGraph/function:Function my-func 7117d38e-4c8f-4624-a505-bd96b97d024c
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:FunctionGraph/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Specifies the agency. This parameter is mandatory if the function needs to access other
     * cloud services.
     */
    public readonly agency!: pulumi.Output<string | undefined>;
    /**
     * Specifies the group to which the function belongs.
     */
    public readonly app!: pulumi.Output<string | undefined>;
    /**
     * Specifies An execution agency enables you to obtain a token or an AK/SK for
     * accessing other cloud services.
     */
    public readonly appAgency!: pulumi.Output<string>;
    /**
     * Specifies the name of a function file, This field is mandatory only when coeType
     * is set to jar or zip.
     */
    public readonly codeFilename!: pulumi.Output<string>;
    /**
     * Specifies the function code type, which can be:
     * + **inline**: inline code.
     * + **zip**: ZIP file.
     * + **jar**: JAR file or java functions.
     * + **obs**: function code stored in an OBS bucket.
     */
    public readonly codeType!: pulumi.Output<string>;
    /**
     * Specifies the code url. This parameter is mandatory when codeType is set to obs.
     */
    public readonly codeUrl!: pulumi.Output<string | undefined>;
    /**
     * Specifies the custom image configuration for creating function.
     * The object structure is documented below.
     * Changing this will create a new resource.
     */
    public readonly customImage!: pulumi.Output<outputs.FunctionGraph.FunctionCustomImage>;
    /**
     * Specifies the ID list of the dependencies.
     */
    public readonly dependLists!: pulumi.Output<string[]>;
    /**
     * Specifies the description of the function.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the key/value information defined to be encrypted for the
     * function. The format is the same as `userData`.
     */
    public readonly encryptedUserData!: pulumi.Output<string | undefined>;
    /**
     * Specifies the enterprise project id of the function.
     * Changing this will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Specifies the function code. When codeType is set to inline, zip, or jar, this
     * parameter is mandatory, and the code can be encoded using Base64 or just with the text code.
     */
    public readonly funcCode!: pulumi.Output<string | undefined>;
    /**
     * Specifies the file system list. The `funcMounts` object structure is documented
     * below.
     */
    public readonly funcMounts!: pulumi.Output<outputs.FunctionGraph.FunctionFuncMount[]>;
    /**
     * Specifies the FunctionGraph version, defaults to **v1**.
     * + **v1**: Hosts event-driven functions in a serverless context.
     * + **v2**: Next-generation function hosting service powered by Huawei YuanRong architecture.
     */
    public readonly functiongraphVersion!: pulumi.Output<string>;
    /**
     * Specifies the entry point of the function.
     */
    public readonly handler!: pulumi.Output<string>;
    /**
     * Specifies the initializer of the function.
     */
    public readonly initializerHandler!: pulumi.Output<string>;
    /**
     * Specifies the maximum duration the function can be initialized. Value range:
     * 1s to 300s.
     */
    public readonly initializerTimeout!: pulumi.Output<number>;
    /**
     * Specifies the memory size(MB) allocated to the function.
     */
    public readonly memorySize!: pulumi.Output<number>;
    /**
     * Specifies the user group ID, a non-0 integer from –1 to 65534. Default to
     * -1.
     */
    public readonly mountUserGroupId!: pulumi.Output<number>;
    /**
     * Specifies the user ID, a non-0 integer from –1 to 65534. Default to -1.
     */
    public readonly mountUserId!: pulumi.Output<number>;
    /**
     * Specifies the name of the function.
     * Changing this will create a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the network ID of subnet.
     */
    public readonly networkId!: pulumi.Output<string | undefined>;
    /**
     * @deprecated use app instead
     */
    public readonly package!: pulumi.Output<string | undefined>;
    /**
     * Specifies the region in which to create the Function resource. If omitted, the
     * provider-level region will be used. Changing this will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the environment for executing the function.
     * If the function is created using a SWR image, set this parameter to `Custom Image`.
     * Changing this will create a new resource.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * Specifies the timeout interval of the function, ranges from 3s to 900s.
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Uniform Resource Name
     */
    public /*out*/ readonly urn!: pulumi.Output<string>;
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * The version of the function
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * Specifies the ID of VPC.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * @deprecated use agency instead
     */
    public readonly xrole!: pulumi.Output<string | undefined>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["agency"] = state ? state.agency : undefined;
            resourceInputs["app"] = state ? state.app : undefined;
            resourceInputs["appAgency"] = state ? state.appAgency : undefined;
            resourceInputs["codeFilename"] = state ? state.codeFilename : undefined;
            resourceInputs["codeType"] = state ? state.codeType : undefined;
            resourceInputs["codeUrl"] = state ? state.codeUrl : undefined;
            resourceInputs["customImage"] = state ? state.customImage : undefined;
            resourceInputs["dependLists"] = state ? state.dependLists : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encryptedUserData"] = state ? state.encryptedUserData : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["funcCode"] = state ? state.funcCode : undefined;
            resourceInputs["funcMounts"] = state ? state.funcMounts : undefined;
            resourceInputs["functiongraphVersion"] = state ? state.functiongraphVersion : undefined;
            resourceInputs["handler"] = state ? state.handler : undefined;
            resourceInputs["initializerHandler"] = state ? state.initializerHandler : undefined;
            resourceInputs["initializerTimeout"] = state ? state.initializerTimeout : undefined;
            resourceInputs["memorySize"] = state ? state.memorySize : undefined;
            resourceInputs["mountUserGroupId"] = state ? state.mountUserGroupId : undefined;
            resourceInputs["mountUserId"] = state ? state.mountUserId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["package"] = state ? state.package : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["urn"] = state ? state.urn : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["xrole"] = state ? state.xrole : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.memorySize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memorySize'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            resourceInputs["agency"] = args ? args.agency : undefined;
            resourceInputs["app"] = args ? args.app : undefined;
            resourceInputs["appAgency"] = args ? args.appAgency : undefined;
            resourceInputs["codeFilename"] = args ? args.codeFilename : undefined;
            resourceInputs["codeType"] = args ? args.codeType : undefined;
            resourceInputs["codeUrl"] = args ? args.codeUrl : undefined;
            resourceInputs["customImage"] = args ? args.customImage : undefined;
            resourceInputs["dependLists"] = args ? args.dependLists : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encryptedUserData"] = args ? args.encryptedUserData : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["funcCode"] = args ? args.funcCode : undefined;
            resourceInputs["funcMounts"] = args ? args.funcMounts : undefined;
            resourceInputs["functiongraphVersion"] = args ? args.functiongraphVersion : undefined;
            resourceInputs["handler"] = args ? args.handler : undefined;
            resourceInputs["initializerHandler"] = args ? args.initializerHandler : undefined;
            resourceInputs["initializerTimeout"] = args ? args.initializerTimeout : undefined;
            resourceInputs["memorySize"] = args ? args.memorySize : undefined;
            resourceInputs["mountUserGroupId"] = args ? args.mountUserGroupId : undefined;
            resourceInputs["mountUserId"] = args ? args.mountUserId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["package"] = args ? args.package : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["xrole"] = args ? args.xrole : undefined;
            resourceInputs["urn"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * Specifies the agency. This parameter is mandatory if the function needs to access other
     * cloud services.
     */
    agency?: pulumi.Input<string>;
    /**
     * Specifies the group to which the function belongs.
     */
    app?: pulumi.Input<string>;
    /**
     * Specifies An execution agency enables you to obtain a token or an AK/SK for
     * accessing other cloud services.
     */
    appAgency?: pulumi.Input<string>;
    /**
     * Specifies the name of a function file, This field is mandatory only when coeType
     * is set to jar or zip.
     */
    codeFilename?: pulumi.Input<string>;
    /**
     * Specifies the function code type, which can be:
     * + **inline**: inline code.
     * + **zip**: ZIP file.
     * + **jar**: JAR file or java functions.
     * + **obs**: function code stored in an OBS bucket.
     */
    codeType?: pulumi.Input<string>;
    /**
     * Specifies the code url. This parameter is mandatory when codeType is set to obs.
     */
    codeUrl?: pulumi.Input<string>;
    /**
     * Specifies the custom image configuration for creating function.
     * The object structure is documented below.
     * Changing this will create a new resource.
     */
    customImage?: pulumi.Input<inputs.FunctionGraph.FunctionCustomImage>;
    /**
     * Specifies the ID list of the dependencies.
     */
    dependLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the key/value information defined to be encrypted for the
     * function. The format is the same as `userData`.
     */
    encryptedUserData?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project id of the function.
     * Changing this will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the function code. When codeType is set to inline, zip, or jar, this
     * parameter is mandatory, and the code can be encoded using Base64 or just with the text code.
     */
    funcCode?: pulumi.Input<string>;
    /**
     * Specifies the file system list. The `funcMounts` object structure is documented
     * below.
     */
    funcMounts?: pulumi.Input<pulumi.Input<inputs.FunctionGraph.FunctionFuncMount>[]>;
    /**
     * Specifies the FunctionGraph version, defaults to **v1**.
     * + **v1**: Hosts event-driven functions in a serverless context.
     * + **v2**: Next-generation function hosting service powered by Huawei YuanRong architecture.
     */
    functiongraphVersion?: pulumi.Input<string>;
    /**
     * Specifies the entry point of the function.
     */
    handler?: pulumi.Input<string>;
    /**
     * Specifies the initializer of the function.
     */
    initializerHandler?: pulumi.Input<string>;
    /**
     * Specifies the maximum duration the function can be initialized. Value range:
     * 1s to 300s.
     */
    initializerTimeout?: pulumi.Input<number>;
    /**
     * Specifies the memory size(MB) allocated to the function.
     */
    memorySize?: pulumi.Input<number>;
    /**
     * Specifies the user group ID, a non-0 integer from –1 to 65534. Default to
     * -1.
     */
    mountUserGroupId?: pulumi.Input<number>;
    /**
     * Specifies the user ID, a non-0 integer from –1 to 65534. Default to -1.
     */
    mountUserId?: pulumi.Input<number>;
    /**
     * Specifies the name of the function.
     * Changing this will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the network ID of subnet.
     */
    networkId?: pulumi.Input<string>;
    /**
     * @deprecated use app instead
     */
    package?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the Function resource. If omitted, the
     * provider-level region will be used. Changing this will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the environment for executing the function.
     * If the function is created using a SWR image, set this parameter to `Custom Image`.
     * Changing this will create a new resource.
     */
    runtime?: pulumi.Input<string>;
    /**
     * Specifies the timeout interval of the function, ranges from 3s to 900s.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Uniform Resource Name
     */
    urn?: pulumi.Input<string>;
    userData?: pulumi.Input<string>;
    /**
     * The version of the function
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies the ID of VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * @deprecated use agency instead
     */
    xrole?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Specifies the agency. This parameter is mandatory if the function needs to access other
     * cloud services.
     */
    agency?: pulumi.Input<string>;
    /**
     * Specifies the group to which the function belongs.
     */
    app?: pulumi.Input<string>;
    /**
     * Specifies An execution agency enables you to obtain a token or an AK/SK for
     * accessing other cloud services.
     */
    appAgency?: pulumi.Input<string>;
    /**
     * Specifies the name of a function file, This field is mandatory only when coeType
     * is set to jar or zip.
     */
    codeFilename?: pulumi.Input<string>;
    /**
     * Specifies the function code type, which can be:
     * + **inline**: inline code.
     * + **zip**: ZIP file.
     * + **jar**: JAR file or java functions.
     * + **obs**: function code stored in an OBS bucket.
     */
    codeType?: pulumi.Input<string>;
    /**
     * Specifies the code url. This parameter is mandatory when codeType is set to obs.
     */
    codeUrl?: pulumi.Input<string>;
    /**
     * Specifies the custom image configuration for creating function.
     * The object structure is documented below.
     * Changing this will create a new resource.
     */
    customImage?: pulumi.Input<inputs.FunctionGraph.FunctionCustomImage>;
    /**
     * Specifies the ID list of the dependencies.
     */
    dependLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the key/value information defined to be encrypted for the
     * function. The format is the same as `userData`.
     */
    encryptedUserData?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project id of the function.
     * Changing this will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the function code. When codeType is set to inline, zip, or jar, this
     * parameter is mandatory, and the code can be encoded using Base64 or just with the text code.
     */
    funcCode?: pulumi.Input<string>;
    /**
     * Specifies the file system list. The `funcMounts` object structure is documented
     * below.
     */
    funcMounts?: pulumi.Input<pulumi.Input<inputs.FunctionGraph.FunctionFuncMount>[]>;
    /**
     * Specifies the FunctionGraph version, defaults to **v1**.
     * + **v1**: Hosts event-driven functions in a serverless context.
     * + **v2**: Next-generation function hosting service powered by Huawei YuanRong architecture.
     */
    functiongraphVersion?: pulumi.Input<string>;
    /**
     * Specifies the entry point of the function.
     */
    handler?: pulumi.Input<string>;
    /**
     * Specifies the initializer of the function.
     */
    initializerHandler?: pulumi.Input<string>;
    /**
     * Specifies the maximum duration the function can be initialized. Value range:
     * 1s to 300s.
     */
    initializerTimeout?: pulumi.Input<number>;
    /**
     * Specifies the memory size(MB) allocated to the function.
     */
    memorySize: pulumi.Input<number>;
    /**
     * Specifies the user group ID, a non-0 integer from –1 to 65534. Default to
     * -1.
     */
    mountUserGroupId?: pulumi.Input<number>;
    /**
     * Specifies the user ID, a non-0 integer from –1 to 65534. Default to -1.
     */
    mountUserId?: pulumi.Input<number>;
    /**
     * Specifies the name of the function.
     * Changing this will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the network ID of subnet.
     */
    networkId?: pulumi.Input<string>;
    /**
     * @deprecated use app instead
     */
    package?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the Function resource. If omitted, the
     * provider-level region will be used. Changing this will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the environment for executing the function.
     * If the function is created using a SWR image, set this parameter to `Custom Image`.
     * Changing this will create a new resource.
     */
    runtime: pulumi.Input<string>;
    /**
     * Specifies the timeout interval of the function, ranges from 3s to 900s.
     */
    timeout: pulumi.Input<number>;
    userData?: pulumi.Input<string>;
    /**
     * Specifies the ID of VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * @deprecated use agency instead
     */
    xrole?: pulumi.Input<string>;
}
