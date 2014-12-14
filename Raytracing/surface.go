//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package scene

import (
	"de/vorlesung/projekt/raytracer/Helper"
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

//the surface implementation
type Surface struct {
	plane             *objects.Plane
	color             *objects.Vector
	diffuse           float64
	specularIntensity float64
	specularPower     float64
	reflectivity      float64
}

//calculaties the intersection between the surface and one ray
func (s *Surface) Intersection(line *objects.Ray) (position,
	color, normal *objects.Vector, diffuse,
	specularIntensity, specularPower, reflectivity float64) {

	position = s.plane.Intersection(line)
	if position == nil {
		return
	}
	color = s.Color().MultiplyVector(s.makeChessBoard(position))
	normal = s.Plane().Normal()
	diffuse = s.Diffuse()
	specularIntensity = s.SpecularIntensity()
	specularPower = s.SpecularPower()
	reflectivity = s.Reflectivity()
	return
}

//returns the color of a texture as a vector
//chessboard is created here
func (s *Surface) makeChessBoard(position *objects.Vector) *objects.Vector {
	position = position.Absolute()
	pX := Helper.Round(position.X(), 0)
	pY := Helper.Round(position.Y(), 0)
	pZ := Helper.Round(position.Z(), 0)
	if int(pX+pY+pZ)%2 == 0 {
		return objects.NewVector(0.0, 0.0, 0.0) //black chess field
	}
	return objects.NewVector(1.0, 1.0, 1.0) //white chess field
}

func NewSurface(plane *objects.Plane, color *objects.Vector, diffuse,
	specularIntensity, specularPower,
	reflectivity float64) *Surface {
	surface := new(Surface)
	surface.SetPlane(plane)
	surface.SetColor(color)
	surface.SetDiffuse(diffuse)
	surface.SetSpecularIntensity(specularIntensity)
	surface.SetSpecularPower(specularPower)
	surface.SetReflectivity(reflectivity)
	return surface
}

func (p *Surface) Plane() *objects.Plane          { return p.plane }
func (p *Surface) Color() *objects.Vector         { return p.color }
func (p *Surface) Diffuse() float64               { return p.diffuse }
func (p *Surface) SpecularIntensity() float64     { return p.specularIntensity }
func (p *Surface) SpecularPower() float64         { return p.specularPower }
func (p *Surface) Reflectivity() float64          { return p.reflectivity }
func (p *Surface) SetPlane(plane *objects.Plane)  { p.plane = plane }
func (p *Surface) SetColor(color *objects.Vector) { p.color = color }
func (p *Surface) SetDiffuse(diffuse float64)     { p.diffuse = diffuse }
func (p *Surface) SetSpecularIntensity(specularIntensity float64) {
	p.specularIntensity = specularIntensity
}
func (p *Surface) SetSpecularPower(specularPower float64) { p.specularPower = specularPower }
func (p *Surface) SetReflectivity(reflectivity float64)   { p.reflectivity = reflectivity }
