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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EscalationRuleObservation) DeepCopyInto(out *EscalationRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EscalationRuleObservation.
func (in *EscalationRuleObservation) DeepCopy() *EscalationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(EscalationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EscalationRuleParameters) DeepCopyInto(out *EscalationRuleParameters) {
	*out = *in
	if in.EscalationDelayInMinutes != nil {
		in, out := &in.EscalationDelayInMinutes, &out.EscalationDelayInMinutes
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]TargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EscalationRuleParameters.
func (in *EscalationRuleParameters) DeepCopy() *EscalationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(EscalationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Play) DeepCopyInto(out *Play) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Play.
func (in *Play) DeepCopy() *Play {
	if in == nil {
		return nil
	}
	out := new(Play)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Play) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayList) DeepCopyInto(out *PlayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Play, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayList.
func (in *PlayList) DeepCopy() *PlayList {
	if in == nil {
		return nil
	}
	out := new(PlayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayObservation) DeepCopyInto(out *PlayObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayObservation.
func (in *PlayObservation) DeepCopy() *PlayObservation {
	if in == nil {
		return nil
	}
	out := new(PlayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayParameters) DeepCopyInto(out *PlayParameters) {
	*out = *in
	if in.ConferenceNumber != nil {
		in, out := &in.ConferenceNumber, &out.ConferenceNumber
		*out = new(string)
		**out = **in
	}
	if in.ConferenceURL != nil {
		in, out := &in.ConferenceURL, &out.ConferenceURL
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Responder != nil {
		in, out := &in.Responder, &out.Responder
		*out = make([]ResponderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RespondersMessage != nil {
		in, out := &in.RespondersMessage, &out.RespondersMessage
		*out = new(string)
		**out = **in
	}
	if in.Runnability != nil {
		in, out := &in.Runnability, &out.Runnability
		*out = new(string)
		**out = **in
	}
	if in.Subscriber != nil {
		in, out := &in.Subscriber, &out.Subscriber
		*out = make([]SubscriberParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubscribersMessage != nil {
		in, out := &in.SubscribersMessage, &out.SubscribersMessage
		*out = new(string)
		**out = **in
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayParameters.
func (in *PlayParameters) DeepCopy() *PlayParameters {
	if in == nil {
		return nil
	}
	out := new(PlayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpec) DeepCopyInto(out *PlaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpec.
func (in *PlaySpec) DeepCopy() *PlaySpec {
	if in == nil {
		return nil
	}
	out := new(PlaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayStatus) DeepCopyInto(out *PlayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayStatus.
func (in *PlayStatus) DeepCopy() *PlayStatus {
	if in == nil {
		return nil
	}
	out := new(PlayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponderObservation) DeepCopyInto(out *ResponderObservation) {
	*out = *in
	if in.EscalationRule != nil {
		in, out := &in.EscalationRule, &out.EscalationRule
		*out = make([]EscalationRuleObservation, len(*in))
		copy(*out, *in)
	}
	if in.NumLoops != nil {
		in, out := &in.NumLoops, &out.NumLoops
		*out = new(float64)
		**out = **in
	}
	if in.OnCallHandoffNotifications != nil {
		in, out := &in.OnCallHandoffNotifications, &out.OnCallHandoffNotifications
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make([]ServiceObservation, len(*in))
		copy(*out, *in)
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = make([]TeamObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponderObservation.
func (in *ResponderObservation) DeepCopy() *ResponderObservation {
	if in == nil {
		return nil
	}
	out := new(ResponderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponderParameters) DeepCopyInto(out *ResponderParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponderParameters.
func (in *ResponderParameters) DeepCopy() *ResponderParameters {
	if in == nil {
		return nil
	}
	out := new(ResponderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceObservation) DeepCopyInto(out *ServiceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceObservation.
func (in *ServiceObservation) DeepCopy() *ServiceObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceParameters) DeepCopyInto(out *ServiceParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceParameters.
func (in *ServiceParameters) DeepCopy() *ServiceParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberObservation) DeepCopyInto(out *SubscriberObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberObservation.
func (in *SubscriberObservation) DeepCopy() *SubscriberObservation {
	if in == nil {
		return nil
	}
	out := new(SubscriberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberParameters) DeepCopyInto(out *SubscriberParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberParameters.
func (in *SubscriberParameters) DeepCopy() *SubscriberParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetObservation) DeepCopyInto(out *TargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetObservation.
func (in *TargetObservation) DeepCopy() *TargetObservation {
	if in == nil {
		return nil
	}
	out := new(TargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetParameters) DeepCopyInto(out *TargetParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetParameters.
func (in *TargetParameters) DeepCopy() *TargetParameters {
	if in == nil {
		return nil
	}
	out := new(TargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamObservation) DeepCopyInto(out *TeamObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamObservation.
func (in *TeamObservation) DeepCopy() *TeamObservation {
	if in == nil {
		return nil
	}
	out := new(TeamObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamParameters) DeepCopyInto(out *TeamParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamParameters.
func (in *TeamParameters) DeepCopy() *TeamParameters {
	if in == nil {
		return nil
	}
	out := new(TeamParameters)
	in.DeepCopyInto(out)
	return out
}
