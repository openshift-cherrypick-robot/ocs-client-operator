/*
Copyright 2022 Red Hat, Inc.

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

package templates

import (
	v1k8scsi "k8s.io/api/storage/v1"
	"k8s.io/utils/ptr"
)

var (
	fileFSGroupPolicy = v1k8scsi.FileFSGroupPolicy
)

var CephFSCSIDriver = &v1k8scsi.CSIDriver{
	Spec: v1k8scsi.CSIDriverSpec{
		AttachRequired: ptr.To(true),
		PodInfoOnMount: ptr.To(false),
		FSGroupPolicy:  &fileFSGroupPolicy,
	},
}

var RbdCSIDriver = &v1k8scsi.CSIDriver{
	Spec: v1k8scsi.CSIDriverSpec{
		AttachRequired: ptr.To(true),
		PodInfoOnMount: ptr.To(false),
		FSGroupPolicy:  &fileFSGroupPolicy,
	},
}
