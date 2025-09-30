package exercises

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 2. Параллельный счётчик: Усложнение (atomic)
// Задача:
// В прошлой задаче мы использовали `sync.Mutex` для защиты общей переменной от гонки данных.
// Теперь реши ту же задачу с помощью **пакета `sync/atomic`, без использования мьютекса.
func main() {
	// int64 для атомарного счётчика
	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {

			defer wg.Done()

			// Атомарное увеливение
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println("Result:", counter)
}
