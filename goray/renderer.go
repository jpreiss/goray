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

func RayTrace(surfaces []Surface, ray Ray) (TraceResult) { 

	mindist := math.MaxFloat32
	closestResult := TraceResult{}
	closestResult.Hit = false

	for _, surface := range surfaces {
		result := surface.Trace(ray)
		if result.Hit {
			surfaceToOrigin := Subtract(ray.Origin, result.Intersection)
			distance := surfaceToOrigin.Length()
			if distance < mindist {
				mindist = distance
				closestResult = result
			}
		}
	}
	return closestResult
}

func RayTraceRecursive(surfaces []Surface, ray Ray, recursionDepth int) (bool, Vector) { 

	maxRecursion := 8

	// TODO: full lighting control.  this really should not be here!
	light := Vector{-1.0, 1.0, 1.0}
	lightColor := Vector{1.0, 1.0, 1.0}

	shadowColorScale := Vector{0.6, 0.7, 1.0}
	ambientAmount := 0.4
	reflectivity := 0.5

	black := Vector{0.0, 0.0, 0.0}

	closestResult := RayTrace(surfaces, ray)
	if !closestResult.Hit {
		return false, black
	}

	ambient := closestResult.Color.ElementScale(shadowColorScale).Scale(ambientAmount)
	myColor := ambient

	surfaceToLight := Subtract(light, closestResult.Intersection).Normalized()
	liftedIntersection := Add(closestResult.Intersection, closestResult.Normal.Scale(0.000001))
	rayToLight := Ray{liftedIntersection, surfaceToLight}

	shadowResult := RayTrace(surfaces, rayToLight)
	if !shadowResult.Hit {
		// not in shadow - use phong lighting
		surfaceToOrigin := Subtract(ray.Origin, closestResult.Intersection).Normalized()
		diffuse := closestResult.Color.Scale(LambertDiffuse(surfaceToLight, closestResult.Normal))
		specular := lightColor.Scale(Specular(surfaceToLight, closestResult.Normal, surfaceToOrigin, 8))
		myColor = Add(Add(ambient, diffuse), specular)
	}

	// reflections. they happen whether we are in shadow or not
	if recursionDepth <= maxRecursion {
		reflection := ray.Direction.Negate().Reflect(closestResult.Normal).Normalized()
		reflectedRay := Ray{liftedIntersection, reflection}
		recursiveHit, recursiveColor := RayTraceRecursive(surfaces, reflectedRay, recursionDepth + 1)
		if recursiveHit {
			myColor = Add(myColor, recursiveColor.Scale(reflectivity))
		} 
	}

	return true, myColor
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

			hit, color := RayTraceRecursive(surfaces, ray, 1)

			if hit {
				img.Set(x, y, vecToNRGBA(color.Scale(0.7)))
			}
		}
	}

	return img
}
