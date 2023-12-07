package main

import "testing"

func rectangleArea(w, h int) int {
	return w * h
}

type input struct{ w, h int }

type testcase struct {
	in   input
	want int
}

func Test_rectangleArea(t *testing.T) {
	testcases := map[string]testcase{
		"ok": {
			in:   input{w: 3, h: 4},
			want: 12,
		},
	}

	for title, tc := range testcases {
		t.Run(title, func(t *testing.T) {
			if got := rectangleArea(tc.in.w, tc.in.h); got != tc.want {
				t.Errorf("w=%d, h=%d, want=%d but got=%d", tc.in.w, tc.in.h, tc.want, got)
			}
		})
	}
}
