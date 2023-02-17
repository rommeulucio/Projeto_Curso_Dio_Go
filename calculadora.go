package main

import "fmt"

func main() {
	var num1, num2 float32
	var op string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scanln(&op)

	resultado := 0.0

	switch op {
	case "+":
		resultado = float64(num1) + float64(num2)
	case "-":
		resultado = float64(num1) - float64(num2)
	case "*":
		resultado = float64(num1) * float64(num2)
	case "/":
		resultado = float64(num1) / float64(num2)
	default:
		fmt.Println("operação invalida.")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, op, num2, resultado)
}
