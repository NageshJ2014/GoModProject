package main

import (
	"fmt"
	"rest/dog"
	"rest/rest"
	// "./dog/dog"
)

func main() {
	fmt.Println("This is Testing")
	rest.Rest_test()
	p2 := rest.Person{"Nagesh Jayaram"}
	p2.SayHello()
	p2.First = "Nagesh"
	p2.SayHello()

	// p1 := person{"Nagesh"}
	// p1.sayHello()
	rest.Another_rest()

	d1 := dog.Dog{}
	d1.Init("Ruby", "Brown", 5)
	d1.Describe()

}
