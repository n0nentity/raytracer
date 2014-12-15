//Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"image"
	imageColor "image/color"
	"math"
)

//global variable
var global = 0.25

//the scene implementation
// a scene has a view (the view attribute), a grid, several elements and so on
type Scene struct {
	view     *objects.Vector
	grid     *Grid
	elements []SceneObject
	ambient  *objects.Vector
	skyColor *objects.Vector
	light    *Light
}

//the general rendering function for the current scene,
//returns the image which one the imageWriter in the helper package saves to disk
func (s *Scene) Render(x, y, supersampling int) image.Image {
	x2 := x * supersampling
	y2 := y * supersampling
	img := make([][]objects.Vector, x2)

	starting := s.Grid().TopLeft()
	sizeZ := math.Abs(starting.Z()-s.Grid().BottomRight().Z()) / float64(x2)
	sizeY := -math.Abs(starting.Y()-s.Grid().BottomRight().Y()) / float64(y2)

	for i := 0; i < x2; i++ {
		img[i] = make([]objects.Vector, y2)
		for j := 0; j < y2; j++ {
			posZ := starting.Z() + (float64(i)+0.5)*sizeZ
			posY := starting.Y() + (float64(j)+0.5)*sizeY
			gridPosition := objects.NewVector(starting.X(), posY, posZ)

			color, _ := s.raytracing(s.View(), gridPosition.SubtractVector(s.View()), nil, 8)
			if color == nil {
				color = s.SkyColor()
			}

			img[i][j] = *color
		}

		//parallel rendering..
		//race condition !!! can be that the image is saved but we are still working on it
		// go func(i, y2 int, img [][]objects.Vector) {
		// 	for j := 0; j < y2; j++ {

		// 		posZ := starting.Z() + (float64(i)+0.5)*sizeZ
		// 		posY := starting.Y() + (float64(j)+0.5)*sizeY
		// 		gridPosition := objects.NewVector(starting.X(), posY, posZ)

		// 		color, _ := s.raytracing(s.View(), gridPosition.SubtractVector(s.View()), nil, 8)
		// 		if color == nil {
		// 			color = s.skyColor
		// 		}

		// 		img[i][j] = *color
		// 	}
		// }(i, y2, img)
	}
	return scale(img, x2, y2, supersampling)
}

//scaling function
func scale(in [][]objects.Vector, size_x, size_y, factor int) *image.RGBA {
	var out = image.NewRGBA(image.Rect(0, 0, size_x/factor, size_y/factor))

	for i := 0; i < size_x; i += factor {
		for j := 0; j < size_y; j += factor {
			var tmpVal = objects.NewVector(0.0, 0.0, 0.0)
			for k := 0; k < factor; k++ {
				for l := 0; l < factor; l++ {
					//race condition
					tmpVal = tmpVal.AddVector(&in[i+k][j+l])
				}
			}

			tmpVal = tmpVal.DivideValue(float64(factor * factor))
			var outColor = new(imageColor.RGBA)
			outColor.R = uint8(tmpVal.X() * 255)
			outColor.G = uint8(tmpVal.Y() * 255)
			outColor.B = uint8(tmpVal.Z() * 255)
			outColor.A = 255
			out.Set(i/factor, j/factor, outColor)
		}
	}
	return out
}

//the recursive raytracing function
func (this *Scene) raytracing(position, direction *objects.Vector, ignored SceneObject, depthLeft uint8) (*objects.Vector, *objects.Vector) {
	Ray := objects.NewRay(position, direction)

	intersectObject, intersectPos, color, normal, diffuse, specularIntensity, specularPower, reflectivity, nearestDist := this.intersectAll(Ray, ignored)

	if math.IsInf(nearestDist, 1) {
		color = nil
	} else {
		var phongColor *objects.Vector
		var phongSpecular *objects.Vector
		if this.light != nil {
			phongColor, phongSpecular = this.phongCalculation(intersectObject, intersectPos, Ray, normal, diffuse, specularIntensity/5, specularPower*2)
		} else {
			phongColor = objects.NewVector(1.0, 1.0, 1.0)
			phongSpecular = objects.NewVector(0.0, 0.0, 0.0)
		}

		color = color.MultiplyVector(phongColor).AddVector(phongSpecular).Limit(0, 1)

		if depthLeft > 0 && reflectivity > 0 {
			reflectColor, reflectPos := this.raytracing(intersectPos, direction.Reflection(normal), intersectObject, depthLeft-1)
			if reflectColor != nil {
				color = color.MultiplyValue(1-reflectivity).AddVector(reflectColor.MultiplyValue(reflectivity)).Limit(0, 1)

				// Ambient occlusion
				reflectDistance := reflectPos.SubtractVector(intersectPos).Length()
				if reflectDistance < global {
					color = color.MultiplyVector(objects.NewVector(1.0, 1.0, 1.0).MultiplyValue(reflectDistance/global).Limit(0.25, 1))
				}
			} else {
				color = color.MultiplyValue(1-reflectivity).AddVector(this.SkyColor().MultiplyValue(reflectivity)).Limit(0, 1)
			}
		}
	}
	return color, intersectPos
}

func (p *Scene) isShadow(intersectPos, directionToLight *objects.Vector, intersectObject SceneObject) bool {
	r := objects.NewRay(intersectPos, directionToLight)
	_, pos, _, _, _, _, _, _, _ := p.intersectAll(r, intersectObject)
	return pos != nil
}

//calculates the phone
func (p *Scene) phongCalculation(intersectObject SceneObject, intersectPos *objects.Vector, ray *objects.Ray, normal *objects.Vector,
	diffuse, specularIntensity,
	specularPower float64) (phongColor, phongSpecular *objects.Vector) {

	phongColor = p.ambient
	phongSpecular = objects.NewVector(0.0, 0.0, 0.0)

	directionToLight := p.light.Position().SubtractVector(intersectPos).Normalized()
	if !p.isShadow(intersectPos, directionToLight, intersectObject) {

		directionFromLight := directionToLight.MultiplyValue(-1)
		normal = normal.Normalized()
		phongDiffuse := p.light.Color().MultiplyValue(diffuse*directionToLight.DotProduct(normal)).Limit(0, 1)

		directionLightOut := directionFromLight.Reflection(normal).Normalized()
		directionIntersectView := ray.Origin().SubtractVector(intersectPos).Normalized()
		specularAmount := directionLightOut.DotProduct(directionIntersectView)
		if specularAmount < 0 {
			specularAmount = 0
		}
		phongSpecular = p.light.Color().MultiplyValue(specularIntensity*math.Pow(specularAmount, specularPower)).Limit(0, 1)

		phongColor = phongColor.AddVector(phongDiffuse).Limit(0, 1)
	}
	return
}

//intersects all elements in the current scene to calculate the shadow
func (p *Scene) intersectAll(Ray *objects.Ray, ignored SceneObject) (intersectObject SceneObject,
	intersectPos, color, normal *objects.Vector,
	diffuse, specularIntensity, specularPower,
	reflectivity, nearestDist float64) {

	nearestDist = math.Inf(1)
	for _, element := range p.elements {
		if ignored != nil && element == ignored {
			continue
		}

		intersectP, col, norm, dif, spInt, spPow, ref := element.Intersection(Ray)
		if intersectP == nil {
			continue
		}
		length := intersectP.SubtractVector(Ray.Origin()).Length()
		if length < nearestDist {
			intersectObject = element
			intersectPos = intersectP
			color = col
			normal = norm
			diffuse = dif
			specularIntensity = spInt
			specularPower = spPow
			reflectivity = ref
			nearestDist = length
		}
	}
	return
}

//creates a new instace of a scene
func NewScene(view *objects.Vector, grid *Grid) *Scene {
	if grid.TopLeft().X() != grid.BottomRight().X() {
		panic("not same x values on corners")
	}
	scene := new(Scene)
	scene.SetView(view)
	scene.SetGrid(grid)
	scene.SetElements(make([]SceneObject, 0, 0))
	scene.SetAmbient(objects.NewVector(0.0, 0.0, 0.0))
	scene.SetSkyColor(objects.NewVector(0.0, 0.0, 0.0))
	scene.SetLight(NewLight(objects.NewVector(0.0, 0.0, 0.0), objects.NewVector(0.0, 0.0, 0.0)))
	return scene
}

//Getters and Setters
//
func (s *Scene) View() *objects.Vector                { return s.view }
func (s *Scene) Grid() *Grid                          { return s.grid }
func (s *Scene) Elements() []SceneObject              { return s.elements }
func (s *Scene) Ambient() *objects.Vector             { return s.ambient }
func (s *Scene) Light() *Light                        { return s.light }
func (s *Scene) SkyColor() *objects.Vector            { return s.skyColor }
func (s *Scene) SetView(view *objects.Vector)         { s.view = view }
func (s *Scene) SetGrid(grid *Grid)                   { s.grid = grid }
func (s *Scene) SetElements(elements []SceneObject)   { s.elements = elements }
func (s *Scene) SetAmbient(ambient *objects.Vector)   { s.ambient = ambient }
func (s *Scene) SetLight(light *Light)                { s.light = light }
func (s *Scene) SetSkyColor(skyColor *objects.Vector) { s.skyColor = skyColor }
func (s *Scene) AddElement(element SceneObject)       { s.elements = append(s.elements, element) }
