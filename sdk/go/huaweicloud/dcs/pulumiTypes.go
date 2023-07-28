// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceBackupPolicy struct {
	// Day in a week on which backup starts, the value ranges from 1 to 7.
	// Where: 1 indicates Monday; 7 indicates Sunday.
	BackupAts []int `pulumi:"backupAts"`
	// Backup type. Default value is `auto`. The valid values are as follows:
	// + `auto`: automatic backup.
	// + `manual`: manual backup.
	BackupType *string `pulumi:"backupType"`
	// Time at which backup starts.
	// Format: `hh24:00-hh24:00`, "00:00-01:00" indicates that backup starts at 00:00:00.
	BeginAt string `pulumi:"beginAt"`
	// Interval at which backup is performed. Default value is `weekly`.
	// Currently, only weekly backup is supported.
	PeriodType *string `pulumi:"periodType"`
	// Retention time. Unit: day, the value ranges from 1 to 7.
	// This parameter is required if the backupType is **auto**.
	SaveDays *int `pulumi:"saveDays"`
}

// InstanceBackupPolicyInput is an input type that accepts InstanceBackupPolicyArgs and InstanceBackupPolicyOutput values.
// You can construct a concrete instance of `InstanceBackupPolicyInput` via:
//
//	InstanceBackupPolicyArgs{...}
type InstanceBackupPolicyInput interface {
	pulumi.Input

	ToInstanceBackupPolicyOutput() InstanceBackupPolicyOutput
	ToInstanceBackupPolicyOutputWithContext(context.Context) InstanceBackupPolicyOutput
}

type InstanceBackupPolicyArgs struct {
	// Day in a week on which backup starts, the value ranges from 1 to 7.
	// Where: 1 indicates Monday; 7 indicates Sunday.
	BackupAts pulumi.IntArrayInput `pulumi:"backupAts"`
	// Backup type. Default value is `auto`. The valid values are as follows:
	// + `auto`: automatic backup.
	// + `manual`: manual backup.
	BackupType pulumi.StringPtrInput `pulumi:"backupType"`
	// Time at which backup starts.
	// Format: `hh24:00-hh24:00`, "00:00-01:00" indicates that backup starts at 00:00:00.
	BeginAt pulumi.StringInput `pulumi:"beginAt"`
	// Interval at which backup is performed. Default value is `weekly`.
	// Currently, only weekly backup is supported.
	PeriodType pulumi.StringPtrInput `pulumi:"periodType"`
	// Retention time. Unit: day, the value ranges from 1 to 7.
	// This parameter is required if the backupType is **auto**.
	SaveDays pulumi.IntPtrInput `pulumi:"saveDays"`
}

func (InstanceBackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceBackupPolicy)(nil)).Elem()
}

func (i InstanceBackupPolicyArgs) ToInstanceBackupPolicyOutput() InstanceBackupPolicyOutput {
	return i.ToInstanceBackupPolicyOutputWithContext(context.Background())
}

func (i InstanceBackupPolicyArgs) ToInstanceBackupPolicyOutputWithContext(ctx context.Context) InstanceBackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupPolicyOutput)
}

func (i InstanceBackupPolicyArgs) ToInstanceBackupPolicyPtrOutput() InstanceBackupPolicyPtrOutput {
	return i.ToInstanceBackupPolicyPtrOutputWithContext(context.Background())
}

func (i InstanceBackupPolicyArgs) ToInstanceBackupPolicyPtrOutputWithContext(ctx context.Context) InstanceBackupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupPolicyOutput).ToInstanceBackupPolicyPtrOutputWithContext(ctx)
}

// InstanceBackupPolicyPtrInput is an input type that accepts InstanceBackupPolicyArgs, InstanceBackupPolicyPtr and InstanceBackupPolicyPtrOutput values.
// You can construct a concrete instance of `InstanceBackupPolicyPtrInput` via:
//
//	        InstanceBackupPolicyArgs{...}
//
//	or:
//
//	        nil
type InstanceBackupPolicyPtrInput interface {
	pulumi.Input

	ToInstanceBackupPolicyPtrOutput() InstanceBackupPolicyPtrOutput
	ToInstanceBackupPolicyPtrOutputWithContext(context.Context) InstanceBackupPolicyPtrOutput
}

type instanceBackupPolicyPtrType InstanceBackupPolicyArgs

func InstanceBackupPolicyPtr(v *InstanceBackupPolicyArgs) InstanceBackupPolicyPtrInput {
	return (*instanceBackupPolicyPtrType)(v)
}

func (*instanceBackupPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupPolicy)(nil)).Elem()
}

func (i *instanceBackupPolicyPtrType) ToInstanceBackupPolicyPtrOutput() InstanceBackupPolicyPtrOutput {
	return i.ToInstanceBackupPolicyPtrOutputWithContext(context.Background())
}

func (i *instanceBackupPolicyPtrType) ToInstanceBackupPolicyPtrOutputWithContext(ctx context.Context) InstanceBackupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupPolicyPtrOutput)
}

type InstanceBackupPolicyOutput struct{ *pulumi.OutputState }

func (InstanceBackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceBackupPolicy)(nil)).Elem()
}

func (o InstanceBackupPolicyOutput) ToInstanceBackupPolicyOutput() InstanceBackupPolicyOutput {
	return o
}

func (o InstanceBackupPolicyOutput) ToInstanceBackupPolicyOutputWithContext(ctx context.Context) InstanceBackupPolicyOutput {
	return o
}

func (o InstanceBackupPolicyOutput) ToInstanceBackupPolicyPtrOutput() InstanceBackupPolicyPtrOutput {
	return o.ToInstanceBackupPolicyPtrOutputWithContext(context.Background())
}

func (o InstanceBackupPolicyOutput) ToInstanceBackupPolicyPtrOutputWithContext(ctx context.Context) InstanceBackupPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceBackupPolicy) *InstanceBackupPolicy {
		return &v
	}).(InstanceBackupPolicyPtrOutput)
}

// Day in a week on which backup starts, the value ranges from 1 to 7.
// Where: 1 indicates Monday; 7 indicates Sunday.
func (o InstanceBackupPolicyOutput) BackupAts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v InstanceBackupPolicy) []int { return v.BackupAts }).(pulumi.IntArrayOutput)
}

// Backup type. Default value is `auto`. The valid values are as follows:
// + `auto`: automatic backup.
// + `manual`: manual backup.
func (o InstanceBackupPolicyOutput) BackupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceBackupPolicy) *string { return v.BackupType }).(pulumi.StringPtrOutput)
}

// Time at which backup starts.
// Format: `hh24:00-hh24:00`, "00:00-01:00" indicates that backup starts at 00:00:00.
func (o InstanceBackupPolicyOutput) BeginAt() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceBackupPolicy) string { return v.BeginAt }).(pulumi.StringOutput)
}

// Interval at which backup is performed. Default value is `weekly`.
// Currently, only weekly backup is supported.
func (o InstanceBackupPolicyOutput) PeriodType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceBackupPolicy) *string { return v.PeriodType }).(pulumi.StringPtrOutput)
}

// Retention time. Unit: day, the value ranges from 1 to 7.
// This parameter is required if the backupType is **auto**.
func (o InstanceBackupPolicyOutput) SaveDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceBackupPolicy) *int { return v.SaveDays }).(pulumi.IntPtrOutput)
}

type InstanceBackupPolicyPtrOutput struct{ *pulumi.OutputState }

func (InstanceBackupPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupPolicy)(nil)).Elem()
}

func (o InstanceBackupPolicyPtrOutput) ToInstanceBackupPolicyPtrOutput() InstanceBackupPolicyPtrOutput {
	return o
}

func (o InstanceBackupPolicyPtrOutput) ToInstanceBackupPolicyPtrOutputWithContext(ctx context.Context) InstanceBackupPolicyPtrOutput {
	return o
}

func (o InstanceBackupPolicyPtrOutput) Elem() InstanceBackupPolicyOutput {
	return o.ApplyT(func(v *InstanceBackupPolicy) InstanceBackupPolicy {
		if v != nil {
			return *v
		}
		var ret InstanceBackupPolicy
		return ret
	}).(InstanceBackupPolicyOutput)
}

// Day in a week on which backup starts, the value ranges from 1 to 7.
// Where: 1 indicates Monday; 7 indicates Sunday.
func (o InstanceBackupPolicyPtrOutput) BackupAts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *InstanceBackupPolicy) []int {
		if v == nil {
			return nil
		}
		return v.BackupAts
	}).(pulumi.IntArrayOutput)
}

// Backup type. Default value is `auto`. The valid values are as follows:
// + `auto`: automatic backup.
// + `manual`: manual backup.
func (o InstanceBackupPolicyPtrOutput) BackupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceBackupPolicy) *string {
		if v == nil {
			return nil
		}
		return v.BackupType
	}).(pulumi.StringPtrOutput)
}

// Time at which backup starts.
// Format: `hh24:00-hh24:00`, "00:00-01:00" indicates that backup starts at 00:00:00.
func (o InstanceBackupPolicyPtrOutput) BeginAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceBackupPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.BeginAt
	}).(pulumi.StringPtrOutput)
}

// Interval at which backup is performed. Default value is `weekly`.
// Currently, only weekly backup is supported.
func (o InstanceBackupPolicyPtrOutput) PeriodType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceBackupPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PeriodType
	}).(pulumi.StringPtrOutput)
}

// Retention time. Unit: day, the value ranges from 1 to 7.
// This parameter is required if the backupType is **auto**.
func (o InstanceBackupPolicyPtrOutput) SaveDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InstanceBackupPolicy) *int {
		if v == nil {
			return nil
		}
		return v.SaveDays
	}).(pulumi.IntPtrOutput)
}

type InstanceWhitelist struct {
	// Specifies the name of IP address group.
	GroupName string `pulumi:"groupName"`
	// Specifies the list of IP address or CIDR which can be whitelisted for an instance.
	// The maximum is 20.
	IpAddresses []string `pulumi:"ipAddresses"`
}

// InstanceWhitelistInput is an input type that accepts InstanceWhitelistArgs and InstanceWhitelistOutput values.
// You can construct a concrete instance of `InstanceWhitelistInput` via:
//
//	InstanceWhitelistArgs{...}
type InstanceWhitelistInput interface {
	pulumi.Input

	ToInstanceWhitelistOutput() InstanceWhitelistOutput
	ToInstanceWhitelistOutputWithContext(context.Context) InstanceWhitelistOutput
}

type InstanceWhitelistArgs struct {
	// Specifies the name of IP address group.
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// Specifies the list of IP address or CIDR which can be whitelisted for an instance.
	// The maximum is 20.
	IpAddresses pulumi.StringArrayInput `pulumi:"ipAddresses"`
}

func (InstanceWhitelistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceWhitelist)(nil)).Elem()
}

func (i InstanceWhitelistArgs) ToInstanceWhitelistOutput() InstanceWhitelistOutput {
	return i.ToInstanceWhitelistOutputWithContext(context.Background())
}

func (i InstanceWhitelistArgs) ToInstanceWhitelistOutputWithContext(ctx context.Context) InstanceWhitelistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceWhitelistOutput)
}

// InstanceWhitelistArrayInput is an input type that accepts InstanceWhitelistArray and InstanceWhitelistArrayOutput values.
// You can construct a concrete instance of `InstanceWhitelistArrayInput` via:
//
//	InstanceWhitelistArray{ InstanceWhitelistArgs{...} }
type InstanceWhitelistArrayInput interface {
	pulumi.Input

	ToInstanceWhitelistArrayOutput() InstanceWhitelistArrayOutput
	ToInstanceWhitelistArrayOutputWithContext(context.Context) InstanceWhitelistArrayOutput
}

type InstanceWhitelistArray []InstanceWhitelistInput

func (InstanceWhitelistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceWhitelist)(nil)).Elem()
}

func (i InstanceWhitelistArray) ToInstanceWhitelistArrayOutput() InstanceWhitelistArrayOutput {
	return i.ToInstanceWhitelistArrayOutputWithContext(context.Background())
}

func (i InstanceWhitelistArray) ToInstanceWhitelistArrayOutputWithContext(ctx context.Context) InstanceWhitelistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceWhitelistArrayOutput)
}

type InstanceWhitelistOutput struct{ *pulumi.OutputState }

func (InstanceWhitelistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceWhitelist)(nil)).Elem()
}

func (o InstanceWhitelistOutput) ToInstanceWhitelistOutput() InstanceWhitelistOutput {
	return o
}

func (o InstanceWhitelistOutput) ToInstanceWhitelistOutputWithContext(ctx context.Context) InstanceWhitelistOutput {
	return o
}

// Specifies the name of IP address group.
func (o InstanceWhitelistOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceWhitelist) string { return v.GroupName }).(pulumi.StringOutput)
}

// Specifies the list of IP address or CIDR which can be whitelisted for an instance.
// The maximum is 20.
func (o InstanceWhitelistOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceWhitelist) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

type InstanceWhitelistArrayOutput struct{ *pulumi.OutputState }

func (InstanceWhitelistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceWhitelist)(nil)).Elem()
}

func (o InstanceWhitelistArrayOutput) ToInstanceWhitelistArrayOutput() InstanceWhitelistArrayOutput {
	return o
}

func (o InstanceWhitelistArrayOutput) ToInstanceWhitelistArrayOutputWithContext(ctx context.Context) InstanceWhitelistArrayOutput {
	return o
}

func (o InstanceWhitelistArrayOutput) Index(i pulumi.IntInput) InstanceWhitelistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceWhitelist {
		return vs[0].([]InstanceWhitelist)[vs[1].(int)]
	}).(InstanceWhitelistOutput)
}

type GetFlavorsFlavor struct {
	// An array of available zones where the cache specification can be used.
	AvailableZones []string `pulumi:"availableZones"`
	// The mode of a cache engine. The valid values are as follows:
	CacheMode string `pulumi:"cacheMode"`
	// The total memory of the cache, in GB.
	// + **Redis4.0, Redis5.0 and Redis6.0**: Stand-alone and active/standby type instance values:
	//   `0.125`, `0.25`, `0.5`, `1`, `2`, `4`, `8`, `16`, `32` and `64`.
	//   Cluster instance specifications support `4`,`8`,`16`,`24`, `32`, `48`, `64`, `96`, `128`, `192`,
	//   `256`, `384`, `512`, `768` and `1024`.
	// + **Redis3.0**: Stand-alone and active/standby type instance values: `2`, `4`, `8`, `16`, `32` and `64`.
	//   Proxy cluster instance specifications support `64`, `128`, `256`, `512`, and `1024`.
	// + **Memcached**: Stand-alone and active/standby type instance values: `2`, `4`, `8`, `16`, `32` and `64`.
	Capacity float64 `pulumi:"capacity"`
	// The charging modes for the specification cache instance.
	ChargingModes []string `pulumi:"chargingModes"`
	// The CPU architecture of cache instance.
	// Valid values *x86_64* and *aarch64*.
	CpuArchitecture string `pulumi:"cpuArchitecture"`
	// The engine of the cache instance. Valid values are *Redis* and *Memcached*.
	// Default value is *Redis*.
	Engine string `pulumi:"engine"`
	// Supported versions of the specification.
	EngineVersions string `pulumi:"engineVersions"`
	// Number of IP addresses corresponding to the specifications.
	IpCount int `pulumi:"ipCount"`
	// The flavor name of the cache instance.
	Name string `pulumi:"name"`
}

// GetFlavorsFlavorInput is an input type that accepts GetFlavorsFlavorArgs and GetFlavorsFlavorOutput values.
// You can construct a concrete instance of `GetFlavorsFlavorInput` via:
//
//	GetFlavorsFlavorArgs{...}
type GetFlavorsFlavorInput interface {
	pulumi.Input

	ToGetFlavorsFlavorOutput() GetFlavorsFlavorOutput
	ToGetFlavorsFlavorOutputWithContext(context.Context) GetFlavorsFlavorOutput
}

type GetFlavorsFlavorArgs struct {
	// An array of available zones where the cache specification can be used.
	AvailableZones pulumi.StringArrayInput `pulumi:"availableZones"`
	// The mode of a cache engine. The valid values are as follows:
	CacheMode pulumi.StringInput `pulumi:"cacheMode"`
	// The total memory of the cache, in GB.
	// + **Redis4.0, Redis5.0 and Redis6.0**: Stand-alone and active/standby type instance values:
	//   `0.125`, `0.25`, `0.5`, `1`, `2`, `4`, `8`, `16`, `32` and `64`.
	//   Cluster instance specifications support `4`,`8`,`16`,`24`, `32`, `48`, `64`, `96`, `128`, `192`,
	//   `256`, `384`, `512`, `768` and `1024`.
	// + **Redis3.0**: Stand-alone and active/standby type instance values: `2`, `4`, `8`, `16`, `32` and `64`.
	//   Proxy cluster instance specifications support `64`, `128`, `256`, `512`, and `1024`.
	// + **Memcached**: Stand-alone and active/standby type instance values: `2`, `4`, `8`, `16`, `32` and `64`.
	Capacity pulumi.Float64Input `pulumi:"capacity"`
	// The charging modes for the specification cache instance.
	ChargingModes pulumi.StringArrayInput `pulumi:"chargingModes"`
	// The CPU architecture of cache instance.
	// Valid values *x86_64* and *aarch64*.
	CpuArchitecture pulumi.StringInput `pulumi:"cpuArchitecture"`
	// The engine of the cache instance. Valid values are *Redis* and *Memcached*.
	// Default value is *Redis*.
	Engine pulumi.StringInput `pulumi:"engine"`
	// Supported versions of the specification.
	EngineVersions pulumi.StringInput `pulumi:"engineVersions"`
	// Number of IP addresses corresponding to the specifications.
	IpCount pulumi.IntInput `pulumi:"ipCount"`
	// The flavor name of the cache instance.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetFlavorsFlavorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlavorsFlavor)(nil)).Elem()
}

func (i GetFlavorsFlavorArgs) ToGetFlavorsFlavorOutput() GetFlavorsFlavorOutput {
	return i.ToGetFlavorsFlavorOutputWithContext(context.Background())
}

func (i GetFlavorsFlavorArgs) ToGetFlavorsFlavorOutputWithContext(ctx context.Context) GetFlavorsFlavorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFlavorsFlavorOutput)
}

// GetFlavorsFlavorArrayInput is an input type that accepts GetFlavorsFlavorArray and GetFlavorsFlavorArrayOutput values.
// You can construct a concrete instance of `GetFlavorsFlavorArrayInput` via:
//
//	GetFlavorsFlavorArray{ GetFlavorsFlavorArgs{...} }
type GetFlavorsFlavorArrayInput interface {
	pulumi.Input

	ToGetFlavorsFlavorArrayOutput() GetFlavorsFlavorArrayOutput
	ToGetFlavorsFlavorArrayOutputWithContext(context.Context) GetFlavorsFlavorArrayOutput
}

type GetFlavorsFlavorArray []GetFlavorsFlavorInput

func (GetFlavorsFlavorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFlavorsFlavor)(nil)).Elem()
}

func (i GetFlavorsFlavorArray) ToGetFlavorsFlavorArrayOutput() GetFlavorsFlavorArrayOutput {
	return i.ToGetFlavorsFlavorArrayOutputWithContext(context.Background())
}

func (i GetFlavorsFlavorArray) ToGetFlavorsFlavorArrayOutputWithContext(ctx context.Context) GetFlavorsFlavorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFlavorsFlavorArrayOutput)
}

type GetFlavorsFlavorOutput struct{ *pulumi.OutputState }

func (GetFlavorsFlavorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlavorsFlavor)(nil)).Elem()
}

func (o GetFlavorsFlavorOutput) ToGetFlavorsFlavorOutput() GetFlavorsFlavorOutput {
	return o
}

func (o GetFlavorsFlavorOutput) ToGetFlavorsFlavorOutputWithContext(ctx context.Context) GetFlavorsFlavorOutput {
	return o
}

// An array of available zones where the cache specification can be used.
func (o GetFlavorsFlavorOutput) AvailableZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) []string { return v.AvailableZones }).(pulumi.StringArrayOutput)
}

// The mode of a cache engine. The valid values are as follows:
func (o GetFlavorsFlavorOutput) CacheMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) string { return v.CacheMode }).(pulumi.StringOutput)
}

// The total memory of the cache, in GB.
//   - **Redis4.0, Redis5.0 and Redis6.0**: Stand-alone and active/standby type instance values:
//     `0.125`, `0.25`, `0.5`, `1`, `2`, `4`, `8`, `16`, `32` and `64`.
//     Cluster instance specifications support `4`,`8`,`16`,`24`, `32`, `48`, `64`, `96`, `128`, `192`,
//     `256`, `384`, `512`, `768` and `1024`.
//   - **Redis3.0**: Stand-alone and active/standby type instance values: `2`, `4`, `8`, `16`, `32` and `64`.
//     Proxy cluster instance specifications support `64`, `128`, `256`, `512`, and `1024`.
//   - **Memcached**: Stand-alone and active/standby type instance values: `2`, `4`, `8`, `16`, `32` and `64`.
func (o GetFlavorsFlavorOutput) Capacity() pulumi.Float64Output {
	return o.ApplyT(func(v GetFlavorsFlavor) float64 { return v.Capacity }).(pulumi.Float64Output)
}

// The charging modes for the specification cache instance.
func (o GetFlavorsFlavorOutput) ChargingModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) []string { return v.ChargingModes }).(pulumi.StringArrayOutput)
}

// The CPU architecture of cache instance.
// Valid values *x86_64* and *aarch64*.
func (o GetFlavorsFlavorOutput) CpuArchitecture() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) string { return v.CpuArchitecture }).(pulumi.StringOutput)
}

// The engine of the cache instance. Valid values are *Redis* and *Memcached*.
// Default value is *Redis*.
func (o GetFlavorsFlavorOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) string { return v.Engine }).(pulumi.StringOutput)
}

// Supported versions of the specification.
func (o GetFlavorsFlavorOutput) EngineVersions() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) string { return v.EngineVersions }).(pulumi.StringOutput)
}

// Number of IP addresses corresponding to the specifications.
func (o GetFlavorsFlavorOutput) IpCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) int { return v.IpCount }).(pulumi.IntOutput)
}

// The flavor name of the cache instance.
func (o GetFlavorsFlavorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) string { return v.Name }).(pulumi.StringOutput)
}

type GetFlavorsFlavorArrayOutput struct{ *pulumi.OutputState }

func (GetFlavorsFlavorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFlavorsFlavor)(nil)).Elem()
}

func (o GetFlavorsFlavorArrayOutput) ToGetFlavorsFlavorArrayOutput() GetFlavorsFlavorArrayOutput {
	return o
}

func (o GetFlavorsFlavorArrayOutput) ToGetFlavorsFlavorArrayOutputWithContext(ctx context.Context) GetFlavorsFlavorArrayOutput {
	return o
}

func (o GetFlavorsFlavorArrayOutput) Index(i pulumi.IntInput) GetFlavorsFlavorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetFlavorsFlavor {
		return vs[0].([]GetFlavorsFlavor)[vs[1].(int)]
	}).(GetFlavorsFlavorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupPolicyInput)(nil)).Elem(), InstanceBackupPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupPolicyPtrInput)(nil)).Elem(), InstanceBackupPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceWhitelistInput)(nil)).Elem(), InstanceWhitelistArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceWhitelistArrayInput)(nil)).Elem(), InstanceWhitelistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFlavorsFlavorInput)(nil)).Elem(), GetFlavorsFlavorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFlavorsFlavorArrayInput)(nil)).Elem(), GetFlavorsFlavorArray{})
	pulumi.RegisterOutputType(InstanceBackupPolicyOutput{})
	pulumi.RegisterOutputType(InstanceBackupPolicyPtrOutput{})
	pulumi.RegisterOutputType(InstanceWhitelistOutput{})
	pulumi.RegisterOutputType(InstanceWhitelistArrayOutput{})
	pulumi.RegisterOutputType(GetFlavorsFlavorOutput{})
	pulumi.RegisterOutputType(GetFlavorsFlavorArrayOutput{})
}
