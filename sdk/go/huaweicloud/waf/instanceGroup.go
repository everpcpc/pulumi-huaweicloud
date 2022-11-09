// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages WAF instance groups within HuaweiCloud. The groups are used to bind the ELB instance to the ELB mode WAF.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vpcId := cfg.RequireObject("vpcId")
//			_, err := Waf.NewInstanceGroup(ctx, "group1", &Waf.InstanceGroupArgs{
//				VpcId: pulumi.Any(vpcId),
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
// The instance group can be imported using the ID, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/instanceGroup:InstanceGroup group_1 0be1e69d-1987-4d9c-9dc5-fc7eed592398
//
// ```
type InstanceGroup struct {
	pulumi.CustomResourceState

	// The body limit of the forwarding policy.
	BodyLimit pulumi.IntOutput `pulumi:"bodyLimit"`
	// The time for connection timeout in the forwarding policy.
	ConnectionTimeout pulumi.IntOutput `pulumi:"connectionTimeout"`
	// Specifies the description of the instance group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The header limit of the forwarding policy.
	HeaderLimit pulumi.IntOutput `pulumi:"headerLimit"`
	// The IDs of the ELB instances that has been bound to the instance group.
	LoadBalancers pulumi.StringArrayOutput `pulumi:"loadBalancers"`
	// Specifies the instance group name.
	// The maximum length is 64 characters. Only letters, digits and underscores (_) are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The time for reading timeout in the forwarding policy.
	ReadTimeout pulumi.IntOutput `pulumi:"readTimeout"`
	// The region in which to create the instance group.
	// If omitted, the provider-level region will be used. Changing this setting will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the id of the VPC that the WAF dedicated instances belongs to.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The time for writing timeout in the forwarding policy.
	WriteTimeout pulumi.IntOutput `pulumi:"writeTimeout"`
}

// NewInstanceGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroup(ctx *pulumi.Context,
	name string, args *InstanceGroupArgs, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceGroup
	err := ctx.RegisterResource("huaweicloud:Waf/instanceGroup:InstanceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroup gets an existing InstanceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupState, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	var resource InstanceGroup
	err := ctx.ReadResource("huaweicloud:Waf/instanceGroup:InstanceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroup resources.
type instanceGroupState struct {
	// The body limit of the forwarding policy.
	BodyLimit *int `pulumi:"bodyLimit"`
	// The time for connection timeout in the forwarding policy.
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Specifies the description of the instance group.
	Description *string `pulumi:"description"`
	// The header limit of the forwarding policy.
	HeaderLimit *int `pulumi:"headerLimit"`
	// The IDs of the ELB instances that has been bound to the instance group.
	LoadBalancers []string `pulumi:"loadBalancers"`
	// Specifies the instance group name.
	// The maximum length is 64 characters. Only letters, digits and underscores (_) are allowed.
	Name *string `pulumi:"name"`
	// The time for reading timeout in the forwarding policy.
	ReadTimeout *int `pulumi:"readTimeout"`
	// The region in which to create the instance group.
	// If omitted, the provider-level region will be used. Changing this setting will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the id of the VPC that the WAF dedicated instances belongs to.
	VpcId *string `pulumi:"vpcId"`
	// The time for writing timeout in the forwarding policy.
	WriteTimeout *int `pulumi:"writeTimeout"`
}

type InstanceGroupState struct {
	// The body limit of the forwarding policy.
	BodyLimit pulumi.IntPtrInput
	// The time for connection timeout in the forwarding policy.
	ConnectionTimeout pulumi.IntPtrInput
	// Specifies the description of the instance group.
	Description pulumi.StringPtrInput
	// The header limit of the forwarding policy.
	HeaderLimit pulumi.IntPtrInput
	// The IDs of the ELB instances that has been bound to the instance group.
	LoadBalancers pulumi.StringArrayInput
	// Specifies the instance group name.
	// The maximum length is 64 characters. Only letters, digits and underscores (_) are allowed.
	Name pulumi.StringPtrInput
	// The time for reading timeout in the forwarding policy.
	ReadTimeout pulumi.IntPtrInput
	// The region in which to create the instance group.
	// If omitted, the provider-level region will be used. Changing this setting will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the id of the VPC that the WAF dedicated instances belongs to.
	VpcId pulumi.StringPtrInput
	// The time for writing timeout in the forwarding policy.
	WriteTimeout pulumi.IntPtrInput
}

func (InstanceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupState)(nil)).Elem()
}

type instanceGroupArgs struct {
	// Specifies the description of the instance group.
	Description *string `pulumi:"description"`
	// Specifies the instance group name.
	// The maximum length is 64 characters. Only letters, digits and underscores (_) are allowed.
	Name *string `pulumi:"name"`
	// The region in which to create the instance group.
	// If omitted, the provider-level region will be used. Changing this setting will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the id of the VPC that the WAF dedicated instances belongs to.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a InstanceGroup resource.
type InstanceGroupArgs struct {
	// Specifies the description of the instance group.
	Description pulumi.StringPtrInput
	// Specifies the instance group name.
	// The maximum length is 64 characters. Only letters, digits and underscores (_) are allowed.
	Name pulumi.StringPtrInput
	// The region in which to create the instance group.
	// If omitted, the provider-level region will be used. Changing this setting will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the id of the VPC that the WAF dedicated instances belongs to.
	VpcId pulumi.StringInput
}

func (InstanceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupArgs)(nil)).Elem()
}

type InstanceGroupInput interface {
	pulumi.Input

	ToInstanceGroupOutput() InstanceGroupOutput
	ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput
}

func (*InstanceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroup)(nil)).Elem()
}

func (i *InstanceGroup) ToInstanceGroupOutput() InstanceGroupOutput {
	return i.ToInstanceGroupOutputWithContext(context.Background())
}

func (i *InstanceGroup) ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupOutput)
}

// InstanceGroupArrayInput is an input type that accepts InstanceGroupArray and InstanceGroupArrayOutput values.
// You can construct a concrete instance of `InstanceGroupArrayInput` via:
//
//	InstanceGroupArray{ InstanceGroupArgs{...} }
type InstanceGroupArrayInput interface {
	pulumi.Input

	ToInstanceGroupArrayOutput() InstanceGroupArrayOutput
	ToInstanceGroupArrayOutputWithContext(context.Context) InstanceGroupArrayOutput
}

type InstanceGroupArray []InstanceGroupInput

func (InstanceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGroup)(nil)).Elem()
}

func (i InstanceGroupArray) ToInstanceGroupArrayOutput() InstanceGroupArrayOutput {
	return i.ToInstanceGroupArrayOutputWithContext(context.Background())
}

func (i InstanceGroupArray) ToInstanceGroupArrayOutputWithContext(ctx context.Context) InstanceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupArrayOutput)
}

// InstanceGroupMapInput is an input type that accepts InstanceGroupMap and InstanceGroupMapOutput values.
// You can construct a concrete instance of `InstanceGroupMapInput` via:
//
//	InstanceGroupMap{ "key": InstanceGroupArgs{...} }
type InstanceGroupMapInput interface {
	pulumi.Input

	ToInstanceGroupMapOutput() InstanceGroupMapOutput
	ToInstanceGroupMapOutputWithContext(context.Context) InstanceGroupMapOutput
}

type InstanceGroupMap map[string]InstanceGroupInput

func (InstanceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGroup)(nil)).Elem()
}

func (i InstanceGroupMap) ToInstanceGroupMapOutput() InstanceGroupMapOutput {
	return i.ToInstanceGroupMapOutputWithContext(context.Background())
}

func (i InstanceGroupMap) ToInstanceGroupMapOutputWithContext(ctx context.Context) InstanceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupMapOutput)
}

type InstanceGroupOutput struct{ *pulumi.OutputState }

func (InstanceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroup)(nil)).Elem()
}

func (o InstanceGroupOutput) ToInstanceGroupOutput() InstanceGroupOutput {
	return o
}

func (o InstanceGroupOutput) ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput {
	return o
}

// The body limit of the forwarding policy.
func (o InstanceGroupOutput) BodyLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.IntOutput { return v.BodyLimit }).(pulumi.IntOutput)
}

// The time for connection timeout in the forwarding policy.
func (o InstanceGroupOutput) ConnectionTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.IntOutput { return v.ConnectionTimeout }).(pulumi.IntOutput)
}

// Specifies the description of the instance group.
func (o InstanceGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The header limit of the forwarding policy.
func (o InstanceGroupOutput) HeaderLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.IntOutput { return v.HeaderLimit }).(pulumi.IntOutput)
}

// The IDs of the ELB instances that has been bound to the instance group.
func (o InstanceGroupOutput) LoadBalancers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.StringArrayOutput { return v.LoadBalancers }).(pulumi.StringArrayOutput)
}

// Specifies the instance group name.
// The maximum length is 64 characters. Only letters, digits and underscores (_) are allowed.
func (o InstanceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The time for reading timeout in the forwarding policy.
func (o InstanceGroupOutput) ReadTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.IntOutput { return v.ReadTimeout }).(pulumi.IntOutput)
}

// The region in which to create the instance group.
// If omitted, the provider-level region will be used. Changing this setting will create a new resource.
func (o InstanceGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the id of the VPC that the WAF dedicated instances belongs to.
func (o InstanceGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The time for writing timeout in the forwarding policy.
func (o InstanceGroupOutput) WriteTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroup) pulumi.IntOutput { return v.WriteTimeout }).(pulumi.IntOutput)
}

type InstanceGroupArrayOutput struct{ *pulumi.OutputState }

func (InstanceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGroup)(nil)).Elem()
}

func (o InstanceGroupArrayOutput) ToInstanceGroupArrayOutput() InstanceGroupArrayOutput {
	return o
}

func (o InstanceGroupArrayOutput) ToInstanceGroupArrayOutputWithContext(ctx context.Context) InstanceGroupArrayOutput {
	return o
}

func (o InstanceGroupArrayOutput) Index(i pulumi.IntInput) InstanceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceGroup {
		return vs[0].([]*InstanceGroup)[vs[1].(int)]
	}).(InstanceGroupOutput)
}

type InstanceGroupMapOutput struct{ *pulumi.OutputState }

func (InstanceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGroup)(nil)).Elem()
}

func (o InstanceGroupMapOutput) ToInstanceGroupMapOutput() InstanceGroupMapOutput {
	return o
}

func (o InstanceGroupMapOutput) ToInstanceGroupMapOutputWithContext(ctx context.Context) InstanceGroupMapOutput {
	return o
}

func (o InstanceGroupMapOutput) MapIndex(k pulumi.StringInput) InstanceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceGroup {
		return vs[0].(map[string]*InstanceGroup)[vs[1].(string)]
	}).(InstanceGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupInput)(nil)).Elem(), &InstanceGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupArrayInput)(nil)).Elem(), InstanceGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupMapInput)(nil)).Elem(), InstanceGroupMap{})
	pulumi.RegisterOutputType(InstanceGroupOutput{})
	pulumi.RegisterOutputType(InstanceGroupArrayOutput{})
	pulumi.RegisterOutputType(InstanceGroupMapOutput{})
}
