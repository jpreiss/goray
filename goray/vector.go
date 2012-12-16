package goray

import "math"
import "math/rand"

// 3D vectors

type Vector struct {
	X float64
	Y float64
	Z float64
}

func Add(a, b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func Subtract(a, b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector) Scale(s float64) Vector {
	return Vector{s * a.X, s * a.Y, s * a.Z}
}

func (a Vector) Divide(d float64) Vector {
	return a.Scale(1.0 / d)
}

func Dot(a Vector, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Length2() float64 {
	return Dot(a, a)
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.Length2())
}

func (a Vector) Normalized() Vector {
	len2 := a.Length2()
	if len2 > 0 {
		return a.Divide(math.Sqrt(len2))
	}
	return a
}

// b must be normalized
func (a Vector) ProjectedOnto(b Vector) Vector {
	return b.Scale(Dot(a, b))
}

// b must be normalized
func (a Vector) OrthogonalTo(b Vector) Vector {
	return Subtract(a, a.ProjectedOnto(b))
}

func RandomVectorUniform(mins, maxes Vector) Vector {
	delta := Subtract(maxes, mins)
	return Vector{
		mins.X + rand.Float64() * delta.X,
		mins.Y + rand.Float64() * delta.Y,
		mins.Z + rand.Float64() * delta.Z,
	}
}

func VecMin(a, b Vector) Vector {
	return Vector{Min(a.X, b.X), Min(a.Y, b.Y), Min(a.Z, b.Z)}
}

func VecMax(a, b Vector) Vector {
	return Vector{Max(a.X, b.X), Max(a.Y, b.Y), Max(a.Z, b.Z)}
}

func (a Vector) ElementScale (b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}
