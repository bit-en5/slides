package main

//START OMIT

type MyType struct {
	Field1 int `json:"field1" sql:"Field1" validate:"required,min=1,max=10"`
}

//END OMIT
