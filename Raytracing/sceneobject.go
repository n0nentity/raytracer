//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

//the interface which every scene object implements
//ball and the surface implements this interface (ducktyping)
type SceneObject interface {
	Intersection(*objects.Ray) (position, color,
		normal *objects.Vector, diffuse,
		specularIntensity, specularPower, reflectivity float64)
	Color() *objects.Vector
}
