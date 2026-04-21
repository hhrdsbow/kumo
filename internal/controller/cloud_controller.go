// Package controller provides Kubernetes controllers for managing cloud resources.
package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kumov1alpha1 "github.com/sivchari/kumo/api/v1alpha1"
)

// CloudReconciler reconciles Cloud objects.
type CloudReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kumo.sivchari.io,resources=clouds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kumo.sivchari.io,resources=clouds/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kumo.sivchari.io,resources=clouds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *CloudReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the Cloud instance
	cloud := &kumov1alpha1.Cloud{}
	if err := r.Get(ctx, req.NamespacedName, cloud); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Reconciling Cloud", "name", cloud.Name, "namespace", cloud.Namespace)

	// Handle deletion
	if !cloud.DeletionTimestamp.IsZero() {
		logger.Info("Cloud is being deleted", "name", cloud.Name)
		return ctrl.Result{}, r.handleDeletion(ctx, cloud)
	}

	// Reconcile the cloud resource
	if err := r.reconcileCloud(ctx, cloud); err != nil {
		logger.Error(err, "Failed to reconcile Cloud", "name", cloud.Name)
		return ctrl.Result{}, fmt.Errorf("reconciling cloud %s/%s: %w", cloud.Namespace, cloud.Name, err)
	}

	return ctrl.Result{}, nil
}

// reconcileCloud handles the main reconciliation logic for a Cloud resource.
func (r *CloudReconciler) reconcileCloud(ctx context.Context, cloud *kumov1alpha1.Cloud) error {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling cloud resources", "provider", cloud.Spec.Provider)

	// Update status to reflect current state
	patch := client.MergeFrom(cloud.DeepCopy())
	cloud.Status.Ready = true
	cloud.Status.Message = fmt.Sprintf("Cloud provider %s is configured", cloud.Spec.Provider)

	if err := r.Status().Patch(ctx, cloud, patch); err != nil {
		return fmt.Errorf("patching cloud status: %w", err)
	}

	return nil
}

// handleDeletion cleans up resources when a Cloud object is deleted.
func (r *CloudReconciler) handleDeletion(ctx context.Context, cloud *kumov1alpha1.Cloud) error {
	logger := log.FromContext(ctx)
	logger.Info("Handling Cloud deletion", "name", cloud.Name)
	// Finalizer cleanup logic would go here
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kumov1alpha1.Cloud{}).
		Complete(r)
}
