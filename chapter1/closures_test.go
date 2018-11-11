package chapter1

import (
	"testing"
)

// TestClosures1 is a testcase about how closure context(primitive) is handled
func TestClosures1(t *testing.T) {
	// Define the closure generator func
	addM := func(m int) func(int) int {
		// Check when this context is released
		defer func() {
			t.Logf("Generator function context has been released")
		}()

		// Return the closure func
		return func(n int) int { return n + m }
	}

	// Generate closure functions
	add5 := addM(5)
	add10 := addM(10)

	// Use closures
	t.Logf("Running add5(3): %d", add5(3))
	t.Logf("Running add10(3): %d", add10(3))
	t.Logf("Running add5(3): %d", add5(3))
	t.Logf("Running add10(3): %d", add10(3))
}

// TestClosures2 is a testcase about how closure context(reference) is handled
func TestClosures2(t *testing.T) {
	// Define struct to pass as context var to closure
	type Params struct {
		operant int
	}

	// Define the closure generator func
	addM := func(m *Params) func(int) int {
		// Check when this context is released
		defer func() {
			t.Logf("Generator function context has been released")
		}()

		// Return the closure func
		return func(n int) int { return n + m.operant }
	}

	// Create Params objects and generate the closures
	confA := Params{5}
	confB := Params{10}
	add5 := addM(&confA)
	add10 := addM(&confB)

	// First run of closures
	t.Logf("Running add5(3): %d", add5(3))
	t.Logf("Running add10(3): %d", add10(3))

	// Modification of the original param structs
	confA.operant = 25
	confB.operant = 100

	// Second run of closures
	t.Logf("Running add5(3): %d", add5(3))
	t.Logf("Running add10(3): %d", add10(3))
}
