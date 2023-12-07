package day_01

import (
	"fmt"
	"strconv"
	"sync"
	"unicode"
)

var numbers = map[string]map[string]string{
	"o": {"one": "1"},
	"t": {"two": "2", "three": "3"},
	"f": {"four": "4", "five": "5"},
	"s": {"six": "6", "seven": "7"},
	"e": {"eight": "8"},
	"n": {"nine": "9"},
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func isNumber(input string) (string, bool) {
	if !unicode.IsLetter(rune(input[0])) {
		return "", false
	}

	possibles := numbers[string(input[0])]
	for k, num := range possibles {
		if len(input) < len(k) {
			continue
		}
		if input[:len(k)] == k {
			return num, true
		}
	}
	return "", false
}

func parseNum(input string) int {
	wg := sync.WaitGroup{}

	wg.Add(1)
	first, second := "", ""
	go func(input string) {
		for i := 0; i < len(input); i++ {
			if unicode.IsLetter(rune(input[i])) {
				if num, ok := isNumber(input[i:]); ok {
					first = num
					break
				}

				continue

			}

			first = string(input[i])
			break
		}
		wg.Done()
	}(input)

	wg.Add(1)
	go func(input string) {
		for i := len(input) - 1; i >= 0; i-- {
			if unicode.IsLetter(rune(input[i])) {
				if num, ok := isNumber(input[i:]); ok {
					second = num
					break
				}
				continue
			}

			second = string(input[i])
			break
		}
		wg.Done()
	}(input)

	wg.Wait()

	out := fmt.Sprintf("%s%s", first, second)
	outI, _ := strconv.Atoi(out)

	return outI
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	res := 0

	for _, line := range input {
		res += parseNum(line)
	}
	return fmt.Sprintf("%d", res)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	res := 0

	for _, line := range input {
		res += parseNum(line)
	}
	return fmt.Sprintf("%d", res)
}
