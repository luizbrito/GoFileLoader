package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
 * type Person Data Trasnfer Object:
 * Used to map data in database.
 */

type PersonDTO struct {
	cpf           string
	private       bool
	incompleto    bool
	dtUltCompra   time.Time
	tckMedio      float64
	tckUltCompra  float64
	lojaFrequente string
	lojaUltCompra string
	hasError      bool
	msgError      string
}

/**
 * parse from VO to DTO.
 */
func parseToPersonDTO(person Person) PersonDTO {
	var p PersonDTO

	p.cpf = person.cpf
	p.lojaFrequente = strings.ReplaceAll(person.lojaFrequente, "NULL", "")
	p.lojaUltCompra = strings.ReplaceAll(person.lojaUltCompra, "NULL", "")
	p.lojaUltCompra = strings.ReplaceAll(p.lojaUltCompra, "0x00", "")

	if len(p.lojaUltCompra) > 18 {
		runes := []rune(p.lojaUltCompra)
		p.lojaUltCompra = string(runes[0:18])
		//fmt.Println(p.lojaUltCompra)
	}

	if person.dtUltCompra != "" {
		p.dtUltCompra = parseStringToDate(person.dtUltCompra)
	}
	// parse to bool value
	pri, error := strconv.ParseBool(person.private)
	if error == nil {
		p.private = pri
	}
	// parse to bool value
	inc, error := strconv.ParseBool(person.incompleto)
	if error == nil {
		p.incompleto = inc
	}

	tm, error := strconv.ParseFloat(strings.ReplaceAll(person.tckMedio, ",", "."), 64)
	if error == nil {
		p.tckMedio = tm
	}

	tl, error := strconv.ParseFloat(strings.ReplaceAll(person.tckUltCompra, ",", "."), 64)
	if error == nil {
		p.tckUltCompra = tl
	}

	return isValidCpfCnpj(p)
}

/**
 * print all data of a PersonDTO
 */
func toPrintPersonDTO(person PersonDTO) {
	fmt.Println("-------- DTO ")
	fmt.Println("person.cpf:", person.cpf)
	fmt.Println("person.private:", person.private)
	fmt.Println("person.incompleto:", person.incompleto)
	fmt.Println("person.dtUltCompra:", person.dtUltCompra)
	fmt.Println("person.tckMedio:", person.tckMedio)
	fmt.Println("person.tckUltCompra:", person.tckUltCompra)
	fmt.Println("person.lojaFrequente:", person.lojaFrequente)
	fmt.Println("person.lojaUltCompra:", person.lojaUltCompra)
}

/**
 * parse object to string
 */
func toString(p PersonDTO) string {
	return p.cpf //+ " | " + p.lojaFrequente + " | " + p.lojaUltCompra
}

/**
 * function to validate some data,
 * with this information will put some record in a error table
 */
func isValidCpfCnpj(person PersonDTO) PersonDTO {

	//	valida se é um cpf valido
	if (person.cpf != "") && (!IsCPF(person.cpf)) {
		person.hasError = true
		person.msgError = "CPF invalido "
		//fmt.Println(person.msgError)
	}
	//	valida se é um cnpj valido
	if (person.lojaFrequente != "") && (!IsCNPJ(person.lojaFrequente)) {
		person.hasError = true
		person.msgError = "CNPJ invalido"
		//fmt.Println(person.msgError)
	}
	//	valida se é um cnpj valido
	if (person.lojaUltCompra != "") && (!IsCNPJ(person.lojaUltCompra)) {
		person.hasError = true
		person.msgError = "CNPJ invalido "
		//fmt.Println(person.msgError)
	}
	return person
}

/**
 * Parse String like 2006-01-02, to a valid time.Time.
 * param string
 * return time.Time
 */
func parseStringToDate(s string) time.Time {

	const (
		layoutISO = "2006-01-02"
	)
	t, _ := time.Parse(layoutISO, s)
	return t
}
