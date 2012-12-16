package main

import "goray"
import "image/png"
import "os"
import "math/rand"
import "time"


func main() {

	numSpheres := 20

	rand.Seed(time.Now().UTC().UnixNano())

	surfaces := make([]goray.Surface, numSpheres)

	black := goray.Vector{0, 0, 0}
	white := goray.Vector{1, 1, 1}

	boxmin := goray.Vector{-1, -1, 0.5}
	boxmax := goray.Vector{1, 1, 5.0}

	for i := range surfaces {
		randomPosition := goray.RandomVectorUniform(boxmin, boxmax)
		randomColor := goray.RandomVectorUniform(black, white)
		surfaces[i] = goray.NewSphere(randomPosition, 0.3, randomColor)
	}

	cam := goray.Camera {
		goray.Vector{0, 0, 0},
		goray.Vector{0, 0, 1},
		goray.Vector{0, 1, 0},
		goray.Vector{1, 0, 0},
		1.0,
	}

	img := goray.Render(surfaces, cam, 640, 480)

	imgfile, _ := os.Create("rays.png")
	defer imgfile.Close()

	png.Encode(imgfile, img)
}
