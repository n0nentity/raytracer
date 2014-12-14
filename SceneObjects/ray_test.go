//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package objects

import "testing"

func TestNewRay(t *testing.T) {
	r := NewRay(NewVector(0.0, 1.0, 2.0), NewVector(3.0, 4.0, 0.0))

	if *r.origin != *NewVector(0.0, 1.0, 2.0) || *r.direction != *NewVector(0.6, 0.8, 0.0) {
		t.Errorf("TestNewRay %v %v %v %v", NewVector(0.0, 1.0, 2.0), NewVector(0.6, 0.8, 0.0), r.origin, r.direction)
	}
}

func TestRayGetSet(t *testing.T) {
	r := NewRay(NewVector(0.0, 1.0, 2.0), NewVector(3.0, 4.0, 0.0))

	if *r.Origin() != *NewVector(0.0, 1.0, 2.0) || *r.Direction() != *NewVector(0.6, 0.8, 0.0) {
		t.Errorf("TestRayGetSet %v %v %v %v", NewVector(0.0, 1.0, 2.0), NewVector(0.6, 0.8, 0.0), r.Origin(), r.Direction())
	}

	r.SetOrigin(NewVector(3.0, 4.0, 5.0))
	r.SetDirection(NewVector(0.0, 3.0, 4.0))

	if *r.Origin() != *NewVector(3.0, 4.0, 5.0) || *r.Direction() != *NewVector(0.0, 0.6, 0.8) {
		t.Errorf("TestRayGetSet %v %v %v %v", NewVector(3.0, 4.0, 5.0), NewVector(0.0, 0.6, 0.8), r.Origin(), r.Direction())
	}
}

func TestAtStep(t *testing.T) {
	r := NewRay(NewVector(0.0, 1.0, 2.0), NewVector(3.0, 4.0, 0.0))

	v := r.AtStep(4.0)

	if *v != *NewVector(2.4, 4.2, 2) {
		t.Errorf("TestAtStep %v %v", NewVector(2.4, 4.2, 2), v)
	}

	v = r.AtStep(-4.0)

	if *v != *NewVector(-2.4, -2.2, 2) {
		t.Errorf("TestAtStep %v %v", NewVector(-2.4, -2.2, 2), v)
	}
}
