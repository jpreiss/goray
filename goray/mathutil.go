package goray

func Max(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}

func Min(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}

func Clamp(x, low, high float64) float64 {
	clampLow := Max(x, low)
	return Min(clampLow, high)
}
