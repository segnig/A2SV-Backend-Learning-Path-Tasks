package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// getStringInput prompts the user for a non-empty string input.
func getStringInput(prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			return input
		}
		fmt.Println("Input cannot be empty. Please try again.")
	}
}

// getFloatInput prompts the user for a float input within the specified range.
func getFloatInput(prompt string, min, max float64) float64 {
	for {
		input := getStringInput(prompt)
		val, err := strconv.ParseFloat(input, 64)
		if err == nil && val >= min && val <= max {
			return val
		}
		fmt.Printf("Please enter a number between %.2f and %.2f.\n", min, max)
	}
}

// getIntInput prompts the user for a positive integer.
func getIntInput(prompt string) int {
	for {
		input := getStringInput(prompt)
		n, err := strconv.Atoi(input)
		if err == nil && n > 0 {
			return n
		}
		fmt.Println("Please enter a valid positive integer.")
	}
}

// getStudentData collects a student's name and their subject grades.
func getStudentData() (string, map[string]float64) {
	fmt.Println("=======================================================")

	name := getStringInput("Enter user name: ")
	numberSubjects := getIntInput("Enter the number of subjects: ")

	subjects := make(map[string]float64)

	for i := 0; i < numberSubjects; i++ {
		fmt.Printf("\nSubject %d:\n", i+1)
		subjectName := getStringInput("Enter subject name: ")
		grade := getFloatInput("Enter grade (0-100): ", 0, 100)
		subjects[subjectName] = grade
	}

	fmt.Println("=======================================================")
	return name, subjects
}
