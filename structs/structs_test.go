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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := rectangle{
			width:  2.0,
			height: 3.0,
		}

		checkArea(t, rect, 6.0)
	})

	t.Run("circles", func(t *testing.T) {
		cir := circle{
			radius: 2.0,
		}

		checkArea(t, cir, 12.56)
	})
}
