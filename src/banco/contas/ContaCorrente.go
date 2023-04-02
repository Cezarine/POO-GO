package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (Conta *ContaCorrente) Sacar(pValorDoSaque float64) string {
	podeSacar := pValorDoSaque > 0 && pValorDoSaque <= Conta.Saldo
	if podeSacar {
		Conta.Saldo -= pValorDoSaque
		return fmt.Sprintf("Saque de %f realizado com sucesso", pValorDoSaque)
	} else {
		return "Saldo insuficiente"
	}
}

func (Conta *ContaCorrente) Depositar(pValorDeposito float64) string {
	var deposito = pValorDeposito > 0
	if deposito {
		Conta.Saldo += pValorDeposito
		return fmt.Sprintf("Depósito de %f realizado com sucesso", pValorDeposito)
	} else {
		return fmt.Sprintf("Valor %f incorreto", pValorDeposito)
	}
}

func (ContaDebito *ContaCorrente) Transferir(pContaCredito *ContaCorrente, pValor float64) bool {
	transferir := ContaDebito.Saldo >= pValor && pValor > 0
	if transferir {
		ContaDebito.Saldo -= pValor
		pContaCredito.Depositar(pValor)
		return true
	} else {
		return false
	}
}
