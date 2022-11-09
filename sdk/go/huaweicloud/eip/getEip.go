// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eip

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of an available EIP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Eip"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Eip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Eip.GetEip(ctx, &eip.GetEipArgs{
//				PublicIp: pulumi.StringRef("123.60.208.163"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetEip(ctx *pulumi.Context, args *GetEipArgs, opts ...pulumi.InvokeOption) (*GetEipResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetEipResult
	err := ctx.Invoke("huaweicloud:Eip/getEip:getEip", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEip.
type GetEipArgs struct {
	// Specifies the enterprise project id of the EIP.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the port id of the EIP.
	PortId *string `pulumi:"portId"`
	// Specifies the public **IPv4** address of the EIP.
	PublicIp *string `pulumi:"publicIp"`
	// Specifies the region in which to obtain the EIP.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getEip.
type GetEipResult struct {
	// The bandwidth id of the EIP.
	BandwidthId string `pulumi:"bandwidthId"`
	// The bandwidth share type of the EIP.
	BandwidthShareType string `pulumi:"bandwidthShareType"`
	// The bandwidth size of the EIP.
	BandwidthSize       int    `pulumi:"bandwidthSize"`
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IP version, either 4 or 6.
	IpVersion int `pulumi:"ipVersion"`
	// The IPv6 address of the EIP.
	Ipv6Address string `pulumi:"ipv6Address"`
	PortId      string `pulumi:"portId"`
	// The private ip of the EIP.
	PrivateIp string `pulumi:"privateIp"`
	PublicIp  string `pulumi:"publicIp"`
	Region    string `pulumi:"region"`
	// The status of the EIP.
	Status string `pulumi:"status"`
	// The type of the EIP.
	Type string `pulumi:"type"`
}

func GetEipOutput(ctx *pulumi.Context, args GetEipOutputArgs, opts ...pulumi.InvokeOption) GetEipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEipResult, error) {
			args := v.(GetEipArgs)
			r, err := GetEip(ctx, &args, opts...)
			var s GetEipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEipResultOutput)
}

// A collection of arguments for invoking getEip.
type GetEipOutputArgs struct {
	// Specifies the enterprise project id of the EIP.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Specifies the port id of the EIP.
	PortId pulumi.StringPtrInput `pulumi:"portId"`
	// Specifies the public **IPv4** address of the EIP.
	PublicIp pulumi.StringPtrInput `pulumi:"publicIp"`
	// Specifies the region in which to obtain the EIP.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetEipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEipArgs)(nil)).Elem()
}

// A collection of values returned by getEip.
type GetEipResultOutput struct{ *pulumi.OutputState }

func (GetEipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEipResult)(nil)).Elem()
}

func (o GetEipResultOutput) ToGetEipResultOutput() GetEipResultOutput {
	return o
}

func (o GetEipResultOutput) ToGetEipResultOutputWithContext(ctx context.Context) GetEipResultOutput {
	return o
}

// The bandwidth id of the EIP.
func (o GetEipResultOutput) BandwidthId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.BandwidthId }).(pulumi.StringOutput)
}

// The bandwidth share type of the EIP.
func (o GetEipResultOutput) BandwidthShareType() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.BandwidthShareType }).(pulumi.StringOutput)
}

// The bandwidth size of the EIP.
func (o GetEipResultOutput) BandwidthSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetEipResult) int { return v.BandwidthSize }).(pulumi.IntOutput)
}

func (o GetEipResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IP version, either 4 or 6.
func (o GetEipResultOutput) IpVersion() pulumi.IntOutput {
	return o.ApplyT(func(v GetEipResult) int { return v.IpVersion }).(pulumi.IntOutput)
}

// The IPv6 address of the EIP.
func (o GetEipResultOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.Ipv6Address }).(pulumi.StringOutput)
}

func (o GetEipResultOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.PortId }).(pulumi.StringOutput)
}

// The private ip of the EIP.
func (o GetEipResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o GetEipResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

func (o GetEipResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.Region }).(pulumi.StringOutput)
}

// The status of the EIP.
func (o GetEipResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.Status }).(pulumi.StringOutput)
}

// The type of the EIP.
func (o GetEipResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEipResultOutput{})
}
