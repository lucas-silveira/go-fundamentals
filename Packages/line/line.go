/*
	Em Go, entende-se por pacote o conjunto de arquivos
	que pertencem a um mesmo diretório. Isso se deve ao fato de que
	quando o programa é compilado, todos os arquivos que pertencem a um
	diretório são unidos num único pacote.

	As funções devem possuir nomes únicos dentro do pacote.

	Se iniciarmos uma struct ou variável com letra maiúscula,
	ela será pública (visível dentro e fora do pacote). Caso iniciemos
	com letra minúscula, ela será privada (visível somente dentro do pacote).
*/

package main

import "math"

// Point representa uma coordenada no plano cartesiano
type Point struct {
	x float64
	y float64
}

func catetos(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distance é responsável por calcular a distância linear entre dois pontos
func Distance(a, b Point) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
