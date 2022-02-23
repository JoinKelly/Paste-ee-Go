package main

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	str := "23-ab-48-caba-56-haha-15-l"
	// testValidity(str)
	// averageNumber(str)
	// wholeStory(str)
	storyStats(str)
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

/**
- Difficulity: Easy
- Estimated time: 15 mins
- Implemented time: 10 mins
*/

func storyStats(str string) (shortestWordList []string, longestWordList []string, avgWordLength float64, sameSizeWordlist []string) {
	words := extractText(str)
	if len(words) == 0 {
		return nil, nil, 0, nil
	}
	m := make(map[int][]string)
	var keys []int
	wLength := 0
	for _, w := range words {
		m[len(w)] = append(m[len(w)], w)
		keys = append(keys, len(w))
		wLength += len(w)
	}
	sort.Ints(keys)
	shortestWordList = m[keys[0]]
	longestWordList = m[keys[len(keys)-1]]
	avgWordLength = float64(wLength) / float64(len(words))
	avgKey := int(math.Round(avgWordLength))
	if val, ok := m[avgKey]; ok {
		sameSizeWordlist = val
	}
	return
}