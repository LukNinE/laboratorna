package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть ваш масив чисел, розділених пробілами: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Видаляємо зайві пробіли

	if input == "" {
		fmt.Println("Ви не ввели жодного числа. Будь ласка, спробуйте ще раз.")
		return
	}

	// Розділяємо введений рядок на частини
	strNumbers := strings.Fields(input)
	arr := make([]int, len(strNumbers))

	// Перетворюємо рядки в цілі числа
	for i, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Невірний ввід. Будь ласка, введіть тільки цілі числа.")
			return
		}
		arr[i] = num
	}

	fmt.Println("Оригінальний масив:", arr)
	selectionSort(arr)
	fmt.Println("Відсортований масив:", arr)
}
