// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package swr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SWR repository resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Swr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			organizationName := cfg.RequireObject("organizationName")
//			_, err := Swr.NewRepository(ctx, "test", &Swr.RepositoryArgs{
//				Organization: pulumi.Any(organizationName),
//				Description:  pulumi.String("Test repository"),
//				Category:     pulumi.String("linux"),
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
// Repository can be imported using the organization name and repository name separated by a slash, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Swr/repository:Repository test org-name/repo-name
//
// ```
type Repository struct {
	pulumi.CustomResourceState

	// Specifies the category of the repository.
	// The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// Specifies the description of the repository.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Intra-cluster image address for docker pull.
	InternalPath pulumi.StringOutput `pulumi:"internalPath"`
	// Specifies whether the repository is public. Default is false.
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// Specifies the name of the repository. Changing this creates a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of image tags in a repository.
	NumImages pulumi.IntOutput `pulumi:"numImages"`
	// Specifies the name of the organization (namespace) the repository belongs.
	// Changing this creates a new resource.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Image address for docker pull.
	Path pulumi.StringOutput `pulumi:"path"`
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Numeric ID of the repository
	RepositoryId pulumi.IntOutput `pulumi:"repositoryId"`
	// Repository size.
	Size pulumi.IntOutput `pulumi:"size"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	var resource Repository
	err := ctx.RegisterResource("huaweicloud:Swr/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("huaweicloud:Swr/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Specifies the category of the repository.
	// The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
	Category *string `pulumi:"category"`
	// Specifies the description of the repository.
	Description *string `pulumi:"description"`
	// Intra-cluster image address for docker pull.
	InternalPath *string `pulumi:"internalPath"`
	// Specifies whether the repository is public. Default is false.
	IsPublic *bool `pulumi:"isPublic"`
	// Specifies the name of the repository. Changing this creates a new resource.
	Name *string `pulumi:"name"`
	// Number of image tags in a repository.
	NumImages *int `pulumi:"numImages"`
	// Specifies the name of the organization (namespace) the repository belongs.
	// Changing this creates a new resource.
	Organization *string `pulumi:"organization"`
	// Image address for docker pull.
	Path *string `pulumi:"path"`
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Numeric ID of the repository
	RepositoryId *int `pulumi:"repositoryId"`
	// Repository size.
	Size *int `pulumi:"size"`
}

type RepositoryState struct {
	// Specifies the category of the repository.
	// The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
	Category pulumi.StringPtrInput
	// Specifies the description of the repository.
	Description pulumi.StringPtrInput
	// Intra-cluster image address for docker pull.
	InternalPath pulumi.StringPtrInput
	// Specifies whether the repository is public. Default is false.
	IsPublic pulumi.BoolPtrInput
	// Specifies the name of the repository. Changing this creates a new resource.
	Name pulumi.StringPtrInput
	// Number of image tags in a repository.
	NumImages pulumi.IntPtrInput
	// Specifies the name of the organization (namespace) the repository belongs.
	// Changing this creates a new resource.
	Organization pulumi.StringPtrInput
	// Image address for docker pull.
	Path pulumi.StringPtrInput
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Numeric ID of the repository
	RepositoryId pulumi.IntPtrInput
	// Repository size.
	Size pulumi.IntPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Specifies the category of the repository.
	// The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
	Category *string `pulumi:"category"`
	// Specifies the description of the repository.
	Description *string `pulumi:"description"`
	// Specifies whether the repository is public. Default is false.
	IsPublic *bool `pulumi:"isPublic"`
	// Specifies the name of the repository. Changing this creates a new resource.
	Name *string `pulumi:"name"`
	// Specifies the name of the organization (namespace) the repository belongs.
	// Changing this creates a new resource.
	Organization string `pulumi:"organization"`
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Specifies the category of the repository.
	// The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
	Category pulumi.StringPtrInput
	// Specifies the description of the repository.
	Description pulumi.StringPtrInput
	// Specifies whether the repository is public. Default is false.
	IsPublic pulumi.BoolPtrInput
	// Specifies the name of the repository. Changing this creates a new resource.
	Name pulumi.StringPtrInput
	// Specifies the name of the organization (namespace) the repository belongs.
	// Changing this creates a new resource.
	Organization pulumi.StringInput
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//	RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//	RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// Specifies the category of the repository.
// The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
func (o RepositoryOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

// Specifies the description of the repository.
func (o RepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Intra-cluster image address for docker pull.
func (o RepositoryOutput) InternalPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.InternalPath }).(pulumi.StringOutput)
}

// Specifies whether the repository is public. Default is false.
func (o RepositoryOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// Specifies the name of the repository. Changing this creates a new resource.
func (o RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of image tags in a repository.
func (o RepositoryOutput) NumImages() pulumi.IntOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntOutput { return v.NumImages }).(pulumi.IntOutput)
}

// Specifies the name of the organization (namespace) the repository belongs.
// Changing this creates a new resource.
func (o RepositoryOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// Image address for docker pull.
func (o RepositoryOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Specifies the region in which to create the resource. If omitted, the
// provider-level region will be used. Changing this creates a new resource.
func (o RepositoryOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Numeric ID of the repository
func (o RepositoryOutput) RepositoryId() pulumi.IntOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntOutput { return v.RepositoryId }).(pulumi.IntOutput)
}

// Repository size.
func (o RepositoryOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].([]*Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].(map[string]*Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryArrayInput)(nil)).Elem(), RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMapInput)(nil)).Elem(), RepositoryMap{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}