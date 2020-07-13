/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by injection-gen. DO NOT EDIT.

package provisionedservice

import (
	context "context"

	v1alpha1 "github.com/vmware-labs/service-bindings/pkg/apis/bindings/v1alpha1"
	provisionedservice "github.com/vmware-labs/service-bindings/pkg/client/injection/reconciler/bindings/v1alpha1/provisionedservice"
	v1 "k8s.io/api/core/v1"
	reconciler "knative.dev/pkg/reconciler"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// newReconciledNormal makes a new reconciler event with event type Normal, and
// reason ProvisionedServiceReconciled.
func newReconciledNormal(namespace, name string) reconciler.Event {
	return reconciler.NewEvent(v1.EventTypeNormal, "ProvisionedServiceReconciled", "ProvisionedService reconciled: \"%s/%s\"", namespace, name)
}

// Reconciler implements controller.Reconciler for ProvisionedService resources.
type Reconciler struct {
	// TODO: add additional requirements here.
}

// Check that our Reconciler implements Interface
var _ provisionedservice.Interface = (*Reconciler)(nil)

// Optionally check that our Reconciler implements Finalizer
//var _ provisionedservice.Finalizer = (*Reconciler)(nil)

// Optionally check that our Reconciler implements ReadOnlyInterface
// Implement this to observe resources even when we are not the leader.
//var _ provisionedservice.ReadOnlyInterface = (*Reconciler)(nil)

// Optionally check that our Reconciler implements ReadOnlyFinalizer
// Implement this to observe tombstoned resources even when we are not
// the leader (best effort).
//var _ provisionedservice.ReadOnlyFinalizer = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *v1alpha1.ProvisionedService) reconciler.Event {
	// TODO: use this if the resource implements InitializeConditions.
	// o.Status.InitializeConditions()

	// TODO: add custom reconciliation logic here.

	// TODO: use this if the object has .status.ObservedGeneration.
	// o.Status.ObservedGeneration = o.Generation
	return newReconciledNormal(o.Namespace, o.Name)
}

// Optionally, use FinalizeKind to add finalizers. FinalizeKind will be called
// when the resource is deleted.
//func (r *Reconciler) FinalizeKind(ctx context.Context, o *v1alpha1.ProvisionedService) reconciler.Event {
//	// TODO: add custom finalization logic here.
//	return nil
//}

// Optionally, use ObserveKind to observe the resource when we are not the leader.
// func (r *Reconciler) ObserveKind(ctx context.Context, o *v1alpha1.ProvisionedService) reconciler.Event {
// 	// TODO: add custom observation logic here.
// 	return nil
// }

// Optionally, use ObserveFinalizeKind to observe resources being finalized when we are no the leader.
//func (r *Reconciler) ObserveFinalizeKind(ctx context.Context, o *v1alpha1.ProvisionedService) reconciler.Event {
// 	// TODO: add custom observation logic here.
//	return nil
//}