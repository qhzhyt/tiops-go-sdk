package utils

import (
	"bufio"
	"encoding/json"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

func UnmarshalYAMLFile(obj interface{}, filePath string) error {
	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	data, _ := ioutil.ReadAll(reader)
	y, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(y, obj)
}
