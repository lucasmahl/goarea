package goarea

import "math"

//Pi é uma proporção numérica definida pela relação entre
//o perímetro de uma circunferencia e seu diametro
const Pi = 3.1416

//Circ é responsavel por calcular a area da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é area do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visível
func _TrianguloEquilatero(base, altura float64) float64 {
	return (base * altura) / 2
}
