package main

import "goray"
import "image"
import "image/png"
import "image/color"
import "os"

func vecToNRGBA(v goray.Vector) color.NRGBA {
	return color.NRGBA {
		uint8(v.X * 255.0),
		uint8(v.Y * 255.0),
		uint8(v.Z * 255.0),
		255,
	}
}

func main() {

	sphere := goray.NewSphere(goray.Vector{0.0, 0.0, 1.0}, 0.1)

	cam := goray.Camera {
		goray.Vector{0, 0, 0},
		goray.Vector{0, 0, 1},
		goray.Vector{0, 1, 0},
		goray.Vector{1, 0, 0},
		1.0,
	}

	light := goray.Vector{-1.0, 0.4, 0}

	// 640 * 480 image
	view := goray.View {cam, 640, 480}
	bounds := image.Rect(0, 0, 640, 480)
	img := image.NewNRGBA64(bounds)
	
	// render
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			ray := view.Ray(int16(x - 320), int16(-y + 240))
			result := sphere.Trace(ray)
			if result.Hit {
				surfaceToLight := goray.Subtract(light, result.Intersection).Normalized()
				diffuse := goray.LambertDiffuse(surfaceToLight, result.Normal)
				shaded := result.Color.Scale(diffuse)
				img.Set(x, y, vecToNRGBA(shaded))
			}
		}
	}

	imgfile, _ := os.Create("rays.png")
	defer imgfile.Close()

	png.Encode(imgfile, img)
}
