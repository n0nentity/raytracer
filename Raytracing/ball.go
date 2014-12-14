//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

//Package Scene
package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

//a ball implementation
type Ball struct {
	sphere            *objects.Sphere
	color             *objects.Vector
	diffuse           float64
	specularIntensity float64
	specularPower     float64
	reflectivity      float64
}

//intersection between a ball and a ray, returns the intersection poisition,
//the color of the intersecting object (also a vector x=r, y=bg, z=b), and more
func (p *Ball) Intersection(line *objects.Ray) (position, color,
	normal *objects.Vector, diffuse,
	specularIntensity, specularPower, reflectivity float64) {

	position = p.sphere.Intersection(line)
	if position == nil {
		return
	}
	color = p.Color()
	normal = position.SubtractVector(p.Sphere().Position())
	diffuse = p.Diffuse()
	specularIntensity = p.SpecularIntensity()
	specularPower = p.SpecularPower()
	reflectivity = p.Reflectivity()
	return
}

//creates a new instance of a ball
func NewBall(sphere *objects.Sphere, color *objects.Vector,
	diffuse, specularIntensity, specularPower,
	reflectivity float64) *Ball {
	tmp := new(Ball)
	tmp.SetSphere(sphere)
	tmp.SetColor(color)
	tmp.SetDiffuse(diffuse)
	tmp.SetSpecularIntensity(specularIntensity)
	tmp.SetSpecularPower(specularPower)
	tmp.SetReflectivity(reflectivity)
	return tmp
}

//Getters and Setters

//returns the sphere
func (p *Ball) Sphere() *objects.Sphere { return p.sphere }

//returns the color of the ball, a vector
func (p *Ball) Color() *objects.Vector { return p.color }

//returns the diffuse value
func (p *Ball) Diffuse() float64 { return p.diffuse }

//returns the intensity value
func (p *Ball) SpecularIntensity() float64 { return p.specularIntensity }

//returns the power
func (p *Ball) SpecularPower() float64 { return p.specularPower }

//returns the reflectivity
func (p *Ball) Reflectivity() float64 { return p.reflectivity }

//sets the sphere
func (p *Ball) SetSphere(sphere *objects.Sphere) { p.sphere = sphere }

//sets the color of the ball
func (p *Ball) SetColor(color *objects.Vector) { p.color = color }

//sets the diffuse value
func (p *Ball) SetDiffuse(diffuse float64) { p.diffuse = diffuse }

//sets the intensity
func (p *Ball) SetSpecularIntensity(specularIntensity float64) {
	p.specularIntensity = specularIntensity
}

//sets the power
func (p *Ball) SetSpecularPower(specularPower float64) { p.specularPower = specularPower }

//sets the reflection
func (p *Ball) SetReflectivity(reflectivity float64) { p.reflectivity = reflectivity }
