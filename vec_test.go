package main

import (
	"testing"
)

func BenchmarkVec3Mag(b *testing.B)	{
	v := Vec3{1, 2, 3}
	for i := 0; i < b.N; i++ {
		v.Mag()
	}
}