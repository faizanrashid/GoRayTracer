package material

import (
	"github.com/faizanrashid/GoRayTracer/ray"
	"github.com/faizanrashid/GoRayTracer/vector"
)

type ScatterEffect struct {
	ScatteredRay *ray.Ray
	Attenuation  *vector.Vec3
}

type Material interface {
	Scatter(rayIn *ray.Ray, hitPoint *vector.Vec3, normal *vector.Vec3) (bool, *ScatterEffect)
}
