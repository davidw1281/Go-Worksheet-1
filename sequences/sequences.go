package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i := range slice {
		slice[i] = f(slice[i])
	}
}

func mapArray(f func(a int) int, array [5]int) [5]int {
	for i := range array {
		array[i] = f(array[i])
	}
	return array
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [5]int{1, 2, 3, 4, 5}
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)

	/*
		nums := [5]int{1, 2, 3, 4, 5}
		numsSlice := nums[1:4]
		numsSlice2 := numsSlice[1:2]
		numsSlice2[0] = 100
		fmt.Println(nums, numsSlice, numsSlice2)
	*/

	intsSlice = []int{2, 3, 4, 5, 6}
	newSlice := intsSlice[1:3]
	for i := range newSlice {
		newSlice[i] = square(newSlice[i])
	}
	fmt.Println(intsSlice, newSlice)

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)

}
