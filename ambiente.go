package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ambiente struct {
	fase         string
	fasecontador int
	faselimite   int

	dia int

	temperatura
	luminosidade
	ventos
	nuvens
}

func ambientenovo() *ambiente {
	ret := ambiente{}

	ret.temperatura = *temperaturaNovo(&ret)
	ret.luminosidade = *luminosidadeNovo(&ret)
	ret.ventos = *ventosNovo(&ret)
	ret.nuvens = *nuvensNovo(&ret)

	ret.fase = "Dia"
	ret.fasecontador = -1
	ret.faselimite = 100

	ret.dia = 0

	return &ret
}

func (a *ambiente) ciclo() {

	rand.Seed(time.Now().UTC().UnixNano())

	a.fasecontador++

	if a.fase == "Noite" && a.fasecontador == a.faselimite {
		a.dia++
	}

	if a.fasecontador >= a.faselimite {
		a.fasecontador = 0
		if a.fase == "Dia" {
			a.fase = "Noite"
		} else {
			a.fase = "Dia"
		}

	}

	a.temperaturaDia()
	a.temperaturaNoite()

	a.claridade()
	a.ventar()

}

func main() {

	var executando bool = true
	var ciclos int = 0

	var amb = ambientenovo()

	for {

		amb.ciclo()

		if amb.fasecontador == 0 {
			fmt.Println("\nDia : ", amb.dia, " Ciclo : ", ciclos, " - FASE : ", amb.fase)
		}

		if amb.fasecontador == 99 {

			//fmt.Println()

			//fmt.Printf("\n\t Temp  :  %.2f NG", amb.temperaturaMedia())

			//fmt.Printf("\n\t Temp 1 :  %.2f NG", amb.temp1)
			//fmt.Printf("\n\t Temp 2 :  %.2f NG", amb.temp2)

			//fmt.Printf("\n\t Temp 3 :  %.2f NG", amb.temp3)
			//fmt.Printf("\n\t Temp 4 :  %.2f NG", amb.temp4)

			//	fmt.Println()

			//		fmt.Printf("\n\t Temp Min :  %.2f NG", amb.tempdiamin)
			//	fmt.Printf("\n\t Temp Max :  %.2f NG", amb.tempdiamax)

			fmt.Printf("\n\t Temperatura :  %.2f NG", amb.temperaturaCorrente)

			if amb.ventorodando == true {
				fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ] - Rodando ", amb.vento, amb.ventoCorrenteNome(), amb.ventoorigem, amb.ventodestino)
			} else {
				fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ]", amb.vento, amb.ventoCorrenteNome(), amb.ventoorigem, amb.ventodestino)
			}

			fmt.Printf("\n\t Luz :  %.2f - %s", amb.luz, amb.luminosidadeCorrenteNome())

			fmt.Println()

			fmt.Println()
		}

		ciclos++
		if ciclos >= 600 {
			executando = false
		}

		if executando == false {
			break
		}

	}

	//fmt.Printf("\n\t Temp Total Min :  %.2f NG", amb.tempmin)
	//fmt.Printf("\n\t Temp Total Max :  %.2f NG", amb.tempmax)

	fmt.Println()

}
