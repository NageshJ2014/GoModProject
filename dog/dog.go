package dog

import "fmt"

type Dog struct {
	name  string
	color string
	age   int
}

func (p Dog) Init(name string, color string, age int) {
	fmt.Println("CAME INSIDE Dog Init Function")
	p.name = name
	p.color = color
	p.age = age
	fmt.Printf(" Hi I am Your %s , And Shining with %s Color \n Of Course I am %d Old", p.name, p.color, p.age)
}

func (p Dog) Describe() {
	fmt.Printf(" Hi I am Your %s , And Shining with %s Color \n Of Course I am %d Old", p.name, p.color, p.age)

}
