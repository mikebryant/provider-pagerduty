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

type AlertGroupingParametersObservation struct {
}

type AlertGroupingParametersParameters struct {

	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AtObservation struct {
}

type AtParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigObservation struct {
}

type ConfigParameters struct {

	// +kubebuilder:validation:Optional
	Aggregate *string `json:"aggregate,omitempty" tf:"aggregate,omitempty"`

	// +kubebuilder:validation:Optional
	Fields []*string `json:"fields,omitempty" tf:"fields,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type DuringSupportHoursObservation struct {
}

type DuringSupportHoursParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency,omitempty"`
}

type IncidentUrgencyRuleObservation struct {
}

type IncidentUrgencyRuleParameters struct {

	// +kubebuilder:validation:Optional
	DuringSupportHours []DuringSupportHoursParameters `json:"duringSupportHours,omitempty" tf:"during_support_hours,omitempty"`

	// +kubebuilder:validation:Optional
	OutsideSupportHours []OutsideSupportHoursParameters `json:"outsideSupportHours,omitempty" tf:"outside_support_hours,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency,omitempty"`
}

type OutsideSupportHoursObservation struct {
}

type OutsideSupportHoursParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency,omitempty"`
}

type ScheduledActionsObservation struct {
}

type ScheduledActionsParameters struct {

	// +kubebuilder:validation:Optional
	At []AtParameters `json:"at,omitempty" tf:"at,omitempty"`

	// +kubebuilder:validation:Optional
	ToUrgency *string `json:"toUrgency,omitempty" tf:"to_urgency,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastIncidentTimestamp *string `json:"lastIncidentTimestamp,omitempty" tf:"last_incident_timestamp,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceParameters struct {

	// +kubebuilder:validation:Optional
	AcknowledgementTimeout *string `json:"acknowledgementTimeout,omitempty" tf:"acknowledgement_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	AlertCreation *string `json:"alertCreation,omitempty" tf:"alert_creation,omitempty"`

	// +kubebuilder:validation:Optional
	AlertGrouping *string `json:"alertGrouping,omitempty" tf:"alert_grouping,omitempty"`

	// +kubebuilder:validation:Optional
	AlertGroupingParameters []AlertGroupingParametersParameters `json:"alertGroupingParameters,omitempty" tf:"alert_grouping_parameters,omitempty"`

	// +kubebuilder:validation:Optional
	AlertGroupingTimeout *string `json:"alertGroupingTimeout,omitempty" tf:"alert_grouping_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	AutoResolveTimeout *string `json:"autoResolveTimeout,omitempty" tf:"auto_resolve_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	EscalationPolicy *string `json:"escalationPolicy" tf:"escalation_policy,omitempty"`

	// +kubebuilder:validation:Optional
	IncidentUrgencyRule []IncidentUrgencyRuleParameters `json:"incidentUrgencyRule,omitempty" tf:"incident_urgency_rule,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduledActions []ScheduledActionsParameters `json:"scheduledActions,omitempty" tf:"scheduled_actions,omitempty"`

	// +kubebuilder:validation:Optional
	SupportHours []SupportHoursParameters `json:"supportHours,omitempty" tf:"support_hours,omitempty"`
}

type SupportHoursObservation struct {
}

type SupportHoursParameters struct {

	// +kubebuilder:validation:Optional
	DaysOfWeek []*float64 `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerdutyjet}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
