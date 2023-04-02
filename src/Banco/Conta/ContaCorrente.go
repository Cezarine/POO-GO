package contas

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (Conta *ContaCorrente) Sacar(pValorDoSaque float64) string {
	podeSacar := pValorDoSaque > 0 && pValorDoSaque <= Conta.saldo
	if podeSacar {
		Conta.saldo -= pValorDoSaque
		return fmt.Sprintf("Saque de %f realizado com sucesso", pValorDoSaque)
	} else {
		return fmt.Sprintf("Saldo insuficiente")
	}
}

func (Conta *ContaCorrente) Depositar(pValorDeposito float64) string {
	var deposito = pValorDeposito > 0
	if deposito {
		Conta.saldo += pValorDeposito
		return fmt.Sprintf("Depósito de %f realizado com sucesso", pValorDeposito)
	} else {
		return fmt.Sprintf("Valor %f incorreto", pValorDeposito)
	}
}

func (ContaDebito *ContaCorrente) Transferir(pContaCredito *ContaCorrente, pValor float64) bool {
	transferir := ContaDebito.saldo >= pValor && pValor > 0
	if transferir {
		ContaDebito.saldo -= pValor
		pContaCredito.Depositar(pValor)
		return true
	} else {
		return false
	}
}
