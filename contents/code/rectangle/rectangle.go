package rectangle

//START1 OMIT

type Rectangle struct {
	Width  float64
	Height float64
}

//END1 OMIT
//START2 OMIT

// SetWidth is a public method
// Because it's a setter, it needs the *
func (r *Rectangle) SetWidth(w float64) { // HL
	r.Width = w
}

// GetWidth is a public method
// Because it's a getter, it doesn't needs the *
func (r Rectangle) GetWidth() float64 { // HL
	return r.Width
}

// getWidth is a private method
func (r Rectangle) getWidth() float64 { // HL
	return r.Height
}

//END2 OMIT
