package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := rectangle{
		width:  10.0,
		height: 10.0,
	}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: rectangle{width: 2.0, height: 3.0},
			want:  6.0,
		},

		{
			name:  "Circle",
			shape: circle{radius: 2.0},
			want:  12.56,
		},
		{
			name:  "Triangle",
			shape: triangle{base: 4.0, height: 3.0},
			want:  6.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want

			if got != want {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, want)
			}
		})
	}
}
