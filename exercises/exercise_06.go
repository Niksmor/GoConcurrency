package exercises

import (
	"context"
	"time"
)

// Задача:
// Запустите горутину, которая каждую секунду выводит `"tick"`.
// Используйте `context.WithTimeout`, чтобы горутина завершилась автоматически через 3 секунды.

func worker(ctx context.Context) {
	for {
		// TODO ЛОГИКА
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(4 * time.Second)
}
