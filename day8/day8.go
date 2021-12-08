package day8

import (
	"bufio"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day8_1() {
	log.Info("day 8-1")
	f, _ := os.Open("day8/input.txt")

	scanner := bufio.NewScanner(f)
	num := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		words := strings.Split(line[1], " ")
		for _, word := range words {
			word_len := len(word)
			if word_len == 7 || word_len == 4 || word_len == 3 || word_len == 2 {
				num++
			}
		}
	}
	log.Infof("res: %d", num)
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Day8_2() {
	log.Info("day 8-2")
	f, _ := os.Open("day8/input.txt")

	scanner := bufio.NewScanner(f)
	res := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		words := strings.Split(line[0], " ")
		num_str := make([]string, 10)
		for _, word := range words {
			word_len := len(word)
			if word_len == 7 {
				num_str[8] = sortString(word)
			} else if word_len == 4 {
				num_str[4] = sortString(word)
			} else if word_len == 3 {
				num_str[7] = sortString(word)
			} else if word_len == 2 {
				num_str[1] = sortString(word)
			}
		}
		// now lets find number 3, since it should have all leters from number 7
		for w_i, word := range words {
			word_len := len(word)
			if word_len == 5 {
				if containsAll(word, num_str[7]) {
					num_str[3] = sortString(word)
					words = remove(words, w_i)
					break
				}
			}
		}
		// now lets find 9, since it should have all letters from 4
		for w_i, word := range words {
			word_len := len(word)
			if word_len == 6 {
				if containsAll(word, num_str[4]) {
					num_str[9] = sortString(word)
					words = remove(words, w_i)
					break
				}
			}
		}
		// now find 0, since it should have all letters from 7
		for w_i, word := range words {
			word_len := len(word)
			if word_len == 6 {
				if containsAll(word, num_str[7]) {
					num_str[0] = sortString(word)
					words = remove(words, w_i)
					break
				}
			}
		}
		// now find 6, it is only one left with 6 digits
		for w_i, word := range words {
			word_len := len(word)
			if word_len == 6 {
				num_str[6] = sortString(word)
				words = remove(words, w_i)
				break
			}
		}
		// now find 5, since it will have all letters from 9
		for w_i, word := range words {
			word_len := len(word)
			if word_len == 5 {
				if containsAll(num_str[6], word) {
					num_str[5] = sortString(word)
					words = remove(words, w_i)
					break
				}
			}
		}
		// now last number is 2
		for _, word := range words {
			word_len := len(word)
			if word_len == 5 {
				num_str[2] = sortString(word)
				break
			}
		}

		words = strings.Split(line[1], " ")
		num := 0
		mul := 1
		for w := len(words) - 1; w >= 0; w-- {
			word := sortString(words[w])
			for i := 0; i < 10; i++ {
				if word == num_str[i] {
					num += i * mul
					mul *= 10
				}
			}
		}
		//log.Infof("num: %d", num)
		res += num
	}
	log.Infof("res: %d", res)
}

func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func containsAll(word string, s string) bool {
	res := true
	runeArray := []rune(s)
	for _, c := range runeArray {
		res = res && strings.ContainsRune(word, c)
	}
	return res
}
