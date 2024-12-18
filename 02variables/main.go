package main

import "fmt"

func main() {
	var username string = "Neel"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float64 = 255.4554634
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Defult values and some aliases
	var variable1 int
	fmt.Println(variable1)
	fmt.Printf("Variable is of type: %T \n", variable1)

	var variable2 string
	fmt.Println(variable2)
	fmt.Printf("Variable is of type: %T \n", variable2)

	// Implicit types
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// No var style
	students := 180.45
	fmt.Println(students)
	fmt.Printf("Variable is of type: %T \n", students)
}
