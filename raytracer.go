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

	grid1 := scene.NewGrid(objects.NewVector(2.0, 1.00, -1.0), objects.NewVector(2.0, -0.50, 1.0))
	ball1 := scene.NewBall(sphere1, objects.NewVector(0.54, 0.60, 0.90), 0.9, 4.0, 30.0, 0.125)
	ball2 := scene.NewBall(sphere2, objects.NewVector(1.0, 2.0, 0.0), 0.9, 4.0, 30.0, 0.125)

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
