//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package objects

//basic plane implementation
type Plane struct {
	pos    *Vector
	normal *Vector
}

//calculates the intersection between the plane and one ray
//returns the position of intersection, a vector
func (p *Plane) Intersection(ray *Ray) *Vector {
	t := p.Position().SubtractVector(ray.Origin()).DotProduct(p.Normal()) / ray.Direction().DotProduct(p.Normal())
	if t <= 0 {
		return nil
	}
	return ray.AtStep(t)
}

//creates a new instance of a plane
func NewPlane(pos, normal *Vector) *Plane {
	var tmp = new(Plane)
	tmp.SetPosition(pos)
	tmp.SetNormal(normal)
	return tmp
}

//Getters and Setters

//returns the position vector
func (p *Plane) Position() *Vector { return p.pos }

//returns the normal vector
func (p *Plane) Normal() *Vector { return p.normal }

//sets the position
func (p *Plane) SetPosition(pos *Vector) { p.pos = pos }

//sets the normal
func (p *Plane) SetNormal(normal *Vector) { p.normal = normal }
