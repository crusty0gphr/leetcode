package zigzagconvert

import (
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s      string
		numRow int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{"PAYPALISHIRING", 3},
			"PAHNAPLSIIGYIR",
		},
		{
			"Test 2",
			args{"PAYPALISHIRING", 4},
			"PINALSIGYAHRPI",
		},
		{
			"Test 3",
			args{"A", 3},
			"A",
		},
		{
			"Test 4",
			args{"AB", 1},
			"AB",
		},
		{
			"Test 5",
			args{"", 0},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRow); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertMatrix(tt.args.s, tt.args.numRow); got != tt.want {
				t.Errorf("convertMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
