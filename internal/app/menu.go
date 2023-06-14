package app

import (
	"calculadora-salario-professores/internal/models"
	"calculadora-salario-professores/internal/sql"
	"calculadora-salario-professores/internal/tools"
	"fmt"
)

func Run() {
	var cliOptions int64
	for cliOptions != 3 {
		cliOptions = tools.GetIntegerValues("Opções\n1 - Cadastrar Professores\n2 - Listar Professores\n3 - SAIR\nEscolha uma das opções: ")
		switch cliOptions {
		case 1:
			professor := professorFactory()
			sql.SaveNewProfessor(professor)
		case 2:
			professores := sql.FindAllProfessores()
			for _, v := range professores {
				fmt.Println(v.Display())
			}
		}
	}
}

func professorFactory() models.ImplProfessor {

	var (
		nome           string
		modeloContrato int64
	)

	for len(nome) == 0 {
		nome = tools.GetStringValues("Informe o nome do professor: ")
	}

	for tools.IntegerToModel(modeloContrato).String() == "invalido" {
		modeloContrato = tools.GetIntegerValues("Modelos de trabalho\n" + tools.ShowOptions() + "Informe o modelo de trabalho desejado: ")
	}

	switch modeloContrato {
	case 1:
		return promptCLT(nome)
	case 2:
		return promptHorista(nome)
	case 3:
		return promptPJ(nome)
	default:
		panic("opcao invalida")
	}
}

func promptCLT(nome string) *models.ProfessorCLT {
	var salario float64
	for salario <= 0.0 {
		salario = tools.GetFloatValues("Informe o salário: ")
	}
	return models.NewProfessorCLT(nome, salario)
}

func promptHorista(nome string) *models.ProfessorHorista {
	var (
		valorHora        float64
		horasTrabalhadas int64
	)

	for valorHora <= 0.0 {
		valorHora = tools.GetFloatValues("Informe o valor da hora: ")
	}

	for horasTrabalhadas <= 0 {
		horasTrabalhadas = tools.GetIntegerValues("Informe as horas trabalhadas: ")
	}

	return models.NewProfessorHorista(nome, valorHora, horasTrabalhadas)
}

func promptPJ(nome string) *models.ProfessorPJ {
	var contrato float64

	for contrato <= 0 {
		contrato = tools.GetFloatValues("Informe o valor do contrato: ")
	}

	return models.NewProfessorPJ(nome, contrato)
}
