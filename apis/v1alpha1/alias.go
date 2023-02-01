// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AliasSpec defines the desired state of Alias.
//

type AliasSpec struct {
	// A description of the alias.
	Description *string `json:"description,omitempty"`
	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - MyFunction.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//    * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	FunctionName *string                                  `json:"functionName,omitempty"`
	FunctionRef  *ackv1alpha1.AWSResourceReferenceWrapper `json:"functionRef,omitempty"`
	// The function version that the alias invokes.
	// +kubebuilder:validation:Required
	FunctionVersion *string `json:"functionVersion"`
	// The name of the alias.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing)
	// of the alias.
	RoutingConfig *AliasRoutingConfiguration `json:"routingConfig,omitempty"`
}

// AliasStatus defines the observed state of Alias
type AliasStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// A unique identifier that changes when you update the alias.
	// +kubebuilder:validation:Optional
	RevisionID *string `json:"revisionID,omitempty"`
}

// Alias is the Schema for the Aliases API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Alias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AliasSpec   `json:"spec,omitempty"`
	Status            AliasStatus `json:"status,omitempty"`
}

// AliasList contains a list of Alias
// +kubebuilder:object:root=true
type AliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alias `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alias{}, &AliasList{})
}
