// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Dedicated Load Balancer resource within HuaweiCloud.
//
// ## Example Usage
// ### Basic Loadbalancer
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := DedicatedElb.NewLoadbalancer(ctx, "basic", &DedicatedElb.LoadbalancerArgs{
//				AvailabilityZones: pulumi.StringArray{
//					pulumi.String("cn-north-4a"),
//					pulumi.String("cn-north-4b"),
//				},
//				CrossVpcBackend:     pulumi.Bool(true),
//				Description:         pulumi.String("basic example"),
//				EnterpriseProjectId: pulumi.String("{{ eps_id }}"),
//				Ipv4SubnetId:        pulumi.String("{{ ipv4_subnet_id }}"),
//				L4FlavorId:          pulumi.String("{{ l4_flavor_id }}"),
//				L7FlavorId:          pulumi.String("{{ l7_flavor_id }}"),
//				VpcId:               pulumi.String("{{ vpc_id }}"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Loadbalancer With Existing EIP
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := DedicatedElb.NewLoadbalancer(ctx, "basic", &DedicatedElb.LoadbalancerArgs{
//				AvailabilityZones: pulumi.StringArray{
//					pulumi.String("cn-north-4a"),
//					pulumi.String("cn-north-4b"),
//				},
//				CrossVpcBackend:     pulumi.Bool(true),
//				Description:         pulumi.String("basic example"),
//				EnterpriseProjectId: pulumi.String("{{ eps_id }}"),
//				Ipv4EipId:           pulumi.String("{{ eip_id }}"),
//				Ipv4SubnetId:        pulumi.String("{{ ipv4_subnet_id }}"),
//				Ipv6BandwidthId:     pulumi.String("{{ ipv6_bandwidth_id }}"),
//				Ipv6NetworkId:       pulumi.String("{{ ipv6_network_id }}"),
//				L4FlavorId:          pulumi.String("{{ l4_flavor_id }}"),
//				L7FlavorId:          pulumi.String("{{ l7_flavor_id }}"),
//				VpcId:               pulumi.String("{{ vpc_id }}"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Loadbalancer With EIP
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := DedicatedElb.NewLoadbalancer(ctx, "basic", &DedicatedElb.LoadbalancerArgs{
//				AvailabilityZones: pulumi.StringArray{
//					pulumi.String("cn-north-4a"),
//					pulumi.String("cn-north-4b"),
//				},
//				BandwidthChargeMode: pulumi.String("traffic"),
//				BandwidthSize:       pulumi.Int(10),
//				CrossVpcBackend:     pulumi.Bool(true),
//				Description:         pulumi.String("basic example"),
//				EnterpriseProjectId: pulumi.String("{{ eps_id }}"),
//				Iptype:              pulumi.String("5_bgp"),
//				Ipv4SubnetId:        pulumi.String("{{ ipv4_subnet_id }}"),
//				Ipv6BandwidthId:     pulumi.String("{{ ipv6_bandwidth_id }}"),
//				Ipv6NetworkId:       pulumi.String("{{ ipv6_network_id }}"),
//				L4FlavorId:          pulumi.String("{{ l4_flavor_id }}"),
//				L7FlavorId:          pulumi.String("{{ l7_flavor_id }}"),
//				Sharetype:           pulumi.String("PER"),
//				VpcId:               pulumi.String("{{ vpc_id }}"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ELB loadbalancer can be imported using the loadbalancer ID, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:DedicatedElb/loadbalancer:Loadbalancer loadbalancer_1 5c20fdad-7288-11eb-b817-0255ac10158b
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`ipv6_bandwidth_id`, `iptype`, `bandwidth_charge_mode`, `sharetype` and `bandwidth_size`. It is generally recommended running `terraform plan` after importing a loadbalancer. You can then decide if changes should be applied to the loadbalancer, or the resource definition should be updated to align with the loadbalancer. Also you can ignore changes as below. resource "huaweicloud_elb_loadbalancer" "loadbalancer_1" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	ipv6_bandwidth_id, iptype, bandwidth_charge_mode, sharetype, bandwidth_size,
//
//	]
//
//	} }
type Loadbalancer struct {
	pulumi.CustomResourceState

	// Deprecated: Deprecated
	AutoPay pulumi.StringPtrOutput `pulumi:"autoPay"`
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew pulumi.StringPtrOutput `pulumi:"autoRenew"`
	// Specifies whether autoscaling is enabled. Valid values are **true** and
	// **false**.
	AutoscalingEnabled pulumi.BoolOutput `pulumi:"autoscalingEnabled"`
	// Specifies the list of AZ names. Changing this parameter will create a
	// new resource.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// Bandwidth billing type. Changing this parameter will create a
	// new resource.
	BandwidthChargeMode pulumi.StringOutput `pulumi:"bandwidthChargeMode"`
	// Bandwidth size. Changing this parameter will create a new resource.
	BandwidthSize pulumi.IntOutput `pulumi:"bandwidthSize"`
	// Specifies the charging mode of the ELB loadbalancer.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new resource.
	ChargingMode pulumi.StringOutput `pulumi:"chargingMode"`
	// Enable this if you want to associate the IP addresses of backend servers with
	// your load balancer. Can only be true when updating.
	CrossVpcBackend pulumi.BoolOutput `pulumi:"crossVpcBackend"`
	// Human-readable description for the loadbalancer.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The enterprise project id of the loadbalancer. Changing this
	// creates a new loadbalancer.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// Elastic IP type. Changing this parameter will create a new resource.
	Iptype pulumi.StringOutput `pulumi:"iptype"`
	// The ipv4 address of the load balancer.
	Ipv4Address pulumi.StringOutput `pulumi:"ipv4Address"`
	// The ipv4 eip address of the Load Balancer.
	Ipv4Eip pulumi.StringOutput `pulumi:"ipv4Eip"`
	// The ID of the EIP. Changing this parameter will create a new resource.
	Ipv4EipId pulumi.StringOutput `pulumi:"ipv4EipId"`
	// The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
	// ipv4 address.
	Ipv4SubnetId pulumi.StringPtrOutput `pulumi:"ipv4SubnetId"`
	// The ipv6 address of the Load Balancer.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// The ipv6 bandwidth id. Only support shared bandwidth.
	Ipv6BandwidthId pulumi.StringPtrOutput `pulumi:"ipv6BandwidthId"`
	// The ipv6 eip address of the Load Balancer.
	Ipv6Eip pulumi.StringOutput `pulumi:"ipv6Eip"`
	// The ipv6 eip id of the Load Balancer.
	Ipv6EipId pulumi.StringOutput `pulumi:"ipv6EipId"`
	// The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
	Ipv6NetworkId pulumi.StringPtrOutput `pulumi:"ipv6NetworkId"`
	// The L4 flavor id of the load balancer.
	L4FlavorId pulumi.StringOutput `pulumi:"l4FlavorId"`
	// The L7 flavor id of the load balancer.
	L7FlavorId pulumi.StringOutput `pulumi:"l7FlavorId"`
	// Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
	// This parameter cannot be left blank if there are HTTP or HTTPS listeners.
	MinL7FlavorId pulumi.StringOutput `pulumi:"minL7FlavorId"`
	// Human-readable name for the loadbalancer.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the charging period of the ELB loadbalancer.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Specifies the charging period unit of the ELB loadbalancer.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	PeriodUnit pulumi.StringPtrOutput `pulumi:"periodUnit"`
	// The region in which to create the loadbalancer resource. If omitted, the
	// provider-level region will be used. Changing this creates a new loadbalancer.
	Region pulumi.StringOutput `pulumi:"region"`
	// Bandwidth sharing type. Changing this parameter will create a new resource.
	Sharetype pulumi.StringOutput `pulumi:"sharetype"`
	// The key/value pairs to associate with the loadbalancer.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The vpc on which to create the loadbalancer. Changing this creates a new
	// loadbalancer.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewLoadbalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadbalancer(ctx *pulumi.Context,
	name string, args *LoadbalancerArgs, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZones == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZones'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Loadbalancer
	err := ctx.RegisterResource("huaweicloud:DedicatedElb/loadbalancer:Loadbalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadbalancer gets an existing Loadbalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadbalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadbalancerState, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	var resource Loadbalancer
	err := ctx.ReadResource("huaweicloud:DedicatedElb/loadbalancer:Loadbalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Loadbalancer resources.
type loadbalancerState struct {
	// Deprecated: Deprecated
	AutoPay *string `pulumi:"autoPay"`
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies whether autoscaling is enabled. Valid values are **true** and
	// **false**.
	AutoscalingEnabled *bool `pulumi:"autoscalingEnabled"`
	// Specifies the list of AZ names. Changing this parameter will create a
	// new resource.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Bandwidth billing type. Changing this parameter will create a
	// new resource.
	BandwidthChargeMode *string `pulumi:"bandwidthChargeMode"`
	// Bandwidth size. Changing this parameter will create a new resource.
	BandwidthSize *int `pulumi:"bandwidthSize"`
	// Specifies the charging mode of the ELB loadbalancer.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new resource.
	ChargingMode *string `pulumi:"chargingMode"`
	// Enable this if you want to associate the IP addresses of backend servers with
	// your load balancer. Can only be true when updating.
	CrossVpcBackend *bool `pulumi:"crossVpcBackend"`
	// Human-readable description for the loadbalancer.
	Description *string `pulumi:"description"`
	// The enterprise project id of the loadbalancer. Changing this
	// creates a new loadbalancer.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Elastic IP type. Changing this parameter will create a new resource.
	Iptype *string `pulumi:"iptype"`
	// The ipv4 address of the load balancer.
	Ipv4Address *string `pulumi:"ipv4Address"`
	// The ipv4 eip address of the Load Balancer.
	Ipv4Eip *string `pulumi:"ipv4Eip"`
	// The ID of the EIP. Changing this parameter will create a new resource.
	Ipv4EipId *string `pulumi:"ipv4EipId"`
	// The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
	// ipv4 address.
	Ipv4SubnetId *string `pulumi:"ipv4SubnetId"`
	// The ipv6 address of the Load Balancer.
	Ipv6Address *string `pulumi:"ipv6Address"`
	// The ipv6 bandwidth id. Only support shared bandwidth.
	Ipv6BandwidthId *string `pulumi:"ipv6BandwidthId"`
	// The ipv6 eip address of the Load Balancer.
	Ipv6Eip *string `pulumi:"ipv6Eip"`
	// The ipv6 eip id of the Load Balancer.
	Ipv6EipId *string `pulumi:"ipv6EipId"`
	// The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
	Ipv6NetworkId *string `pulumi:"ipv6NetworkId"`
	// The L4 flavor id of the load balancer.
	L4FlavorId *string `pulumi:"l4FlavorId"`
	// The L7 flavor id of the load balancer.
	L7FlavorId *string `pulumi:"l7FlavorId"`
	// Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
	// This parameter cannot be left blank if there are HTTP or HTTPS listeners.
	MinL7FlavorId *string `pulumi:"minL7FlavorId"`
	// Human-readable name for the loadbalancer.
	Name *string `pulumi:"name"`
	// Specifies the charging period of the ELB loadbalancer.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the ELB loadbalancer.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The region in which to create the loadbalancer resource. If omitted, the
	// provider-level region will be used. Changing this creates a new loadbalancer.
	Region *string `pulumi:"region"`
	// Bandwidth sharing type. Changing this parameter will create a new resource.
	Sharetype *string `pulumi:"sharetype"`
	// The key/value pairs to associate with the loadbalancer.
	Tags map[string]string `pulumi:"tags"`
	// The vpc on which to create the loadbalancer. Changing this creates a new
	// loadbalancer.
	VpcId *string `pulumi:"vpcId"`
}

type LoadbalancerState struct {
	// Deprecated: Deprecated
	AutoPay pulumi.StringPtrInput
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew pulumi.StringPtrInput
	// Specifies whether autoscaling is enabled. Valid values are **true** and
	// **false**.
	AutoscalingEnabled pulumi.BoolPtrInput
	// Specifies the list of AZ names. Changing this parameter will create a
	// new resource.
	AvailabilityZones pulumi.StringArrayInput
	// Bandwidth billing type. Changing this parameter will create a
	// new resource.
	BandwidthChargeMode pulumi.StringPtrInput
	// Bandwidth size. Changing this parameter will create a new resource.
	BandwidthSize pulumi.IntPtrInput
	// Specifies the charging mode of the ELB loadbalancer.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new resource.
	ChargingMode pulumi.StringPtrInput
	// Enable this if you want to associate the IP addresses of backend servers with
	// your load balancer. Can only be true when updating.
	CrossVpcBackend pulumi.BoolPtrInput
	// Human-readable description for the loadbalancer.
	Description pulumi.StringPtrInput
	// The enterprise project id of the loadbalancer. Changing this
	// creates a new loadbalancer.
	EnterpriseProjectId pulumi.StringPtrInput
	// Elastic IP type. Changing this parameter will create a new resource.
	Iptype pulumi.StringPtrInput
	// The ipv4 address of the load balancer.
	Ipv4Address pulumi.StringPtrInput
	// The ipv4 eip address of the Load Balancer.
	Ipv4Eip pulumi.StringPtrInput
	// The ID of the EIP. Changing this parameter will create a new resource.
	Ipv4EipId pulumi.StringPtrInput
	// The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
	// ipv4 address.
	Ipv4SubnetId pulumi.StringPtrInput
	// The ipv6 address of the Load Balancer.
	Ipv6Address pulumi.StringPtrInput
	// The ipv6 bandwidth id. Only support shared bandwidth.
	Ipv6BandwidthId pulumi.StringPtrInput
	// The ipv6 eip address of the Load Balancer.
	Ipv6Eip pulumi.StringPtrInput
	// The ipv6 eip id of the Load Balancer.
	Ipv6EipId pulumi.StringPtrInput
	// The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
	Ipv6NetworkId pulumi.StringPtrInput
	// The L4 flavor id of the load balancer.
	L4FlavorId pulumi.StringPtrInput
	// The L7 flavor id of the load balancer.
	L7FlavorId pulumi.StringPtrInput
	// Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
	// This parameter cannot be left blank if there are HTTP or HTTPS listeners.
	MinL7FlavorId pulumi.StringPtrInput
	// Human-readable name for the loadbalancer.
	Name pulumi.StringPtrInput
	// Specifies the charging period of the ELB loadbalancer.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the ELB loadbalancer.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	PeriodUnit pulumi.StringPtrInput
	// The region in which to create the loadbalancer resource. If omitted, the
	// provider-level region will be used. Changing this creates a new loadbalancer.
	Region pulumi.StringPtrInput
	// Bandwidth sharing type. Changing this parameter will create a new resource.
	Sharetype pulumi.StringPtrInput
	// The key/value pairs to associate with the loadbalancer.
	Tags pulumi.StringMapInput
	// The vpc on which to create the loadbalancer. Changing this creates a new
	// loadbalancer.
	VpcId pulumi.StringPtrInput
}

func (LoadbalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerState)(nil)).Elem()
}

type loadbalancerArgs struct {
	// Deprecated: Deprecated
	AutoPay *string `pulumi:"autoPay"`
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies whether autoscaling is enabled. Valid values are **true** and
	// **false**.
	AutoscalingEnabled *bool `pulumi:"autoscalingEnabled"`
	// Specifies the list of AZ names. Changing this parameter will create a
	// new resource.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Bandwidth billing type. Changing this parameter will create a
	// new resource.
	BandwidthChargeMode *string `pulumi:"bandwidthChargeMode"`
	// Bandwidth size. Changing this parameter will create a new resource.
	BandwidthSize *int `pulumi:"bandwidthSize"`
	// Specifies the charging mode of the ELB loadbalancer.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new resource.
	ChargingMode *string `pulumi:"chargingMode"`
	// Enable this if you want to associate the IP addresses of backend servers with
	// your load balancer. Can only be true when updating.
	CrossVpcBackend *bool `pulumi:"crossVpcBackend"`
	// Human-readable description for the loadbalancer.
	Description *string `pulumi:"description"`
	// The enterprise project id of the loadbalancer. Changing this
	// creates a new loadbalancer.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Elastic IP type. Changing this parameter will create a new resource.
	Iptype *string `pulumi:"iptype"`
	// The ipv4 address of the load balancer.
	Ipv4Address *string `pulumi:"ipv4Address"`
	// The ID of the EIP. Changing this parameter will create a new resource.
	Ipv4EipId *string `pulumi:"ipv4EipId"`
	// The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
	// ipv4 address.
	Ipv4SubnetId *string `pulumi:"ipv4SubnetId"`
	// The ipv6 bandwidth id. Only support shared bandwidth.
	Ipv6BandwidthId *string `pulumi:"ipv6BandwidthId"`
	// The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
	Ipv6NetworkId *string `pulumi:"ipv6NetworkId"`
	// The L4 flavor id of the load balancer.
	L4FlavorId *string `pulumi:"l4FlavorId"`
	// The L7 flavor id of the load balancer.
	L7FlavorId *string `pulumi:"l7FlavorId"`
	// Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
	// This parameter cannot be left blank if there are HTTP or HTTPS listeners.
	MinL7FlavorId *string `pulumi:"minL7FlavorId"`
	// Human-readable name for the loadbalancer.
	Name *string `pulumi:"name"`
	// Specifies the charging period of the ELB loadbalancer.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the ELB loadbalancer.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The region in which to create the loadbalancer resource. If omitted, the
	// provider-level region will be used. Changing this creates a new loadbalancer.
	Region *string `pulumi:"region"`
	// Bandwidth sharing type. Changing this parameter will create a new resource.
	Sharetype *string `pulumi:"sharetype"`
	// The key/value pairs to associate with the loadbalancer.
	Tags map[string]string `pulumi:"tags"`
	// The vpc on which to create the loadbalancer. Changing this creates a new
	// loadbalancer.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Loadbalancer resource.
type LoadbalancerArgs struct {
	// Deprecated: Deprecated
	AutoPay pulumi.StringPtrInput
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew pulumi.StringPtrInput
	// Specifies whether autoscaling is enabled. Valid values are **true** and
	// **false**.
	AutoscalingEnabled pulumi.BoolPtrInput
	// Specifies the list of AZ names. Changing this parameter will create a
	// new resource.
	AvailabilityZones pulumi.StringArrayInput
	// Bandwidth billing type. Changing this parameter will create a
	// new resource.
	BandwidthChargeMode pulumi.StringPtrInput
	// Bandwidth size. Changing this parameter will create a new resource.
	BandwidthSize pulumi.IntPtrInput
	// Specifies the charging mode of the ELB loadbalancer.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new resource.
	ChargingMode pulumi.StringPtrInput
	// Enable this if you want to associate the IP addresses of backend servers with
	// your load balancer. Can only be true when updating.
	CrossVpcBackend pulumi.BoolPtrInput
	// Human-readable description for the loadbalancer.
	Description pulumi.StringPtrInput
	// The enterprise project id of the loadbalancer. Changing this
	// creates a new loadbalancer.
	EnterpriseProjectId pulumi.StringPtrInput
	// Elastic IP type. Changing this parameter will create a new resource.
	Iptype pulumi.StringPtrInput
	// The ipv4 address of the load balancer.
	Ipv4Address pulumi.StringPtrInput
	// The ID of the EIP. Changing this parameter will create a new resource.
	Ipv4EipId pulumi.StringPtrInput
	// The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
	// ipv4 address.
	Ipv4SubnetId pulumi.StringPtrInput
	// The ipv6 bandwidth id. Only support shared bandwidth.
	Ipv6BandwidthId pulumi.StringPtrInput
	// The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
	Ipv6NetworkId pulumi.StringPtrInput
	// The L4 flavor id of the load balancer.
	L4FlavorId pulumi.StringPtrInput
	// The L7 flavor id of the load balancer.
	L7FlavorId pulumi.StringPtrInput
	// Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
	// This parameter cannot be left blank if there are HTTP or HTTPS listeners.
	MinL7FlavorId pulumi.StringPtrInput
	// Human-readable name for the loadbalancer.
	Name pulumi.StringPtrInput
	// Specifies the charging period of the ELB loadbalancer.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the ELB loadbalancer.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new resource.
	PeriodUnit pulumi.StringPtrInput
	// The region in which to create the loadbalancer resource. If omitted, the
	// provider-level region will be used. Changing this creates a new loadbalancer.
	Region pulumi.StringPtrInput
	// Bandwidth sharing type. Changing this parameter will create a new resource.
	Sharetype pulumi.StringPtrInput
	// The key/value pairs to associate with the loadbalancer.
	Tags pulumi.StringMapInput
	// The vpc on which to create the loadbalancer. Changing this creates a new
	// loadbalancer.
	VpcId pulumi.StringPtrInput
}

func (LoadbalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerArgs)(nil)).Elem()
}

type LoadbalancerInput interface {
	pulumi.Input

	ToLoadbalancerOutput() LoadbalancerOutput
	ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput
}

func (*Loadbalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (i *Loadbalancer) ToLoadbalancerOutput() LoadbalancerOutput {
	return i.ToLoadbalancerOutputWithContext(context.Background())
}

func (i *Loadbalancer) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerOutput)
}

// LoadbalancerArrayInput is an input type that accepts LoadbalancerArray and LoadbalancerArrayOutput values.
// You can construct a concrete instance of `LoadbalancerArrayInput` via:
//
//	LoadbalancerArray{ LoadbalancerArgs{...} }
type LoadbalancerArrayInput interface {
	pulumi.Input

	ToLoadbalancerArrayOutput() LoadbalancerArrayOutput
	ToLoadbalancerArrayOutputWithContext(context.Context) LoadbalancerArrayOutput
}

type LoadbalancerArray []LoadbalancerInput

func (LoadbalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Loadbalancer)(nil)).Elem()
}

func (i LoadbalancerArray) ToLoadbalancerArrayOutput() LoadbalancerArrayOutput {
	return i.ToLoadbalancerArrayOutputWithContext(context.Background())
}

func (i LoadbalancerArray) ToLoadbalancerArrayOutputWithContext(ctx context.Context) LoadbalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerArrayOutput)
}

// LoadbalancerMapInput is an input type that accepts LoadbalancerMap and LoadbalancerMapOutput values.
// You can construct a concrete instance of `LoadbalancerMapInput` via:
//
//	LoadbalancerMap{ "key": LoadbalancerArgs{...} }
type LoadbalancerMapInput interface {
	pulumi.Input

	ToLoadbalancerMapOutput() LoadbalancerMapOutput
	ToLoadbalancerMapOutputWithContext(context.Context) LoadbalancerMapOutput
}

type LoadbalancerMap map[string]LoadbalancerInput

func (LoadbalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Loadbalancer)(nil)).Elem()
}

func (i LoadbalancerMap) ToLoadbalancerMapOutput() LoadbalancerMapOutput {
	return i.ToLoadbalancerMapOutputWithContext(context.Background())
}

func (i LoadbalancerMap) ToLoadbalancerMapOutputWithContext(ctx context.Context) LoadbalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerMapOutput)
}

type LoadbalancerOutput struct{ *pulumi.OutputState }

func (LoadbalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerOutput) ToLoadbalancerOutput() LoadbalancerOutput {
	return o
}

func (o LoadbalancerOutput) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return o
}

// Deprecated: Deprecated
func (o LoadbalancerOutput) AutoPay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.AutoPay }).(pulumi.StringPtrOutput)
}

// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
func (o LoadbalancerOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

// Specifies whether autoscaling is enabled. Valid values are **true** and
// **false**.
func (o LoadbalancerOutput) AutoscalingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.BoolOutput { return v.AutoscalingEnabled }).(pulumi.BoolOutput)
}

// Specifies the list of AZ names. Changing this parameter will create a
// new resource.
func (o LoadbalancerOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// Bandwidth billing type. Changing this parameter will create a
// new resource.
func (o LoadbalancerOutput) BandwidthChargeMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.BandwidthChargeMode }).(pulumi.StringOutput)
}

// Bandwidth size. Changing this parameter will create a new resource.
func (o LoadbalancerOutput) BandwidthSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.IntOutput { return v.BandwidthSize }).(pulumi.IntOutput)
}

// Specifies the charging mode of the ELB loadbalancer.
// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
// Changing this parameter will create a new resource.
func (o LoadbalancerOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

// Enable this if you want to associate the IP addresses of backend servers with
// your load balancer. Can only be true when updating.
func (o LoadbalancerOutput) CrossVpcBackend() pulumi.BoolOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.BoolOutput { return v.CrossVpcBackend }).(pulumi.BoolOutput)
}

// Human-readable description for the loadbalancer.
func (o LoadbalancerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The enterprise project id of the loadbalancer. Changing this
// creates a new loadbalancer.
func (o LoadbalancerOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Elastic IP type. Changing this parameter will create a new resource.
func (o LoadbalancerOutput) Iptype() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Iptype }).(pulumi.StringOutput)
}

// The ipv4 address of the load balancer.
func (o LoadbalancerOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Ipv4Address }).(pulumi.StringOutput)
}

// The ipv4 eip address of the Load Balancer.
func (o LoadbalancerOutput) Ipv4Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Ipv4Eip }).(pulumi.StringOutput)
}

// The ID of the EIP. Changing this parameter will create a new resource.
func (o LoadbalancerOutput) Ipv4EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Ipv4EipId }).(pulumi.StringOutput)
}

// The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
// ipv4 address.
func (o LoadbalancerOutput) Ipv4SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.Ipv4SubnetId }).(pulumi.StringPtrOutput)
}

// The ipv6 address of the Load Balancer.
func (o LoadbalancerOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

// The ipv6 bandwidth id. Only support shared bandwidth.
func (o LoadbalancerOutput) Ipv6BandwidthId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.Ipv6BandwidthId }).(pulumi.StringPtrOutput)
}

// The ipv6 eip address of the Load Balancer.
func (o LoadbalancerOutput) Ipv6Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Ipv6Eip }).(pulumi.StringOutput)
}

// The ipv6 eip id of the Load Balancer.
func (o LoadbalancerOutput) Ipv6EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Ipv6EipId }).(pulumi.StringOutput)
}

// The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
func (o LoadbalancerOutput) Ipv6NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.Ipv6NetworkId }).(pulumi.StringPtrOutput)
}

// The L4 flavor id of the load balancer.
func (o LoadbalancerOutput) L4FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.L4FlavorId }).(pulumi.StringOutput)
}

// The L7 flavor id of the load balancer.
func (o LoadbalancerOutput) L7FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.L7FlavorId }).(pulumi.StringOutput)
}

// Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
// This parameter cannot be left blank if there are HTTP or HTTPS listeners.
func (o LoadbalancerOutput) MinL7FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.MinL7FlavorId }).(pulumi.StringOutput)
}

// Human-readable name for the loadbalancer.
func (o LoadbalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the charging period of the ELB loadbalancer.
// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
// This parameter is mandatory if `chargingMode` is set to **prePaid**.
// Changing this parameter will create a new resource.
func (o LoadbalancerOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Specifies the charging period unit of the ELB loadbalancer.
// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
// Changing this parameter will create a new resource.
func (o LoadbalancerOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

// The region in which to create the loadbalancer resource. If omitted, the
// provider-level region will be used. Changing this creates a new loadbalancer.
func (o LoadbalancerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Bandwidth sharing type. Changing this parameter will create a new resource.
func (o LoadbalancerOutput) Sharetype() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Sharetype }).(pulumi.StringOutput)
}

// The key/value pairs to associate with the loadbalancer.
func (o LoadbalancerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The vpc on which to create the loadbalancer. Changing this creates a new
// loadbalancer.
func (o LoadbalancerOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type LoadbalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadbalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerArrayOutput) ToLoadbalancerArrayOutput() LoadbalancerArrayOutput {
	return o
}

func (o LoadbalancerArrayOutput) ToLoadbalancerArrayOutputWithContext(ctx context.Context) LoadbalancerArrayOutput {
	return o
}

func (o LoadbalancerArrayOutput) Index(i pulumi.IntInput) LoadbalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Loadbalancer {
		return vs[0].([]*Loadbalancer)[vs[1].(int)]
	}).(LoadbalancerOutput)
}

type LoadbalancerMapOutput struct{ *pulumi.OutputState }

func (LoadbalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerMapOutput) ToLoadbalancerMapOutput() LoadbalancerMapOutput {
	return o
}

func (o LoadbalancerMapOutput) ToLoadbalancerMapOutputWithContext(ctx context.Context) LoadbalancerMapOutput {
	return o
}

func (o LoadbalancerMapOutput) MapIndex(k pulumi.StringInput) LoadbalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Loadbalancer {
		return vs[0].(map[string]*Loadbalancer)[vs[1].(string)]
	}).(LoadbalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerInput)(nil)).Elem(), &Loadbalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerArrayInput)(nil)).Elem(), LoadbalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerMapInput)(nil)).Elem(), LoadbalancerMap{})
	pulumi.RegisterOutputType(LoadbalancerOutput{})
	pulumi.RegisterOutputType(LoadbalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadbalancerMapOutput{})
}
