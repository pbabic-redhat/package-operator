package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ObjectDeploymentSpec defines the desired state of an ObjectDeployment.
type ObjectDeploymentSpec struct {
	// Number of old revisions in the form of archived ObjectSets to keep.
	// +kubebuilder:default=10
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
	// Selector targets ObjectSets managed by this Deployment.
	// +example={"matchLabels": {"test": "test"}}
	Selector metav1.LabelSelector `json:"selector"`
	// Template to create new ObjectSets from.
	Template ObjectSetTemplate `json:"template"`
	// If Paused is true, the object and its children will not be reconciled.
	Paused bool `json:"paused,omitempty"`
}

// ObjectSetTemplate describes the template to create new ObjectSets from.
type ObjectSetTemplate struct {
	// Common Object Metadata.
	// +example={"labels": {"test": "test"}}
	Metadata metav1.ObjectMeta `json:"metadata"`
	// ObjectSet specification.
	Spec ObjectSetTemplateSpec `json:"spec"`
}

// ObjectDeploymentStatus defines the observed state of an ObjectDeployment.
type ObjectDeploymentStatus struct {
	// Conditions is a list of status conditions ths object is in.
	// +example=[{"type": "Available", "status": "True", "reason": "Available",  "message": ""}]
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Count of hash collisions of the ObjectDeployment.
	CollisionCount *int32 `json:"collisionCount,omitempty"`
	// Computed TemplateHash.
	TemplateHash string `json:"templateHash,omitempty"`
	// Deployment revision.
	Revision int64 `json:"revision,omitempty"`
	// ControllerOf references the owned ObjectSet revisions.
	ControllerOf []ControlledObjectReference `json:"controllerOf,omitempty"`
}

// ObjectDeployment Condition Types.
const (
	ObjectDeploymentAvailable   = "Available"
	ObjectDeploymentProgressing = "Progressing"
	ObjectDeploymentPaused      = "Paused"
)

// ObjectDeployment is the Schema for the ObjectDeployments API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={"objdeploy","od"}
// +kubebuilder:printcolumn:name="Available",type=string,JSONPath=`.status.conditions[?(@.type=="Available")].status`
// +kubebuilder:printcolumn:name="Progressing",type=string,JSONPath=`.status.conditions[?(@.type=="Progressing")].status`
// +kubebuilder:printcolumn:name="Paused",type=string,JSONPath=`.status.conditions[?(@.type=="Paused")].status`
// +kubebuilder:printcolumn:name="Revision",type=string,JSONPath=`.status.revision`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type ObjectDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectDeploymentSpec   `json:"spec,omitempty"`
	Status ObjectDeploymentStatus `json:"status,omitempty"`
}

// ObjectDeploymentList contains a list of ObjectDeployments
// +kubebuilder:object:root=true
type ObjectDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectDeployment `json:"items"`
}

func init() { register(&ObjectDeployment{}, &ObjectDeploymentList{}) }
