package main

import (
	"testing"
	"math"
)

func BenchmarkVec3Normalize(b *testing.B)	{
	v := Vec3{1, 2, 3}
	for i := 0; i < b.N; i++ {
		v.Normalize()
	}
}

func TestVec3Normalize(t *testing.T) {
	v := Vec3{1, 2, 3}
	v.Normalize()
	if math.Abs(v.X-1) > 1e-2 {
		t.Errorf("v.X = %v, want %v", v.X, 1.0)
	}
}

func TestFastInvSqrt(t *testing.T) {
	x := 4.0
	y := fastInvSqrt(x)
	if math.Abs(y-0.5) > 1e-2 {
		t.Errorf("y = %v, want %v", y, 1.0)
	}
}