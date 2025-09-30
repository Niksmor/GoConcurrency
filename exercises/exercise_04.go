package exercises

import (
	"fmt"
	"sync"
)

// 4. Передача данных по каналу
// Задача:
// Представьте, что у вас есть передатчик (генератор чисел) и приёмник (читатель).
// Передатчик генерирует числа от 1 до 5 и отправляет их по каналу.
// Приёмник читает эти числа и выводит их в консоль.
// После того как все числа отправлены, необходимо корректно закрыть канал.
// Реализуйте программу с двумя горутинами:
//   - одна пишет числа в канал
//   - вторая читает их и печатает

func generator(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // закрываю какнал после отправки

	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch { // читаб из канала пока он не закрыт
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go generator(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
