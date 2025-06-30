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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BowhunterSpec defines the desired state of Bowhunter
type BowhunterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Bow Bow `json:"bow,omitempty"`
}

type Bow struct {
	Brand      string `json:"brand,omitempty"`
	DrawWeight int    `json:"drawWeight,omitempty"`
}

// BowhunterStatus defines the observed state of Bowhunter
type BowhunterStatus struct {
	ReadyToHunt bool `json:"readyToHunt,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Bowhunter is the Schema for the bowhunters API
type Bowhunter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BowhunterSpec   `json:"spec,omitempty"`
	Status BowhunterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BowhunterList contains a list of Bowhunter
type BowhunterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bowhunter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bowhunter{}, &BowhunterList{})
}
