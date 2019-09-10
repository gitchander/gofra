package gofra

import (
	"encoding/json"
	"io/ioutil"
)

func LoadFromJsonFile(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func SaveToJsonFile(filename string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0666)
}
