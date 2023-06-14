package tools

import "fmt"

type Modelos int64

const (
	INVALIDO Modelos = iota
	CLT
	HORISTA
	PJ
)

func ObterModelos(o Modelos) {
	fmt.Println(o)
}

func IntegerToModel(o int64) Modelos {
	switch o {
	case 1:
		return CLT
	case 2:
		return HORISTA
	case 3:
		return PJ
	}
	return INVALIDO
}

func ModelToInteger(o Modelos) int64 {
	switch o {
	case CLT:
		return 1
	case HORISTA:
		return 2
	case PJ:
		return 3
	}
	return 0
}

func (o Modelos) String() string {
	switch o {
	case CLT:
		return "CLT"
	case HORISTA:
		return "HORISTA"
	case PJ:
		return "PJ"
	}
	return "invalido"
}

func ShowOptions() string {
	var t string
	for i := 1; i != 0; i++ {
		v := IntegerToModel(int64(i)).String()

		if v == "invalido" {
			break
		}
		t += fmt.Sprintf("%d - %s\n", i, v)
	}
	return t
}
