package main

import (
	"banco/clientes"
	"banco/contas"
	"banco/interfaces/metodosContas"
	"fmt"
)

func main() {
	clienteGuilherme := clientes.Titular{
		Nome:      "Guilherme",
		CPF:       "466.746.648-04",
		RG:        "57.759.867-04",
		Profissao: "Desenvolvedor"}

	contaDoGuilherme := contas.ContaCorrente{
		Titular:       clienteGuilherme,
		NumeroAgencia: 589,
		NumeroConta:   123456}
	contaDoGuilherme.Depositar(300)

	teste := metodosContas.PagarBoleto(&contaDoGuilherme, 100)
	fmt.Println(teste)

	/*	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

		fmt.Println(contaDaBruna)

		var contaDaCris *ContaCorrente
		contaDaCris = new(ContaCorrente)
		contaDaCris.Titular = "Cris"

		fmt.Println(*contaDaCris)*/
}
