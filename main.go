package main

import "regexp"

func main() {
	str := "23-ab-48-caba-56-haha-15-l"
	testValidity(str)
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