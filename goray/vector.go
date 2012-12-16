package goray

import("math")

// 3D vectors

type Vector struct {
	X float64
	Y float64
	Z float64
}

func add(a, b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func subtract(a, b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector) scale(s float64) Vector {
	return Vector{s * a.X, s * a.Y, s * a.Z}
}

func (a Vector) divide(d float64) Vector {
	return a.scale(1.0 / d)
}

func dot(a Vector, b Vector) float64 {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func (a Vector) length2() float64 {
	return dot(a, a)
}

func (a Vector) length() float64 {
	return math.Sqrt(a.length2())
}

func (a Vector) normalized() Vector {
	len2 := a.length2()
	if len2 > 0 {
		return a.divide(math.Sqrt(len2))
	}
	return a
}

// b must be normalized
func (a Vector) projectedOnto(b Vector) Vector {
	return b.scale(dot(a, b))
}

// b must be normalized
func (a Vector) orthogonalTo(b Vector) Vector {
	return subtract(a, a.projectedOnto(b))
}
