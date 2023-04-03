package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia,
	NumeroConta,
	Operacao int
	saldo float64
}

func (Conta *ContaPoupanca) Sacar(pValorDoSaque float64) string {
	podeSacar := pValorDoSaque > 0 && pValorDoSaque <= Conta.saldo
	if podeSacar {
		Conta.saldo -= pValorDoSaque
		return fmt.Sprintf("Saque de %f realizado com sucesso", pValorDoSaque)
	} else {
		return "saldo insuficiente"
	}
}

func (Conta *ContaPoupanca) Depositar(pValorDeposito float64) string {
	var deposito = pValorDeposito > 0
	if deposito {
		Conta.saldo += pValorDeposito
		return fmt.Sprintf("Depósito de %f realizado com sucesso", pValorDeposito)
	} else {
		return fmt.Sprintf("Valor %f incorreto", pValorDeposito)
	}
}

func (Conta *ContaPoupanca) GetSaldo() float64 {
	return Conta.saldo
}
