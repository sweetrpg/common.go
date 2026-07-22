package util

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMapOmitsNilResults(t *testing.T) {
	in := []int{1, 2, 3, 4}
	out := Map(in, func(v int) *string {
		if v%2 != 0 {
			return nil
		}
		s := strconv.Itoa(v)
		return &s
	})

	want := []string{"2", "4"}
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("Map() = %v, want %v", out, want)
	}
}

func TestMapEmptyInput(t *testing.T) {
	out := Map([]int{}, func(v int) *int { return &v })
	if len(out) != 0 {
		t.Fatalf("Map() on empty input = %v, want empty", out)
	}
}

func TestNullMapPreservesLength(t *testing.T) {
	in := []int{1, 2, 3}
	out := NullMap(in, func(v int) string { return strconv.Itoa(v * 2) })

	want := []string{"2", "4", "6"}
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("NullMap() = %v, want %v", out, want)
	}
}

func TestNullMapEmptyInput(t *testing.T) {
	out := NullMap([]int{}, func(v int) int { return v })
	if len(out) != 0 {
		t.Fatalf("NullMap() on empty input = %v, want empty", out)
	}
}
