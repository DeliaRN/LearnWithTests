package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaWithoutRefactoring(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		//By declaring the interface, the helper is decoupled
		//the concrete types, and only focuses on the method it needs.
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	//using the function
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	//without using the function
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func RefactoredTestArea(t *testing.T) {
	//USING ANONYMOUS STRUCTS
	areaTests := []struct {
		//name String --> we can add it
		shape  Shape
		areaOf float64 //the "want"
	}{
		{Rectangle{12, 6}, 72.0},                                                     //sahpe properties & want
		{Circle{10}, 314.1592653589793},                                              //shape props. & want
		{ /*name: "Triangle", */ shape: Triangle{Base: 12, Height: 6}, areaOf: 36.0}, //named
	}
	for _, test := range areaTests { //for each test in areaTests
		got := test.shape.Area()
		if got != test.areaOf {
			t.Errorf("got %g want %g", got, test.areaOf)
		}
	}
}
