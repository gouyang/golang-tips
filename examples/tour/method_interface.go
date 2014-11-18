package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) Say() {
	fmt.Println("Say hello!")
}

func (p Person) SayName() {
	fmt.Printf("My name is %s\n", p.Name)
}

type Emp struct {
	Name, Title string
}

func (e Emp) Say() {
	fmt.Println("I'm employee now!")
}

func (e Emp) SayName() {
	fmt.Printf("My name is %s and my title is %s\n", e.Name, e.Title)
}

type person interface {
	Say()
	SayName()
}

func main() {
	var p person
	me := Person{"guohua"}
	p = me
	p.Say()
	p.SayName()
	e := Emp{"shao", "Boss"}
	p = e
	p.Say()
	p.SayName()
}
