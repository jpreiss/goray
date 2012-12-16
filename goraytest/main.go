package main

import "goray"
import "image"
import "image/png"
import "image/color"
import "os"
import "math/rand"
import "time"

func vecToNRGBA(v goray.Vector) color.NRGBA {
	return color.NRGBA {
		uint8(v.X * 255.0),
		uint8(v.Y * 255.0),
		uint8(v.Z * 255.0),
		255,
	}
}

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

	light := goray.Vector{-3.0, 1.0, 1.0}

	// 640 * 480 image
	view := goray.View {cam, 640, 480}
	bounds := image.Rect(0, 0, 640, 480)
	img := image.NewNRGBA64(bounds)
	
	// render
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

			ray := view.Ray(int16(x - 320), int16(-y + 240))

			mindist := -1.0

			for _, surface := range surfaces {
				result := surface.Trace(ray)
				if result.Hit {
					distance := goray.Subtract(result.Intersection, cam.Position).Length()
					iAmClosest := (mindist == -1.0) || distance < mindist
					if iAmClosest {
						surfaceToLight := goray.Subtract(light, result.Intersection).Normalized()
						diffuse := goray.LambertDiffuse(surfaceToLight, result.Normal)
						shaded := result.Color.Scale(diffuse)
						img.Set(x, y, vecToNRGBA(shaded))
						mindist = distance
					}
				}
			}
		}
	}

	imgfile, _ := os.Create("rays.png")
	defer imgfile.Close()

	png.Encode(imgfile, img)
}
