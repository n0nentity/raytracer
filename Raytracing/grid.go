//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

//a grid implementation
type Grid struct {
	top_left     *objects.Vector
	bottom_right *objects.Vector
}

//creates a new Instance of a grid
func NewGrid(top_left *objects.Vector, bottom_right *objects.Vector) *Grid {
	this := new(Grid)
	this.SetTopLeft(top_left)
	this.SetBottomRight(bottom_right)
	return this
}

//Getters and Setters

//returns the position of the topleft corner
func (p *Grid) TopLeft() *objects.Vector { return p.top_left }

//returns the position of the bottom right corner
func (p *Grid) BottomRight() *objects.Vector { return p.bottom_right }

//sets the topleft corner
func (p *Grid) SetTopLeft(top_left *objects.Vector) { p.top_left = top_left }

//sets the bottomright corner
func (p *Grid) SetBottomRight(bottom_right *objects.Vector) { p.bottom_right = bottom_right }
