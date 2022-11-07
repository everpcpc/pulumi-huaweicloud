// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaussdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get available HuaweiCloud gaussdb mysql instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := GaussDB.GetMysqlInstance(ctx, &gaussdb.GetMysqlInstanceArgs{
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
func LookupMysqlInstance(ctx *pulumi.Context, args *LookupMysqlInstanceArgs, opts ...pulumi.InvokeOption) (*LookupMysqlInstanceResult, error) {
	var rv LookupMysqlInstanceResult
	err := ctx.Invoke("huaweicloud:GaussDB/getMysqlInstance:getMysqlInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMysqlInstance.
type LookupMysqlInstanceArgs struct {
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

// A collection of values returned by getMysqlInstance.
type LookupMysqlInstanceResult struct {
	// Indicates the availability zone mode: "single" or "multi".
	AvailabilityZoneMode string `pulumi:"availabilityZoneMode"`
	// Indicates the advanced backup policy. Structure is documented below.
	BackupStrategies []GetMysqlInstanceBackupStrategy `pulumi:"backupStrategies"`
	// Indicates the configuration ID.
	ConfigurationId string `pulumi:"configurationId"`
	// Indicates the database information. Structure is documented below.
	Datastores []GetMysqlInstanceDatastore `pulumi:"datastores"`
	// Indicates the default username.
	DbUserName string `pulumi:"dbUserName"`
	// Indicates the enterprise project id.
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	// Indicates the instance specifications.
	Flavor string `pulumi:"flavor"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates the availability zone where the master node resides.
	MasterAvailabilityZone string `pulumi:"masterAvailabilityZone"`
	// Indicates the instance mode.
	Mode string `pulumi:"mode"`
	// Indicates the node name.
	Name string `pulumi:"name"`
	// Indicates the instance nodes information. Structure is documented below.
	Nodes []GetMysqlInstanceNode `pulumi:"nodes"`
	// Indicates the database port.
	Port int `pulumi:"port"`
	// Indicates the private IP address of the DB instance.
	PrivateWriteIp string `pulumi:"privateWriteIp"`
	// Indicates the count of read replicas.
	ReadReplicas int    `pulumi:"readReplicas"`
	Region       string `pulumi:"region"`
	// Indicates the security group ID.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Indicates the node status.
	Status   string `pulumi:"status"`
	SubnetId string `pulumi:"subnetId"`
	// Indicates the time zone.
	TimeZone string `pulumi:"timeZone"`
	VpcId    string `pulumi:"vpcId"`
}

func LookupMysqlInstanceOutput(ctx *pulumi.Context, args LookupMysqlInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupMysqlInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMysqlInstanceResult, error) {
			args := v.(LookupMysqlInstanceArgs)
			r, err := LookupMysqlInstance(ctx, &args, opts...)
			var s LookupMysqlInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMysqlInstanceResultOutput)
}

// A collection of arguments for invoking getMysqlInstance.
type LookupMysqlInstanceOutputArgs struct {
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

func (LookupMysqlInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMysqlInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getMysqlInstance.
type LookupMysqlInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupMysqlInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMysqlInstanceResult)(nil)).Elem()
}

func (o LookupMysqlInstanceResultOutput) ToLookupMysqlInstanceResultOutput() LookupMysqlInstanceResultOutput {
	return o
}

func (o LookupMysqlInstanceResultOutput) ToLookupMysqlInstanceResultOutputWithContext(ctx context.Context) LookupMysqlInstanceResultOutput {
	return o
}

// Indicates the availability zone mode: "single" or "multi".
func (o LookupMysqlInstanceResultOutput) AvailabilityZoneMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.AvailabilityZoneMode }).(pulumi.StringOutput)
}

// Indicates the advanced backup policy. Structure is documented below.
func (o LookupMysqlInstanceResultOutput) BackupStrategies() GetMysqlInstanceBackupStrategyArrayOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) []GetMysqlInstanceBackupStrategy { return v.BackupStrategies }).(GetMysqlInstanceBackupStrategyArrayOutput)
}

// Indicates the configuration ID.
func (o LookupMysqlInstanceResultOutput) ConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.ConfigurationId }).(pulumi.StringOutput)
}

// Indicates the database information. Structure is documented below.
func (o LookupMysqlInstanceResultOutput) Datastores() GetMysqlInstanceDatastoreArrayOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) []GetMysqlInstanceDatastore { return v.Datastores }).(GetMysqlInstanceDatastoreArrayOutput)
}

// Indicates the default username.
func (o LookupMysqlInstanceResultOutput) DbUserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.DbUserName }).(pulumi.StringOutput)
}

// Indicates the enterprise project id.
func (o LookupMysqlInstanceResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Indicates the instance specifications.
func (o LookupMysqlInstanceResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMysqlInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates the availability zone where the master node resides.
func (o LookupMysqlInstanceResultOutput) MasterAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.MasterAvailabilityZone }).(pulumi.StringOutput)
}

// Indicates the instance mode.
func (o LookupMysqlInstanceResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.Mode }).(pulumi.StringOutput)
}

// Indicates the node name.
func (o LookupMysqlInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates the instance nodes information. Structure is documented below.
func (o LookupMysqlInstanceResultOutput) Nodes() GetMysqlInstanceNodeArrayOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) []GetMysqlInstanceNode { return v.Nodes }).(GetMysqlInstanceNodeArrayOutput)
}

// Indicates the database port.
func (o LookupMysqlInstanceResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) int { return v.Port }).(pulumi.IntOutput)
}

// Indicates the private IP address of the DB instance.
func (o LookupMysqlInstanceResultOutput) PrivateWriteIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.PrivateWriteIp }).(pulumi.StringOutput)
}

// Indicates the count of read replicas.
func (o LookupMysqlInstanceResultOutput) ReadReplicas() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) int { return v.ReadReplicas }).(pulumi.IntOutput)
}

func (o LookupMysqlInstanceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the security group ID.
func (o LookupMysqlInstanceResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Indicates the node status.
func (o LookupMysqlInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupMysqlInstanceResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// Indicates the time zone.
func (o LookupMysqlInstanceResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

func (o LookupMysqlInstanceResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMysqlInstanceResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMysqlInstanceResultOutput{})
}