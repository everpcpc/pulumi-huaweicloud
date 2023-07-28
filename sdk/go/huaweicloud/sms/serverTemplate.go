// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an SMS server template resource within HuaweiCloud.
//
// ## Example Usage
// ### A template will create networks during migration
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud"
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Sms"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			demoAvailabilityZones, err := huaweicloud.GetAvailabilityZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Sms.NewServerTemplate(ctx, "demoServerTemplate", &Sms.ServerTemplateArgs{
//				AvailabilityZone: pulumi.String(demoAvailabilityZones.Names[0]),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### A template will use the existing networks during migration
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud"
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Sms"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vpcId := cfg.RequireObject("vpcId")
//			subentId := cfg.RequireObject("subentId")
//			secgroupId := cfg.RequireObject("secgroupId")
//			demoAvailabilityZones, err := huaweicloud.GetAvailabilityZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Sms.NewServerTemplate(ctx, "demoServerTemplate", &Sms.ServerTemplateArgs{
//				AvailabilityZone: pulumi.String(demoAvailabilityZones.Names[0]),
//				VpcId:            pulumi.Any(vpcId),
//				SubnetIds: pulumi.StringArray{
//					pulumi.Any(subentId),
//				},
//				SecurityGroupIds: pulumi.StringArray{
//					pulumi.Any(secgroupId),
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
// SMS server templates can be imported by `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Sms/serverTemplate:ServerTemplate demo 4618ccaf-b4d7-43b9-b958-3df3b885126d
//
// ```
type ServerTemplate struct {
	pulumi.CustomResourceState

	// Specifies the availability zone where the target server is located.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Specifies the bandwidth size in Mbit/s about the public IP address
	// that will be used for migration.
	BandwidthSize pulumi.IntPtrOutput `pulumi:"bandwidthSize"`
	// Specifies the flavor ID for the target server.
	Flavor pulumi.StringPtrOutput `pulumi:"flavor"`
	// Specifies the server template name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the project ID where the target server is located.
	// If omitted, the default project in the region will be used.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Specifies the region where the target server is located.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies an array of one or more security group IDs to associate with
	// the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Specifies an array of one or more subnet IDs to attach to the target server.
	// If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Specifies the name of the target server. Defaults to the template name.
	TargetServerName pulumi.StringOutput `pulumi:"targetServerName"`
	// Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
	// defaults to **SAS**.
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
	// Specifies the ID of the VPC which the target server belongs to.
	// If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The name of the VPC which the target server belongs to.
	VpcName pulumi.StringOutput `pulumi:"vpcName"`
}

// NewServerTemplate registers a new resource with the given unique name, arguments, and options.
func NewServerTemplate(ctx *pulumi.Context,
	name string, args *ServerTemplateArgs, opts ...pulumi.ResourceOption) (*ServerTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServerTemplate
	err := ctx.RegisterResource("huaweicloud:Sms/serverTemplate:ServerTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerTemplate gets an existing ServerTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerTemplateState, opts ...pulumi.ResourceOption) (*ServerTemplate, error) {
	var resource ServerTemplate
	err := ctx.ReadResource("huaweicloud:Sms/serverTemplate:ServerTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerTemplate resources.
type serverTemplateState struct {
	// Specifies the availability zone where the target server is located.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the bandwidth size in Mbit/s about the public IP address
	// that will be used for migration.
	BandwidthSize *int `pulumi:"bandwidthSize"`
	// Specifies the flavor ID for the target server.
	Flavor *string `pulumi:"flavor"`
	// Specifies the server template name.
	Name *string `pulumi:"name"`
	// Specifies the project ID where the target server is located.
	// If omitted, the default project in the region will be used.
	ProjectId *string `pulumi:"projectId"`
	// Specifies the region where the target server is located.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifies an array of one or more security group IDs to associate with
	// the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Specifies an array of one or more subnet IDs to attach to the target server.
	// If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
	SubnetIds []string `pulumi:"subnetIds"`
	// Specifies the name of the target server. Defaults to the template name.
	TargetServerName *string `pulumi:"targetServerName"`
	// Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
	// defaults to **SAS**.
	VolumeType *string `pulumi:"volumeType"`
	// Specifies the ID of the VPC which the target server belongs to.
	// If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
	VpcId *string `pulumi:"vpcId"`
	// The name of the VPC which the target server belongs to.
	VpcName *string `pulumi:"vpcName"`
}

type ServerTemplateState struct {
	// Specifies the availability zone where the target server is located.
	AvailabilityZone pulumi.StringPtrInput
	// Specifies the bandwidth size in Mbit/s about the public IP address
	// that will be used for migration.
	BandwidthSize pulumi.IntPtrInput
	// Specifies the flavor ID for the target server.
	Flavor pulumi.StringPtrInput
	// Specifies the server template name.
	Name pulumi.StringPtrInput
	// Specifies the project ID where the target server is located.
	// If omitted, the default project in the region will be used.
	ProjectId pulumi.StringPtrInput
	// Specifies the region where the target server is located.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput
	// Specifies an array of one or more security group IDs to associate with
	// the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
	SecurityGroupIds pulumi.StringArrayInput
	// Specifies an array of one or more subnet IDs to attach to the target server.
	// If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
	SubnetIds pulumi.StringArrayInput
	// Specifies the name of the target server. Defaults to the template name.
	TargetServerName pulumi.StringPtrInput
	// Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
	// defaults to **SAS**.
	VolumeType pulumi.StringPtrInput
	// Specifies the ID of the VPC which the target server belongs to.
	// If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
	VpcId pulumi.StringPtrInput
	// The name of the VPC which the target server belongs to.
	VpcName pulumi.StringPtrInput
}

func (ServerTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTemplateState)(nil)).Elem()
}

type serverTemplateArgs struct {
	// Specifies the availability zone where the target server is located.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Specifies the bandwidth size in Mbit/s about the public IP address
	// that will be used for migration.
	BandwidthSize *int `pulumi:"bandwidthSize"`
	// Specifies the flavor ID for the target server.
	Flavor *string `pulumi:"flavor"`
	// Specifies the server template name.
	Name *string `pulumi:"name"`
	// Specifies the project ID where the target server is located.
	// If omitted, the default project in the region will be used.
	ProjectId *string `pulumi:"projectId"`
	// Specifies the region where the target server is located.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifies an array of one or more security group IDs to associate with
	// the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Specifies an array of one or more subnet IDs to attach to the target server.
	// If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
	SubnetIds []string `pulumi:"subnetIds"`
	// Specifies the name of the target server. Defaults to the template name.
	TargetServerName *string `pulumi:"targetServerName"`
	// Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
	// defaults to **SAS**.
	VolumeType *string `pulumi:"volumeType"`
	// Specifies the ID of the VPC which the target server belongs to.
	// If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ServerTemplate resource.
type ServerTemplateArgs struct {
	// Specifies the availability zone where the target server is located.
	AvailabilityZone pulumi.StringInput
	// Specifies the bandwidth size in Mbit/s about the public IP address
	// that will be used for migration.
	BandwidthSize pulumi.IntPtrInput
	// Specifies the flavor ID for the target server.
	Flavor pulumi.StringPtrInput
	// Specifies the server template name.
	Name pulumi.StringPtrInput
	// Specifies the project ID where the target server is located.
	// If omitted, the default project in the region will be used.
	ProjectId pulumi.StringPtrInput
	// Specifies the region where the target server is located.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput
	// Specifies an array of one or more security group IDs to associate with
	// the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
	SecurityGroupIds pulumi.StringArrayInput
	// Specifies an array of one or more subnet IDs to attach to the target server.
	// If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
	SubnetIds pulumi.StringArrayInput
	// Specifies the name of the target server. Defaults to the template name.
	TargetServerName pulumi.StringPtrInput
	// Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
	// defaults to **SAS**.
	VolumeType pulumi.StringPtrInput
	// Specifies the ID of the VPC which the target server belongs to.
	// If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
	VpcId pulumi.StringPtrInput
}

func (ServerTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTemplateArgs)(nil)).Elem()
}

type ServerTemplateInput interface {
	pulumi.Input

	ToServerTemplateOutput() ServerTemplateOutput
	ToServerTemplateOutputWithContext(ctx context.Context) ServerTemplateOutput
}

func (*ServerTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTemplate)(nil)).Elem()
}

func (i *ServerTemplate) ToServerTemplateOutput() ServerTemplateOutput {
	return i.ToServerTemplateOutputWithContext(context.Background())
}

func (i *ServerTemplate) ToServerTemplateOutputWithContext(ctx context.Context) ServerTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerTemplateOutput)
}

// ServerTemplateArrayInput is an input type that accepts ServerTemplateArray and ServerTemplateArrayOutput values.
// You can construct a concrete instance of `ServerTemplateArrayInput` via:
//
//	ServerTemplateArray{ ServerTemplateArgs{...} }
type ServerTemplateArrayInput interface {
	pulumi.Input

	ToServerTemplateArrayOutput() ServerTemplateArrayOutput
	ToServerTemplateArrayOutputWithContext(context.Context) ServerTemplateArrayOutput
}

type ServerTemplateArray []ServerTemplateInput

func (ServerTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerTemplate)(nil)).Elem()
}

func (i ServerTemplateArray) ToServerTemplateArrayOutput() ServerTemplateArrayOutput {
	return i.ToServerTemplateArrayOutputWithContext(context.Background())
}

func (i ServerTemplateArray) ToServerTemplateArrayOutputWithContext(ctx context.Context) ServerTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerTemplateArrayOutput)
}

// ServerTemplateMapInput is an input type that accepts ServerTemplateMap and ServerTemplateMapOutput values.
// You can construct a concrete instance of `ServerTemplateMapInput` via:
//
//	ServerTemplateMap{ "key": ServerTemplateArgs{...} }
type ServerTemplateMapInput interface {
	pulumi.Input

	ToServerTemplateMapOutput() ServerTemplateMapOutput
	ToServerTemplateMapOutputWithContext(context.Context) ServerTemplateMapOutput
}

type ServerTemplateMap map[string]ServerTemplateInput

func (ServerTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerTemplate)(nil)).Elem()
}

func (i ServerTemplateMap) ToServerTemplateMapOutput() ServerTemplateMapOutput {
	return i.ToServerTemplateMapOutputWithContext(context.Background())
}

func (i ServerTemplateMap) ToServerTemplateMapOutputWithContext(ctx context.Context) ServerTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerTemplateMapOutput)
}

type ServerTemplateOutput struct{ *pulumi.OutputState }

func (ServerTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTemplate)(nil)).Elem()
}

func (o ServerTemplateOutput) ToServerTemplateOutput() ServerTemplateOutput {
	return o
}

func (o ServerTemplateOutput) ToServerTemplateOutputWithContext(ctx context.Context) ServerTemplateOutput {
	return o
}

// Specifies the availability zone where the target server is located.
func (o ServerTemplateOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Specifies the bandwidth size in Mbit/s about the public IP address
// that will be used for migration.
func (o ServerTemplateOutput) BandwidthSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.IntPtrOutput { return v.BandwidthSize }).(pulumi.IntPtrOutput)
}

// Specifies the flavor ID for the target server.
func (o ServerTemplateOutput) Flavor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringPtrOutput { return v.Flavor }).(pulumi.StringPtrOutput)
}

// Specifies the server template name.
func (o ServerTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the project ID where the target server is located.
// If omitted, the default project in the region will be used.
func (o ServerTemplateOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Specifies the region where the target server is located.
// If omitted, the provider-level region will be used.
func (o ServerTemplateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies an array of one or more security group IDs to associate with
// the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
func (o ServerTemplateOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Specifies an array of one or more subnet IDs to attach to the target server.
// If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
func (o ServerTemplateOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Specifies the name of the target server. Defaults to the template name.
func (o ServerTemplateOutput) TargetServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringOutput { return v.TargetServerName }).(pulumi.StringOutput)
}

// Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
// defaults to **SAS**.
func (o ServerTemplateOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

// Specifies the ID of the VPC which the target server belongs to.
// If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
func (o ServerTemplateOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The name of the VPC which the target server belongs to.
func (o ServerTemplateOutput) VpcName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTemplate) pulumi.StringOutput { return v.VpcName }).(pulumi.StringOutput)
}

type ServerTemplateArrayOutput struct{ *pulumi.OutputState }

func (ServerTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerTemplate)(nil)).Elem()
}

func (o ServerTemplateArrayOutput) ToServerTemplateArrayOutput() ServerTemplateArrayOutput {
	return o
}

func (o ServerTemplateArrayOutput) ToServerTemplateArrayOutputWithContext(ctx context.Context) ServerTemplateArrayOutput {
	return o
}

func (o ServerTemplateArrayOutput) Index(i pulumi.IntInput) ServerTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerTemplate {
		return vs[0].([]*ServerTemplate)[vs[1].(int)]
	}).(ServerTemplateOutput)
}

type ServerTemplateMapOutput struct{ *pulumi.OutputState }

func (ServerTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerTemplate)(nil)).Elem()
}

func (o ServerTemplateMapOutput) ToServerTemplateMapOutput() ServerTemplateMapOutput {
	return o
}

func (o ServerTemplateMapOutput) ToServerTemplateMapOutputWithContext(ctx context.Context) ServerTemplateMapOutput {
	return o
}

func (o ServerTemplateMapOutput) MapIndex(k pulumi.StringInput) ServerTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerTemplate {
		return vs[0].(map[string]*ServerTemplate)[vs[1].(string)]
	}).(ServerTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerTemplateInput)(nil)).Elem(), &ServerTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerTemplateArrayInput)(nil)).Elem(), ServerTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerTemplateMapInput)(nil)).Elem(), ServerTemplateMap{})
	pulumi.RegisterOutputType(ServerTemplateOutput{})
	pulumi.RegisterOutputType(ServerTemplateArrayOutput{})
	pulumi.RegisterOutputType(ServerTemplateMapOutput{})
}
