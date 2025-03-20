package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// chamada direta pela thread principal da aplicacao
	fun("Direct Call")

	// chamada da funcao por meio de uma go routine
	go fun("goroutine - 2")

	// as go routines nao se conhecem, entao executam independente de qualquer outro processo

	go fun("goroutine - 1")

	// funcao anonima rodando dentro de uma go routine
	go func() {
		fun("goroutine - 3")
	}()

	// time sleep para dar tempo de iniciar e rodar a goroutine antes de finalizar o processo principal (main)
	time.Sleep(5 * time.Millisecond)

	fmt.Println("Done...")

}
