package fungsi_test

import (
	"reflect"
	"testing"

	"01group/fungsi"
)

func TestMap(t *testing.T) {
	input, expectedResult := []int32{2, 4, 8, 16}, []int32{4, 8, 16, 32}
	result := fungsi.Map(input, func(v int32) int32 {
		return v * 2
	})

	if reflect.DeepEqual(result, expectedResult) != true {
		t.Error("Slice not the same")
	}
}
