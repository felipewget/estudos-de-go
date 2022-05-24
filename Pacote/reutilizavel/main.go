package main

import (
	"fmt"
	"github/cod3cursos/area" // Esse kra ta dentro da pasta do gopath
)

// Vamos criar o workspage, no Gopath... e vamos poder usar no nosso projeto algo la do gopath
// na pasta go raiz tem a pasta
// BIN -> Executaveis |
// PKG -> pacotes compilados
// src -> source C:\Program Files\Go\src

// Vamos usar o pacote da area main que coloquei no GO path

func main() {

	fmt.Println(area.Circulo(6.0))
	fmt.Println(area.Retangulo(6.0, 6.0))
	// fmt.Println(area._privada()) // <- daria erro porque essa funcao nao e publica
}
