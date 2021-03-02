package main

import (
	"fmt"
)

/**
 * Save a array of object's in database
 */
func savePeople(people []PersonDTO) {
	id := 0

	db := postgresConnection()

	for i, p := range people {
		if !p.hasError {
			sqlStatement := `INSERT INTO people (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, cnpj_loja_mais_frequente, cnpj_loja_ultima_compra) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
			_, err := db.Exec(sqlStatement, p.cpf, p.private, p.incompleto, p.dtUltCompra, p.tckMedio, p.tckUltCompra, p.lojaFrequente, p.lojaUltCompra)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			sqlStatement := `INSERT INTO error (chave, message) VALUES ($1, $2)`
			_, err := db.Exec(sqlStatement, toString(p), p.msgError)
			if err != nil {
				fmt.Println(err)
			}
		}

		id = i
	}
	fmt.Println("Total record ", id)
}
