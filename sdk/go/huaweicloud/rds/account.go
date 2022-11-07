// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages RDS Mysql account resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			instanceId := cfg.RequireObject("instanceId")
//			_, err := Rds.NewAccount(ctx, "test", &Rds.AccountArgs{
//				InstanceId: pulumi.Any(instanceId),
//				Password:   pulumi.String("Test@12345678"),
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
// RDS account can be imported using the `instance id` and `account name`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/account:Account user_1 instance_id/account_name
//
// ```
type Account struct {
	pulumi.CustomResourceState

	// Specifies the rds instance id. Changing this will create a new resource.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the username of the db account. Only lowercase letters, digits,
	// hyphens (-), and userscores (_) are allowed. Changing this will create a new resource.
	// + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
	// + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the password of the db account. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
	// different from name or name spelled backwards.
	Password pulumi.StringOutput `pulumi:"password"`
	// The region in which to create the rds account resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	var resource Account
	err := ctx.RegisterResource("huaweicloud:Rds/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("huaweicloud:Rds/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Specifies the rds instance id. Changing this will create a new resource.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the username of the db account. Only lowercase letters, digits,
	// hyphens (-), and userscores (_) are allowed. Changing this will create a new resource.
	// + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
	// + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
	Name *string `pulumi:"name"`
	// Specifies the password of the db account. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
	// different from name or name spelled backwards.
	Password *string `pulumi:"password"`
	// The region in which to create the rds account resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
}

type AccountState struct {
	// Specifies the rds instance id. Changing this will create a new resource.
	InstanceId pulumi.StringPtrInput
	// Specifies the username of the db account. Only lowercase letters, digits,
	// hyphens (-), and userscores (_) are allowed. Changing this will create a new resource.
	// + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
	// + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
	Name pulumi.StringPtrInput
	// Specifies the password of the db account. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
	// different from name or name spelled backwards.
	Password pulumi.StringPtrInput
	// The region in which to create the rds account resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Specifies the rds instance id. Changing this will create a new resource.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the username of the db account. Only lowercase letters, digits,
	// hyphens (-), and userscores (_) are allowed. Changing this will create a new resource.
	// + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
	// + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
	Name *string `pulumi:"name"`
	// Specifies the password of the db account. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
	// different from name or name spelled backwards.
	Password string `pulumi:"password"`
	// The region in which to create the rds account resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Specifies the rds instance id. Changing this will create a new resource.
	InstanceId pulumi.StringInput
	// Specifies the username of the db account. Only lowercase letters, digits,
	// hyphens (-), and userscores (_) are allowed. Changing this will create a new resource.
	// + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
	// + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
	Name pulumi.StringPtrInput
	// Specifies the password of the db account. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
	// different from name or name spelled backwards.
	Password pulumi.StringInput
	// The region in which to create the rds account resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//	AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//	AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

// Specifies the rds instance id. Changing this will create a new resource.
func (o AccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the username of the db account. Only lowercase letters, digits,
// hyphens (-), and userscores (_) are allowed. Changing this will create a new resource.
// + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
// + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the password of the db account. The parameter must be 8 to 32 characters
// long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
// different from name or name spelled backwards.
func (o AccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The region in which to create the rds account resource. If omitted, the
// provider-level region will be used. Changing this creates a new resource.
func (o AccountOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Account {
		return vs[0].([]*Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Account {
		return vs[0].(map[string]*Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountArrayInput)(nil)).Elem(), AccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountMapInput)(nil)).Elem(), AccountMap{})
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}