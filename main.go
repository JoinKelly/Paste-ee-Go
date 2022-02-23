package main

import (
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "23-ab-48-caba-56-haha-15-l"
	testValidity(str)
	averageNumber(str)
	wholeStory(str)
}

func extractText(str string) []string {
	const pattern = "[a-zA-Z]+"
	r := regexp.MustCompile(pattern)
	return r.FindAllString(str, -1)
}

func extractNumbers(str string) (numberList []int) {
	const pattern = "\\d+"
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(str, -1)

	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		numberList = append(numberList, num)
	}

	return
}

/**
- Difficulity: Easy
- Estimated time: 5 mins
- Implemented time: 7 mins
*/

func testValidity(str string) bool {
	const pattern = "^(\\d+\\-[a-zA-Z]+)((\\-\\d+\\-[a-zA-Z]+)+)?$"
	isValid, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	return isValid
}

/**
- Difficulity: Easy
- Estimated time: 5 mins
- Implemented time: 3 mins
*/

func averageNumber(str string) float64 {
	matches := extractNumbers(str)
	sum := 0
	for _, num := range matches {
		sum += num
	}

	return (float64(sum)) / (float64(len(matches)))
}

/**
- Difficulity: Easy
- Estimated time: 1 mins
- Implemented time: 1 mins
*/

func wholeStory(str string) string {
	matches := extractText(str)

	return strings.Join(matches, " ")
}