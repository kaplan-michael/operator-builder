/*
Copyright 2024 Michael Kaplan.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	buildsv1alpha1 "github.com/kaplan-michael/operator-builder/api/v1alpha1"
)

// RecipeReconciler reconciles a Recipe object
type RecipeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=builds.okderators.io,resources=recipes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=builds.okderators.io,resources=recipes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=builds.okderators.io,resources=recipes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *RecipeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// Fetch the Recipe instance
	recipe := &buildsv1alpha1.Recipe{}
	err := r.Get(ctx, req.NamespacedName, recipe)
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			// Error fetching the object
			l.Error(err, "unable to fetch Recipe")
			return ctrl.Result{}, err
		}
		// Recipe resource not found, could have been deleted after reconcile request.
		// Return and don't requeue
		l.Info("Recipe resource not found. Ignoring...", "Recipe", req.NamespacedName)
		return ctrl.Result{}, nil
	}

	// Log the current status
	l.Info("Successfully fetched Recipe", "Recipe", recipe)

	// Check if the spec is empty and log a warning
	if isSpecEmpty(recipe.Spec) {
		l.Info("Recipe spec is empty, skipping processing", "Recipe", req.NamespacedName)
		return ctrl.Result{}, nil
	}

	// Placeholder for watching PipelineRun events and updating Recipe status
	if err := r.updateRecipeStatus(ctx, recipe); err != nil {
		l.Error(err, "error updating Recipe status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{RequeueAfter: 10 * time.Minute}, nil
}

func (r *RecipeReconciler) pollRepositories(ctx context.Context, recipe *buildsv1alpha1.Recipe) (bool, error) {
	// Placeholder logic for polling repositories
	// Return true if there are updates, false otherwise
	// Implement actual logic later
	return true, nil
}

func (r *RecipeReconciler) triggerTektonPipeline(ctx context.Context, recipe *buildsv1alpha1.Recipe) error {
	// Placeholder logic for triggering Tekton pipeline
	// Implement actual logic later
	return nil
}

func (r *RecipeReconciler) updateRecipeStatus(ctx context.Context, recipe *buildsv1alpha1.Recipe) error {
	l := log.FromContext(ctx)

	// Placeholder logic for updating Recipe status
	// Implement actual logic later
	recipe.Status.LastBuildStatus = "Updated"
	recipe.Status.LastBuildTime = metav1.Now()
	recipe.Status.LastVersionBuilt = recipe.Spec.Operator.Version

	if err := r.Status().Update(ctx, recipe); err != nil {
		l.Error(err, "unable to update Recipe status")
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RecipeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&buildsv1alpha1.Recipe{}).
		Complete(r)
}
