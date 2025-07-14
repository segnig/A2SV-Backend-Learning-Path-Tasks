package main

import "fmt"

// displayToUser prints the average grades of all students in a formatted list.
func displayToUser(students map[string]float64) {
	fmt.Println("\n--- All Students Average Grades ---")
	for name, avg := range students {
		fmt.Printf("%-10s : %.2f\n", name, avg)
	}
}
