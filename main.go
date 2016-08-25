package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/faizanrashid/GoRayTracer/camera"
	"github.com/faizanrashid/GoRayTracer/hitable"
	"github.com/faizanrashid/GoRayTracer/material"
	"github.com/faizanrashid/GoRayTracer/ray"
	"github.com/faizanrashid/GoRayTracer/vector"
)

const (
	TMin = 0.001 // Minimum T to avoid collisions with the object itself.
)

/*
 Linearly blends white and blue depending on the up and downess of the
 Y coordinate and whether we hit a sphere.
*/
func colour(r *ray.Ray, world *hitable.HitableList, depth int) *vector.Vec3 {
	hitRecord := &hitable.HitRecord{}
	if world.Hit(r, hitRecord, TMin, math.MaxFloat64) {
		isScatter, scatterEffect := hitRecord.Material.Scatter(r, hitRecord.P, hitRecord.Normal)
		if depth < 50 && isScatter {
			return colour(scatterEffect.ScatteredRay, world, depth+1).MultiplyVector(scatterEffect.Attenuation)
		}
		return vector.New()
	}
	// Background colour
	unitDirection := vector.UnitVector(r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	// v(1.0, 1.0, 1.0) * (1.0-t) +  v(0.5, 0.7, 1.0)*t
	return vector.New(1.0, 1.0, 1.0).MultiplyScaler(1.0 - t).PlusVector(vector.New(0.5, 0.7, 1.0).MultiplyScaler(t))
}

/* Generate output to write to file*/
func createOutput(world *hitable.HitableList, cam *camera.Camera, nx, ny, ns int) string {
	// Initial metadata metadata.
	output := "P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n255\n"

	// Create and write the rgb values.
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			colourVec := vector.New()
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				colourVec.PlusEqualVector(colour(r, world, 0))
			}

			colourVec.DivideEqualScaler(float64(ns))
			// Gamma 2
			colourVec = vector.New(
				math.Sqrt(colourVec.R()),
				math.Sqrt(colourVec.G()),
				math.Sqrt(colourVec.B()))
			var ir int = int(255.99 * colourVec.R())
			var ig int = int(255.99 * colourVec.G())
			var ib int = int(255.99 * colourVec.B())
			output += strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n"
		}
	}

	return output
}

/*Display the progress bar*/
func progressBar(stopChan chan bool, doneChan chan bool) {
	progressTicker := time.NewTicker(100 * time.Millisecond)
	startTime := time.Now()
	for {
		select {
		case <-progressTicker.C:
			fmt.Print(".")
		case <-stopChan:
			progressTicker.Stop()
			fmt.Println("\nDone!\nTime elapsed: ", time.Since(startTime))
			doneChan <- true
			return
		}
	}
}

/* Opens a file named filename and writes the bytes to it.*/
func writeToFile(filename string, output []byte) error {
	// Open file.
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write to file.
	_, err = f.Write(output)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	doneChan := make(chan bool, 1)
	stopChan := make(chan bool, 1)
	go progressBar(stopChan, doneChan) // Start the progress bar.
	// Define picture size.
	nx, ny := 200, 100

	// Number of rays per pixel (Anti Aliasing)
	ns := 500

	// Define camera.
	cam := camera.New(vector.New(),
		vector.New(-2.0, -1.0, -1.0),
		vector.New(4.0, 0.0, 0.0),
		vector.New(0.0, 2.0, 0.0))

	// Define the hitable objects.
	world := &hitable.HitableList{}
	world.Add(hitable.NewSphere(vector.New(0, 0, -1),
		0.5,
		material.NewLambertian(vector.New(0.8, 0.3, 0.3))))
	world.Add(hitable.NewSphere(
		vector.New(0, -100.5, -1),
		100.0,
		material.NewLambertian(vector.New(0.8, 0.8, 0.0))))
	world.Add(hitable.NewSphere(
		vector.New(1.0, 0.0, -1),
		0.5,
		material.NewMetal(vector.New(0.8, 0.6, 0.2), 1.0)))
	world.Add(hitable.NewSphere(
		vector.New(-1, 0.0, -1),
		0.5,
		material.NewMetal(vector.New(0.8, 0.8, 0.8), 0.3)))

	output := createOutput(world, cam, nx, ny, ns)
	writeToFile("hello_world.ppm", []byte(output))

	// Stop progress bar.
	stopChan <- true
	<-doneChan // Block until everything is printed
}
