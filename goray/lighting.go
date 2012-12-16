package goray

import "math"

func LambertDiffuse(surfaceToLight Vector, normal Vector) float64 {
	return Max(0, Dot(surfaceToLight, normal))
}

func Specular(surfaceToLight, normal, surfaceToCamera Vector, power float64) float64 {
	reflection := Subtract(normal, surfaceToLight.OrthogonalTo(normal)).Normalized()
	dot := Dot(reflection, surfaceToCamera)
	if dot < 0 {
		return 0
	}
	return math.Pow(dot, power)
}
