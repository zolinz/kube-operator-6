/*
Copyright 2024.

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

package controller

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"
	huntingv1 "zoli.com/hunting/api/v1"
)

// BowhunterReconciler reconciles a Bowhunter object
type BowhunterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=hunting.zoli.com,resources=bowhunters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=hunting.zoli.com,resources=bowhunters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=hunting.zoli.com,resources=bowhunters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Bowhunter object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *BowhunterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	bowhunter := &huntingv1.Bowhunter{}

	if err := r.Get(ctx, req.NamespacedName, bowhunter); err != nil {
		log.Error(err, "unable to fetch CronJob")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	bowhunter.Spec.Bow.Brand = "Bowtech"
	log.Info("Reconciling Bowhunter's Bow ", bowhunter.Spec.Bow.Brand, bowhunter.Spec.Bow.DrawWeight)

	if err := r.Client.Update(ctx, bowhunter); err != nil {
		log.Error(err, "unable to update Bowhunter")
	}

	bowhunter.Status.ReadyToHunt = true
	if err := r.Client.Status().Update(ctx, bowhunter); err != nil {
		log.Error(err, "unable to update status")
	}

	err := errors.New("reconciliation Error")
	log.Error(err, "reconciliation ERROR")

	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *BowhunterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&huntingv1.Bowhunter{}).
		Complete(r)
}
