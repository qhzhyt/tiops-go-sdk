package stores

import "tiops/common/config"

type BasicDataStore interface {
	// PutValue 保存k,v到当前store
	PutValue(k string, v interface{})
	// SetValue 向上寻找k，找到设k为v，返回true，否则返回false
	SetValue(k string, v interface{}) bool
	// GetValue 从当前store获取
	GetValue(k string) interface{}

	// LookupValue 向上寻找k，找到返回v
	LookupValue(k string) interface{}
	SaveValue(k string, v interface{}) bool
	LoadValue(k string) interface{}
	SetValueWithName(n, k string, v interface{}) bool
	GetValueWithName(n, k string) interface{}
	LoadAll() bool
	SaveAll() bool
	Upstream() DataStore
	Path() string
}

type StoreMode uint8

const (
	Never = iota
	Once
	Default
	OnSet
	Always
)

var (
	GlobalStoreName    = "tiops"
	WorkspaceStoreName = config.WorkspaceId
	WorkflowStoreName  = config.WorkflowId
	JobStoreName       = config.JobId
	ExecutionStoreName = config.ExecutionId
)

//tiopsApiClient := apiClient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)

type PersistentStore interface {
	LoadValue(p, k string) (interface{}, error)
	SaveValue(p, k string, v interface{}) error
	LoadAll(p string) (map[string]interface{}, error)
	SaveAll(p string, v map[string]interface{}) error
}

type TypedGetter interface {
	GetInt(k string) int
	GetString(k string) string
	GetIntOrDefault(k string, d int) int
	GetStringOrDefault(k string, d string) string
}

type DataStore interface {
	BasicDataStore
	TypedGetter
	GetValueOrDefault(k string, d interface{}) interface{}
}
