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
	Color         Vector    
}

func NewSphere(center Vector, radius float64, color Vector) Sphere {
	return Sphere{center, radius, radius * radius, color}
}

func (s Sphere) Trace(r Ray) TraceResult {
	result := TraceResult{}
	originToCenter := Subtract(s.Center, r.Origin)
	rayToCenter := originToCenter.OrthogonalTo(r.Direction)
	secantMidpoint := Subtract(s.Center, rayToCenter)
	result.Hit = rayToCenter.Length2() < s.RadiusSquared && Dot(secantMidpoint, r.Direction) > 0
	if result.Hit {
		halfSecantLength := math.Sqrt(rayToCenter.Length2() + s.RadiusSquared)
		stepBack := r.Direction.Scale(halfSecantLength)
		result.Intersection = Subtract(secantMidpoint, stepBack)
		result.Normal = Subtract(result.Intersection, s.Center).Normalized()
		result.Color = s.Color
	}
	return result
}
