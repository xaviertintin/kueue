/*
Copyright 2022 The Kubernetes Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	v1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	kueuev1beta1 "sigs.k8s.io/kueue/client-go/applyconfiguration/kueue/v1beta1"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=kueue.x-k8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Admission"):
		return &kueuev1beta1.AdmissionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheck"):
		return &kueuev1beta1.AdmissionCheckApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckParametersReference"):
		return &kueuev1beta1.AdmissionCheckParametersReferenceApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckSpec"):
		return &kueuev1beta1.AdmissionCheckSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckState"):
		return &kueuev1beta1.AdmissionCheckStateApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckStatus"):
		return &kueuev1beta1.AdmissionCheckStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueue"):
		return &kueuev1beta1.ClusterQueueApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueuePendingWorkload"):
		return &kueuev1beta1.ClusterQueuePendingWorkloadApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueuePendingWorkloadsStatus"):
		return &kueuev1beta1.ClusterQueuePendingWorkloadsStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueuePreemption"):
		return &kueuev1beta1.ClusterQueuePreemptionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueueSpec"):
		return &kueuev1beta1.ClusterQueueSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueueStatus"):
		return &kueuev1beta1.ClusterQueueStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FlavorFungibility"):
		return &kueuev1beta1.FlavorFungibilityApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FlavorQuotas"):
		return &kueuev1beta1.FlavorQuotasApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FlavorUsage"):
		return &kueuev1beta1.FlavorUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueue"):
		return &kueuev1beta1.LocalQueueApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueFlavorUsage"):
		return &kueuev1beta1.LocalQueueFlavorUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueResourceUsage"):
		return &kueuev1beta1.LocalQueueResourceUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueSpec"):
		return &kueuev1beta1.LocalQueueSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueStatus"):
		return &kueuev1beta1.LocalQueueStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodSet"):
		return &kueuev1beta1.PodSetApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodSetAssignment"):
		return &kueuev1beta1.PodSetAssignmentApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodSetUpdate"):
		return &kueuev1beta1.PodSetUpdateApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ReclaimablePod"):
		return &kueuev1beta1.ReclaimablePodApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceFlavor"):
		return &kueuev1beta1.ResourceFlavorApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceFlavorSpec"):
		return &kueuev1beta1.ResourceFlavorSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceGroup"):
		return &kueuev1beta1.ResourceGroupApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceQuota"):
		return &kueuev1beta1.ResourceQuotaApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceUsage"):
		return &kueuev1beta1.ResourceUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Workload"):
		return &kueuev1beta1.WorkloadApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WorkloadPriorityClass"):
		return &kueuev1beta1.WorkloadPriorityClassApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WorkloadSpec"):
		return &kueuev1beta1.WorkloadSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WorkloadStatus"):
		return &kueuev1beta1.WorkloadStatusApplyConfiguration{}

	}
	return nil
}
