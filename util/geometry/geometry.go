package geometry

import "math"

type Point3D struct {
	X int
	Y int
	Z int
}

func (p1 Point3D) Distance(p2 Point3D) float64 {
	return math.Sqrt(float64(
		(p2.X-p1.X)*(p2.X-p1.X) +
			(p2.Y-p1.Y)*(p2.Y-p1.Y) +
			(p2.Z-p1.Z)*(p2.Z-p1.Z),
	))
}

