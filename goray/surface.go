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
	
	// direction dot direction, but direction is normalized
	a := 1.0
	
	b := 2.0 * (-Dot(s.Center, r.Direction) + Dot(r.Origin, r.Direction))

	c := Dot(s.Center, s.Center) - (2.0 * Dot(s.Center, r.Origin)) + Dot(r.Origin, r.Origin) - s.RadiusSquared;

	discriminant := (b * b) - (4 * a * c)

	if discriminant < 0 {
		result.Hit = false
		return result
	}

	hitTime := 0.0

	// todo: clean up conditionals...
	if math.Abs(discriminant) < math.SmallestNonzeroFloat64 {
		result.Hit = true
		hitTime = -b / (2.0 * a)
	} else {
		sqrtDiscriminant := math.Sqrt(discriminant)
		lowRoot := (-b - sqrtDiscriminant) / (2.0 * a)
		highRoot := (-b + sqrtDiscriminant) / (2.0 * a)
		if lowRoot > 0 {
			result.Hit = true
			hitTime = lowRoot
		} else if highRoot > 0 {
			result.Hit = true
			hitTime = highRoot
		} else {
			result.Hit = false
			return result
		}
	}

	result.Intersection = Add(r.Origin, r.Direction.Scale(hitTime))
	result.Normal = Subtract(result.Intersection, s.Center).Normalized()
	result.Color = s.Color
	return result
}
