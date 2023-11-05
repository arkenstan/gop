package main

import "fmt"

func slices() {

	fmt.Println("========== SLICES")

	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "s"
	s[1] = "h"
	s[2] = "e"
	s = append(s, "t")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	l := s[1:len(s)]
	fmt.Println("emp:", l, "len:", len(l), "cap:", cap(l))

}
