//go:build !ignore_autogenerated
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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDefaultIngress) DeepCopyInto(out *AstarteDefaultIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
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
func (in *AstarteDefaultIngressAPISpec) DeepCopyInto(out *AstarteDefaultIngressAPISpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = new(bool)
		**out = **in
	}
	if in.ExposeHousekeeping != nil {
		in, out := &in.ExposeHousekeeping, &out.ExposeHousekeeping
		*out = new(bool)
		**out = **in
	}
	if in.ServeMetrics != nil {
		in, out := &in.ServeMetrics, &out.ServeMetrics
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDefaultIngressAPISpec.
func (in *AstarteDefaultIngressAPISpec) DeepCopy() *AstarteDefaultIngressAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDefaultIngressAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDefaultIngressBrokerSpec) DeepCopyInto(out *AstarteDefaultIngressBrokerSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDefaultIngressBrokerSpec.
func (in *AstarteDefaultIngressBrokerSpec) DeepCopy() *AstarteDefaultIngressBrokerSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDefaultIngressBrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDefaultIngressDashboardSpec) DeepCopyInto(out *AstarteDefaultIngressDashboardSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDefaultIngressDashboardSpec.
func (in *AstarteDefaultIngressDashboardSpec) DeepCopy() *AstarteDefaultIngressDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDefaultIngressDashboardSpec)
	in.DeepCopyInto(out)
	return out
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
	out.TypeMeta = in.TypeMeta
	in.API.DeepCopyInto(&out.API)
	in.Dashboard.DeepCopyInto(&out.Dashboard)
	in.Broker.DeepCopyInto(&out.Broker)
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
	out.TypeMeta = in.TypeMeta
	in.APIStatus.DeepCopyInto(&out.APIStatus)
	in.BrokerStatus.DeepCopyInto(&out.BrokerStatus)
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
