// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package version

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "istio.io/api/meta/v1alpha1"
	defaults "github.com/110y/kubetype-gen/cmd/kubetype-gen/testdata/test_input/positive/defaults"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Defaulted is for test
// +kubetype-gen
type Defaulted struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec defaults.Defaulted `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DefaultedList is a collection of Defaulteds.
type DefaultedList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Defaulted `json:"items" protobuf:"bytes,2,rep,name=items"`
}
