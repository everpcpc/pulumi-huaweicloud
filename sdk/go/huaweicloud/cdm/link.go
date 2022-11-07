// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a link resource within HuaweiCloud. A link enables the CDM cluster to read data from and write data to
//
//	a data source.
//
// ## Example Usage
// ### Link to MySql
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cdm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			mysqlLinkName := cfg.RequireObject("mysqlLinkName")
//			cdmClusterId := cfg.RequireObject("cdmClusterId")
//			mysqlHost := cfg.RequireObject("mysqlHost")
//			dbName := cfg.RequireObject("dbName")
//			dbPassword := cfg.RequireObject("dbPassword")
//			_, err := Cdm.NewLink(ctx, "mysqlLink", &Cdm.LinkArgs{
//				Connector: pulumi.String("generic-jdbc-connector"),
//				ClusterId: pulumi.Any(cdmClusterId),
//				Config: pulumi.StringMap{
//					"databaseType": pulumi.String("MYSQL"),
//					"host":         pulumi.Any(mysqlHost),
//					"port":         pulumi.String("3306"),
//					"database":     pulumi.Any(dbName),
//					"username":     pulumi.Any(_var.Username),
//				},
//				Password: pulumi.Any(dbPassword),
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
// The link can be imported by `id`, It is composed of the ID of CDM cluster and the name of job, separated by a slash.
//
// For example,
//
// ```sh
//
//	$ pulumi import huaweicloud:Cdm/link:Link test b11b407c-e604-4e8d-8bc4-92398320b847/linkName
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`password` and `secret_key`. It is generally recommended running `terraform plan` after importing an instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. resource "huaweicloud_cdm_link" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	password, secret_key,
//
//	]
//
//	} }
type Link struct {
	pulumi.CustomResourceState

	// Specifies access key for accessing the data sources.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// Specifies the id of CDM cluster which this link belongs to.
	// Changing this parameter will create a new resource.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Specifies the link configuration parameters. Each type of the data source to be connected
	// has different configuration parameters, please refer to the document link below.
	Config pulumi.StringMapOutput `pulumi:"config"`
	// Specifies the connector that is classified based on the type of the data
	// source to be connected. Changing this parameter will create a new resource. The options are as follows:
	Connector pulumi.StringOutput `pulumi:"connector"`
	// Specifies whether to enable the link. The default value is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the name of the link.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the password for accessing the data sources.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies security key for accessing the data sources.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
}

// NewLink registers a new resource with the given unique name, arguments, and options.
func NewLink(ctx *pulumi.Context,
	name string, args *LinkArgs, opts ...pulumi.ResourceOption) (*Link, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Connector == nil {
		return nil, errors.New("invalid value for required argument 'Connector'")
	}
	var resource Link
	err := ctx.RegisterResource("huaweicloud:Cdm/link:Link", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLink gets an existing Link resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkState, opts ...pulumi.ResourceOption) (*Link, error) {
	var resource Link
	err := ctx.ReadResource("huaweicloud:Cdm/link:Link", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Link resources.
type linkState struct {
	// Specifies access key for accessing the data sources.
	AccessKey *string `pulumi:"accessKey"`
	// Specifies the id of CDM cluster which this link belongs to.
	// Changing this parameter will create a new resource.
	ClusterId *string `pulumi:"clusterId"`
	// Specifies the link configuration parameters. Each type of the data source to be connected
	// has different configuration parameters, please refer to the document link below.
	Config map[string]string `pulumi:"config"`
	// Specifies the connector that is classified based on the type of the data
	// source to be connected. Changing this parameter will create a new resource. The options are as follows:
	Connector *string `pulumi:"connector"`
	// Specifies whether to enable the link. The default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the link.
	Name *string `pulumi:"name"`
	// Specifies the password for accessing the data sources.
	Password *string `pulumi:"password"`
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies security key for accessing the data sources.
	SecretKey *string `pulumi:"secretKey"`
}

type LinkState struct {
	// Specifies access key for accessing the data sources.
	AccessKey pulumi.StringPtrInput
	// Specifies the id of CDM cluster which this link belongs to.
	// Changing this parameter will create a new resource.
	ClusterId pulumi.StringPtrInput
	// Specifies the link configuration parameters. Each type of the data source to be connected
	// has different configuration parameters, please refer to the document link below.
	Config pulumi.StringMapInput
	// Specifies the connector that is classified based on the type of the data
	// source to be connected. Changing this parameter will create a new resource. The options are as follows:
	Connector pulumi.StringPtrInput
	// Specifies whether to enable the link. The default value is `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the link.
	Name pulumi.StringPtrInput
	// Specifies the password for accessing the data sources.
	Password pulumi.StringPtrInput
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies security key for accessing the data sources.
	SecretKey pulumi.StringPtrInput
}

func (LinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkState)(nil)).Elem()
}

type linkArgs struct {
	// Specifies access key for accessing the data sources.
	AccessKey *string `pulumi:"accessKey"`
	// Specifies the id of CDM cluster which this link belongs to.
	// Changing this parameter will create a new resource.
	ClusterId string `pulumi:"clusterId"`
	// Specifies the link configuration parameters. Each type of the data source to be connected
	// has different configuration parameters, please refer to the document link below.
	Config map[string]string `pulumi:"config"`
	// Specifies the connector that is classified based on the type of the data
	// source to be connected. Changing this parameter will create a new resource. The options are as follows:
	Connector string `pulumi:"connector"`
	// Specifies whether to enable the link. The default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the link.
	Name *string `pulumi:"name"`
	// Specifies the password for accessing the data sources.
	Password *string `pulumi:"password"`
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies security key for accessing the data sources.
	SecretKey *string `pulumi:"secretKey"`
}

// The set of arguments for constructing a Link resource.
type LinkArgs struct {
	// Specifies access key for accessing the data sources.
	AccessKey pulumi.StringPtrInput
	// Specifies the id of CDM cluster which this link belongs to.
	// Changing this parameter will create a new resource.
	ClusterId pulumi.StringInput
	// Specifies the link configuration parameters. Each type of the data source to be connected
	// has different configuration parameters, please refer to the document link below.
	Config pulumi.StringMapInput
	// Specifies the connector that is classified based on the type of the data
	// source to be connected. Changing this parameter will create a new resource. The options are as follows:
	Connector pulumi.StringInput
	// Specifies whether to enable the link. The default value is `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the link.
	Name pulumi.StringPtrInput
	// Specifies the password for accessing the data sources.
	Password pulumi.StringPtrInput
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies security key for accessing the data sources.
	SecretKey pulumi.StringPtrInput
}

func (LinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkArgs)(nil)).Elem()
}

type LinkInput interface {
	pulumi.Input

	ToLinkOutput() LinkOutput
	ToLinkOutputWithContext(ctx context.Context) LinkOutput
}

func (*Link) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (i *Link) ToLinkOutput() LinkOutput {
	return i.ToLinkOutputWithContext(context.Background())
}

func (i *Link) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkOutput)
}

// LinkArrayInput is an input type that accepts LinkArray and LinkArrayOutput values.
// You can construct a concrete instance of `LinkArrayInput` via:
//
//	LinkArray{ LinkArgs{...} }
type LinkArrayInput interface {
	pulumi.Input

	ToLinkArrayOutput() LinkArrayOutput
	ToLinkArrayOutputWithContext(context.Context) LinkArrayOutput
}

type LinkArray []LinkInput

func (LinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Link)(nil)).Elem()
}

func (i LinkArray) ToLinkArrayOutput() LinkArrayOutput {
	return i.ToLinkArrayOutputWithContext(context.Background())
}

func (i LinkArray) ToLinkArrayOutputWithContext(ctx context.Context) LinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkArrayOutput)
}

// LinkMapInput is an input type that accepts LinkMap and LinkMapOutput values.
// You can construct a concrete instance of `LinkMapInput` via:
//
//	LinkMap{ "key": LinkArgs{...} }
type LinkMapInput interface {
	pulumi.Input

	ToLinkMapOutput() LinkMapOutput
	ToLinkMapOutputWithContext(context.Context) LinkMapOutput
}

type LinkMap map[string]LinkInput

func (LinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Link)(nil)).Elem()
}

func (i LinkMap) ToLinkMapOutput() LinkMapOutput {
	return i.ToLinkMapOutputWithContext(context.Background())
}

func (i LinkMap) ToLinkMapOutputWithContext(ctx context.Context) LinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkMapOutput)
}

type LinkOutput struct{ *pulumi.OutputState }

func (LinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (o LinkOutput) ToLinkOutput() LinkOutput {
	return o
}

func (o LinkOutput) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return o
}

// Specifies access key for accessing the data sources.
func (o LinkOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Link) pulumi.StringPtrOutput { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// Specifies the id of CDM cluster which this link belongs to.
// Changing this parameter will create a new resource.
func (o LinkOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Specifies the link configuration parameters. Each type of the data source to be connected
// has different configuration parameters, please refer to the document link below.
func (o LinkOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Link) pulumi.StringMapOutput { return v.Config }).(pulumi.StringMapOutput)
}

// Specifies the connector that is classified based on the type of the data
// source to be connected. Changing this parameter will create a new resource. The options are as follows:
func (o LinkOutput) Connector() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Connector }).(pulumi.StringOutput)
}

// Specifies whether to enable the link. The default value is `true`.
func (o LinkOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Link) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies the name of the link.
func (o LinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the password for accessing the data sources.
func (o LinkOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Link) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The region in which to create the resource. If omitted, the
// provider-level region will be used. Changing this parameter will create a new resource.
func (o LinkOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies security key for accessing the data sources.
func (o LinkOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Link) pulumi.StringPtrOutput { return v.SecretKey }).(pulumi.StringPtrOutput)
}

type LinkArrayOutput struct{ *pulumi.OutputState }

func (LinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Link)(nil)).Elem()
}

func (o LinkArrayOutput) ToLinkArrayOutput() LinkArrayOutput {
	return o
}

func (o LinkArrayOutput) ToLinkArrayOutputWithContext(ctx context.Context) LinkArrayOutput {
	return o
}

func (o LinkArrayOutput) Index(i pulumi.IntInput) LinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Link {
		return vs[0].([]*Link)[vs[1].(int)]
	}).(LinkOutput)
}

type LinkMapOutput struct{ *pulumi.OutputState }

func (LinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Link)(nil)).Elem()
}

func (o LinkMapOutput) ToLinkMapOutput() LinkMapOutput {
	return o
}

func (o LinkMapOutput) ToLinkMapOutputWithContext(ctx context.Context) LinkMapOutput {
	return o
}

func (o LinkMapOutput) MapIndex(k pulumi.StringInput) LinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Link {
		return vs[0].(map[string]*Link)[vs[1].(string)]
	}).(LinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkInput)(nil)).Elem(), &Link{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkArrayInput)(nil)).Elem(), LinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkMapInput)(nil)).Elem(), LinkMap{})
	pulumi.RegisterOutputType(LinkOutput{})
	pulumi.RegisterOutputType(LinkArrayOutput{})
	pulumi.RegisterOutputType(LinkMapOutput{})
}