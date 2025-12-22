package utils

import "encoding/json"

func Struct2Map(obj interface{}) map[string]string {
	j, _ := json.Marshal(obj)
	var m map[string]string
	json.Unmarshal(j, &m)
	return m
}
