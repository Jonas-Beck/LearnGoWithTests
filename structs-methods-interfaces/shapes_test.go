package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := Perimeter(rectangle)
	want := float64(40)

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		Shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{Width: 12, Height: 6}, 72},
		{"Circle", Circle{Radius: 10}, 314.1592653589793},
		{"Triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.Shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt, got, tt.want)
			}
		})
	}
}
