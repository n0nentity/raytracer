//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package SceneObjects

//ray implementation
type Ray struct {
	origin    *Vector
	direction *Vector
}

func (this *Ray) AtStep(count float64) *Vector {
	return this.Origin().AddVector(this.Direction().MultiplyValue(count))
}

//creats a new instance of a ray
func NewRay(origin, direction *Vector) *Ray {
	ray := new(Ray)
	ray.SetOrigin(origin)
	ray.SetDirection(direction.Normalized())
	return ray
}

//Getters and Setters
//returns the origin of the ray
func (this *Ray) Origin() *Vector { return this.origin }

//returns the direction of the ray
func (this *Ray) Direction() *Vector { return this.direction }

//sets the origin
func (this *Ray) SetOrigin(origin *Vector) { this.origin = origin }

//sets the direction
func (this *Ray) SetDirection(direction *Vector) { this.direction = direction.Normalized() }
