// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspace

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to register or unregister the Workspace service in HuaweiCloud.
//
// ## Example Usage
// ### Register the Workspace service and use local authentication
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Workspace"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vpcId := cfg.RequireObject("vpcId")
//			networkIds := cfg.Require("networkIds")
//			_, err := Workspace.NewService(ctx, "test", &Workspace.ServiceArgs{
//				AccessMode: pulumi.String("INTERNET"),
//				VpcId:      pulumi.Any(vpcId),
//				NetworkIds: networkIds,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Appendix
//
// <a name="secgroupRulesForAdDomainConnection"></a>
// If a firewall is deployed between Windows AD and the Workspace service, you need to open the following ports on the
// firewall for the desktops of Workspace service to connect to Windows AD or DNS:
//
// | Protocol | Ports | Usage |
// | ---- | ---- | ---- |
// | TCP | 135 | RPC protocol (required for LDAP, Distributed File System, and Distributed File Replication) |
// | UDP | 137 | NetBIOS name resolution (required by the network login service) |
// | UDP | 138 | NetBIOS datagram service (distributed file system, network login and other services need to use this port) |
// | TCP | 139 | NetBIOS-SSN Service (Network Basic I/O Interface) |
// | TCP | 445 | NetBIOS-SSN Service (Network Basic I/O Interface) |
// | UDP | 445 | NetBIOS-SSN Service (Network Basic I/O Interface) |
// | TCP | 49152-65535 | RPC dynamic ports (ports that are not hardened and open by AD. If AD is hardened, ports 50152-51151 need to be opened) |
// | UDP | 49152-65535 | RPC dynamic ports (ports that are not hardened and open by AD. If AD is hardened, ports 50152-51151 need to be opened) |
// | TCP | 88 | Kerberos Key Distribution Center Service |
// | UDP | 88 | Kerberos Key Distribution Center Service |
// | UDP | 123 | Port used by NTP service |
// | TCP | 389 | LDAP server |
// | UDP | 389 | LDAP server |
// | TCP | 464 | Kerberos authentication protocol |
// | UDP | 464 | Kerberos Authentication Protocol |
// | UDP | 500 | isakmp |
// | TCP | 593 | RPC over HTTP |
// | TCP | 636 | LDAP SSL |
// | TCP | 53 | DNS server |
// | UDP | 53 | DNS server |
//
// ## Import
//
// Service can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Workspace/service:Service test fd3f81cb-d95f-43ce-b342-81b6b5dcadda
//
// ```
type Service struct {
	pulumi.CustomResourceState

	// Specifies the access mode of Workspace service.
	// The valid values are as follows:
	// + **INTERNET**: internet access.
	// + **DEDICATED**: dedicated line access.
	// + **BOTH**: both internet access and dedicated access are supported.
	AccessMode pulumi.StringOutput `pulumi:"accessMode"`
	// Specifies the configuration of AD domain.
	// Required if `authType` is **LOCAL_AD**. Make sure that the selected VPC network and the network to which AD
	// belongs can be connected. The object structure is documented below.
	AdDomain ServiceAdDomainOutput `pulumi:"adDomain"`
	// Specifies the authentication type of Workspace service.
	// The valid values are as follows:
	// + **LITE_AS**: Local authentication.
	// + **LOCAL_AD**: Connect to AD domain.
	AuthType pulumi.StringPtrOutput `pulumi:"authType"`
	// The subnet segments of the dedicated access.
	DedicatedSubnets pulumi.StringArrayOutput `pulumi:"dedicatedSubnets"`
	// The desktop security group automatically created under the specified VPC after the service
	// is registered. The object structure is documented below.
	DesktopSecurityGroups ServiceDesktopSecurityGroupArrayOutput `pulumi:"desktopSecurityGroups"`
	// Specifies the enterprise ID.
	// The enterprise ID is the unique identification in the Workspace service.
	// If omitted, the system will automatically generate an enterprise ID.
	// The ID can contain `1` to `32` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	EnterpriseId pulumi.StringOutput `pulumi:"enterpriseId"`
	// The management component security group automatically created under the specified
	// VPC after the service is registered. The object structure is documented below.
	InfrastructureSecurityGroups ServiceInfrastructureSecurityGroupArrayOutput `pulumi:"infrastructureSecurityGroups"`
	// The internet access address.
	// This attribute is returned only when the accessMode is **INTERNET** or **BOTH**.
	InternetAccessAddress pulumi.StringOutput `pulumi:"internetAccessAddress"`
	// Specifies the internet access port.
	// The valid value is range from `1,025` to `65,535`.
	InternetAccessPort pulumi.IntOutput `pulumi:"internetAccessPort"`
	// The subnet segment of the management component.
	ManagementSubnetCidr pulumi.StringOutput `pulumi:"managementSubnetCidr"`
	// The network ID list of subnets that the service have.
	// The subnets corresponding to this parameter must be included in the VPC resource corresponding to `vpcId`.
	// These subnet segments cannot conflict with `172.16.0.0/12`.
	NetworkIds pulumi.StringArrayOutput `pulumi:"networkIds"`
	// The region in which to register the Workspace service.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The current status of the Workspace service.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the VPC ID to which the service belongs.
	// Changing this will create a new resource.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessMode == nil {
		return nil, errors.New("invalid value for required argument 'AccessMode'")
	}
	if args.NetworkIds == nil {
		return nil, errors.New("invalid value for required argument 'NetworkIds'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("huaweicloud:Workspace/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("huaweicloud:Workspace/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Specifies the access mode of Workspace service.
	// The valid values are as follows:
	// + **INTERNET**: internet access.
	// + **DEDICATED**: dedicated line access.
	// + **BOTH**: both internet access and dedicated access are supported.
	AccessMode *string `pulumi:"accessMode"`
	// Specifies the configuration of AD domain.
	// Required if `authType` is **LOCAL_AD**. Make sure that the selected VPC network and the network to which AD
	// belongs can be connected. The object structure is documented below.
	AdDomain *ServiceAdDomain `pulumi:"adDomain"`
	// Specifies the authentication type of Workspace service.
	// The valid values are as follows:
	// + **LITE_AS**: Local authentication.
	// + **LOCAL_AD**: Connect to AD domain.
	AuthType *string `pulumi:"authType"`
	// The subnet segments of the dedicated access.
	DedicatedSubnets []string `pulumi:"dedicatedSubnets"`
	// The desktop security group automatically created under the specified VPC after the service
	// is registered. The object structure is documented below.
	DesktopSecurityGroups []ServiceDesktopSecurityGroup `pulumi:"desktopSecurityGroups"`
	// Specifies the enterprise ID.
	// The enterprise ID is the unique identification in the Workspace service.
	// If omitted, the system will automatically generate an enterprise ID.
	// The ID can contain `1` to `32` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	EnterpriseId *string `pulumi:"enterpriseId"`
	// The management component security group automatically created under the specified
	// VPC after the service is registered. The object structure is documented below.
	InfrastructureSecurityGroups []ServiceInfrastructureSecurityGroup `pulumi:"infrastructureSecurityGroups"`
	// The internet access address.
	// This attribute is returned only when the accessMode is **INTERNET** or **BOTH**.
	InternetAccessAddress *string `pulumi:"internetAccessAddress"`
	// Specifies the internet access port.
	// The valid value is range from `1,025` to `65,535`.
	InternetAccessPort *int `pulumi:"internetAccessPort"`
	// The subnet segment of the management component.
	ManagementSubnetCidr *string `pulumi:"managementSubnetCidr"`
	// The network ID list of subnets that the service have.
	// The subnets corresponding to this parameter must be included in the VPC resource corresponding to `vpcId`.
	// These subnet segments cannot conflict with `172.16.0.0/12`.
	NetworkIds []string `pulumi:"networkIds"`
	// The region in which to register the Workspace service.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// The current status of the Workspace service.
	Status *string `pulumi:"status"`
	// Specifies the VPC ID to which the service belongs.
	// Changing this will create a new resource.
	VpcId *string `pulumi:"vpcId"`
}

type ServiceState struct {
	// Specifies the access mode of Workspace service.
	// The valid values are as follows:
	// + **INTERNET**: internet access.
	// + **DEDICATED**: dedicated line access.
	// + **BOTH**: both internet access and dedicated access are supported.
	AccessMode pulumi.StringPtrInput
	// Specifies the configuration of AD domain.
	// Required if `authType` is **LOCAL_AD**. Make sure that the selected VPC network and the network to which AD
	// belongs can be connected. The object structure is documented below.
	AdDomain ServiceAdDomainPtrInput
	// Specifies the authentication type of Workspace service.
	// The valid values are as follows:
	// + **LITE_AS**: Local authentication.
	// + **LOCAL_AD**: Connect to AD domain.
	AuthType pulumi.StringPtrInput
	// The subnet segments of the dedicated access.
	DedicatedSubnets pulumi.StringArrayInput
	// The desktop security group automatically created under the specified VPC after the service
	// is registered. The object structure is documented below.
	DesktopSecurityGroups ServiceDesktopSecurityGroupArrayInput
	// Specifies the enterprise ID.
	// The enterprise ID is the unique identification in the Workspace service.
	// If omitted, the system will automatically generate an enterprise ID.
	// The ID can contain `1` to `32` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	EnterpriseId pulumi.StringPtrInput
	// The management component security group automatically created under the specified
	// VPC after the service is registered. The object structure is documented below.
	InfrastructureSecurityGroups ServiceInfrastructureSecurityGroupArrayInput
	// The internet access address.
	// This attribute is returned only when the accessMode is **INTERNET** or **BOTH**.
	InternetAccessAddress pulumi.StringPtrInput
	// Specifies the internet access port.
	// The valid value is range from `1,025` to `65,535`.
	InternetAccessPort pulumi.IntPtrInput
	// The subnet segment of the management component.
	ManagementSubnetCidr pulumi.StringPtrInput
	// The network ID list of subnets that the service have.
	// The subnets corresponding to this parameter must be included in the VPC resource corresponding to `vpcId`.
	// These subnet segments cannot conflict with `172.16.0.0/12`.
	NetworkIds pulumi.StringArrayInput
	// The region in which to register the Workspace service.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// The current status of the Workspace service.
	Status pulumi.StringPtrInput
	// Specifies the VPC ID to which the service belongs.
	// Changing this will create a new resource.
	VpcId pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Specifies the access mode of Workspace service.
	// The valid values are as follows:
	// + **INTERNET**: internet access.
	// + **DEDICATED**: dedicated line access.
	// + **BOTH**: both internet access and dedicated access are supported.
	AccessMode string `pulumi:"accessMode"`
	// Specifies the configuration of AD domain.
	// Required if `authType` is **LOCAL_AD**. Make sure that the selected VPC network and the network to which AD
	// belongs can be connected. The object structure is documented below.
	AdDomain *ServiceAdDomain `pulumi:"adDomain"`
	// Specifies the authentication type of Workspace service.
	// The valid values are as follows:
	// + **LITE_AS**: Local authentication.
	// + **LOCAL_AD**: Connect to AD domain.
	AuthType *string `pulumi:"authType"`
	// The subnet segments of the dedicated access.
	DedicatedSubnets []string `pulumi:"dedicatedSubnets"`
	// Specifies the enterprise ID.
	// The enterprise ID is the unique identification in the Workspace service.
	// If omitted, the system will automatically generate an enterprise ID.
	// The ID can contain `1` to `32` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	EnterpriseId *string `pulumi:"enterpriseId"`
	// Specifies the internet access port.
	// The valid value is range from `1,025` to `65,535`.
	InternetAccessPort *int `pulumi:"internetAccessPort"`
	// The subnet segment of the management component.
	ManagementSubnetCidr *string `pulumi:"managementSubnetCidr"`
	// The network ID list of subnets that the service have.
	// The subnets corresponding to this parameter must be included in the VPC resource corresponding to `vpcId`.
	// These subnet segments cannot conflict with `172.16.0.0/12`.
	NetworkIds []string `pulumi:"networkIds"`
	// The region in which to register the Workspace service.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the VPC ID to which the service belongs.
	// Changing this will create a new resource.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Specifies the access mode of Workspace service.
	// The valid values are as follows:
	// + **INTERNET**: internet access.
	// + **DEDICATED**: dedicated line access.
	// + **BOTH**: both internet access and dedicated access are supported.
	AccessMode pulumi.StringInput
	// Specifies the configuration of AD domain.
	// Required if `authType` is **LOCAL_AD**. Make sure that the selected VPC network and the network to which AD
	// belongs can be connected. The object structure is documented below.
	AdDomain ServiceAdDomainPtrInput
	// Specifies the authentication type of Workspace service.
	// The valid values are as follows:
	// + **LITE_AS**: Local authentication.
	// + **LOCAL_AD**: Connect to AD domain.
	AuthType pulumi.StringPtrInput
	// The subnet segments of the dedicated access.
	DedicatedSubnets pulumi.StringArrayInput
	// Specifies the enterprise ID.
	// The enterprise ID is the unique identification in the Workspace service.
	// If omitted, the system will automatically generate an enterprise ID.
	// The ID can contain `1` to `32` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	EnterpriseId pulumi.StringPtrInput
	// Specifies the internet access port.
	// The valid value is range from `1,025` to `65,535`.
	InternetAccessPort pulumi.IntPtrInput
	// The subnet segment of the management component.
	ManagementSubnetCidr pulumi.StringPtrInput
	// The network ID list of subnets that the service have.
	// The subnets corresponding to this parameter must be included in the VPC resource corresponding to `vpcId`.
	// These subnet segments cannot conflict with `172.16.0.0/12`.
	NetworkIds pulumi.StringArrayInput
	// The region in which to register the Workspace service.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the VPC ID to which the service belongs.
	// Changing this will create a new resource.
	VpcId pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//	ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//	ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Specifies the access mode of Workspace service.
// The valid values are as follows:
// + **INTERNET**: internet access.
// + **DEDICATED**: dedicated line access.
// + **BOTH**: both internet access and dedicated access are supported.
func (o ServiceOutput) AccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.AccessMode }).(pulumi.StringOutput)
}

// Specifies the configuration of AD domain.
// Required if `authType` is **LOCAL_AD**. Make sure that the selected VPC network and the network to which AD
// belongs can be connected. The object structure is documented below.
func (o ServiceOutput) AdDomain() ServiceAdDomainOutput {
	return o.ApplyT(func(v *Service) ServiceAdDomainOutput { return v.AdDomain }).(ServiceAdDomainOutput)
}

// Specifies the authentication type of Workspace service.
// The valid values are as follows:
// + **LITE_AS**: Local authentication.
// + **LOCAL_AD**: Connect to AD domain.
func (o ServiceOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.AuthType }).(pulumi.StringPtrOutput)
}

// The subnet segments of the dedicated access.
func (o ServiceOutput) DedicatedSubnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.StringArrayOutput { return v.DedicatedSubnets }).(pulumi.StringArrayOutput)
}

// The desktop security group automatically created under the specified VPC after the service
// is registered. The object structure is documented below.
func (o ServiceOutput) DesktopSecurityGroups() ServiceDesktopSecurityGroupArrayOutput {
	return o.ApplyT(func(v *Service) ServiceDesktopSecurityGroupArrayOutput { return v.DesktopSecurityGroups }).(ServiceDesktopSecurityGroupArrayOutput)
}

// Specifies the enterprise ID.
// The enterprise ID is the unique identification in the Workspace service.
// If omitted, the system will automatically generate an enterprise ID.
// The ID can contain `1` to `32` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
func (o ServiceOutput) EnterpriseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.EnterpriseId }).(pulumi.StringOutput)
}

// The management component security group automatically created under the specified
// VPC after the service is registered. The object structure is documented below.
func (o ServiceOutput) InfrastructureSecurityGroups() ServiceInfrastructureSecurityGroupArrayOutput {
	return o.ApplyT(func(v *Service) ServiceInfrastructureSecurityGroupArrayOutput { return v.InfrastructureSecurityGroups }).(ServiceInfrastructureSecurityGroupArrayOutput)
}

// The internet access address.
// This attribute is returned only when the accessMode is **INTERNET** or **BOTH**.
func (o ServiceOutput) InternetAccessAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.InternetAccessAddress }).(pulumi.StringOutput)
}

// Specifies the internet access port.
// The valid value is range from `1,025` to `65,535`.
func (o ServiceOutput) InternetAccessPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.InternetAccessPort }).(pulumi.IntOutput)
}

// The subnet segment of the management component.
func (o ServiceOutput) ManagementSubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ManagementSubnetCidr }).(pulumi.StringOutput)
}

// The network ID list of subnets that the service have.
// The subnets corresponding to this parameter must be included in the VPC resource corresponding to `vpcId`.
// These subnet segments cannot conflict with `172.16.0.0/12`.
func (o ServiceOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.StringArrayOutput { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

// The region in which to register the Workspace service.
// If omitted, the provider-level region will be used. Changing this will create a new resource.
func (o ServiceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current status of the Workspace service.
func (o ServiceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the VPC ID to which the service belongs.
// Changing this will create a new resource.
func (o ServiceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
