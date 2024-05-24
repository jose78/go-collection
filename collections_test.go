package gocollection

import (
	"testing"

	"fmt"
)

// Test functions
// Test functions
func TestMap(t *testing.T) {
	source := []int{1, 2, 3, 4}
	var dest []any
	builder := Map(func(n int) any { return n * 2 }, source, &dest)

	if err := builder.Error(); err != nil {
		t.Fatalf("Map failed: %v", err)
	}

	expected := []any{2, 4, 6, 8}
	for i, v := range dest {
		if v != expected[i] {
			t.Errorf("Map result mismatch at index %d: got %v, want %v", i, v, expected[i])
		}
	}
}

func TestFilter(t *testing.T) {
	source := []int{1, 2, 3, 4}
	var dest []int
	builder := Filter(func(n int) bool { return n%2 == 0 }, source, &dest)

	if err := builder.Error(); err != nil {
		t.Fatalf("Filter failed: %v", err)
	}

	expected := []int{2, 4}
	for i, v := range dest {
		if v != expected[i] {
			t.Errorf("Filter result mismatch at index %d: got %v, want %v", i, v, expected[i])
		}
	}
}

func TestForEach(t *testing.T) {
	source := []int{1, 2, 3, 4}
	var output []string
	builder := ForEach(func(i int, n int) { output = append(output, fmt.Sprintf("Index: %d, Value: %d", i, n)) }, source)

	if err := builder.Error(); err != nil {
		t.Fatalf("ForEach failed: %v", err)
	}

	expected := []string{
		"Index: 0, Value: 1",
		"Index: 1, Value: 2",
		"Index: 2, Value: 3",
		"Index: 3, Value: 4",
	}
	for i, v := range output {
		if v != expected[i] {
			t.Errorf("ForEach result mismatch at index %d: got %v, want %v", i, v, expected[i])
		}
	}
}

func TestZip(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := []int{1, 2, 3}
	result := make(map[string]int)
	builder := Zip(keys, values, result)

	if err := builder.Error(); err != nil {
		t.Fatalf("Zip failed: %v", err)
	}

	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Zip result mismatch for key %v: got %v, want %v", k, result[k], v)
		}
	}
}

func TestWithErrorMessage(t *testing.T) {
	source := []int{1, 2, 3, 4}
	var dest []any

	mapper := func(n int) any {
		if n == 3 {
			panic("unexpected value")
		}
		return n * 2
	}
	builder := Map(mapper, source, &dest)

	customErrFunc := func(item int) error {
		return fmt.Errorf("custom error for item %d", item)
	}

	err := builder.WithErrorMessage(customErrFunc).Error()
	if err == nil || err.Error() != "custom error for item 3" {
		t.Fatalf("expected custom error for item 3, got %v", err)
	}
}

