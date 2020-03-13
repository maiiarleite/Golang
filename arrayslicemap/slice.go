package main

import (
"log"
"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	log.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3]
	log.Println(a2, s2)

	s3 := a2[:0]
	log.Println(a2, s3)

	s4 := s2[:1]
	log.Println(s2, s4)
}
