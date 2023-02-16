package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

// Este metodo área possui um tipo retangulo
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

// perimetro a soma dos dois lados
func (r *retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 10, altura: 5}

	//Aqui chamamos os 2 metodos definidos para a nossa estrutura
	fmt.Println("Área:", r.area())
	fmt.Println("Perimetro:", r.perimetro())

}

//	for i := 1; i <= 10; i++ {
//		if i%2 == 0 {
//			fmt.Println("par")
//		} else {
//			fmt.Println("impar")
//		}
//	}
