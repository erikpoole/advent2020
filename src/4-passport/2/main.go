package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../utils"
)

type passport struct {
	birthYear      int
	issueYear      int
	expirationYear int
	height         string
	hairColor      string
	eyeColor       string
	rawPassportID  string
	passportID     int
	countryID      int
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
		intValue, err := strconv.Atoi(value)
		if err == nil {
			p.birthYear = intValue
		}
	case "iyr":
		intValue, err := strconv.Atoi(value)
		if err == nil {
			p.issueYear = intValue
		}
	case "eyr":
		intValue, err := strconv.Atoi(value)
		if err == nil {
			p.expirationYear = intValue
		}
	case "hgt":
		p.height = value
	case "hcl":
		p.hairColor = value
	case "ecl":
		p.eyeColor = value
	case "pid":
		p.rawPassportID = value
		intValue, err := strconv.Atoi(value)
		if err == nil {
			p.passportID = intValue
		}
	case "cid":
		intValue, err := strconv.Atoi(value)
		if err == nil {
			p.countryID = intValue
		}
	}
}

func (p passport) validate() bool {
	if p.birthYear < 1920 || p.birthYear > 2002 {
		return false
	}

	if p.issueYear < 2010 || p.issueYear > 2020 {
		return false
	}

	if p.expirationYear < 2020 || p.expirationYear > 2030 {
		return false
	}

	if !validateHeight(p.height) {
		return false
	}

	if !validateHairColor(p.hairColor) {
		return false
	}

	if !validateEyeColor(p.eyeColor) {
		return false
	}

	if len(p.rawPassportID) != 9 {
		return false
	}

	return true
}

func validateHeight(height string) bool {
	if len(height) < 2 {
		return false
	}

	value := height[:len(height)-2]
	measurementType := height[len(height)-2:]

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	switch measurementType {
	case "cm":
		if intValue < 150 || intValue > 193 {
			return false
		}
	case "in":
		if intValue < 59 || intValue > 76 {
			return false
		}
	}

	return true
}

func validateHairColor(hairColor string) bool {
	if len(hairColor) != 7 {
		return false
	}

	if hairColor[0] != '#' {
		return false
	}

	for i := 1; i < 7; i++ {
		if isNum(hairColor[i]) || isHexChar(hairColor[i]) {
			continue
		}
		return false
	}
	return true
}

func isNum(b byte) bool {
	return b > 47 && b < 58
}

func isHexChar(b byte) bool {
	return b > 96 && b < 103
}

func validateEyeColor(eyeColor string) bool {
	return eyeColor == "amb" ||
		eyeColor == "blu" ||
		eyeColor == "brn" ||
		eyeColor == "gry" ||
		eyeColor == "grn" ||
		eyeColor == "hzl" ||
		eyeColor == "oth"
}
