/*
Copyright 2021 The Pixiu Authors.

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

package types

type IdOptions struct {
	Id int64 `uri:"id" binding:"required"`
}

type CloudOptions struct {
	CloudName string `uri:"cloud_name" binding:"required"`
}

type ObjectOptions struct {
	ObjectName string `uri:"object_name" binding:"required"`
}

type NamespaceOptions struct {
	CloudOptions `json:",inline"`

	ObjectOptions `json:",inline"`
}

// NodeOptions todo: 后续整合优化
type NodeOptions struct {
	CloudOptions `json:",inline"`

	ObjectOptions `json:",inline"`
}

type ListOptions struct {
	CloudName string `uri:"cloud_name" binding:"required"`
	Namespace string `uri:"namespace" binding:"required"`
}

type GetOrDeleteOptions struct {
	ListOptions `json:",inline"`

	ObjectName string `uri:"object_name" binding:"required"`
}

type GetOrCreateOptions struct {
	ListOptions `json:",inline,omitempty"`

	ObjectName string `uri:"object_name" binding:"required"`
}

type CreateOptions struct {
	ListOptions `json:",inline,omitempty"`
}

// LogsOptions 日志
type LogsOptions struct {
	GetOrCreateOptions
	ContainerName string `form:"container_name"`
}

type Git struct {
	GitUrl        string `json:"gitUrl,omitempty"`
	Branch        string `json:"branch,omitempty"`
	CredentialsId string `json:"credentialsId,omitempty"`
	ScriptPath    string `json:"scriptPath,omitempty"`
}

type Cicd struct {
	Name     string `json:"name,omitempty"`
	OldName  string `json:"oldName,omitempty"`
	NewName  string `json:"newName,omitempty"`
	ViewName string `json:"viewname,omitempty"`
	Version  string `json:"version,omitempty"`
	Type     string `json:"type,omitempty"`

	Git
}

type User struct {
	Id              int64  `json:"id"`
	ResourceVersion int64  `json:"resource_version"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	Status          int8   `json:"status"`
	Role            string `json:"role"`
	Email           string `json:"email"`
	Description     string `json:"description"`

	TimeOption `json:",inline"`
}

type Password struct {
	UserId          int64  `json:"user_id"`
	OriginPassword  string `json:"origin_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Cloud struct {
	IdMeta     `json:",inline"`
	TimeOption `json:",inline"`

	Name        string `json:"name"`
	Status      int    `json:"status"`     // 0: 正常 1: 异常 2: 正在初始化 3: 删除中
	CloudType   string `json:"cloud_type"` // 0：导入集群（前端又名标准集群） 1: 自建集群
	KubeVersion string `json:"kube_version"`
	KubeConfig  []byte `json:"kube_config"`
	NodeNumber  int    `json:"node_number"`
	Resources   string `json:"resources"`
	Description string `json:"description"`
}

// Node k8s node属性
type Node struct {
	Name             string `json:"name"`
	Status           string `json:"status"`
	Roles            string `json:"roles"`
	CreateAt         string `json:"create_at"`
	Version          string `json:"version"`
	InternalIP       string `json:"internal_ip"`
	OsImage          string `json:"osImage"`
	KernelVersion    string `json:"kernel_version"`
	ContainerRuntime string `json:"container_runtime"`
}

type KubeConfigOptions struct {
	Id                  int64  `json:"id"`
	CloudName           string `json:"cloud_name"`
	ServiceAccount      string `json:"service_account"`
	ClusterRole         string `json:"cluster_role"`
	Config              string `json:"config"`
	ExpirationTimestamp string `json:"expiration_timestamp"`
}
