// +build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSIdentifiers) DeepCopyInto(out *AWSIdentifiers) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(AWSResourceName)
		**out = **in
	}
	if in.AdditionalKeys != nil {
		in, out := &in.AdditionalKeys, &out.AdditionalKeys
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSIdentifiers.
func (in *AWSIdentifiers) DeepCopy() *AWSIdentifiers {
	if in == nil {
		return nil
	}
	out := new(AWSIdentifiers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptedResource) DeepCopyInto(out *AdoptedResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptedResource.
func (in *AdoptedResource) DeepCopy() *AdoptedResource {
	if in == nil {
		return nil
	}
	out := new(AdoptedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdoptedResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptedResourceList) DeepCopyInto(out *AdoptedResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdoptedResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptedResourceList.
func (in *AdoptedResourceList) DeepCopy() *AdoptedResourceList {
	if in == nil {
		return nil
	}
	out := new(AdoptedResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdoptedResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptedResourceSpec) DeepCopyInto(out *AdoptedResourceSpec) {
	*out = *in
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(TargetKubernetesResource)
		(*in).DeepCopyInto(*out)
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSIdentifiers)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptedResourceSpec.
func (in *AdoptedResourceSpec) DeepCopy() *AdoptedResourceSpec {
	if in == nil {
		return nil
	}
	out := new(AdoptedResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptedResourceStatus) DeepCopyInto(out *AdoptedResourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptedResourceStatus.
func (in *AdoptedResourceStatus) DeepCopy() *AdoptedResourceStatus {
	if in == nil {
		return nil
	}
	out := new(AdoptedResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialObjectMeta) DeepCopyInto(out *PartialObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]v1.OwnerReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialObjectMeta.
func (in *PartialObjectMeta) DeepCopy() *PartialObjectMeta {
	if in == nil {
		return nil
	}
	out := new(PartialObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetadata) DeepCopyInto(out *ResourceMetadata) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(AWSResourceName)
		**out = **in
	}
	if in.OwnerAccountID != nil {
		in, out := &in.OwnerAccountID, &out.OwnerAccountID
		*out = new(AWSAccountID)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetadata.
func (in *ResourceMetadata) DeepCopy() *ResourceMetadata {
	if in == nil {
		return nil
	}
	out := new(ResourceMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
	out.SecretReference = in.SecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetKubernetesResource) DeepCopyInto(out *TargetKubernetesResource) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(PartialObjectMeta)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetKubernetesResource.
func (in *TargetKubernetesResource) DeepCopy() *TargetKubernetesResource {
	if in == nil {
		return nil
	}
	out := new(TargetKubernetesResource)
	in.DeepCopyInto(out)
	return out
}
