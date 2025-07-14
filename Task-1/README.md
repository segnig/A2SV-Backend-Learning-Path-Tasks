# 🎓 Student Grade Calculator (Go Console Application)

This is a simple Go console application that allows students to calculate their average grades based on subjects they've taken. It supports multiple students and outputs a report of average scores.

---

## 🧾 Features

- Accepts input for multiple students.
- Prompts for each subject name and grade.
- Calculates and displays average grades per student.
- Input validation ensures clean, safe data (e.g., grade must be between 0–100).

---

## 🚀 Getting Started

### 🔧 Requirements

- [Go 1.20+](https://golang.org/dl/) installed on your system.

### 📦 Installation & Running

1. **Clone the repository** (if applicable) or download the `.go` files:

   ```bash
   git clone https://github.com/yourusername/student-grade-calculator.git
   cd student-grade-calculator
````

2. **Run the application**:

   ```bash
   go run main.go user_input.go display_to_user.go
   ```

---

## 🗂️ Project Structure

```bash
.
├── main.go                # Entry point for the application
├── user_input.go          # Functions for handling user input
├── display_to_user.go     # Function to display the report
└── README.md              # This file
```

---

## 💡 Example Usage

```bash
Enter number of students: 2

=======================================================
Enter user name: Lime
Enter the number of subjects: 2

Subject 1:
Enter subject name: Math
Enter grade (0-100): 90

Subject 2:
Enter subject name: English
Enter grade (0-100): 85
=======================================================

--- All Students Average Grades ---
Lemi      : 87.50
```

---

## 📘 Documentation

* `getStringInput(prompt string)`: Gets non-empty string input.
* `getFloatInput(prompt string, min, max float64)`: Gets float input within range.
* `getIntInput(prompt string)`: Gets positive integer.
* `getStudentData()`: Gathers name and subject/grade entries.
* `calculateAverage(map[string]map[string]float64)`: Returns student → average map.
* `displayToUser(map[string]float64)`: Displays nicely formatted output.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## 👩‍💻 Author

Developed by **Segni Girma**
[GitHub Profile](https://github.com/segnig)