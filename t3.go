package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Ініціалізація списку продуктів
	products := []string{"Яблука", "Банани", "Апельсини", "Груші", "Виноград"}
	fmt.Println("Список доступних продуктів:")
	for i, product := range products {
		fmt.Printf("%d. %s\n", i+1, product)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введіть номери продуктів, які хочете забрати (максимум 2), розділені пробілами:")

	input, _ := reader.ReadString('\n') // Зчитує рядок до нового рядка
	strNumbers := strings.Fields(input)

	if len(strNumbers) == 0 || (len(strNumbers) == 1 && (strNumbers[0] == "0" || strNumbers[0] == "")) {
		fmt.Println("Виберіть що потрібно забрати, або йдіть звідси.")
		return
	}

	// Перевірка на максимум 2 продукти
	if len(strNumbers) > 2 {
		fmt.Println("Виберіть не більше ніж 2 продукти.")
		return
	}

	var removedProducts []string
	for _, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil || num < 1 || num > len(products) {
			fmt.Printf("Неправильний номер продукту: %s\n", strNum)
			continue
		}
		removedProducts = append(removedProducts, products[num-1]) // Додаємо обраний продукт до списку видалених
	}

	if len(removedProducts) > 0 {
		fmt.Println("Ви забрали наступні продукти:")
		for _, product := range removedProducts {
			fmt.Println(product)
		}
	} else {
		fmt.Println("Ви нічого не забрали.")
	}
}
