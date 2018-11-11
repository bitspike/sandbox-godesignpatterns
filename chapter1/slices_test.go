package chapter1

import (
	"testing"
)

// TestSlices1 is a testcase about how slice references behave after appending
// to the original slice
func TestSlices1(t *testing.T) {
	// Initialize the slice to test
	t.Logf("Initialize the slice to test")
	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Logf("slice1: %+v", slice1)
	t.Logf("\n")

	// Declare "new" slice and change one item
	t.Logf("Declare 'new' slice and change one item")
	slice2 := slice1
	slice2[2] = 12
	t.Logf("slice1: %+v", slice1)
	t.Logf("slice2: %+v", slice2)
	t.Logf("\n")

	// Declare "new" subslice and change one item
	t.Logf("Declare 'new' slice and change one item")
	slice2 = slice1[5:]
	slice2[2] = 12
	t.Logf("slice1: %+v", slice1)
	t.Logf("slice2: %+v", slice2)
	t.Logf("\n")

	// Declare empty slice, append old slice and change one item
	t.Logf("Declare empty slice, append old slice and change one item")
	slice2 = append([]int{}, slice1...)
	slice2[2] = 22
	t.Logf("slice1: %+v", slice1)
	t.Logf("slice2: %+v", slice2)
	t.Logf("\n")
}

// // TestClosures2 is a testcase about how closure context(reference) is handled
// func TestClosures2(t *testing.T) {
// 	// Second run of closures
// 	t.Logf("Running add5(3): %d", add5(3))
// 	t.Logf("Running add10(3): %d", add10(3))
// }
