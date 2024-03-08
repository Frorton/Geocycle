package main

import (
	"Geocycle/Comparing"
	"Geocycle/Filling"
	"fmt"
)

func main() {
	fmt.Println("Vélo 1")
	fmt.Println("")
	vélo1 := Filling.FillTable()

	fmt.Println("Vélo 2")
	fmt.Println("")
	vélo2 := Filling.FillTable()

	// Compare tables
	if vélo1 != nil && vélo2 != nil {
		Comparing.CompareTables(vélo1, vélo2)
	}
}
