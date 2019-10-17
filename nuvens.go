package main

type nuvens struct {
	ambienteC *ambiente

	nuvem         float32
	nuvemcontador int
	nuvemlimte    int
}

func nuvensNovo(a *ambiente) *nuvens {
	ret := nuvens{}
	ret.ambienteC = a

	ret.nuvem = 0
	ret.nuvemcontador = 0
	ret.nuvemlimte = 20

	return &ret
}
