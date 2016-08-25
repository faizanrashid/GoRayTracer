package material

import (
	"github.com/faizanrashid/GoRayTracer/ray"
	"github.com/faizanrashid/GoRayTracer/vector"
)

type Metal struct {
	Albedo *vector.Vec3
	Fuzz   float64
}

func NewMetal(albedo *vector.Vec3, fuzz float64) *Metal {
	if fuzz > 1 {
		fuzz = 1.0
	}

	return &Metal{
		Albedo: albedo,
		Fuzz:   fuzz,
	}
}

func (m *Metal) Scatter(rayIn *ray.Ray, hitPoint *vector.Vec3, normal *vector.Vec3) (bool, *ScatterEffect) {
	reflected := reflect(vector.UnitVector(rayIn.Direction()), normal)
	scatteredRay := ray.New(hitPoint, reflected.PlusVector(vector.RandomInUnitSphere().MultiplyScaler(m.Fuzz)))
	scatteredEffect := &ScatterEffect{
		ScatteredRay: scatteredRay,
		Attenuation:  m.Albedo,
	}
	return scatteredRay.Direction().Dot(normal) > 0, scatteredEffect
}

func reflect(v, normal *vector.Vec3) *vector.Vec3 {
	return v.MinusVector(normal.MultiplyScaler(2 * v.Dot(normal)))
}
