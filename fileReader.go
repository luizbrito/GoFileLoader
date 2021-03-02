package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * reader the file and return an array of Data Transfer Object
 */
func fileReader(namefile string) []PersonDTO {
	fmt.Println("fileReader", namefile)
	csvFile, _ := os.Open(namefile)
	scanner := bufio.NewScanner(csvFile)
	var people []PersonDTO
	count := 0
	// read line by line
	for scanner.Scan() {
		line := scanner.Text()
		if count > 0 {
			// parse the array to object and do the fist clean
			person := lineToPerson(line)
			// parse the DTO and valid some data
			personDTO := parseToPersonDTO(person)
			// add the DTO in vector to return
			people = append(people, personDTO)
		}
		count++
	}
	//fmt.Println("lines", count-1)
	return people
}
