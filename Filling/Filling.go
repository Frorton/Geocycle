package Filling

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

// getUserInput reads a line of input from the user
func getUserInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}
