// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedapig

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a VPC channel resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedApig"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedApig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			instanceId := cfg.RequireObject("instanceId")
//			channelName := cfg.RequireObject("channelName")
//			ecsId1 := cfg.RequireObject("ecsId1")
//			ecsId2 := cfg.RequireObject("ecsId2")
//			_, err := DedicatedApig.NewVpcChannel(ctx, "test", &DedicatedApig.VpcChannelArgs{
//				InstanceId: pulumi.Any(instanceId),
//				Port:       pulumi.Int(8080),
//				Protocol:   pulumi.String("HTTPS"),
//				Path:       pulumi.String("/"),
//				HttpCode:   pulumi.String("201,202,203"),
//				Members: dedicatedapig.VpcChannelMemberArray{
//					&dedicatedapig.VpcChannelMemberArgs{
//						Id:     pulumi.Any(ecsId1),
//						Weight: pulumi.Int(30),
//					},
//					&dedicatedapig.VpcChannelMemberArgs{
//						Id:     pulumi.Any(ecsId2),
//						Weight: pulumi.Int(70),
//					},
//				},
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
// VPC Channels can be imported using their `name` and ID of the APIG dedicated instance to which the channel belongs, separated by a slash, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:DedicatedApig/vpcChannel:VpcChannel test <instance id>/<channel name>
//
// ```
type VpcChannel struct {
	pulumi.CustomResourceState

	// Specifies the type of the backend service.
	// The valid types are *WRR*, *WLC*, *SH* and *URI hashing*, default to *WRR*.
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// Time when the channel created, in UTC format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Specifies the healthy threshold, which refers to the number of consecutive
	// successful checks required for a backend server to be considered healthy.
	// The valid value is range from 2 to 10, default to 2.
	HealthyThreshold pulumi.IntPtrOutput    `pulumi:"healthyThreshold"`
	HttpCode         pulumi.StringPtrOutput `pulumi:"httpCode"`
	// Specifies an ID of the APIG dedicated instance to which the APIG
	// vpc channel belongs to.
	// Changing this will create a new VPC channel resource.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the interval between consecutive checks, in second.
	// The valid value is range from 5 to 300, default to 10.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Specifies the type of the backend service.
	// The valid types are *ECS* and *EIP*, default to *ECS*.
	MemberType pulumi.StringPtrOutput `pulumi:"memberType"`
	// Specifies an array of one or more backend server IDs or IP addresses that bind the VPC
	// channel.
	// The object structure is documented below.
	Members VpcChannelMemberArrayOutput `pulumi:"members"`
	// Specifies the name of the VPC channel.
	// The channel name consists of 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the destination path for health checks.
	// Required if `protocol` is *HTTP* or *HTTPS*.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Specifies the host port of the VPC channel.
	// The valid value is range from 1 to 65535.
	Port pulumi.IntOutput `pulumi:"port"`
	// Specifies the protocol for performing health checks on backend servers in the VPC
	// channel.
	// The valid values are *TCP*, *HTTP* and *HTTPS*, default to *TCP*.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// Specifies the region in which to create the VPC channel resource.
	// If omitted, the provider-level region will be used.
	// Changing this will create a new VPC channel resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The status of VPC channel, supports *Normal* and *Abnormal*.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the timeout for determining whether a health check fails, in second.
	// The value must be less than the value of time_interval.
	// The valid value is range from 2 to 30, default to 5.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Specifies the unhealthy threshold, which refers to the number of consecutive
	// failed checks required for a backend server to be considered unhealthy.
	// The valid value is range from 2 to 10, default to 5.
	UnhealthyThreshold pulumi.IntPtrOutput `pulumi:"unhealthyThreshold"`
}

// NewVpcChannel registers a new resource with the given unique name, arguments, and options.
func NewVpcChannel(ctx *pulumi.Context,
	name string, args *VpcChannelArgs, opts ...pulumi.ResourceOption) (*VpcChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	var resource VpcChannel
	err := ctx.RegisterResource("huaweicloud:DedicatedApig/vpcChannel:VpcChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcChannel gets an existing VpcChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcChannelState, opts ...pulumi.ResourceOption) (*VpcChannel, error) {
	var resource VpcChannel
	err := ctx.ReadResource("huaweicloud:DedicatedApig/vpcChannel:VpcChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcChannel resources.
type vpcChannelState struct {
	// Specifies the type of the backend service.
	// The valid types are *WRR*, *WLC*, *SH* and *URI hashing*, default to *WRR*.
	Algorithm *string `pulumi:"algorithm"`
	// Time when the channel created, in UTC format.
	CreateTime *string `pulumi:"createTime"`
	// Specifies the healthy threshold, which refers to the number of consecutive
	// successful checks required for a backend server to be considered healthy.
	// The valid value is range from 2 to 10, default to 2.
	HealthyThreshold *int    `pulumi:"healthyThreshold"`
	HttpCode         *string `pulumi:"httpCode"`
	// Specifies an ID of the APIG dedicated instance to which the APIG
	// vpc channel belongs to.
	// Changing this will create a new VPC channel resource.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the interval between consecutive checks, in second.
	// The valid value is range from 5 to 300, default to 10.
	Interval *int `pulumi:"interval"`
	// Specifies the type of the backend service.
	// The valid types are *ECS* and *EIP*, default to *ECS*.
	MemberType *string `pulumi:"memberType"`
	// Specifies an array of one or more backend server IDs or IP addresses that bind the VPC
	// channel.
	// The object structure is documented below.
	Members []VpcChannelMember `pulumi:"members"`
	// Specifies the name of the VPC channel.
	// The channel name consists of 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name *string `pulumi:"name"`
	// Specifies the destination path for health checks.
	// Required if `protocol` is *HTTP* or *HTTPS*.
	Path *string `pulumi:"path"`
	// Specifies the host port of the VPC channel.
	// The valid value is range from 1 to 65535.
	Port *int `pulumi:"port"`
	// Specifies the protocol for performing health checks on backend servers in the VPC
	// channel.
	// The valid values are *TCP*, *HTTP* and *HTTPS*, default to *TCP*.
	Protocol *string `pulumi:"protocol"`
	// Specifies the region in which to create the VPC channel resource.
	// If omitted, the provider-level region will be used.
	// Changing this will create a new VPC channel resource.
	Region *string `pulumi:"region"`
	// The status of VPC channel, supports *Normal* and *Abnormal*.
	Status *string `pulumi:"status"`
	// Specifies the timeout for determining whether a health check fails, in second.
	// The value must be less than the value of time_interval.
	// The valid value is range from 2 to 30, default to 5.
	Timeout *int `pulumi:"timeout"`
	// Specifies the unhealthy threshold, which refers to the number of consecutive
	// failed checks required for a backend server to be considered unhealthy.
	// The valid value is range from 2 to 10, default to 5.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

type VpcChannelState struct {
	// Specifies the type of the backend service.
	// The valid types are *WRR*, *WLC*, *SH* and *URI hashing*, default to *WRR*.
	Algorithm pulumi.StringPtrInput
	// Time when the channel created, in UTC format.
	CreateTime pulumi.StringPtrInput
	// Specifies the healthy threshold, which refers to the number of consecutive
	// successful checks required for a backend server to be considered healthy.
	// The valid value is range from 2 to 10, default to 2.
	HealthyThreshold pulumi.IntPtrInput
	HttpCode         pulumi.StringPtrInput
	// Specifies an ID of the APIG dedicated instance to which the APIG
	// vpc channel belongs to.
	// Changing this will create a new VPC channel resource.
	InstanceId pulumi.StringPtrInput
	// Specifies the interval between consecutive checks, in second.
	// The valid value is range from 5 to 300, default to 10.
	Interval pulumi.IntPtrInput
	// Specifies the type of the backend service.
	// The valid types are *ECS* and *EIP*, default to *ECS*.
	MemberType pulumi.StringPtrInput
	// Specifies an array of one or more backend server IDs or IP addresses that bind the VPC
	// channel.
	// The object structure is documented below.
	Members VpcChannelMemberArrayInput
	// Specifies the name of the VPC channel.
	// The channel name consists of 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name pulumi.StringPtrInput
	// Specifies the destination path for health checks.
	// Required if `protocol` is *HTTP* or *HTTPS*.
	Path pulumi.StringPtrInput
	// Specifies the host port of the VPC channel.
	// The valid value is range from 1 to 65535.
	Port pulumi.IntPtrInput
	// Specifies the protocol for performing health checks on backend servers in the VPC
	// channel.
	// The valid values are *TCP*, *HTTP* and *HTTPS*, default to *TCP*.
	Protocol pulumi.StringPtrInput
	// Specifies the region in which to create the VPC channel resource.
	// If omitted, the provider-level region will be used.
	// Changing this will create a new VPC channel resource.
	Region pulumi.StringPtrInput
	// The status of VPC channel, supports *Normal* and *Abnormal*.
	Status pulumi.StringPtrInput
	// Specifies the timeout for determining whether a health check fails, in second.
	// The value must be less than the value of time_interval.
	// The valid value is range from 2 to 30, default to 5.
	Timeout pulumi.IntPtrInput
	// Specifies the unhealthy threshold, which refers to the number of consecutive
	// failed checks required for a backend server to be considered unhealthy.
	// The valid value is range from 2 to 10, default to 5.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (VpcChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcChannelState)(nil)).Elem()
}

type vpcChannelArgs struct {
	// Specifies the type of the backend service.
	// The valid types are *WRR*, *WLC*, *SH* and *URI hashing*, default to *WRR*.
	Algorithm *string `pulumi:"algorithm"`
	// Specifies the healthy threshold, which refers to the number of consecutive
	// successful checks required for a backend server to be considered healthy.
	// The valid value is range from 2 to 10, default to 2.
	HealthyThreshold *int    `pulumi:"healthyThreshold"`
	HttpCode         *string `pulumi:"httpCode"`
	// Specifies an ID of the APIG dedicated instance to which the APIG
	// vpc channel belongs to.
	// Changing this will create a new VPC channel resource.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the interval between consecutive checks, in second.
	// The valid value is range from 5 to 300, default to 10.
	Interval *int `pulumi:"interval"`
	// Specifies the type of the backend service.
	// The valid types are *ECS* and *EIP*, default to *ECS*.
	MemberType *string `pulumi:"memberType"`
	// Specifies an array of one or more backend server IDs or IP addresses that bind the VPC
	// channel.
	// The object structure is documented below.
	Members []VpcChannelMember `pulumi:"members"`
	// Specifies the name of the VPC channel.
	// The channel name consists of 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name *string `pulumi:"name"`
	// Specifies the destination path for health checks.
	// Required if `protocol` is *HTTP* or *HTTPS*.
	Path *string `pulumi:"path"`
	// Specifies the host port of the VPC channel.
	// The valid value is range from 1 to 65535.
	Port int `pulumi:"port"`
	// Specifies the protocol for performing health checks on backend servers in the VPC
	// channel.
	// The valid values are *TCP*, *HTTP* and *HTTPS*, default to *TCP*.
	Protocol *string `pulumi:"protocol"`
	// Specifies the region in which to create the VPC channel resource.
	// If omitted, the provider-level region will be used.
	// Changing this will create a new VPC channel resource.
	Region *string `pulumi:"region"`
	// Specifies the timeout for determining whether a health check fails, in second.
	// The value must be less than the value of time_interval.
	// The valid value is range from 2 to 30, default to 5.
	Timeout *int `pulumi:"timeout"`
	// Specifies the unhealthy threshold, which refers to the number of consecutive
	// failed checks required for a backend server to be considered unhealthy.
	// The valid value is range from 2 to 10, default to 5.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

// The set of arguments for constructing a VpcChannel resource.
type VpcChannelArgs struct {
	// Specifies the type of the backend service.
	// The valid types are *WRR*, *WLC*, *SH* and *URI hashing*, default to *WRR*.
	Algorithm pulumi.StringPtrInput
	// Specifies the healthy threshold, which refers to the number of consecutive
	// successful checks required for a backend server to be considered healthy.
	// The valid value is range from 2 to 10, default to 2.
	HealthyThreshold pulumi.IntPtrInput
	HttpCode         pulumi.StringPtrInput
	// Specifies an ID of the APIG dedicated instance to which the APIG
	// vpc channel belongs to.
	// Changing this will create a new VPC channel resource.
	InstanceId pulumi.StringInput
	// Specifies the interval between consecutive checks, in second.
	// The valid value is range from 5 to 300, default to 10.
	Interval pulumi.IntPtrInput
	// Specifies the type of the backend service.
	// The valid types are *ECS* and *EIP*, default to *ECS*.
	MemberType pulumi.StringPtrInput
	// Specifies an array of one or more backend server IDs or IP addresses that bind the VPC
	// channel.
	// The object structure is documented below.
	Members VpcChannelMemberArrayInput
	// Specifies the name of the VPC channel.
	// The channel name consists of 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name pulumi.StringPtrInput
	// Specifies the destination path for health checks.
	// Required if `protocol` is *HTTP* or *HTTPS*.
	Path pulumi.StringPtrInput
	// Specifies the host port of the VPC channel.
	// The valid value is range from 1 to 65535.
	Port pulumi.IntInput
	// Specifies the protocol for performing health checks on backend servers in the VPC
	// channel.
	// The valid values are *TCP*, *HTTP* and *HTTPS*, default to *TCP*.
	Protocol pulumi.StringPtrInput
	// Specifies the region in which to create the VPC channel resource.
	// If omitted, the provider-level region will be used.
	// Changing this will create a new VPC channel resource.
	Region pulumi.StringPtrInput
	// Specifies the timeout for determining whether a health check fails, in second.
	// The value must be less than the value of time_interval.
	// The valid value is range from 2 to 30, default to 5.
	Timeout pulumi.IntPtrInput
	// Specifies the unhealthy threshold, which refers to the number of consecutive
	// failed checks required for a backend server to be considered unhealthy.
	// The valid value is range from 2 to 10, default to 5.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (VpcChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcChannelArgs)(nil)).Elem()
}

type VpcChannelInput interface {
	pulumi.Input

	ToVpcChannelOutput() VpcChannelOutput
	ToVpcChannelOutputWithContext(ctx context.Context) VpcChannelOutput
}

func (*VpcChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcChannel)(nil)).Elem()
}

func (i *VpcChannel) ToVpcChannelOutput() VpcChannelOutput {
	return i.ToVpcChannelOutputWithContext(context.Background())
}

func (i *VpcChannel) ToVpcChannelOutputWithContext(ctx context.Context) VpcChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcChannelOutput)
}

// VpcChannelArrayInput is an input type that accepts VpcChannelArray and VpcChannelArrayOutput values.
// You can construct a concrete instance of `VpcChannelArrayInput` via:
//
//	VpcChannelArray{ VpcChannelArgs{...} }
type VpcChannelArrayInput interface {
	pulumi.Input

	ToVpcChannelArrayOutput() VpcChannelArrayOutput
	ToVpcChannelArrayOutputWithContext(context.Context) VpcChannelArrayOutput
}

type VpcChannelArray []VpcChannelInput

func (VpcChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcChannel)(nil)).Elem()
}

func (i VpcChannelArray) ToVpcChannelArrayOutput() VpcChannelArrayOutput {
	return i.ToVpcChannelArrayOutputWithContext(context.Background())
}

func (i VpcChannelArray) ToVpcChannelArrayOutputWithContext(ctx context.Context) VpcChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcChannelArrayOutput)
}

// VpcChannelMapInput is an input type that accepts VpcChannelMap and VpcChannelMapOutput values.
// You can construct a concrete instance of `VpcChannelMapInput` via:
//
//	VpcChannelMap{ "key": VpcChannelArgs{...} }
type VpcChannelMapInput interface {
	pulumi.Input

	ToVpcChannelMapOutput() VpcChannelMapOutput
	ToVpcChannelMapOutputWithContext(context.Context) VpcChannelMapOutput
}

type VpcChannelMap map[string]VpcChannelInput

func (VpcChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcChannel)(nil)).Elem()
}

func (i VpcChannelMap) ToVpcChannelMapOutput() VpcChannelMapOutput {
	return i.ToVpcChannelMapOutputWithContext(context.Background())
}

func (i VpcChannelMap) ToVpcChannelMapOutputWithContext(ctx context.Context) VpcChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcChannelMapOutput)
}

type VpcChannelOutput struct{ *pulumi.OutputState }

func (VpcChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcChannel)(nil)).Elem()
}

func (o VpcChannelOutput) ToVpcChannelOutput() VpcChannelOutput {
	return o
}

func (o VpcChannelOutput) ToVpcChannelOutputWithContext(ctx context.Context) VpcChannelOutput {
	return o
}

// Specifies the type of the backend service.
// The valid types are *WRR*, *WLC*, *SH* and *URI hashing*, default to *WRR*.
func (o VpcChannelOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringPtrOutput { return v.Algorithm }).(pulumi.StringPtrOutput)
}

// Time when the channel created, in UTC format.
func (o VpcChannelOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Specifies the healthy threshold, which refers to the number of consecutive
// successful checks required for a backend server to be considered healthy.
// The valid value is range from 2 to 10, default to 2.
func (o VpcChannelOutput) HealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.IntPtrOutput { return v.HealthyThreshold }).(pulumi.IntPtrOutput)
}

func (o VpcChannelOutput) HttpCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringPtrOutput { return v.HttpCode }).(pulumi.StringPtrOutput)
}

// Specifies an ID of the APIG dedicated instance to which the APIG
// vpc channel belongs to.
// Changing this will create a new VPC channel resource.
func (o VpcChannelOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the interval between consecutive checks, in second.
// The valid value is range from 5 to 300, default to 10.
func (o VpcChannelOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// Specifies the type of the backend service.
// The valid types are *ECS* and *EIP*, default to *ECS*.
func (o VpcChannelOutput) MemberType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringPtrOutput { return v.MemberType }).(pulumi.StringPtrOutput)
}

// Specifies an array of one or more backend server IDs or IP addresses that bind the VPC
// channel.
// The object structure is documented below.
func (o VpcChannelOutput) Members() VpcChannelMemberArrayOutput {
	return o.ApplyT(func(v *VpcChannel) VpcChannelMemberArrayOutput { return v.Members }).(VpcChannelMemberArrayOutput)
}

// Specifies the name of the VPC channel.
// The channel name consists of 3 to 64 characters, starting with a letter.
// Only letters, digits and underscores (_) are allowed.
// Chinese characters must be in UTF-8 or Unicode format.
func (o VpcChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the destination path for health checks.
// Required if `protocol` is *HTTP* or *HTTPS*.
func (o VpcChannelOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Specifies the host port of the VPC channel.
// The valid value is range from 1 to 65535.
func (o VpcChannelOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Specifies the protocol for performing health checks on backend servers in the VPC
// channel.
// The valid values are *TCP*, *HTTP* and *HTTPS*, default to *TCP*.
func (o VpcChannelOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// Specifies the region in which to create the VPC channel resource.
// If omitted, the provider-level region will be used.
// Changing this will create a new VPC channel resource.
func (o VpcChannelOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The status of VPC channel, supports *Normal* and *Abnormal*.
func (o VpcChannelOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the timeout for determining whether a health check fails, in second.
// The value must be less than the value of time_interval.
// The valid value is range from 2 to 30, default to 5.
func (o VpcChannelOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// Specifies the unhealthy threshold, which refers to the number of consecutive
// failed checks required for a backend server to be considered unhealthy.
// The valid value is range from 2 to 10, default to 5.
func (o VpcChannelOutput) UnhealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcChannel) pulumi.IntPtrOutput { return v.UnhealthyThreshold }).(pulumi.IntPtrOutput)
}

type VpcChannelArrayOutput struct{ *pulumi.OutputState }

func (VpcChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcChannel)(nil)).Elem()
}

func (o VpcChannelArrayOutput) ToVpcChannelArrayOutput() VpcChannelArrayOutput {
	return o
}

func (o VpcChannelArrayOutput) ToVpcChannelArrayOutputWithContext(ctx context.Context) VpcChannelArrayOutput {
	return o
}

func (o VpcChannelArrayOutput) Index(i pulumi.IntInput) VpcChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcChannel {
		return vs[0].([]*VpcChannel)[vs[1].(int)]
	}).(VpcChannelOutput)
}

type VpcChannelMapOutput struct{ *pulumi.OutputState }

func (VpcChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcChannel)(nil)).Elem()
}

func (o VpcChannelMapOutput) ToVpcChannelMapOutput() VpcChannelMapOutput {
	return o
}

func (o VpcChannelMapOutput) ToVpcChannelMapOutputWithContext(ctx context.Context) VpcChannelMapOutput {
	return o
}

func (o VpcChannelMapOutput) MapIndex(k pulumi.StringInput) VpcChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcChannel {
		return vs[0].(map[string]*VpcChannel)[vs[1].(string)]
	}).(VpcChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcChannelInput)(nil)).Elem(), &VpcChannel{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcChannelArrayInput)(nil)).Elem(), VpcChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcChannelMapInput)(nil)).Elem(), VpcChannelMap{})
	pulumi.RegisterOutputType(VpcChannelOutput{})
	pulumi.RegisterOutputType(VpcChannelArrayOutput{})
	pulumi.RegisterOutputType(VpcChannelMapOutput{})
}