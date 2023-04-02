package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoGuilherme := contas.ContaCorrente{
		Titular:       "Guilherme",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         500.25}

	contaDoGustavo := contas.ContaCorrente{}
	contaDoGustavo.Titular = "Gustavo"
	contaDoGustavo.Saldo = 20

	fmt.Println(contaDoGustavo)
	fmt.Println("")
	fmt.Println(contaDoGuilherme)
	fmt.Println("")
	valorDoSaque := 200.0
	valorDoDeposito := 1000.
	valorTransferencia := 250.

	fmt.Println(contaDoGuilherme.Sacar(valorDoSaque))
	fmt.Println("")
	fmt.Println(contaDoGuilherme.Depositar(valorDoDeposito))
	fmt.Println(contaDoGuilherme.Saldo)
	fmt.Println("")
	status := contaDoGuilherme.Transferir(&contaDoGustavo, valorTransferencia)
	if status {
		fmt.Println("Transferencia de ", valorTransferencia, "realizada com sucesso")
		fmt.Println("")
		fmt.Println("Conta do Guilherme:", contaDoGuilherme.Saldo)
		fmt.Println("")
		fmt.Println("Conta do Gustavo:", contaDoGustavo.Saldo)
	} else {
		fmt.Println("Transferencia de ", valorTransferencia, "não foi realizado")
	}

	/*	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

		fmt.Println(contaDaBruna)

		var contaDaCris *ContaCorrente
		contaDaCris = new(ContaCorrente)
		contaDaCris.Titular = "Cris"

		fmt.Println(*contaDaCris)*/
}
