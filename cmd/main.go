package main

import (
	"calculadora-salario-professores/internal/app"
	"calculadora-salario-professores/internal/connection"
)

func init() {
	db := connection.GetDBInstance()

	defer db.Close()

	err := db.Ping()

	if err != nil {
		panic("database connection refused")
	}
}

func main() {
	app.Run()
}
