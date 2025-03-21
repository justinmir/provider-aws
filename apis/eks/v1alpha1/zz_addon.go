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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// AddonParameters defines the desired state of Addon
type AddonParameters struct {
	// Region is which region the Addon will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The name of the add-on. The name must match one of the names that DescribeAddonVersions
	// (https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeAddonVersions.html)
	// returns.
	// +kubebuilder:validation:Required
	AddonName *string `json:"addonName"`
	// The version of the add-on. The version must match one of the versions returned
	// by DescribeAddonVersions (https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeAddonVersions.html).
	AddonVersion *string `json:"addonVersion,omitempty"`
	// The set of configuration values for the add-on that's created. The values
	// that you provide are validated against the schema in DescribeAddonConfiguration
	// (https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeAddonConfiguration.html).
	ConfigurationValues *string `json:"configurationValues,omitempty"`
	// How to resolve field value conflicts for an Amazon EKS add-on. Conflicts
	// are handled based on the value you choose:
	//
	//    * None – If the self-managed version of the add-on is installed on your
	//    cluster, Amazon EKS doesn't change the value. Creation of the add-on might
	//    fail.
	//
	//    * Overwrite – If the self-managed version of the add-on is installed
	//    on your cluster and the Amazon EKS default value is different than the
	//    existing value, Amazon EKS changes the value to the Amazon EKS default
	//    value.
	//
	//    * Preserve – This is similar to the NONE option. If the self-managed
	//    version of the add-on is installed on your cluster Amazon EKS doesn't
	//    change the add-on resource properties. Creation of the add-on might fail
	//    if conflicts are detected. This option works differently during the update
	//    operation. For more information, see UpdateAddon (https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html).
	//
	// If you don't currently have the self-managed version of the add-on installed
	// on your cluster, the Amazon EKS add-on is installed. Amazon EKS sets all
	// values to default values, regardless of the option that you specify.
	ResolveConflicts *string `json:"resolveConflicts,omitempty"`
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's
	// service account. The role must be assigned the IAM permissions required by
	// the add-on. If you don't specify an existing IAM role, then the add-on uses
	// the permissions assigned to the node IAM role. For more information, see
	// Amazon EKS node IAM role (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// in the Amazon EKS User Guide.
	//
	// To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC)
	// provider created for your cluster. For more information, see Enabling IAM
	// roles for service accounts on your cluster (https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleARN *string `json:"serviceAccountRoleARN,omitempty"`
	// The metadata to apply to the cluster to assist with categorization and organization.
	// Each tag consists of a key and an optional value. You define both.
	Tags                  map[string]*string `json:"tags,omitempty"`
	CustomAddonParameters `json:",inline"`
}

// AddonSpec defines the desired state of Addon
type AddonSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AddonParameters `json:"forProvider"`
}

// AddonObservation defines the observed state of Addon
type AddonObservation struct {
	// The Amazon Resource Name (ARN) of the add-on.
	AddonARN *string `json:"addonARN,omitempty"`
	// The name of the cluster.
	ClusterName *string `json:"clusterName,omitempty"`
	// The date and time that the add-on was created.
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// An object that represents the health of the add-on.
	Health *AddonHealth `json:"health,omitempty"`
	// Information about an Amazon EKS add-on from the Amazon Web Services Marketplace.
	MarketplaceInformation *MarketplaceInformation `json:"marketplaceInformation,omitempty"`
	// The date and time that the add-on was last modified.
	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// The owner of the add-on.
	Owner *string `json:"owner,omitempty"`
	// The publisher of the add-on.
	Publisher *string `json:"publisher,omitempty"`
	// The status of the add-on.
	Status *string `json:"status,omitempty"`

	CustomAddonObservation `json:",inline"`
}

// AddonStatus defines the observed state of Addon.
type AddonStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AddonObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Addon is the Schema for the Addons API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Addon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddonSpec   `json:"spec"`
	Status            AddonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonList contains a list of Addons
type AddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Addon `json:"items"`
}

// Repository type metadata.
var (
	AddonKind             = "Addon"
	AddonGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddonKind}.String()
	AddonKindAPIVersion   = AddonKind + "." + GroupVersion.String()
	AddonGroupVersionKind = GroupVersion.WithKind(AddonKind)
)

func init() {
	SchemeBuilder.Register(&Addon{}, &AddonList{})
}
