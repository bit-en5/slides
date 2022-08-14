package main

//START1 OMIT

type User struct {
	Name string
	Age  int
}

type Gamer struct {
	User     // HL
	NickName string
}

//END1 OMIT
