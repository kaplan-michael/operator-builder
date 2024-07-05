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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RecipeSpec defines the desired state of Recipe
type RecipeSpec struct {
	Operand  *OperandSpec  `json:"operand,omitempty"`
	Operator *OperatorSpec `json:"operator,omitempty"`
	Image    *ImageSpec    `json:"image,omitempty"`
}

// OperandSpec defines the parameters for building the operand
type OperandSpec struct {
	RepoUrl           string  `json:"repoUrl"`
	RepoRef           string  `json:"repoRef"`
	BaseImageRegistry string  `json:"baseImageRegistry,omitempty"`
	ImageName         string  `json:"imageName,omitempty"`
	Version           string  `json:"version,omitempty"`
	BuildImage        string  `json:"buildImage,omitempty"`
	EnvMap            string  `json:"envMap,omitempty"`
	Patches           []Patch `json:"patches,omitempty"`
}

// OperatorSpec defines the parameters for building the operator
type OperatorSpec struct {
	RepoUrl           string  `json:"repoUrl"`
	RepoRef           string  `json:"repoRef"`
	BaseImageRegistry string  `json:"baseImageRegistry,omitempty"`
	ImageName         string  `json:"imageName,omitempty"`
	Version           string  `json:"version,omitempty"`
	Channel           string  `json:"channel,omitempty"`
	DefaultChannel    string  `json:"defaultChannel,omitempty"`
	BuildImage        string  `json:"buildImage,omitempty"`
	EnvMap            string  `json:"envMap,omitempty"`
	Patches           []Patch `json:"patches,omitempty"`
}

// ImageSpec defines the parameters for building an image
type ImageSpec struct {
	RepoURL           string `json:"repoURL"`
	RepoRef           string `json:"repoRef"`
	BaseImageRegistry string `json:"baseImageRegistry"`
	ImageName         string `json:"imageName"`
	Version           string `json:"version"`
	Dockerfile        string `json:"dockerfile"`
}

// Patch defines a patch which can be inline or referenced from a ConfigMap
type Patch struct {
	Name         string `json:"name"`
	Content      string `json:"inline,omitempty"`
	ConfigMapRef string `json:"configMap,omitempty"`
}

// RecipeStatus defines the observed state of BuildDefinition
type RecipeStatus struct {
	LastVersionBuilt string      `json:"lastVersionBuilt,omitempty"`
	LastBuildStatus  string      `json:"lastBuildStatus,omitempty"`
	LastBuildTime    metav1.Time `json:"lastBuildTime,omitempty"`
	Error            bool        `json:"error,omitempty"`
	ErrorMessage     string      `json:"errorMessage,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Recipe is the Schema for the Recipe API
type Recipe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecipeSpec   `json:"spec,omitempty"`
	Status RecipeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RecipeList contains a list of Recipe
type RecipeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Recipe `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Recipe{}, &RecipeList{})
}
