// ❗ Does not yet support angles like seat tube angle etc ❗
package Comparing

import (
	"Geocycle/Filling"
	"fmt"
)

// Comparision is based on vélo1
func CompareTables(vélo1, vélo2 Filling.Vélo) {

	fmt.Println("")
	fmt.Println("The results are :")
	fmt.Println("")

	for name, value1 := range vélo1 {
		if value2, ok := vélo2[name]; ok {
			switch {
			case value1 == value2:
				fmt.Printf("%s:\n Same lenght\n", name)
				fmt.Println("-------")
			case value1 < value2:
				diff1 := value2 - value1
				fmt.Printf("%s:\n Table 1 is %v mm shorter\n", name, diff1)
				fmt.Println("-------")
			case value1 > value2:
				diff2 := value1 - value2
				fmt.Printf("%s:\n Table 1 is %v mm longer\n", name, diff2)
				fmt.Println("-------")
			}
		} else {
			fmt.Printf("%s: Name not found in Table 2\n", name)
		}
	}

	// Checking for names in Vélo 2 that are not in Vélo 1
	for name := range vélo2 {
		if _, ok := vélo1[name]; !ok {
			fmt.Printf("%s: Name not found in Table 1\n", name)
		}
	}
}
