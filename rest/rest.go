package rest

import "fmt"

// Declaring a A Strcut with name person having first as string attribute

type Person struct {
	First string
}

func (p Person) SayHello() {

	// fmt.Println("Hi My Name is ", p.First)
	length := len(p.First)
	fmt.Println("Hi My Name is ", p.First, "And The Length = ", length)

}

func Rest_test() {
	fmt.Println("This is Coming From Rest Module")
}
