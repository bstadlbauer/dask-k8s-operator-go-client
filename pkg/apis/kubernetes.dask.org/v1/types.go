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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=true
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=true
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

type DaskJobSpec struct {
	Job     JobSpec        `json:"job"`
	Cluster JobClusterSpec `json:"cluster"`
}

type JobStatus string

const (
	DaskJobCreated        JobStatus = "JobCreated"
	DaskJobClusterCreated JobStatus = "ClusterCreated"
	DaskJobRunning        JobStatus = "Running"
	DaskJobSuccessful     JobStatus = "Successful"
	DaskJobFailed         JobStatus = "Failed"
)

// DaskJobStatus describes the current status of a Dask Job
type DaskJobStatus struct {
	// The name of the cluster the job is executed on
	ClusterName string `json:"clusterName,omitempty"`
	// The time the job runner pod changed to either Successful or Failing
	EndTime metav1.Time `json:"endTime,omitempty"`
	// The name of the job-runner pod
	JobRunnerPodName string `json:"jobRunnerPodName,omitempty"`
	// JobStatus describes the current status of the job
	JobStatus JobStatus `json:"jobStatus"`
	// Start time records the time the job-runner pod changed into a `running` state
	StartTime metav1.Time `json:"startTime"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=true
type DaskJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DaskJobSpec   `json:"spec"`
	Status            DaskJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DaskJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaskJob `json:"items,omitempty"`
}
