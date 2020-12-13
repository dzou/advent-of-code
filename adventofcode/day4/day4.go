package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var passports []map[string]string

	currMap := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine := scanner.Text()

		if len(currLine) == 0 {
			passports = append(passports, currMap)
			currMap = make(map[string]string)
		} else {
			tokens := strings.Split(currLine, " ")
			for _, token := range tokens {
				parts := strings.Split(token, ":")
				currMap[parts[0]] = parts[1]
			}
		}
	}
	if len(currMap) > 0 {
		passports = append(passports, currMap)
	}

	answer := solve(passports)
	fmt.Println(answer)
}

func solve(passports []map[string]string) int {

	var ALL_EYE_COLORS = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	count := 0

	for _, passport := range passports {
		_, ok := passport["cid"]
		if len(passport) == 8 || len(passport) == 7 && !ok {

			year, err := strconv.Atoi(passport["byr"])
			if err != nil || year < 1920 || year > 2002 {
				continue
			}

			year, err = strconv.Atoi(passport["iyr"])
			if err != nil || year < 2010 || year > 2020 {
				continue
			}

			year, err = strconv.Atoi(passport["eyr"])
			if err != nil || year < 2020 || year > 2030 {
				continue
			}

			heightStr := passport["hgt"]
			if strings.HasSuffix(heightStr, "in") {
				heightStr = strings.Replace(heightStr, "in", "", 1)
				heightValue, err := strconv.Atoi(heightStr)
				if err != nil || heightValue < 59 || heightValue > 76 {
					continue
				}
			} else if strings.HasSuffix(heightStr, "cm") {
				heightStr = strings.Replace(heightStr, "cm", "", 1)
				heightValue, err := strconv.Atoi(heightStr)
				if err != nil || heightValue < 150 || heightValue > 193 {
					continue
				}
			} else {
				continue
			}

			hairColor := passport["hcl"]
			if strings.HasPrefix(hairColor, "#") && len(hairColor) == 7 {
				for i := 1; i < len(hairColor); i++ {
					if (hairColor[i] >= '0' && hairColor[i] <= '9') || (hairColor[i] >= 'a' && hairColor[i] <= 'f') {

					} else {
						continue
					}
				}
			} else {
				continue
			}

			eyeColor := passport["ecl"]
			if !ALL_EYE_COLORS[eyeColor] {
				continue
			}

			passportId := passport["pid"]
			parsedPassportId, err := strconv.Atoi(passport["pid"])
			if err != nil || len(passportId) != 9 || parsedPassportId < 0 {
				continue
			}
			count++

		}
	}

	return count
}
