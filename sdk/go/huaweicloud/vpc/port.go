// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Port struct {
	pulumi.CustomResourceState

	AdminStateUp        pulumi.BoolOutput                 `pulumi:"adminStateUp"`
	AllFixedIps         pulumi.StringArrayOutput          `pulumi:"allFixedIps"`
	AllSecurityGroupIds pulumi.StringArrayOutput          `pulumi:"allSecurityGroupIds"`
	AllowedAddressPairs PortAllowedAddressPairArrayOutput `pulumi:"allowedAddressPairs"`
	DeviceId            pulumi.StringOutput               `pulumi:"deviceId"`
	DeviceOwner         pulumi.StringOutput               `pulumi:"deviceOwner"`
	ExtraDhcpOptions    PortExtraDhcpOptionArrayOutput    `pulumi:"extraDhcpOptions"`
	FixedIps            PortFixedIpArrayOutput            `pulumi:"fixedIps"`
	MacAddress          pulumi.StringOutput               `pulumi:"macAddress"`
	Name                pulumi.StringOutput               `pulumi:"name"`
	NetworkId           pulumi.StringOutput               `pulumi:"networkId"`
	NoSecurityGroups    pulumi.BoolPtrOutput              `pulumi:"noSecurityGroups"`
	Region              pulumi.StringOutput               `pulumi:"region"`
	SecurityGroupIds    pulumi.StringArrayOutput          `pulumi:"securityGroupIds"`
	// Deprecated: tenant_id is deprecated
	TenantId   pulumi.StringOutput    `pulumi:"tenantId"`
	ValueSpecs pulumi.StringMapOutput `pulumi:"valueSpecs"`
}

// NewPort registers a new resource with the given unique name, arguments, and options.
func NewPort(ctx *pulumi.Context,
	name string, args *PortArgs, opts ...pulumi.ResourceOption) (*Port, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Port
	err := ctx.RegisterResource("huaweicloud:Vpc/port:Port", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPort gets an existing Port resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortState, opts ...pulumi.ResourceOption) (*Port, error) {
	var resource Port
	err := ctx.ReadResource("huaweicloud:Vpc/port:Port", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Port resources.
type portState struct {
	AdminStateUp        *bool                    `pulumi:"adminStateUp"`
	AllFixedIps         []string                 `pulumi:"allFixedIps"`
	AllSecurityGroupIds []string                 `pulumi:"allSecurityGroupIds"`
	AllowedAddressPairs []PortAllowedAddressPair `pulumi:"allowedAddressPairs"`
	DeviceId            *string                  `pulumi:"deviceId"`
	DeviceOwner         *string                  `pulumi:"deviceOwner"`
	ExtraDhcpOptions    []PortExtraDhcpOption    `pulumi:"extraDhcpOptions"`
	FixedIps            []PortFixedIp            `pulumi:"fixedIps"`
	MacAddress          *string                  `pulumi:"macAddress"`
	Name                *string                  `pulumi:"name"`
	NetworkId           *string                  `pulumi:"networkId"`
	NoSecurityGroups    *bool                    `pulumi:"noSecurityGroups"`
	Region              *string                  `pulumi:"region"`
	SecurityGroupIds    []string                 `pulumi:"securityGroupIds"`
	// Deprecated: tenant_id is deprecated
	TenantId   *string           `pulumi:"tenantId"`
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
}

type PortState struct {
	AdminStateUp        pulumi.BoolPtrInput
	AllFixedIps         pulumi.StringArrayInput
	AllSecurityGroupIds pulumi.StringArrayInput
	AllowedAddressPairs PortAllowedAddressPairArrayInput
	DeviceId            pulumi.StringPtrInput
	DeviceOwner         pulumi.StringPtrInput
	ExtraDhcpOptions    PortExtraDhcpOptionArrayInput
	FixedIps            PortFixedIpArrayInput
	MacAddress          pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	NetworkId           pulumi.StringPtrInput
	NoSecurityGroups    pulumi.BoolPtrInput
	Region              pulumi.StringPtrInput
	SecurityGroupIds    pulumi.StringArrayInput
	// Deprecated: tenant_id is deprecated
	TenantId   pulumi.StringPtrInput
	ValueSpecs pulumi.StringMapInput
}

func (PortState) ElementType() reflect.Type {
	return reflect.TypeOf((*portState)(nil)).Elem()
}

type portArgs struct {
	AdminStateUp        *bool                    `pulumi:"adminStateUp"`
	AllowedAddressPairs []PortAllowedAddressPair `pulumi:"allowedAddressPairs"`
	DeviceId            *string                  `pulumi:"deviceId"`
	DeviceOwner         *string                  `pulumi:"deviceOwner"`
	ExtraDhcpOptions    []PortExtraDhcpOption    `pulumi:"extraDhcpOptions"`
	FixedIps            []PortFixedIp            `pulumi:"fixedIps"`
	MacAddress          *string                  `pulumi:"macAddress"`
	Name                *string                  `pulumi:"name"`
	NetworkId           string                   `pulumi:"networkId"`
	NoSecurityGroups    *bool                    `pulumi:"noSecurityGroups"`
	Region              *string                  `pulumi:"region"`
	SecurityGroupIds    []string                 `pulumi:"securityGroupIds"`
	// Deprecated: tenant_id is deprecated
	TenantId   *string           `pulumi:"tenantId"`
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Port resource.
type PortArgs struct {
	AdminStateUp        pulumi.BoolPtrInput
	AllowedAddressPairs PortAllowedAddressPairArrayInput
	DeviceId            pulumi.StringPtrInput
	DeviceOwner         pulumi.StringPtrInput
	ExtraDhcpOptions    PortExtraDhcpOptionArrayInput
	FixedIps            PortFixedIpArrayInput
	MacAddress          pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	NetworkId           pulumi.StringInput
	NoSecurityGroups    pulumi.BoolPtrInput
	Region              pulumi.StringPtrInput
	SecurityGroupIds    pulumi.StringArrayInput
	// Deprecated: tenant_id is deprecated
	TenantId   pulumi.StringPtrInput
	ValueSpecs pulumi.StringMapInput
}

func (PortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portArgs)(nil)).Elem()
}

type PortInput interface {
	pulumi.Input

	ToPortOutput() PortOutput
	ToPortOutputWithContext(ctx context.Context) PortOutput
}

func (*Port) ElementType() reflect.Type {
	return reflect.TypeOf((**Port)(nil)).Elem()
}

func (i *Port) ToPortOutput() PortOutput {
	return i.ToPortOutputWithContext(context.Background())
}

func (i *Port) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortOutput)
}

// PortArrayInput is an input type that accepts PortArray and PortArrayOutput values.
// You can construct a concrete instance of `PortArrayInput` via:
//
//	PortArray{ PortArgs{...} }
type PortArrayInput interface {
	pulumi.Input

	ToPortArrayOutput() PortArrayOutput
	ToPortArrayOutputWithContext(context.Context) PortArrayOutput
}

type PortArray []PortInput

func (PortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Port)(nil)).Elem()
}

func (i PortArray) ToPortArrayOutput() PortArrayOutput {
	return i.ToPortArrayOutputWithContext(context.Background())
}

func (i PortArray) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortArrayOutput)
}

// PortMapInput is an input type that accepts PortMap and PortMapOutput values.
// You can construct a concrete instance of `PortMapInput` via:
//
//	PortMap{ "key": PortArgs{...} }
type PortMapInput interface {
	pulumi.Input

	ToPortMapOutput() PortMapOutput
	ToPortMapOutputWithContext(context.Context) PortMapOutput
}

type PortMap map[string]PortInput

func (PortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Port)(nil)).Elem()
}

func (i PortMap) ToPortMapOutput() PortMapOutput {
	return i.ToPortMapOutputWithContext(context.Background())
}

func (i PortMap) ToPortMapOutputWithContext(ctx context.Context) PortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortMapOutput)
}

type PortOutput struct{ *pulumi.OutputState }

func (PortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Port)(nil)).Elem()
}

func (o PortOutput) ToPortOutput() PortOutput {
	return o
}

func (o PortOutput) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return o
}

func (o PortOutput) AdminStateUp() pulumi.BoolOutput {
	return o.ApplyT(func(v *Port) pulumi.BoolOutput { return v.AdminStateUp }).(pulumi.BoolOutput)
}

func (o PortOutput) AllFixedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.AllFixedIps }).(pulumi.StringArrayOutput)
}

func (o PortOutput) AllSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.AllSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o PortOutput) AllowedAddressPairs() PortAllowedAddressPairArrayOutput {
	return o.ApplyT(func(v *Port) PortAllowedAddressPairArrayOutput { return v.AllowedAddressPairs }).(PortAllowedAddressPairArrayOutput)
}

func (o PortOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

func (o PortOutput) DeviceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.DeviceOwner }).(pulumi.StringOutput)
}

func (o PortOutput) ExtraDhcpOptions() PortExtraDhcpOptionArrayOutput {
	return o.ApplyT(func(v *Port) PortExtraDhcpOptionArrayOutput { return v.ExtraDhcpOptions }).(PortExtraDhcpOptionArrayOutput)
}

func (o PortOutput) FixedIps() PortFixedIpArrayOutput {
	return o.ApplyT(func(v *Port) PortFixedIpArrayOutput { return v.FixedIps }).(PortFixedIpArrayOutput)
}

func (o PortOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

func (o PortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PortOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

func (o PortOutput) NoSecurityGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Port) pulumi.BoolPtrOutput { return v.NoSecurityGroups }).(pulumi.BoolPtrOutput)
}

func (o PortOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o PortOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Deprecated: tenant_id is deprecated
func (o PortOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o PortOutput) ValueSpecs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Port) pulumi.StringMapOutput { return v.ValueSpecs }).(pulumi.StringMapOutput)
}

type PortArrayOutput struct{ *pulumi.OutputState }

func (PortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Port)(nil)).Elem()
}

func (o PortArrayOutput) ToPortArrayOutput() PortArrayOutput {
	return o
}

func (o PortArrayOutput) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return o
}

func (o PortArrayOutput) Index(i pulumi.IntInput) PortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Port {
		return vs[0].([]*Port)[vs[1].(int)]
	}).(PortOutput)
}

type PortMapOutput struct{ *pulumi.OutputState }

func (PortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Port)(nil)).Elem()
}

func (o PortMapOutput) ToPortMapOutput() PortMapOutput {
	return o
}

func (o PortMapOutput) ToPortMapOutputWithContext(ctx context.Context) PortMapOutput {
	return o
}

func (o PortMapOutput) MapIndex(k pulumi.StringInput) PortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Port {
		return vs[0].(map[string]*Port)[vs[1].(string)]
	}).(PortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortInput)(nil)).Elem(), &Port{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortArrayInput)(nil)).Elem(), PortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortMapInput)(nil)).Elem(), PortMap{})
	pulumi.RegisterOutputType(PortOutput{})
	pulumi.RegisterOutputType(PortArrayOutput{})
	pulumi.RegisterOutputType(PortMapOutput{})
}
