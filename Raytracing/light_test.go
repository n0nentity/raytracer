//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357
package Raytracing

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"testing"
)

func TestNewLight(t *testing.T) {
	l := NewLight(objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0))

	if *l.position != *objects.NewVector(0.0, 1.0, 2.0) || *l.color != *objects.NewVector(3.0, 4.0, 5.0) {
		t.Errorf("TestNewLight %v %v %v %v", objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0), l.position, l.color)
	}
}

func TestLightGetSet(t *testing.T) {
	l := NewLight(objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0))

	if *l.Position() != *objects.NewVector(0.0, 1.0, 2.0) || *l.Color() != *objects.NewVector(3.0, 4.0, 5.0) {
		t.Errorf("TestLightGetSet %v %v %v %v", objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0), l.Position(), l.Color())
	}

	l.SetPosition(objects.NewVector(6.0, 7.0, 8.0))
	l.SetColor(objects.NewVector(9.0, 10.0, 11.0))

	if *l.Position() != *objects.NewVector(6.0, 7.0, 8.0) || *l.Color() != *objects.NewVector(9.0, 10.0, 11.0) {
		t.Errorf("TestLightGetSet %v %v %v %v", objects.NewVector(6.0, 7.0, 8.0), objects.NewVector(9.0, 10.0, 11.0), l.Position(), l.Color())
	}
}
