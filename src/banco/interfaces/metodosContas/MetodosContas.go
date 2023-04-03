package metodosContas

type verificaConta interface {
	Sacar(pValor float64) string
	Depositar(pValor float64) string
}

func PagarBoleto(pConta verificaConta, pValorBoleto float64) string {
	return pConta.Sacar(pValorBoleto)
}

func Depositar(pConta verificaConta, pValorDeposito float64) string {
	return pConta.Depositar(pValorDeposito)
}
