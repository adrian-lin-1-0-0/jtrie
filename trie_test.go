package jtrie

import (
	"reflect"
	"testing"
)

func TestJsonTrie_Set(t *testing.T) {
	tests := []struct {
		name     string
		initial  JsonTrie
		path     []string
		value    interface{}
		expected JsonTrie
	}{
		{
			name:     "set value at root",
			initial:  JsonTrie{},
			path:     []string{"key1"},
			value:    "value1",
			expected: JsonTrie{"key1": "value1"},
		},
		{
			name:     "set nested value",
			initial:  JsonTrie{},
			path:     []string{"key1", "key2"},
			value:    "value2",
			expected: JsonTrie{"key1": map[string]interface{}{"key2": "value2"}},
		},
		{
			name:     "overwrite existing value",
			initial:  JsonTrie{"key1": "oldValue"},
			path:     []string{"key1"},
			value:    "newValue",
			expected: JsonTrie{"key1": "newValue"},
		},
		{
			name:     "set value in existing nested map",
			initial:  JsonTrie{"key1": map[string]interface{}{"key2": "value2"}},
			path:     []string{"key1", "key3"},
			value:    "value3",
			expected: JsonTrie{"key1": map[string]interface{}{"key2": "value2", "key3": "value3"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.Set(tt.value, tt.path...)
			if !reflect.DeepEqual(tt.initial, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, tt.initial)
			}
		})
	}
}
func TestJsonTrie_Get(t *testing.T) {
	tests := []struct {
		name     string
		initial  JsonTrie
		path     []string
		expected interface{}
		found    bool
	}{
		{
			name:     "get value at root",
			initial:  JsonTrie{"key1": "value1"},
			path:     []string{"key1"},
			expected: "value1",
			found:    true,
		},
		{
			name:     "get nested value",
			initial:  JsonTrie{"key1": map[string]interface{}{"key2": "value2"}},
			path:     []string{"key1", "key2"},
			expected: "value2",
			found:    true,
		},
		{
			name:     "get non-existent value",
			initial:  JsonTrie{"key1": "value1"},
			path:     []string{"key2"},
			expected: nil,
			found:    false,
		},
		{
			name:     "get value from non-map type",
			initial:  JsonTrie{"key1": "value1"},
			path:     []string{"key1", "key2"},
			expected: nil,
			found:    false,
		},
		{
			name:     "get value in deeper nested map",
			initial:  JsonTrie{"key1": map[string]interface{}{"key2": map[string]interface{}{"key3": "value3"}}},
			path:     []string{"key1", "key2", "key3"},
			expected: "value3",
			found:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := tt.initial.Get(tt.path...)
			if !reflect.DeepEqual(result, tt.expected) || found != tt.found {
				t.Errorf("%s expected %v, %v, got %v, %v", tt.name, tt.expected, tt.found, result, found)
			}
		})
	}
}
func TestJsonTrie_Delete(t *testing.T) {
	tests := []struct {
		name     string
		initial  JsonTrie
		path     []string
		expected JsonTrie
	}{
		{
			name:     "delete value at root",
			initial:  JsonTrie{"key1": "value1"},
			path:     []string{"key1"},
			expected: JsonTrie{},
		},
		{
			name:     "delete nested value",
			initial:  JsonTrie{"key1": map[string]interface{}{"key2": "value2"}},
			path:     []string{"key1", "key2"},
			expected: JsonTrie{"key1": map[string]interface{}{}},
		},
		{
			name:     "delete non-existent value",
			initial:  JsonTrie{"key1": "value1"},
			path:     []string{"key2"},
			expected: JsonTrie{"key1": "value1"},
		},
		{
			name:     "delete value from non-map type",
			initial:  JsonTrie{"key1": "value1"},
			path:     []string{"key1", "key2"},
			expected: JsonTrie{"key1": "value1"},
		},
		{
			name:     "delete value in deeper nested map",
			initial:  JsonTrie{"key1": map[string]interface{}{"key2": map[string]interface{}{"key3": "value3"}}},
			path:     []string{"key1", "key2", "key3"},
			expected: JsonTrie{"key1": map[string]interface{}{"key2": map[string]interface{}{}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.Delete(tt.path...)
			if !reflect.DeepEqual(tt.initial, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, tt.initial)
			}
		})
	}
}
func TestJsonTrie_ToJSON(t *testing.T) {
	tests := []struct {
		name     string
		initial  JsonTrie
		expected string
	}{
		{
			name:     "empty trie",
			initial:  JsonTrie{},
			expected: `{}`,
		},
		{
			name:     "single key-value pair",
			initial:  JsonTrie{"key1": "value1"},
			expected: `{"key1":"value1"}`,
		},
		{
			name:     "nested key-value pairs",
			initial:  JsonTrie{"key1": map[string]interface{}{"key2": "value2"}},
			expected: `{"key1":{"key2":"value2"}}`,
		},
		{
			name:     "multiple key-value pairs",
			initial:  JsonTrie{"key1": "value1", "key2": "value2"},
			expected: `{"key1":"value1","key2":"value2"}`,
		},
		{
			name:     "complex nested structure",
			initial:  JsonTrie{"key1": map[string]interface{}{"key2": map[string]interface{}{"key3": "value3"}}},
			expected: `{"key1":{"key2":{"key3":"value3"}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.initial.ToJSON()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if string(result) != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, string(result))
			}
		})
	}
}
