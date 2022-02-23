package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/**
- Difficulity: Easy
- Estimated time: 10 mins
- Implemented time: 7 mins
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