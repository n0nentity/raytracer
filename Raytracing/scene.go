package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"image"
	imageColor "image/color"
	"log"
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
func (s *Scene) Render(x, y int, superSample int) image.Image {
	log.Println("Start rendering (function) Image\n")
	tmp_x := x * superSample
	tmp_y := y * superSample
	tmp_img := make([][]objects.Vector, tmp_x)

	rasterStart := s.grid.TopLeft()
	rasterSizeZ := math.Abs(rasterStart.Z()-s.grid.BottomRight().Z()) / float64(tmp_x)
	rasterSizeY := -math.Abs(rasterStart.Y()-s.grid.BottomRight().Y()) / float64(tmp_y)

	for i := 0; i < tmp_x; i++ {
		tmp_img[i] = make([]objects.Vector, tmp_y)
		for j := 0; j < tmp_y; j++ {

			posZ := rasterStart.Z() + (float64(i)+0.5)*rasterSizeZ
			posY := rasterStart.Y() + (float64(j)+0.5)*rasterSizeY
			gridPos := objects.NewVector(rasterStart.X(), posY, posZ)

			var color, _ = s.raytracing(s.View(), gridPos.SubtractVector(s.View()), nil, 8)
			if color == nil {
				color = s.skyColor
			}

			tmp_img[i][j] = *color
		}

		//parallel rendering
		// go func(i, tmp_y int, tmp_img [][]objects.Vector) {
		// 	for j := 0; j < tmp_y; j++ {

		// 		posZ := rasterStart.Z() + (float64(i)+0.5)*rasterSizeZ
		// 		posY := rasterStart.Y() + (float64(j)+0.5)*rasterSizeY
		// 		gridPos := objects.NewVector(rasterStart.X(), posY, posZ)

		// 		var color, _ = s.raytracing(s.View(), gridPos.SubtractVector(s.View()), nil, 8)
		// 		if color == nil {
		// 			color = s.skyColor
		// 		}

		// 		tmp_img[i][j] = *color
		// 	}
		// }(i, tmp_y, tmp_img)
	}
	return scale(tmp_img, tmp_x, tmp_y, superSample)
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

		var intersectP, col, norm, dif, spInt, spPow, ref = element.Intersection(Ray)
		if intersectP == nil {
			continue
		}
		var length = intersectP.SubtractVector(Ray.Origin()).Length()
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
func (p *Scene) View() *objects.Vector                { return p.view }
func (p *Scene) Grid() *Grid                          { return p.grid }
func (p *Scene) Elements() []SceneObject              { return p.elements }
func (p *Scene) Ambient() *objects.Vector             { return p.ambient }
func (p *Scene) Light() *Light                        { return p.light }
func (p *Scene) SkyColor() *objects.Vector            { return p.skyColor }
func (p *Scene) SetView(view *objects.Vector)         { p.view = view }
func (p *Scene) SetGrid(grid *Grid)                   { p.grid = grid }
func (p *Scene) SetElements(elements []SceneObject)   { p.elements = elements }
func (p *Scene) SetAmbient(ambient *objects.Vector)   { p.ambient = ambient }
func (p *Scene) SetLight(light *Light)                { p.light = light }
func (p *Scene) SetSkyColor(skyColor *objects.Vector) { p.skyColor = skyColor }
func (p *Scene) AddElement(element SceneObject)       { p.elements = append(p.elements, element) }
