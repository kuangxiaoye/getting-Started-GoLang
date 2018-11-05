package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) Sayhi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (s *Human) Sayhi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", s.name, s.phone)

}

func main() {
	mark := Student{
		Human{
			"Mark", 25, "222-222-YYYY",
		},
		"yizhong",
	}

	sam := Employee{
		Human{
			"Sam", 45, "111-888-XXXX",
		},
		"Golang Inc",
	}
	mark.Sayhi()
	sam.Sayhi()
}
