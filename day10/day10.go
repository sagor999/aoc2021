package day10

import (
	"bufio"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day10_1() {
	log.Info("day 10-1")
	f, _ := os.Open("day10/input.txt")

	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		stack := make([]string, 0)
		//log.Infof("%v", chars)
		for _, c := range chars {
			if c == "(" {
				stack = append(stack, ")")
			} else if c == "[" {
				stack = append(stack, "]")
			} else if c == "{" {
				stack = append(stack, "}")
			} else if c == "<" {
				stack = append(stack, ">")
			} else {
				//log.Infof("stack: %s, %s", c, stack)
				if stack[len(stack)-1] == c {
					stack = stack[:len(stack)-1]
				} else {
					switch c {
					case ")":
						score += 3
					case "]":
						score += 57
					case "}":
						score += 1197
					case ">":
						score += 25137
					}
					break
				}
			}
		}
	}
	log.Infof("score: %d", score)
}

func Day10_2() {
	log.Info("day 10-2")
	f, _ := os.Open("day10/input.txt")

	scanner := bufio.NewScanner(f)
	scores := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		stack := make([]string, 0)
		//log.Infof("%v", chars)
		corrupted := false
		for _, c := range chars {
			if c == "(" {
				stack = append(stack, ")")
			} else if c == "[" {
				stack = append(stack, "]")
			} else if c == "{" {
				stack = append(stack, "}")
			} else if c == "<" {
				stack = append(stack, ">")
			} else {
				//log.Infof("stack: %s, %s", c, stack)
				if stack[len(stack)-1] == c {
					stack = stack[:len(stack)-1]
				} else {
					corrupted = true
					break
				}
			}
		}
		if corrupted {
			continue
		}
		//log.Infof("stack: %v", stack)
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score *= 5
			switch stack[i] {
			case ")":
				score += 1
			case "]":
				score += 2
			case "}":
				score += 3
			case ">":
				score += 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	log.Infof("score: %d", scores[len(scores)/2])
}
