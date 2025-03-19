package main

import "fmt"

type Address4 struct {
	Region string
	Street string
	No     string
}

type User7 struct {
	ID   int
	Name string
	Addr Address4
}

func change(u User7) {
	u.Name = "xxxx"
}

func changePoint(u *User7) {
	u.Name = "yyyy"
}

func vars() {
	me := User7{}
	me2 := me
	me2.Name = "kk"
	fmt.Printf("%#v\n", me.Name)  // ""
	fmt.Printf("%#v\n", me2.Name) // "kk"

	change(me2)
	fmt.Printf("%#v\n", me2) //   Name:"kk"

	changePoint(&me2)
	fmt.Printf("%#v\n", me2) //   Name:"yyyy"
}
