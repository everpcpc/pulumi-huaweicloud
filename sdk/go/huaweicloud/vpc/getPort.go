// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available HuaweiCloud port.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.GetPort(ctx, &vpc.GetPortArgs{
//				NetworkId: pulumi.StringRef(_var.Network_id),
//				FixedIp:   pulumi.StringRef("192.168.0.100"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPort(ctx *pulumi.Context, args *LookupPortArgs, opts ...pulumi.InvokeOption) (*LookupPortResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupPortResult
	err := ctx.Invoke("huaweicloud:Vpc/getPort:getPort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPort.
type LookupPortArgs struct {
	// The administrative state of the port.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The ID of the device the port belongs to.
	DeviceId *string `pulumi:"deviceId"`
	// The device owner of the port.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// Specifies the port IP address filter.
	FixedIp *string `pulumi:"fixedIp"`
	// Specifies the MAC address of the port.
	MacAddress *string `pulumi:"macAddress"`
	// The name of the port.
	Name *string `pulumi:"name"`
	// Specifies the ID of the network the port belongs to.
	NetworkId *string `pulumi:"networkId"`
	// Specifies the ID of the port.
	PortId    *string `pulumi:"portId"`
	ProjectId *string `pulumi:"projectId"`
	// Specifies the region in which to obtain the port. If omitted, the provider-level region
	// will be used.
	Region *string `pulumi:"region"`
	// The list of port security group IDs to filter.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Specifies the status of the port.
	Status   *string `pulumi:"status"`
	TenantId *string `pulumi:"tenantId"`
}

// A collection of values returned by getPort.
type LookupPortResult struct {
	// The administrative state of the port.
	AdminStateUp bool `pulumi:"adminStateUp"`
	// The collection of Fixed IP addresses on the port.
	AllFixedIps []string `pulumi:"allFixedIps"`
	// The collection of security group IDs applied on the port.
	AllSecurityGroupIds []string `pulumi:"allSecurityGroupIds"`
	// The ID of the device the port belongs to.
	DeviceId string `pulumi:"deviceId"`
	// The device owner of the port.
	DeviceOwner string  `pulumi:"deviceOwner"`
	FixedIp     *string `pulumi:"fixedIp"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	MacAddress string `pulumi:"macAddress"`
	// The name of the port.
	Name             string   `pulumi:"name"`
	NetworkId        string   `pulumi:"networkId"`
	PortId           string   `pulumi:"portId"`
	ProjectId        *string  `pulumi:"projectId"`
	Region           string   `pulumi:"region"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	Status           string   `pulumi:"status"`
	TenantId         *string  `pulumi:"tenantId"`
}

func LookupPortOutput(ctx *pulumi.Context, args LookupPortOutputArgs, opts ...pulumi.InvokeOption) LookupPortResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPortResult, error) {
			args := v.(LookupPortArgs)
			r, err := LookupPort(ctx, &args, opts...)
			var s LookupPortResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPortResultOutput)
}

// A collection of arguments for invoking getPort.
type LookupPortOutputArgs struct {
	// The administrative state of the port.
	AdminStateUp pulumi.BoolPtrInput `pulumi:"adminStateUp"`
	// The ID of the device the port belongs to.
	DeviceId pulumi.StringPtrInput `pulumi:"deviceId"`
	// The device owner of the port.
	DeviceOwner pulumi.StringPtrInput `pulumi:"deviceOwner"`
	// Specifies the port IP address filter.
	FixedIp pulumi.StringPtrInput `pulumi:"fixedIp"`
	// Specifies the MAC address of the port.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The name of the port.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the ID of the network the port belongs to.
	NetworkId pulumi.StringPtrInput `pulumi:"networkId"`
	// Specifies the ID of the port.
	PortId    pulumi.StringPtrInput `pulumi:"portId"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// Specifies the region in which to obtain the port. If omitted, the provider-level region
	// will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The list of port security group IDs to filter.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// Specifies the status of the port.
	Status   pulumi.StringPtrInput `pulumi:"status"`
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupPortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortArgs)(nil)).Elem()
}

// A collection of values returned by getPort.
type LookupPortResultOutput struct{ *pulumi.OutputState }

func (LookupPortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortResult)(nil)).Elem()
}

func (o LookupPortResultOutput) ToLookupPortResultOutput() LookupPortResultOutput {
	return o
}

func (o LookupPortResultOutput) ToLookupPortResultOutputWithContext(ctx context.Context) LookupPortResultOutput {
	return o
}

// The administrative state of the port.
func (o LookupPortResultOutput) AdminStateUp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPortResult) bool { return v.AdminStateUp }).(pulumi.BoolOutput)
}

// The collection of Fixed IP addresses on the port.
func (o LookupPortResultOutput) AllFixedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.AllFixedIps }).(pulumi.StringArrayOutput)
}

// The collection of security group IDs applied on the port.
func (o LookupPortResultOutput) AllSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.AllSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The ID of the device the port belongs to.
func (o LookupPortResultOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.DeviceId }).(pulumi.StringOutput)
}

// The device owner of the port.
func (o LookupPortResultOutput) DeviceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.DeviceOwner }).(pulumi.StringOutput)
}

func (o LookupPortResultOutput) FixedIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.FixedIp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPortResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPortResultOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.MacAddress }).(pulumi.StringOutput)
}

// The name of the port.
func (o LookupPortResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPortResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

func (o LookupPortResultOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.PortId }).(pulumi.StringOutput)
}

func (o LookupPortResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupPortResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupPortResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupPortResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupPortResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPortResultOutput{})
}
