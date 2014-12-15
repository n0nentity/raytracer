//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package SceneObjects

import "testing"

func TestNewPlane(t *testing.T) {
	p := NewPlane(NewVector(1.0, 2.0, 3.0), NewVector(4.0, 5.0, 6.0))

	if *p.pos != *NewVector(1.0, 2.0, 3.0) || *p.normal != *NewVector(4.0, 5.0, 6.0) {
		t.Errorf("TestNewRay %v %v %v %v", NewVector(1.0, 2.0, 3.0), NewVector(4.0, 5.0, 6.0), p.pos, p.normal)
	}
}

func TestPlaneGetSet(t *testing.T) {
	p := NewPlane(NewVector(1.0, 2.0, 3.0), NewVector(4.0, 5.0, 6.0))
	if *p.Position() != *NewVector(1.0, 2.0, 3.0) || *p.Normal() != *NewVector(4.0, 5.0, 6.0) {
		t.Errorf("TestNewRay %v %v %v %v", NewVector(1.0, 2.0, 3.0), NewVector(4.0, 5.0, 6.0), p.Position(), p.Normal())
	}

	p.SetPosition(NewVector(7.0, 8.0, 9.0))
	p.SetNormal(NewVector(10.0, 11.0, 12.0))
	if *p.Position() != *NewVector(7.0, 8.0, 9.0) || *p.Normal() != *NewVector(10.0, 11.0, 12.0) {
		t.Errorf("TestNewRay %v %v %v %v", NewVector(7.0, 8.0, 9.0), NewVector(10.0, 11.0, 12.0), p.Position(), p.Normal())
	}
}

func TestPlaneFindIntersection(t *testing.T) {
	p := NewPlane(NewVector(0, 0, 0), NewVector(0, 1, 0))

	pos := p.Intersection(NewRay(NewVector(0, 0, 0), NewVector(0, -1, 0)))

	if pos != nil {
		t.Errorf("TestPlaneFindIntersection 1 %v %v", NewVector(3.0, 0.0, 0.0), pos)
	}

	pos = p.Intersection(NewRay(NewVector(0, 1, 0), NewVector(0, -1, 0)))

	if *pos != *NewVector(0.0, 0.0, 0.0) {
		t.Errorf("TestPlaneFindIntersection 2 %v %v", NewVector(0.0, 0.0, 0.0), pos)
	}

	pos = p.Intersection(NewRay(NewVector(0, -1, 0), NewVector(0, -1, 0)))

	if pos != nil {
		t.Errorf("TestPlaneFindIntersection 3 %v", pos)
	}
}
