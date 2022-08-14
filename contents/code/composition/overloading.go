package main

//START1 OMIT

func (u User) GetName() string { // HL
	return u.Name
}

func (g Gamer) GetName() string { // HL
	return g.NickName
}

//END1 OMIT
