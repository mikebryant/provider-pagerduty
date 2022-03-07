//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servicenow) DeepCopyInto(out *Servicenow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servicenow.
func (in *Servicenow) DeepCopy() *Servicenow {
	if in == nil {
		return nil
	}
	out := new(Servicenow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Servicenow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicenowList) DeepCopyInto(out *ServicenowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Servicenow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicenowList.
func (in *ServicenowList) DeepCopy() *ServicenowList {
	if in == nil {
		return nil
	}
	out := new(ServicenowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicenowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicenowObservation) DeepCopyInto(out *ServicenowObservation) {
	*out = *in
	if in.HTMLURL != nil {
		in, out := &in.HTMLURL, &out.HTMLURL
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicenowObservation.
func (in *ServicenowObservation) DeepCopy() *ServicenowObservation {
	if in == nil {
		return nil
	}
	out := new(ServicenowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicenowParameters) DeepCopyInto(out *ServicenowParameters) {
	*out = *in
	if in.EndpointURLSecretRef != nil {
		in, out := &in.EndpointURLSecretRef, &out.EndpointURLSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ExtensionObjects != nil {
		in, out := &in.ExtensionObjects, &out.ExtensionObjects
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExtensionSchema != nil {
		in, out := &in.ExtensionSchema, &out.ExtensionSchema
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Referer != nil {
		in, out := &in.Referer, &out.Referer
		*out = new(string)
		**out = **in
	}
	out.SnowPasswordSecretRef = in.SnowPasswordSecretRef
	if in.SnowUser != nil {
		in, out := &in.SnowUser, &out.SnowUser
		*out = new(string)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(string)
		**out = **in
	}
	if in.SyncOptions != nil {
		in, out := &in.SyncOptions, &out.SyncOptions
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.TaskType != nil {
		in, out := &in.TaskType, &out.TaskType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicenowParameters.
func (in *ServicenowParameters) DeepCopy() *ServicenowParameters {
	if in == nil {
		return nil
	}
	out := new(ServicenowParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicenowSpec) DeepCopyInto(out *ServicenowSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicenowSpec.
func (in *ServicenowSpec) DeepCopy() *ServicenowSpec {
	if in == nil {
		return nil
	}
	out := new(ServicenowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicenowStatus) DeepCopyInto(out *ServicenowStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicenowStatus.
func (in *ServicenowStatus) DeepCopy() *ServicenowStatus {
	if in == nil {
		return nil
	}
	out := new(ServicenowStatus)
	in.DeepCopyInto(out)
	return out
}
