package main

//START1 OMIT

func New() *Person {
	return &Person{}
	// return new(Person) → would be identical
}

//END1 OMIT
//START2 OMIT

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

//END2 OMIT
