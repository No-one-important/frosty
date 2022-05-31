package main

import (
	"encoding/json"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{
		X: v.X + w.X,
		Y: v.Y + w.Y,
		Z: v.Z + w.Z,
	}
}

func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{
		X: v.X - w.X,
		Y: v.Y - w.Y,
		Z: v.Z - w.Z,
	}
}

// Div divides v by a scalar.
func (v Vec3) Div(x float64) Vec3 {
	return Vec3{
		X: v.X / x,
		Y: v.Y / x,
		Z: v.Z / x,
	}
}

// Mul multiplies v by a scalar.
func (v Vec3) Mul(x float64) Vec3 {
	return Vec3{
		X: v.X * x,
		Y: v.Y * x,
		Z: v.Z * x,
	}
}

// Mag returns the magnitude of v.
func (v Vec3) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns the normalized (unit) vector of v.
func (v Vec3) Normalize() Vec3 {
	x := fastInvSqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

	return Vec3{
		X: v.X * x,
		Y: v.Y * x,
		Z: v.Z * x,
	}
}

// Cross returns the cross product of v and w.
func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		X: v.Y*w.Z - v.Z*w.Y,
		Y: v.Z*w.X - v.X*w.Z,
		Z: v.X*w.Y - v.Y*w.X,
	}
}

// Dot returns the dot product of v and w.
func (v Vec3) Dot(w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v *Vec3) UnmarshalJSON(b []byte) error {
	a := [3]float64{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	v.X, v.Y, v.Z = a[0], a[1], a[2]
	return nil
}

func fastInvSqrt(n float64) float64 {
	n2, th := n*0.5, float64(1.5)
    b := math.Float64bits(n)
    b = 0x5FE6EB50C7B537A9 - (b >> 1)
    f := math.Float64frombits(b)
    f *= th - (n2 * f * f)
    return f
}