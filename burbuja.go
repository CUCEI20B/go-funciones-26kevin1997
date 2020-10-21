package main

import "fmt"

func burbuja(s []int64) {
	fmt.Println("No ordenada")
	for a := 0; a < len(s); a++ {
		fmt.Print(s[a], " ")
	}
	var aux int64
	for i := 1; i < len(s); i++ {
		for j := 0; j < (len(s) - i); j++ {
			if s[j] > s[j+1] {
				aux = s[j]
				s[j] = s[j+1]
				s[j+1] = aux
			}
		}
	}
	fmt.Println()
	fmt.Println("Ordenada")
	for a := 0; a < len(s); a++ {
		fmt.Print(s[a], " ")
	}
}

func main() {
	s := []int64{8, 2, 4, 5, 7, 1, 4, 8, 6, 0}
	burbuja(s)
}
