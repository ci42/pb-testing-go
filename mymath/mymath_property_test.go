// +build property

package mymath_test

import (
	"testing"
	"testing/quick"

	"example.org/mymath"
)

// func int16Generator(values []reflect.Value, r *rand.Rand) {
// 	for i := range values {
// 		//values[i] = reflect.ValueOf(int16(0))
// 		values[i] = reflect.ValueOf(int16(r.Int31()))
// 	}
// }

// var retainSignConfig = &quick.Config{
// 	MaxCount:      0,
// 	MaxCountScale: 0,
// 	Rand:          nil,
// 	Values:        int16Generator,
// }

// func float32Generator(values []reflect.Value, r *rand.Rand) {
// 	for i := range values {
// 		// values[i] = reflect.ValueOf(float32(0.0))
// 		values[i] = reflect.ValueOf(float32(r.NormFloat64()))
// 	}
// }

// var add3FloatsThreeConfig = &quick.Config{
// 	MaxCount:      0,
// 	MaxCountScale: 0,
// 	Rand:          nil,
// 	Values:        float32Generator,
// }

func TestCubeRetainsSign(t *testing.T) {
	property := func(x int64) bool {
		result := mymath.CubeInt64(x)
		if x < 0 {
			return result < 0
		}
		return result >= 0
	}

	if err := quick.Check(property, nil); err != nil {
		t.Fatalf("%s", err)
	}
}

func TestAdd3FloatsCommutative(t *testing.T) {
	f1 := func(x, y, z float32) float32 {
		return mymath.Add3Floats(x, y, z)
	}

	f2 := func(x, y, z float32) float32 {
		return mymath.Add3Floats(z, y, x)
	}

	// if err := quick.CheckEqual(f1, f2, add3FloatsThreeConfig); err != nil {
	if err := quick.CheckEqual(f1, f2, nil); err != nil {
		t.Fatalf("%s", err)
	}
}
