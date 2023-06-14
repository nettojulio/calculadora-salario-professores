package sql

import (
	"calculadora-salario-professores/internal/connection"
	"calculadora-salario-professores/internal/models"
	"calculadora-salario-professores/internal/tools"
)

func FindAllProfessores() []models.ImplProfessor {
	db := connection.GetDBInstance()

	defer db.Close()

	querySelector := "SELECT professor.id, professor.nome, contratos.tipo, contrato_professor.id_detalhes_contrato, contrato_professor.inativo FROM professor LEFT JOIN contratos ON contratos.id = professor.id_detalhes_contrato LEFT JOIN contrato_professor ON contrato_professor.id_professor = professor.id"

	selectAllProducts, err := db.Query(querySelector)

	if err != nil {
		panic(err.Error())
	}

	professorList := []models.ImplProfessor{}

	for selectAllProducts.Next() {
		var id, idDetalhes int64
		var nome, tipoContrato string
		var inativo bool

		err = selectAllProducts.Scan(&id, &nome, &tipoContrato, &idDetalhes, &inativo)

		if err != nil {
			panic(err.Error())
		}
		professor := models.NewProfessorDTO(id, nome, tipoContrato, idDetalhes, inativo)

		specificProfessor := newSpecificProfessor(professor)

		professorList = append(professorList, specificProfessor)
	}

	return professorList
}

func newSpecificProfessor(professor models.ProfessorDTO) models.ImplProfessor {
	switch professor.TipoContrato {
	case "CLT":
		salario := findCltData(professor.IdDetalhes)
		return models.NewProfessorCLT(professor.Nome, salario)
	case "HORISTA":
		valorHora, horasTrabalhadas := findHoristaData(professor.IdDetalhes)
		return models.NewProfessorHorista(professor.Nome, valorHora, horasTrabalhadas)
	case "PJ":
		contrato := findPjData(professor.IdDetalhes)
		return models.NewProfessorPJ(professor.Nome, contrato)
	}
	return nil
}

func findCltData(id int64) float64 {
	db := connection.GetDBInstance()

	defer db.Close()

	querySelector := "SELECT professor_clt.salario FROM professor_clt where id=$1"

	selectAllProducts, err := db.Query(querySelector, id)
	if err != nil {
		panic(err.Error())
	}

	var salario float64

	for selectAllProducts.Next() {
		err = selectAllProducts.Scan(&salario)
		if err != nil {
			panic(err.Error())
		}
	}

	return salario
}

func findHoristaData(id int64) (float64, int64) {
	db := connection.GetDBInstance()

	defer db.Close()

	querySelector := "SELECT professor_horista.valor_hora, professor_horista.horas_trabalhadas FROM professor_horista where id=$1"

	selectAllProducts, err := db.Query(querySelector, id)
	if err != nil {
		panic(err.Error())
	}

	var valorHora float64
	var horasTrabalhadas int64

	for selectAllProducts.Next() {
		err = selectAllProducts.Scan(&valorHora, &horasTrabalhadas)
		if err != nil {
			panic(err.Error())
		}
	}

	return valorHora, horasTrabalhadas
}

func findPjData(id int64) float64 {
	db := connection.GetDBInstance()

	defer db.Close()

	querySelector := "SELECT professor_pj.contrato FROM professor_pj where id=$1"

	selectAllProducts, err := db.Query(querySelector, id)
	if err != nil {
		panic(err.Error())
	}

	var contrato float64

	for selectAllProducts.Next() {
		err = selectAllProducts.Scan(&contrato)
		if err != nil {
			panic(err.Error())
		}
	}

	return contrato
}

func SaveNewProfessor(professor models.ImplProfessor) {
	db := connection.GetDBInstance()
	defer db.Close()

	professorId := 0
	err := db.QueryRow("INSERT INTO professor(nome, id_detalhes_contrato) values($1,$2) RETURNING id", professor.GetNome(), professor.GetTipoContrato()).Scan(&professorId)
	if err != nil {
		panic(err.Error())
	}

	professorSpecificId := 0
	switch professor.GetTipoContrato().String() {
	case "CLT":
		clt := professor.(*models.ProfessorCLT)
		querySelector := "INSERT INTO professor_clt (salario) values($1) RETURNING id"
		err := db.QueryRow(querySelector, clt.Salario).Scan(&professorSpecificId)
		if err != nil {
			panic(err.Error())
		}
	case "HORISTA":
		horista := professor.(*models.ProfessorHorista)
		querySelector := "INSERT INTO professor_horista (valor_hora, horas_trabalhadas) values($1,$2) RETURNING id"
		err := db.QueryRow(querySelector, horista.ValorHora, horista.HorasTrabalhadas).Scan(&professorSpecificId)
		if err != nil {
			panic(err.Error())
		}
	case "PJ":
		pj := professor.(*models.ProfessorPJ)
		querySelector := "INSERT INTO professor_pj (contrato) values($1) RETURNING id"
		err := db.QueryRow(querySelector, pj.Contrato).Scan(&professorSpecificId)
		if err != nil {
			panic(err.Error())
		}
	}

	insertDataDB, err := db.Prepare("insert into contrato_professor(id_contrato, id_professor, id_detalhes_contrato, inativo) values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insertDataDB.Exec(tools.ModelToInteger(professor.GetTipoContrato()), professorId, professorSpecificId, false)
}
