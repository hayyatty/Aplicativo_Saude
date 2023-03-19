package main

import (
	. "fmt"
	"os"
	"time"
)

func decidir() {

	//Selecionar oq o usuário vai querer usar
	var escolha int
	time.Sleep(5 * time.Second)
	Println("Olá eu sou o app da sua saúde")
	Println("Selecione que sessão do app você quer usar ")
	Println("[1]- Cálculo de IMC (Indice de massa corporal) ")
	Println("[2]- Cálculo da Taxa de metabolismo basal")
	Println("[3]-  ")
	Println("[0]- Sair do programa")
	Scan(&escolha)
	es(escolha)
	if escolha == 0 {
		os.Exit(0)
	}
}
func es(escolha int) {
	switch escolha {
	case 1:
		Println("Você selecionou a Calculadora de IMC")
		Println("Iniciando...")
		for i := 0; i < 5; i++ {
			Print("|")
			time.Sleep(250 * time.Millisecond)
			Print("\b")
			Print(" ")
			Print("\b")
			time.Sleep(250 * time.Millisecond)
		}
		Println()
		peso, altura, _, _, _ := ColetarDados()
		imc := peso / (altura * altura)
		nivelIMC(imc)
	case 2:
		Println("Você selecionou a calculadora de taxa de metabolismo basal")
		for i := 0; i < 5; i++ {
			Print("|")
			time.Sleep(250 * time.Millisecond)
			Print("\b")
			Print(" ")
			Print("\b")
			time.Sleep(250 * time.Millisecond)
		}
		Println()
		peso, altura, idade, nome, sexo := ColetarDados()
		var txexer int
		var tmb float32
		Println(nome, "Selecione a sua taxa de exercicio")
		Println("[1]Sedentário")
		Println("[2]Levemente Ativo")
		Println("[3]Moderadamente Ativo")
		Println("[4]Altamente Ativo")
		Println("[5]Extremamente Ativo")
		Scan(&txexer)
		tx := map[int]float32{1: 1.2,
			2: 1.375,
			3: 1.55,
			4: 1.725,
			5: 1.9,
		}
		if sexo == 1 {
			tmb = tx[int(txexer)] * (66.47 + ((13.75 * peso) + (5.003 * altura) - (float32(idade) * 6.755)))
		} else if sexo == 2 {
			tmb = tx[int(txexer)] * (655.1 + ((9.563 * peso) + (1.85 * altura) - (float32(idade) * 4.676)))
		}
		Println("A sua Taxa de metabolismo basal é de :", tmb)
	case 3:
		Println()
		for i := 0; i < 5; i++ {
			Print("|")
			time.Sleep(250 * time.Millisecond)
			Print("\b")
			Print(" ")
			Print("\b")
			time.Sleep(250 * time.Millisecond)
		}
		Println()

	}
}

// ColetarDados Coletar os dados do usuario para faer os calcúlos
func ColetarDados() (peso float32, altura float32, idade int8, nome string, sexo int) {
	Println("Precisaremos de alguns dados de sua saúde para prosseguir")
	Println("Qual o seu nome ?")
	Scan(&nome)
	Println("Qual a sua altura ? (em METROS)")
	Scan(&altura)
	Println("Qual a sua idade ?")
	Scan(&idade)
	Println("Qual o seu peso ? (em KG)")
	Scan(&peso)
	Println("Qual o sexo")
	Println("[1]masculino")
	Println("[2]feminino")
	Scan(&sexo)
	return
}
func nivelIMC(imc float32) {
	var class string
	println("O seu nivel de IMC indica que :")
	if imc <= 18.5 {
		class = "Abaixo do peso normal"
		println(class)
	} else if imc >= 18.5 && imc <= 24.9 {
		class = "Peso normal"
		println(class)
	} else if imc >= 25.0 && imc <= 29.9 {
		class = "Excesso de peso"
		println(class)
	} else if imc >= 30.0 && imc <= 34.9 {
		class = "Obesidade GRAU 1"
		println(class)
	} else if imc >= 35.0 && imc <= 39.9 {
		class = "Obesidade GRAU 2 (Severa)"
		println(class)
	} else {
		class = "Obesidade GRAU 3 (Morbida)"
	}
}
