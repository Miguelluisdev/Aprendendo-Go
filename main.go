package main

import "fmt"

type Person struct {
	name       string
}

func main() {
	primitivo := Person{
		name: "darois",
	}


	fmt.Printf("primitivo = %+v\n", primitivo)
	printName(primitivo)
}

func printName(p Person){
	print(p.name)
}