package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hey there! We are learning slices.")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruit list is: %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 400
	highScores[2] = 200
	highScores[3] = 700
	highScores = append(highScores, 246, 535, 532)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	// How to remove a value from slices based on index
	courses := []string{"React", "Javascript", "Swift", "Python"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
