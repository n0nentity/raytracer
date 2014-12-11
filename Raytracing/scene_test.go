package scene

import (
	"de/vorlesung/projekt/raytracer/Helper"
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"testing"
)

func TestRender(t *testing.T) {

}

func Builder() *Scene {
	sphere1 := objects.NewSphere(objects.NewVector(0.0, 0.0, 1.0), 1.0)
	sphere2 := objects.NewSphere(objects.NewVector(-2.5, 0.0, -0.75), 1.0)
	plane := objects.NewPlane(objects.NewVector(0.0, -1.0, 0.0), objects.NewVector(0.0, 1.0, 0.0))

	basicBlue := objects.NewVector(0.54, 0.60, 0.90)
	neonYellow := objects.NewVector(1.0, 2.0, 0.0)

	grid1 := NewGrid(objects.NewVector(2.0, 1.00, -1.0), objects.NewVector(2.0, -0.50, 1.0))
	ball1 := NewBall(sphere1, basicBlue, 0.9, 4.0, 30.0, 0.125)
	ball2 := NewBall(sphere2, neonYellow, 0.9, 4.0, 30.0, 0.125)

	surface1 := NewSurface(plane, objects.NewVector(1.0, 1.0, 1.0), 1.0, 1.0, 8.0, 0.05)
	light1 := NewLight(objects.NewVector(1.0, 4.0, 0.5), objects.NewVector(1.0, 1.0, 1.0))

	colorSky := objects.NewVector(0.85, 0.85, 0.95)

	currentScene := NewScene(objects.NewVector(4.0, 0.5, 0.0), grid1)
	currentScene.AddElement(SceneObject(ball1))
	currentScene.AddElement(SceneObject(ball2))
	currentScene.AddElement(SceneObject(surface1))
	currentScene.SetAmbient(objects.NewVector(0.25, 0.25, 0.3))
	currentScene.SetLight(light1)
	currentScene.SetSkyColor(colorSky)

	return currentScene
}

func TestIntersection(t *testing.T) {
	scene := Builder()
	r := objects.NewRay(objects.NewVector(2.0, 0.0, 0.0), objects.NewVector(-3.0, 0.0, 0.0))

	_, p, _, _, _, _, _, _, _ := scene.intersectAll(r, nil)
	x := Helper.Round(p.X(), 5)
	if x != -1.83857 || p.Y() != 0.0 || p.Z() != 0.0 {
		t.Errorf("Intersect expected VEC(1.0, 0.0, 0.0), actual(%f, %f, %f", x, p.Y(), p.Z())
	}
}

func TestRaytracing(t *testing.T) {
	scene := Builder()
	color, position := scene.raytracing(objects.NewVector(4.0, 0.0, 0.0), objects.NewVector(-1.0, 0.0, 0.0), nil, 8)
	r := Helper.Round(color.X(), 4)
	g := Helper.Round(color.Y(), 4)
	b := Helper.Round(color.Z(), 4)
	//if color.X() != 0.203724 || color.Y() != 0.253906 || color.Z() != 0.251094 {
	if r != 0.2037 || g != 0.2539 || b != 0.2511 {
		t.Errorf("Color expected VEC(0.203724, 0.253906, 0.251094) actual(%f, %f, %f)", r, g, b)
	}
	if position.X() != 0.0 || position.Y() != 0.0 || position.Z() != 0.0 {
		t.Errorf("Position expected VEC(0.0, 0.0, 0.0) actual(%f, %f, %f)", position.X(), position.Y(), position.Z())
	}
}

func TestNewScene(t *testing.T) {
	defer func() {
		e := recover()
		if e == nil {
			t.Errorf("TestNewScene error is nil")
		}
	}()
	g := NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(1.0, 5.0, 6.0))
	s := NewScene(objects.NewVector(7.0, 8.0, 9.0), g)

	if *s.View() != *objects.NewVector(7.0, 8.0, 9.0) ||
		*s.Grid().TopLeft() != *objects.NewVector(1.0, 2.0, 3.0) ||
		*s.Grid().BottomRight() != *objects.NewVector(1.0, 5.0, 6.0) ||
		s.Elements() == nil ||
		*s.Ambient() != *objects.NewVector(0.0, 0.0, 0.0) ||
		*s.SkyColor() != *objects.NewVector(0.0, 0.0, 0.0) ||
		s.Light() == nil {
		t.Errorf("TestNewScene %v %v %v %v %v %v %v", s.View(), s.Grid().TopLeft(),
			s.Grid().BottomRight(), s.Elements(),
			s.Ambient(), s.SkyColor(), s.Light())
	}

	g = NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(2.0, 5.0, 6.0))
	s = NewScene(objects.NewVector(7.0, 8.0, 9.0), g)
}

func TestSceneGetSet(t *testing.T) {
	g := NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(1.0, 5.0, 6.0))
	s := NewScene(objects.NewVector(7.0, 8.0, 9.0), g)

	if *s.View() != *objects.NewVector(7.0, 8.0, 9.0) ||
		*s.Grid().TopLeft() != *objects.NewVector(1.0, 2.0, 3.0) ||
		*s.Grid().BottomRight() != *objects.NewVector(1.0, 5.0, 6.0) ||
		s.Elements() == nil || *s.Ambient() != *objects.NewVector(0.0, 0.0, 0.0) ||
		*s.SkyColor() != *objects.NewVector(0.0, 0.0, 0.0) ||
		s.Light() == nil {
		t.Errorf("TestSceneGetSet 1 %v %v %v %v %v %v %v", s.View(), s.Grid().TopLeft(),
			s.Grid().BottomRight(),
			s.Elements(), s.Ambient(),
			s.SkyColor(), s.Light())
	}

	g = NewGrid(objects.NewVector(10.0, 11.0, 12.0), objects.NewVector(10.0, 14.0, 15.0))
	s.SetGrid(g)
	s.SetView(objects.NewVector(16.0, 17.0, 18.0))
	//objects := make([]SceneObject)
	//objects = append(objects, element)
	s.SetAmbient(objects.NewVector(19.0, 20.0, 21.0))
	s.SetSkyColor(objects.NewVector(22.0, 23.0, 24.0))

	if *s.View() != *objects.NewVector(16.0, 17.0, 18.0) ||
		*s.Grid().TopLeft() != *objects.NewVector(10.0, 11.0, 12.0) ||
		*s.Grid().BottomRight() != *objects.NewVector(10.0, 14.0, 15.0) ||
		s.Elements() == nil ||
		*s.Ambient() != *objects.NewVector(19.0, 20.0, 21.0) ||
		*s.SkyColor() != *objects.NewVector(22.0, 23.0, 24.0) {
		t.Errorf("TestSceneGetSet 2 %v %v %v %v %v %v %v", s.View(), s.Grid().TopLeft(),
			s.Grid().BottomRight(), s.Elements(),
			s.Ambient(), s.SkyColor(),
			s.Light())
	}
}

func TestSceneScale(t *testing.T) {
	x, y := 100, 100
	tmp_in := make([][]objects.Vector, x)
	for i := 0; i < x; i++ {
		tmp_in[i] = make([]objects.Vector, y)
		for j := 0; j < y; j++ {
			v := objects.NewVector(float64(i+1), float64(j+1), 0)
			tmp_in[i][j] = *v
		}
	}

	bRes := false
	res := scale(tmp_in, x, y, 1)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			r, g, b, a := res.At(i, j).RGBA()
			if r != uint32(256*(256-i)-1-i) || g != uint32(256*(256-j)-1-j) || b != 0 || a != uint32(256*256-1) {
				bRes = true
			}
		}
	}
	if bRes {
		t.Errorf("TestSceneScale 1 %v", res)
	}

	res = scale(tmp_in, x, y, 50)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			r, g, b, a := res.At(0, 0).RGBA()
			if r != 26214 || g != 26214 || b != 0 || a != uint32(256*256-1) {
				t.Errorf("TestSceneScale 2 -1 %v %v %v %v", r, 26214, g, 26214)
			}
			r, g, b, a = res.At(0, 1).RGBA()
			if r != 26214 || g != 13364 || b != 0 || a != uint32(256*256-1) {
				t.Errorf("TestSceneScale 2 -2 %v %v %v %v", r, 26214, g, 13364)
			}
			r, g, b, a = res.At(1, 0).RGBA()
			if r != 13364 || g != 26214 || b != 0 || a != uint32(256*256-1) {
				t.Errorf("TestSceneScale 2 -3 %v %v %v %v", r, 13364, g, 26214)
			}
			r, g, b, a = res.At(1, 1).RGBA()
			if r != 13364 || g != 13364 || b != 0 || a != uint32(256*256-1) {
				t.Errorf("TestSceneScale 2 -4 %v %v %v %v", r, 13364, g, 13364)
			}
		}
	}
}
