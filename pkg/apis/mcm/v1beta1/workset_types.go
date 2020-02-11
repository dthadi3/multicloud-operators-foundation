// licensed Materials - Property of IBM
// 5737-E67
// (C) Copyright IBM Corporation 2016, 2019 All Rights Reserved
// US Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule Contract with IBM Corp.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkSetList is a list of all the work set
type WorkSetList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// List of Cluster objects.
	Items []WorkSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkSet is the work set that will be done on a set of cluster
type WorkSet struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of the work.
	// +optional
	Spec WorkSetSpec `json:"spec,omitempty"`

	// Status describes the result of a work
	// +optional
	Status WorkSetStatus `json:"status,omitempty"`
}

// WorkSetSpec is the spec for workset
type WorkSetSpec struct {
	// Selector for clusters.
	ClusterSelector *metav1.LabelSelector `json:"clusterSelector,omitempty"`

	// Selector for works.
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// Template describes the works that will be created.
	Template WorkTemplateSpec `json:"template,omitempty"`
}

// WorkTemplateSpec describes work created from a template
type WorkTemplateSpec struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the work.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec WorkSpec `json:"spec,omitempty"`
}

// WorkSetStatus describes the work set status
type WorkSetStatus struct {
	// Status of the work set
	Status WorkStatusType `json:"status,omitempty"`

	// Reason is the reason of the status
	Reason string `json:"reason,omitempty"`
}