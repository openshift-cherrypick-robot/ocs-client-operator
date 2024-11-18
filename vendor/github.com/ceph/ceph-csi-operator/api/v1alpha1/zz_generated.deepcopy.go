//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in BlockPoolIdPair) DeepCopyInto(out *BlockPoolIdPair) {
	{
		in := &in
		*out = make(BlockPoolIdPair, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockPoolIdPair.
func (in BlockPoolIdPair) DeepCopy() BlockPoolIdPair {
	if in == nil {
		return nil
	}
	out := new(BlockPoolIdPair)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephConnection) DeepCopyInto(out *CephConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephConnection.
func (in *CephConnection) DeepCopy() *CephConnection {
	if in == nil {
		return nil
	}
	out := new(CephConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CephConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephConnectionList) DeepCopyInto(out *CephConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CephConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephConnectionList.
func (in *CephConnectionList) DeepCopy() *CephConnectionList {
	if in == nil {
		return nil
	}
	out := new(CephConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CephConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephConnectionSpec) DeepCopyInto(out *CephConnectionSpec) {
	*out = *in
	if in.Monitors != nil {
		in, out := &in.Monitors, &out.Monitors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReadAffinity != nil {
		in, out := &in.ReadAffinity, &out.ReadAffinity
		*out = new(ReadAffinitySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephConnectionSpec.
func (in *CephConnectionSpec) DeepCopy() *CephConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(CephConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephConnectionStatus) DeepCopyInto(out *CephConnectionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephConnectionStatus.
func (in *CephConnectionStatus) DeepCopy() *CephConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(CephConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephFsConfigSpec) DeepCopyInto(out *CephFsConfigSpec) {
	*out = *in
	if in.KernelMountOptions != nil {
		in, out := &in.KernelMountOptions, &out.KernelMountOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FuseMountOptions != nil {
		in, out := &in.FuseMountOptions, &out.FuseMountOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephFsConfigSpec.
func (in *CephFsConfigSpec) DeepCopy() *CephFsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CephFsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfile) DeepCopyInto(out *ClientProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfile.
func (in *ClientProfile) DeepCopy() *ClientProfile {
	if in == nil {
		return nil
	}
	out := new(ClientProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfileList) DeepCopyInto(out *ClientProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClientProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfileList.
func (in *ClientProfileList) DeepCopy() *ClientProfileList {
	if in == nil {
		return nil
	}
	out := new(ClientProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfileMapping) DeepCopyInto(out *ClientProfileMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfileMapping.
func (in *ClientProfileMapping) DeepCopy() *ClientProfileMapping {
	if in == nil {
		return nil
	}
	out := new(ClientProfileMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientProfileMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfileMappingList) DeepCopyInto(out *ClientProfileMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClientProfileMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfileMappingList.
func (in *ClientProfileMappingList) DeepCopy() *ClientProfileMappingList {
	if in == nil {
		return nil
	}
	out := new(ClientProfileMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientProfileMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfileMappingSpec) DeepCopyInto(out *ClientProfileMappingSpec) {
	*out = *in
	if in.Mappings != nil {
		in, out := &in.Mappings, &out.Mappings
		*out = make([]MappingsSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfileMappingSpec.
func (in *ClientProfileMappingSpec) DeepCopy() *ClientProfileMappingSpec {
	if in == nil {
		return nil
	}
	out := new(ClientProfileMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfileMappingStatus) DeepCopyInto(out *ClientProfileMappingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfileMappingStatus.
func (in *ClientProfileMappingStatus) DeepCopy() *ClientProfileMappingStatus {
	if in == nil {
		return nil
	}
	out := new(ClientProfileMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfileSpec) DeepCopyInto(out *ClientProfileSpec) {
	*out = *in
	out.CephConnectionRef = in.CephConnectionRef
	if in.CephFs != nil {
		in, out := &in.CephFs, &out.CephFs
		*out = new(CephFsConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Rbd != nil {
		in, out := &in.Rbd, &out.Rbd
		*out = new(RbdConfigSpec)
		**out = **in
	}
	if in.Nfs != nil {
		in, out := &in.Nfs, &out.Nfs
		*out = new(NfsConfigSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfileSpec.
func (in *ClientProfileSpec) DeepCopy() *ClientProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ClientProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientProfileStatus) DeepCopyInto(out *ClientProfileStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientProfileStatus.
func (in *ClientProfileStatus) DeepCopy() *ClientProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ClientProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerPluginResourcesSpec) DeepCopyInto(out *ControllerPluginResourcesSpec) {
	*out = *in
	if in.Attacher != nil {
		in, out := &in.Attacher, &out.Attacher
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshotter != nil {
		in, out := &in.Snapshotter, &out.Snapshotter
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Resizer != nil {
		in, out := &in.Resizer, &out.Resizer
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Provisioner != nil {
		in, out := &in.Provisioner, &out.Provisioner
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.OMapGenerator != nil {
		in, out := &in.OMapGenerator, &out.OMapGenerator
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Liveness != nil {
		in, out := &in.Liveness, &out.Liveness
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.LogRotator != nil {
		in, out := &in.LogRotator, &out.LogRotator
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Plugin != nil {
		in, out := &in.Plugin, &out.Plugin
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerPluginResourcesSpec.
func (in *ControllerPluginResourcesSpec) DeepCopy() *ControllerPluginResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerPluginResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerPluginSpec) DeepCopyInto(out *ControllerPluginSpec) {
	*out = *in
	in.PodCommonSpec.DeepCopyInto(&out.PodCommonSpec)
	if in.DeploymentStrategy != nil {
		in, out := &in.DeploymentStrategy, &out.DeploymentStrategy
		*out = new(appsv1.DeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Privileged != nil {
		in, out := &in.Privileged, &out.Privileged
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerPluginSpec.
func (in *ControllerPluginSpec) DeepCopy() *ControllerPluginSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerPluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Driver) DeepCopyInto(out *Driver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Driver.
func (in *Driver) DeepCopy() *Driver {
	if in == nil {
		return nil
	}
	out := new(Driver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Driver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverList) DeepCopyInto(out *DriverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Driver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverList.
func (in *DriverList) DeepCopy() *DriverList {
	if in == nil {
		return nil
	}
	out := new(DriverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DriverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverSpec) DeepCopyInto(out *DriverSpec) {
	*out = *in
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(LogSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageSet != nil {
		in, out := &in.ImageSet, &out.ImageSet
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.EnableMetadata != nil {
		in, out := &in.EnableMetadata, &out.EnableMetadata
		*out = new(bool)
		**out = **in
	}
	if in.GenerateOMapInfo != nil {
		in, out := &in.GenerateOMapInfo, &out.GenerateOMapInfo
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionSpec)
		**out = **in
	}
	if in.NodePlugin != nil {
		in, out := &in.NodePlugin, &out.NodePlugin
		*out = new(NodePluginSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerPlugin != nil {
		in, out := &in.ControllerPlugin, &out.ControllerPlugin
		*out = new(ControllerPluginSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AttachRequired != nil {
		in, out := &in.AttachRequired, &out.AttachRequired
		*out = new(bool)
		**out = **in
	}
	if in.Liveness != nil {
		in, out := &in.Liveness, &out.Liveness
		*out = new(LivenessSpec)
		**out = **in
	}
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(LeaderElectionSpec)
		**out = **in
	}
	if in.DeployCsiAddons != nil {
		in, out := &in.DeployCsiAddons, &out.DeployCsiAddons
		*out = new(bool)
		**out = **in
	}
	if in.KernelMountOptions != nil {
		in, out := &in.KernelMountOptions, &out.KernelMountOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FuseMountOptions != nil {
		in, out := &in.FuseMountOptions, &out.FuseMountOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverSpec.
func (in *DriverSpec) DeepCopy() *DriverSpec {
	if in == nil {
		return nil
	}
	out := new(DriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverStatus) DeepCopyInto(out *DriverStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverStatus.
func (in *DriverStatus) DeepCopy() *DriverStatus {
	if in == nil {
		return nil
	}
	out := new(DriverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSpec) DeepCopyInto(out *EncryptionSpec) {
	*out = *in
	out.ConfigMapRef = in.ConfigMapRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSpec.
func (in *EncryptionSpec) DeepCopy() *EncryptionSpec {
	if in == nil {
		return nil
	}
	out := new(EncryptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionSpec) DeepCopyInto(out *LeaderElectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionSpec.
func (in *LeaderElectionSpec) DeepCopy() *LeaderElectionSpec {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LivenessSpec) DeepCopyInto(out *LivenessSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LivenessSpec.
func (in *LivenessSpec) DeepCopy() *LivenessSpec {
	if in == nil {
		return nil
	}
	out := new(LivenessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogRotationSpec) DeepCopyInto(out *LogRotationSpec) {
	*out = *in
	out.MaxLogSize = in.MaxLogSize.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogRotationSpec.
func (in *LogRotationSpec) DeepCopy() *LogRotationSpec {
	if in == nil {
		return nil
	}
	out := new(LogRotationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSpec) DeepCopyInto(out *LogSpec) {
	*out = *in
	if in.Rotation != nil {
		in, out := &in.Rotation, &out.Rotation
		*out = new(LogRotationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSpec.
func (in *LogSpec) DeepCopy() *LogSpec {
	if in == nil {
		return nil
	}
	out := new(LogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MappingsSpec) DeepCopyInto(out *MappingsSpec) {
	*out = *in
	if in.BlockPoolIdMapping != nil {
		in, out := &in.BlockPoolIdMapping, &out.BlockPoolIdMapping
		*out = make([]BlockPoolIdPair, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(BlockPoolIdPair, len(*in))
				copy(*out, *in)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MappingsSpec.
func (in *MappingsSpec) DeepCopy() *MappingsSpec {
	if in == nil {
		return nil
	}
	out := new(MappingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NfsConfigSpec) DeepCopyInto(out *NfsConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NfsConfigSpec.
func (in *NfsConfigSpec) DeepCopy() *NfsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NfsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePluginResourcesSpec) DeepCopyInto(out *NodePluginResourcesSpec) {
	*out = *in
	if in.Registrar != nil {
		in, out := &in.Registrar, &out.Registrar
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Liveness != nil {
		in, out := &in.Liveness, &out.Liveness
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.LogRotator != nil {
		in, out := &in.LogRotator, &out.LogRotator
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Plugin != nil {
		in, out := &in.Plugin, &out.Plugin
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePluginResourcesSpec.
func (in *NodePluginResourcesSpec) DeepCopy() *NodePluginResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(NodePluginResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePluginSpec) DeepCopyInto(out *NodePluginSpec) {
	*out = *in
	in.PodCommonSpec.DeepCopyInto(&out.PodCommonSpec)
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(appsv1.DaemonSetUpdateStrategy)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.EnableSeLinuxHostMount != nil {
		in, out := &in.EnableSeLinuxHostMount, &out.EnableSeLinuxHostMount
		*out = new(bool)
		**out = **in
	}
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = new(TopologySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePluginSpec.
func (in *NodePluginSpec) DeepCopy() *NodePluginSpec {
	if in == nil {
		return nil
	}
	out := new(NodePluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfig) DeepCopyInto(out *OperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfig.
func (in *OperatorConfig) DeepCopy() *OperatorConfig {
	if in == nil {
		return nil
	}
	out := new(OperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfigList) DeepCopyInto(out *OperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfigList.
func (in *OperatorConfigList) DeepCopy() *OperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(OperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfigSpec) DeepCopyInto(out *OperatorConfigSpec) {
	*out = *in
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(OperatorLogSpec)
		**out = **in
	}
	if in.DriverSpecDefaults != nil {
		in, out := &in.DriverSpecDefaults, &out.DriverSpecDefaults
		*out = new(DriverSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfigSpec.
func (in *OperatorConfigSpec) DeepCopy() *OperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfigStatus) DeepCopyInto(out *OperatorConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfigStatus.
func (in *OperatorConfigStatus) DeepCopy() *OperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorLogSpec) DeepCopyInto(out *OperatorLogSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorLogSpec.
func (in *OperatorLogSpec) DeepCopy() *OperatorLogSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorLogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodCommonSpec) DeepCopyInto(out *PodCommonSpec) {
	*out = *in
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.PrioritylClassName != nil {
		in, out := &in.PrioritylClassName, &out.PrioritylClassName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VolumeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodCommonSpec.
func (in *PodCommonSpec) DeepCopy() *PodCommonSpec {
	if in == nil {
		return nil
	}
	out := new(PodCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdConfigSpec) DeepCopyInto(out *RbdConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdConfigSpec.
func (in *RbdConfigSpec) DeepCopy() *RbdConfigSpec {
	if in == nil {
		return nil
	}
	out := new(RbdConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadAffinitySpec) DeepCopyInto(out *ReadAffinitySpec) {
	*out = *in
	if in.CrushLocationLabels != nil {
		in, out := &in.CrushLocationLabels, &out.CrushLocationLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadAffinitySpec.
func (in *ReadAffinitySpec) DeepCopy() *ReadAffinitySpec {
	if in == nil {
		return nil
	}
	out := new(ReadAffinitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopologySpec) DeepCopyInto(out *TopologySpec) {
	*out = *in
	if in.DomainLabels != nil {
		in, out := &in.DomainLabels, &out.DomainLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologySpec.
func (in *TopologySpec) DeepCopy() *TopologySpec {
	if in == nil {
		return nil
	}
	out := new(TopologySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
	in.Mount.DeepCopyInto(&out.Mount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
