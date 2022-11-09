// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sharedapig

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an API gateway API resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/SharedApig"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/SharedApig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tfApigwGroup, err := SharedApig.NewGroup(ctx, "tfApigwGroup", &SharedApig.GroupArgs{
//				Description: pulumi.String("your descpiption"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = SharedApig.NewApi(ctx, "tfApigwApi", &SharedApig.ApiArgs{
//				GroupId:     tfApigwGroup.ID(),
//				Description: pulumi.String("your descpiption"),
//				Tags: pulumi.StringArray{
//					pulumi.String("tag1"),
//					pulumi.String("tag2"),
//				},
//				Visibility:             pulumi.Int(2),
//				AuthType:               pulumi.String("IAM"),
//				BackendType:            pulumi.String("HTTP"),
//				RequestProtocol:        pulumi.String("HTTPS"),
//				RequestMethod:          pulumi.String("GET"),
//				RequestUri:             pulumi.String("/test/path1"),
//				ExampleSuccessResponse: pulumi.String("example response"),
//				HttpBackend: &sharedapig.ApiHttpBackendArgs{
//					Protocol:  pulumi.String("HTTPS"),
//					Method:    pulumi.String("GET"),
//					Uri:       pulumi.String("/web/openapi"),
//					UrlDomain: pulumi.String("myhuaweicloud.com"),
//					Timeout:   pulumi.Int(10000),
//				},
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
// API can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:SharedApig/api:Api api "774438a28a574ac8a496325d1bf51807"
//
// ```
type Api struct {
	pulumi.CustomResourceState

	// Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
	// NONE'.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// the backend parameter list (documented below).
	BackendParameters ApiBackendParameterArrayOutput `pulumi:"backendParameters"`
	// Specifies the service backend type. The value can be:
	// + 'HTTP': the web service backend
	// + 'FUNCTION': the FunctionGraph service backend
	// + 'MOCK': the Mock service backend
	BackendType pulumi.StringOutput `pulumi:"backendType"`
	// Specifies whether CORS is supported or not.
	Cors pulumi.BoolPtrOutput `pulumi:"cors"`
	// Specifies the description of the API. The description cannot exceed 255 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the example response for a failed request The length cannot
	// exceed 20,480 characters.
	ExampleFailureResponse pulumi.StringPtrOutput `pulumi:"exampleFailureResponse"`
	// Specifies the example response for a successful request. The length
	// cannot exceed 20,480 characters.
	ExampleSuccessResponse pulumi.StringOutput `pulumi:"exampleSuccessResponse"`
	// Specifies the configuration when backendType selected 'FUNCTION' (documented
	// below).
	FunctionBackend ApiFunctionBackendPtrOutput `pulumi:"functionBackend"`
	// Specifies the ID of the API group. Changing this creates a new resource.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the API group to which the API belongs.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Specifies the configuration when backendType selected 'HTTP' (documented below).
	HttpBackend ApiHttpBackendPtrOutput `pulumi:"httpBackend"`
	// Specifies the configuration when backendType selected 'MOCK' (documented below).
	MockBackend ApiMockBackendPtrOutput `pulumi:"mockBackend"`
	// Specifies the name of the API. An API name consists of 3–64 characters, starting with a
	// letter. Only letters, digits, and underscores (_) are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the API resource. If omitted, the provider-level
	// region will be used. Changing this creates a new API resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the request method, including 'GET','POST','PUT' and etc..
	RequestMethod pulumi.StringOutput `pulumi:"requestMethod"`
	// the request parameter list (documented below).
	RequestParameters ApiRequestParameterArrayOutput `pulumi:"requestParameters"`
	// Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
	// which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
	RequestProtocol pulumi.StringPtrOutput `pulumi:"requestProtocol"`
	// Specifies the request path of the API. The value must comply with URI
	// specifications.
	RequestUri pulumi.StringOutput `pulumi:"requestUri"`
	// the tags of API in format of string list.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Specifies the version of the API. A maximum of 16 characters are allowed.
	Version pulumi.StringOutput `pulumi:"version"`
	// Specifies whether the API is available to the public. The value can be 1 (public) and
	// 2 (private). Defaults to 2.
	Visibility pulumi.IntPtrOutput `pulumi:"visibility"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthType == nil {
		return nil, errors.New("invalid value for required argument 'AuthType'")
	}
	if args.BackendType == nil {
		return nil, errors.New("invalid value for required argument 'BackendType'")
	}
	if args.ExampleSuccessResponse == nil {
		return nil, errors.New("invalid value for required argument 'ExampleSuccessResponse'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.RequestMethod == nil {
		return nil, errors.New("invalid value for required argument 'RequestMethod'")
	}
	if args.RequestUri == nil {
		return nil, errors.New("invalid value for required argument 'RequestUri'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Api
	err := ctx.RegisterResource("huaweicloud:SharedApig/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("huaweicloud:SharedApig/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
	// NONE'.
	AuthType *string `pulumi:"authType"`
	// the backend parameter list (documented below).
	BackendParameters []ApiBackendParameter `pulumi:"backendParameters"`
	// Specifies the service backend type. The value can be:
	// + 'HTTP': the web service backend
	// + 'FUNCTION': the FunctionGraph service backend
	// + 'MOCK': the Mock service backend
	BackendType *string `pulumi:"backendType"`
	// Specifies whether CORS is supported or not.
	Cors *bool `pulumi:"cors"`
	// Specifies the description of the API. The description cannot exceed 255 characters.
	Description *string `pulumi:"description"`
	// Specifies the example response for a failed request The length cannot
	// exceed 20,480 characters.
	ExampleFailureResponse *string `pulumi:"exampleFailureResponse"`
	// Specifies the example response for a successful request. The length
	// cannot exceed 20,480 characters.
	ExampleSuccessResponse *string `pulumi:"exampleSuccessResponse"`
	// Specifies the configuration when backendType selected 'FUNCTION' (documented
	// below).
	FunctionBackend *ApiFunctionBackend `pulumi:"functionBackend"`
	// Specifies the ID of the API group. Changing this creates a new resource.
	GroupId *string `pulumi:"groupId"`
	// The name of the API group to which the API belongs.
	GroupName *string `pulumi:"groupName"`
	// Specifies the configuration when backendType selected 'HTTP' (documented below).
	HttpBackend *ApiHttpBackend `pulumi:"httpBackend"`
	// Specifies the configuration when backendType selected 'MOCK' (documented below).
	MockBackend *ApiMockBackend `pulumi:"mockBackend"`
	// Specifies the name of the API. An API name consists of 3–64 characters, starting with a
	// letter. Only letters, digits, and underscores (_) are allowed.
	Name *string `pulumi:"name"`
	// The region in which to create the API resource. If omitted, the provider-level
	// region will be used. Changing this creates a new API resource.
	Region *string `pulumi:"region"`
	// Specifies the request method, including 'GET','POST','PUT' and etc..
	RequestMethod *string `pulumi:"requestMethod"`
	// the request parameter list (documented below).
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
	// which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
	RequestProtocol *string `pulumi:"requestProtocol"`
	// Specifies the request path of the API. The value must comply with URI
	// specifications.
	RequestUri *string `pulumi:"requestUri"`
	// the tags of API in format of string list.
	Tags []string `pulumi:"tags"`
	// Specifies the version of the API. A maximum of 16 characters are allowed.
	Version *string `pulumi:"version"`
	// Specifies whether the API is available to the public. The value can be 1 (public) and
	// 2 (private). Defaults to 2.
	Visibility *int `pulumi:"visibility"`
}

type ApiState struct {
	// Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
	// NONE'.
	AuthType pulumi.StringPtrInput
	// the backend parameter list (documented below).
	BackendParameters ApiBackendParameterArrayInput
	// Specifies the service backend type. The value can be:
	// + 'HTTP': the web service backend
	// + 'FUNCTION': the FunctionGraph service backend
	// + 'MOCK': the Mock service backend
	BackendType pulumi.StringPtrInput
	// Specifies whether CORS is supported or not.
	Cors pulumi.BoolPtrInput
	// Specifies the description of the API. The description cannot exceed 255 characters.
	Description pulumi.StringPtrInput
	// Specifies the example response for a failed request The length cannot
	// exceed 20,480 characters.
	ExampleFailureResponse pulumi.StringPtrInput
	// Specifies the example response for a successful request. The length
	// cannot exceed 20,480 characters.
	ExampleSuccessResponse pulumi.StringPtrInput
	// Specifies the configuration when backendType selected 'FUNCTION' (documented
	// below).
	FunctionBackend ApiFunctionBackendPtrInput
	// Specifies the ID of the API group. Changing this creates a new resource.
	GroupId pulumi.StringPtrInput
	// The name of the API group to which the API belongs.
	GroupName pulumi.StringPtrInput
	// Specifies the configuration when backendType selected 'HTTP' (documented below).
	HttpBackend ApiHttpBackendPtrInput
	// Specifies the configuration when backendType selected 'MOCK' (documented below).
	MockBackend ApiMockBackendPtrInput
	// Specifies the name of the API. An API name consists of 3–64 characters, starting with a
	// letter. Only letters, digits, and underscores (_) are allowed.
	Name pulumi.StringPtrInput
	// The region in which to create the API resource. If omitted, the provider-level
	// region will be used. Changing this creates a new API resource.
	Region pulumi.StringPtrInput
	// Specifies the request method, including 'GET','POST','PUT' and etc..
	RequestMethod pulumi.StringPtrInput
	// the request parameter list (documented below).
	RequestParameters ApiRequestParameterArrayInput
	// Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
	// which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
	RequestProtocol pulumi.StringPtrInput
	// Specifies the request path of the API. The value must comply with URI
	// specifications.
	RequestUri pulumi.StringPtrInput
	// the tags of API in format of string list.
	Tags pulumi.StringArrayInput
	// Specifies the version of the API. A maximum of 16 characters are allowed.
	Version pulumi.StringPtrInput
	// Specifies whether the API is available to the public. The value can be 1 (public) and
	// 2 (private). Defaults to 2.
	Visibility pulumi.IntPtrInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
	// NONE'.
	AuthType string `pulumi:"authType"`
	// the backend parameter list (documented below).
	BackendParameters []ApiBackendParameter `pulumi:"backendParameters"`
	// Specifies the service backend type. The value can be:
	// + 'HTTP': the web service backend
	// + 'FUNCTION': the FunctionGraph service backend
	// + 'MOCK': the Mock service backend
	BackendType string `pulumi:"backendType"`
	// Specifies whether CORS is supported or not.
	Cors *bool `pulumi:"cors"`
	// Specifies the description of the API. The description cannot exceed 255 characters.
	Description *string `pulumi:"description"`
	// Specifies the example response for a failed request The length cannot
	// exceed 20,480 characters.
	ExampleFailureResponse *string `pulumi:"exampleFailureResponse"`
	// Specifies the example response for a successful request. The length
	// cannot exceed 20,480 characters.
	ExampleSuccessResponse string `pulumi:"exampleSuccessResponse"`
	// Specifies the configuration when backendType selected 'FUNCTION' (documented
	// below).
	FunctionBackend *ApiFunctionBackend `pulumi:"functionBackend"`
	// Specifies the ID of the API group. Changing this creates a new resource.
	GroupId string `pulumi:"groupId"`
	// Specifies the configuration when backendType selected 'HTTP' (documented below).
	HttpBackend *ApiHttpBackend `pulumi:"httpBackend"`
	// Specifies the configuration when backendType selected 'MOCK' (documented below).
	MockBackend *ApiMockBackend `pulumi:"mockBackend"`
	// Specifies the name of the API. An API name consists of 3–64 characters, starting with a
	// letter. Only letters, digits, and underscores (_) are allowed.
	Name *string `pulumi:"name"`
	// The region in which to create the API resource. If omitted, the provider-level
	// region will be used. Changing this creates a new API resource.
	Region *string `pulumi:"region"`
	// Specifies the request method, including 'GET','POST','PUT' and etc..
	RequestMethod string `pulumi:"requestMethod"`
	// the request parameter list (documented below).
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
	// which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
	RequestProtocol *string `pulumi:"requestProtocol"`
	// Specifies the request path of the API. The value must comply with URI
	// specifications.
	RequestUri string `pulumi:"requestUri"`
	// the tags of API in format of string list.
	Tags []string `pulumi:"tags"`
	// Specifies the version of the API. A maximum of 16 characters are allowed.
	Version *string `pulumi:"version"`
	// Specifies whether the API is available to the public. The value can be 1 (public) and
	// 2 (private). Defaults to 2.
	Visibility *int `pulumi:"visibility"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
	// NONE'.
	AuthType pulumi.StringInput
	// the backend parameter list (documented below).
	BackendParameters ApiBackendParameterArrayInput
	// Specifies the service backend type. The value can be:
	// + 'HTTP': the web service backend
	// + 'FUNCTION': the FunctionGraph service backend
	// + 'MOCK': the Mock service backend
	BackendType pulumi.StringInput
	// Specifies whether CORS is supported or not.
	Cors pulumi.BoolPtrInput
	// Specifies the description of the API. The description cannot exceed 255 characters.
	Description pulumi.StringPtrInput
	// Specifies the example response for a failed request The length cannot
	// exceed 20,480 characters.
	ExampleFailureResponse pulumi.StringPtrInput
	// Specifies the example response for a successful request. The length
	// cannot exceed 20,480 characters.
	ExampleSuccessResponse pulumi.StringInput
	// Specifies the configuration when backendType selected 'FUNCTION' (documented
	// below).
	FunctionBackend ApiFunctionBackendPtrInput
	// Specifies the ID of the API group. Changing this creates a new resource.
	GroupId pulumi.StringInput
	// Specifies the configuration when backendType selected 'HTTP' (documented below).
	HttpBackend ApiHttpBackendPtrInput
	// Specifies the configuration when backendType selected 'MOCK' (documented below).
	MockBackend ApiMockBackendPtrInput
	// Specifies the name of the API. An API name consists of 3–64 characters, starting with a
	// letter. Only letters, digits, and underscores (_) are allowed.
	Name pulumi.StringPtrInput
	// The region in which to create the API resource. If omitted, the provider-level
	// region will be used. Changing this creates a new API resource.
	Region pulumi.StringPtrInput
	// Specifies the request method, including 'GET','POST','PUT' and etc..
	RequestMethod pulumi.StringInput
	// the request parameter list (documented below).
	RequestParameters ApiRequestParameterArrayInput
	// Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
	// which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
	RequestProtocol pulumi.StringPtrInput
	// Specifies the request path of the API. The value must comply with URI
	// specifications.
	RequestUri pulumi.StringInput
	// the tags of API in format of string list.
	Tags pulumi.StringArrayInput
	// Specifies the version of the API. A maximum of 16 characters are allowed.
	Version pulumi.StringPtrInput
	// Specifies whether the API is available to the public. The value can be 1 (public) and
	// 2 (private). Defaults to 2.
	Visibility pulumi.IntPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

// ApiArrayInput is an input type that accepts ApiArray and ApiArrayOutput values.
// You can construct a concrete instance of `ApiArrayInput` via:
//
//	ApiArray{ ApiArgs{...} }
type ApiArrayInput interface {
	pulumi.Input

	ToApiArrayOutput() ApiArrayOutput
	ToApiArrayOutputWithContext(context.Context) ApiArrayOutput
}

type ApiArray []ApiInput

func (ApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (i ApiArray) ToApiArrayOutput() ApiArrayOutput {
	return i.ToApiArrayOutputWithContext(context.Background())
}

func (i ApiArray) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiArrayOutput)
}

// ApiMapInput is an input type that accepts ApiMap and ApiMapOutput values.
// You can construct a concrete instance of `ApiMapInput` via:
//
//	ApiMap{ "key": ApiArgs{...} }
type ApiMapInput interface {
	pulumi.Input

	ToApiMapOutput() ApiMapOutput
	ToApiMapOutputWithContext(context.Context) ApiMapOutput
}

type ApiMap map[string]ApiInput

func (ApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (i ApiMap) ToApiMapOutput() ApiMapOutput {
	return i.ToApiMapOutputWithContext(context.Background())
}

func (i ApiMap) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMapOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

// Specifies the security authentication mode. The value can be 'APP', 'IAM', and '
// NONE'.
func (o ApiOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// the backend parameter list (documented below).
func (o ApiOutput) BackendParameters() ApiBackendParameterArrayOutput {
	return o.ApplyT(func(v *Api) ApiBackendParameterArrayOutput { return v.BackendParameters }).(ApiBackendParameterArrayOutput)
}

// Specifies the service backend type. The value can be:
// + 'HTTP': the web service backend
// + 'FUNCTION': the FunctionGraph service backend
// + 'MOCK': the Mock service backend
func (o ApiOutput) BackendType() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.BackendType }).(pulumi.StringOutput)
}

// Specifies whether CORS is supported or not.
func (o ApiOutput) Cors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.Cors }).(pulumi.BoolPtrOutput)
}

// Specifies the description of the API. The description cannot exceed 255 characters.
func (o ApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the example response for a failed request The length cannot
// exceed 20,480 characters.
func (o ApiOutput) ExampleFailureResponse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ExampleFailureResponse }).(pulumi.StringPtrOutput)
}

// Specifies the example response for a successful request. The length
// cannot exceed 20,480 characters.
func (o ApiOutput) ExampleSuccessResponse() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ExampleSuccessResponse }).(pulumi.StringOutput)
}

// Specifies the configuration when backendType selected 'FUNCTION' (documented
// below).
func (o ApiOutput) FunctionBackend() ApiFunctionBackendPtrOutput {
	return o.ApplyT(func(v *Api) ApiFunctionBackendPtrOutput { return v.FunctionBackend }).(ApiFunctionBackendPtrOutput)
}

// Specifies the ID of the API group. Changing this creates a new resource.
func (o ApiOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the API group to which the API belongs.
func (o ApiOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// Specifies the configuration when backendType selected 'HTTP' (documented below).
func (o ApiOutput) HttpBackend() ApiHttpBackendPtrOutput {
	return o.ApplyT(func(v *Api) ApiHttpBackendPtrOutput { return v.HttpBackend }).(ApiHttpBackendPtrOutput)
}

// Specifies the configuration when backendType selected 'MOCK' (documented below).
func (o ApiOutput) MockBackend() ApiMockBackendPtrOutput {
	return o.ApplyT(func(v *Api) ApiMockBackendPtrOutput { return v.MockBackend }).(ApiMockBackendPtrOutput)
}

// Specifies the name of the API. An API name consists of 3–64 characters, starting with a
// letter. Only letters, digits, and underscores (_) are allowed.
func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the API resource. If omitted, the provider-level
// region will be used. Changing this creates a new API resource.
func (o ApiOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the request method, including 'GET','POST','PUT' and etc..
func (o ApiOutput) RequestMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.RequestMethod }).(pulumi.StringOutput)
}

// the request parameter list (documented below).
func (o ApiOutput) RequestParameters() ApiRequestParameterArrayOutput {
	return o.ApplyT(func(v *Api) ApiRequestParameterArrayOutput { return v.RequestParameters }).(ApiRequestParameterArrayOutput)
}

// Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH'
// which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.
func (o ApiOutput) RequestProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.RequestProtocol }).(pulumi.StringPtrOutput)
}

// Specifies the request path of the API. The value must comply with URI
// specifications.
func (o ApiOutput) RequestUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.RequestUri }).(pulumi.StringOutput)
}

// the tags of API in format of string list.
func (o ApiOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Api) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Specifies the version of the API. A maximum of 16 characters are allowed.
func (o ApiOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// Specifies whether the API is available to the public. The value can be 1 (public) and
// 2 (private). Defaults to 2.
func (o ApiOutput) Visibility() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.IntPtrOutput { return v.Visibility }).(pulumi.IntPtrOutput)
}

type ApiArrayOutput struct{ *pulumi.OutputState }

func (ApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (o ApiArrayOutput) ToApiArrayOutput() ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) Index(i pulumi.IntInput) ApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Api {
		return vs[0].([]*Api)[vs[1].(int)]
	}).(ApiOutput)
}

type ApiMapOutput struct{ *pulumi.OutputState }

func (ApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (o ApiMapOutput) ToApiMapOutput() ApiMapOutput {
	return o
}

func (o ApiMapOutput) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return o
}

func (o ApiMapOutput) MapIndex(k pulumi.StringInput) ApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Api {
		return vs[0].(map[string]*Api)[vs[1].(string)]
	}).(ApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiArrayInput)(nil)).Elem(), ApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiMapInput)(nil)).Elem(), ApiMap{})
	pulumi.RegisterOutputType(ApiOutput{})
	pulumi.RegisterOutputType(ApiArrayOutput{})
	pulumi.RegisterOutputType(ApiMapOutput{})
}
