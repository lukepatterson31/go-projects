package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

func foo(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!", v)
	case string:
		fmt.Println("I'm a string!", v)
	default:
		fmt.Println("Unknown type!", v)
	}
}

func main() {
	var x = 1
	var y = "foo"
	var guy = new(Person)
	guy.Name = "Dave"
	guy.SayHello()
	Greet(guy)

	var dog = new(Dog)
	Greet(dog)

	var count = int(42)

	// can also be declared
	// ptr := &count
	var ptr *int = &count

	fmt.Println("Dereferenced ptr:", *ptr)

	*ptr = int(100)

	fmt.Println("Dereferenced ptr:", *ptr)
	fmt.Println("Variable count:", count)
	fmt.Println("Hello, Black Hat Gophers!")

	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}

	switch y {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}

	foo(guy)

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	var nums = []int{1, 2, 3, 4, 5}
	var testmap = make(map[int]int)
	for idx, val := range nums {
		fmt.Println(idx, val)
		testmap[idx] = val
	}

	for idx, val := range testmap {
		fmt.Printf("key %d, val %d\n", idx, val)
	}

}
