package polymorphism

//START1 OMIT

type Rectangle struct{ Width, Height float64 }
type Circle struct{ Radius float64 }

type Geometry interface { // HL
	Area() float64 // HL
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (c Circle) Area() float64    { return 3.14159 * c.Radius * c.Radius }

//END1 OMIT
