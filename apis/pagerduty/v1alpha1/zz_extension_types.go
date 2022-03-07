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

type ExtensionObservation struct {
	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Summary *string `json:"summary,omitempty" tf:"summary,omitempty"`
}

type ExtensionParameters struct {

	// +kubebuilder:validation:Optional
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointURLSecretRef *v1.SecretKeySelector `json:"endpointUrlSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ExtensionObjects []*string `json:"extensionObjects" tf:"extension_objects,omitempty"`

	// +kubebuilder:validation:Required
	ExtensionSchema *string `json:"extensionSchema" tf:"extension_schema,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ExtensionSpec defines the desired state of Extension
type ExtensionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExtensionParameters `json:"forProvider"`
}

// ExtensionStatus defines the observed state of Extension.
type ExtensionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExtensionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Extension is the Schema for the Extensions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerdutyjet}
type Extension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExtensionSpec   `json:"spec"`
	Status            ExtensionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExtensionList contains a list of Extensions
type ExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Extension `json:"items"`
}

// Repository type metadata.
var (
	Extension_Kind             = "Extension"
	Extension_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Extension_Kind}.String()
	Extension_KindAPIVersion   = Extension_Kind + "." + CRDGroupVersion.String()
	Extension_GroupVersionKind = CRDGroupVersion.WithKind(Extension_Kind)
)

func init() {
	SchemeBuilder.Register(&Extension{}, &ExtensionList{})
}
