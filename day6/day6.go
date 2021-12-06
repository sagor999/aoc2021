package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day6_1() {
	log.Info("day 6-1")
	f, _ := os.Open("day6/input.txt")

	fishes := make([]int8, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, ",")
		for _, v := range sep {
			i, _ := strconv.Atoi(v)
			fishes = append(fishes, int8(i))
		}
	}
	log.Infof("Initial state: %v", fishes)
	num_days := 80
	for i := 0; i < num_days; i++ {
		cur_range := len(fishes)
		for fish := 0; fish < cur_range; fish++ {
			fishes[fish]--
			if fishes[fish] < 0 {
				fishes[fish] = 6
				fishes = append(fishes, 8)
			}
		}
		log.Infof("Iter: %d, Fish count: %d", i, len(fishes))
	}
	log.Infof("Final fish count: %d", len(fishes))
}

func Day6_2() {
	log.Info("day 6-2")
	f, _ := os.Open("day6/input.txt")

	fishes := make([]int, 9)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, ",")
		for _, v := range sep {
			i, _ := strconv.Atoi(v)
			fishes[i]++
		}
	}
	log.Infof("Initial state: %v", fishes)
	num_days := 256
	for i := 0; i < num_days; i++ {
		new_fishes := make([]int, 9)
		new_fishes[8] = fishes[0]
		for j := 1; j < 9; j++ {
			new_fishes[j-1] = fishes[j]
		}
		new_fishes[6] += fishes[0]
		fishes = new_fishes
		total_fish := 0
		for j := 0; j < 9; j++ {
			total_fish += fishes[j]
		}
		log.Infof("Iter: %d, Fish count: %d", i, total_fish)
	}
	total_fish := 0
	for j := 0; j < 9; j++ {
		total_fish += fishes[j]
	}
	log.Infof("Final fish count: %d", total_fish)
}
