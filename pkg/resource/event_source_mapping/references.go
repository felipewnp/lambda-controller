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

package event_source_mapping

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mqapitypes "github.com/aws-controllers-k8s/mq-controller/apis/v1alpha1"
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"

	svcapitypes "github.com/aws-controllers-k8s/lambda-controller/apis/v1alpha1"
)

// +kubebuilder:rbac:groups=mq.services.k8s.aws,resources=brokers,verbs=get;list
// +kubebuilder:rbac:groups=mq.services.k8s.aws,resources=brokers/status,verbs=get;list

// ClearResolvedReferences removes any reference values that were made
// concrete in the spec. It returns a copy of the input AWSResource which
// contains the original *Ref values, but none of their respective concrete
// values.
func (rm *resourceManager) ClearResolvedReferences(res acktypes.AWSResource) acktypes.AWSResource {
	ko := rm.concreteResource(res).ko.DeepCopy()

	if ko.Spec.FunctionRef != nil {
		ko.Spec.FunctionName = nil
	}

	if len(ko.Spec.QueueRefs) > 0 {
		ko.Spec.Queues = nil
	}

	return &resource{ko}
}

// ResolveReferences finds if there are any Reference field(s) present
// inside AWSResource passed in the parameter and attempts to resolve those
// reference field(s) into their respective target field(s). It returns a
// copy of the input AWSResource with resolved reference(s), a boolean which
// is set to true if the resource contains any references (regardless of if
// they are resolved successfully) and an error if the passed AWSResource's
// reference field(s) could not be resolved.
func (rm *resourceManager) ResolveReferences(
	ctx context.Context,
	apiReader client.Reader,
	res acktypes.AWSResource,
) (acktypes.AWSResource, bool, error) {
	namespace := res.MetaObject().GetNamespace()
	ko := rm.concreteResource(res).ko

	resourceHasReferences := false
	err := validateReferenceFields(ko)
	if fieldHasReferences, err := rm.resolveReferenceForFunctionName(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	if fieldHasReferences, err := rm.resolveReferenceForQueues(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	return &resource{ko}, resourceHasReferences, err
}

// validateReferenceFields validates the reference field and corresponding
// identifier field.
func validateReferenceFields(ko *svcapitypes.EventSourceMapping) error {

	if ko.Spec.FunctionRef != nil && ko.Spec.FunctionName != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("FunctionName", "FunctionRef")
	}
	if ko.Spec.FunctionRef == nil && ko.Spec.FunctionName == nil {
		return ackerr.ResourceReferenceOrIDRequiredFor("FunctionName", "FunctionRef")
	}

	if len(ko.Spec.QueueRefs) > 0 && len(ko.Spec.Queues) > 0 {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("Queues", "QueueRefs")
	}
	return nil
}

// resolveReferenceForFunctionName reads the resource referenced
// from FunctionRef field and sets the FunctionName
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForFunctionName(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.EventSourceMapping,
) (hasReferences bool, err error) {
	if ko.Spec.FunctionRef != nil && ko.Spec.FunctionRef.From != nil {
		hasReferences = true
		arr := ko.Spec.FunctionRef.From
		if arr.Name == nil || *arr.Name == "" {
			return hasReferences, fmt.Errorf("provided resource reference is nil or empty: FunctionRef")
		}
		obj := &svcapitypes.Function{}
		if err := getReferencedResourceState_Function(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return hasReferences, err
		}
		ko.Spec.FunctionName = (*string)(obj.Spec.Name)
	}

	return hasReferences, nil
}

// getReferencedResourceState_Function looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_Function(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.Function,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"Function",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"Function",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"Function",
			namespace, name)
	}
	if obj.Spec.Name == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"Function",
			namespace, name,
			"Spec.Name")
	}
	return nil
}

// resolveReferenceForQueues reads the resource referenced
// from QueueRefs field and sets the Queues
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForQueues(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.EventSourceMapping,
) (hasReferences bool, err error) {
	for _, f0iter := range ko.Spec.QueueRefs {
		if f0iter != nil && f0iter.From != nil {
			hasReferences = true
			arr := f0iter.From
			if arr.Name == nil || *arr.Name == "" {
				return hasReferences, fmt.Errorf("provided resource reference is nil or empty: QueueRefs")
			}
			obj := &mqapitypes.Broker{}
			if err := getReferencedResourceState_Broker(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
				return hasReferences, err
			}
			if ko.Spec.Queues == nil {
				ko.Spec.Queues = make([]*string, 0, 1)
			}
			ko.Spec.Queues = append(ko.Spec.Queues, (*string)(obj.Status.BrokerID))
		}
	}

	return hasReferences, nil
}

// getReferencedResourceState_Broker looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_Broker(
	ctx context.Context,
	apiReader client.Reader,
	obj *mqapitypes.Broker,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"Broker",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"Broker",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"Broker",
			namespace, name)
	}
	if obj.Status.BrokerID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"Broker",
			namespace, name,
			"Status.BrokerID")
	}
	return nil
}
