package utils

import (
	"encoding/json"
	"io/ioutil"
)

func LoadConfigJSON(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func SaveConfigJSON(filename string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0666)
}
