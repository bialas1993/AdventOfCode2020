package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

const (
	UnitCentimeters = "cm"
	UnitInches      = "in"
)

type Passport struct {
	PID            string
	CID            int
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         Height
	HairColor      string
	EyeColor       string
}

type Height struct {
	Value int
	Unit  string
}

func (h Height) IsValid() bool {
	if h.Value > 0 {
		if h.Unit == UnitInches {
			return h.Value >= 59 && h.Value <= 76
		}

		if h.Unit == UnitCentimeters {
			return h.Value >= 150 && h.Value <= 193
		}
	}

	return false
}

// IsValid returns true if the passport data is correct.
func (p Passport) IsValid() (bool, error) {
	if !(p.BirthYear >= 1920 && p.BirthYear <= 2002) {
		return false, errors.New("birth year is not valid")
	}

	if !(p.IssueYear >= 2010 && p.IssueYear <= 2020) {
		return false, errors.New("issue year is not valid")
	}

	if !(p.ExpirationYear >= 2020 && p.ExpirationYear <= 2030) {
		return false, errors.New("expiration year is not valid")
	}

	if !p.Height.IsValid() {
		return false, errors.New("height is not valid")
	}

	var validHair = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if !(len(p.HairColor) > 0 && validHair.MatchString(p.HairColor)) {
		return false, errors.New("hair color is not valid")
	}

	var validEyeColor = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	if !(validEyeColor.MatchString(p.EyeColor)) {
		return false, errors.New("eye color is not valid")
	}

	var validPID = regexp.MustCompile(`^[0-9]{9}$`)
	if !(validPID.MatchString(p.PID)) {
		return false, errors.New("PID is not valid")
	}

	return true, nil
}

func checkErr(err error) {
	if err != nil {
		panic("cannot parse data " + err.Error())
	}
}

func createPassport(data string) Passport {
	p := Passport{}

	for _, field := range strings.Fields(data) {
		fieldData := strings.Split(field, ":")
		key, value := fieldData[0], fieldData[1]

		switch key {
		case "pid":
			p.PID = value
		case "cid":
			i, err := strconv.Atoi(value)
			checkErr(err)
			p.CID = i
		case "byr":
			i, err := strconv.Atoi(value)
			checkErr(err)
			p.BirthYear = i
		case "iyr":
			i, err := strconv.Atoi(value)
			checkErr(err)
			p.IssueYear = i
		case "eyr":
			i, err := strconv.Atoi(value[:4])
			checkErr(err)
			p.ExpirationYear = i
		case "hgt":
			if strings.HasSuffix(value, UnitInches) || strings.HasSuffix(value, UnitCentimeters) {
				i, err := strconv.Atoi(value[:len(value)-2])
				if err == nil {
					p.Height = Height{i, value[len(value)-2:]}
				}
			}
		case "hcl":
			if strings.HasPrefix(value, "#") {
				p.HairColor = value
			}
		case "ecl":
			p.EyeColor = value
		default:
			panic(fmt.Sprintf("not support [%s] %s = %s", field, key, value))
		}
	}

	return p
}

func main() {
	passports := []Passport{}
	data := ""
	rows := utils.FileRows("/day04/1/input.txt")

	for _, row := range rows {
		if len(strings.Trim(row, " ")) == 0 {
			passport := createPassport(strings.Trim(data, " "))

			if ok, _ := passport.IsValid(); ok {
				passports = append(passports, passport)
			}

			data = ""
		}

		data = data + row + " "
	}

	fmt.Printf("Correct passports: %d\n", len(passports))
}
