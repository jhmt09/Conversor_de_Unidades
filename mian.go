package main

import "fmt"

func peso_do_usuario(){
	var peso float64
	var libras float64 
	fmt.Println("Digite aqui seu peso: ")
	fmt.Scanln(&peso)
	libras = 2.20462
	x := (libras * peso)
	fmt.Println("Seu peso em libras aproximado é: ", x)
}

func tranformando_celsius_em_fharenheit_e_kelvin(){
	var temperaturacelsius float64
	var temperaturafharenheit float64
	fmt.Println("Digite a temperatura: ")
	fmt.Scanln(&temperaturacelsius)
	temperaturafharenheit = 1.8
	y := (temperaturacelsius * temperaturafharenheit)
	fmt.Println("A temperatura em Fharenheit é: ", y+32)
	fmt.Println("A temperatura em Kelvin é: ", temperaturacelsius + 273)
	
}








func main(){
	peso_do_usuario()
	tranformando_celsius_em_fharenheit_e_kelvin()
}