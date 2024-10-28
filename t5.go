package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverseSlice(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}

func main() {
	fmt.Println("Введіть ваш масив чисел, розділених пробілами:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Зчитує рядок до нового рядка

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
	reversed := reverseSlice(arr)
	fmt.Println("Реверсований масив:", reversed)
}
