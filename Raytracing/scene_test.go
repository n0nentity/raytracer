package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"testing"
)

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
