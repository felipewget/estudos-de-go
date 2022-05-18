// tambem chamada de stack a pilha
package main

// Fun 1 chama a f2  e a f3... e coloca tudo isso na fila... pra terminar todas as funcoes dali precisam acabar
// A func terminou, vai pra outra fila

// tecnica chamada recursao... metodos recursivos
// pode estourar a memoria e ai da xabu

import (
	"fmt"
	"runtime/debug"
)

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {

	f1()
	fmt.Println("Terminou")

}
