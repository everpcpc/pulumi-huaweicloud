// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific VPC subnet.
//
// This resource can prove useful when a module accepts a subnet id as an input variable and needs to, for example,
// determine the id of the VPC that the subnet belongs to.
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
//			subnet, err := Vpc.GetSubnet(ctx, &vpc.GetSubnetArgs{
//				Id: pulumi.StringRef(_var.Subnet_id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("subnetVpcId", subnet.VpcId)
//			return nil
//		})
//	}
//
// ```
// ## **Attributes Reference**
//
// In addition to all arguments above, the following attributes are exported:
//
// * `dnsList` - The IP address list of DNS servers on the subnet.
//
// * `dhcpEnable` - Whether the DHCP is enabled.
//
// * `description` - The description of the subnet.
//
// * `ipv4SubnetId` - The ID of the IPv4 subnet (Native OpenStack API).
//
// * `ipv6Enable` - Whether the IPv6 is enabled.
//
// * `ipv6SubnetId` - The ID of the IPv6 subnet (Native OpenStack API).
//
// * `ipv6Cidr` - The IPv6 subnet CIDR block.
//
// * `ipv6Gateway` - The IPv6 subnet gateway.
func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSubnetResult
	err := ctx.Invoke("huaweicloud:Vpc/getSubnet:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetArgs struct {
	// Specifies the availability zone (AZ) to which the subnet should belong.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the network segment of specific subnet to retrieve. The value must be in CIDR
	// format.
	Cidr *string `pulumi:"cidr"`
	// Specifies the subnet gateway address of specific subnet.
	GatewayIp *string `pulumi:"gatewayIp"`
	// - Specifies a resource ID in UUID format.
	Id *string `pulumi:"id"`
	// Specifies the name of the specific subnet to retrieve.
	Name *string `pulumi:"name"`
	// Specifies the IP address of DNS server 1 on the specific subnet.
	PrimaryDns *string `pulumi:"primaryDns"`
	// Specifies the region in which to obtain the subnet. If omitted, the provider-level
	// region will be used.
	Region *string `pulumi:"region"`
	// Specifies the IP address of DNS server 2 on the specific subnet.
	SecondaryDns *string `pulumi:"secondaryDns"`
	// Specifies the value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
	Status *string `pulumi:"status"`
	// Specifies the id of the VPC that the desired subnet belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getSubnet.
type LookupSubnetResult struct {
	AvailabilityZone string   `pulumi:"availabilityZone"`
	Cidr             string   `pulumi:"cidr"`
	Description      string   `pulumi:"description"`
	DhcpEnable       bool     `pulumi:"dhcpEnable"`
	DnsLists         []string `pulumi:"dnsLists"`
	GatewayIp        string   `pulumi:"gatewayIp"`
	Id               string   `pulumi:"id"`
	Ipv4SubnetId     string   `pulumi:"ipv4SubnetId"`
	Ipv6Cidr         string   `pulumi:"ipv6Cidr"`
	Ipv6Enable       bool     `pulumi:"ipv6Enable"`
	Ipv6Gateway      string   `pulumi:"ipv6Gateway"`
	Ipv6SubnetId     string   `pulumi:"ipv6SubnetId"`
	Name             string   `pulumi:"name"`
	PrimaryDns       string   `pulumi:"primaryDns"`
	Region           string   `pulumi:"region"`
	SecondaryDns     string   `pulumi:"secondaryDns"`
	Status           string   `pulumi:"status"`
	SubnetId         string   `pulumi:"subnetId"`
	VpcId            string   `pulumi:"vpcId"`
}

func LookupSubnetOutput(ctx *pulumi.Context, args LookupSubnetOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetResult, error) {
			args := v.(LookupSubnetArgs)
			r, err := LookupSubnet(ctx, &args, opts...)
			var s LookupSubnetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetResultOutput)
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetOutputArgs struct {
	// Specifies the availability zone (AZ) to which the subnet should belong.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// Specifies the network segment of specific subnet to retrieve. The value must be in CIDR
	// format.
	Cidr pulumi.StringPtrInput `pulumi:"cidr"`
	// Specifies the subnet gateway address of specific subnet.
	GatewayIp pulumi.StringPtrInput `pulumi:"gatewayIp"`
	// - Specifies a resource ID in UUID format.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Specifies the name of the specific subnet to retrieve.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the IP address of DNS server 1 on the specific subnet.
	PrimaryDns pulumi.StringPtrInput `pulumi:"primaryDns"`
	// Specifies the region in which to obtain the subnet. If omitted, the provider-level
	// region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the IP address of DNS server 2 on the specific subnet.
	SecondaryDns pulumi.StringPtrInput `pulumi:"secondaryDns"`
	// Specifies the value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Specifies the id of the VPC that the desired subnet belongs to.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupSubnetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetArgs)(nil)).Elem()
}

// A collection of values returned by getSubnet.
type LookupSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetResult)(nil)).Elem()
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutput() LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutputWithContext(ctx context.Context) LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Cidr }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) DhcpEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetResult) bool { return v.DhcpEnable }).(pulumi.BoolOutput)
}

func (o LookupSubnetResultOutput) DnsLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.DnsLists }).(pulumi.StringArrayOutput)
}

func (o LookupSubnetResultOutput) GatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.GatewayIp }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Ipv4SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Ipv4SubnetId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Ipv6Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Ipv6Cidr }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Ipv6Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetResult) bool { return v.Ipv6Enable }).(pulumi.BoolOutput)
}

func (o LookupSubnetResultOutput) Ipv6Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Ipv6Gateway }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Ipv6SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Ipv6SubnetId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) PrimaryDns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.PrimaryDns }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) SecondaryDns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.SecondaryDns }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetResultOutput{})
}
