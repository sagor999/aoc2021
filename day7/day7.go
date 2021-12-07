package day7

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day7_1() {
	log.Info("day 7-1")
	f, _ := os.Open("day7/input.txt")

	min_pos := 1000000
	max_pos := 0
	crabs := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, ",")
		for _, v := range sep {
			i, _ := strconv.Atoi(v)
			crabs = append(crabs, i)
			if i < min_pos {
				min_pos = i
			}
			if i > max_pos {
				max_pos = i
			}
		}
	}

	min_fuel := 100000000
	for i := min_pos; i <= max_pos; i++ {
		fuel := 0
		for _, crab_pos := range crabs {
			fuel += int(math.Abs(float64(crab_pos - i)))
		}
		if fuel < min_fuel {
			min_fuel = fuel
		}
	}

	log.Infof("min fuel: %d", min_fuel)

}

func Day7_2() {
	log.Info("day 7-2")
	f, _ := os.Open("day7/input.txt")

	min_pos := 1000000
	max_pos := 0
	crabs := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, ",")
		for _, v := range sep {
			i, _ := strconv.Atoi(v)
			crabs = append(crabs, i)
			if i < min_pos {
				min_pos = i
			}
			if i > max_pos {
				max_pos = i
			}
		}
	}

	min_fuel := 100000000
	for i := min_pos; i <= max_pos; i++ {
		fuel := 0
		for _, crab_pos := range crabs {
			num_steps := int(math.Abs(float64(crab_pos - i)))
			for s := 0; s < num_steps; s++ {
				fuel += (s + 1)
			}
		}
		if fuel < min_fuel {
			min_fuel = fuel
		}
	}

	log.Infof("min fuel: %d", min_fuel)

}
