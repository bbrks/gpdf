package document

import "testing"

func TestSizeLandscape(t *testing.T) {
	tests := []struct {
		name string
		in   Size
		want Size
	}{
		{"portrait A4 swaps", A4, Size{Width: A4.Height, Height: A4.Width}},
		{"portrait Letter swaps", Letter, Size{Width: Letter.Height, Height: Letter.Width}},
		{"already landscape unchanged", Size{Width: 800, Height: 600}, Size{Width: 800, Height: 600}},
		{"square unchanged", Size{Width: 500, Height: 500}, Size{Width: 500, Height: 500}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.in.Landscape()
			if got != tc.want {
				t.Errorf("Landscape() = %+v, want %+v", got, tc.want)
			}
		})
	}
}

