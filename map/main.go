package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Alice"] = 23
	m["Bob"] = 34
	m["Charlie"] = 29

	fmt.Println("idade de bob", m["Alice"])

	if age, exists := m["Charlie"]; exists {
		fmt.Println("Idade de Charlie:", age)
	} else {
		fmt.Println("Charlie não está no map")
	}

	delete(m, "Bob")

	if age, exists := m["Bob"]; exists {
        fmt.Println("Idade de Bob:", age)
    } else {
        fmt.Println("Bob não está no map")
    }

    // Iterando sobre o map
    for key, value := range m {
        fmt.Println("Chave:", key, "Valor:", value)
    }
}
