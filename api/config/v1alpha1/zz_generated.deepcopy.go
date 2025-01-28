//go:build !ignore_autogenerated

/*
Copyright 2025 The Kubernetes Authors.
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
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnection) DeepCopyInto(out *ClientConnection) {
	*out = *in
	if in.QPS != nil {
		in, out := &in.QPS, &out.QPS
		*out = new(float32)
		**out = **in
	}
	if in.Burst != nil {
		in, out := &in.Burst, &out.Burst
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnection.
func (in *ClientConnection) DeepCopy() *ClientConnection {
	if in == nil {
		return nil
	}
	out := new(ClientConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManager.DeepCopyInto(&out.ControllerManager)
	if in.InternalCertManagement != nil {
		in, out := &in.InternalCertManagement, &out.InternalCertManagement
		*out = new(InternalCertManagement)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientConnection != nil {
		in, out := &in.ClientConnection, &out.ClientConnection
		*out = new(ClientConnection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerHealth) DeepCopyInto(out *ControllerHealth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerHealth.
func (in *ControllerHealth) DeepCopy() *ControllerHealth {
	if in == nil {
		return nil
	}
	out := new(ControllerHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManager) DeepCopyInto(out *ControllerManager) {
	*out = *in
	in.Webhook.DeepCopyInto(&out.Webhook)
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(configv1alpha1.LeaderElectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	out.Metrics = in.Metrics
	out.Health = in.Health
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManager.
func (in *ControllerManager) DeepCopy() *ControllerManager {
	if in == nil {
		return nil
	}
	out := new(ControllerManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerMetrics) DeepCopyInto(out *ControllerMetrics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerMetrics.
func (in *ControllerMetrics) DeepCopy() *ControllerMetrics {
	if in == nil {
		return nil
	}
	out := new(ControllerMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerWebhook) DeepCopyInto(out *ControllerWebhook) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerWebhook.
func (in *ControllerWebhook) DeepCopy() *ControllerWebhook {
	if in == nil {
		return nil
	}
	out := new(ControllerWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalCertManagement) DeepCopyInto(out *InternalCertManagement) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.WebhookServiceName != nil {
		in, out := &in.WebhookServiceName, &out.WebhookServiceName
		*out = new(string)
		**out = **in
	}
	if in.WebhookSecretName != nil {
		in, out := &in.WebhookSecretName, &out.WebhookSecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalCertManagement.
func (in *InternalCertManagement) DeepCopy() *InternalCertManagement {
	if in == nil {
		return nil
	}
	out := new(InternalCertManagement)
	in.DeepCopyInto(out)
	return out
}