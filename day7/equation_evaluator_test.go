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
			ce:   CalibrationEquation{190, []uint64{10, 19}},
			want: true,
		},
		{
			ce:   CalibrationEquation{83, []uint64{17, 5}},
			want: false,
		},
		{
			ce:   CalibrationEquation{3267, []uint64{81, 40, 27}},
			want: true,
		},
		{
			ce:   CalibrationEquation{292, []uint64{11, 6, 16, 20}},
			want: true,
		},
		{
			ce:   CalibrationEquation{7290, []uint64{6, 8, 6, 15}},
			want: false,
		},
		{
			ce:   CalibrationEquation{35, []uint64{6, 8, 6, 15}},
			want: true,
		},
		{
			ce:   CalibrationEquation{4320, []uint64{6, 8, 6, 15}},
			want: true,
		},
		{
			ce:   CalibrationEquation{479027832, []uint64{8, 9, 69, 659, 96, 634}},
			want: false,
		},
		{
			ce:   CalibrationEquation{12524863130276, []uint64{5, 30, 6, 928, 6, 2, 969, 73, 5}},
			want: false,
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
