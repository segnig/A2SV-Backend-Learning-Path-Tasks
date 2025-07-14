package main

// main is the entry point for the Student Grade Calculator.
// It collects input for multiple students, calculates their average grades, and displays them.
func main() {
	students := make(map[string]map[string]float64)
	numberStudents := getIntInput("Enter number of students: ")

	for i := 0; i < numberStudents; i++ {
		name, subjects := getStudentData()
		students[name] = subjects
	}

	studentsAverage := calculateAverage(students)
	displayToUser(studentsAverage)
}

// calculateAverage computes the average grade for each student.
func calculateAverage(students map[string]map[string]float64) map[string]float64 {
	studentsAverage := make(map[string]float64)

	for name, subjects := range students {
		var total float64
		for _, grade := range subjects {
			total += grade
		}
		studentsAverage[name] = total / float64(len(subjects))
	}

	return studentsAverage
}
