// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterInstance struct {
	// Instance ID.
	Id *string `pulumi:"id"`
	// Management IP address.
	ManageIp *string `pulumi:"manageIp"`
	// Specifies cluster name. Changing this parameter will create a new resource.
	Name *string `pulumi:"name"`
	// Private IP.
	PrivateIp *string `pulumi:"privateIp"`
	// Public IP.
	PublicIp *string `pulumi:"publicIp"`
	// Instance role.
	Role *string `pulumi:"role"`
	// Traffic IP.
	TrafficIp *string `pulumi:"trafficIp"`
	// Instance type.
	Type *string `pulumi:"type"`
}

// ClusterInstanceInput is an input type that accepts ClusterInstanceArgs and ClusterInstanceOutput values.
// You can construct a concrete instance of `ClusterInstanceInput` via:
//
//	ClusterInstanceArgs{...}
type ClusterInstanceInput interface {
	pulumi.Input

	ToClusterInstanceOutput() ClusterInstanceOutput
	ToClusterInstanceOutputWithContext(context.Context) ClusterInstanceOutput
}

type ClusterInstanceArgs struct {
	// Instance ID.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Management IP address.
	ManageIp pulumi.StringPtrInput `pulumi:"manageIp"`
	// Specifies cluster name. Changing this parameter will create a new resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Private IP.
	PrivateIp pulumi.StringPtrInput `pulumi:"privateIp"`
	// Public IP.
	PublicIp pulumi.StringPtrInput `pulumi:"publicIp"`
	// Instance role.
	Role pulumi.StringPtrInput `pulumi:"role"`
	// Traffic IP.
	TrafficIp pulumi.StringPtrInput `pulumi:"trafficIp"`
	// Instance type.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ClusterInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInstance)(nil)).Elem()
}

func (i ClusterInstanceArgs) ToClusterInstanceOutput() ClusterInstanceOutput {
	return i.ToClusterInstanceOutputWithContext(context.Background())
}

func (i ClusterInstanceArgs) ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInstanceOutput)
}

// ClusterInstanceArrayInput is an input type that accepts ClusterInstanceArray and ClusterInstanceArrayOutput values.
// You can construct a concrete instance of `ClusterInstanceArrayInput` via:
//
//	ClusterInstanceArray{ ClusterInstanceArgs{...} }
type ClusterInstanceArrayInput interface {
	pulumi.Input

	ToClusterInstanceArrayOutput() ClusterInstanceArrayOutput
	ToClusterInstanceArrayOutputWithContext(context.Context) ClusterInstanceArrayOutput
}

type ClusterInstanceArray []ClusterInstanceInput

func (ClusterInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterInstance)(nil)).Elem()
}

func (i ClusterInstanceArray) ToClusterInstanceArrayOutput() ClusterInstanceArrayOutput {
	return i.ToClusterInstanceArrayOutputWithContext(context.Background())
}

func (i ClusterInstanceArray) ToClusterInstanceArrayOutputWithContext(ctx context.Context) ClusterInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInstanceArrayOutput)
}

type ClusterInstanceOutput struct{ *pulumi.OutputState }

func (ClusterInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInstance)(nil)).Elem()
}

func (o ClusterInstanceOutput) ToClusterInstanceOutput() ClusterInstanceOutput {
	return o
}

func (o ClusterInstanceOutput) ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput {
	return o
}

// Instance ID.
func (o ClusterInstanceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Management IP address.
func (o ClusterInstanceOutput) ManageIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.ManageIp }).(pulumi.StringPtrOutput)
}

// Specifies cluster name. Changing this parameter will create a new resource.
func (o ClusterInstanceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Private IP.
func (o ClusterInstanceOutput) PrivateIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.PrivateIp }).(pulumi.StringPtrOutput)
}

// Public IP.
func (o ClusterInstanceOutput) PublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.PublicIp }).(pulumi.StringPtrOutput)
}

// Instance role.
func (o ClusterInstanceOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.Role }).(pulumi.StringPtrOutput)
}

// Traffic IP.
func (o ClusterInstanceOutput) TrafficIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.TrafficIp }).(pulumi.StringPtrOutput)
}

// Instance type.
func (o ClusterInstanceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInstance) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ClusterInstanceArrayOutput struct{ *pulumi.OutputState }

func (ClusterInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterInstance)(nil)).Elem()
}

func (o ClusterInstanceArrayOutput) ToClusterInstanceArrayOutput() ClusterInstanceArrayOutput {
	return o
}

func (o ClusterInstanceArrayOutput) ToClusterInstanceArrayOutputWithContext(ctx context.Context) ClusterInstanceArrayOutput {
	return o
}

func (o ClusterInstanceArrayOutput) Index(i pulumi.IntInput) ClusterInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterInstance {
		return vs[0].([]ClusterInstance)[vs[1].(int)]
	}).(ClusterInstanceOutput)
}

type JobConfig struct {
	// Specifies group to which a job belongs. The default group is `DEFAULT`.
	GroupName *string `pulumi:"groupName"`
	// Specifies whether to automatically retry if a job fails to be executed.
	// The options are as follows:
	// + **NONE**: Do not retry.
	// + **RETRY_TRIPLE**: Retry three times.
	RetryType *string `pulumi:"retryType"`
	// Specifies cycle of a scheduled task. If `schedulerCycleType` is set to minute
	// and `schedulerCycle` is set to 10, the scheduled task is executed every 10 minutes.
	SchedulerCycle *int `pulumi:"schedulerCycle"`
	// Specifies cycle type of a scheduled task. The options are as follows:
	// `minute`, `hour`, `day`, `week`, `month`.
	SchedulerCycleType *string `pulumi:"schedulerCycleType"`
	// Specifies whether to delete a job after the job is executed.
	// The options are as follows:
	// + **NONE**: The job will not be deleted after it is executed.
	// + **DELETE_AFTER_SUCCEED**: The job will be deleted only after it is successfully executed. It is applicable to
	//   massive one-time jobs.
	// + **DELETE**: Thejob will be deleted after it is executed, regardless of the execution result.
	SchedulerDisposableType *string `pulumi:"schedulerDisposableType"`
	// Specifies whether to enable a scheduled task.  Default value is `false`.
	SchedulerEnabled *bool `pulumi:"schedulerEnabled"`
	// Specifies time when a scheduled task is triggered in a cycle. This parameter
	// is valid only when `schedulerCycleType` is set to `hour`, `week`, or `month`.
	// + If `schedulerCycleType` is set to month, cycle is set to 1, and runAt is set to 15, the scheduled task is executed
	//   on the 15th day of each month. You can set runAt to multiple values and separate the values with commas (,).
	//   For example, if runAt is set to 1,2,3,4,5, the scheduled task is executed on the first day, second day, third day,
	//   fourth day, and fifth day of each month.
	// + If `schedulerCycleType` is set to week and runAt is set to mon,tue,wed,thu,fri, the scheduled task is executed on
	//   Monday to Friday.
	// + If `schedulerCycleType` is set to hour and runAt is set to 27,57, the scheduled task is executed at the 27th and
	//   57th minute in the cycle.
	SchedulerRunAt *string `pulumi:"schedulerRunAt"`
	// Specifies start time of a scheduled task.
	// For example, `2018-01-24 19:56:19`
	SchedulerStartDate *string `pulumi:"schedulerStartDate"`
	// Specifies End time of a scheduled task. For example, `2018-01-27 23:59:00`.
	// If you do not set the end time, the scheduled task is always executed and will never stop.
	SchedulerStopDate *string `pulumi:"schedulerStopDate"`
	// Specifies name of the OBS bucket to which dirty data is
	// written. This parameter is valid only when dirty data is written to `OBS`.
	ThrottlingDirtyWriteToBucket *string `pulumi:"throttlingDirtyWriteToBucket"`
	// Specifies the directory in the OBS bucket or HDFS which
	// dirty data is written to. For example, `/data/dirtydata/`.
	ThrottlingDirtyWriteToDirectory *string `pulumi:"throttlingDirtyWriteToDirectory"`
	// Specifies the link name to which dirty data is written to.
	// The Dirty data can be written only to `OBS` or `HDFS`.
	ThrottlingDirtyWriteToLink *string `pulumi:"throttlingDirtyWriteToLink"`
	// Specifies maximum number of concurrent extraction jobs.
	ThrottlingExtractorsNumber *int `pulumi:"throttlingExtractorsNumber"`
	// Specifies maximum number of loading jobs. This parameter is available
	// only when HBase or Hive serves as the destination data source.
	ThrottlingLoaderNumber *int `pulumi:"throttlingLoaderNumber"`
	// Specifies maximum number of error records in a single
	// shard. When the number of error records of a map exceeds the upper limit, the task automatically ends.
	ThrottlingMaxErrorRecords *int `pulumi:"throttlingMaxErrorRecords"`
	// Specifies whether to write dirty data.
	ThrottlingRecordDirtyData *bool `pulumi:"throttlingRecordDirtyData"`
}

// JobConfigInput is an input type that accepts JobConfigArgs and JobConfigOutput values.
// You can construct a concrete instance of `JobConfigInput` via:
//
//	JobConfigArgs{...}
type JobConfigInput interface {
	pulumi.Input

	ToJobConfigOutput() JobConfigOutput
	ToJobConfigOutputWithContext(context.Context) JobConfigOutput
}

type JobConfigArgs struct {
	// Specifies group to which a job belongs. The default group is `DEFAULT`.
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
	// Specifies whether to automatically retry if a job fails to be executed.
	// The options are as follows:
	// + **NONE**: Do not retry.
	// + **RETRY_TRIPLE**: Retry three times.
	RetryType pulumi.StringPtrInput `pulumi:"retryType"`
	// Specifies cycle of a scheduled task. If `schedulerCycleType` is set to minute
	// and `schedulerCycle` is set to 10, the scheduled task is executed every 10 minutes.
	SchedulerCycle pulumi.IntPtrInput `pulumi:"schedulerCycle"`
	// Specifies cycle type of a scheduled task. The options are as follows:
	// `minute`, `hour`, `day`, `week`, `month`.
	SchedulerCycleType pulumi.StringPtrInput `pulumi:"schedulerCycleType"`
	// Specifies whether to delete a job after the job is executed.
	// The options are as follows:
	// + **NONE**: The job will not be deleted after it is executed.
	// + **DELETE_AFTER_SUCCEED**: The job will be deleted only after it is successfully executed. It is applicable to
	//   massive one-time jobs.
	// + **DELETE**: Thejob will be deleted after it is executed, regardless of the execution result.
	SchedulerDisposableType pulumi.StringPtrInput `pulumi:"schedulerDisposableType"`
	// Specifies whether to enable a scheduled task.  Default value is `false`.
	SchedulerEnabled pulumi.BoolPtrInput `pulumi:"schedulerEnabled"`
	// Specifies time when a scheduled task is triggered in a cycle. This parameter
	// is valid only when `schedulerCycleType` is set to `hour`, `week`, or `month`.
	// + If `schedulerCycleType` is set to month, cycle is set to 1, and runAt is set to 15, the scheduled task is executed
	//   on the 15th day of each month. You can set runAt to multiple values and separate the values with commas (,).
	//   For example, if runAt is set to 1,2,3,4,5, the scheduled task is executed on the first day, second day, third day,
	//   fourth day, and fifth day of each month.
	// + If `schedulerCycleType` is set to week and runAt is set to mon,tue,wed,thu,fri, the scheduled task is executed on
	//   Monday to Friday.
	// + If `schedulerCycleType` is set to hour and runAt is set to 27,57, the scheduled task is executed at the 27th and
	//   57th minute in the cycle.
	SchedulerRunAt pulumi.StringPtrInput `pulumi:"schedulerRunAt"`
	// Specifies start time of a scheduled task.
	// For example, `2018-01-24 19:56:19`
	SchedulerStartDate pulumi.StringPtrInput `pulumi:"schedulerStartDate"`
	// Specifies End time of a scheduled task. For example, `2018-01-27 23:59:00`.
	// If you do not set the end time, the scheduled task is always executed and will never stop.
	SchedulerStopDate pulumi.StringPtrInput `pulumi:"schedulerStopDate"`
	// Specifies name of the OBS bucket to which dirty data is
	// written. This parameter is valid only when dirty data is written to `OBS`.
	ThrottlingDirtyWriteToBucket pulumi.StringPtrInput `pulumi:"throttlingDirtyWriteToBucket"`
	// Specifies the directory in the OBS bucket or HDFS which
	// dirty data is written to. For example, `/data/dirtydata/`.
	ThrottlingDirtyWriteToDirectory pulumi.StringPtrInput `pulumi:"throttlingDirtyWriteToDirectory"`
	// Specifies the link name to which dirty data is written to.
	// The Dirty data can be written only to `OBS` or `HDFS`.
	ThrottlingDirtyWriteToLink pulumi.StringPtrInput `pulumi:"throttlingDirtyWriteToLink"`
	// Specifies maximum number of concurrent extraction jobs.
	ThrottlingExtractorsNumber pulumi.IntPtrInput `pulumi:"throttlingExtractorsNumber"`
	// Specifies maximum number of loading jobs. This parameter is available
	// only when HBase or Hive serves as the destination data source.
	ThrottlingLoaderNumber pulumi.IntPtrInput `pulumi:"throttlingLoaderNumber"`
	// Specifies maximum number of error records in a single
	// shard. When the number of error records of a map exceeds the upper limit, the task automatically ends.
	ThrottlingMaxErrorRecords pulumi.IntPtrInput `pulumi:"throttlingMaxErrorRecords"`
	// Specifies whether to write dirty data.
	ThrottlingRecordDirtyData pulumi.BoolPtrInput `pulumi:"throttlingRecordDirtyData"`
}

func (JobConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobConfig)(nil)).Elem()
}

func (i JobConfigArgs) ToJobConfigOutput() JobConfigOutput {
	return i.ToJobConfigOutputWithContext(context.Background())
}

func (i JobConfigArgs) ToJobConfigOutputWithContext(ctx context.Context) JobConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobConfigOutput)
}

func (i JobConfigArgs) ToJobConfigPtrOutput() JobConfigPtrOutput {
	return i.ToJobConfigPtrOutputWithContext(context.Background())
}

func (i JobConfigArgs) ToJobConfigPtrOutputWithContext(ctx context.Context) JobConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobConfigOutput).ToJobConfigPtrOutputWithContext(ctx)
}

// JobConfigPtrInput is an input type that accepts JobConfigArgs, JobConfigPtr and JobConfigPtrOutput values.
// You can construct a concrete instance of `JobConfigPtrInput` via:
//
//	        JobConfigArgs{...}
//
//	or:
//
//	        nil
type JobConfigPtrInput interface {
	pulumi.Input

	ToJobConfigPtrOutput() JobConfigPtrOutput
	ToJobConfigPtrOutputWithContext(context.Context) JobConfigPtrOutput
}

type jobConfigPtrType JobConfigArgs

func JobConfigPtr(v *JobConfigArgs) JobConfigPtrInput {
	return (*jobConfigPtrType)(v)
}

func (*jobConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobConfig)(nil)).Elem()
}

func (i *jobConfigPtrType) ToJobConfigPtrOutput() JobConfigPtrOutput {
	return i.ToJobConfigPtrOutputWithContext(context.Background())
}

func (i *jobConfigPtrType) ToJobConfigPtrOutputWithContext(ctx context.Context) JobConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobConfigPtrOutput)
}

type JobConfigOutput struct{ *pulumi.OutputState }

func (JobConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobConfig)(nil)).Elem()
}

func (o JobConfigOutput) ToJobConfigOutput() JobConfigOutput {
	return o
}

func (o JobConfigOutput) ToJobConfigOutputWithContext(ctx context.Context) JobConfigOutput {
	return o
}

func (o JobConfigOutput) ToJobConfigPtrOutput() JobConfigPtrOutput {
	return o.ToJobConfigPtrOutputWithContext(context.Background())
}

func (o JobConfigOutput) ToJobConfigPtrOutputWithContext(ctx context.Context) JobConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobConfig) *JobConfig {
		return &v
	}).(JobConfigPtrOutput)
}

// Specifies group to which a job belongs. The default group is `DEFAULT`.
func (o JobConfigOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

// Specifies whether to automatically retry if a job fails to be executed.
// The options are as follows:
// + **NONE**: Do not retry.
// + **RETRY_TRIPLE**: Retry three times.
func (o JobConfigOutput) RetryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.RetryType }).(pulumi.StringPtrOutput)
}

// Specifies cycle of a scheduled task. If `schedulerCycleType` is set to minute
// and `schedulerCycle` is set to 10, the scheduled task is executed every 10 minutes.
func (o JobConfigOutput) SchedulerCycle() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobConfig) *int { return v.SchedulerCycle }).(pulumi.IntPtrOutput)
}

// Specifies cycle type of a scheduled task. The options are as follows:
// `minute`, `hour`, `day`, `week`, `month`.
func (o JobConfigOutput) SchedulerCycleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.SchedulerCycleType }).(pulumi.StringPtrOutput)
}

// Specifies whether to delete a job after the job is executed.
// The options are as follows:
//   - **NONE**: The job will not be deleted after it is executed.
//   - **DELETE_AFTER_SUCCEED**: The job will be deleted only after it is successfully executed. It is applicable to
//     massive one-time jobs.
//   - **DELETE**: Thejob will be deleted after it is executed, regardless of the execution result.
func (o JobConfigOutput) SchedulerDisposableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.SchedulerDisposableType }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable a scheduled task.  Default value is `false`.
func (o JobConfigOutput) SchedulerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobConfig) *bool { return v.SchedulerEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies time when a scheduled task is triggered in a cycle. This parameter
// is valid only when `schedulerCycleType` is set to `hour`, `week`, or `month`.
//   - If `schedulerCycleType` is set to month, cycle is set to 1, and runAt is set to 15, the scheduled task is executed
//     on the 15th day of each month. You can set runAt to multiple values and separate the values with commas (,).
//     For example, if runAt is set to 1,2,3,4,5, the scheduled task is executed on the first day, second day, third day,
//     fourth day, and fifth day of each month.
//   - If `schedulerCycleType` is set to week and runAt is set to mon,tue,wed,thu,fri, the scheduled task is executed on
//     Monday to Friday.
//   - If `schedulerCycleType` is set to hour and runAt is set to 27,57, the scheduled task is executed at the 27th and
//     57th minute in the cycle.
func (o JobConfigOutput) SchedulerRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.SchedulerRunAt }).(pulumi.StringPtrOutput)
}

// Specifies start time of a scheduled task.
// For example, `2018-01-24 19:56:19`
func (o JobConfigOutput) SchedulerStartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.SchedulerStartDate }).(pulumi.StringPtrOutput)
}

// Specifies End time of a scheduled task. For example, `2018-01-27 23:59:00`.
// If you do not set the end time, the scheduled task is always executed and will never stop.
func (o JobConfigOutput) SchedulerStopDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.SchedulerStopDate }).(pulumi.StringPtrOutput)
}

// Specifies name of the OBS bucket to which dirty data is
// written. This parameter is valid only when dirty data is written to `OBS`.
func (o JobConfigOutput) ThrottlingDirtyWriteToBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.ThrottlingDirtyWriteToBucket }).(pulumi.StringPtrOutput)
}

// Specifies the directory in the OBS bucket or HDFS which
// dirty data is written to. For example, `/data/dirtydata/`.
func (o JobConfigOutput) ThrottlingDirtyWriteToDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.ThrottlingDirtyWriteToDirectory }).(pulumi.StringPtrOutput)
}

// Specifies the link name to which dirty data is written to.
// The Dirty data can be written only to `OBS` or `HDFS`.
func (o JobConfigOutput) ThrottlingDirtyWriteToLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobConfig) *string { return v.ThrottlingDirtyWriteToLink }).(pulumi.StringPtrOutput)
}

// Specifies maximum number of concurrent extraction jobs.
func (o JobConfigOutput) ThrottlingExtractorsNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobConfig) *int { return v.ThrottlingExtractorsNumber }).(pulumi.IntPtrOutput)
}

// Specifies maximum number of loading jobs. This parameter is available
// only when HBase or Hive serves as the destination data source.
func (o JobConfigOutput) ThrottlingLoaderNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobConfig) *int { return v.ThrottlingLoaderNumber }).(pulumi.IntPtrOutput)
}

// Specifies maximum number of error records in a single
// shard. When the number of error records of a map exceeds the upper limit, the task automatically ends.
func (o JobConfigOutput) ThrottlingMaxErrorRecords() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobConfig) *int { return v.ThrottlingMaxErrorRecords }).(pulumi.IntPtrOutput)
}

// Specifies whether to write dirty data.
func (o JobConfigOutput) ThrottlingRecordDirtyData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobConfig) *bool { return v.ThrottlingRecordDirtyData }).(pulumi.BoolPtrOutput)
}

type JobConfigPtrOutput struct{ *pulumi.OutputState }

func (JobConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobConfig)(nil)).Elem()
}

func (o JobConfigPtrOutput) ToJobConfigPtrOutput() JobConfigPtrOutput {
	return o
}

func (o JobConfigPtrOutput) ToJobConfigPtrOutputWithContext(ctx context.Context) JobConfigPtrOutput {
	return o
}

func (o JobConfigPtrOutput) Elem() JobConfigOutput {
	return o.ApplyT(func(v *JobConfig) JobConfig {
		if v != nil {
			return *v
		}
		var ret JobConfig
		return ret
	}).(JobConfigOutput)
}

// Specifies group to which a job belongs. The default group is `DEFAULT`.
func (o JobConfigPtrOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.GroupName
	}).(pulumi.StringPtrOutput)
}

// Specifies whether to automatically retry if a job fails to be executed.
// The options are as follows:
// + **NONE**: Do not retry.
// + **RETRY_TRIPLE**: Retry three times.
func (o JobConfigPtrOutput) RetryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.RetryType
	}).(pulumi.StringPtrOutput)
}

// Specifies cycle of a scheduled task. If `schedulerCycleType` is set to minute
// and `schedulerCycle` is set to 10, the scheduled task is executed every 10 minutes.
func (o JobConfigPtrOutput) SchedulerCycle() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobConfig) *int {
		if v == nil {
			return nil
		}
		return v.SchedulerCycle
	}).(pulumi.IntPtrOutput)
}

// Specifies cycle type of a scheduled task. The options are as follows:
// `minute`, `hour`, `day`, `week`, `month`.
func (o JobConfigPtrOutput) SchedulerCycleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.SchedulerCycleType
	}).(pulumi.StringPtrOutput)
}

// Specifies whether to delete a job after the job is executed.
// The options are as follows:
//   - **NONE**: The job will not be deleted after it is executed.
//   - **DELETE_AFTER_SUCCEED**: The job will be deleted only after it is successfully executed. It is applicable to
//     massive one-time jobs.
//   - **DELETE**: Thejob will be deleted after it is executed, regardless of the execution result.
func (o JobConfigPtrOutput) SchedulerDisposableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.SchedulerDisposableType
	}).(pulumi.StringPtrOutput)
}

// Specifies whether to enable a scheduled task.  Default value is `false`.
func (o JobConfigPtrOutput) SchedulerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobConfig) *bool {
		if v == nil {
			return nil
		}
		return v.SchedulerEnabled
	}).(pulumi.BoolPtrOutput)
}

// Specifies time when a scheduled task is triggered in a cycle. This parameter
// is valid only when `schedulerCycleType` is set to `hour`, `week`, or `month`.
//   - If `schedulerCycleType` is set to month, cycle is set to 1, and runAt is set to 15, the scheduled task is executed
//     on the 15th day of each month. You can set runAt to multiple values and separate the values with commas (,).
//     For example, if runAt is set to 1,2,3,4,5, the scheduled task is executed on the first day, second day, third day,
//     fourth day, and fifth day of each month.
//   - If `schedulerCycleType` is set to week and runAt is set to mon,tue,wed,thu,fri, the scheduled task is executed on
//     Monday to Friday.
//   - If `schedulerCycleType` is set to hour and runAt is set to 27,57, the scheduled task is executed at the 27th and
//     57th minute in the cycle.
func (o JobConfigPtrOutput) SchedulerRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.SchedulerRunAt
	}).(pulumi.StringPtrOutput)
}

// Specifies start time of a scheduled task.
// For example, `2018-01-24 19:56:19`
func (o JobConfigPtrOutput) SchedulerStartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.SchedulerStartDate
	}).(pulumi.StringPtrOutput)
}

// Specifies End time of a scheduled task. For example, `2018-01-27 23:59:00`.
// If you do not set the end time, the scheduled task is always executed and will never stop.
func (o JobConfigPtrOutput) SchedulerStopDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.SchedulerStopDate
	}).(pulumi.StringPtrOutput)
}

// Specifies name of the OBS bucket to which dirty data is
// written. This parameter is valid only when dirty data is written to `OBS`.
func (o JobConfigPtrOutput) ThrottlingDirtyWriteToBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.ThrottlingDirtyWriteToBucket
	}).(pulumi.StringPtrOutput)
}

// Specifies the directory in the OBS bucket or HDFS which
// dirty data is written to. For example, `/data/dirtydata/`.
func (o JobConfigPtrOutput) ThrottlingDirtyWriteToDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.ThrottlingDirtyWriteToDirectory
	}).(pulumi.StringPtrOutput)
}

// Specifies the link name to which dirty data is written to.
// The Dirty data can be written only to `OBS` or `HDFS`.
func (o JobConfigPtrOutput) ThrottlingDirtyWriteToLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobConfig) *string {
		if v == nil {
			return nil
		}
		return v.ThrottlingDirtyWriteToLink
	}).(pulumi.StringPtrOutput)
}

// Specifies maximum number of concurrent extraction jobs.
func (o JobConfigPtrOutput) ThrottlingExtractorsNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobConfig) *int {
		if v == nil {
			return nil
		}
		return v.ThrottlingExtractorsNumber
	}).(pulumi.IntPtrOutput)
}

// Specifies maximum number of loading jobs. This parameter is available
// only when HBase or Hive serves as the destination data source.
func (o JobConfigPtrOutput) ThrottlingLoaderNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobConfig) *int {
		if v == nil {
			return nil
		}
		return v.ThrottlingLoaderNumber
	}).(pulumi.IntPtrOutput)
}

// Specifies maximum number of error records in a single
// shard. When the number of error records of a map exceeds the upper limit, the task automatically ends.
func (o JobConfigPtrOutput) ThrottlingMaxErrorRecords() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobConfig) *int {
		if v == nil {
			return nil
		}
		return v.ThrottlingMaxErrorRecords
	}).(pulumi.IntPtrOutput)
}

// Specifies whether to write dirty data.
func (o JobConfigPtrOutput) ThrottlingRecordDirtyData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobConfig) *bool {
		if v == nil {
			return nil
		}
		return v.ThrottlingRecordDirtyData
	}).(pulumi.BoolPtrOutput)
}

type GetFlavorsFlavor struct {
	// The id of the CDM flavor.
	Id string `pulumi:"id"`
	// The name of the CDM flavor.
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
	// The id of the CDM flavor.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the CDM flavor.
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

// The id of the CDM flavor.
func (o GetFlavorsFlavorOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsFlavor) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the CDM flavor.
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
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInstanceInput)(nil)).Elem(), ClusterInstanceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInstanceArrayInput)(nil)).Elem(), ClusterInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobConfigInput)(nil)).Elem(), JobConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobConfigPtrInput)(nil)).Elem(), JobConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFlavorsFlavorInput)(nil)).Elem(), GetFlavorsFlavorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFlavorsFlavorArrayInput)(nil)).Elem(), GetFlavorsFlavorArray{})
	pulumi.RegisterOutputType(ClusterInstanceOutput{})
	pulumi.RegisterOutputType(ClusterInstanceArrayOutput{})
	pulumi.RegisterOutputType(JobConfigOutput{})
	pulumi.RegisterOutputType(JobConfigPtrOutput{})
	pulumi.RegisterOutputType(GetFlavorsFlavorOutput{})
	pulumi.RegisterOutputType(GetFlavorsFlavorArrayOutput{})
}
