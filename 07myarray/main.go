package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Fruitlist is ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is ", len(vegList))
}
