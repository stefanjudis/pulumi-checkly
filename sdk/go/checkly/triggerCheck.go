// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # TriggerCheck
//
// `TriggerCheck` allows users to manage Checkly trigger checks. Add a `TriggerCheck` resource to your resource file.
//
// ## Example Usage
//
// Trigger check example
//
// ```go
// package main
//
// import (
// 	"github.com/checkly/pulumi-checkly/sdk/go/checkly"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := checkly.NewTriggerCheck(ctx, "test-trigger-check", &checkly.TriggerCheckArgs{
// 			CheckId: pulumi.String("c1ff95c5-d7f6-4a90-9ce2-1e605f117592"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("test-trigger-check-url", test_trigger_check.Url)
// 		return nil
// 	})
// }
// ```
type TriggerCheck struct {
	pulumi.CustomResourceState

	// The id of the check that you want to attach the trigger to.
	CheckId pulumi.StringOutput `pulumi:"checkId"`
	Token   pulumi.StringOutput `pulumi:"token"`
	Url     pulumi.StringOutput `pulumi:"url"`
}

// NewTriggerCheck registers a new resource with the given unique name, arguments, and options.
func NewTriggerCheck(ctx *pulumi.Context,
	name string, args *TriggerCheckArgs, opts ...pulumi.ResourceOption) (*TriggerCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CheckId == nil {
		return nil, errors.New("invalid value for required argument 'CheckId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TriggerCheck
	err := ctx.RegisterResource("checkly:index/triggerCheck:TriggerCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerCheck gets an existing TriggerCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerCheckState, opts ...pulumi.ResourceOption) (*TriggerCheck, error) {
	var resource TriggerCheck
	err := ctx.ReadResource("checkly:index/triggerCheck:TriggerCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerCheck resources.
type triggerCheckState struct {
	// The id of the check that you want to attach the trigger to.
	CheckId *string `pulumi:"checkId"`
	Token   *string `pulumi:"token"`
	Url     *string `pulumi:"url"`
}

type TriggerCheckState struct {
	// The id of the check that you want to attach the trigger to.
	CheckId pulumi.StringPtrInput
	Token   pulumi.StringPtrInput
	Url     pulumi.StringPtrInput
}

func (TriggerCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerCheckState)(nil)).Elem()
}

type triggerCheckArgs struct {
	// The id of the check that you want to attach the trigger to.
	CheckId string  `pulumi:"checkId"`
	Token   *string `pulumi:"token"`
	Url     *string `pulumi:"url"`
}

// The set of arguments for constructing a TriggerCheck resource.
type TriggerCheckArgs struct {
	// The id of the check that you want to attach the trigger to.
	CheckId pulumi.StringInput
	Token   pulumi.StringPtrInput
	Url     pulumi.StringPtrInput
}

func (TriggerCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerCheckArgs)(nil)).Elem()
}

type TriggerCheckInput interface {
	pulumi.Input

	ToTriggerCheckOutput() TriggerCheckOutput
	ToTriggerCheckOutputWithContext(ctx context.Context) TriggerCheckOutput
}

func (*TriggerCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerCheck)(nil)).Elem()
}

func (i *TriggerCheck) ToTriggerCheckOutput() TriggerCheckOutput {
	return i.ToTriggerCheckOutputWithContext(context.Background())
}

func (i *TriggerCheck) ToTriggerCheckOutputWithContext(ctx context.Context) TriggerCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCheckOutput)
}

// TriggerCheckArrayInput is an input type that accepts TriggerCheckArray and TriggerCheckArrayOutput values.
// You can construct a concrete instance of `TriggerCheckArrayInput` via:
//
//          TriggerCheckArray{ TriggerCheckArgs{...} }
type TriggerCheckArrayInput interface {
	pulumi.Input

	ToTriggerCheckArrayOutput() TriggerCheckArrayOutput
	ToTriggerCheckArrayOutputWithContext(context.Context) TriggerCheckArrayOutput
}

type TriggerCheckArray []TriggerCheckInput

func (TriggerCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TriggerCheck)(nil)).Elem()
}

func (i TriggerCheckArray) ToTriggerCheckArrayOutput() TriggerCheckArrayOutput {
	return i.ToTriggerCheckArrayOutputWithContext(context.Background())
}

func (i TriggerCheckArray) ToTriggerCheckArrayOutputWithContext(ctx context.Context) TriggerCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCheckArrayOutput)
}

// TriggerCheckMapInput is an input type that accepts TriggerCheckMap and TriggerCheckMapOutput values.
// You can construct a concrete instance of `TriggerCheckMapInput` via:
//
//          TriggerCheckMap{ "key": TriggerCheckArgs{...} }
type TriggerCheckMapInput interface {
	pulumi.Input

	ToTriggerCheckMapOutput() TriggerCheckMapOutput
	ToTriggerCheckMapOutputWithContext(context.Context) TriggerCheckMapOutput
}

type TriggerCheckMap map[string]TriggerCheckInput

func (TriggerCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TriggerCheck)(nil)).Elem()
}

func (i TriggerCheckMap) ToTriggerCheckMapOutput() TriggerCheckMapOutput {
	return i.ToTriggerCheckMapOutputWithContext(context.Background())
}

func (i TriggerCheckMap) ToTriggerCheckMapOutputWithContext(ctx context.Context) TriggerCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCheckMapOutput)
}

type TriggerCheckOutput struct{ *pulumi.OutputState }

func (TriggerCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerCheck)(nil)).Elem()
}

func (o TriggerCheckOutput) ToTriggerCheckOutput() TriggerCheckOutput {
	return o
}

func (o TriggerCheckOutput) ToTriggerCheckOutputWithContext(ctx context.Context) TriggerCheckOutput {
	return o
}

type TriggerCheckArrayOutput struct{ *pulumi.OutputState }

func (TriggerCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TriggerCheck)(nil)).Elem()
}

func (o TriggerCheckArrayOutput) ToTriggerCheckArrayOutput() TriggerCheckArrayOutput {
	return o
}

func (o TriggerCheckArrayOutput) ToTriggerCheckArrayOutputWithContext(ctx context.Context) TriggerCheckArrayOutput {
	return o
}

func (o TriggerCheckArrayOutput) Index(i pulumi.IntInput) TriggerCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TriggerCheck {
		return vs[0].([]*TriggerCheck)[vs[1].(int)]
	}).(TriggerCheckOutput)
}

type TriggerCheckMapOutput struct{ *pulumi.OutputState }

func (TriggerCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TriggerCheck)(nil)).Elem()
}

func (o TriggerCheckMapOutput) ToTriggerCheckMapOutput() TriggerCheckMapOutput {
	return o
}

func (o TriggerCheckMapOutput) ToTriggerCheckMapOutputWithContext(ctx context.Context) TriggerCheckMapOutput {
	return o
}

func (o TriggerCheckMapOutput) MapIndex(k pulumi.StringInput) TriggerCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TriggerCheck {
		return vs[0].(map[string]*TriggerCheck)[vs[1].(string)]
	}).(TriggerCheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerCheckInput)(nil)).Elem(), &TriggerCheck{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerCheckArrayInput)(nil)).Elem(), TriggerCheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerCheckMapInput)(nil)).Elem(), TriggerCheckMap{})
	pulumi.RegisterOutputType(TriggerCheckOutput{})
	pulumi.RegisterOutputType(TriggerCheckArrayOutput{})
	pulumi.RegisterOutputType(TriggerCheckMapOutput{})
}
