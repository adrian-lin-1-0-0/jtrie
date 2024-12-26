package jtrie

import (
	"encoding/json"
	"fmt"
)

type JsonTrie map[string]interface{}

func (t JsonTrie) Get(keys ...string) (interface{}, bool) {
	var obj interface{} = map[string]interface{}(t)
	for _, key := range keys {
		switch v := obj.(type) {
		case map[string]interface{}:
			obj = v[key]
		default:
			return nil, false
		}
	}

	if obj == nil {
		return nil, false
	}

	return obj, true
}

func (t JsonTrie) Set(value interface{}, keys ...string) {
	current := t
	for index, key := range keys {
		if index == len(keys)-1 {
			current[key] = value
			return
		}
		switch child := current[key].(type) {
		case map[string]interface{}:
			current = child
		default:
			current[key] = map[string]interface{}{}
			current = current[key].(map[string]interface{})
		}
	}
}

func (t JsonTrie) Delete(path ...string) error {
	current := t
	for index, key := range path {
		if index == len(path)-1 {
			if _, exists := current[key]; !exists {
				return fmt.Errorf("path does not exist")
			}
			delete(current, key)
			return nil
		}
		switch v := current[key].(type) {
		case map[string]interface{}:
			current = v
		default:
			return fmt.Errorf("path does not exist")
		}
	}
	return nil
}

func (t JsonTrie) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}
