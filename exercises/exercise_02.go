package exercises

import (
	"fmt"
)

// 2. Параллельный счётчик: Усложнение (atomic)
// Задача:
// В прошлой задаче мы использовали `sync.Mutex` для защиты общей переменной от гонки данных.
// Теперь реши ту же задачу с помощью **пакета `sync/atomic`, без использования мьютекса.
func main() {
	counter := 0

	for i := 0; i < 10; i++ {
		// TODO использовать atomic
		// TODO: Запустить горутину, которая увеличит counter
	}

	fmt.Println("Result:", counter)
}
