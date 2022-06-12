package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const targetPodName = "nginx-clean"

type CleanUpController struct {
	client client.Client
	scheme *runtime.Scheme
}

func NewCleanUpController(client client.Client, scheme *runtime.Scheme) *CleanUpController {
	return &CleanUpController{client: client,
		scheme: scheme}
}

func (c *CleanUpController) SetupWithManager(mgr ctrl.Manager) error {
	err := ctrl.NewControllerManagedBy(mgr).
		For(&v1.Pod{}).
		WithEventFilter(predicate.ResourceVersionChangedPredicate{}).
		Complete(c)
	if err != nil {
		klog.Info("CleanUpController setup error.", err)
		return err
	}
	return err
}

func (c *CleanUpController) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var actual v1.Pod
	key := client.ObjectKey{Namespace: req.Namespace, Name: targetPodName}
	if err := c.client.Get(ctx, key, &actual); err != nil {
		klog.Info("nginx-clean pod was deleted")

		klog.Info("Let's start clean up task ...")
	}
	return reconcile.Result{}, nil
}
