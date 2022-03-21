package stores



type basicDataStore struct {
	store     map[string]interface{}
	upstream  DataStore
	name      string
	path      string
	storeMode StoreMode
	apiStore  APIStore
	changed   bool
}

func (c *basicDataStore) LoadAll() bool {
	data, err := c.apiStore.LoadAll(c.path)
	if err != nil {
		return false
	}

	for k, v := range data {
		c.store[k] = v
	}

	return true
}

func (c *basicDataStore) SaveAll() bool {
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

func (c *basicDataStore) SetValueWithName(n, k string, v interface{}) bool {
	if c.name == n {
		c.PutValue(k, v)
		return true
	}
	if c.upstream != nil {
		return c.upstream.SetValueWithName(n, k, v)
	}
	return false
}

func (c *basicDataStore) GetValueWithName(n, k string) interface{} {
	if c.name == n {
		return c.GetValue(k)
	}
	if c.upstream != nil {
		return c.upstream.GetValueWithName(n, k)
	}
	return nil
}

func (c *basicDataStore) loadValue(k string) interface{} {
	value, err := c.apiStore.LoadValue(c.path, k)
	if err != nil {
		return nil
	}
	return value
}

func (c *basicDataStore) saveValue(k string, v interface{}) bool {
	err := c.apiStore.SaveValue(c.path, k, v)
	if err != nil {
		return false
	}
	return true
}

func (c *basicDataStore) LoadValue(k string) interface{} {
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

func (c *basicDataStore) Path() string {
	return c.path
}

func (c *basicDataStore) LookupValue(k string) interface{} {

	if _, ok := c.store[k]; ok {
		return c.GetValue(k)
	}
	if c.upstream != nil {
		return c.upstream.LookupValue(k)
	}
	return nil
}

func (c *basicDataStore) SetValue(k string, v interface{}) bool {
	if _, ok := c.store[k]; ok {
		c.PutValue(k, v)
		return true
	}
	if c.upstream != nil {
		return c.upstream.SetValue(k, v)
	}
	return false
}

func (c *basicDataStore) PutValue(k string, v interface{}) {
	v0 := c.store[k]
	if v0 != v {
		c.changed = true
	}
	c.store[k] = v
	if c.storeMode >= OnSet {
		c.saveValue(k, v)
	}
}

func (c *basicDataStore) GetValue(k string) interface{} {

	if c.storeMode >= Always {
		return c.LoadValue(k)
	}

	v := c.store[k]

	if v == nil && c.storeMode >= Default {
		return c.LoadValue(k)
	}

	return nil

}

func (c *basicDataStore) SaveValue(k string, v interface{}) bool {
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

//func (c *basicDataStore) CanSaveValue() bool {
//	return false
//}

func (c *basicDataStore) Upstream() DataStore {
	return c.upstream
}

func NewBasicDataStore(upstream DataStore, name string, mode StoreMode) BasicDataStore {
	storePath := name
	if upstream != nil {
		storePath = upstream.Path() + "/" + name
	}
	store := &basicDataStore{store: map[string]interface{}{}, upstream: upstream, name: name, path: storePath, storeMode: mode}
	if mode == Once {
		store.LoadAll()
	}
	return store
}



//func DownstreamCommonDataStore(upstream DataStore) DataStore {
//	return &basicDataStore{store: map[string]interface{}{}, upstream: upstream}
//}
