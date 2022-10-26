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

type PagesObservation struct {
	Custom404 *bool `json:"custom404,omitempty" tf:"custom_404,omitempty"`

	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PagesParameters struct {

	// +kubebuilder:validation:Optional
	Cname *string `json:"cname,omitempty" tf:"cname,omitempty"`

	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`
}

type RepositoryObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	FullName *string `json:"fullName,omitempty" tf:"full_name,omitempty"`

	GitCloneURL *string `json:"gitCloneUrl,omitempty" tf:"git_clone_url,omitempty"`

	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	HTTPCloneURL *string `json:"httpCloneUrl,omitempty" tf:"http_clone_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NodeID *string `json:"nodeId,omitempty" tf:"node_id,omitempty"`

	RepoID *float64 `json:"repoId,omitempty" tf:"repo_id,omitempty"`

	SSHCloneURL *string `json:"sshCloneUrl,omitempty" tf:"ssh_clone_url,omitempty"`

	SvnURL *string `json:"svnUrl,omitempty" tf:"svn_url,omitempty"`
}

type RepositoryParameters struct {

	// +kubebuilder:validation:Optional
	AllowAutoMerge *bool `json:"allowAutoMerge,omitempty" tf:"allow_auto_merge,omitempty"`

	// +kubebuilder:validation:Optional
	AllowMergeCommit *bool `json:"allowMergeCommit,omitempty" tf:"allow_merge_commit,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRebaseMerge *bool `json:"allowRebaseMerge,omitempty" tf:"allow_rebase_merge,omitempty"`

	// +kubebuilder:validation:Optional
	AllowSquashMerge *bool `json:"allowSquashMerge,omitempty" tf:"allow_squash_merge,omitempty"`

	// +kubebuilder:validation:Optional
	ArchiveOnDestroy *bool `json:"archiveOnDestroy,omitempty" tf:"archive_on_destroy,omitempty"`

	// +kubebuilder:validation:Optional
	Archived *bool `json:"archived,omitempty" tf:"archived,omitempty"`

	// +kubebuilder:validation:Optional
	AutoInit *bool `json:"autoInit,omitempty" tf:"auto_init,omitempty"`

	// Can only be set after initial repository creation, and only if the target branch exists
	// +kubebuilder:validation:Optional
	DefaultBranch *string `json:"defaultBranch,omitempty" tf:"default_branch,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteBranchOnMerge *bool `json:"deleteBranchOnMerge,omitempty" tf:"delete_branch_on_merge,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	GitignoreTemplate *string `json:"gitignoreTemplate,omitempty" tf:"gitignore_template,omitempty"`

	// +kubebuilder:validation:Optional
	HasDownloads *bool `json:"hasDownloads,omitempty" tf:"has_downloads,omitempty"`

	// +kubebuilder:validation:Optional
	HasIssues *bool `json:"hasIssues,omitempty" tf:"has_issues,omitempty"`

	// +kubebuilder:validation:Optional
	HasProjects *bool `json:"hasProjects,omitempty" tf:"has_projects,omitempty"`

	// +kubebuilder:validation:Optional
	HasWiki *bool `json:"hasWiki,omitempty" tf:"has_wiki,omitempty"`

	// +kubebuilder:validation:Optional
	HomepageURL *string `json:"homepageUrl,omitempty" tf:"homepage_url,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreVulnerabilityAlertsDuringRead *bool `json:"ignoreVulnerabilityAlertsDuringRead,omitempty" tf:"ignore_vulnerability_alerts_during_read,omitempty"`

	// +kubebuilder:validation:Optional
	IsTemplate *bool `json:"isTemplate,omitempty" tf:"is_template,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseTemplate *string `json:"licenseTemplate,omitempty" tf:"license_template,omitempty"`

	// +kubebuilder:validation:Optional
	MergeCommitMessage *string `json:"mergeCommitMessage,omitempty" tf:"merge_commit_message,omitempty"`

	// +kubebuilder:validation:Optional
	MergeCommitTitle *string `json:"mergeCommitTitle,omitempty" tf:"merge_commit_title,omitempty"`

	// +kubebuilder:validation:Optional
	Pages []PagesParameters `json:"pages,omitempty" tf:"pages,omitempty"`

	// +kubebuilder:validation:Optional
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// +kubebuilder:validation:Optional
	SquashMergeCommitMessage *string `json:"squashMergeCommitMessage,omitempty" tf:"squash_merge_commit_message,omitempty"`

	// +kubebuilder:validation:Optional
	SquashMergeCommitTitle *string `json:"squashMergeCommitTitle,omitempty" tf:"squash_merge_commit_title,omitempty"`

	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`

	// +kubebuilder:validation:Optional
	Topics []*string `json:"topics,omitempty" tf:"topics,omitempty"`

	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// +kubebuilder:validation:Optional
	VulnerabilityAlerts *bool `json:"vulnerabilityAlerts,omitempty" tf:"vulnerability_alerts,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Required
	Branch *string `json:"branch" tf:"branch,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type TemplateObservation struct {
}

type TemplateParameters struct {

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubjet}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
