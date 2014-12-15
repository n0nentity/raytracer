//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package Raytracing

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"testing"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid(objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0))

	if *g.top_left != *objects.NewVector(0.0, 1.0, 2.0) || *g.bottom_right != *objects.NewVector(3.0, 4.0, 5.0) {
		t.Errorf("TestNewGrid %v %v %v %v", objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0), g.top_left, g.bottom_right)
	}
}

func TestGridGetSet(t *testing.T) {
	g := NewGrid(objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0))

	if *g.TopLeft() != *objects.NewVector(0.0, 1.0, 2.0) || *g.BottomRight() != *objects.NewVector(3.0, 4.0, 5.0) {
		t.Errorf("TestNewGrid %v %v %v %v", objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0), g.TopLeft(), g.BottomRight())
	}

	g.SetTopLeft(objects.NewVector(6.0, 7.0, 8.0))
	g.SetBottomRight(objects.NewVector(9.0, 10.0, 11.0))

	if *g.TopLeft() != *objects.NewVector(6.0, 7.0, 8.0) || *g.BottomRight() != *objects.NewVector(9.0, 10.0, 11.0) {
		t.Errorf("TestNewGrid %v %v %v %v", objects.NewVector(6.0, 7.0, 8.0), objects.NewVector(9.0, 10.0, 11.0), g.TopLeft(), g.BottomRight())
	}
}
