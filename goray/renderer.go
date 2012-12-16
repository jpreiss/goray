package goray

import "image"
import "image/color"


func vecToNRGBA(v Vector) color.NRGBA {
	v = VecMin(v, Vector{1,1,1})
	return color.NRGBA {
		uint8(v.X * 255.0),
		uint8(v.Y * 255.0),
		uint8(v.Z * 255.0),
		255,
	}
}

func Render(surfaces []Surface, camera Camera, width, height int) image.Image {

	// TODO: full lighting control
	light := Vector{-1.0, 1.0, 1.0}
	lightColor := Vector{1.0, 1.0, 1.0}

	shadowColorScale := Vector{0.6, 0.7, 1.0}
	ambientAmount := 0.4

	view := View {camera, width, height}
	bounds := image.Rect(0, 0, width, height)
	img := image.NewNRGBA64(bounds)

	halfwidth := width / 2
	halfheight := height / 2

	// render
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

			// compensate for image y direction
			ray := view.Ray(x - halfwidth, -y + halfheight)

			mindist := -1.0

			for _, surface := range surfaces {
				result := surface.Trace(ray)
				if result.Hit {
					surfaceToCamera := Subtract(camera.Position, result.Intersection)
					distance := surfaceToCamera.Length()
					iAmClosest := (mindist == -1.0) || distance < mindist
					if iAmClosest {
						surfaceToLight := Subtract(light, result.Intersection).Normalized()
						surfaceToCamera = surfaceToCamera.Normalized()
						ambient := result.Color.ElementScale(shadowColorScale).Scale(ambientAmount)
						diffuse := result.Color.Scale(LambertDiffuse(surfaceToLight, result.Normal))
						specular := lightColor.Scale(Specular(surfaceToLight, result.Normal, surfaceToCamera, 8))
						shaded := Add(Add(ambient, diffuse), specular)
						img.Set(x, y, vecToNRGBA(shaded))
						mindist = distance
					}
				}
			}
		}
	}

	return img
}
