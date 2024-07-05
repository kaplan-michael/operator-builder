package controller

import buildsv1alpha1 "github.com/kaplan-michael/operator-builder/api/v1alpha1"

func isSpecEmpty(spec buildsv1alpha1.RecipeSpec) bool {
	if spec.Operator == nil && spec.Operand == nil && spec.Image == nil {
		return true
	}
	return false
}
