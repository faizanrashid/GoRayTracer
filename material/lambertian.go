package material

import (
	"github.com/faizanrashid/GoRayTracer/ray"
	"github.com/faizanrashid/GoRayTracer/vector"
)

type Lambertian struct {
	Albedo *vector.Vec3
}

func NewLambertian(albedo *vector.Vec3) *Lambertian {
	return &Lambertian{
		Albedo: albedo,
	}
}

func (l *Lambertian) Scatter(rayIn *ray.Ray, hitPoint *vector.Vec3, normal *vector.Vec3) (bool, *ScatterEffect) {
	target := hitPoint.PlusVector(normal).PlusVector(vector.RandomInUnitSphere())
	scatterEffect := &ScatterEffect{
		ScatteredRay: ray.New(hitPoint, target.MinusVector(hitPoint)),
		Attenuation:  l.Albedo,
	}
	return true, scatterEffect
}
