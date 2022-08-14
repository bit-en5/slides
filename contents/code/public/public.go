package main

//START OMIT

// Functions // HL
func Test() {} // → public
func test() {} // → private

// Structs // HL
type Partner struct{} // → public
type partner struct{} // → private

// Properties // HL
type User struct {
	Name     string // → public
	nickName string // → private
}

//END OMIT
