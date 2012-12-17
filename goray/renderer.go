package goray

import "image"
import "image/color"
import "math"


func vecToNRGBA(v Vector) color.NRGBA {
	v = VecMin(v, Vector{1,1,1}).Scale(255.0)
	return color.NRGBA {
		uint8(v.X),
		uint8(v.Y),
		uint8(v.Z),
		255,
	}
}

func RayTrace(surfaces []Surface, ray Ray, recursionDepth int) (bool, Vector) {

	maxRecursion := 8

	// TODO: full lighting control.  this really should not be here!
	light := Vector{-1.0, 1.0, 1.0}
	lightColor := Vector{1.0, 1.0, 1.0}

	shadowColorScale := Vector{0.6, 0.7, 1.0}
	ambientAmount := 0.3
	reflectivity := 0.5

	hit := false
	mindist := math.MaxFloat32
	color := Vector{}

	closestResult := TraceResult{}

	for _, surface := range surfaces {
		result := surface.Trace(ray)
		if result.Hit {
			hit = true
			surfaceToOrigin := Subtract(ray.Origin, result.Intersection)
			distance := surfaceToOrigin.Length()
			if distance < mindist {
				mindist = distance
				closestResult = result
				surfaceToLight := Subtract(light, result.Intersection).Normalized()
				surfaceToOrigin = surfaceToOrigin.Normalized()
				ambient := result.Color.ElementScale(shadowColorScale).Scale(ambientAmount)
				diffuse := result.Color.Scale(LambertDiffuse(surfaceToLight, result.Normal))
				specular := lightColor.Scale(Specular(surfaceToLight, result.Normal, surfaceToOrigin, 8))
				color = Add(Add(ambient, diffuse), specular)
			}
		}
	}

	if hit && (recursionDepth <= maxRecursion) {
		reflection := ray.Direction.Negate().Reflect(closestResult.Normal).Normalized()
		liftedOrigin := Add(closestResult.Intersection, reflection.Scale(0.000001))
		reflectedRay := Ray{liftedOrigin, reflection}
		recursiveHit, recursiveColor := RayTrace(surfaces, reflectedRay, recursionDepth + 1)
		if recursiveHit {
			color = Add(color, recursiveColor.Scale(reflectivity))
		} 
	}

	return hit, color
}


func Render(surfaces []Surface, camera Camera, width, height int) image.Image {

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

			hit, color := RayTrace(surfaces, ray, 1)

			if hit {
				img.Set(x, y, vecToNRGBA(color.Scale(0.7)))
			}
		}
	}

	return img
}
