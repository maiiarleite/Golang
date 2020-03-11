package main

import "log"

func main(){
	s:= make([]int, 10)
	s[9] = 12
	log.Println(s)

	s = make([]int, 10, 20)
	log.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	log.Println(s, len(s), cap(s))

	s = append(s, 1)
	log.Println(s, len(s), cap(s))

}
