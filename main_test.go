package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

type testTestValidity struct {
	arg      string
	expected bool
}

var testTestValidities = []testTestValidity{
	{"23-ab-48-caba-56-haha", true},
	{"23-ab", true},
	{"23-ab-48-caba-56-", false},
	{"23-ab-48-caba-56-ha1", false},
}

/**
- Difficulity: Easy
- Estimated time: 7 mins
- Implemented time: 7 mins
*/

func TestTestValidity(t *testing.T) {
	for _, test := range testTestValidities {
		if output := testValidity(test.arg); output != test.expected {
			fmt.Println(output, test)
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}

type testAverageNumber struct {
	arg      string
	expected float64
}

var testAverageNumbers = []testAverageNumber{
	{"23-ab-48-caba-56-haha", float64(23+48+56) / float64(3)},
	{"23-ab", 23},
	{"1-hello-2-world", 1.5},
}

/**
- Difficulity: Easy
- Estimated time: 7 mins
- Implemented time: 7 mins
*/
func TestAverageNumber(t *testing.T) {
	for _, test := range testAverageNumbers {
		if output := averageNumber(test.arg); output != test.expected {
			fmt.Println(output, test)
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}

type testWholeStory struct {
	arg      string
	expected string
}

var testWholeStories = []testWholeStory{
	{"23-ab-48-caba-56-haha", "ab caba haha"},
	{"23-ab", "ab"},
	{"1-hello-2-world", "hello world"},
}

/**
- Difficulity: Easy
- Estimated time: 7 mins
- Implemented time: 5 mins
*/
func TestWholeStory(t *testing.T) {
	for _, test := range testWholeStories {
		if output := wholeStory(test.arg); output != test.expected {
			fmt.Println(output, test)
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

type testStoryStats struct {
	arg      string
	expected expectedStoryStats
}

type expectedStoryStats struct {
	shortestWordList []string
	longestWordList  []string
	avgWordLength    float64
	sameSizeWordlist []string
}

var testStoryStatsList = []testStoryStats{
	{"23-ab-48-caba-56-haha", expectedStoryStats{
		[]string{"ab"}, []string{"caba", "haha"}, float64(2+4+4) / float64(3), nil,
	}},
	{"1-hello-2-world", expectedStoryStats{
		[]string{"hello", "world"}, []string{"hello", "world"}, 5, []string{"hello", "world"},
	}},
}

/**
- Difficulity: Easy
- Estimated time: 10 mins
- Implemented time: 10 mins
*/
func TestStoryStats(t *testing.T) {
	for _, test := range testStoryStatsList {
		output1, output2, output3, output4 := storyStats(test.arg)

		if !reflect.DeepEqual(output1, test.expected.shortestWordList) {
			fmt.Println(output1, test)
			t.Errorf("Output %q not equal to expected %q", output1, test.expected.shortestWordList)
		}

		if !reflect.DeepEqual(output2, test.expected.longestWordList) {
			fmt.Println(output2, test)
			t.Errorf("Output %q not equal to expected %q", output2, test.expected.longestWordList)
		}

		if output3 != test.expected.avgWordLength {
			fmt.Println(output3, test)
			t.Errorf("Output %f not equal to expected %f", output3, test.expected.avgWordLength)
		}

		if !reflect.DeepEqual(output4, test.expected.sameSizeWordlist) {
			fmt.Println(output2, test)
			t.Errorf("Output %q not equal to expected %q", output4, test.expected.sameSizeWordlist)
		}
	}
}

/**
- Difficulity: Easy
- Estimated time: 15 mins
- Implemented time: 20 mins
*/
func randomStrList(flag bool) (randomStrList []string) {
	rand.Seed(time.Now().UnixNano())
	stringCount := randInt(1, 3)

	for i := 0; i < stringCount; i++ {
		j := 0
		var sequence []string

		for j < randInt(1, 3) {
			s := randStr(randInt(1, 3))
			n := strconv.Itoa(rand.Intn(60))

			if flag {
				sequence = append(sequence, fmt.Sprintf("%s-%s", n, s))
			} else {
				sequence = append(sequence, fmt.Sprintf("%s-%s", s, n))
			}

			j++
		}

		randomStrList = append(randomStrList, strings.Join(sequence, "-"))
	}

	return
}

func randStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	if max-min <= 0 {
		return rand.Intn(5) + 1
	}

	return min + rand.Intn(max-min)
}