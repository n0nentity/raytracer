//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"testing"
)

func TestNewSurface(t *testing.T) {
	p := objects.NewPlane(objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0))

	s := NewSurface(p, objects.NewVector(6.0, 7.0, 8.0), 9.0, 10.0, 11.0, 12.0)

	if *s.plane.Position() != *objects.NewVector(0.0, 1.0, 2.0) || *s.plane.Normal() != *objects.NewVector(3.0, 4.0, 5.0) || *s.color != *objects.NewVector(6.0, 7.0, 8.0) || s.diffuse != 9.0 || s.specularIntensity != 10.0 || s.specularPower != 11.0 || s.reflectivity != 12.0 {
		t.Errorf("TestNewSurface %v %v %v %v %v %v %v %v %v", objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0), s.plane.Position(), s.plane.Normal(), s.color, s.diffuse, s.specularIntensity, s.specularPower, s.reflectivity)
	}
}

func TestSurfaceGetSet(t *testing.T) {
	p := objects.NewPlane(objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0))

	s := NewSurface(p, objects.NewVector(6.0, 7.0, 8.0), 9.0, 10.0, 11.0, 12.0)

	if *s.Plane().Position() != *objects.NewVector(0.0, 1.0, 2.0) || *s.Plane().Normal() != *objects.NewVector(3.0, 4.0, 5.0) || *s.Color() != *objects.NewVector(6.0, 7.0, 8.0) || s.Diffuse() != 9.0 || s.SpecularIntensity() != 10.0 || s.SpecularPower() != 11.0 || s.Reflectivity() != 12.0 {
		t.Errorf("TestNewSurface %v %v %v %v %v %v %v %v %v", objects.NewVector(0.0, 1.0, 2.0), objects.NewVector(3.0, 4.0, 5.0), s.plane.Position(), s.plane.Normal(), s.Color, s.Diffuse(), s.SpecularIntensity(), s.SpecularPower(), s.Reflectivity())
	}

	p = objects.NewPlane(objects.NewVector(13.0, 14.0, 15.0), objects.NewVector(16.0, 17.0, 18.0))
	s.SetPlane(p)
	s.SetColor(objects.NewVector(19.0, 20.0, 21.0))
	s.SetDiffuse(22.0)
	s.SetSpecularIntensity(23.0)
	s.SetSpecularPower(24.0)
	s.SetReflectivity(25.0)

	if *s.Plane().Position() != *objects.NewVector(13.0, 14.0, 15.0) || *s.Plane().Normal() != *objects.NewVector(16.0, 17.0, 18.0) || *s.Color() != *objects.NewVector(19.0, 20.0, 21.0) || s.Diffuse() != 22.0 || s.SpecularIntensity() != 23.0 || s.SpecularPower() != 24.0 || s.Reflectivity() != 25.0 {
		t.Errorf("TestNewSurface %v %v %v %v %v %v %v %v %v", objects.NewVector(13.0, 14.0, 15.0), objects.NewVector(16.0, 17.0, 18.0), s.plane.Position(), s.plane.Normal(), s.Color, s.Diffuse(), s.SpecularIntensity(), s.SpecularPower(), s.Reflectivity())
	}
}

func TestSurfaceIntersection(t *testing.T) {
	p := objects.NewPlane(objects.NewVector(0.0, 0.0, 0.0), objects.NewVector(0.0, 1.0, 0.0))

	s := NewSurface(p, objects.NewVector(6.0, 7.0, 8.0), 9.0, 10.0, 11.0, 12.0)

	var intersectP, col, norm, dif, spInt, spPow, ref = s.Intersection(objects.NewRay(objects.NewVector(0, 1, 0), objects.NewVector(0, -1, 0)))
	if *intersectP != *objects.NewVector(0.0, 0.0, 0.0) || *col != *objects.NewVector(0.0, 0.0, 0.0) || *norm != *objects.NewVector(0.0, 1.0, 0.0) || dif != 9.0 || spInt != 10.0 || spPow != 11.0 || ref != 12.0 {
		t.Errorf("TestSurfaceIntersection 1 %v %v %v %v %v %v %v", intersectP, col, norm, dif, spInt, spPow, ref)
	}
	intersectP, col, norm, dif, spInt, spPow, ref = s.Intersection(objects.NewRay(objects.NewVector(1, 1, 0), objects.NewVector(0, -1, 0)))
	if *intersectP != *objects.NewVector(1.0, 0.0, 0.0) || *col != *objects.NewVector(6.0, 7.0, 8.0) || *norm != *objects.NewVector(0.0, 1.0, 0.0) || dif != 9.0 || spInt != 10.0 || spPow != 11.0 || ref != 12.0 {
		t.Errorf("TestSurfaceIntersection 2 %v %v %v %v %v %v %v", intersectP, col, norm, dif, spInt, spPow, ref)
	}

	intersectP, col, norm, dif, spInt, spPow, ref = s.Intersection(objects.NewRay(objects.NewVector(0, -1, 0), objects.NewVector(0, -1, 0)))
	if intersectP != nil {
		t.Errorf("TestSurfaceIntersection 3 %v", intersectP)
	}

}
