//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package SceneObjects

import "testing"

func TestNewSphere(t *testing.T) {
	x, y, z := 4.0, 2.0, 8.0
	vec := NewVector(x, y, z)

	r := 4.0

	s := NewSphere(vec, r)

	if *s.position != *vec || s.radius != r {
		t.Errorf("TestNewSphere %v %v %v %v", vec, r, s.position, s.radius)
	}
}

func TestSphereGetSet(t *testing.T) {
	vec := NewVector(1.0, 2.0, 3.0)
	r := 4.0
	s := NewSphere(vec, r)

	if *s.Position() != *vec || s.Radius() != r {
		t.Errorf("TestSphereGetSet %v %v %v %v", vec, r, s.Position(), s.Radius())
	}

	vec = NewVector(5.0, 6.0, 7.0)
	r = 8.0
	s.SetPosition(vec)
	s.SetRadius(r)
	if *s.Position() != *vec || s.Radius() != r {
		t.Errorf("TestSphereGetSet %v %v %v %v", vec, r, s.Position(), s.Radius())
	}
}

func TestSphereIntersection(t *testing.T) {
	s := NewSphere(NewVector(0.0, 0.0, 0.0), 1.0)
	r := NewRay(NewVector(3.0, 0.0, 0.0), NewVector(-1.0, 0.0, 0.0))
	pos := s.Intersection(r)
	if *pos != *NewVector(1.0, 0.0, 0.0) {
		t.Errorf("TestSphereIntersection 1 %v %v", NewVector(1.0, 0.0, 0.0), pos)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 4.0)
	r = NewRay(NewVector(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0))
	pos = s.Intersection(r)

	if *pos != *NewVector(0.0, 0.0, 0.0) {
		t.Errorf("TestSphereIntersection 2 %v %v", NewVector(0.0, 0.0, 0.0), pos)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 4.0)
	r = NewRay(NewVector(4.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0))
	pos = s.Intersection(r)

	if *pos != *NewVector(4.0, 0.0, 0.0) {
		t.Errorf("TestSphereIntersection 3 %v %v", NewVector(4.0, 0.0, 0.0), pos)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 4.0)
	r = NewRay(NewVector(3.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0))
	pos = s.Intersection(r)

	if *pos != *NewVector(3.0, 0.0, 0.0) {
		t.Errorf("TestSphereIntersection 4 %v %v", NewVector(3.0, 0.0, 0.0), pos)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 4.0)
	r = NewRay(NewVector(5.0, 0.0, 0.0), NewVector(0.0, 1.0, 0.0))
	pos = s.Intersection(r)

	if pos != nil {
		t.Errorf("TestSphereIntersection 5 %v", pos)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 0.1)
	r = NewRay(NewVector(10.0, 10.0, 10.0), NewVector(10.0, 10.0, 10.0))
	pos = s.Intersection(r)

	if pos != nil {
		t.Errorf("TestSphereIntersection 6 %v", pos)
	}
}
