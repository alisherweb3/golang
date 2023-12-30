package main

import "fmt"

func mars_age(age int) int {
	age *= 365
	age /= 687
	return age
}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}
