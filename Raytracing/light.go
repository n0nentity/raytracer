//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357
package Raytracing

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

//a light implementation
type Light struct {
	position *objects.Vector
	color    *objects.Vector
}

//creates a new instance of a light
func NewLight(position, color *objects.Vector) *Light {
	light := new(Light)
	light.SetPosition(position)
	light.SetColor(color)
	return light
}

//returns the position
func (this *Light) Position() *objects.Vector { return this.position }

//returns the color
func (this *Light) Color() *objects.Vector { return this.color }

//sets the position
func (this *Light) SetPosition(position *objects.Vector) { this.position = position }

//sets the color
func (this *Light) SetColor(color *objects.Vector) { this.color = color }
