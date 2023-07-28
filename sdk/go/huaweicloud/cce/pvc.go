// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cce

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a CCE Persistent Volume Claim resource within HuaweiCloud.
//
// ## Example Usage
// ### Create PVC with EVS
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterId := cfg.RequireObject("clusterId")
//			namespace := cfg.RequireObject("namespace")
//			pvcName := cfg.RequireObject("pvcName")
//			_, err := Cce.NewPvc(ctx, "test", &Cce.PvcArgs{
//				ClusterId: pulumi.Any(clusterId),
//				Namespace: pulumi.Any(namespace),
//				Annotations: pulumi.StringMap{
//					"everest.io/disk-volume-type": pulumi.String("SSD"),
//				},
//				StorageClassName: pulumi.String("csi-disk"),
//				AccessModes: pulumi.StringArray{
//					pulumi.String("ReadWriteOnce"),
//				},
//				Storage: pulumi.String("10Gi"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create PVC with OBS
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterId := cfg.RequireObject("clusterId")
//			namespace := cfg.RequireObject("namespace")
//			pvcName := cfg.RequireObject("pvcName")
//			_, err := Cce.NewPvc(ctx, "test", &Cce.PvcArgs{
//				ClusterId: pulumi.Any(clusterId),
//				Namespace: pulumi.Any(namespace),
//				Annotations: pulumi.StringMap{
//					"everest.io/obs-volume-type": pulumi.String("STANDARD"),
//					"csi.storage.k8s.io/fstype":  pulumi.String("obsfs"),
//				},
//				StorageClassName: pulumi.String("csi-obs"),
//				AccessModes: pulumi.StringArray{
//					pulumi.String("ReadWriteMany"),
//				},
//				Storage: pulumi.String("1Gi"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create PVC with SFS
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterId := cfg.RequireObject("clusterId")
//			namespace := cfg.RequireObject("namespace")
//			pvcName := cfg.RequireObject("pvcName")
//			_, err := Cce.NewPvc(ctx, "test", &Cce.PvcArgs{
//				ClusterId:        pulumi.Any(clusterId),
//				Namespace:        pulumi.Any(namespace),
//				StorageClassName: pulumi.String("csi-nas"),
//				AccessModes: pulumi.StringArray{
//					pulumi.String("ReadWriteMany"),
//				},
//				Storage: pulumi.String("10Gi"),
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
// CCE PVC can be imported using the cluster ID, namespace and ID separated by slashes, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Cce/pvc:Pvc test 5c20fdad-7288-11eb-b817-0255ac10158b/default/fa540f3b-12d9-40e5-8268-04bcfed95a46
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`annotations`. It is generally recommended running `terraform plan` after importing a PVC. You can then decide if changes should be applied to the PVC, or the resource definition should be updated to align with the PVC. Also you can ignore changes as below. resource "huaweicloud_cce_pvc" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	annotations,
//
//	]
//
//	} }
type Pvc struct {
	pulumi.CustomResourceState

	// Specifies the desired access modes the volume should have.
	// The valid values are as follows:
	// + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
	// + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
	// + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
	AccessModes pulumi.StringArrayOutput `pulumi:"accessModes"`
	// Specifies the unstructured key value map for external parameters.
	// Changing this will create a new PVC resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Specifies the cluster ID to which the CCE PVC belongs.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The server time when PVC was created.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new PVC resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Specifies the unique name of the PVC resource. This parameter can contain a
	// maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
	// lowercase letters and digits. Changing this will create a new PVC resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the namespace to logically divide your containers into different
	// group. Changing this will create a new PVC resource.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// Specifies the region in which to create the PVC resource.
	// If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The current phase of the PVC.
	// + **Pending**: Not yet bound.
	// + **Bound**: Already bound.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the minimum amount of storage resources required.
	// Changing this creates a new PVC resource.
	Storage pulumi.StringOutput `pulumi:"storage"`
	// Specifies the type of the storage bound to the CCE PVC.
	// The valid values are as follows:
	// + **csi-disk**: EVS.
	// + **csi-obs**: OBS.
	// + **csi-nas**: SFS.
	StorageClassName pulumi.StringOutput `pulumi:"storageClassName"`
}

// NewPvc registers a new resource with the given unique name, arguments, and options.
func NewPvc(ctx *pulumi.Context,
	name string, args *PvcArgs, opts ...pulumi.ResourceOption) (*Pvc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessModes == nil {
		return nil, errors.New("invalid value for required argument 'AccessModes'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	if args.StorageClassName == nil {
		return nil, errors.New("invalid value for required argument 'StorageClassName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Pvc
	err := ctx.RegisterResource("huaweicloud:Cce/pvc:Pvc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPvc gets an existing Pvc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPvc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PvcState, opts ...pulumi.ResourceOption) (*Pvc, error) {
	var resource Pvc
	err := ctx.ReadResource("huaweicloud:Cce/pvc:Pvc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pvc resources.
type pvcState struct {
	// Specifies the desired access modes the volume should have.
	// The valid values are as follows:
	// + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
	// + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
	// + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
	AccessModes []string `pulumi:"accessModes"`
	// Specifies the unstructured key value map for external parameters.
	// Changing this will create a new PVC resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Specifies the cluster ID to which the CCE PVC belongs.
	ClusterId *string `pulumi:"clusterId"`
	// The server time when PVC was created.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new PVC resource.
	Labels map[string]string `pulumi:"labels"`
	// Specifies the unique name of the PVC resource. This parameter can contain a
	// maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
	// lowercase letters and digits. Changing this will create a new PVC resource.
	Name *string `pulumi:"name"`
	// Specifies the namespace to logically divide your containers into different
	// group. Changing this will create a new PVC resource.
	Namespace *string `pulumi:"namespace"`
	// Specifies the region in which to create the PVC resource.
	// If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
	Region *string `pulumi:"region"`
	// The current phase of the PVC.
	// + **Pending**: Not yet bound.
	// + **Bound**: Already bound.
	Status *string `pulumi:"status"`
	// Specifies the minimum amount of storage resources required.
	// Changing this creates a new PVC resource.
	Storage *string `pulumi:"storage"`
	// Specifies the type of the storage bound to the CCE PVC.
	// The valid values are as follows:
	// + **csi-disk**: EVS.
	// + **csi-obs**: OBS.
	// + **csi-nas**: SFS.
	StorageClassName *string `pulumi:"storageClassName"`
}

type PvcState struct {
	// Specifies the desired access modes the volume should have.
	// The valid values are as follows:
	// + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
	// + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
	// + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
	AccessModes pulumi.StringArrayInput
	// Specifies the unstructured key value map for external parameters.
	// Changing this will create a new PVC resource.
	Annotations pulumi.StringMapInput
	// Specifies the cluster ID to which the CCE PVC belongs.
	ClusterId pulumi.StringPtrInput
	// The server time when PVC was created.
	CreationTimestamp pulumi.StringPtrInput
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new PVC resource.
	Labels pulumi.StringMapInput
	// Specifies the unique name of the PVC resource. This parameter can contain a
	// maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
	// lowercase letters and digits. Changing this will create a new PVC resource.
	Name pulumi.StringPtrInput
	// Specifies the namespace to logically divide your containers into different
	// group. Changing this will create a new PVC resource.
	Namespace pulumi.StringPtrInput
	// Specifies the region in which to create the PVC resource.
	// If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
	Region pulumi.StringPtrInput
	// The current phase of the PVC.
	// + **Pending**: Not yet bound.
	// + **Bound**: Already bound.
	Status pulumi.StringPtrInput
	// Specifies the minimum amount of storage resources required.
	// Changing this creates a new PVC resource.
	Storage pulumi.StringPtrInput
	// Specifies the type of the storage bound to the CCE PVC.
	// The valid values are as follows:
	// + **csi-disk**: EVS.
	// + **csi-obs**: OBS.
	// + **csi-nas**: SFS.
	StorageClassName pulumi.StringPtrInput
}

func (PvcState) ElementType() reflect.Type {
	return reflect.TypeOf((*pvcState)(nil)).Elem()
}

type pvcArgs struct {
	// Specifies the desired access modes the volume should have.
	// The valid values are as follows:
	// + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
	// + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
	// + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
	AccessModes []string `pulumi:"accessModes"`
	// Specifies the unstructured key value map for external parameters.
	// Changing this will create a new PVC resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Specifies the cluster ID to which the CCE PVC belongs.
	ClusterId string `pulumi:"clusterId"`
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new PVC resource.
	Labels map[string]string `pulumi:"labels"`
	// Specifies the unique name of the PVC resource. This parameter can contain a
	// maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
	// lowercase letters and digits. Changing this will create a new PVC resource.
	Name *string `pulumi:"name"`
	// Specifies the namespace to logically divide your containers into different
	// group. Changing this will create a new PVC resource.
	Namespace string `pulumi:"namespace"`
	// Specifies the region in which to create the PVC resource.
	// If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
	Region *string `pulumi:"region"`
	// Specifies the minimum amount of storage resources required.
	// Changing this creates a new PVC resource.
	Storage string `pulumi:"storage"`
	// Specifies the type of the storage bound to the CCE PVC.
	// The valid values are as follows:
	// + **csi-disk**: EVS.
	// + **csi-obs**: OBS.
	// + **csi-nas**: SFS.
	StorageClassName string `pulumi:"storageClassName"`
}

// The set of arguments for constructing a Pvc resource.
type PvcArgs struct {
	// Specifies the desired access modes the volume should have.
	// The valid values are as follows:
	// + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
	// + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
	// + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
	AccessModes pulumi.StringArrayInput
	// Specifies the unstructured key value map for external parameters.
	// Changing this will create a new PVC resource.
	Annotations pulumi.StringMapInput
	// Specifies the cluster ID to which the CCE PVC belongs.
	ClusterId pulumi.StringInput
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new PVC resource.
	Labels pulumi.StringMapInput
	// Specifies the unique name of the PVC resource. This parameter can contain a
	// maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
	// lowercase letters and digits. Changing this will create a new PVC resource.
	Name pulumi.StringPtrInput
	// Specifies the namespace to logically divide your containers into different
	// group. Changing this will create a new PVC resource.
	Namespace pulumi.StringInput
	// Specifies the region in which to create the PVC resource.
	// If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
	Region pulumi.StringPtrInput
	// Specifies the minimum amount of storage resources required.
	// Changing this creates a new PVC resource.
	Storage pulumi.StringInput
	// Specifies the type of the storage bound to the CCE PVC.
	// The valid values are as follows:
	// + **csi-disk**: EVS.
	// + **csi-obs**: OBS.
	// + **csi-nas**: SFS.
	StorageClassName pulumi.StringInput
}

func (PvcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pvcArgs)(nil)).Elem()
}

type PvcInput interface {
	pulumi.Input

	ToPvcOutput() PvcOutput
	ToPvcOutputWithContext(ctx context.Context) PvcOutput
}

func (*Pvc) ElementType() reflect.Type {
	return reflect.TypeOf((**Pvc)(nil)).Elem()
}

func (i *Pvc) ToPvcOutput() PvcOutput {
	return i.ToPvcOutputWithContext(context.Background())
}

func (i *Pvc) ToPvcOutputWithContext(ctx context.Context) PvcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PvcOutput)
}

// PvcArrayInput is an input type that accepts PvcArray and PvcArrayOutput values.
// You can construct a concrete instance of `PvcArrayInput` via:
//
//	PvcArray{ PvcArgs{...} }
type PvcArrayInput interface {
	pulumi.Input

	ToPvcArrayOutput() PvcArrayOutput
	ToPvcArrayOutputWithContext(context.Context) PvcArrayOutput
}

type PvcArray []PvcInput

func (PvcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pvc)(nil)).Elem()
}

func (i PvcArray) ToPvcArrayOutput() PvcArrayOutput {
	return i.ToPvcArrayOutputWithContext(context.Background())
}

func (i PvcArray) ToPvcArrayOutputWithContext(ctx context.Context) PvcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PvcArrayOutput)
}

// PvcMapInput is an input type that accepts PvcMap and PvcMapOutput values.
// You can construct a concrete instance of `PvcMapInput` via:
//
//	PvcMap{ "key": PvcArgs{...} }
type PvcMapInput interface {
	pulumi.Input

	ToPvcMapOutput() PvcMapOutput
	ToPvcMapOutputWithContext(context.Context) PvcMapOutput
}

type PvcMap map[string]PvcInput

func (PvcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pvc)(nil)).Elem()
}

func (i PvcMap) ToPvcMapOutput() PvcMapOutput {
	return i.ToPvcMapOutputWithContext(context.Background())
}

func (i PvcMap) ToPvcMapOutputWithContext(ctx context.Context) PvcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PvcMapOutput)
}

type PvcOutput struct{ *pulumi.OutputState }

func (PvcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pvc)(nil)).Elem()
}

func (o PvcOutput) ToPvcOutput() PvcOutput {
	return o
}

func (o PvcOutput) ToPvcOutputWithContext(ctx context.Context) PvcOutput {
	return o
}

// Specifies the desired access modes the volume should have.
// The valid values are as follows:
// + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
// + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
// + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
func (o PvcOutput) AccessModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringArrayOutput { return v.AccessModes }).(pulumi.StringArrayOutput)
}

// Specifies the unstructured key value map for external parameters.
// Changing this will create a new PVC resource.
func (o PvcOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Specifies the cluster ID to which the CCE PVC belongs.
func (o PvcOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The server time when PVC was created.
func (o PvcOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Specifies the map of string keys and values for labels.
// Changing this will create a new PVC resource.
func (o PvcOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Specifies the unique name of the PVC resource. This parameter can contain a
// maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with
// lowercase letters and digits. Changing this will create a new PVC resource.
func (o PvcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the namespace to logically divide your containers into different
// group. Changing this will create a new PVC resource.
func (o PvcOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// Specifies the region in which to create the PVC resource.
// If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
func (o PvcOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current phase of the PVC.
// + **Pending**: Not yet bound.
// + **Bound**: Already bound.
func (o PvcOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the minimum amount of storage resources required.
// Changing this creates a new PVC resource.
func (o PvcOutput) Storage() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.Storage }).(pulumi.StringOutput)
}

// Specifies the type of the storage bound to the CCE PVC.
// The valid values are as follows:
// + **csi-disk**: EVS.
// + **csi-obs**: OBS.
// + **csi-nas**: SFS.
func (o PvcOutput) StorageClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pvc) pulumi.StringOutput { return v.StorageClassName }).(pulumi.StringOutput)
}

type PvcArrayOutput struct{ *pulumi.OutputState }

func (PvcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pvc)(nil)).Elem()
}

func (o PvcArrayOutput) ToPvcArrayOutput() PvcArrayOutput {
	return o
}

func (o PvcArrayOutput) ToPvcArrayOutputWithContext(ctx context.Context) PvcArrayOutput {
	return o
}

func (o PvcArrayOutput) Index(i pulumi.IntInput) PvcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pvc {
		return vs[0].([]*Pvc)[vs[1].(int)]
	}).(PvcOutput)
}

type PvcMapOutput struct{ *pulumi.OutputState }

func (PvcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pvc)(nil)).Elem()
}

func (o PvcMapOutput) ToPvcMapOutput() PvcMapOutput {
	return o
}

func (o PvcMapOutput) ToPvcMapOutputWithContext(ctx context.Context) PvcMapOutput {
	return o
}

func (o PvcMapOutput) MapIndex(k pulumi.StringInput) PvcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pvc {
		return vs[0].(map[string]*Pvc)[vs[1].(string)]
	}).(PvcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PvcInput)(nil)).Elem(), &Pvc{})
	pulumi.RegisterInputType(reflect.TypeOf((*PvcArrayInput)(nil)).Elem(), PvcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PvcMapInput)(nil)).Elem(), PvcMap{})
	pulumi.RegisterOutputType(PvcOutput{})
	pulumi.RegisterOutputType(PvcArrayOutput{})
	pulumi.RegisterOutputType(PvcMapOutput{})
}
