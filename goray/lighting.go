package goray

func LambertDiffuse(surfaceToLight Vector, normal Vector) float64 {
	return Max(0, Dot(surfaceToLight, normal))
}
