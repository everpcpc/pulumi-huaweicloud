// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of a specified IEC server.
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
//			serverName := cfg.RequireObject("serverName")
//			_, err := Iec.GetServer(ctx, &iec.GetServerArgs{
//				Name: serverName,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupServerResult
	err := ctx.Invoke("huaweicloud:Iec/getServer:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServer.
type LookupServerArgs struct {
	// Specifies the ID of the edgecloud service.
	EdgecloudId *string `pulumi:"edgecloudId"`
	// Specifies the IEC server name, which can be queried with a regular expression.
	Name string `pulumi:"name"`
	// Specifies the status of IEC server.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getServer.
type LookupServerResult struct {
	// An array of site ID and operator for the IEC server. The object structure is documented below.
	CoverageSites []GetServerCoverageSite `pulumi:"coverageSites"`
	EdgecloudId   string                  `pulumi:"edgecloudId"`
	// The Name of the edgecloud service.
	EdgecloudName string `pulumi:"edgecloudName"`
	// The flavor ID of the IEC server.
	FlavorId string `pulumi:"flavorId"`
	// The flavor name of the IEC server.
	FlavorName string `pulumi:"flavorName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The image ID of the IEC server.
	ImageId string `pulumi:"imageId"`
	// The image name of the IEC server.
	ImageName string `pulumi:"imageName"`
	// The name of a key pair to put on the IEC server.
	KeyPair string `pulumi:"keyPair"`
	Name    string `pulumi:"name"`
	// An array of one or more networks to attach to the IEC server. The object structure is documented below.
	Nics []GetServerNic `pulumi:"nics"`
	// The EIP address that is associted to the IEC server.
	PublicIp string `pulumi:"publicIp"`
	// An array of one or more security group IDs to associate with the IEC server.
	SecurityGroups []string `pulumi:"securityGroups"`
	Status         string   `pulumi:"status"`
	// The system disk voume ID.
	SystemDiskId string `pulumi:"systemDiskId"`
	// The user data (information after encoding) configured during IEC server creation.
	UserData string `pulumi:"userData"`
	// An array of one or more disks to attach to the IEC server. The object structure is documented
	// below.
	VolumeAttacheds []GetServerVolumeAttached `pulumi:"volumeAttacheds"`
	// The ID of vpc for the IEC server.
	VpcId string `pulumi:"vpcId"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			var s LookupServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerResultOutput)
}

// A collection of arguments for invoking getServer.
type LookupServerOutputArgs struct {
	// Specifies the ID of the edgecloud service.
	EdgecloudId pulumi.StringPtrInput `pulumi:"edgecloudId"`
	// Specifies the IEC server name, which can be queried with a regular expression.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the status of IEC server.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

// A collection of values returned by getServer.
type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

// An array of site ID and operator for the IEC server. The object structure is documented below.
func (o LookupServerResultOutput) CoverageSites() GetServerCoverageSiteArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []GetServerCoverageSite { return v.CoverageSites }).(GetServerCoverageSiteArrayOutput)
}

func (o LookupServerResultOutput) EdgecloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.EdgecloudId }).(pulumi.StringOutput)
}

// The Name of the edgecloud service.
func (o LookupServerResultOutput) EdgecloudName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.EdgecloudName }).(pulumi.StringOutput)
}

// The flavor ID of the IEC server.
func (o LookupServerResultOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FlavorId }).(pulumi.StringOutput)
}

// The flavor name of the IEC server.
func (o LookupServerResultOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FlavorName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The image ID of the IEC server.
func (o LookupServerResultOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ImageId }).(pulumi.StringOutput)
}

// The image name of the IEC server.
func (o LookupServerResultOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ImageName }).(pulumi.StringOutput)
}

// The name of a key pair to put on the IEC server.
func (o LookupServerResultOutput) KeyPair() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.KeyPair }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// An array of one or more networks to attach to the IEC server. The object structure is documented below.
func (o LookupServerResultOutput) Nics() GetServerNicArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []GetServerNic { return v.Nics }).(GetServerNicArrayOutput)
}

// The EIP address that is associted to the IEC server.
func (o LookupServerResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

// An array of one or more security group IDs to associate with the IEC server.
func (o LookupServerResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o LookupServerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Status }).(pulumi.StringOutput)
}

// The system disk voume ID.
func (o LookupServerResultOutput) SystemDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.SystemDiskId }).(pulumi.StringOutput)
}

// The user data (information after encoding) configured during IEC server creation.
func (o LookupServerResultOutput) UserData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.UserData }).(pulumi.StringOutput)
}

// An array of one or more disks to attach to the IEC server. The object structure is documented
// below.
func (o LookupServerResultOutput) VolumeAttacheds() GetServerVolumeAttachedArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []GetServerVolumeAttached { return v.VolumeAttacheds }).(GetServerVolumeAttachedArrayOutput)
}

// The ID of vpc for the IEC server.
func (o LookupServerResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
