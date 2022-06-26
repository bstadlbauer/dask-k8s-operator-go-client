package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkerSpec struct {
	Replicas int        `json:"replicas"`
	Spec     v1.PodSpec `json:"spec"`
}

type SchedulerSpec struct {
	Spec    v1.PodSpec     `json:"spec"`
	Service v1.ServiceSpec `json:"service"`
}

type DaskClusterSpec struct {
	Worker    WorkerSpec    `json:"worker"`
	Scheduler SchedulerSpec `json:"scheduler"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DaskCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DaskClusterSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DaskClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaskCluster `json:"items,omitempty"`
}

type DaskWorkerGroupSpec struct {
	Cluster string     `json:"cluster"`
	Worker  WorkerSpec `json:"worker"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DaskWorkerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DaskWorkerGroupSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DaskWorkerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaskWorkerGroup `json:"items,omitempty"`
}

type JobSpec struct {
	Spec v1.PodSpec `json:"spec"`
}

type JobClusterSpec struct {
	Spec DaskClusterSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DaskJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Job               JobSpec        `json:"job"`
	Cluster           JobClusterSpec `json:"cluster"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DaskJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaskJob `json:"items,omitempty"`
}
