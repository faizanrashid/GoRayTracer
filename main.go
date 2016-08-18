package main

import (
	"code/vector"
	"os"
	"strconv"
)

func main() {
	// Define picture size.
	nx, ny := 200, 100

	// Write metadata
	output := "P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n255\n"

	// Create and write the rgb values.
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			v := vector.New(
				float32(i)/float32(nx),
				float32(j)/float32(ny),
				0.2,
			)
			var ir int = int(255.99 * v.R())
			var ig int = int(255.99 * v.G())
			var ib int = int(255.99 * v.B())
			output += strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n"
		}
	}
	// Open file.
	f, err := os.Create("hello_world.ppm")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write to file.
	bytesWritten, err := f.Write([]byte(output))
	if err != nil {
		panic(err)
	}
}
