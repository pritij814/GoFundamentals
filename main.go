package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name..")
	fmt.Scanln(&name)

	fmt.Print("Enter Your age:")
	fmt.Scanln(&age)

	fmt.Println("Hello ", name, " your age is ", age)
}

//create 2 file
//in 1st file call 2nd function
//
