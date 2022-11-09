// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package advancedantiddos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a forward rule resource of Advanced Anti-DDos service within HuaweiCloud.
//
// ## Import
//
// Rule can be imported using the `id` (combination of `instance_id`, `ip`, `forward_protocol` and `forward_port`), separated by slashes (/), e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:AdvancedAntiDDos/forwardRule:ForwardRule test &ltinstance_id&gt/&ltip&gt/&ltforward_protocol&gt/&ltforward_port&gt
//
// ```
type ForwardRule struct {
	pulumi.CustomResourceState

	// Specifies the forward port.
	// The valid value is range from **1** to **65535**.
	ForwardPort pulumi.IntOutput `pulumi:"forwardPort"`
	// Specifies the forward protocol.
	// The valid values are **tcp** and **udp**.
	ForwardProtocol pulumi.StringOutput `pulumi:"forwardProtocol"`
	// Specifies the ID of advanced Anti-DDoS instance.
	// Changing this will create a new rule resource.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the public IP addresss to which Advanced Anti-DDoS instance
	// belongs. Changing this will create a new rule resource.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The LVS forward policy.
	LbMethod pulumi.StringOutput `pulumi:"lbMethod"`
	// schema: Deprecated
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Specifies the source IP addresses, separated by commas (,).
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Specifies the source port.
	// The valid value is range from **1** to **65535**.
	SourcePort pulumi.IntOutput `pulumi:"sourcePort"`
	// The status of forward rule.
	Status pulumi.IntOutput `pulumi:"status"`
}

// NewForwardRule registers a new resource with the given unique name, arguments, and options.
func NewForwardRule(ctx *pulumi.Context,
	name string, args *ForwardRuleArgs, opts ...pulumi.ResourceOption) (*ForwardRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ForwardPort == nil {
		return nil, errors.New("invalid value for required argument 'ForwardPort'")
	}
	if args.ForwardProtocol == nil {
		return nil, errors.New("invalid value for required argument 'ForwardProtocol'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.SourceIp == nil {
		return nil, errors.New("invalid value for required argument 'SourceIp'")
	}
	if args.SourcePort == nil {
		return nil, errors.New("invalid value for required argument 'SourcePort'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ForwardRule
	err := ctx.RegisterResource("huaweicloud:AdvancedAntiDDos/forwardRule:ForwardRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForwardRule gets an existing ForwardRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardRuleState, opts ...pulumi.ResourceOption) (*ForwardRule, error) {
	var resource ForwardRule
	err := ctx.ReadResource("huaweicloud:AdvancedAntiDDos/forwardRule:ForwardRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ForwardRule resources.
type forwardRuleState struct {
	// Specifies the forward port.
	// The valid value is range from **1** to **65535**.
	ForwardPort *int `pulumi:"forwardPort"`
	// Specifies the forward protocol.
	// The valid values are **tcp** and **udp**.
	ForwardProtocol *string `pulumi:"forwardProtocol"`
	// Specifies the ID of advanced Anti-DDoS instance.
	// Changing this will create a new rule resource.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the public IP addresss to which Advanced Anti-DDoS instance
	// belongs. Changing this will create a new rule resource.
	Ip *string `pulumi:"ip"`
	// The LVS forward policy.
	LbMethod *string `pulumi:"lbMethod"`
	// schema: Deprecated
	RuleId *string `pulumi:"ruleId"`
	// Specifies the source IP addresses, separated by commas (,).
	SourceIp *string `pulumi:"sourceIp"`
	// Specifies the source port.
	// The valid value is range from **1** to **65535**.
	SourcePort *int `pulumi:"sourcePort"`
	// The status of forward rule.
	Status *int `pulumi:"status"`
}

type ForwardRuleState struct {
	// Specifies the forward port.
	// The valid value is range from **1** to **65535**.
	ForwardPort pulumi.IntPtrInput
	// Specifies the forward protocol.
	// The valid values are **tcp** and **udp**.
	ForwardProtocol pulumi.StringPtrInput
	// Specifies the ID of advanced Anti-DDoS instance.
	// Changing this will create a new rule resource.
	InstanceId pulumi.StringPtrInput
	// Specifies the public IP addresss to which Advanced Anti-DDoS instance
	// belongs. Changing this will create a new rule resource.
	Ip pulumi.StringPtrInput
	// The LVS forward policy.
	LbMethod pulumi.StringPtrInput
	// schema: Deprecated
	RuleId pulumi.StringPtrInput
	// Specifies the source IP addresses, separated by commas (,).
	SourceIp pulumi.StringPtrInput
	// Specifies the source port.
	// The valid value is range from **1** to **65535**.
	SourcePort pulumi.IntPtrInput
	// The status of forward rule.
	Status pulumi.IntPtrInput
}

func (ForwardRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardRuleState)(nil)).Elem()
}

type forwardRuleArgs struct {
	// Specifies the forward port.
	// The valid value is range from **1** to **65535**.
	ForwardPort int `pulumi:"forwardPort"`
	// Specifies the forward protocol.
	// The valid values are **tcp** and **udp**.
	ForwardProtocol string `pulumi:"forwardProtocol"`
	// Specifies the ID of advanced Anti-DDoS instance.
	// Changing this will create a new rule resource.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the public IP addresss to which Advanced Anti-DDoS instance
	// belongs. Changing this will create a new rule resource.
	Ip string `pulumi:"ip"`
	// Specifies the source IP addresses, separated by commas (,).
	SourceIp string `pulumi:"sourceIp"`
	// Specifies the source port.
	// The valid value is range from **1** to **65535**.
	SourcePort int `pulumi:"sourcePort"`
}

// The set of arguments for constructing a ForwardRule resource.
type ForwardRuleArgs struct {
	// Specifies the forward port.
	// The valid value is range from **1** to **65535**.
	ForwardPort pulumi.IntInput
	// Specifies the forward protocol.
	// The valid values are **tcp** and **udp**.
	ForwardProtocol pulumi.StringInput
	// Specifies the ID of advanced Anti-DDoS instance.
	// Changing this will create a new rule resource.
	InstanceId pulumi.StringInput
	// Specifies the public IP addresss to which Advanced Anti-DDoS instance
	// belongs. Changing this will create a new rule resource.
	Ip pulumi.StringInput
	// Specifies the source IP addresses, separated by commas (,).
	SourceIp pulumi.StringInput
	// Specifies the source port.
	// The valid value is range from **1** to **65535**.
	SourcePort pulumi.IntInput
}

func (ForwardRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardRuleArgs)(nil)).Elem()
}

type ForwardRuleInput interface {
	pulumi.Input

	ToForwardRuleOutput() ForwardRuleOutput
	ToForwardRuleOutputWithContext(ctx context.Context) ForwardRuleOutput
}

func (*ForwardRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardRule)(nil)).Elem()
}

func (i *ForwardRule) ToForwardRuleOutput() ForwardRuleOutput {
	return i.ToForwardRuleOutputWithContext(context.Background())
}

func (i *ForwardRule) ToForwardRuleOutputWithContext(ctx context.Context) ForwardRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardRuleOutput)
}

// ForwardRuleArrayInput is an input type that accepts ForwardRuleArray and ForwardRuleArrayOutput values.
// You can construct a concrete instance of `ForwardRuleArrayInput` via:
//
//	ForwardRuleArray{ ForwardRuleArgs{...} }
type ForwardRuleArrayInput interface {
	pulumi.Input

	ToForwardRuleArrayOutput() ForwardRuleArrayOutput
	ToForwardRuleArrayOutputWithContext(context.Context) ForwardRuleArrayOutput
}

type ForwardRuleArray []ForwardRuleInput

func (ForwardRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ForwardRule)(nil)).Elem()
}

func (i ForwardRuleArray) ToForwardRuleArrayOutput() ForwardRuleArrayOutput {
	return i.ToForwardRuleArrayOutputWithContext(context.Background())
}

func (i ForwardRuleArray) ToForwardRuleArrayOutputWithContext(ctx context.Context) ForwardRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardRuleArrayOutput)
}

// ForwardRuleMapInput is an input type that accepts ForwardRuleMap and ForwardRuleMapOutput values.
// You can construct a concrete instance of `ForwardRuleMapInput` via:
//
//	ForwardRuleMap{ "key": ForwardRuleArgs{...} }
type ForwardRuleMapInput interface {
	pulumi.Input

	ToForwardRuleMapOutput() ForwardRuleMapOutput
	ToForwardRuleMapOutputWithContext(context.Context) ForwardRuleMapOutput
}

type ForwardRuleMap map[string]ForwardRuleInput

func (ForwardRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ForwardRule)(nil)).Elem()
}

func (i ForwardRuleMap) ToForwardRuleMapOutput() ForwardRuleMapOutput {
	return i.ToForwardRuleMapOutputWithContext(context.Background())
}

func (i ForwardRuleMap) ToForwardRuleMapOutputWithContext(ctx context.Context) ForwardRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardRuleMapOutput)
}

type ForwardRuleOutput struct{ *pulumi.OutputState }

func (ForwardRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardRule)(nil)).Elem()
}

func (o ForwardRuleOutput) ToForwardRuleOutput() ForwardRuleOutput {
	return o
}

func (o ForwardRuleOutput) ToForwardRuleOutputWithContext(ctx context.Context) ForwardRuleOutput {
	return o
}

// Specifies the forward port.
// The valid value is range from **1** to **65535**.
func (o ForwardRuleOutput) ForwardPort() pulumi.IntOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.IntOutput { return v.ForwardPort }).(pulumi.IntOutput)
}

// Specifies the forward protocol.
// The valid values are **tcp** and **udp**.
func (o ForwardRuleOutput) ForwardProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.ForwardProtocol }).(pulumi.StringOutput)
}

// Specifies the ID of advanced Anti-DDoS instance.
// Changing this will create a new rule resource.
func (o ForwardRuleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the public IP addresss to which Advanced Anti-DDoS instance
// belongs. Changing this will create a new rule resource.
func (o ForwardRuleOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The LVS forward policy.
func (o ForwardRuleOutput) LbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.LbMethod }).(pulumi.StringOutput)
}

// schema: Deprecated
func (o ForwardRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Specifies the source IP addresses, separated by commas (,).
func (o ForwardRuleOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Specifies the source port.
// The valid value is range from **1** to **65535**.
func (o ForwardRuleOutput) SourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.IntOutput { return v.SourcePort }).(pulumi.IntOutput)
}

// The status of forward rule.
func (o ForwardRuleOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

type ForwardRuleArrayOutput struct{ *pulumi.OutputState }

func (ForwardRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ForwardRule)(nil)).Elem()
}

func (o ForwardRuleArrayOutput) ToForwardRuleArrayOutput() ForwardRuleArrayOutput {
	return o
}

func (o ForwardRuleArrayOutput) ToForwardRuleArrayOutputWithContext(ctx context.Context) ForwardRuleArrayOutput {
	return o
}

func (o ForwardRuleArrayOutput) Index(i pulumi.IntInput) ForwardRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ForwardRule {
		return vs[0].([]*ForwardRule)[vs[1].(int)]
	}).(ForwardRuleOutput)
}

type ForwardRuleMapOutput struct{ *pulumi.OutputState }

func (ForwardRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ForwardRule)(nil)).Elem()
}

func (o ForwardRuleMapOutput) ToForwardRuleMapOutput() ForwardRuleMapOutput {
	return o
}

func (o ForwardRuleMapOutput) ToForwardRuleMapOutputWithContext(ctx context.Context) ForwardRuleMapOutput {
	return o
}

func (o ForwardRuleMapOutput) MapIndex(k pulumi.StringInput) ForwardRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ForwardRule {
		return vs[0].(map[string]*ForwardRule)[vs[1].(string)]
	}).(ForwardRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardRuleInput)(nil)).Elem(), &ForwardRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardRuleArrayInput)(nil)).Elem(), ForwardRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardRuleMapInput)(nil)).Elem(), ForwardRuleMap{})
	pulumi.RegisterOutputType(ForwardRuleOutput{})
	pulumi.RegisterOutputType(ForwardRuleArrayOutput{})
	pulumi.RegisterOutputType(ForwardRuleMapOutput{})
}
