package goray

import "math"

type Ray struct {
	Origin    Vector
	Direction Vector
}

type TraceResult struct {
	Hit          bool
	Intersection Vector
	Normal       Vector
	Color        Vector
}

type Surface interface {
	Trace(r Ray) TraceResult
}

type Sphere struct {
	Center        Vector
	Radius        float64
	RadiusSquared float64
}

func NewSphere(center Vector, radius float64) Sphere {
	return Sphere{center, radius, radius * radius}
}

func (s Sphere) Trace(r Ray) TraceResult {
	result := TraceResult{}
	originToCenter := Subtract(s.Center, r.Origin)
	rayToCenter := originToCenter.OrthogonalTo(r.Direction)
	result.Hit = rayToCenter.Length2() < s.RadiusSquared
	if result.Hit {
		secantMidpoint := Subtract(s.Center, rayToCenter)
		halfSecantLength := math.Sqrt(rayToCenter.Length2() + s.RadiusSquared)
		stepBack := r.Direction.Scale(halfSecantLength)
		result.Intersection = Subtract(secantMidpoint, stepBack)
		result.Normal = Subtract(result.Intersection, s.Center).Normalized()
		result.Color = Vector{0.8, 0.0, 0.2}
	}
	return result
}
