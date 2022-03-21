package stores

var (
	GlobalStore    = NewCommonDataStore(nil, GlobalStoreName, Always)
	WorkspaceStore = NewCommonDataStore(GlobalStore, WorkspaceStoreName, Always)
	WorkflowStore  = NewCommonDataStore(WorkspaceStore, WorkflowStoreName, OnSet)
	JobStore       = NewCommonDataStore(WorkflowStore, JobStoreName, OnSet)
	ExecutionStore = NewCommonDataStore(JobStore, ExecutionStoreName, Once)
)

type commonDataStore struct {
	store     map[string]interface{}
	upstream  DataStore
	name      string
	path      string
	storeMode StoreMode
	apiStore  APIStore
	changed   bool
}

func (c *commonDataStore) LoadAll() bool {
	data, err := c.apiStore.LoadAll(c.path)
	if err != nil {
		return false
	}

	for k, v := range data {
		c.store[k] = v
	}

	return true
}

func (c *commonDataStore) SaveAll() bool {
	if !c.changed {
		return true
	}
	err := c.apiStore.SaveAll(c.path, c.store)
	if err != nil {
		return false
	}
	c.changed = false
	return true
}

func (c *commonDataStore) SetValueWithName(n, k string, v interface{}) bool {
	if c.name == n {
		c.PutValue(k, v)
		return true
	}
	if c.upstream != nil {
		return c.upstream.SetValueWithName(n, k, v)
	}
	return false
}

func (c *commonDataStore) GetValueWithName(n, k string) interface{} {
	if c.name == n {
		return c.GetValue(k)
	}
	if c.upstream != nil {
		return c.upstream.GetValueWithName(n, k)
	}
	return nil
}

func (c *commonDataStore) loadValue(k string) interface{} {
	value, err := c.apiStore.LoadValue(c.path, k)
	if err != nil {
		return nil
	}
	return value
}

func (c *commonDataStore) saveValue(k string, v interface{}) bool {
	err := c.apiStore.SaveValue(c.path, k, v)
	if err != nil {
		return false
	}
	return true
}

func (c *commonDataStore) LoadValue(k string) interface{} {
	if c.storeMode > Once {
		v := c.loadValue(k)
		if v != nil {
			c.PutValue(k, v)
		}
	}
	if v, ok := c.store[k]; ok {
		return v
	}
	if c.upstream != nil {
		return c.upstream.LoadValue(k)
	}
	return nil
}

func (c *commonDataStore) Path() string {
	return c.path
}

func (c *commonDataStore) LookupValue(k string) interface{} {

	if _, ok := c.store[k]; ok {
		return c.GetValue(k)
	}
	if c.upstream != nil {
		return c.upstream.LookupValue(k)
	}
	return nil
}

func (c *commonDataStore) SetValue(k string, v interface{}) bool {
	if _, ok := c.store[k]; ok {
		c.PutValue(k, v)
		return true
	}
	if c.upstream != nil {
		return c.upstream.SetValue(k, v)
	}
	return false
}

func (c *commonDataStore) PutValue(k string, v interface{}) {
	v0 := c.store[k]
	if v0 != v {
		c.changed = true
	}
	c.store[k] = v
	if c.storeMode >= OnSet {
		c.saveValue(k, v)
	}
}

func (c *commonDataStore) GetValue(k string) interface{} {

	if c.storeMode >= Always {
		return c.LoadValue(k)
	}

	v := c.store[k]

	if v == nil && c.storeMode >= Default {
		return c.LoadValue(k)
	}

	return nil

}

func (c *commonDataStore) SaveValue(k string, v interface{}) bool {
	if _, ok := c.store[k]; ok {
		c.PutValue(k, v)
	} else {
		if c.upstream == nil || !c.upstream.SaveValue(k, v) {
			c.PutValue(k, v)
		}
	}
	if v, ok := c.store[k]; ok {
		if c.storeMode > Never {
			return c.saveValue(k, v)
		}
	}
	return false
}

//func (c *commonDataStore) CanSaveValue() bool {
//	return false
//}

func (c *commonDataStore) Upstream() DataStore {
	return c.upstream
}

func NewCommonDataStore(upstream DataStore, name string, mode StoreMode) DataStore {
	storePath := name
	if upstream != nil {
		storePath = upstream.Path() + "/" + name
	}
	store := &commonDataStore{store: map[string]interface{}{}, upstream: upstream, name: name, path: storePath, storeMode: mode}
	if mode == Once {
		store.LoadAll()
	}
	return store
}

func NewNodeStore(nodeId string) DataStore {
	return NewCommonDataStore(ExecutionStore, nodeId, Never)
}

//func DownstreamCommonDataStore(upstream DataStore) DataStore {
//	return &commonDataStore{store: map[string]interface{}{}, upstream: upstream}
//}
