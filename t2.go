package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть ім'я студента: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Видаляємо зайві пробіли

	fmt.Print("Введіть оцінки, розділені пробілами: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Видаляємо зайві пробіли

	if input == "" {
		fmt.Println("Ви не ввели оцінки. Будь ласка, спробуйте ще раз.")
		return
	}

	strGrades := strings.Fields(input)
	grades := make([]float64, len(strGrades))

	for i, strGrade := range strGrades {
		grade, err := strconv.ParseFloat(strGrade, 64)
		if err != nil {
			fmt.Println("Невірний ввід. Будь ласка, введіть дійсні оцінки.")
			return
		}
		grades[i] = grade
	}

	student := Student{Name: name, Grades: grades}
	fmt.Printf("Середня оцінка для %s: %.2f\n", student.Name, student.AverageGrade())
}
