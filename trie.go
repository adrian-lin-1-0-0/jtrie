package jtrie

import "encoding/json"

type JsonTrie map[string]interface{}

func (t JsonTrie) Get(path ...string) (interface{}, bool) {
	var obj interface{} = t
	for _, key := range path {
		switch v := obj.(type) {
		case JsonTrie:
			obj = v[key]
		case map[string]interface{}:
			obj = v[key]
		default:
			return nil, false
		}
	}
	return obj, true
}

func (t JsonTrie) Set(value interface{}, path ...string) {
	obj := t
	for i, key := range path {
		if i == len(path)-1 {
			obj[key] = value
			return
		}
		switch v := obj[key].(type) {
		case JsonTrie:
			obj = v
		case map[string]interface{}:
			obj = v
		default:
			obj[key] = map[string]interface{}{}
			obj = obj[key].(map[string]interface{})
		}
	}
}

func (t JsonTrie) Delete(path ...string) {
	obj := t
	for i, key := range path {
		if i == len(path)-1 {
			delete(obj, key)
			return
		}
		switch v := obj[key].(type) {
		case map[string]interface{}:
			obj = v
		case JsonTrie:
			obj = v
		default:
			return
		}
	}
}

func (t JsonTrie) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(t))
}
