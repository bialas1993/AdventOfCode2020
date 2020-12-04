package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

type Passport struct {
	PID            string
	CID            int
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
}

// IsValid returns true if the passport data is correct.
func (p Passport) IsValid() bool {
	return !(p.PID == "" || p.BirthYear == 0 || p.IssueYear == 0 || p.ExpirationYear == 0 || p.Height == "" || p.HairColor == "" || p.EyeColor == "")
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
			i, err := strconv.Atoi(value)
			checkErr(err)
			p.ExpirationYear = i
		case "hgt":
			p.Height = value
		case "hcl":
			p.HairColor = value
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
	rows := utils.FileRows("/day04/2/input.txt")

	for _, row := range rows {
		if len(strings.Trim(row, " ")) == 0 {
			passport := createPassport(strings.Trim(data, " "))
			if passport.IsValid() {
				passports = append(passports, passport)
			}

			data = ""
		}

		data = data + row + " "
	}

	fmt.Printf("Correct passports: %d\n", len(passports))
}
