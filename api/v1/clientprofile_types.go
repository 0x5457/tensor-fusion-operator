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

// +kubebuilder:validation:Enum=low;medium;high;critical
type QoSLevel string

const (
	QoSLow      QoSLevel = "low"
	QoSMedium   QoSLevel = "medium"
	QoSHigh     QoSLevel = "high"
	QoSCritical QoSLevel = "critical"
)

// ClientProfileSpec defines the desired state of ClientProfile.
type ClientProfileSpec struct {
	// +optional
	PoolName string `json:"poolName,omitempty"`

	// +optional
	Resources Resources `json:"resources,omitempty"`

	// +optional
	// Qos defines the quality of service level for the client.
	Qos QoSLevel `json:"qos,omitempty"`

	IsLocalGPU bool `json:"isLocalGPU"`
}

// ClientProfileStatus defines the observed state of ClientProfile.
type ClientProfileStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClientProfile is the Schema for the clientprofiles API.
type ClientProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClientProfileSpec   `json:"spec,omitempty"`
	Status ClientProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientProfileList contains a list of ClientProfile.
type ClientProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClientProfile{}, &ClientProfileList{})
}
