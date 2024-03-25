//go:build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	apiv1beta2 "github.com/kuadrant/authorino/api/v1beta2"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apisv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthPolicy) DeepCopyInto(out *AuthPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthPolicy.
func (in *AuthPolicy) DeepCopy() *AuthPolicy {
	if in == nil {
		return nil
	}
	out := new(AuthPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthPolicyCommonSpec) DeepCopyInto(out *AuthPolicyCommonSpec) {
	*out = *in
	if in.RouteSelectors != nil {
		in, out := &in.RouteSelectors, &out.RouteSelectors
		*out = make([]RouteSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamedPatterns != nil {
		in, out := &in.NamedPatterns, &out.NamedPatterns
		*out = make(map[string]apiv1beta2.PatternExpressions, len(*in))
		for key, val := range *in {
			var outVal []apiv1beta2.PatternExpression
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(apiv1beta2.PatternExpressions, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1beta2.PatternExpressionOrRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthScheme != nil {
		in, out := &in.AuthScheme, &out.AuthScheme
		*out = new(AuthSchemeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthPolicyCommonSpec.
func (in *AuthPolicyCommonSpec) DeepCopy() *AuthPolicyCommonSpec {
	if in == nil {
		return nil
	}
	out := new(AuthPolicyCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthPolicyList) DeepCopyInto(out *AuthPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthPolicyList.
func (in *AuthPolicyList) DeepCopy() *AuthPolicyList {
	if in == nil {
		return nil
	}
	out := new(AuthPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthPolicySpec) DeepCopyInto(out *AuthPolicySpec) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = new(AuthPolicyCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	in.AuthPolicyCommonSpec.DeepCopyInto(&out.AuthPolicyCommonSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthPolicySpec.
func (in *AuthPolicySpec) DeepCopy() *AuthPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AuthPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthPolicyStatus) DeepCopyInto(out *AuthPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthPolicyStatus.
func (in *AuthPolicyStatus) DeepCopy() *AuthPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AuthPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthSchemeSpec) DeepCopyInto(out *AuthSchemeSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = make(map[string]AuthenticationSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]MetadataSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = make(map[string]AuthorizationSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(ResponseSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Callbacks != nil {
		in, out := &in.Callbacks, &out.Callbacks
		*out = make(map[string]CallbackSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthSchemeSpec.
func (in *AuthSchemeSpec) DeepCopy() *AuthSchemeSpec {
	if in == nil {
		return nil
	}
	out := new(AuthSchemeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	in.AuthenticationSpec.DeepCopyInto(&out.AuthenticationSpec)
	in.CommonAuthRuleSpec.DeepCopyInto(&out.CommonAuthRuleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationSpec) DeepCopyInto(out *AuthorizationSpec) {
	*out = *in
	in.AuthorizationSpec.DeepCopyInto(&out.AuthorizationSpec)
	in.CommonAuthRuleSpec.DeepCopyInto(&out.CommonAuthRuleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationSpec.
func (in *AuthorizationSpec) DeepCopy() *AuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CallbackSpec) DeepCopyInto(out *CallbackSpec) {
	*out = *in
	in.CallbackSpec.DeepCopyInto(&out.CallbackSpec)
	in.CommonAuthRuleSpec.DeepCopyInto(&out.CommonAuthRuleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CallbackSpec.
func (in *CallbackSpec) DeepCopy() *CallbackSpec {
	if in == nil {
		return nil
	}
	out := new(CallbackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonAuthRuleSpec) DeepCopyInto(out *CommonAuthRuleSpec) {
	*out = *in
	if in.RouteSelectors != nil {
		in, out := &in.RouteSelectors, &out.RouteSelectors
		*out = make([]RouteSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonAuthRuleSpec.
func (in *CommonAuthRuleSpec) DeepCopy() *CommonAuthRuleSpec {
	if in == nil {
		return nil
	}
	out := new(CommonAuthRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderSuccessResponseSpec) DeepCopyInto(out *HeaderSuccessResponseSpec) {
	*out = *in
	in.SuccessResponseSpec.DeepCopyInto(&out.SuccessResponseSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderSuccessResponseSpec.
func (in *HeaderSuccessResponseSpec) DeepCopy() *HeaderSuccessResponseSpec {
	if in == nil {
		return nil
	}
	out := new(HeaderSuccessResponseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limit) DeepCopyInto(out *Limit) {
	*out = *in
	if in.RouteSelectors != nil {
		in, out := &in.RouteSelectors, &out.RouteSelectors
		*out = make([]RouteSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.When != nil {
		in, out := &in.When, &out.When
		*out = make([]WhenCondition, len(*in))
		copy(*out, *in)
	}
	if in.Counters != nil {
		in, out := &in.Counters, &out.Counters
		*out = make([]ContextSelector, len(*in))
		copy(*out, *in)
	}
	if in.Rates != nil {
		in, out := &in.Rates, &out.Rates
		*out = make([]Rate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limit.
func (in *Limit) DeepCopy() *Limit {
	if in == nil {
		return nil
	}
	out := new(Limit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataSpec) DeepCopyInto(out *MetadataSpec) {
	*out = *in
	in.MetadataSpec.DeepCopyInto(&out.MetadataSpec)
	in.CommonAuthRuleSpec.DeepCopyInto(&out.CommonAuthRuleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataSpec.
func (in *MetadataSpec) DeepCopy() *MetadataSpec {
	if in == nil {
		return nil
	}
	out := new(MetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rate) DeepCopyInto(out *Rate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rate.
func (in *Rate) DeepCopy() *Rate {
	if in == nil {
		return nil
	}
	out := new(Rate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicy) DeepCopyInto(out *RateLimitPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicy.
func (in *RateLimitPolicy) DeepCopy() *RateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyCommonSpec) DeepCopyInto(out *RateLimitPolicyCommonSpec) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(map[string]Limit, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyCommonSpec.
func (in *RateLimitPolicyCommonSpec) DeepCopy() *RateLimitPolicyCommonSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyList) DeepCopyInto(out *RateLimitPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimitPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyList.
func (in *RateLimitPolicyList) DeepCopy() *RateLimitPolicyList {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicySpec) DeepCopyInto(out *RateLimitPolicySpec) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = new(RateLimitPolicyCommonSpec)
		(*in).DeepCopyInto(*out)
	}
	in.RateLimitPolicyCommonSpec.DeepCopyInto(&out.RateLimitPolicyCommonSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicySpec.
func (in *RateLimitPolicySpec) DeepCopy() *RateLimitPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyStatus) DeepCopyInto(out *RateLimitPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyStatus.
func (in *RateLimitPolicyStatus) DeepCopy() *RateLimitPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponseSpec) DeepCopyInto(out *ResponseSpec) {
	*out = *in
	if in.Unauthenticated != nil {
		in, out := &in.Unauthenticated, &out.Unauthenticated
		*out = new(apiv1beta2.DenyWithSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Unauthorized != nil {
		in, out := &in.Unauthorized, &out.Unauthorized
		*out = new(apiv1beta2.DenyWithSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Success.DeepCopyInto(&out.Success)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponseSpec.
func (in *ResponseSpec) DeepCopy() *ResponseSpec {
	if in == nil {
		return nil
	}
	out := new(ResponseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSelector) DeepCopyInto(out *RouteSelector) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]apisv1.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Matches != nil {
		in, out := &in.Matches, &out.Matches
		*out = make([]apisv1.HTTPRouteMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSelector.
func (in *RouteSelector) DeepCopy() *RouteSelector {
	if in == nil {
		return nil
	}
	out := new(RouteSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuccessResponseSpec) DeepCopyInto(out *SuccessResponseSpec) {
	*out = *in
	in.SuccessResponseSpec.DeepCopyInto(&out.SuccessResponseSpec)
	in.CommonAuthRuleSpec.DeepCopyInto(&out.CommonAuthRuleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuccessResponseSpec.
func (in *SuccessResponseSpec) DeepCopy() *SuccessResponseSpec {
	if in == nil {
		return nil
	}
	out := new(SuccessResponseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhenCondition) DeepCopyInto(out *WhenCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhenCondition.
func (in *WhenCondition) DeepCopy() *WhenCondition {
	if in == nil {
		return nil
	}
	out := new(WhenCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrappedSuccessResponseSpec) DeepCopyInto(out *WrappedSuccessResponseSpec) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]HeaderSuccessResponseSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.DynamicMetadata != nil {
		in, out := &in.DynamicMetadata, &out.DynamicMetadata
		*out = make(map[string]SuccessResponseSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrappedSuccessResponseSpec.
func (in *WrappedSuccessResponseSpec) DeepCopy() *WrappedSuccessResponseSpec {
	if in == nil {
		return nil
	}
	out := new(WrappedSuccessResponseSpec)
	in.DeepCopyInto(out)
	return out
}
