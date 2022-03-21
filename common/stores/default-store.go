package stores

type defaultDataStore struct {
	BasicDataStore
}

func (s *defaultDataStore) GetInt(k string) int {
	switch v := s.GetValue(k).(type) {
	case int:
		return v
	default:
		return 0
	}
}

func (s *defaultDataStore) GetString(k string) string {
	switch v := s.GetValue(k).(type) {
	case string:
		return v
	default:
		return ""
	}
}

func (s *defaultDataStore) GetIntOrDefault(k string, d int) int {
	switch v := s.GetValue(k).(type) {
	case int:
		return v
	default:
		return d
	}
}

func (s *defaultDataStore) GetStringOrDefault(k string, d string) string {
	switch v := s.GetValue(k).(type) {
	case string:
		return v
	default:
		return d
	}
}

func (s *defaultDataStore) GetValueOrDefault(k string, d interface{}) interface{} {
	v := s.GetValue(k)
	if v != nil {
		return v
	}
	return d
}

func NewDefaultDataStore(upstream DataStore, name string, mode StoreMode) DataStore {
	return &defaultDataStore{BasicDataStore: NewBasicDataStore(upstream, name, mode)}
}

func NewActionNodeStore(nodeId string) DataStore {
	return NewDefaultDataStore(ExecutionStore, nodeId, Never)
}

var (
	GlobalStore    = NewDefaultDataStore(nil, GlobalStoreName, Always)
	WorkspaceStore = NewDefaultDataStore(GlobalStore, WorkspaceStoreName, Always)
	WorkflowStore  = NewDefaultDataStore(WorkspaceStore, WorkflowStoreName, OnSet)
	JobStore       = NewDefaultDataStore(WorkflowStore, JobStoreName, OnSet)
	ExecutionStore = NewDefaultDataStore(JobStore, ExecutionStoreName, Once)
)
