package main

import (
	"fmt"
)

// ConvertePeso transforma o peso em quilogramas para libras.
func ConvertePeso() {
	var pesoEmKg float64

	fmt.Println("Digite aqui o seu peso em quilogramas: ")
	_, err := fmt.Scanln(&pesoEmKg)
	if err != nil {
		fmt.Println("Entrada inválida! Por favor, digite um número válido.")
		return
	}

	const librasPorKg = 2.20462
	pesoEmLb := librasPorKg * pesoEmKg

	fmt.Printf("Seu peso em libras é de: %.2f lbs\n", pesoEmLb)
}

// ConverteTemperatura transforma a temperatura em graus Celsius para graus Fahrenheit e Kelvin.
func ConverteTemperatura() {
	var temperaturaEmCelsius float64

	fmt.Println("Digite a temperatura em graus Celsius: ")
	_, err := fmt.Scanln(&temperaturaEmCelsius)
	if err != nil {
		fmt.Println("Entrada inválida! Por favor, digite um número válido.")
		return
	}

	const fatorDeConversaoFahrenheit = 1.8
	temperaturaEmFahrenheit := fatorDeConversaoFahrenheit*temperaturaEmCelsius + 32
	temperaturaEmKelvin := temperaturaEmCelsius + 273.15

	fmt.Printf("A temperatura em Fahrenheit é de: %.2f °F\n", temperaturaEmFahrenheit)
	fmt.Printf("A temperatura em Kelvin é de: %.2f K\n", temperaturaEmKelvin)
}

func main() {
	ConvertePeso()
	ConverteTemperatura()
}
