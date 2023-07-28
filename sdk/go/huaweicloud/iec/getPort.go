// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of a specific IEC subnet port.
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			subnetId := cfg.RequireObject("subnetId")
//			_, err := Iec.GetPort(ctx, &iec.GetPortArgs{
//				SubnetId: pulumi.StringRef(subnetId),
//				FixedIp:  pulumi.StringRef("192.168.1.123"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPort(ctx *pulumi.Context, args *GetPortArgs, opts ...pulumi.InvokeOption) (*GetPortResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPortResult
	err := ctx.Invoke("huaweicloud:Iec/getPort:getPort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPort.
type GetPortArgs struct {
	// The IP address of the port.
	FixedIp *string `pulumi:"fixedIp"`
	// The ID of the port.
	Id *string `pulumi:"id"`
	// The MAC address of the port.
	MacAddress *string `pulumi:"macAddress"`
	// The region in which to obtain the port. If omitted, the provider-level region will be
	// used.
	Region *string `pulumi:"region"`
	// The ID of the subnet which the port belongs to.
	SubnetId *string `pulumi:"subnetId"`
}

// A collection of values returned by getPort.
type GetPortResult struct {
	FixedIp string `pulumi:"fixedIp"`
	// Specifies a data source ID in UUID format.
	Id         string `pulumi:"id"`
	MacAddress string `pulumi:"macAddress"`
	Region     string `pulumi:"region"`
	// Indicates the list of security group IDs applied on the port.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Indicates the ID of the IEC site.
	SiteId string `pulumi:"siteId"`
	// Indicates the status of the port.
	Status   string `pulumi:"status"`
	SubnetId string `pulumi:"subnetId"`
}

func GetPortOutput(ctx *pulumi.Context, args GetPortOutputArgs, opts ...pulumi.InvokeOption) GetPortResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPortResult, error) {
			args := v.(GetPortArgs)
			r, err := GetPort(ctx, &args, opts...)
			var s GetPortResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPortResultOutput)
}

// A collection of arguments for invoking getPort.
type GetPortOutputArgs struct {
	// The IP address of the port.
	FixedIp pulumi.StringPtrInput `pulumi:"fixedIp"`
	// The ID of the port.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The MAC address of the port.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The region in which to obtain the port. If omitted, the provider-level region will be
	// used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The ID of the subnet which the port belongs to.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (GetPortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPortArgs)(nil)).Elem()
}

// A collection of values returned by getPort.
type GetPortResultOutput struct{ *pulumi.OutputState }

func (GetPortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPortResult)(nil)).Elem()
}

func (o GetPortResultOutput) ToGetPortResultOutput() GetPortResultOutput {
	return o
}

func (o GetPortResultOutput) ToGetPortResultOutputWithContext(ctx context.Context) GetPortResultOutput {
	return o
}

func (o GetPortResultOutput) FixedIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortResult) string { return v.FixedIp }).(pulumi.StringOutput)
}

// Specifies a data source ID in UUID format.
func (o GetPortResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPortResultOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortResult) string { return v.MacAddress }).(pulumi.StringOutput)
}

func (o GetPortResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the list of security group IDs applied on the port.
func (o GetPortResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPortResult) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Indicates the ID of the IEC site.
func (o GetPortResultOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortResult) string { return v.SiteId }).(pulumi.StringOutput)
}

// Indicates the status of the port.
func (o GetPortResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetPortResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPortResultOutput{})
}
