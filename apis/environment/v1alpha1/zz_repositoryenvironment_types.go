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

type DeploymentBranchPolicyObservation struct {
}

type DeploymentBranchPolicyParameters struct {

	// +kubebuilder:validation:Required
	CustomBranchPolicies *bool `json:"customBranchPolicies" tf:"custom_branch_policies,omitempty"`

	// +kubebuilder:validation:Required
	ProtectedBranches *bool `json:"protectedBranches" tf:"protected_branches,omitempty"`
}

type RepositoryEnvironmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RepositoryEnvironmentParameters struct {

	// +kubebuilder:validation:Optional
	DeploymentBranchPolicy []DeploymentBranchPolicyParameters `json:"deploymentBranchPolicy,omitempty" tf:"deployment_branch_policy,omitempty"`

	// +kubebuilder:validation:Required
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`

	// +kubebuilder:validation:Optional
	Reviewers []ReviewersParameters `json:"reviewers,omitempty" tf:"reviewers,omitempty"`

	// +kubebuilder:validation:Optional
	WaitTimer *float64 `json:"waitTimer,omitempty" tf:"wait_timer,omitempty"`
}

type ReviewersObservation struct {
}

type ReviewersParameters struct {

	// +crossplane:generate:reference:type=github.com/joakimhew/provider-jet-github/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	Teams []*float64 `json:"teams,omitempty" tf:"teams,omitempty"`

	// +kubebuilder:validation:Optional
	TeamsRefs []v1.Reference `json:"teamsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TeamsSelector *v1.Selector `json:"teamsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Users []*float64 `json:"users,omitempty" tf:"users,omitempty"`
}

// RepositoryEnvironmentSpec defines the desired state of RepositoryEnvironment
type RepositoryEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryEnvironmentParameters `json:"forProvider"`
}

// RepositoryEnvironmentStatus defines the observed state of RepositoryEnvironment.
type RepositoryEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryEnvironment is the Schema for the RepositoryEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type RepositoryEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositoryEnvironmentSpec   `json:"spec"`
	Status            RepositoryEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryEnvironmentList contains a list of RepositoryEnvironments
type RepositoryEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryEnvironment `json:"items"`
}

// Repository type metadata.
var (
	RepositoryEnvironment_Kind             = "RepositoryEnvironment"
	RepositoryEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryEnvironment_Kind}.String()
	RepositoryEnvironment_KindAPIVersion   = RepositoryEnvironment_Kind + "." + CRDGroupVersion.String()
	RepositoryEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryEnvironment{}, &RepositoryEnvironmentList{})
}
