package day17

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day17_1() {
	log.Info("day 17-1")
	f, _ := os.Open("day17/input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := scanner.Text()
	sep := strings.Split(input, ", ")
	x1, x2, y1, y2 := 0, 0, 0, 0
	x_str := strings.Split(strings.Split(sep[0], "x=")[1], "..")
	x1, _ = strconv.Atoi(x_str[0])
	x2, _ = strconv.Atoi(x_str[1])
	y_str := strings.Split(strings.Split(sep[1], "y=")[1], "..")
	y1, _ = strconv.Atoi(y_str[0])
	y2, _ = strconv.Atoi(y_str[1])
	log.Infof("target area: %d..%d, %d..%d", x1, x2, y1, y2)

	// brute force
	max_y := math.MinInt32
	num_hits := 0
	for x := 0; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			in_target, my := solve(x, y, x1, x2, y1, y2)
			if in_target {
				num_hits++
				if my > max_y {
					max_y = my
				}
			}
		}
	}
	log.Infof("max y: %d, num_hits: %d", max_y, num_hits)
}

// solve, return true if within range and max Y achieved
func solve(x, y int, x1, x2, y1, y2 int) (bool, int) {
	max_y := math.MinInt32

	pos_x := 0
	pos_y := 0
	vel_x := x
	vel_y := y
	for step := 0; step < 1000; step++ {
		pos_x = pos_x + vel_x
		pos_y = pos_y + vel_y
		hit_target := pos_x >= x1 && pos_x <= x2 && pos_y >= y1 && pos_y <= y2
		if pos_y > max_y {
			max_y = pos_y
		}
		if hit_target {
			//log.Infof("%d, (%d,%d)", max_y, x, y)
			return true, max_y
		}
		if vel_x > 0 {
			vel_x = vel_x - 1
		} else if vel_x < 0 {
			vel_x = vel_x + 1
		}
		vel_y = vel_y - 1
	}

	return false, 0
}
