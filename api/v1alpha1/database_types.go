/*
Copyright 2021-2022 Nutanix, Inc.

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

/*
GENERATED by operator-sdk
Changes added
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	NDB      NDB      `json:"ndb"`
	Instance Instance `json:"databaseInstance"`
}

// DatabaseStatus defines the observed state of Database
type DatabaseStatus struct {
	IPAddress        string `json:"ipAddress"`
	Id               string `json:"id"`
	Status           string `json:"status"`
	DatabaseServerId string `json:"dbServerId"`
}

// Database is the Schema for the databases API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={"db","dbs"}
// +kubebuilder:printcolumn:name="IP Address",type=string,JSONPath=`.status.ipAddress`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Database Instance ID",type=string,JSONPath=`.status.id`
// +kubebuilder:printcolumn:name="Database Server ID",type=string,JSONPath=`.status.dbServerId`
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseSpec   `json:"spec,omitempty"`
	Status DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// DatabaseList contains a list of Database
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}

// These are required to have a deep copy, object interface implementation
// These are the structs for the Spec and Status

// Details of the NDB installation
type NDB struct {
	ClusterId   string      `json:"clusterId"`
	Credentials Credentials `json:"credentials"`
	Server      string      `json:"server"`
}

type Credentials struct {
	// Username for NDB
	LoginUser string `json:"loginUser"`
	// Password for NDB
	Password string `json:"password"`
	// SSH public key for the database vm
	SSHPublicKey string `json:"sshPublicKey"`
}

// Database instance specific details
type Instance struct {
	// +kubebuilder:default:=database_instance_name
	// Name of the database instance
	DatabaseInstanceName string `json:"databaseInstanceName"`
	// +kubebuilder:default:={"database_one", "database_two", "database_three"}
	// +kubebuilder:validation:MinItems:=1
	// Name of the database to be provisiond in the database instance
	DatabaseNames []string `json:"databaseNames"`
	// Password of the database instance
	Password string `json:"password"`
	// +kubebuilder:validation:Minimum:=10
	// +kubebuilder:default:=10
	// +optional
	// Size of the database instance
	Size int `json:"size"`
	// +kubebuilder:default:=UTC
	// +optional
	TimeZone string `json:"timezone"`
	// +kubebuilder:validation:Enum=mysql;postgres;mongodb
	// +kubebuilder:default:=postgres
	Type string `json:"type"`
	// +optional
	Profiles Profiles `json:"profiles"`
}

type Profiles struct {
	// +optional
	Software Profile `json:"software"`

	// +optional
	Compute Profile `json:"compute"`

	// +optional
	Network Profile `json:"network"`

	// +optional
	DbParam Profile `json:"dbParam"`
}

type Profile struct {
	// +optional
	Id string `json:"id"`
	// +optional
	VersionId string `json:"versionId"`
}

// NDB API client model types
// These are not required to have a deep copy or object interface implementation
// These are essentially 'internal' types

// ##### REQUESTS #####

// +kubebuilder:object:generate:=false
type DatabaseProvisionRequest struct {
	DatabaseType             string           `json:"databaseType"`
	Name                     string           `json:"name"`
	DatabaseDescription      string           `json:"databaseDescription"`
	SoftwareProfileId        string           `json:"softwareProfileId"`
	SoftwareProfileVersionId string           `json:"softwareProfileVersionId"`
	ComputeProfileId         string           `json:"computeProfileId"`
	NetworkProfileId         string           `json:"networkProfileId"`
	DbParameterProfileId     string           `json:"dbParameterProfileId"`
	NewDbServerTimeZone      string           `json:"newDbServerTimeZone"`
	CreateDbServer           bool             `json:"createDbserver"`
	NodeCount                int              `json:"nodeCount"`
	NxClusterId              string           `json:"nxClusterId"`
	SSHPublicKey             string           `json:"sshPublicKey"`
	Clustered                bool             `json:"clustered"`
	AutoTuneStagingDrive     bool             `json:"autoTuneStagingDrive"`
	TimeMachineInfo          TimeMachineInfo  `json:"timeMachineInfo"`
	ActionArguments          []ActionArgument `json:"actionArguments"`
	Nodes                    []Node           `json:"nodes"`
}

// +kubebuilder:object:generate:=false
type DatabaseDeprovisionRequest struct {
	Delete               bool `json:"delete"`
	Remove               bool `json:"remove"`
	SoftRemove           bool `json:"softRemove"`
	Forced               bool `json:"forced"`
	DeleteTimeMachine    bool `json:"deleteTimeMachine"`
	DeleteLogicalCluster bool `json:"deleteLogicalCluster"`
}

// +kubebuilder:object:generate:=false
type DatabaseServerDeprovisionRequest struct {
	Delete            bool `json:"delete"`
	Remove            bool `json:"remove"`
	SoftRemove        bool `json:"softRemove"`
	DeleteVgs         bool `json:"deleteVgs"`
	DeleteVmSnapshots bool `json:"deleteVmSnapshots"`
}

// ##### RESPONSES #####

// +kubebuilder:object:generate:=false
type DatabaseResponse struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Status        string         `json:"status"`
	DatabaseNodes []DatabaseNode `json:"databaseNodes"`
	Properties    []Property     `json:"properties"`
}

// +kubebuilder:object:generate:=false
type TaskInfoSummaryResponse struct {
	Name                 string                    `json:"name"`
	WorkId               string                    `json:"workId"`
	OperationId          string                    `json:"operationId"`
	DbServerId           string                    `json:"dbserverId"`
	Message              string                    `json:"messgae"`
	EntityId             string                    `json:"entityId"`
	EntityName           string                    `json:"entityName"`
	EntityType           string                    `json:"entityType"`
	Status               string                    `json:"status"`
	AssociatedOperations []TaskInfoSummaryResponse `json:"associatedOperations"`
	DependencyReport     interface{}               `json:"dependencyReport"`
}

// +kubebuilder:object:generate:=false
type ProfileResponse struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	EngineType      string `json:"engineType"`
	LatestVersionId string `json:"latestVersionId"`
	Topology        string `json:"topology"`
}

// +kubebuilder:object:generate:=false
type SLAResponse struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	UniqueName         string `json:"uniqueName"`
	Description        string `json:"description"`
	DailyRetention     int    `json:"dailyRetention"`
	WeeklyRetention    int    `json:"weeklyRetention"`
	MonthlyRetention   int    `json:"monthlyRetention"`
	QuarterlyRetention int    `json:"quarterlyRetention"`
	YearlyRetention    int    `json:"yearlyRetention"`
}

// ##### OTHER STRUCTS #####

// +kubebuilder:object:generate:=false
type DatabaseNode struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	DatabaseServerId string `json:"dbServerId"`
}

// +kubebuilder:object:generate:=false
type TimeMachineInfo struct {
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	SlaId            string            `json:"slaId"`
	Schedule         map[string]string `json:"schedule"`
	Tags             []string          `json:"tags"`
	AutoTuneLogDrive bool              `json:"autoTuneLogDrive"`
}

// +kubebuilder:object:generate:=false
type ActionArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +kubebuilder:object:generate:=false
type Node struct {
	Properties []string `json:"properties"`
	VmName     string   `json:"vmName"`
}

//used by database response
// +kubebuilder:object:generate:=false
type Property struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
