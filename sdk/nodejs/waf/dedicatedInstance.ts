// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a WAF dedicated instance resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const azName = config.requireObject("azName");
 * const ecsFlavorId = config.requireObject("ecsFlavorId");
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const securityGroupId = config.requireObject("securityGroupId");
 * const enterpriseProjectId = config.requireObject("enterpriseProjectId");
 * const instance1 = new huaweicloud.waf.DedicatedInstance("instance1", {
 *     availableZone: azName,
 *     specificationCode: "waf.instance.professional",
 *     ecsFlavor: ecsFlavorId,
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     enterpriseProjectId: enterpriseProjectId,
 *     securityGroups: [securityGroupId],
 * });
 * ```
 *
 * ## Import
 *
 * There are two ways to import WAF dedicated instance state. * Using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Waf/dedicatedInstance:DedicatedInstance test <id>
 * ```
 *
 *  * Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Waf/dedicatedInstance:DedicatedInstance test <id>/<enterprise_project_id>
 * ```
 */
export class DedicatedInstance extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedInstanceState, opts?: pulumi.CustomResourceOptions): DedicatedInstance {
        return new DedicatedInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Waf/dedicatedInstance:DedicatedInstance';

    /**
     * Returns true if the given object is an instance of DedicatedInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedInstance.__pulumiType;
    }

    /**
     * The access status of the instance. `0`: inaccessible, `1`: accessible.
     */
    public /*out*/ readonly accessStatus!: pulumi.Output<number>;
    /**
     * The available zone names for the dedicated instances. It can be
     * obtained through this data source `huaweicloud.getAvailabilityZones`. Changing this will create a new instance.
     */
    public readonly availableZone!: pulumi.Output<string>;
    /**
     * The ECS cpu architecture of instance, Default value is `x86`.
     * Changing this will create a new instance.
     */
    public readonly cpuArchitecture!: pulumi.Output<string | undefined>;
    /**
     * The flavor of the ECS used by the WAF instance. Flavors can be obtained
     * through this data source `huaweicloud.Ecs.getFlavors`. Changing this will create a new instance.
     */
    public readonly ecsFlavor!: pulumi.Output<string>;
    /**
     * The enterprise project ID of WAF dedicated instance. Changing this
     * will migrate the WAF instance to a new enterprise project.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string | undefined>;
    /**
     * The instance group ID used by the WAF dedicated instance in ELB mode.
     * Changing this will create a new instance.
     */
    public readonly groupId!: pulumi.Output<string | undefined>;
    /**
     * The name of WAF dedicated instance. Duplicate names are allowed, we suggest to keeping the
     * name unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the WAF dedicated instance. If omitted, the
     * provider-level region will be used. Changing this setting will create a new instance.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * schema: Internal; Specifies whether this is resource tenant.
     */
    public readonly resTenant!: pulumi.Output<boolean | undefined>;
    /**
     * The running status of the instance. Values are:
     */
    public /*out*/ readonly runStatus!: pulumi.Output<number>;
    /**
     * The security group of the instance. This is an array of security group
     * ids. Changing this will create a new instance.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * The id of the instance server.
     */
    public /*out*/ readonly serverId!: pulumi.Output<string>;
    /**
     * The ip of the instance service.
     */
    public /*out*/ readonly serviceIp!: pulumi.Output<string>;
    /**
     * The specification code of instance. Different specifications have
     * different throughput. Changing this will create a new instance. Values are:
     * + `waf.instance.professional` - The professional edition, throughput: 100 Mbit/s; QPS: 2,000 (Reference only).
     * + `waf.instance.enterprise` - The enterprise edition, throughput: 500 Mbit/s; QPS: 10,000 (Reference only).
     */
    public readonly specificationCode!: pulumi.Output<string>;
    /**
     * The subnet id of WAF dedicated instance VPC. Changing this will create a
     * new instance.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The instance is to support upgrades. `0`: Cannot be upgraded, `1`: Can be upgraded.
     */
    public /*out*/ readonly upgradable!: pulumi.Output<number>;
    /**
     * The VPC id of WAF dedicated instance. Changing this will create a new
     * instance.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a DedicatedInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedInstanceArgs | DedicatedInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DedicatedInstanceState | undefined;
            resourceInputs["accessStatus"] = state ? state.accessStatus : undefined;
            resourceInputs["availableZone"] = state ? state.availableZone : undefined;
            resourceInputs["cpuArchitecture"] = state ? state.cpuArchitecture : undefined;
            resourceInputs["ecsFlavor"] = state ? state.ecsFlavor : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["resTenant"] = state ? state.resTenant : undefined;
            resourceInputs["runStatus"] = state ? state.runStatus : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
            resourceInputs["serviceIp"] = state ? state.serviceIp : undefined;
            resourceInputs["specificationCode"] = state ? state.specificationCode : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["upgradable"] = state ? state.upgradable : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DedicatedInstanceArgs | undefined;
            if ((!args || args.availableZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availableZone'");
            }
            if ((!args || args.ecsFlavor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ecsFlavor'");
            }
            if ((!args || args.securityGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroups'");
            }
            if ((!args || args.specificationCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'specificationCode'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availableZone"] = args ? args.availableZone : undefined;
            resourceInputs["cpuArchitecture"] = args ? args.cpuArchitecture : undefined;
            resourceInputs["ecsFlavor"] = args ? args.ecsFlavor : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["resTenant"] = args ? args.resTenant : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["specificationCode"] = args ? args.specificationCode : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["accessStatus"] = undefined /*out*/;
            resourceInputs["runStatus"] = undefined /*out*/;
            resourceInputs["serverId"] = undefined /*out*/;
            resourceInputs["serviceIp"] = undefined /*out*/;
            resourceInputs["upgradable"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DedicatedInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedInstance resources.
 */
export interface DedicatedInstanceState {
    /**
     * The access status of the instance. `0`: inaccessible, `1`: accessible.
     */
    accessStatus?: pulumi.Input<number>;
    /**
     * The available zone names for the dedicated instances. It can be
     * obtained through this data source `huaweicloud.getAvailabilityZones`. Changing this will create a new instance.
     */
    availableZone?: pulumi.Input<string>;
    /**
     * The ECS cpu architecture of instance, Default value is `x86`.
     * Changing this will create a new instance.
     */
    cpuArchitecture?: pulumi.Input<string>;
    /**
     * The flavor of the ECS used by the WAF instance. Flavors can be obtained
     * through this data source `huaweicloud.Ecs.getFlavors`. Changing this will create a new instance.
     */
    ecsFlavor?: pulumi.Input<string>;
    /**
     * The enterprise project ID of WAF dedicated instance. Changing this
     * will migrate the WAF instance to a new enterprise project.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The instance group ID used by the WAF dedicated instance in ELB mode.
     * Changing this will create a new instance.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The name of WAF dedicated instance. Duplicate names are allowed, we suggest to keeping the
     * name unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the WAF dedicated instance. If omitted, the
     * provider-level region will be used. Changing this setting will create a new instance.
     */
    region?: pulumi.Input<string>;
    /**
     * schema: Internal; Specifies whether this is resource tenant.
     */
    resTenant?: pulumi.Input<boolean>;
    /**
     * The running status of the instance. Values are:
     */
    runStatus?: pulumi.Input<number>;
    /**
     * The security group of the instance. This is an array of security group
     * ids. Changing this will create a new instance.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the instance server.
     */
    serverId?: pulumi.Input<string>;
    /**
     * The ip of the instance service.
     */
    serviceIp?: pulumi.Input<string>;
    /**
     * The specification code of instance. Different specifications have
     * different throughput. Changing this will create a new instance. Values are:
     * + `waf.instance.professional` - The professional edition, throughput: 100 Mbit/s; QPS: 2,000 (Reference only).
     * + `waf.instance.enterprise` - The enterprise edition, throughput: 500 Mbit/s; QPS: 10,000 (Reference only).
     */
    specificationCode?: pulumi.Input<string>;
    /**
     * The subnet id of WAF dedicated instance VPC. Changing this will create a
     * new instance.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The instance is to support upgrades. `0`: Cannot be upgraded, `1`: Can be upgraded.
     */
    upgradable?: pulumi.Input<number>;
    /**
     * The VPC id of WAF dedicated instance. Changing this will create a new
     * instance.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DedicatedInstance resource.
 */
export interface DedicatedInstanceArgs {
    /**
     * The available zone names for the dedicated instances. It can be
     * obtained through this data source `huaweicloud.getAvailabilityZones`. Changing this will create a new instance.
     */
    availableZone: pulumi.Input<string>;
    /**
     * The ECS cpu architecture of instance, Default value is `x86`.
     * Changing this will create a new instance.
     */
    cpuArchitecture?: pulumi.Input<string>;
    /**
     * The flavor of the ECS used by the WAF instance. Flavors can be obtained
     * through this data source `huaweicloud.Ecs.getFlavors`. Changing this will create a new instance.
     */
    ecsFlavor: pulumi.Input<string>;
    /**
     * The enterprise project ID of WAF dedicated instance. Changing this
     * will migrate the WAF instance to a new enterprise project.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The instance group ID used by the WAF dedicated instance in ELB mode.
     * Changing this will create a new instance.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The name of WAF dedicated instance. Duplicate names are allowed, we suggest to keeping the
     * name unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the WAF dedicated instance. If omitted, the
     * provider-level region will be used. Changing this setting will create a new instance.
     */
    region?: pulumi.Input<string>;
    /**
     * schema: Internal; Specifies whether this is resource tenant.
     */
    resTenant?: pulumi.Input<boolean>;
    /**
     * The security group of the instance. This is an array of security group
     * ids. Changing this will create a new instance.
     */
    securityGroups: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The specification code of instance. Different specifications have
     * different throughput. Changing this will create a new instance. Values are:
     * + `waf.instance.professional` - The professional edition, throughput: 100 Mbit/s; QPS: 2,000 (Reference only).
     * + `waf.instance.enterprise` - The enterprise edition, throughput: 500 Mbit/s; QPS: 10,000 (Reference only).
     */
    specificationCode: pulumi.Input<string>;
    /**
     * The subnet id of WAF dedicated instance VPC. Changing this will create a
     * new instance.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The VPC id of WAF dedicated instance. Changing this will create a new
     * instance.
     */
    vpcId: pulumi.Input<string>;
}
