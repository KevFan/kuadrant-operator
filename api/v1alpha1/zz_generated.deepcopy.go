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

package v1alpha1

import (
	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	out.IssuerRef = in.IssuerRef
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]certmanagerv1.KeyUsage, len(*in))
		copy(*out, *in)
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(certmanagerv1.CertificatePrivateKey)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomWeight) DeepCopyInto(out *CustomWeight) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomWeight.
func (in *CustomWeight) DeepCopy() *CustomWeight {
	if in == nil {
		return nil
	}
	out := new(CustomWeight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSPolicy) DeepCopyInto(out *DNSPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSPolicy.
func (in *DNSPolicy) DeepCopy() *DNSPolicy {
	if in == nil {
		return nil
	}
	out := new(DNSPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSPolicyList) DeepCopyInto(out *DNSPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSPolicyList.
func (in *DNSPolicyList) DeepCopy() *DNSPolicyList {
	if in == nil {
		return nil
	}
	out := new(DNSPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSPolicySpec) DeepCopyInto(out *DNSPolicySpec) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LoadBalancing != nil {
		in, out := &in.LoadBalancing, &out.LoadBalancing
		*out = new(LoadBalancingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSPolicySpec.
func (in *DNSPolicySpec) DeepCopy() *DNSPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DNSPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSPolicyStatus) DeepCopyInto(out *DNSPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSPolicyStatus.
func (in *DNSPolicyStatus) DeepCopy() *DNSPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DNSPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckSpec) DeepCopyInto(out *HealthCheckSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.FailureThreshold != nil {
		in, out := &in.FailureThreshold, &out.FailureThreshold
		*out = new(int)
		**out = **in
	}
	if in.AdditionalHeadersRef != nil {
		in, out := &in.AdditionalHeadersRef, &out.AdditionalHeadersRef
		*out = new(string)
		**out = **in
	}
	if in.ExpectedResponses != nil {
		in, out := &in.ExpectedResponses, &out.ExpectedResponses
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckSpec.
func (in *HealthCheckSpec) DeepCopy() *HealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(HealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckStatus) DeepCopyInto(out *HealthCheckStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckStatus.
func (in *HealthCheckStatus) DeepCopy() *HealthCheckStatus {
	if in == nil {
		return nil
	}
	out := new(HealthCheckStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancingGeo) DeepCopyInto(out *LoadBalancingGeo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancingGeo.
func (in *LoadBalancingGeo) DeepCopy() *LoadBalancingGeo {
	if in == nil {
		return nil
	}
	out := new(LoadBalancingGeo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancingSpec) DeepCopyInto(out *LoadBalancingSpec) {
	*out = *in
	in.Weighted.DeepCopyInto(&out.Weighted)
	out.Geo = in.Geo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancingSpec.
func (in *LoadBalancingSpec) DeepCopy() *LoadBalancingSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancingWeighted) DeepCopyInto(out *LoadBalancingWeighted) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = make([]*CustomWeight, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CustomWeight)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancingWeighted.
func (in *LoadBalancingWeighted) DeepCopy() *LoadBalancingWeighted {
	if in == nil {
		return nil
	}
	out := new(LoadBalancingWeighted)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSPolicy) DeepCopyInto(out *TLSPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSPolicy.
func (in *TLSPolicy) DeepCopy() *TLSPolicy {
	if in == nil {
		return nil
	}
	out := new(TLSPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TLSPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSPolicyList) DeepCopyInto(out *TLSPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TLSPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSPolicyList.
func (in *TLSPolicyList) DeepCopy() *TLSPolicyList {
	if in == nil {
		return nil
	}
	out := new(TLSPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TLSPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSPolicySpec) DeepCopyInto(out *TLSPolicySpec) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	in.CertificateSpec.DeepCopyInto(&out.CertificateSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSPolicySpec.
func (in *TLSPolicySpec) DeepCopy() *TLSPolicySpec {
	if in == nil {
		return nil
	}
	out := new(TLSPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSPolicyStatus) DeepCopyInto(out *TLSPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSPolicyStatus.
func (in *TLSPolicyStatus) DeepCopy() *TLSPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(TLSPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
