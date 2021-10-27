// +build !ignore_autogenerated

/*
  This file is part of Astarte.

  Copyright 2020 Ispirata Srl

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDefaultIngress) DeepCopyInto(out *AstarteDefaultIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDefaultIngress.
func (in *AstarteDefaultIngress) DeepCopy() *AstarteDefaultIngress {
	if in == nil {
		return nil
	}
	out := new(AstarteDefaultIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AstarteDefaultIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDefaultIngressList) DeepCopyInto(out *AstarteDefaultIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AstarteDefaultIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDefaultIngressList.
func (in *AstarteDefaultIngressList) DeepCopy() *AstarteDefaultIngressList {
	if in == nil {
		return nil
	}
	out := new(AstarteDefaultIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AstarteDefaultIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDefaultIngressSpec) DeepCopyInto(out *AstarteDefaultIngressSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDefaultIngressSpec.
func (in *AstarteDefaultIngressSpec) DeepCopy() *AstarteDefaultIngressSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDefaultIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDefaultIngressStatus) DeepCopyInto(out *AstarteDefaultIngressStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDefaultIngressStatus.
func (in *AstarteDefaultIngressStatus) DeepCopy() *AstarteDefaultIngressStatus {
	if in == nil {
		return nil
	}
	out := new(AstarteDefaultIngressStatus)
	in.DeepCopyInto(out)
	return out
}
