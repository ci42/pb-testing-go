package mymath_test

import (
	"testing"

	"example.org/mymath"
)

func TestCubeRetainsSign(t *testing.T) {

	testcases := []struct {
		name   string
		in     int64
		result int64
	}{
		{name: "3**3", in: 3, result: 27},
		{name: "4**3", in: 4, result: 64},
		{name: "5**3", in: 5, result: 125},
		{name: "-5**3", in: -5, result: -125},
		{name: "2000**3", in: 2000, result: 8000000000},
		{name: "-4000**3", in: -4000, result: -64000000000},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			testcase := testcase
			t.Parallel()

			if r := mymath.CubeInt64(testcase.in); r != testcase.result {
				t.Fatalf("Actual r: %d, Wanted: %d", r, testcase.result)
			}
		})
	}
}

func TestAdd3FloatsCommutative(t *testing.T) {

	testcases := []struct {
		name string
		x    float32
		y    float32
		z    float32
	}{
		{name: "1.2+1.4+1.6", x: 1.2, y: 1.4, z: 1.6},
		{name: "1.246+1.482+1.622", x: 1.246, y: 1.482, z: 1.622},
		{name: "(-1.2)+(1.4)+(-1.6)", x: -1.2, y: 1.4, z: -1.6},
		{name: "-1.246+15.48232+0.62211", x: 1.246, y: 15.48232, z: 0.62211},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			testcase := testcase
			t.Parallel()

			r1 := mymath.Add3Floats(testcase.x, testcase.y, testcase.z)
			r2 := mymath.Add3Floats(testcase.z, testcase.y, testcase.x)
			if r1 != r2 {
				t.Fatalf("r1: %e, r2: %e", r1, r2)
			}
		})
	}
}
