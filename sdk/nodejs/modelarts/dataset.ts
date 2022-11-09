// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages ModelArts dataset resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const name = config.requireObject("name");
 * const outputObsPath = config.requireObject("outputObsPath");
 * const inputObsPath = config.requireObject("inputObsPath");
 * const test = new huaweicloud.modelarts.Dataset("test", {
 *     type: 1,
 *     outputPath: outputObsPath,
 *     description: "Terraform Demo",
 *     dataSource: {
 *         path: inputObsPath,
 *     },
 *     labels: [{
 *         name: "foo",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * The datasets can be imported by `id`.
 *
 * ```sh
 *  $ pulumi import huaweicloud:ModelArts/dataset:Dataset test yiROKoTTjtwjvP71yLG
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`data_source.0.path`, `data_source.0.queue_name`, `data_source.0.database_name`, `data_source.0.table_name`, `data_source.0.cluster_id`, `data_source.0.user_name` and `data_source.0.password`. It is generally recommended running `terraform plan` after importing a dataset. You can then decide if changes should be applied to the dataset, or the resource definition should be updated to align with the dataset. Also you can ignore changes as below. resource "huaweicloud_modelarts_dataset" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  data_source.0.path, data_source.0.queue_name, data_source.0.database_name, data_source.0.table_name,
 *
 *  data_source.0.cluster_id, data_source.0.user_name, data_source.0.password,
 *
 *  ]
 *
 *  } }
 */
export class Dataset extends pulumi.CustomResource {
    /**
     * Get an existing Dataset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetState, opts?: pulumi.CustomResourceOptions): Dataset {
        return new Dataset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:ModelArts/dataset:Dataset';

    /**
     * Returns true if the given object is an instance of Dataset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dataset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dataset.__pulumiType;
    }

    /**
     * The dataset creation time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * dataset format. Valid values include: `Default`, `CarbonData`: Carbon format(Supported only for
     * table type datasets).
     */
    public /*out*/ readonly dataFormat!: pulumi.Output<string>;
    /**
     * Specifies the data sources which be used to imported the source data (such
     * as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    public readonly dataSource!: pulumi.Output<outputs.ModelArts.DatasetDataSource>;
    /**
     * Specifies the description of dataset. It contains a maximum of 256 characters and
     * cannot contain special characters `!<>=&"'`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to import labeled files.
     * Default value is `true`. Changing this parameter will create a new resource.
     */
    public readonly importLabeledEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the custom format information of labeled files when import
     * labeled files for Text classification. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    public readonly labelFormat!: pulumi.Output<outputs.ModelArts.DatasetLabelFormat | undefined>;
    /**
     * Specifies labels information. Structure is documented below.
     */
    public readonly labels!: pulumi.Output<outputs.ModelArts.DatasetLabel[] | undefined>;
    /**
     * Specifies the name of label.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the OBS path for storing output files such as labeled files.
     * The path cannot be the same as the import path or subdirectory of the import path.
     * Changing this parameter will create a new resource.
     */
    public readonly outputPath!: pulumi.Output<string>;
    /**
     * The region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the schema information of source data when `type` is `400`.
     * Structure is documented below. Changing this parameter will create a new resource.
     */
    public readonly schemas!: pulumi.Output<outputs.ModelArts.DatasetSchema[] | undefined>;
    /**
     * Dataset status. Valid values are as follows:
     * + **0**: Creating.
     * + **1**: Completed.
     * + **2**: Deleting.
     * + **3**: Deleted.
     * + **4**: Exception.
     * + **5**: Syncing.
     * + **6**: Releasing.
     * + **7**: Version switching.
     * + **8**: Importing.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Specifies Label type for text classification.
     * The optional values are as follows:
     */
    public readonly type!: pulumi.Output<number>;

    /**
     * Create a Dataset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetArgs | DatasetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatasetState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dataFormat"] = state ? state.dataFormat : undefined;
            resourceInputs["dataSource"] = state ? state.dataSource : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["importLabeledEnabled"] = state ? state.importLabeledEnabled : undefined;
            resourceInputs["labelFormat"] = state ? state.labelFormat : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outputPath"] = state ? state.outputPath : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["schemas"] = state ? state.schemas : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DatasetArgs | undefined;
            if ((!args || args.dataSource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataSource'");
            }
            if ((!args || args.outputPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'outputPath'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["dataSource"] = args ? args.dataSource : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["importLabeledEnabled"] = args ? args.importLabeledEnabled : undefined;
            resourceInputs["labelFormat"] = args ? args.labelFormat : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputPath"] = args ? args.outputPath : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["schemas"] = args ? args.schemas : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dataFormat"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Dataset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dataset resources.
 */
export interface DatasetState {
    /**
     * The dataset creation time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * dataset format. Valid values include: `Default`, `CarbonData`: Carbon format(Supported only for
     * table type datasets).
     */
    dataFormat?: pulumi.Input<string>;
    /**
     * Specifies the data sources which be used to imported the source data (such
     * as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    dataSource?: pulumi.Input<inputs.ModelArts.DatasetDataSource>;
    /**
     * Specifies the description of dataset. It contains a maximum of 256 characters and
     * cannot contain special characters `!<>=&"'`.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to import labeled files.
     * Default value is `true`. Changing this parameter will create a new resource.
     */
    importLabeledEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the custom format information of labeled files when import
     * labeled files for Text classification. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    labelFormat?: pulumi.Input<inputs.ModelArts.DatasetLabelFormat>;
    /**
     * Specifies labels information. Structure is documented below.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ModelArts.DatasetLabel>[]>;
    /**
     * Specifies the name of label.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the OBS path for storing output files such as labeled files.
     * The path cannot be the same as the import path or subdirectory of the import path.
     * Changing this parameter will create a new resource.
     */
    outputPath?: pulumi.Input<string>;
    /**
     * The region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the schema information of source data when `type` is `400`.
     * Structure is documented below. Changing this parameter will create a new resource.
     */
    schemas?: pulumi.Input<pulumi.Input<inputs.ModelArts.DatasetSchema>[]>;
    /**
     * Dataset status. Valid values are as follows:
     * + **0**: Creating.
     * + **1**: Completed.
     * + **2**: Deleting.
     * + **3**: Deleted.
     * + **4**: Exception.
     * + **5**: Syncing.
     * + **6**: Releasing.
     * + **7**: Version switching.
     * + **8**: Importing.
     */
    status?: pulumi.Input<number>;
    /**
     * Specifies Label type for text classification.
     * The optional values are as follows:
     */
    type?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Dataset resource.
 */
export interface DatasetArgs {
    /**
     * Specifies the data sources which be used to imported the source data (such
     * as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    dataSource: pulumi.Input<inputs.ModelArts.DatasetDataSource>;
    /**
     * Specifies the description of dataset. It contains a maximum of 256 characters and
     * cannot contain special characters `!<>=&"'`.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to import labeled files.
     * Default value is `true`. Changing this parameter will create a new resource.
     */
    importLabeledEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the custom format information of labeled files when import
     * labeled files for Text classification. Structure is documented below.
     * Changing this parameter will create a new resource.
     */
    labelFormat?: pulumi.Input<inputs.ModelArts.DatasetLabelFormat>;
    /**
     * Specifies labels information. Structure is documented below.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ModelArts.DatasetLabel>[]>;
    /**
     * Specifies the name of label.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the OBS path for storing output files such as labeled files.
     * The path cannot be the same as the import path or subdirectory of the import path.
     * Changing this parameter will create a new resource.
     */
    outputPath: pulumi.Input<string>;
    /**
     * The region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the schema information of source data when `type` is `400`.
     * Structure is documented below. Changing this parameter will create a new resource.
     */
    schemas?: pulumi.Input<pulumi.Input<inputs.ModelArts.DatasetSchema>[]>;
    /**
     * Specifies Label type for text classification.
     * The optional values are as follows:
     */
    type: pulumi.Input<number>;
}
