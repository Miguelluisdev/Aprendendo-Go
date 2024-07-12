package main

import "fmt"

func findMax(arr []int) int {
	if len(arr) == 0 {
		panic("o array não pode ser vazio")
	}
	min, max := arr[0], arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return max
}

func findMin(arr []int) int {
	if len(arr) == 0 {
		panic("o array não pode ser vazio")
	}
	min, max := arr[0], arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	 max:=findMax(arr)
	 min:= findMin(arr)

	fmt.Println("Maior numero" , max)
	fmt.Println("Maior numero" , min)
}
