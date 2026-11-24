package main

import "testing"

func TestIntersection(t *testing.T) {
	shared := &Node{Val: 8, Next: &Node{Val: 10}}
	a := &Node{Val: 3, Next: shared}
	b := &Node{Val: 99, Next: &Node{Val: 1, Next: shared}}
	if got := Intersection(a, b); got != shared {
		t.Errorf("expected the shared node, got %v", got)
	}

	// No intersection.
	c := &Node{Val: 1}
	d := &Node{Val: 1} // same value, different node
	if got := Intersection(c, d); got != nil {
		t.Errorf("expected nil for disjoint lists, got %v", got)
	}

	// One list empty.
	if got := Intersection(nil, a); got != nil {
		t.Errorf("expected nil with an empty list, got %v", got)
	}

	// Identical lists intersect at the head.
	if got := Intersection(a, a); got != a {
		t.Errorf("expected head for identical lists")
	}
}
