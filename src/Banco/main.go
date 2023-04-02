package main

import (
	"fmt"
	"C:\Desenvolvimento\Cursos\POO-GO\src\Banco\Conta"
)

func main() {
	contaDoGuilherme := contas.ContaCorrente{
		titular:       "Guilherme",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         500.25}

	contaDoGustavo := contas.ContaCorrente{}
	contaDoGustavo.titular = "Gustavo"
	contaDoGustavo.saldo = 20

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
	fmt.Println(contaDoGuilherme.saldo)
	fmt.Println("")
	status := contaDoGuilherme.Transferir(&contaDoGustavo, valorTransferencia)
	if status {
		fmt.Println("Transferencia de ", valorTransferencia, "realizada com sucesso")
		fmt.Println("")
		fmt.Println("Conta do Guilherme:", contaDoGuilherme.saldo)
		fmt.Println("")
		fmt.Println("Conta do Gustavo:", contaDoGustavo.saldo)
	} else {
		fmt.Println("Transferencia de ", valorTransferencia, "não foi realizado")
	}

	/*	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

		fmt.Println(contaDaBruna)

		var contaDaCris *ContaCorrente
		contaDaCris = new(ContaCorrente)
		contaDaCris.titular = "Cris"

		fmt.Println(*contaDaCris)*/
}
