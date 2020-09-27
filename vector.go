package vector

import "math"

// Vector represents a Euclidean bound vector that is represented in cartesian
// space with an origin point defined to (0, 0, 0).
type Vector struct {
	X float64
	Y float64
	Z float64
}

// MakeVector creates a vector with the given coordinates that represents the
// end of the vector in a cartesian space. The origin is assumed to be (0, 0, 0),
func MakeVector(x, y, z float64) Vector {
	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

// MakeVectorWithOrigin creates a vector by specifying its origin o and its end
// e.
func MakeVectorWithOrigin(ox, oy, oz, ex, ey, ez float64) Vector {
	return Vector{
		X: ex - ox,
		Y: ey - oy,
		Z: ez - oz,
	}
}

// Length return the length (or magnitude) of the vector.
func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Equal reports whether vectors a and b are equal.
func (a Vector) Equal(b Vector) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

// Add adds vector b to vector a.
func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// Sub substract vector b from vector a.
func (a Vector) Sub(b Vector) Vector {
	return a.Add(b.ScalarProduct(-1))
}

// ScalarProduct multiplies the vector by the given number.
func (a Vector) ScalarProduct(n float64) Vector {
	return Vector{
		X: a.X * n,
		Y: a.Y * n,
		Z: a.Z * n,
	}
}

// DotProduct returns the dot product of vectors a and b.
func (a Vector) DotProduct(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// CrossProduct returns the cross product of vectors a and b.
func (a Vector) CrossProduct(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}
