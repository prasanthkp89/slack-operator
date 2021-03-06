/*


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

// ChannelSpec defines the desired state of Channel
type ChannelSpec struct {
	// Name of the slack channel
	// +required
	Name string `json:"name"`

	// Make the channel private or public
	// +optional
	Private bool `json:"private,omitempty"`

	// List of user IDs of the users to invite
	// +kubebuilder:validation:MinItems=1
	// +required
	Users []string `json:"users,omitempty"`

	// Description of the channel
	// +optional
	Description string `json:"description,omitempty"`

	// Topic of the channel
	// +optional
	Topic string `json:"topic,omitempty"`
}

// ChannelStatus defines the observed state of Channel
type ChannelStatus struct {
	// ID of the slack channel
	ID string `json:"id"`

	// Status conditions
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Channel is the Schema for the channels API
type Channel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChannelSpec   `json:"spec,omitempty"`
	Status ChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelList contains a list of Channel
type ChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Channel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Channel{}, &ChannelList{})
}

// GetReconcileStatus - returns conditions, required for making Channel ConditionsStatusAware
func (channel *Channel) GetReconcileStatus() []metav1.Condition {
	return channel.Status.Conditions
}

// SetReconcileStatus - sets status, required for making Channel ConditionsStatusAware
func (channel *Channel) SetReconcileStatus(reconcileStatus []metav1.Condition) {
	channel.Status.Conditions = reconcileStatus
}
