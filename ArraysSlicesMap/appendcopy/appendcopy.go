package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	var slice []int

	// Esse copigo da erro pois nao pode ser um array, array é fixo
	//array = append(array, 1, 2, 3, 4)
	slice = append(array[:], 4, 5, 6, 7)
	slice = append(slice, 4, 5, 6, 7)

	fmt.Println(array, slice)

	slice2 := make([]int, 2)
	copy(slice2, slice)

	fmt.Println(slice2)

}
