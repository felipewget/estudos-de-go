package main

import "fmt"

// Usar o go run *.go
// TODOS os init() vao executar antes do main
func init() {
	fmt.Println("INICIALIZANDO EM OUTRO ARQUIVO... func init")
}
