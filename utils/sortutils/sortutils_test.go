package sortutils

import (
	"sort"
	"testing"
)

func TestColumnSort(t *testing.T) {
	lines := []string{
		"foo 2",
		"bar 1",
		"baz 3",
		"baz 30",
		"baz 10",
	}
	sort.Slice(lines, Comparator(lines, 1, true, false, false, false))

	want := []string{"bar 1", "foo 2", "baz 3", "baz 10", "baz 30"}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("Column sort failed, got %v, want %v", lines, want)
		}
	}
}

func TestNumericSort(t *testing.T) {
	lines := []string{"10", "2", "30"}
	sort.Slice(lines, Comparator(lines, -1, true, false, false, false))

	want := []string{"2", "10", "30"}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("Numeric sort failed, got %v, want %v", lines, want)
		}
	}
}

func TestNumericReverseSort(t *testing.T) {
	lines := []string{"10", "2", "30"}
	sort.Slice(lines, Comparator(lines, -1, true, true, false, false))

	want := []string{"30", "10", "2"}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("Numeric reverse sort failed, got %v, want %v", lines, want)
		}
	}
}

func TestMonthSort(t *testing.T) {
	lines := []string{"Mar", "Jan", "Feb"}
	sort.Slice(lines, Comparator(lines, -1, false, false, true, false))

	want := []string{"Jan", "Feb", "Mar"}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("Month sort failed, got %v, want %v", lines, want)
		}
	}
}

func TestHumanSort(t *testing.T) {
	lines := []string{"10K", "5M", "500"}
	sort.Slice(lines, Comparator(lines, -1, false, false, false, true))

	want := []string{"500", "10K", "5M"}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("Human-readable sort failed, got %v, want %v", lines, want)
		}
	}
}

func TestUniqueSort(t *testing.T) {
	lines := []string{"apple", "banana", "apple"}
	lines = RemoveDuplicates(lines)
	want := []string{"apple", "banana"}
	if len(lines) != len(want) {
		t.Fatalf("RemoveDuplicates failed: got %v, want %v", lines, want)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("RemoveDuplicates failed: got %v, want %v", lines, want)
		}
	}
}

func TestIsSorted(t *testing.T) {
	lines := []string{"1", "2", "3"}
	cmp := Comparator(lines, -1, true, false, false, false)
	if !IsSorted(lines, cmp) {
		t.Errorf("IsSorted failed: lines should be sorted")
	}
	lines = []string{"2", "1", "3"}
	cmp = Comparator(lines, -1, true, false, false, false)
	if IsSorted(lines, cmp) {
		t.Errorf("IsSorted failed: lines should not be sorted")
	}
}
