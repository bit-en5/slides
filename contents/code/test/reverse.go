package main

func Reverse(s string) (res string) {
	for _, r := range s {
		res = string(r) + res
	}

	return
}
