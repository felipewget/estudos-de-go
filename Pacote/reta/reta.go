package main

import "math"

// Inicializando qualquer letra maiuscula, vai ser publico, quando e minuscula, ela e privada
// nao pode ter funcoes duplicadas e serao compuladas num arquivo so
// vc nao tem visibilidade privada dentro do arquivo, apenas dentro do pacote

// pacotes que tem arquivos que tem codigo la dentro

// Ponto -> publico
// ponto ou _ponto -> privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Todas as funcoes publicas exigem um comentario
// calcula a distancia linear entre 2 pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	distancia := math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
	return float64(distancia)
}
