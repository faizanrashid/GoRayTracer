package hitable

import (
	"github.com/faizanrashid/GoRayTracer/material"
	"github.com/faizanrashid/GoRayTracer/ray"
	"github.com/faizanrashid/GoRayTracer/vector"
)

type HitRecord struct {
	T        float64
	P        *vector.Vec3
	Normal   *vector.Vec3
	Material material.Material
}

type Hitable interface {
	Hit(r *ray.Ray, hitRecord *HitRecord, tMin, tMax float64) bool
	Material() material.Material
}

type HitableList struct {
	hitables []Hitable
}

func (htl *HitableList) Add(hitable Hitable) {
	htl.hitables = append(htl.hitables, hitable)
}

func (htl *HitableList) Hit(r *ray.Ray, hitRecord *HitRecord, tMin, tMax float64) bool {
	hitAnything := false
	closestSoFar := tMax
	tempRec := &HitRecord{}
	for _, hitable := range htl.hitables {
		if hitable.Hit(r, tempRec, tMin, closestSoFar) {
			hitAnything = true
			closestSoFar = tempRec.T
			*hitRecord = *tempRec
		}
	}
	return hitAnything
}
