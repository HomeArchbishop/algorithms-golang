package structure

import (
	"testing"
)

// Mock type that implements Comparable interface
type MockComparable struct {
	id    int
	value int
}

func (m MockComparable) Id() int {
	return m.id
}

func (m MockComparable) LargerThan(other interface{}) bool {
	return m.value > other.(MockComparable).value
}

func TestMaxHeapInit(t *testing.T) {
	tests := []struct {
		name     string
		input    []MockComparable
		expected []MockComparable
	}{
		{
			name: "Non-empty slice",
			input: []MockComparable{
				{id: 1, value: 10},
				{id: 2, value: 20},
				{id: 3, value: 5},
			},
			expected: []MockComparable{
				{id: 2, value: 20},
				{id: 1, value: 10},
				{id: 3, value: 5},
			},
		},
		{
			name:     "Empty slice",
			input:    []MockComparable{},
			expected: []MockComparable{},
		},
		{
			name:     "Nil slice",
			input:    nil,
			expected: []MockComparable{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var heap MaxHeap[MockComparable]
			heap.Init(tt.input)

			if heap.heapSize != len(tt.expected) {
				t.Errorf("Expected heap size %d, got %d", len(tt.expected), heap.heapSize)
			}

			for i, v := range tt.expected {
				if heap.slice[i] != v {
					t.Errorf("Expected heap element %v at index %d, got %v", v, i, heap.slice[i])
				}
			}
		})
	}
}
