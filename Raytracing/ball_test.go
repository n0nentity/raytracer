//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"testing"
)

func TestNewBall(t *testing.T) {
	s := objects.NewSphere(objects.NewVector(1.0, 2.0, 3.0), 4.0)

	b := NewBall(s, objects.NewVector(5.0, 6.0, 7.0), 8.0, 9.0, 10.0, 11.0)

	if *b.sphere.Position() != *objects.NewVector(1.0, 2.0, 3.0) || b.sphere.Radius() != 4.0 || *b.color != *objects.NewVector(5.0, 6.0, 7.0) || b.diffuse != 8.0 || b.specularIntensity != 9.0 || b.specularPower != 10.0 || b.reflectivity != 11.0 {
		t.Errorf("TestNewSurface %v %v %v %v %v %v %v", b.sphere.Position(), b.sphere.Radius(), b.color, b.diffuse, b.specularIntensity, b.specularPower, b.reflectivity)
	}
}

func TestBallGetSet(t *testing.T) {
	s := objects.NewSphere(objects.NewVector(1.0, 2.0, 3.0), 4.0)

	b := NewBall(s, objects.NewVector(5.0, 6.0, 7.0), 8.0, 9.0, 10.0, 11.0)

	if *b.Sphere().Position() != *objects.NewVector(1.0, 2.0, 3.0) || b.Sphere().Radius() != 4.0 || *b.Color() != *objects.NewVector(5.0, 6.0, 7.0) || b.Diffuse() != 8.0 || b.SpecularIntensity() != 9.0 || b.SpecularPower() != 10.0 || b.Reflectivity() != 11.0 {
		t.Errorf("TestBallGetSet 1 %v %v %v %v %v %v %v", b.Sphere().Position(), b.Sphere().Radius(), b.Color(), b.Diffuse(), b.SpecularIntensity(), b.SpecularPower(), b.Reflectivity())
	}

	s = objects.NewSphere(objects.NewVector(12.0, 13.0, 14.0), 15.0)
	b.SetSphere(s)
	b.SetColor(objects.NewVector(16.0, 17.0, 18.0))
	b.SetDiffuse(19.0)
	b.SetSpecularIntensity(20.0)
	b.SetSpecularPower(21.0)
	b.SetReflectivity(22.0)

	if *b.Sphere().Position() != *objects.NewVector(12.0, 13.0, 14.0) || b.Sphere().Radius() != 15.0 || *b.Color() != *objects.NewVector(16.0, 17.0, 18.0) || b.Diffuse() != 19.0 || b.SpecularIntensity() != 20.0 || b.SpecularPower() != 21.0 || b.Reflectivity() != 22.0 {
		t.Errorf("TestBallGetSet 2 %v %v %v %v %v %v %v", b.Sphere().Position(), b.Sphere().Radius(), b.Color(), b.Diffuse(), b.SpecularIntensity(), b.SpecularPower(), b.Reflectivity())
	}
}

func TestBallIntersection(t *testing.T) {
	s := objects.NewSphere(objects.NewVector(0.0, 0.0, 0.0), 4.0)
	b := NewBall(s, objects.NewVector(5.0, 6.0, 7.0), 8.0, 9.0, 10.0, 11.0)

	var intersectP, col, norm, dif, spInt, spPow, ref = b.Intersection(objects.NewRay(objects.NewVector(3.0, 0.0, 0.0), objects.NewVector(0.0, 0.0, 0.0)))
	if *intersectP != *objects.NewVector(3.0, 0.0, 0.0) || *col != *objects.NewVector(5.0, 6.0, 7.0) || *norm != *objects.NewVector(3.0, 0.0, 0.0) || dif != 8.0 || spInt != 9.0 || spPow != 10.0 || ref != 11.0 {
		t.Errorf("TestBallIntersection 1 %v %v %v %v %v %v %v", intersectP, col, norm, dif, spInt, spPow, ref)
	}

	intersectP, col, norm, dif, spInt, spPow, ref = b.Intersection(objects.NewRay(objects.NewVector(5.0, 0.0, 0.0), objects.NewVector(0.0, 1.0, 0.0)))
	if intersectP != nil {
		t.Errorf("TestBallIntersection 2 %v", intersectP)
	}
}
