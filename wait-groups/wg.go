package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	//se adicionar mais goroutines a do que o valor inicial do WaitGroup, o programa irá falhar

	//se colocar menos goroutine no WaitGroup do que o valor real, finalizará antes de todas as goroutines terminarem
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			demoraParaExecutar(2 * time.Second)
		}()
	}

	wg.Wait()

	//go demoraParaExecutar(2 * time.Second)
	//go demoraParaExecutar(2 * time.Second)
	//go demoraParaExecutar(2 * time.Second)
	//go demoraParaExecutar(2 * time.Second)
	//go demoraParaExecutar(2 * time.Second)
}

func demoraParaExecutar(t time.Duration) {
	fmt.Println("Começando a executar")

	time.Sleep(t)

	fmt.Println("Finalizando a execução")
}
