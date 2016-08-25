package camera

import (
	"github.com/faizanrashid/GoRayTracer/ray"
	"github.com/faizanrashid/GoRayTracer/vector"
)

type Camera struct {
	Origin          *vector.Vec3
	LowerLeftCorner *vector.Vec3
	Horizontal      *vector.Vec3
	Vertical        *vector.Vec3
}

func New(origin, lowerLeftCorner, horizontal, vertical *vector.Vec3) *Camera {
	return &Camera{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
	}
}

func (c *Camera) GetRay(u, v float64) *ray.Ray {
	directionVec := c.LowerLeftCorner.PlusVector(c.Horizontal.MultiplyScaler(u)).PlusVector(c.Vertical.MultiplyScaler(v))
	return ray.New(c.Origin, directionVec)
}
