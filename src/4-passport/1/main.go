package main

import (
	"fmt"
	"strings"

	"../../utils"
)

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	passportRows := utils.GetLines(file)
	passports := getPassports(passportRows)

	validPassportcount := 0
	for _, passport := range passports {
		if passport.validate() {
			validPassportcount++
		}
	}

	fmt.Println(validPassportcount)
}

func getPassports(rows []string) []passport {
	passports := []passport{}

	workingPassport := passport{}
	for _, row := range rows {
		if row == "" {
			passports = append(passports, workingPassport)
			workingPassport = passport{}
			continue
		}
		attributes := strings.Split(row, " ")
		for _, attribute := range attributes {
			workingPassport.assignAttribute(attribute)
		}
	}

	passports = append(passports, workingPassport)

	return passports
}

func (p *passport) assignAttribute(attribute string) {
	splitAttribute := strings.Split(attribute, ":")
	key := splitAttribute[0]
	value := splitAttribute[1]

	switch key {
	case "byr":
		p.birthYear = value
	case "iyr":
		p.issueYear = value
	case "eyr":
		p.expirationYear = value
	case "hgt":
		p.height = value
	case "hcl":
		p.hairColor = value
	case "ecl":
		p.eyeColor = value
	case "pid":
		p.passportID = value
	case "cid":
		p.countryID = value
	}
}

func (p passport) validate() bool {
	return p.birthYear != "" &&
		p.issueYear != "" &&
		p.expirationYear != "" &&
		p.height != "" &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportID != ""
}
