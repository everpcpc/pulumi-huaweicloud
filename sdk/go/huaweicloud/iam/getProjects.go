// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query the project list within HuaweiCloud.
//
// > You *must* have IAM read privileges to use this data source.
//
// ## Example Usage
// ### Obtain project information by name
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Iam"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Iam.GetProjects(ctx, &iam.GetProjectsArgs{
//				Name: pulumi.StringRef("cn-north-4_demo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Obtain special project information by name
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Iam"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Iam.GetProjects(ctx, &iam.GetProjectsArgs{
//				Name: pulumi.StringRef("MOS"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("huaweicloud:Iam/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// Specifies the project name to query.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The project name.
	Name *string `pulumi:"name"`
	// The details of the query projects. The structure is documented below.
	Projects []GetProjectsProject `pulumi:"projects"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectsResult, error) {
			args := v.(GetProjectsArgs)
			r, err := GetProjects(ctx, &args, opts...)
			var s GetProjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	// Specifies the project name to query.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The project name.
func (o GetProjectsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The details of the query projects. The structure is documented below.
func (o GetProjectsResultOutput) Projects() GetProjectsProjectArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsProject { return v.Projects }).(GetProjectsProjectArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}
