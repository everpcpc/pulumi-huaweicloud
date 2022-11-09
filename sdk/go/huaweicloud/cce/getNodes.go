// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cce

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of CCE nodes.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterId := cfg.RequireObject("clusterId")
//			nodeName := cfg.RequireObject("nodeName")
//			_, err := Cce.GetNodes(ctx, &cce.GetNodesArgs{
//				ClusterId: clusterId,
//				Name:      pulumi.StringRef(nodeName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNodes(ctx *pulumi.Context, args *GetNodesArgs, opts ...pulumi.InvokeOption) (*GetNodesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetNodesResult
	err := ctx.Invoke("huaweicloud:Cce/getNodes:getNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodes.
type GetNodesArgs struct {
	// Specifies the ID of CCE cluster.
	ClusterId string `pulumi:"clusterId"`
	// Specifies the of the node.
	Name *string `pulumi:"name"`
	// Specifies the ID of the node.
	NodeId *string `pulumi:"nodeId"`
	// Specifies the region in which to obtain the CCE nodes. If omitted, the provider-level
	// region will be used.
	Region *string `pulumi:"region"`
	// Specifies the status of the node.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getNodes.
type GetNodesResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates a list of IDs of all CCE nodes found.
	Ids []string `pulumi:"ids"`
	// The name of the node.
	Name   *string `pulumi:"name"`
	NodeId *string `pulumi:"nodeId"`
	// Indicates a list of CCE nodes found. Structure is documented below.
	Nodes  []GetNodesNode `pulumi:"nodes"`
	Region string         `pulumi:"region"`
	// The state of the node.
	Status *string `pulumi:"status"`
}

func GetNodesOutput(ctx *pulumi.Context, args GetNodesOutputArgs, opts ...pulumi.InvokeOption) GetNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNodesResult, error) {
			args := v.(GetNodesArgs)
			r, err := GetNodes(ctx, &args, opts...)
			var s GetNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNodesResultOutput)
}

// A collection of arguments for invoking getNodes.
type GetNodesOutputArgs struct {
	// Specifies the ID of CCE cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Specifies the of the node.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the ID of the node.
	NodeId pulumi.StringPtrInput `pulumi:"nodeId"`
	// Specifies the region in which to obtain the CCE nodes. If omitted, the provider-level
	// region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the status of the node.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesArgs)(nil)).Elem()
}

// A collection of values returned by getNodes.
type GetNodesResultOutput struct{ *pulumi.OutputState }

func (GetNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesResult)(nil)).Elem()
}

func (o GetNodesResultOutput) ToGetNodesResultOutput() GetNodesResultOutput {
	return o
}

func (o GetNodesResultOutput) ToGetNodesResultOutputWithContext(ctx context.Context) GetNodesResultOutput {
	return o
}

func (o GetNodesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates a list of IDs of all CCE nodes found.
func (o GetNodesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The name of the node.
func (o GetNodesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetNodesResultOutput) NodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodesResult) *string { return v.NodeId }).(pulumi.StringPtrOutput)
}

// Indicates a list of CCE nodes found. Structure is documented below.
func (o GetNodesResultOutput) Nodes() GetNodesNodeArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []GetNodesNode { return v.Nodes }).(GetNodesNodeArrayOutput)
}

func (o GetNodesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesResult) string { return v.Region }).(pulumi.StringOutput)
}

// The state of the node.
func (o GetNodesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNodesResultOutput{})
}
