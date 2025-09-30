package exercises

import (
	"fmt"
	"math/rand"
)

// 3. Сумма чисел с распределением задач
// Задача:
// Разделите массив из 1000 случайных чисел на 10 равных частей.
// Запустите 10 горутин, каждая должна посчитать сумму своей части и отправить результат в канал.
// Главная функция должна собрать все результаты и вывести общую сумму.

func sumPart(nums []int, resultCh chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	resultCh <- sum
}

func main() {
	const parts = 10
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}

	resultCh := make(chan int, parts)

	chunkSize := len(nums) / parts

	for i := 0; i < parts; i++ {
		start := i * chunkSize
		end := start + chunkSize
		go sumPart(nums[start:end], resultCh)
	}

	totalSum := 0
	for i := 0; i < parts; i++ {
		totalSum += <-resultCh
	}

	fmt.Println("Total sum:", totalSum)
}
