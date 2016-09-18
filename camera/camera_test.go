package camera

import (
	"github.com/faizanrashid/GoRayTracer/vector"

	"testing"
)

func vectorEqual(u, v *vector.Vec3) bool {
	return u.X() == v.X() && u.Y() == v.Y() && u.Z() == v.Z()
}

func TestNewCamera(t *testing.T) {
	origin := vector.New()
	lowerLeftCorner := vector.New(1, 2, 3)
	horizontal := vector.New(3, 2, 1)
	vertical := vector.New(1, 3, 5)
	c := New(origin, lowerLeftCorner, horizontal, vertical)
	if c.Origin != origin {
		t.Error("Origin vector is incorrect.")
	}
	if c.LowerLeftCorner != lowerLeftCorner {
		t.Error("Lower left corner vector is incorrect.")
	}
	if c.Horizontal != horizontal {
		t.Error("Horizontal vector is incorrect.")
	}
	if c.Vertical != vertical {
		t.Error("Vertical vector is incorrect.")
	}
}

func TestGetRay(t *testing.T) {
	origin := vector.New()
	lowerLeftCorner := vector.New(1, 2, 3)
	horizontal := vector.New(3, 2, 1)
	vertical := vector.New(1, 3, 5)
	c := New(origin, lowerLeftCorner, horizontal, vertical)
	r := c.GetRay(5, 6)
	expected := vector.New()
	if !vectorEqual(r.Origin(), expected) {
		t.Errorf("New ray origin is incorrect. Expected %v \n Actual %v", expected, r.Origin())
	}
	expected = vector.New(22, 30, 38)
	if !vectorEqual(r.Direction(), expected) {
		t.Errorf("New ray origin is incorrect. Expected %v \n Actual %v", expected, r.Direction())
	}
}
