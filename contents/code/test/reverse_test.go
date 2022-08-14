package main

import (
	"fmt"
	"testing"
)

//START1 OMIT

func TestReverse(t *testing.T) {
	if Reverse("Hello") != "olleH" {
		t.Error("reverse goes wrong")
	}
}

//END1 OMIT
//START2 OMIT

func TestReverse2(t *testing.T) {
	s := `Hello 🙂`
	want := `🙂 olleH`
	has := Reverse(s)

	if has != want {
		t.Errorf("want %s / has %s", want, has)
	}
}

//END2 OMIT
//START3 OMIT

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Reverse("Hello")
	}
}

//END3 OMIT
//START4 OMIT

func ExampleReverse() {
	fmt.Println(Reverse("Hello"))
	// Output: olleH
}

//END4 OMIT
