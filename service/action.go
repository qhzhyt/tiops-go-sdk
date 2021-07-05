package service

import (
	"encoding/json"
	"strconv"
)

func (a ActionOptions) GetString(name string) string {
	return a[name]
}

func (a ActionOptions) GetInt(name string) int {
	result, err := strconv.Atoi(a[name])
	if err != nil {
		return 0
	}
	return result
}

func (a ActionOptions) GetStringList(name string) []string {
	result := []string{}

	err := json.Unmarshal([]byte(a[name]), &result)

	if err != nil {
		return nil
	}

	return result
}
