/*The MIT License (MIT)

Copyright (c) 2014 Michael Heier 8311689, Patrick Dahlke 2458357

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"de/vorlesung/projekt/raytracer/Helper"
	scene "de/vorlesung/projekt/raytracer/Raytracing"
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"log"
	"os"
	"runtime"
	"time"
)

//the raytracer main method
func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	h := new(Helper.Helper)
	width := 640
	height := 480
	filename := "out.png"

	f, _ := os.Create("logfile")
	log.SetOutput(f)
	log.Println("################################")
	log.Println("start rendering")

	sphere1 := objects.NewSphere(objects.NewVector(0.0, 0.0, 1.0), 1.0)
	sphere2 := objects.NewSphere(objects.NewVector(-2.5, 0.0, -0.75), 1.0)
	plane := objects.NewPlane(objects.NewVector(0.0, -1.0, 0.0), objects.NewVector(0.0, 1.0, 0.0))

	basicBlue := objects.NewVector(0.54, 0.60, 0.90)
	neonYellow := objects.NewVector(1.0, 2.0, 0.0)

	grid1 := scene.NewGrid(objects.NewVector(2.0, 1.00, -1.0), objects.NewVector(2.0, -0.50, 1.0))
	ball1 := scene.NewBall(sphere1, basicBlue, 0.9, 4.0, 30.0, 0.125)
	ball2 := scene.NewBall(sphere2, neonYellow, 0.9, 4.0, 30.0, 0.125)

	surface1 := scene.NewSurface(plane, objects.NewVector(1.0, 1.0, 1.0), 1.0, 1.0, 8.0, 0.05)
	light1 := scene.NewLight(objects.NewVector(1.0, 4.0, 0.5), objects.NewVector(1.0, 1.0, 1.0))

	colorSky := objects.NewVector(0.85, 0.85, 0.95)

	currentScene := scene.NewScene(objects.NewVector(4.0, 0.5, 0.0), grid1)
	currentScene.AddElement(scene.SceneObject(ball1))
	currentScene.AddElement(scene.SceneObject(ball2))
	currentScene.AddElement(scene.SceneObject(surface1))
	currentScene.SetAmbient(objects.NewVector(0.25, 0.25, 0.3))
	currentScene.SetLight(light1)
	currentScene.SetSkyColor(colorSky)

	//benchmark
	t1 := time.Now()
	//get better result with 3 or 4 instead of 1
	//i := currentScene.Render(width, height, 5)
	i := currentScene.Render(width, height, 1)
	log.Println("Rendering time: ", time.Since(t1))
	err := h.ImageWriter(filename, i)
	if err != nil {
		log.Println("Error in image write: ", err)
	}
}
