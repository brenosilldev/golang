package main

import (
	"testing"
)

func TestPersonStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	tests := []struct {
		name     string
		input    person
		expected person
	}{
		{"Test Alice", person{name: "Alice", age: 30}, person{name: "Alice", age: 30}},
		{"Test Bob", person{name: "Bob", age: 25}, person{name: "Bob", age: 25}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input != tt.expected {
				t.Errorf("got %v, want %v", tt.input, tt.expected)
			}
		})
	}
}
