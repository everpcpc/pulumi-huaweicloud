// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspace

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Workspace desktop resource within HuaweiCloud.
//
// ## Example Usage
//
// ## Import
//
// Desktops can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Workspace/desktop:Desktop test 339d2539-e945-4090-a08d-c16badc0c6bb
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response. The missing attributes include`nic` and `user_email`. It is generally recommended running `terraform plan` after importing a desktop. You can then decide if changes should be applied to the desktop, or the resource definition should be updated to align with the desktop. Also you can ignore changes as below. resource "huaweicloud_workspace_desktop" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	user_email, nic,
//
//	]
//
//	} }
type Desktop struct {
	pulumi.CustomResourceState

	// Specifies the availability zone where the desktop is located.
	// Changing this will create a new resource.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Specifies the configuration of data volumes.
	// The object structure is documented below.
	DataVolumes DesktopDataVolumeArrayOutput `pulumi:"dataVolumes"`
	// Specifies whether to delete user associated with this desktop after deleting it.
	// The user can only be successfully deleted if the user has no other desktops.
	DeleteUser pulumi.BoolPtrOutput `pulumi:"deleteUser"`
	// Specifies whether to send emails to user mailbox during important
	// operations.
	// Defaults to **false**. Changing this will create a new resource.
	EmailNotification pulumi.BoolPtrOutput `pulumi:"emailNotification"`
	// Specifies the flavor ID of desktop.
	FlavorId pulumi.StringOutput `pulumi:"flavorId"`
	// Specifies the image ID to create the desktop.
	// Changing this will create a new resource.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// Specifies the image type. The valid values are as follows:
	// + **market**: The market image.
	// + **gold**: The public image.
	// + **private**: The private image.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// Specifies the desktop name.
	// The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
	// The name must start with a letter or digit and cannot end with a hyphen.
	// Changing this will create a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the NIC information corresponding to the desktop.
	// The object structure is documented below. Changing this will create a new resource.
	Nics DesktopNicArrayOutput `pulumi:"nics"`
	// The region in which to create the Workspace desktop resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the configuration of system volume.
	// The object structure is documented below.
	RootVolume DesktopRootVolumeOutput `pulumi:"rootVolume"`
	// Specifies the ID list of security groups.
	// In addition to the custom security group, it must also contain a security group called **WorkspaceUserSecurityGroup**.
	// Changing this will create a new resource.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Specifies the key/value pairs of the desktop.
	// Changing this will create a new resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the user email.
	// Some operations on the desktop (such as creation, deletion) will notify the user by sending an email.
	// Changing this will create a new resource.
	UserEmail pulumi.StringOutput `pulumi:"userEmail"`
	// Specifies the user group to which the desktop belongs.
	// The valid values are as follows:
	// + **sudo**: Linux administrator group.
	// + **default**: Linux default user group.
	// + **administrators**: Windows administrator group.
	// + **users**: Windows standard user group.
	UserGroup pulumi.StringOutput `pulumi:"userGroup"`
	// Specifies the user name to which the desktop belongs.
	// The name can contain `1` to `20` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The name must start with a letter. Changing this will create a new resource.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Specifies the VPC ID to which the desktop belongs.
	// Changing this will create a new resource.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDesktop registers a new resource with the given unique name, arguments, and options.
func NewDesktop(ctx *pulumi.Context,
	name string, args *DesktopArgs, opts ...pulumi.ResourceOption) (*Desktop, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlavorId == nil {
		return nil, errors.New("invalid value for required argument 'FlavorId'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.ImageType == nil {
		return nil, errors.New("invalid value for required argument 'ImageType'")
	}
	if args.RootVolume == nil {
		return nil, errors.New("invalid value for required argument 'RootVolume'")
	}
	if args.UserEmail == nil {
		return nil, errors.New("invalid value for required argument 'UserEmail'")
	}
	if args.UserGroup == nil {
		return nil, errors.New("invalid value for required argument 'UserGroup'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Desktop
	err := ctx.RegisterResource("huaweicloud:Workspace/desktop:Desktop", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDesktop gets an existing Desktop resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDesktop(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DesktopState, opts ...pulumi.ResourceOption) (*Desktop, error) {
	var resource Desktop
	err := ctx.ReadResource("huaweicloud:Workspace/desktop:Desktop", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Desktop resources.
type desktopState struct {
	// Specifies the availability zone where the desktop is located.
	// Changing this will create a new resource.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the configuration of data volumes.
	// The object structure is documented below.
	DataVolumes []DesktopDataVolume `pulumi:"dataVolumes"`
	// Specifies whether to delete user associated with this desktop after deleting it.
	// The user can only be successfully deleted if the user has no other desktops.
	DeleteUser *bool `pulumi:"deleteUser"`
	// Specifies whether to send emails to user mailbox during important
	// operations.
	// Defaults to **false**. Changing this will create a new resource.
	EmailNotification *bool `pulumi:"emailNotification"`
	// Specifies the flavor ID of desktop.
	FlavorId *string `pulumi:"flavorId"`
	// Specifies the image ID to create the desktop.
	// Changing this will create a new resource.
	ImageId *string `pulumi:"imageId"`
	// Specifies the image type. The valid values are as follows:
	// + **market**: The market image.
	// + **gold**: The public image.
	// + **private**: The private image.
	ImageType *string `pulumi:"imageType"`
	// Specifies the desktop name.
	// The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
	// The name must start with a letter or digit and cannot end with a hyphen.
	// Changing this will create a new resource.
	Name *string `pulumi:"name"`
	// Specifies the NIC information corresponding to the desktop.
	// The object structure is documented below. Changing this will create a new resource.
	Nics []DesktopNic `pulumi:"nics"`
	// The region in which to create the Workspace desktop resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the configuration of system volume.
	// The object structure is documented below.
	RootVolume *DesktopRootVolume `pulumi:"rootVolume"`
	// Specifies the ID list of security groups.
	// In addition to the custom security group, it must also contain a security group called **WorkspaceUserSecurityGroup**.
	// Changing this will create a new resource.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Specifies the key/value pairs of the desktop.
	// Changing this will create a new resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the user email.
	// Some operations on the desktop (such as creation, deletion) will notify the user by sending an email.
	// Changing this will create a new resource.
	UserEmail *string `pulumi:"userEmail"`
	// Specifies the user group to which the desktop belongs.
	// The valid values are as follows:
	// + **sudo**: Linux administrator group.
	// + **default**: Linux default user group.
	// + **administrators**: Windows administrator group.
	// + **users**: Windows standard user group.
	UserGroup *string `pulumi:"userGroup"`
	// Specifies the user name to which the desktop belongs.
	// The name can contain `1` to `20` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The name must start with a letter. Changing this will create a new resource.
	UserName *string `pulumi:"userName"`
	// Specifies the VPC ID to which the desktop belongs.
	// Changing this will create a new resource.
	VpcId *string `pulumi:"vpcId"`
}

type DesktopState struct {
	// Specifies the availability zone where the desktop is located.
	// Changing this will create a new resource.
	AvailabilityZone pulumi.StringPtrInput
	// Specifies the configuration of data volumes.
	// The object structure is documented below.
	DataVolumes DesktopDataVolumeArrayInput
	// Specifies whether to delete user associated with this desktop after deleting it.
	// The user can only be successfully deleted if the user has no other desktops.
	DeleteUser pulumi.BoolPtrInput
	// Specifies whether to send emails to user mailbox during important
	// operations.
	// Defaults to **false**. Changing this will create a new resource.
	EmailNotification pulumi.BoolPtrInput
	// Specifies the flavor ID of desktop.
	FlavorId pulumi.StringPtrInput
	// Specifies the image ID to create the desktop.
	// Changing this will create a new resource.
	ImageId pulumi.StringPtrInput
	// Specifies the image type. The valid values are as follows:
	// + **market**: The market image.
	// + **gold**: The public image.
	// + **private**: The private image.
	ImageType pulumi.StringPtrInput
	// Specifies the desktop name.
	// The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
	// The name must start with a letter or digit and cannot end with a hyphen.
	// Changing this will create a new resource.
	Name pulumi.StringPtrInput
	// Specifies the NIC information corresponding to the desktop.
	// The object structure is documented below. Changing this will create a new resource.
	Nics DesktopNicArrayInput
	// The region in which to create the Workspace desktop resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the configuration of system volume.
	// The object structure is documented below.
	RootVolume DesktopRootVolumePtrInput
	// Specifies the ID list of security groups.
	// In addition to the custom security group, it must also contain a security group called **WorkspaceUserSecurityGroup**.
	// Changing this will create a new resource.
	SecurityGroups pulumi.StringArrayInput
	// Specifies the key/value pairs of the desktop.
	// Changing this will create a new resource.
	Tags pulumi.StringMapInput
	// Specifies the user email.
	// Some operations on the desktop (such as creation, deletion) will notify the user by sending an email.
	// Changing this will create a new resource.
	UserEmail pulumi.StringPtrInput
	// Specifies the user group to which the desktop belongs.
	// The valid values are as follows:
	// + **sudo**: Linux administrator group.
	// + **default**: Linux default user group.
	// + **administrators**: Windows administrator group.
	// + **users**: Windows standard user group.
	UserGroup pulumi.StringPtrInput
	// Specifies the user name to which the desktop belongs.
	// The name can contain `1` to `20` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The name must start with a letter. Changing this will create a new resource.
	UserName pulumi.StringPtrInput
	// Specifies the VPC ID to which the desktop belongs.
	// Changing this will create a new resource.
	VpcId pulumi.StringPtrInput
}

func (DesktopState) ElementType() reflect.Type {
	return reflect.TypeOf((*desktopState)(nil)).Elem()
}

type desktopArgs struct {
	// Specifies the availability zone where the desktop is located.
	// Changing this will create a new resource.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the configuration of data volumes.
	// The object structure is documented below.
	DataVolumes []DesktopDataVolume `pulumi:"dataVolumes"`
	// Specifies whether to delete user associated with this desktop after deleting it.
	// The user can only be successfully deleted if the user has no other desktops.
	DeleteUser *bool `pulumi:"deleteUser"`
	// Specifies whether to send emails to user mailbox during important
	// operations.
	// Defaults to **false**. Changing this will create a new resource.
	EmailNotification *bool `pulumi:"emailNotification"`
	// Specifies the flavor ID of desktop.
	FlavorId string `pulumi:"flavorId"`
	// Specifies the image ID to create the desktop.
	// Changing this will create a new resource.
	ImageId string `pulumi:"imageId"`
	// Specifies the image type. The valid values are as follows:
	// + **market**: The market image.
	// + **gold**: The public image.
	// + **private**: The private image.
	ImageType string `pulumi:"imageType"`
	// Specifies the desktop name.
	// The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
	// The name must start with a letter or digit and cannot end with a hyphen.
	// Changing this will create a new resource.
	Name *string `pulumi:"name"`
	// Specifies the NIC information corresponding to the desktop.
	// The object structure is documented below. Changing this will create a new resource.
	Nics []DesktopNic `pulumi:"nics"`
	// The region in which to create the Workspace desktop resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the configuration of system volume.
	// The object structure is documented below.
	RootVolume DesktopRootVolume `pulumi:"rootVolume"`
	// Specifies the ID list of security groups.
	// In addition to the custom security group, it must also contain a security group called **WorkspaceUserSecurityGroup**.
	// Changing this will create a new resource.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Specifies the key/value pairs of the desktop.
	// Changing this will create a new resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the user email.
	// Some operations on the desktop (such as creation, deletion) will notify the user by sending an email.
	// Changing this will create a new resource.
	UserEmail string `pulumi:"userEmail"`
	// Specifies the user group to which the desktop belongs.
	// The valid values are as follows:
	// + **sudo**: Linux administrator group.
	// + **default**: Linux default user group.
	// + **administrators**: Windows administrator group.
	// + **users**: Windows standard user group.
	UserGroup string `pulumi:"userGroup"`
	// Specifies the user name to which the desktop belongs.
	// The name can contain `1` to `20` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The name must start with a letter. Changing this will create a new resource.
	UserName string `pulumi:"userName"`
	// Specifies the VPC ID to which the desktop belongs.
	// Changing this will create a new resource.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Desktop resource.
type DesktopArgs struct {
	// Specifies the availability zone where the desktop is located.
	// Changing this will create a new resource.
	AvailabilityZone pulumi.StringPtrInput
	// Specifies the configuration of data volumes.
	// The object structure is documented below.
	DataVolumes DesktopDataVolumeArrayInput
	// Specifies whether to delete user associated with this desktop after deleting it.
	// The user can only be successfully deleted if the user has no other desktops.
	DeleteUser pulumi.BoolPtrInput
	// Specifies whether to send emails to user mailbox during important
	// operations.
	// Defaults to **false**. Changing this will create a new resource.
	EmailNotification pulumi.BoolPtrInput
	// Specifies the flavor ID of desktop.
	FlavorId pulumi.StringInput
	// Specifies the image ID to create the desktop.
	// Changing this will create a new resource.
	ImageId pulumi.StringInput
	// Specifies the image type. The valid values are as follows:
	// + **market**: The market image.
	// + **gold**: The public image.
	// + **private**: The private image.
	ImageType pulumi.StringInput
	// Specifies the desktop name.
	// The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
	// The name must start with a letter or digit and cannot end with a hyphen.
	// Changing this will create a new resource.
	Name pulumi.StringPtrInput
	// Specifies the NIC information corresponding to the desktop.
	// The object structure is documented below. Changing this will create a new resource.
	Nics DesktopNicArrayInput
	// The region in which to create the Workspace desktop resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the configuration of system volume.
	// The object structure is documented below.
	RootVolume DesktopRootVolumeInput
	// Specifies the ID list of security groups.
	// In addition to the custom security group, it must also contain a security group called **WorkspaceUserSecurityGroup**.
	// Changing this will create a new resource.
	SecurityGroups pulumi.StringArrayInput
	// Specifies the key/value pairs of the desktop.
	// Changing this will create a new resource.
	Tags pulumi.StringMapInput
	// Specifies the user email.
	// Some operations on the desktop (such as creation, deletion) will notify the user by sending an email.
	// Changing this will create a new resource.
	UserEmail pulumi.StringInput
	// Specifies the user group to which the desktop belongs.
	// The valid values are as follows:
	// + **sudo**: Linux administrator group.
	// + **default**: Linux default user group.
	// + **administrators**: Windows administrator group.
	// + **users**: Windows standard user group.
	UserGroup pulumi.StringInput
	// Specifies the user name to which the desktop belongs.
	// The name can contain `1` to `20` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The name must start with a letter. Changing this will create a new resource.
	UserName pulumi.StringInput
	// Specifies the VPC ID to which the desktop belongs.
	// Changing this will create a new resource.
	VpcId pulumi.StringInput
}

func (DesktopArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*desktopArgs)(nil)).Elem()
}

type DesktopInput interface {
	pulumi.Input

	ToDesktopOutput() DesktopOutput
	ToDesktopOutputWithContext(ctx context.Context) DesktopOutput
}

func (*Desktop) ElementType() reflect.Type {
	return reflect.TypeOf((**Desktop)(nil)).Elem()
}

func (i *Desktop) ToDesktopOutput() DesktopOutput {
	return i.ToDesktopOutputWithContext(context.Background())
}

func (i *Desktop) ToDesktopOutputWithContext(ctx context.Context) DesktopOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DesktopOutput)
}

// DesktopArrayInput is an input type that accepts DesktopArray and DesktopArrayOutput values.
// You can construct a concrete instance of `DesktopArrayInput` via:
//
//	DesktopArray{ DesktopArgs{...} }
type DesktopArrayInput interface {
	pulumi.Input

	ToDesktopArrayOutput() DesktopArrayOutput
	ToDesktopArrayOutputWithContext(context.Context) DesktopArrayOutput
}

type DesktopArray []DesktopInput

func (DesktopArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Desktop)(nil)).Elem()
}

func (i DesktopArray) ToDesktopArrayOutput() DesktopArrayOutput {
	return i.ToDesktopArrayOutputWithContext(context.Background())
}

func (i DesktopArray) ToDesktopArrayOutputWithContext(ctx context.Context) DesktopArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DesktopArrayOutput)
}

// DesktopMapInput is an input type that accepts DesktopMap and DesktopMapOutput values.
// You can construct a concrete instance of `DesktopMapInput` via:
//
//	DesktopMap{ "key": DesktopArgs{...} }
type DesktopMapInput interface {
	pulumi.Input

	ToDesktopMapOutput() DesktopMapOutput
	ToDesktopMapOutputWithContext(context.Context) DesktopMapOutput
}

type DesktopMap map[string]DesktopInput

func (DesktopMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Desktop)(nil)).Elem()
}

func (i DesktopMap) ToDesktopMapOutput() DesktopMapOutput {
	return i.ToDesktopMapOutputWithContext(context.Background())
}

func (i DesktopMap) ToDesktopMapOutputWithContext(ctx context.Context) DesktopMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DesktopMapOutput)
}

type DesktopOutput struct{ *pulumi.OutputState }

func (DesktopOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Desktop)(nil)).Elem()
}

func (o DesktopOutput) ToDesktopOutput() DesktopOutput {
	return o
}

func (o DesktopOutput) ToDesktopOutputWithContext(ctx context.Context) DesktopOutput {
	return o
}

// Specifies the availability zone where the desktop is located.
// Changing this will create a new resource.
func (o DesktopOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Specifies the configuration of data volumes.
// The object structure is documented below.
func (o DesktopOutput) DataVolumes() DesktopDataVolumeArrayOutput {
	return o.ApplyT(func(v *Desktop) DesktopDataVolumeArrayOutput { return v.DataVolumes }).(DesktopDataVolumeArrayOutput)
}

// Specifies whether to delete user associated with this desktop after deleting it.
// The user can only be successfully deleted if the user has no other desktops.
func (o DesktopOutput) DeleteUser() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Desktop) pulumi.BoolPtrOutput { return v.DeleteUser }).(pulumi.BoolPtrOutput)
}

// Specifies whether to send emails to user mailbox during important
// operations.
// Defaults to **false**. Changing this will create a new resource.
func (o DesktopOutput) EmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Desktop) pulumi.BoolPtrOutput { return v.EmailNotification }).(pulumi.BoolPtrOutput)
}

// Specifies the flavor ID of desktop.
func (o DesktopOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.FlavorId }).(pulumi.StringOutput)
}

// Specifies the image ID to create the desktop.
// Changing this will create a new resource.
func (o DesktopOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// Specifies the image type. The valid values are as follows:
// + **market**: The market image.
// + **gold**: The public image.
// + **private**: The private image.
func (o DesktopOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.ImageType }).(pulumi.StringOutput)
}

// Specifies the desktop name.
// The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
// The name must start with a letter or digit and cannot end with a hyphen.
// Changing this will create a new resource.
func (o DesktopOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the NIC information corresponding to the desktop.
// The object structure is documented below. Changing this will create a new resource.
func (o DesktopOutput) Nics() DesktopNicArrayOutput {
	return o.ApplyT(func(v *Desktop) DesktopNicArrayOutput { return v.Nics }).(DesktopNicArrayOutput)
}

// The region in which to create the Workspace desktop resource.
// If omitted, the provider-level region will be used. Changing this will create a new resource.
func (o DesktopOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the configuration of system volume.
// The object structure is documented below.
func (o DesktopOutput) RootVolume() DesktopRootVolumeOutput {
	return o.ApplyT(func(v *Desktop) DesktopRootVolumeOutput { return v.RootVolume }).(DesktopRootVolumeOutput)
}

// Specifies the ID list of security groups.
// In addition to the custom security group, it must also contain a security group called **WorkspaceUserSecurityGroup**.
// Changing this will create a new resource.
func (o DesktopOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Specifies the key/value pairs of the desktop.
// Changing this will create a new resource.
func (o DesktopOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the user email.
// Some operations on the desktop (such as creation, deletion) will notify the user by sending an email.
// Changing this will create a new resource.
func (o DesktopOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.UserEmail }).(pulumi.StringOutput)
}

// Specifies the user group to which the desktop belongs.
// The valid values are as follows:
// + **sudo**: Linux administrator group.
// + **default**: Linux default user group.
// + **administrators**: Windows administrator group.
// + **users**: Windows standard user group.
func (o DesktopOutput) UserGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.UserGroup }).(pulumi.StringOutput)
}

// Specifies the user name to which the desktop belongs.
// The name can contain `1` to `20` characters, only letters, digits, hyphens (-) and underscores (_) are allowed.
// The name must start with a letter. Changing this will create a new resource.
func (o DesktopOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// Specifies the VPC ID to which the desktop belongs.
// Changing this will create a new resource.
func (o DesktopOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Desktop) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type DesktopArrayOutput struct{ *pulumi.OutputState }

func (DesktopArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Desktop)(nil)).Elem()
}

func (o DesktopArrayOutput) ToDesktopArrayOutput() DesktopArrayOutput {
	return o
}

func (o DesktopArrayOutput) ToDesktopArrayOutputWithContext(ctx context.Context) DesktopArrayOutput {
	return o
}

func (o DesktopArrayOutput) Index(i pulumi.IntInput) DesktopOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Desktop {
		return vs[0].([]*Desktop)[vs[1].(int)]
	}).(DesktopOutput)
}

type DesktopMapOutput struct{ *pulumi.OutputState }

func (DesktopMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Desktop)(nil)).Elem()
}

func (o DesktopMapOutput) ToDesktopMapOutput() DesktopMapOutput {
	return o
}

func (o DesktopMapOutput) ToDesktopMapOutputWithContext(ctx context.Context) DesktopMapOutput {
	return o
}

func (o DesktopMapOutput) MapIndex(k pulumi.StringInput) DesktopOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Desktop {
		return vs[0].(map[string]*Desktop)[vs[1].(string)]
	}).(DesktopOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DesktopInput)(nil)).Elem(), &Desktop{})
	pulumi.RegisterInputType(reflect.TypeOf((*DesktopArrayInput)(nil)).Elem(), DesktopArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DesktopMapInput)(nil)).Elem(), DesktopMap{})
	pulumi.RegisterOutputType(DesktopOutput{})
	pulumi.RegisterOutputType(DesktopArrayOutput{})
	pulumi.RegisterOutputType(DesktopMapOutput{})
}
