// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedapig

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Publish a new version of the API
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedApig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			instanceId := cfg.RequireObject("instanceId")
//			envId := cfg.RequireObject("envId")
//			apiId := cfg.RequireObject("apiId")
//			_, err := DedicatedApig.NewApiPublishment(ctx, "default", &DedicatedApig.ApiPublishmentArgs{
//				InstanceId: pulumi.Any(instanceId),
//				EnvId:      pulumi.Any(envId),
//				ApiId:      pulumi.Any(apiId),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Switch to a specified version of the API which is published
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedApig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			instanceId := cfg.RequireObject("instanceId")
//			envId := cfg.RequireObject("envId")
//			apiId := cfg.RequireObject("apiId")
//			versionId := cfg.RequireObject("versionId")
//			_, err := DedicatedApig.NewApiPublishment(ctx, "default", &DedicatedApig.ApiPublishmentArgs{
//				InstanceId: pulumi.Any(instanceId),
//				EnvId:      pulumi.Any(envId),
//				ApiId:      pulumi.Any(apiId),
//				VersionId:  pulumi.Any(versionId),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// APIs can be imported using their `instance_id`, `env_id` and `api_id`, separated by slashes, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:DedicatedApig/apiPublishment:ApiPublishment test
//
// ```
//
//	9b0a0a2f97aa43afbf7d852e3ba6a6f9/c5b32727186c4fe6b60408a8a297be09/9a3b3484c08545f9b9b0dcb2de0f5b8a
type ApiPublishment struct {
	pulumi.CustomResourceState

	// Specifies the API ID to be published or already published.
	// Changing this will create a new publishment resource.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Specifies the description of the current publishment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the environment ID to which the current version of the API will be
	// published or has been published. Changing this will create a new publishment resource.
	EnvId pulumi.StringOutput `pulumi:"envId"`
	// Environment name to which the current version of the API is published.
	EnvName pulumi.StringOutput `pulumi:"envName"`
	// All publish informations of the API. The structure is documented below.
	Histories ApiPublishmentHistoryArrayOutput `pulumi:"histories"`
	// Specifies an ID of the APIG dedicated instance to which the API belongs
	// to. Changing this will create a new publishment resource.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The publish ID of the API in current environment.
	PublishId pulumi.StringOutput `pulumi:"publishId"`
	// Time when the current version was published.
	PublishTime pulumi.StringOutput `pulumi:"publishTime"`
	// Specifies the region in which to publish APIs.
	// If omitted, the provider-level region will be used. Changing this will create a new publishment resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the version ID of the current publishment.
	VersionId pulumi.StringPtrOutput `pulumi:"versionId"`
}

// NewApiPublishment registers a new resource with the given unique name, arguments, and options.
func NewApiPublishment(ctx *pulumi.Context,
	name string, args *ApiPublishmentArgs, opts ...pulumi.ResourceOption) (*ApiPublishment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.EnvId == nil {
		return nil, errors.New("invalid value for required argument 'EnvId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource ApiPublishment
	err := ctx.RegisterResource("huaweicloud:DedicatedApig/apiPublishment:ApiPublishment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiPublishment gets an existing ApiPublishment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiPublishment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiPublishmentState, opts ...pulumi.ResourceOption) (*ApiPublishment, error) {
	var resource ApiPublishment
	err := ctx.ReadResource("huaweicloud:DedicatedApig/apiPublishment:ApiPublishment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiPublishment resources.
type apiPublishmentState struct {
	// Specifies the API ID to be published or already published.
	// Changing this will create a new publishment resource.
	ApiId *string `pulumi:"apiId"`
	// Specifies the description of the current publishment.
	Description *string `pulumi:"description"`
	// Specifies the environment ID to which the current version of the API will be
	// published or has been published. Changing this will create a new publishment resource.
	EnvId *string `pulumi:"envId"`
	// Environment name to which the current version of the API is published.
	EnvName *string `pulumi:"envName"`
	// All publish informations of the API. The structure is documented below.
	Histories []ApiPublishmentHistory `pulumi:"histories"`
	// Specifies an ID of the APIG dedicated instance to which the API belongs
	// to. Changing this will create a new publishment resource.
	InstanceId *string `pulumi:"instanceId"`
	// The publish ID of the API in current environment.
	PublishId *string `pulumi:"publishId"`
	// Time when the current version was published.
	PublishTime *string `pulumi:"publishTime"`
	// Specifies the region in which to publish APIs.
	// If omitted, the provider-level region will be used. Changing this will create a new publishment resource.
	Region *string `pulumi:"region"`
	// Specifies the version ID of the current publishment.
	VersionId *string `pulumi:"versionId"`
}

type ApiPublishmentState struct {
	// Specifies the API ID to be published or already published.
	// Changing this will create a new publishment resource.
	ApiId pulumi.StringPtrInput
	// Specifies the description of the current publishment.
	Description pulumi.StringPtrInput
	// Specifies the environment ID to which the current version of the API will be
	// published or has been published. Changing this will create a new publishment resource.
	EnvId pulumi.StringPtrInput
	// Environment name to which the current version of the API is published.
	EnvName pulumi.StringPtrInput
	// All publish informations of the API. The structure is documented below.
	Histories ApiPublishmentHistoryArrayInput
	// Specifies an ID of the APIG dedicated instance to which the API belongs
	// to. Changing this will create a new publishment resource.
	InstanceId pulumi.StringPtrInput
	// The publish ID of the API in current environment.
	PublishId pulumi.StringPtrInput
	// Time when the current version was published.
	PublishTime pulumi.StringPtrInput
	// Specifies the region in which to publish APIs.
	// If omitted, the provider-level region will be used. Changing this will create a new publishment resource.
	Region pulumi.StringPtrInput
	// Specifies the version ID of the current publishment.
	VersionId pulumi.StringPtrInput
}

func (ApiPublishmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPublishmentState)(nil)).Elem()
}

type apiPublishmentArgs struct {
	// Specifies the API ID to be published or already published.
	// Changing this will create a new publishment resource.
	ApiId string `pulumi:"apiId"`
	// Specifies the description of the current publishment.
	Description *string `pulumi:"description"`
	// Specifies the environment ID to which the current version of the API will be
	// published or has been published. Changing this will create a new publishment resource.
	EnvId string `pulumi:"envId"`
	// Specifies an ID of the APIG dedicated instance to which the API belongs
	// to. Changing this will create a new publishment resource.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the region in which to publish APIs.
	// If omitted, the provider-level region will be used. Changing this will create a new publishment resource.
	Region *string `pulumi:"region"`
	// Specifies the version ID of the current publishment.
	VersionId *string `pulumi:"versionId"`
}

// The set of arguments for constructing a ApiPublishment resource.
type ApiPublishmentArgs struct {
	// Specifies the API ID to be published or already published.
	// Changing this will create a new publishment resource.
	ApiId pulumi.StringInput
	// Specifies the description of the current publishment.
	Description pulumi.StringPtrInput
	// Specifies the environment ID to which the current version of the API will be
	// published or has been published. Changing this will create a new publishment resource.
	EnvId pulumi.StringInput
	// Specifies an ID of the APIG dedicated instance to which the API belongs
	// to. Changing this will create a new publishment resource.
	InstanceId pulumi.StringInput
	// Specifies the region in which to publish APIs.
	// If omitted, the provider-level region will be used. Changing this will create a new publishment resource.
	Region pulumi.StringPtrInput
	// Specifies the version ID of the current publishment.
	VersionId pulumi.StringPtrInput
}

func (ApiPublishmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPublishmentArgs)(nil)).Elem()
}

type ApiPublishmentInput interface {
	pulumi.Input

	ToApiPublishmentOutput() ApiPublishmentOutput
	ToApiPublishmentOutputWithContext(ctx context.Context) ApiPublishmentOutput
}

func (*ApiPublishment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPublishment)(nil)).Elem()
}

func (i *ApiPublishment) ToApiPublishmentOutput() ApiPublishmentOutput {
	return i.ToApiPublishmentOutputWithContext(context.Background())
}

func (i *ApiPublishment) ToApiPublishmentOutputWithContext(ctx context.Context) ApiPublishmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPublishmentOutput)
}

// ApiPublishmentArrayInput is an input type that accepts ApiPublishmentArray and ApiPublishmentArrayOutput values.
// You can construct a concrete instance of `ApiPublishmentArrayInput` via:
//
//	ApiPublishmentArray{ ApiPublishmentArgs{...} }
type ApiPublishmentArrayInput interface {
	pulumi.Input

	ToApiPublishmentArrayOutput() ApiPublishmentArrayOutput
	ToApiPublishmentArrayOutputWithContext(context.Context) ApiPublishmentArrayOutput
}

type ApiPublishmentArray []ApiPublishmentInput

func (ApiPublishmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiPublishment)(nil)).Elem()
}

func (i ApiPublishmentArray) ToApiPublishmentArrayOutput() ApiPublishmentArrayOutput {
	return i.ToApiPublishmentArrayOutputWithContext(context.Background())
}

func (i ApiPublishmentArray) ToApiPublishmentArrayOutputWithContext(ctx context.Context) ApiPublishmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPublishmentArrayOutput)
}

// ApiPublishmentMapInput is an input type that accepts ApiPublishmentMap and ApiPublishmentMapOutput values.
// You can construct a concrete instance of `ApiPublishmentMapInput` via:
//
//	ApiPublishmentMap{ "key": ApiPublishmentArgs{...} }
type ApiPublishmentMapInput interface {
	pulumi.Input

	ToApiPublishmentMapOutput() ApiPublishmentMapOutput
	ToApiPublishmentMapOutputWithContext(context.Context) ApiPublishmentMapOutput
}

type ApiPublishmentMap map[string]ApiPublishmentInput

func (ApiPublishmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiPublishment)(nil)).Elem()
}

func (i ApiPublishmentMap) ToApiPublishmentMapOutput() ApiPublishmentMapOutput {
	return i.ToApiPublishmentMapOutputWithContext(context.Background())
}

func (i ApiPublishmentMap) ToApiPublishmentMapOutputWithContext(ctx context.Context) ApiPublishmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPublishmentMapOutput)
}

type ApiPublishmentOutput struct{ *pulumi.OutputState }

func (ApiPublishmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPublishment)(nil)).Elem()
}

func (o ApiPublishmentOutput) ToApiPublishmentOutput() ApiPublishmentOutput {
	return o
}

func (o ApiPublishmentOutput) ToApiPublishmentOutputWithContext(ctx context.Context) ApiPublishmentOutput {
	return o
}

// Specifies the API ID to be published or already published.
// Changing this will create a new publishment resource.
func (o ApiPublishmentOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// Specifies the description of the current publishment.
func (o ApiPublishmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the environment ID to which the current version of the API will be
// published or has been published. Changing this will create a new publishment resource.
func (o ApiPublishmentOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringOutput { return v.EnvId }).(pulumi.StringOutput)
}

// Environment name to which the current version of the API is published.
func (o ApiPublishmentOutput) EnvName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringOutput { return v.EnvName }).(pulumi.StringOutput)
}

// All publish informations of the API. The structure is documented below.
func (o ApiPublishmentOutput) Histories() ApiPublishmentHistoryArrayOutput {
	return o.ApplyT(func(v *ApiPublishment) ApiPublishmentHistoryArrayOutput { return v.Histories }).(ApiPublishmentHistoryArrayOutput)
}

// Specifies an ID of the APIG dedicated instance to which the API belongs
// to. Changing this will create a new publishment resource.
func (o ApiPublishmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The publish ID of the API in current environment.
func (o ApiPublishmentOutput) PublishId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringOutput { return v.PublishId }).(pulumi.StringOutput)
}

// Time when the current version was published.
func (o ApiPublishmentOutput) PublishTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringOutput { return v.PublishTime }).(pulumi.StringOutput)
}

// Specifies the region in which to publish APIs.
// If omitted, the provider-level region will be used. Changing this will create a new publishment resource.
func (o ApiPublishmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the version ID of the current publishment.
func (o ApiPublishmentOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPublishment) pulumi.StringPtrOutput { return v.VersionId }).(pulumi.StringPtrOutput)
}

type ApiPublishmentArrayOutput struct{ *pulumi.OutputState }

func (ApiPublishmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiPublishment)(nil)).Elem()
}

func (o ApiPublishmentArrayOutput) ToApiPublishmentArrayOutput() ApiPublishmentArrayOutput {
	return o
}

func (o ApiPublishmentArrayOutput) ToApiPublishmentArrayOutputWithContext(ctx context.Context) ApiPublishmentArrayOutput {
	return o
}

func (o ApiPublishmentArrayOutput) Index(i pulumi.IntInput) ApiPublishmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiPublishment {
		return vs[0].([]*ApiPublishment)[vs[1].(int)]
	}).(ApiPublishmentOutput)
}

type ApiPublishmentMapOutput struct{ *pulumi.OutputState }

func (ApiPublishmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiPublishment)(nil)).Elem()
}

func (o ApiPublishmentMapOutput) ToApiPublishmentMapOutput() ApiPublishmentMapOutput {
	return o
}

func (o ApiPublishmentMapOutput) ToApiPublishmentMapOutputWithContext(ctx context.Context) ApiPublishmentMapOutput {
	return o
}

func (o ApiPublishmentMapOutput) MapIndex(k pulumi.StringInput) ApiPublishmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiPublishment {
		return vs[0].(map[string]*ApiPublishment)[vs[1].(string)]
	}).(ApiPublishmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiPublishmentInput)(nil)).Elem(), &ApiPublishment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiPublishmentArrayInput)(nil)).Elem(), ApiPublishmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiPublishmentMapInput)(nil)).Elem(), ApiPublishmentMap{})
	pulumi.RegisterOutputType(ApiPublishmentOutput{})
	pulumi.RegisterOutputType(ApiPublishmentArrayOutput{})
	pulumi.RegisterOutputType(ApiPublishmentMapOutput{})
}