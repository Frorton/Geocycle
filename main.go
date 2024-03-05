package main

import (
	"fmt"
	"strconv"
)

// Vélo represents a collection of values with associated names
type Vélo map[string]int

// FillTable prompts the user to fill in values for a table
func FillTable() Vélo {
	fmt.Println("Fill in the table:")
	vélo := make(Vélo)

	// Define the names of the segments
	segmentNames := []string{"TopTube", "SeatTube", "Chainstay", "BB Height", "Wheelbase", "Head Angle", "SeatTube Angle", "Reach", "Stack"}

	for _, name := range segmentNames {
		fmt.Printf("%s : ", name)
		input, err := getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		// Convert the user input to an integer
		length, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			return nil
		}

		// Add the value to vélo
		vélo[name] = length
		fmt.Println("-------")
	}

	return vélo
}

// Vélo represents a collection of values with associated names
// Comparision is based on Vélo 1
func CompareTables(vélo1, vélo2 Vélo) {
	fmt.Println("The results are :")
	for name, value1 := range vélo1 {
		if value2, ok := vélo2[name]; ok {
			switch {
			case value1 == value2:
				fmt.Printf("%s:\n Identiqual values\n", name)
				fmt.Println("-------")
			case value1 < value2:
				diff1 := value2 - value1
				fmt.Printf("%s:\n Table 2 is %v mm longer\n", name, diff1)
				fmt.Println("-------")
			case value1 > value2:
				diff2 := value1 - value2
				fmt.Printf("%s:\n Table 2 is %v mm shorter\n", name, diff2)
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

// getUserInput reads a line of input from the user
func getUserInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}

func main() {
	// Prompt user to fill in the tables
	fmt.Println("Vélo 1:")
	vélo1 := FillTable()

	fmt.Println("Vélo 2:")
	vélo2 := FillTable()

	// Compare tables
	if vélo1 != nil && vélo2 != nil {
		CompareTables(vélo1, vélo2)
	}
}
