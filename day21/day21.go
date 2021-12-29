package day21

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day21_1() {
	log.Info("day 21-1")
	f, _ := os.Open("day21/input.txt")

	pl1_pos := 0
	pl2_pos := 0
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	pl1_pos, _ = strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
	scanner.Scan()
	pl2_pos, _ = strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])

	die := make([]int, 100)
	for i := range die {
		die[i] = i + 1
	}

	die_ind := 0
	pl1_score := 0
	pl2_score := 0
	for {
		d1 := deterministic_die(die_ind, die)
		die_ind++
		d2 := deterministic_die(die_ind, die)
		die_ind++
		d3 := deterministic_die(die_ind, die)
		die_ind++

		pl1_pos += d1 + d2 + d3
		for {
			if pl1_pos > 10 {
				pl1_pos -= 10
			}
			if pl1_pos <= 10 {
				break
			}
		}
		pl1_score += pl1_pos
		if pl1_score >= 1000 {
			break
		}

		d4 := deterministic_die(die_ind, die)
		die_ind++
		d5 := deterministic_die(die_ind, die)
		die_ind++
		d6 := deterministic_die(die_ind, die)
		die_ind++
		pl2_pos += d4 + d5 + d6
		for {
			if pl2_pos > 10 {
				pl2_pos -= 10
			}
			if pl2_pos <= 10 {
				break
			}
		}
		pl2_score += pl2_pos
		if pl2_score >= 1000 {
			break
		}
	}
	log.Infof("scores: %d, %d", pl1_score, pl2_score)
	if pl1_score > pl2_score {
		log.Infof("result: %d", die_ind*pl2_score)
	} else {
		log.Infof("result: %d", die_ind*pl1_score)
	}
}

func deterministic_die(ind int, die []int) int {
	return die[ind%100]
}
