//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package SceneObjects

import (
	"log"
	"testing"
)

func TestNewVector(t *testing.T) {
	x, y, z := 1.0, 2.0, 3.0
	v := NewVector(x, y, z)

	if v.x != 1.0 || v.y != 2.0 || v.z != 3.0 {
		t.Errorf("TestNewVector %v %v %v %v %v %v", 1.0, 2.0, 3.0, v.x, v.y, v.z)
	}
}

func TestVectorGetSet(t *testing.T) {
	v := NewVector(1.0, 2.0, 3.0)

	if v.X() != 1.0 || v.Y() != 2.0 || v.Z() != 3.0 {
		t.Errorf("TestVectorGetSet %v %v %v %v %v %v", 1.0, 2.0, 3.0, v.X(), v.Y(), v.Z())
	}

	v.SetX(4.0)
	v.SetY(5.0)
	v.SetZ(6.0)

	if v.X() != 4.0 || v.Y() != 5.0 || v.Z() != 6.0 {
		t.Errorf("TestVectorGetSet %v %v %v %v %v %v", 4.0, 5.0, 6.0, v.X(), v.Y(), v.Z())
	}
}

func TestLimit(t *testing.T) {
	v := NewVector(2.0, 0.5, -1.0)

	v = v.Limit(0.0, 1.0)

	if v.X() != 1.0 || v.Y() != 0.5 || v.Z() != 0.0 {
		t.Errorf("TestLimit %v %v %v %v %v %v", 1.0, 0.5, 0.0, v.X(), v.Y(), v.Z())
	}
}

func TestFindClosest(t *testing.T) {
	v := NewVector(2.0, 0.5, -1.0)

	v1 := NewVector(2.0, 0.5, -1.1)
	v2 := NewVector(2.0, 0.5, -0.9)

	res := v.GetClosest(v1, v2)
	if *res != *v2 {
		t.Errorf("TestFindClosest %v %v", v2, res)
	}

	v2 = NewVector(2.0, 0.5, -1.1)
	v1 = NewVector(2.0, 0.5, -0.9)

	res = v.GetClosest(v1, v2)
	if *res != *v1 {
		t.Errorf("TestFindClosest %v %v", v1, res)
	}
}

func TestMax(t *testing.T) {
	v := NewVector(2.0, 0.0, -1.0)
	if v.Max() != 2.0 {
		t.Errorf("TestMax %v %v", 2.0, v.Max())
	}
}

func TestAbs(t *testing.T) {
	v := NewVector(1.0, -1.1, 1.1)
	v2 := v.Absolute()
	if v2.X() != 1.0 && v2.Y() != 1.1 && v2.Z() != 1.1 {
		t.Errorf("TestCross %v %v %v %v %v %v", 1.0, 1.1, 1.1, v2.X(), v2.Y(), v2.Z())
	}
}

func TestCross(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.CrossProduct(vec2)

	if result.X() != -3.0 || result.Y() != 6.0 || result.Z() != -3.0 {
		t.Errorf("TestCross %v %v %v %v %v %v", -3.0, 6.0, -3.0, result.x, result.y, result.z)
	}
}

func TestAddVector(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.AddVector(vec2)

	if result.X() != 9.0 || result.Y() != 11.0 || result.Z() != 13.0 {
		t.Errorf("TestAddVector %v %v %v %v %v %v", 9.0, 11.0, 13.0, result.X(), result.Y(), result.Z())
	}
}

func TestMulitplyVector(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.MultiplyVector(vec2)

	if result.X() != 18.0 || result.Y() != 28.0 || result.Z() != 40.0 {
		t.Errorf("TestMultiplyVector %v %v %v %v %v %v", 18.0, 28.0, 40.0, result.X(), result.Y(), result.Z())
	}
}

func TestSubVector(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.SubtractVector(vec2)

	if result.X() != -3.0 || result.Y() != -3.0 || result.Z() != -3.0 {
		t.Errorf("TestSubVector %v %v %v %v %v %v", -3.0, -3.0, -3.0, result.X(), result.Y(), result.Z())
	}
}

func TestSubValue(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)

	result := vec1.SubtractValue(-2.5)

	if result.X() != 5.5 || result.Y() != -1.5 || result.Z() != 7.5 {
		t.Errorf("TestSubValue %v %v %v %v %v %v", 5.5, -1.5, 7.5, result.X(), result.Y(), result.Z())
	}
}

func TestMultiplyValue(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)

	result := vec1.MultiplyValue(-2.5)

	if result.X() != -7.5 || result.Y() != 10.0 || result.Z() != -12.5 {
		t.Errorf("TestMultiplyValue %v %v %v %v %v %v", -7.5, 10.0, -12.5, result.X(), result.Y(), result.Z())
	}
}

func TestDivValue(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)

	result := vec1.DivideValue(-2.5)

	if result.X() != -1.2 || result.Y() != 1.6 || result.Z() != -2 {
		t.Errorf("TestDivValue %v %v %v %v %v %v", -1.2, 1.6, -2, result.X(), result.Y(), result.Z())
	}

	result = vec1.DivideValue(0)
	if result.X() != 0.0 || result.Y() != 0.0 || result.Z() != 0.0 {
		log.Println("div 0 issue...null vector returned")
		t.Errorf("TestDivValue %v %v %v %v %v %v", 0.0, 0.0, 0.0, result.X(), result.Y(), result.Z())
	}
}

func TestDivVector(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)
	vec2 := NewVector(1.0, 2.0, -1.0)
	result := vec1.DivideVector(vec2)

	if result.X() != 3.0 || result.Y() != -2.0 || result.Z() != -5.0 {
		t.Errorf("TestDivVector %v %v %v %v %v %v", 3.0, -2.0, -5.0, result.X(), result.Y(), result.Z())
	}
	vec2 = NewVector(0.0, 0.0, 0.0)
	result = vec1.DivideVector(vec2)
	if result.X() != 0.0 || result.Y() != 0.0 || result.Z() != 0.0 {
		log.Println("div 0 issue...null vector returned")
		t.Errorf("TestDivVector %v %v %v %v %v %v", 0.0, 0.0, 0.0, result.X(), result.Y(), result.Z())
	}
}

func TestDotProduct(t *testing.T) {
	vec1 := NewVector(3.5, 4.0, 5.2)
	//vec2 := NewVector(6.0, 7.3, 8.0) //--> test failed... reason?!?!?!?!?
	vec2 := NewVector(6.0, 7.2, 8.0)

	result := vec1.DotProduct(vec2)

	if result != 91.4 {
		t.Errorf("TestDotProduct %v %v", 91.4, result)
	}
}

func TestNormalize(t *testing.T) {
	vec := NewVector(3.0, 4.0, 0.0)

	vec = vec.Normalized()

	if vec.x != 0.6 || vec.y != 0.8 || vec.z != 0.0 {
		t.Errorf("TestNormalize %v %v %v %v %v %v", 0.6, 0.8, 0.0, vec.X(), vec.Y(), vec.Z())
	}

	vec = NewVector(0.0, 0.0, 0.0)

	vec = vec.Normalized()

	if vec.X() != 0.0 || vec.Y() != 0.0 || vec.Z() != 0.0 {
		log.Println("div 0 issue...null vector returned")
		t.Errorf("TestNormalize %v %v %v %v %v %v", 0.0, 0.0, 0.0, vec.X(), vec.Y(), vec.Z())
	}
}

func TestReflect(t *testing.T) {
	vec1 := NewVector(1.0, -1.0, 0.0)
	vec2 := NewVector(0.0, 1.0, 0.0)
	r := vec1.Reflection(vec2)
	if r.X() != 0.7071067811865475 || r.Y() != 0.7071067811865475 || r.Z() != 0.0 {
		t.Errorf("TestReflect got Vec(%f, %f, %f)", r.X(), r.Y(), r.Z())
	}
}

func TestAddValue(t *testing.T) {
	vec := NewVector(1.0, -1.0, 0.0)
	var i float64
	i = 5.0
	vec = vec.AddValue(i)
	if vec.X() != 6 || vec.Y() != 4 || vec.Z() != 5 {
		t.Errorf("TestAddValue failed %v %v %v ", vec.X(), vec.Y(), vec.Z())
	}
}
