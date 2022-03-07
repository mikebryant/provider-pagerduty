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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NotificationRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NotificationRuleParameters struct {

	// +kubebuilder:validation:Required
	ContactMethod map[string]*string `json:"contactMethod" tf:"contact_method,omitempty"`

	// +kubebuilder:validation:Required
	StartDelayInMinutes *float64 `json:"startDelayInMinutes" tf:"start_delay_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	Urgency *string `json:"urgency" tf:"urgency,omitempty"`

	// +kubebuilder:validation:Required
	UserID *string `json:"userId" tf:"user_id,omitempty"`
}

// NotificationRuleSpec defines the desired state of NotificationRule
type NotificationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationRuleParameters `json:"forProvider"`
}

// NotificationRuleStatus defines the observed state of NotificationRule.
type NotificationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRule is the Schema for the NotificationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerdutyjet}
type NotificationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationRuleSpec   `json:"spec"`
	Status            NotificationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRuleList contains a list of NotificationRules
type NotificationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationRule `json:"items"`
}

// Repository type metadata.
var (
	NotificationRule_Kind             = "NotificationRule"
	NotificationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationRule_Kind}.String()
	NotificationRule_KindAPIVersion   = NotificationRule_Kind + "." + CRDGroupVersion.String()
	NotificationRule_GroupVersionKind = CRDGroupVersion.WithKind(NotificationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationRule{}, &NotificationRuleList{})
}
