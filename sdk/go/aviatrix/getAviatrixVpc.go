// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAviatrixVpc(ctx *pulumi.Context, args *LookupAviatrixVpcArgs, opts ...pulumi.InvokeOption) (*LookupAviatrixVpcResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAviatrixVpcResult
	err := ctx.Invoke("aviatrix:index/getAviatrixVpc:getAviatrixVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAviatrixVpc.
type LookupAviatrixVpcArgs struct {
	Name              string  `pulumi:"name"`
	RouteTablesFilter *string `pulumi:"routeTablesFilter"`
}

// A collection of values returned by getAviatrixVpc.
type LookupAviatrixVpcResult struct {
	AccountName         string   `pulumi:"accountName"`
	AvailabilityDomains []string `pulumi:"availabilityDomains"`
	AviatrixFirenetVpc  bool     `pulumi:"aviatrixFirenetVpc"`
	AviatrixTransitVpc  bool     `pulumi:"aviatrixTransitVpc"`
	AzureVnetResourceId string   `pulumi:"azureVnetResourceId"`
	Cidr                string   `pulumi:"cidr"`
	CloudType           int      `pulumi:"cloudType"`
	FaultDomains        []string `pulumi:"faultDomains"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                        `pulumi:"id"`
	Name              string                        `pulumi:"name"`
	NumOfSubnetPairs  int                           `pulumi:"numOfSubnetPairs"`
	PrivateSubnets    []GetAviatrixVpcPrivateSubnet `pulumi:"privateSubnets"`
	PublicSubnets     []GetAviatrixVpcPublicSubnet  `pulumi:"publicSubnets"`
	Region            string                        `pulumi:"region"`
	ResourceGroup     string                        `pulumi:"resourceGroup"`
	RouteTables       []string                      `pulumi:"routeTables"`
	RouteTablesFilter *string                       `pulumi:"routeTablesFilter"`
	SubnetSize        int                           `pulumi:"subnetSize"`
	Subnets           []GetAviatrixVpcSubnet        `pulumi:"subnets"`
	VpcId             string                        `pulumi:"vpcId"`
}

func LookupAviatrixVpcOutput(ctx *pulumi.Context, args LookupAviatrixVpcOutputArgs, opts ...pulumi.InvokeOption) LookupAviatrixVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAviatrixVpcResult, error) {
			args := v.(LookupAviatrixVpcArgs)
			r, err := LookupAviatrixVpc(ctx, &args, opts...)
			var s LookupAviatrixVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAviatrixVpcResultOutput)
}

// A collection of arguments for invoking getAviatrixVpc.
type LookupAviatrixVpcOutputArgs struct {
	Name              pulumi.StringInput    `pulumi:"name"`
	RouteTablesFilter pulumi.StringPtrInput `pulumi:"routeTablesFilter"`
}

func (LookupAviatrixVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixVpcArgs)(nil)).Elem()
}

// A collection of values returned by getAviatrixVpc.
type LookupAviatrixVpcResultOutput struct{ *pulumi.OutputState }

func (LookupAviatrixVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixVpcResult)(nil)).Elem()
}

func (o LookupAviatrixVpcResultOutput) ToLookupAviatrixVpcResultOutput() LookupAviatrixVpcResultOutput {
	return o
}

func (o LookupAviatrixVpcResultOutput) ToLookupAviatrixVpcResultOutputWithContext(ctx context.Context) LookupAviatrixVpcResultOutput {
	return o
}

func (o LookupAviatrixVpcResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o LookupAviatrixVpcResultOutput) AvailabilityDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []string { return v.AvailabilityDomains }).(pulumi.StringArrayOutput)
}

func (o LookupAviatrixVpcResultOutput) AviatrixFirenetVpc() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) bool { return v.AviatrixFirenetVpc }).(pulumi.BoolOutput)
}

func (o LookupAviatrixVpcResultOutput) AviatrixTransitVpc() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) bool { return v.AviatrixTransitVpc }).(pulumi.BoolOutput)
}

func (o LookupAviatrixVpcResultOutput) AzureVnetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.AzureVnetResourceId }).(pulumi.StringOutput)
}

func (o LookupAviatrixVpcResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Cidr }).(pulumi.StringOutput)
}

func (o LookupAviatrixVpcResultOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) int { return v.CloudType }).(pulumi.IntOutput)
}

func (o LookupAviatrixVpcResultOutput) FaultDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []string { return v.FaultDomains }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAviatrixVpcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAviatrixVpcResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAviatrixVpcResultOutput) NumOfSubnetPairs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) int { return v.NumOfSubnetPairs }).(pulumi.IntOutput)
}

func (o LookupAviatrixVpcResultOutput) PrivateSubnets() GetAviatrixVpcPrivateSubnetArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []GetAviatrixVpcPrivateSubnet { return v.PrivateSubnets }).(GetAviatrixVpcPrivateSubnetArrayOutput)
}

func (o LookupAviatrixVpcResultOutput) PublicSubnets() GetAviatrixVpcPublicSubnetArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []GetAviatrixVpcPublicSubnet { return v.PublicSubnets }).(GetAviatrixVpcPublicSubnetArrayOutput)
}

func (o LookupAviatrixVpcResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupAviatrixVpcResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupAviatrixVpcResultOutput) RouteTables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []string { return v.RouteTables }).(pulumi.StringArrayOutput)
}

func (o LookupAviatrixVpcResultOutput) RouteTablesFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) *string { return v.RouteTablesFilter }).(pulumi.StringPtrOutput)
}

func (o LookupAviatrixVpcResultOutput) SubnetSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) int { return v.SubnetSize }).(pulumi.IntOutput)
}

func (o LookupAviatrixVpcResultOutput) Subnets() GetAviatrixVpcSubnetArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []GetAviatrixVpcSubnet { return v.Subnets }).(GetAviatrixVpcSubnetArrayOutput)
}

func (o LookupAviatrixVpcResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAviatrixVpcResultOutput{})
}
