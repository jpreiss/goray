package main

import "goray"
import "image"
import "image/png"
import "image/color"
import "os"

func main() {

	sphere := goray.Sphere{goray.Vector{0.0, 0.0, 1.0}, 0.1 }

	cam := goray.Camera {
		goray.Vector{0, 0, 0},
		goray.Vector{0, 0, 1},
		goray.Vector{0, 1, 0},
		goray.Vector{1, 0, 0},
		1.0,
	}

	// 640 * 480 image
	view := goray.View {cam, 640, 480}

	bounds := image.Rect(0, 0, 640, 480)

	img := image.NewNRGBA64(bounds)

	
	// render
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			
			ray := view.Ray(int16(x - 320), int16(y - 240))

			if sphere.Hit(ray) {
				cv := sphere.Color(ray)
				col := color.NRGBA {
					uint8(cv.X * 255.0),
					uint8(cv.Y * 255.0),
					uint8(cv.Z * 255.0),
					255,
				}
				img.Set(x, y, col)
			}
		}
	}

	imgfile, _ := os.Create("rays.png")
	defer imgfile.Close()

	png.Encode(imgfile, img)
}
