package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent

	ikechukwu := User{"Ikechukwu", "Ikechukwu@go.dev", true, 23}
	fmt.Println(ikechukwu)
	fmt.Printf("ikechukwu details are: %+v\n", ikechukwu)
	fmt.Printf("Name is %v and email is %v", ikechukwu.Name, ikechukwu.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
