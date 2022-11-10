// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a WAF domain resource within HuaweiCloud.
 *
 * > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
 * used. The domain name resource can be used in Cloud Mode.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const certificate1 = new huaweicloud.waf.Certificate("certificate1", {
 *     certificate: `-----BEGIN CERTIFICATE-----
 * MIIFmQl5dh2QUAeo39TIKtadgAgh4zHx09kSgayS9Wph9LEqq7MA+2042L3J9aOa
 * DAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUR+SosWwALt6PkP0J9iOIxA6RW8gVsLwq
 * ...
 * +HhDvD/VeOHytX3RAs2GeTOtxyAV5XpKY5r+PkyUqPJj04t3d0Fopi0gNtLpMF=
 * -----END CERTIFICATE-----
 * `,
 *     privateKey: `-----BEGIN PRIVATE KEY-----
 * MIIJwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAM
 * ATAwMC4GCCsGAQUFBwIBFiJodHRwOi8vY3BzLnJvb3QteDEubGV0c2VuY3J5cHQu
 * ...
 * he8Y4IWS6wY7bCkjCWDcRQJMEhg76fsO3txE+FiYruq9RUWhiF1myv4Q6W+CyBFC
 * 1qoJFlcDyqSMo5iHq3HLjs
 * -----END PRIVATE KEY-----
 * `,
 * });
 * const domain1 = new huaweicloud.waf.Domain("domain1", {
 *     domain: "www.example.com",
 *     certificateId: certificate1.id,
 *     certificateName: certificate1.name,
 *     proxy: true,
 *     servers: [{
 *         clientProtocol: "HTTPS",
 *         serverProtocol: "HTTP",
 *         address: "119.8.0.13",
 *         port: 8080,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Domains can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Waf/domain:Domain domain_2 7902bd9e01104cb794dcb668f235e0c5
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Waf/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * Whether a domain name is connected to WAF. 0: The domain name is not connected to WAF, 1: The domain
     * name is connected to WAF.
     */
    public /*out*/ readonly accessStatus!: pulumi.Output<number>;
    /**
     * Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
     * is set to HTTPS.
     */
    public readonly certificateId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the certificate name. This parameter is mandatory
     * when `clientProtocol` is set to HTTPS.
     */
    public readonly certificateName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the domain name to be protected. For example, www.example.com or
     * *.example.com. Changing this creates a new domain.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Specifies whether to retain the policy when deleting a domain name.
     * Defaults to true.
     */
    public readonly keepPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the policy ID associated with the domain. If not specified, a new
     * policy will be created automatically. Changing this create a new domain.
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * The WAF mode. -1: bypassed, 0: disabled, 1: enabled.
     */
    public /*out*/ readonly protectStatus!: pulumi.Output<number>;
    /**
     * The protocol type of the client. The options are HTTP, HTTPS, and HTTP&HTTPS.
     */
    public /*out*/ readonly protocol!: pulumi.Output<string>;
    /**
     * Specifies whether a proxy is configured.
     */
    public readonly proxy!: pulumi.Output<boolean | undefined>;
    /**
     * The region in which to create the WAF domain resource. If omitted, the
     * provider-level region will be used. Changing this setting will push a new certificate.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies an array of origin web servers. The object structure is documented below.
     */
    public readonly servers!: pulumi.Output<outputs.Waf.DomainServer[]>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            resourceInputs["accessStatus"] = state ? state.accessStatus : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["certificateName"] = state ? state.certificateName : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["keepPolicy"] = state ? state.keepPolicy : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["protectStatus"] = state ? state.protectStatus : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["servers"] = state ? state.servers : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.servers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servers'");
            }
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["certificateName"] = args ? args.certificateName : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["keepPolicy"] = args ? args.keepPolicy : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["servers"] = args ? args.servers : undefined;
            resourceInputs["accessStatus"] = undefined /*out*/;
            resourceInputs["protectStatus"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * Whether a domain name is connected to WAF. 0: The domain name is not connected to WAF, 1: The domain
     * name is connected to WAF.
     */
    accessStatus?: pulumi.Input<number>;
    /**
     * Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
     * is set to HTTPS.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * Specifies the certificate name. This parameter is mandatory
     * when `clientProtocol` is set to HTTPS.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * Specifies the domain name to be protected. For example, www.example.com or
     * *.example.com. Changing this creates a new domain.
     */
    domain?: pulumi.Input<string>;
    /**
     * Specifies whether to retain the policy when deleting a domain name.
     * Defaults to true.
     */
    keepPolicy?: pulumi.Input<boolean>;
    /**
     * Specifies the policy ID associated with the domain. If not specified, a new
     * policy will be created automatically. Changing this create a new domain.
     */
    policyId?: pulumi.Input<string>;
    /**
     * The WAF mode. -1: bypassed, 0: disabled, 1: enabled.
     */
    protectStatus?: pulumi.Input<number>;
    /**
     * The protocol type of the client. The options are HTTP, HTTPS, and HTTP&HTTPS.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Specifies whether a proxy is configured.
     */
    proxy?: pulumi.Input<boolean>;
    /**
     * The region in which to create the WAF domain resource. If omitted, the
     * provider-level region will be used. Changing this setting will push a new certificate.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies an array of origin web servers. The object structure is documented below.
     */
    servers?: pulumi.Input<pulumi.Input<inputs.Waf.DomainServer>[]>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
     * is set to HTTPS.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * Specifies the certificate name. This parameter is mandatory
     * when `clientProtocol` is set to HTTPS.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * Specifies the domain name to be protected. For example, www.example.com or
     * *.example.com. Changing this creates a new domain.
     */
    domain: pulumi.Input<string>;
    /**
     * Specifies whether to retain the policy when deleting a domain name.
     * Defaults to true.
     */
    keepPolicy?: pulumi.Input<boolean>;
    /**
     * Specifies the policy ID associated with the domain. If not specified, a new
     * policy will be created automatically. Changing this create a new domain.
     */
    policyId?: pulumi.Input<string>;
    /**
     * Specifies whether a proxy is configured.
     */
    proxy?: pulumi.Input<boolean>;
    /**
     * The region in which to create the WAF domain resource. If omitted, the
     * provider-level region will be used. Changing this setting will push a new certificate.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies an array of origin web servers. The object structure is documented below.
     */
    servers: pulumi.Input<pulumi.Input<inputs.Waf.DomainServer>[]>;
}
