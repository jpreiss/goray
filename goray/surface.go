package goray

type Ray struct {
	Origin Vector
	Direction Vector
}

type Surface interface {
	Hit(r Ray) bool
	Color(r Ray) Vector
}

type Sphere struct {
	Center Vector
	Radius float64
}

func (s Sphere) Hit(r Ray) bool {
	originToCenter := subtract(r.Origin, s.Center)
	rayToCenter := originToCenter.orthogonalTo(r.Direction)
	return rayToCenter.length2() < (s.Radius * s.Radius)
}

// fixed color
func (s Sphere) Color(r Ray) Vector {
	return Vector {0.8, 0.0, 0.2}
}