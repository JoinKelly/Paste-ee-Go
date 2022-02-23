package main

import (
	"fmt"
	"math/rand"
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