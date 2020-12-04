package day4

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	birthYear      = "byr"
	issueYear      = "iyr"
	expirationYear = "eyr"
	height         = "hgt"
	hairColor      = "hcl"
	eyeColor       = "ecl"
	passportID     = "pid"
	countryID      = "cid"
)

type passportData struct {
	birthYear      int
	issueYear      int
	expirationYear int
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      int
}

func (p *passportData) isValid() bool {
	return p.birthYear != 0 &&
		p.expirationYear != 0 &&
		p.eyeColor != "" &&
		p.hairColor != "" &&
		p.height != "" &&
		p.issueYear != 0 &&
		p.passportID != ""
}

func (p *passportData) isValidFields() bool {
	return p.isValid() &&
		p.isBirthYearValid() &&
		p.isIssueYearValid() &&
		p.isExpirationYearValid() &&
		p.isHeightValid() &&
		p.isHairColorValid() &&
		p.isEyeColorValid() &&
		p.isPassportIDValid()
}

// isBirthYearValid returns true if birthYear is four digits; at least 1920 and at most 2002.
func (p *passportData) isBirthYearValid() bool {
	return p.birthYear >= 1920 && p.birthYear <= 2002
}

// isIssueYearValid returns true if issueYear is four digits; at least 2010 and at most 2020.
func (p *passportData) isIssueYearValid() bool {
	return p.issueYear >= 2010 && p.issueYear <= 2020
}

// isExpirationYearValid returns true if expirationYear is four digits; at least 2020 and at most 2030.
func (p *passportData) isExpirationYearValid() bool {
	return p.expirationYear >= 2020 && p.expirationYear <= 2030
}

// isHeightValid returns true if height is a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func (p *passportData) isHeightValid() bool {
	if in := "in"; strings.HasSuffix(p.height, "in") {
		trimmedHeight := strings.TrimSuffix(p.height, in)
		height := stringToInt(trimmedHeight)
		return height >= 59 && height <= 76
	} else if cm := "cm"; strings.HasSuffix(p.height, cm) {
		trimmedHeight := strings.TrimSuffix(p.height, cm)
		height := stringToInt(trimmedHeight)
		return height >= 150 && height <= 193
	}
	return false
}

// isHairColorValid returns true if hairColor is a # followed by exactly six characters 0-9 or a-f.
func (p *passportData) isHairColorValid() bool {
	return len(p.hairColor) == 7 && regexp.MustCompile(`^#[0-9-a-f]{6}$`).MatchString(p.hairColor)
}

// isEyeColorValid returns true if eyeColor is exactly one of: amb blu brn gry grn hzl oth.
func (p *passportData) isEyeColorValid() bool {
	switch p.eyeColor {
	case
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth":
		return true
	}
	return false
}

// isPassportIDValid returns true if passportID is a nine-digit number, including leading zeroes.
func (p *passportData) isPassportIDValid() bool {
	return len(p.passportID) == 9 && regexp.MustCompile(`^[0-9]{9}$`).MatchString(p.passportID)
}

func stringToInt(s string) int {
	converted, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("could not convert to int %s", s)
		return 1
	}
	return converted
}

func parsePassport(passport []string) passportData {
	var result passportData
	for _, data := range passport {
		for _, field := range strings.Split(data, " ") {
			kvPair := strings.Split(field, ":")
			key, value := kvPair[0], kvPair[1]

			switch key {
			case birthYear:
				result.birthYear = stringToInt(value)
			case issueYear:
				result.issueYear = stringToInt(value)
			case expirationYear:
				result.expirationYear = stringToInt(value)
			case height:
				result.height = value
			case hairColor:
				result.hairColor = value
			case eyeColor:
				result.eyeColor = value
			case passportID:
				result.passportID = value
			case countryID:
				result.countryID = stringToInt(value)
			}
		}

	}

	return result

}
