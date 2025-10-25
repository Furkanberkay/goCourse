package goroutins

import (
	"fmt"
	"sync"
)

func Worker(id int, wq *sync.WaitGroup) {
	defer wq.Done()
	fmt.Println("worker : ", id, "çalıştı")
}
