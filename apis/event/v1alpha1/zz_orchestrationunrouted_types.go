/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CatchAllActionsExtractionObservation struct {
}

type CatchAllActionsExtractionParameters struct {

	// A RE2 regular expression that will be matched against field specified via the source argument. If the regex contains one or more capture groups, their values will be extracted and appended together. If it contains no capture groups, the whole match is used. This field can be ignored for template based extractions.
	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// The path to the event field where the regex will be applied to extract a value. You can use any valid PCL path like event.summary and you can reference previously-defined variables using a path like variables.hostname. This field can be ignored for template based extractions.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The PagerDuty Common Event Format PD-CEF field that will be set with the value from the template or based on regex and source fields.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// A string that will be used to populate the target field. You can reference variables or event data within your template using double curly braces. For example:
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type CatchAllActionsVariableObservation struct {
}

type CatchAllActionsVariableParameters struct {

	// The name of the variable
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Path to a field in an event, in dot-notation. This supports both PD-CEF and non-CEF fields. Eg: Use event.summary for the summary CEF field. Use raw_event.fieldname to read from the original event fieldname data.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Only regex is supported
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The Regex expression to match against. Must use valid RE2 regular expression syntax.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type OrchestrationUnroutedCatchAllActionsObservation struct {
	Suppress *bool `json:"suppress,omitempty" tf:"suppress,omitempty"`
}

type OrchestrationUnroutedCatchAllActionsParameters struct {

	// sets whether the resulting alert status is trigger or resolve. Allowed values are: trigger, resolve
	// +kubebuilder:validation:Optional
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Replace any CEF field or Custom Details object field using custom variables.
	// +kubebuilder:validation:Optional
	Extraction []CatchAllActionsExtractionParameters `json:"extraction,omitempty" tf:"extraction,omitempty"`

	// sets Severity of the resulting alert. Allowed values are: info, error, warning, critical
	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// Populate variables from event payloads and use those variables in other event actions.
	// +kubebuilder:validation:Optional
	Variable []CatchAllActionsVariableParameters `json:"variable,omitempty" tf:"variable,omitempty"`
}

type OrchestrationUnroutedCatchAllObservation struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	// +kubebuilder:validation:Required
	Actions []OrchestrationUnroutedCatchAllActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`
}

type OrchestrationUnroutedCatchAllParameters struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	// +kubebuilder:validation:Required
	Actions []OrchestrationUnroutedCatchAllActionsParameters `json:"actions" tf:"actions,omitempty"`
}

type OrchestrationUnroutedObservation struct {

	// the catch_all actions will be applied if an Event reaches the end of any set without matching any rules in that set.
	// +kubebuilder:validation:Required
	CatchAll []OrchestrationUnroutedCatchAllObservation `json:"catchAll,omitempty" tf:"catch_all,omitempty"`

	// The ID of this set of rules. Rules in other sets can route events into this set using the rule's route_to property.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An Unrouted Orchestration must contain at least a "start" set, but can contain any number of additional sets that are routed to by other rules to form a directional graph.
	// +kubebuilder:validation:Required
	Set []OrchestrationUnroutedSetObservation `json:"set,omitempty" tf:"set,omitempty"`
}

type OrchestrationUnroutedParameters struct {

	// the catch_all actions will be applied if an Event reaches the end of any set without matching any rules in that set.
	// +kubebuilder:validation:Required
	CatchAll []OrchestrationUnroutedCatchAllParameters `json:"catchAll" tf:"catch_all,omitempty"`

	// The Event Orchestration to which this Unrouted Orchestration belongs to.
	// +crossplane:generate:reference:type=Orchestration
	// +kubebuilder:validation:Optional
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// Reference to a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationRef *v1.Reference `json:"eventOrchestrationRef,omitempty" tf:"-"`

	// Selector for a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationSelector *v1.Selector `json:"eventOrchestrationSelector,omitempty" tf:"-"`

	// An Unrouted Orchestration must contain at least a "start" set, but can contain any number of additional sets that are routed to by other rules to form a directional graph.
	// +kubebuilder:validation:Required
	Set []OrchestrationUnroutedSetParameters `json:"set" tf:"set,omitempty"`
}

type OrchestrationUnroutedSetObservation struct {

	// The Unrouted Orchestration evaluates Events against these Rules, one at a time, and applies all the actions for first rule it finds where the event matches the rule's conditions.
	// +kubebuilder:validation:Optional
	Rule []OrchestrationUnroutedSetRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type OrchestrationUnroutedSetParameters struct {

	// The ID of this set of rules. Rules in other sets can route events into this set using the rule's route_to property.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// The Unrouted Orchestration evaluates Events against these Rules, one at a time, and applies all the actions for first rule it finds where the event matches the rule's conditions.
	// +kubebuilder:validation:Optional
	Rule []OrchestrationUnroutedSetRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type OrchestrationUnroutedSetRuleActionsObservation struct {
}

type OrchestrationUnroutedSetRuleActionsParameters struct {

	// sets whether the resulting alert status is trigger or resolve. Allowed values are: trigger, resolve
	// +kubebuilder:validation:Optional
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Replace any CEF field or Custom Details object field using custom variables.
	// +kubebuilder:validation:Optional
	Extraction []RuleActionsExtractionParameters `json:"extraction,omitempty" tf:"extraction,omitempty"`

	// The ID of a Set from this Unrouted Orchestration whose rules you also want to use with events that match this rule.
	// +kubebuilder:validation:Optional
	RouteTo *string `json:"routeTo,omitempty" tf:"route_to,omitempty"`

	// sets Severity of the resulting alert. Allowed values are: info, error, warning, critical
	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// Populate variables from event payloads and use those variables in other event actions.
	// +kubebuilder:validation:Optional
	Variable []RuleActionsVariableParameters `json:"variable,omitempty" tf:"variable,omitempty"`
}

type OrchestrationUnroutedSetRuleObservation struct {

	// The ID of this set of rules. Rules in other sets can route events into this set using the rule's route_to property.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OrchestrationUnroutedSetRuleParameters struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	// +kubebuilder:validation:Required
	Actions []OrchestrationUnroutedSetRuleActionsParameters `json:"actions" tf:"actions,omitempty"`

	// Each of these conditions is evaluated to check if an event matches this rule. The rule is considered a match if any of these conditions match. If none are provided, the event will always match against the rule.
	// +kubebuilder:validation:Optional
	Condition []SetRuleConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Indicates whether the rule is disabled and would therefore not be evaluated.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A description of this rule's purpose.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type RuleActionsExtractionObservation struct {
}

type RuleActionsExtractionParameters struct {

	// A RE2 regular expression that will be matched against field specified via the source argument. If the regex contains one or more capture groups, their values will be extracted and appended together. If it contains no capture groups, the whole match is used. This field can be ignored for template based extractions.
	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// The path to the event field where the regex will be applied to extract a value. You can use any valid PCL path like event.summary and you can reference previously-defined variables using a path like variables.hostname. This field can be ignored for template based extractions.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The PagerDuty Common Event Format PD-CEF field that will be set with the value from the template or based on regex and source fields.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// A string that will be used to populate the target field. You can reference variables or event data within your template using double curly braces. For example:
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type RuleActionsVariableObservation struct {
}

type RuleActionsVariableParameters struct {

	// The name of the variable
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Path to a field in an event, in dot-notation. This supports both PD-CEF and non-CEF fields. Eg: Use event.summary for the summary CEF field. Use raw_event.fieldname to read from the original event fieldname data.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Only regex is supported
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The Regex expression to match against. Must use valid RE2 regular expression syntax.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type SetRuleConditionObservation struct {
}

type SetRuleConditionParameters struct {

	// A PCL condition string.
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`
}

// OrchestrationUnroutedSpec defines the desired state of OrchestrationUnrouted
type OrchestrationUnroutedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrchestrationUnroutedParameters `json:"forProvider"`
}

// OrchestrationUnroutedStatus defines the observed state of OrchestrationUnrouted.
type OrchestrationUnroutedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrchestrationUnroutedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestrationUnrouted is the Schema for the OrchestrationUnrouteds API. Creates and manages an Unrouted Orchestration for a Global Event Orchestration in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type OrchestrationUnrouted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrchestrationUnroutedSpec   `json:"spec"`
	Status            OrchestrationUnroutedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestrationUnroutedList contains a list of OrchestrationUnrouteds
type OrchestrationUnroutedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrchestrationUnrouted `json:"items"`
}

// Repository type metadata.
var (
	OrchestrationUnrouted_Kind             = "OrchestrationUnrouted"
	OrchestrationUnrouted_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrchestrationUnrouted_Kind}.String()
	OrchestrationUnrouted_KindAPIVersion   = OrchestrationUnrouted_Kind + "." + CRDGroupVersion.String()
	OrchestrationUnrouted_GroupVersionKind = CRDGroupVersion.WithKind(OrchestrationUnrouted_Kind)
)

func init() {
	SchemeBuilder.Register(&OrchestrationUnrouted{}, &OrchestrationUnroutedList{})
}
