package main

import (
	"fmt"
	"strings"
)

/**
 * Type Person Value Object:
 * used to map data from file
 */
type Person struct {
	cpf           string
	private       string
	incompleto    string
	dtUltCompra   string
	tckMedio      string
	tckUltCompra  string
	lojaFrequente string
	lojaUltCompra string
}

/**
 * parse a row from file to object person
 */
func lineToPerson(line string) Person {
	var person Person
	runes := []rune(line)

	if len(runes) > 36 {
		person.cpf = strings.TrimSpace(string(runes[0:19]))
		person.private = strings.TrimSpace(string(runes[19:30]))
		person.incompleto = strings.TrimSpace(string(runes[31:42]))
		person.dtUltCompra = strings.TrimSpace(string(runes[43:64]))
		person.tckMedio = strings.TrimSpace(string(runes[65:86]))
		person.tckUltCompra = strings.TrimSpace(string(runes[87:110]))
		person.lojaFrequente = strings.TrimSpace(string(runes[111:130]))
		if len(runes) > 144 {
			person.lojaUltCompra = strings.TrimSpace(string(runes[131:150]))
		} else {
			person.lojaUltCompra = ""
		}
	}
	return person
}

/**
 * print all data of a Person
 */
func toPrintPerson(person Person) {
	fmt.Println("-------- VO ")
	fmt.Println("person.cpf:", person.cpf)
	fmt.Println("person.private:", person.private)
	fmt.Println("person.incompleto:", person.incompleto)
	fmt.Println("person.dtUltCompra:", person.dtUltCompra)
	fmt.Println("person.tckMedio:", person.tckMedio)
	fmt.Println("person.tckUltCompra:", person.tckUltCompra)
	fmt.Println("person.lojaFrequente:", person.lojaFrequente)
	fmt.Println("person.lojaUltCompra:", person.lojaUltCompra)
}
