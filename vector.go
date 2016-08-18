package vector

import (
	"math"
)

type vec3 struct {
	e0 float32
	e1 float32
	e2 float32
}

func New(args ...float32) *vec3 {
	x, y, z := 0, 0, 0
	switch len(args) {
	case 1:
		x := args[0]
	case 2:
		x, y = args[0], args[1]
	case 3:
		x, y, z = args[0], args[1], args[2]
	}

	return &vec3{
		e0: x,
		e1: y,
		e2: z,
	}

}
func UnitVector(v *vec3) *vec3 {
	return v.Divide(v.Length())
}

func (v *vec3) MakeUnitVector() {
	k := 1.0 / (math.Sqrt(v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2))
	v.e0 *= k
	v.e1 *= k
	v.e2 *= k
}

func (v *vec3) PlusEqualVector(w *vec3) {
	v.e0 += w.e0
	v.e1 += w.e1
	v.e2 += w.e2
}

func (v *vec3) MinusEqualVector(w *vec3) {
	v.e0 -= w.e0
	v.e1 -= w.e0
	v.e2 -= w.e0
}

func (v *vec3) DivideEqualVector(w *vec3) {
	v.e0 /= w.e0
	v.e1 /= w.e1
	v.e2 /= w.e2
}

func (v *vec3) MultiplyEqualVector(w *vec3) {
	v.e0 *= w.e0
	v.e1 *= w.e1
	v.e2 *= w.e2
}

func (v *vec3) MultiplyEqual(k float32) {
	v.e0 *= k
	v.e1 *= k
	v.e2 *= k
}

func (v *vec3) DivideEqual(k float32) {
	v.e0 /= k
	v.e1 /= k
	v.e2 /= k
}

func (v *vec3) Cross(w *vec3) *vec3 {
	return &vec3{
		e0: (v.e1 * w.e2) - (v.e2 * w.e1),
		e1: -1 * ((v.e0 * w.e2) - (v.e2 * w.e0)),
		e2: (v.e0 * w.e1) - (v.e1 * w.e0),
	}
}

func (v *vec3) Dot(w *vec3) float32 {
	return v.e0*w.e0 + v.e1*w.e1 + v.e2*w.e2
}

func (v *vec3) Multiply(t float32) *vec3 {
	return &vec3{
		e0: v.e0 * t,
		e1: v.e1 * t,
		e2: v.e2 * t,
	}
}

func (v *vec3) Divide(t float32) *vec3 {
	return &vec3{
		e0: v.e0 / t,
		e1: v.e1 / t,
		e2: v.e2 / t,
	}
}

func (v *vec3) VectorPlus(w *vec3) *vec3 {
	return &vec3{
		e0: v.e0 + w.e0,
		e1: v.e1 + w.e1,
		e2: v.e2 + w.e2,
	}
}

func (v *vec3) VectorMinus(w *vec3) *vec3 {
	return &vec3{
		e0: v.e0 - w.e0,
		e1: v.e1 - w.e1,
		e2: v.e2 - w.e2,
	}
}

func (v *vec3) VectorMultiply(w *vec3) *vec3 {
	return &vec3{
		e0: v.e0 * w.e0,
		e1: v.e1 * w.e1,
		e2: v.e2 * w.e2,
	}
}

func (v *vec3) VectorDivide(w *vec3) *vec3 {
	return &vec3{
		e0: v.e0 / w.e0,
		e1: v.e1 / w.e1,
		e2: v.e2 / w.e2,
	}
}

func (v *vec3) Length() float32 {
	l = math.Sqrt(float64(v.e1*v.e1 + v.e1*v.e1 + v.e2*v.e2))
	return float32(l)
}

func (v *vec3) SquaredLength() float32 {
	return v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2
}

func (v *vec3) X() float32 {
	return v.e0
}

func (v *vec3) Y() float32 {
	return v.e1
}

func (v *vec3) Z() float32 {
	return v.e2
}

func (v *vec3) R() float32 {
	return v.e0
}

func (v *vec3) G() float32 {
	return v.e1
}

func (v *vec3) B() float32 {
	return v.e2
}
