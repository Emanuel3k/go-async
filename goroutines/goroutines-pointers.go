package main

import (
	"fmt"
	"time"
)

func funPointer(value *string) {
	for {
		fmt.Println(*value)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// variavel test recebendo um valor
	var test string = "test"

	//passando o endereco de memoria (ponteiro) de test para outra variavel
	var pointerTest *string = &test

	// chamada da funcao que printa o valor dentro do endereco de memoria
	go funPointer(pointerTest)

	time.Sleep(time.Second)

	// atualizacao do valor da variavel que esta sendo printado
	// e de fato isso funciona, pois estou alterando o valor que esta no endereco de memoria, e enquanto a goroutine roda consigo manipular o endereco que esta la dentro
	test = "test 2"

	time.Sleep(time.Second)
}
