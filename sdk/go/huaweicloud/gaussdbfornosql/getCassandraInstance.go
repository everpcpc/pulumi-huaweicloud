// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaussdbfornosql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get available HuaweiCloud gaussdb cassandra instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDBforNoSQL"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDBforNoSQL"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := GaussDBforNoSQL.GetCassandraInstance(ctx, &gaussdbfornosql.GetCassandraInstanceArgs{
//				Name: pulumi.StringRef("gaussdb-instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCassandraInstance(ctx *pulumi.Context, args *LookupCassandraInstanceArgs, opts ...pulumi.InvokeOption) (*LookupCassandraInstanceResult, error) {
	var rv LookupCassandraInstanceResult
	err := ctx.Invoke("huaweicloud:GaussDBforNoSQL/getCassandraInstance:getCassandraInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCassandraInstance.
type LookupCassandraInstanceArgs struct {
	// Specifies the name of the instance.
	Name *string `pulumi:"name"`
	// The region in which to obtain the instance. If omitted, the provider-level region will
	// be used.
	Region *string `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Specifies the VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getCassandraInstance.
type LookupCassandraInstanceResult struct {
	// Indicates the availability zone where the node resides.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Indicates the advanced backup policy. Structure is documented below.
	BackupStrategies []GetCassandraInstanceBackupStrategy `pulumi:"backupStrategies"`
	// Indicates the database information. Structure is documented below.
	Datastores []GetCassandraInstanceDatastore `pulumi:"datastores"`
	// Indicates the default username.
	DbUserName string `pulumi:"dbUserName"`
	// Indicates the enterprise project id.
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	// Indicates the instance specifications.
	Flavor string `pulumi:"flavor"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates the instance mode.
	Mode string `pulumi:"mode"`
	// Indicates the node name.
	Name string `pulumi:"name"`
	// Indicates the count of the nodes.
	NodeNum int `pulumi:"nodeNum"`
	// Indicates the instance nodes information. Structure is documented below.
	Nodes []GetCassandraInstanceNode `pulumi:"nodes"`
	// Indicates the database port.
	Port int `pulumi:"port"`
	// Indicates the list of private IP address of the nodes.
	PrivateIps []string `pulumi:"privateIps"`
	Region     string   `pulumi:"region"`
	// Indicates the security group ID.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Indicates the node status.
	Status   string `pulumi:"status"`
	SubnetId string `pulumi:"subnetId"`
	// Indicates the key/value tags of the instance.
	Tags map[string]string `pulumi:"tags"`
	// Indicates the size of the volume.
	VolumeSize int    `pulumi:"volumeSize"`
	VpcId      string `pulumi:"vpcId"`
}

func LookupCassandraInstanceOutput(ctx *pulumi.Context, args LookupCassandraInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraInstanceResult, error) {
			args := v.(LookupCassandraInstanceArgs)
			r, err := LookupCassandraInstance(ctx, &args, opts...)
			var s LookupCassandraInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraInstanceResultOutput)
}

// A collection of arguments for invoking getCassandraInstance.
type LookupCassandraInstanceOutputArgs struct {
	// Specifies the name of the instance.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the instance. If omitted, the provider-level region will
	// be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Specifies the VPC ID.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupCassandraInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getCassandraInstance.
type LookupCassandraInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraInstanceResult)(nil)).Elem()
}

func (o LookupCassandraInstanceResultOutput) ToLookupCassandraInstanceResultOutput() LookupCassandraInstanceResultOutput {
	return o
}

func (o LookupCassandraInstanceResultOutput) ToLookupCassandraInstanceResultOutputWithContext(ctx context.Context) LookupCassandraInstanceResultOutput {
	return o
}

// Indicates the availability zone where the node resides.
func (o LookupCassandraInstanceResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Indicates the advanced backup policy. Structure is documented below.
func (o LookupCassandraInstanceResultOutput) BackupStrategies() GetCassandraInstanceBackupStrategyArrayOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) []GetCassandraInstanceBackupStrategy { return v.BackupStrategies }).(GetCassandraInstanceBackupStrategyArrayOutput)
}

// Indicates the database information. Structure is documented below.
func (o LookupCassandraInstanceResultOutput) Datastores() GetCassandraInstanceDatastoreArrayOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) []GetCassandraInstanceDatastore { return v.Datastores }).(GetCassandraInstanceDatastoreArrayOutput)
}

// Indicates the default username.
func (o LookupCassandraInstanceResultOutput) DbUserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.DbUserName }).(pulumi.StringOutput)
}

// Indicates the enterprise project id.
func (o LookupCassandraInstanceResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Indicates the instance specifications.
func (o LookupCassandraInstanceResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCassandraInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates the instance mode.
func (o LookupCassandraInstanceResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.Mode }).(pulumi.StringOutput)
}

// Indicates the node name.
func (o LookupCassandraInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates the count of the nodes.
func (o LookupCassandraInstanceResultOutput) NodeNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) int { return v.NodeNum }).(pulumi.IntOutput)
}

// Indicates the instance nodes information. Structure is documented below.
func (o LookupCassandraInstanceResultOutput) Nodes() GetCassandraInstanceNodeArrayOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) []GetCassandraInstanceNode { return v.Nodes }).(GetCassandraInstanceNodeArrayOutput)
}

// Indicates the database port.
func (o LookupCassandraInstanceResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) int { return v.Port }).(pulumi.IntOutput)
}

// Indicates the list of private IP address of the nodes.
func (o LookupCassandraInstanceResultOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) []string { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

func (o LookupCassandraInstanceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the security group ID.
func (o LookupCassandraInstanceResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Indicates the node status.
func (o LookupCassandraInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupCassandraInstanceResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// Indicates the key/value tags of the instance.
func (o LookupCassandraInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Indicates the size of the volume.
func (o LookupCassandraInstanceResultOutput) VolumeSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) int { return v.VolumeSize }).(pulumi.IntOutput)
}

func (o LookupCassandraInstanceResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraInstanceResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraInstanceResultOutput{})
}