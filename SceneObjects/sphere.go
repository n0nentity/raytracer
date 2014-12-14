//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package objects

import (
	"math"
)

//the sphere
type Sphere struct {
	position *Vector
	radius   float64
}

//the function to calculate the intersection point between a ray and the sphere
//returns the position of the intersection, a vector
func (s *Sphere) Intersection(ray *Ray) *Vector {
	d := ray.Origin().SubtractVector(s.Position())
	db := ray.Direction().DotProduct(d)
	//subSqrt := math.Pow(db, 2) + math.Pow(p.Radius(), 2) - math.Pow(d.Length(), 2)
	subSqrt := db*db + s.Radius()*s.Radius() - d.Length()*d.Length()
	if subSqrt < 0 {
		return nil
	}
	sqrt := math.Sqrt(subSqrt)
	t1 := -db + sqrt
	t2 := -db - sqrt
	if t1 < 0 && t2 < 0 {
		return nil
	}
	var jumps = math.Min(t1, t2)
	if t1 < 0 {
		jumps = t2
	}
	if t2 < 0 {
		jumps = t1
	}
	return ray.AtStep(jumps)
}

//Creates a new Instance of the Sphere struct
func NewSphere(position *Vector, radius float64) *Sphere {
	tmp := new(Sphere)
	tmp.SetPosition(position)
	tmp.SetRadius(radius)
	return tmp
}

//Getters and Setters

//returns the Position of the Sphere
func (s *Sphere) Position() *Vector { return s.position }

//return the Radius of the Sphere
func (s *Sphere) Radius() float64 { return s.radius }

//sets the position
func (s *Sphere) SetPosition(position *Vector) { s.position = position }

//sets the radius
func (s *Sphere) SetRadius(radius float64) { s.radius = radius }
