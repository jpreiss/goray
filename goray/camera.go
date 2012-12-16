package goray

import "math"

type Camera struct {
	Position Vector
	To       Vector
	Up       Vector
	Right    Vector
	FovWidth float64
}

type View struct {
	Camera Camera
	Width  uint16
	Height uint16
}

// gets the ray corresponding to an image pixel
// TODO: is totally wrong right now! completely ignores camera orientation!
// only works when camera is pointing in +Z.
func (v View) Ray(x int16, y int16) Ray {

	// put our sensor at the distance
	// such that its width is 2 (-1..1)
	sensorDistance := 1.0 / math.Tan(v.Camera.FovWidth/2.0)

	pixelSize := 2.0 / float64(v.Width)

	ray := Vector{float64(x) * pixelSize, float64(y) * pixelSize, sensorDistance}

	return Ray{v.Camera.Position, ray.Normalized()}
}
