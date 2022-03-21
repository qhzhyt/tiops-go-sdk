package stores

import (
	"context"
	"encoding/json"
	apiClient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/services"
)

type APIStore struct {
	apiClient *apiClient.APIClient
}

func (a *APIStore) LoadValue(p, k string) (interface{}, error) {
	res, err := a.apiClient.LoadValue(context.TODO(), &services.LoadValueRequest{Path: p, Key: k})
	if err != nil {
		return nil, err
	}
	var value interface{}
	err = json.Unmarshal(res.Value, &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (a *APIStore) SaveValue(p, k string, v interface{}) error {
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = a.apiClient.SaveValue(context.TODO(), &services.SaveValueRequest{
		Path:  p,
		Key:   k,
		Value: value,
	})
	return err
}

func (a *APIStore) LoadAll(p string) (map[string]interface{}, error) {
	allValue, err := a.apiClient.LoadAllValue(context.TODO(), &services.LoadAllValueRequest{Path: p})
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	if allValue.Data != nil {

		for k, v := range allValue.Data {
			var value interface{}
			err = json.Unmarshal(v, &value)
			if err != nil {
				result[k] = v
			} else {
				result[k] = value
			}
		}
	}

	return result, nil
}

func (a *APIStore) SaveAll(p string, d map[string]interface{}) error {
	data := map[string][]byte{}

	for k, v := range d {
		value, err := json.Marshal(v)
		if err == nil {
			data[k] = value
		}
	}

	_, err := a.apiClient.SaveAllValue(context.TODO(), &services.SaveAllValueRequest{Path: p, Data: data})

	return err
}

func NewAPIStore() PersistentStore {
	return &APIStore{apiClient: apiClient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)}
}
