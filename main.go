package main

import "fmt"

func main() {

	var size int
	fmt.Println("Enter the range")
	fmt.Scan(&size)
	fmt.Println("The odd no in the range is")
	for i := 1; i <= size; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
