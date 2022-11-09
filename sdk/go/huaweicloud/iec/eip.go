// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iec

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a eip resource within HuaweiCloud IEC.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Iec"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Iec"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			iecSites, err := Iec.GetSites(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Iec.NewEip(ctx, "eipTest", &Iec.EipArgs{
//				SiteId: pulumi.String(iecSites.Sites[0].Id),
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
// IEC EIPs can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Iec/eip:Eip eip_test b5ad19d1-57d1-48fd-aab7-1378f9bee169
//
// ```
type Eip struct {
	pulumi.CustomResourceState

	BandwidthId        pulumi.StringOutput `pulumi:"bandwidthId"`
	BandwidthName      pulumi.StringOutput `pulumi:"bandwidthName"`
	BandwidthShareType pulumi.StringOutput `pulumi:"bandwidthShareType"`
	BandwidthSize      pulumi.IntOutput    `pulumi:"bandwidthSize"`
	// The version of elastic IP address.
	IpVersion pulumi.IntOutput `pulumi:"ipVersion"`
	// Specifies the line ID of IEC sevice site.
	// Changing this parameter creates a new resource.
	LineId pulumi.StringOutput `pulumi:"lineId"`
	// Specifies the port ID which this eip will associate with.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The address of private IP.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The address of elastic IP.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	Region   pulumi.StringOutput `pulumi:"region"`
	// Specifies the ID of IEC sevice site. Changing this parameter creates a new
	// resource.
	SiteId pulumi.StringOutput `pulumi:"siteId"`
	// The located information of the IEC site. It contains area, province and city.
	SiteInfo pulumi.StringOutput `pulumi:"siteInfo"`
	// The status of elastic IP.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewEip registers a new resource with the given unique name, arguments, and options.
func NewEip(ctx *pulumi.Context,
	name string, args *EipArgs, opts ...pulumi.ResourceOption) (*Eip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Eip
	err := ctx.RegisterResource("huaweicloud:Iec/eip:Eip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEip gets an existing Eip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipState, opts ...pulumi.ResourceOption) (*Eip, error) {
	var resource Eip
	err := ctx.ReadResource("huaweicloud:Iec/eip:Eip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Eip resources.
type eipState struct {
	BandwidthId        *string `pulumi:"bandwidthId"`
	BandwidthName      *string `pulumi:"bandwidthName"`
	BandwidthShareType *string `pulumi:"bandwidthShareType"`
	BandwidthSize      *int    `pulumi:"bandwidthSize"`
	// The version of elastic IP address.
	IpVersion *int `pulumi:"ipVersion"`
	// Specifies the line ID of IEC sevice site.
	// Changing this parameter creates a new resource.
	LineId *string `pulumi:"lineId"`
	// Specifies the port ID which this eip will associate with.
	PortId *string `pulumi:"portId"`
	// The address of private IP.
	PrivateIp *string `pulumi:"privateIp"`
	// The address of elastic IP.
	PublicIp *string `pulumi:"publicIp"`
	Region   *string `pulumi:"region"`
	// Specifies the ID of IEC sevice site. Changing this parameter creates a new
	// resource.
	SiteId *string `pulumi:"siteId"`
	// The located information of the IEC site. It contains area, province and city.
	SiteInfo *string `pulumi:"siteInfo"`
	// The status of elastic IP.
	Status *string `pulumi:"status"`
}

type EipState struct {
	BandwidthId        pulumi.StringPtrInput
	BandwidthName      pulumi.StringPtrInput
	BandwidthShareType pulumi.StringPtrInput
	BandwidthSize      pulumi.IntPtrInput
	// The version of elastic IP address.
	IpVersion pulumi.IntPtrInput
	// Specifies the line ID of IEC sevice site.
	// Changing this parameter creates a new resource.
	LineId pulumi.StringPtrInput
	// Specifies the port ID which this eip will associate with.
	PortId pulumi.StringPtrInput
	// The address of private IP.
	PrivateIp pulumi.StringPtrInput
	// The address of elastic IP.
	PublicIp pulumi.StringPtrInput
	Region   pulumi.StringPtrInput
	// Specifies the ID of IEC sevice site. Changing this parameter creates a new
	// resource.
	SiteId pulumi.StringPtrInput
	// The located information of the IEC site. It contains area, province and city.
	SiteInfo pulumi.StringPtrInput
	// The status of elastic IP.
	Status pulumi.StringPtrInput
}

func (EipState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipState)(nil)).Elem()
}

type eipArgs struct {
	// The version of elastic IP address.
	IpVersion *int `pulumi:"ipVersion"`
	// Specifies the line ID of IEC sevice site.
	// Changing this parameter creates a new resource.
	LineId *string `pulumi:"lineId"`
	// Specifies the port ID which this eip will associate with.
	PortId *string `pulumi:"portId"`
	Region *string `pulumi:"region"`
	// Specifies the ID of IEC sevice site. Changing this parameter creates a new
	// resource.
	SiteId string `pulumi:"siteId"`
}

// The set of arguments for constructing a Eip resource.
type EipArgs struct {
	// The version of elastic IP address.
	IpVersion pulumi.IntPtrInput
	// Specifies the line ID of IEC sevice site.
	// Changing this parameter creates a new resource.
	LineId pulumi.StringPtrInput
	// Specifies the port ID which this eip will associate with.
	PortId pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// Specifies the ID of IEC sevice site. Changing this parameter creates a new
	// resource.
	SiteId pulumi.StringInput
}

func (EipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipArgs)(nil)).Elem()
}

type EipInput interface {
	pulumi.Input

	ToEipOutput() EipOutput
	ToEipOutputWithContext(ctx context.Context) EipOutput
}

func (*Eip) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (i *Eip) ToEipOutput() EipOutput {
	return i.ToEipOutputWithContext(context.Background())
}

func (i *Eip) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipOutput)
}

// EipArrayInput is an input type that accepts EipArray and EipArrayOutput values.
// You can construct a concrete instance of `EipArrayInput` via:
//
//	EipArray{ EipArgs{...} }
type EipArrayInput interface {
	pulumi.Input

	ToEipArrayOutput() EipArrayOutput
	ToEipArrayOutputWithContext(context.Context) EipArrayOutput
}

type EipArray []EipInput

func (EipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (i EipArray) ToEipArrayOutput() EipArrayOutput {
	return i.ToEipArrayOutputWithContext(context.Background())
}

func (i EipArray) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipArrayOutput)
}

// EipMapInput is an input type that accepts EipMap and EipMapOutput values.
// You can construct a concrete instance of `EipMapInput` via:
//
//	EipMap{ "key": EipArgs{...} }
type EipMapInput interface {
	pulumi.Input

	ToEipMapOutput() EipMapOutput
	ToEipMapOutputWithContext(context.Context) EipMapOutput
}

type EipMap map[string]EipInput

func (EipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (i EipMap) ToEipMapOutput() EipMapOutput {
	return i.ToEipMapOutputWithContext(context.Background())
}

func (i EipMap) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipMapOutput)
}

type EipOutput struct{ *pulumi.OutputState }

func (EipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (o EipOutput) ToEipOutput() EipOutput {
	return o
}

func (o EipOutput) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return o
}

func (o EipOutput) BandwidthId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.BandwidthId }).(pulumi.StringOutput)
}

func (o EipOutput) BandwidthName() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.BandwidthName }).(pulumi.StringOutput)
}

func (o EipOutput) BandwidthShareType() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.BandwidthShareType }).(pulumi.StringOutput)
}

func (o EipOutput) BandwidthSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Eip) pulumi.IntOutput { return v.BandwidthSize }).(pulumi.IntOutput)
}

// The version of elastic IP address.
func (o EipOutput) IpVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *Eip) pulumi.IntOutput { return v.IpVersion }).(pulumi.IntOutput)
}

// Specifies the line ID of IEC sevice site.
// Changing this parameter creates a new resource.
func (o EipOutput) LineId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.LineId }).(pulumi.StringOutput)
}

// Specifies the port ID which this eip will associate with.
func (o EipOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// The address of private IP.
func (o EipOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// The address of elastic IP.
func (o EipOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

func (o EipOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the ID of IEC sevice site. Changing this parameter creates a new
// resource.
func (o EipOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.SiteId }).(pulumi.StringOutput)
}

// The located information of the IEC site. It contains area, province and city.
func (o EipOutput) SiteInfo() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.SiteInfo }).(pulumi.StringOutput)
}

// The status of elastic IP.
func (o EipOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type EipArrayOutput struct{ *pulumi.OutputState }

func (EipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (o EipArrayOutput) ToEipArrayOutput() EipArrayOutput {
	return o
}

func (o EipArrayOutput) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return o
}

func (o EipArrayOutput) Index(i pulumi.IntInput) EipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].([]*Eip)[vs[1].(int)]
	}).(EipOutput)
}

type EipMapOutput struct{ *pulumi.OutputState }

func (EipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (o EipMapOutput) ToEipMapOutput() EipMapOutput {
	return o
}

func (o EipMapOutput) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return o
}

func (o EipMapOutput) MapIndex(k pulumi.StringInput) EipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].(map[string]*Eip)[vs[1].(string)]
	}).(EipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EipInput)(nil)).Elem(), &Eip{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipArrayInput)(nil)).Elem(), EipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipMapInput)(nil)).Elem(), EipMap{})
	pulumi.RegisterOutputType(EipOutput{})
	pulumi.RegisterOutputType(EipArrayOutput{})
	pulumi.RegisterOutputType(EipMapOutput{})
}
