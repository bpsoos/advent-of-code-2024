package day7

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	var tests = []struct {
		ce   CalibrationEquation
		want bool
	}{
		{
			ce:   CalibrationEquation{190, []int{10, 19}},
			want: true,
		},
		{
			ce:   CalibrationEquation{83, []int{17, 5}},
			want: false,
		},
		{
			ce:   CalibrationEquation{3267, []int{81, 40, 27}},
			want: true,
		},
		{
			ce:   CalibrationEquation{292, []int{11, 6, 16, 20}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.ce, tt.want), func(t *testing.T) {
			e := EquationEvaluator{}
			valid := e.Evaluate(&tt.ce)
			if valid != tt.want {
				t.Fatalf("result (%v) does not equal expected (%v) for equation %v", valid, tt.want, tt.ce)
			}
		})
	}
}
