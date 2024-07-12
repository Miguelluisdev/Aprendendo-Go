package main

import "fmt"

// Escreva um algoritmo que dado um vetor qualquer (exemplo [5,4,3,2,1]) retorne o vetor ordenado (exemplo [1,2,3,4,5]) sem utilizar funções tipo sort.

func main() {

	arr:= []int{10,9,8,7,6,5,4,3,2,1}

	sortArray(arr)

	fmt.Println("array ordenado" , arr)
}

func sortArray(arr []int) {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		for j:= 0; j < n-i-1 ; j++ {
			if arr[j] > arr[j+1]{
				arr[j] , arr[j+1] = arr[j+1] , arr[j]
			}
		}
	}
}
