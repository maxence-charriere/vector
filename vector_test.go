package vector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVectorLength(t *testing.T) {
	v := MakeVector(2, 0, 0)
	require.Equal(t, 2.0, v.Length())
}

func TestVectorEqual(t *testing.T) {
	utests := []struct {
		scenario string
		a        Vector
		b        Vector
		equals   bool
	}{
		{
			scenario: "vectors with same origin are equals",
			a:        MakeVector(2, 2, 2),
			b:        MakeVector(2, 2, 2),
			equals:   true,
		},
		{
			scenario: "vectors with different origin are equals",
			a:        MakeVector(2, 2, 2),
			b:        MakeVectorWithOrigin(2, 2, -2, 4, 4, 0),
			equals:   true,
		},
		{
			scenario: "vectors are no equals",
			a:        MakeVector(2, 2, 2),
			b:        MakeVector(2, -2, 2),
			equals:   false,
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			require.Equal(t, u.equals, u.a.Equal(u.b))
		})
	}
}

func TestVectorAdd(t *testing.T) {
	utests := []struct {
		scenario string
		a        Vector
		b        Vector
		result   Vector
	}{
		{
			scenario: "adding a vector without length",
			a:        MakeVector(2, 2, 2),
			b:        Vector{},
			result:   MakeVector(2, 2, 2),
		},
		{
			scenario: "adding a vector with positive coordinates",
			a:        MakeVector(8, 13, 2),
			b:        MakeVector(26, 7, 40),
			result:   MakeVector(34, 20, 42),
		},
		{
			scenario: "adding a vector with negative coordinates",
			a:        MakeVector(2, 2, 2),
			b:        MakeVector(2, -2, 2),
			result:   MakeVector(4, 0, 4),
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			require.Equal(t, u.result, u.a.Add(u.b))
		})
	}
}

func TestVectorSub(t *testing.T) {
	utests := []struct {
		scenario string
		a        Vector
		b        Vector
		result   Vector
	}{
		{
			scenario: "substracting a vector with length by a vector without length",
			a:        MakeVector(2, 2, 2),
			b:        Vector{},
			result:   MakeVector(2, 2, 2),
		},
		{
			scenario: "substracting a vector without length by a vector with length",
			a:        Vector{},
			b:        MakeVector(2, 2, 2),
			result:   MakeVector(-2, -2, -2),
		},
		{
			scenario: "substracting a vector with positive coordinates",
			a:        MakeVector(12, 2, 2),
			b:        MakeVector(4, 5, 2),
			result:   MakeVector(8, -3, 0),
		},
		{
			scenario: "substracting a vector with negative coordinates",
			a:        MakeVector(12, 2, 2),
			b:        MakeVector(4, -5, -2),
			result:   MakeVector(8, 7, 4),
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			require.Equal(t, u.result, u.a.Sub(u.b))
		})
	}
}

func TestVectorScalarProduct(t *testing.T) {
	utests := []struct {
		scenario string
		a        Vector
		n        float64
		result   Vector
	}{
		{
			scenario: "multiplying a vector by 0",
			a:        MakeVector(2, 2, 2),
			n:        0,
			result:   Vector{},
		},
		{
			scenario: "multiplying a vector by positive number",
			a:        MakeVector(2, 2, 2),
			n:        2,
			result:   MakeVector(4, 4, 4),
		},
		{
			scenario: "multiplying a vector by negative number",
			a:        MakeVector(2, 2, -2),
			n:        -2,
			result:   MakeVector(-4, -4, 4),
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			require.Equal(t, u.result, u.a.ScalarProduct(u.n))
		})
	}
}

func TestVectorDotProduct(t *testing.T) {
	utests := []struct {
		scenario string
		a        Vector
		b        Vector
		result   float64
	}{
		{
			scenario: "product from vectors with positive coodinates",
			a:        MakeVector(4, 8, 10),
			b:        MakeVector(9, 2, 7),
			result:   122,
		},
		{
			scenario: "product from vectors with negative coodinates",
			a:        MakeVector(-6, 8, 0),
			b:        MakeVector(5, 12, 42),
			result:   66,
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			require.Equal(t, u.result, u.a.DotProduct(u.b))
		})
	}
}

func TestVectorCrossProduct(t *testing.T) {
	utests := []struct {
		scenario string
		a        Vector
		b        Vector
		result   Vector
	}{
		{
			scenario: "product from vectors with positive coodinates",
			a:        MakeVector(2, 3, 4),
			b:        MakeVector(5, 6, 7),
			result:   MakeVector(-3, 6, -3),
		},
		{
			scenario: "product from vectors with negative coodinates",
			a:        MakeVector(2, -3, 4),
			b:        MakeVector(5, 6, 7),
			result:   MakeVector(-45, 6, 27),
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			require.Equal(t, u.result, u.a.CrossProduct(u.b))
		})
	}
}
