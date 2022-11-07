// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of a specific IEC security group.
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
//			sgName := cfg.RequireObject("sgName")
//			_, err := Iec.GetSecurityGroup(ctx, &iec.GetSecurityGroupArgs{
//				Name: sgName,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSecurityGroup(ctx *pulumi.Context, args *LookupSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupResult, error) {
	var rv LookupSecurityGroupResult
	err := ctx.Invoke("huaweicloud:Iec/getSecurityGroup:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroup.
type LookupSecurityGroupArgs struct {
	// Specifies the name of the security group with a maximum of 64 characters.
	Name string `pulumi:"name"`
}

// A collection of values returned by getSecurityGroup.
type LookupSecurityGroupResult struct {
	// The description for the IEC security group rules.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// An Array of one or more security group rules. The object is documented below.
	SecurityGroupRules []GetSecurityGroupSecurityGroupRule `pulumi:"securityGroupRules"`
}

func LookupSecurityGroupOutput(ctx *pulumi.Context, args LookupSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupResult, error) {
			args := v.(LookupSecurityGroupArgs)
			r, err := LookupSecurityGroup(ctx, &args, opts...)
			var s LookupSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityGroupResultOutput)
}

// A collection of arguments for invoking getSecurityGroup.
type LookupSecurityGroupOutputArgs struct {
	// Specifies the name of the security group with a maximum of 64 characters.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityGroup.
type LookupSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupResult)(nil)).Elem()
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutput() LookupSecurityGroupResultOutput {
	return o
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutputWithContext(ctx context.Context) LookupSecurityGroupResultOutput {
	return o
}

// The description for the IEC security group rules.
func (o LookupSecurityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// An Array of one or more security group rules. The object is documented below.
func (o LookupSecurityGroupResultOutput) SecurityGroupRules() GetSecurityGroupSecurityGroupRuleArrayOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) []GetSecurityGroupSecurityGroupRule { return v.SecurityGroupRules }).(GetSecurityGroupSecurityGroupRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupResultOutput{})
}