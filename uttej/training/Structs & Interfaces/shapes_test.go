package shape

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)

	}
}

func TestArea(t *testing.T) {

	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})

	t.Run("Circle Area", func(t *testing.T) {

		circle := Circle{10.0}

		got := circle.Area()

		want := 314.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)

		}

	})

}
