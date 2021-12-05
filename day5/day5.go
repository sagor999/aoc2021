package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day5_1() {
	log.Info("day 5-1")
	f, _ := os.Open("day5/input.txt")

	straight_lines_only := false
	grid_size := 1000
	grid := make([][]int, grid_size)
	for i := range grid {
		grid[i] = make([]int, grid_size)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")
		line_start := strings.Split(coords[0], ",")
		line_end := strings.Split(coords[1], ",")
		x1, _ := strconv.Atoi(line_start[0])
		y1, _ := strconv.Atoi(line_start[1])
		x2, _ := strconv.Atoi(line_end[0])
		y2, _ := strconv.Atoi(line_end[1])

		if straight_lines_only && !(x1 == x2 || y1 == y2) {
			continue
		}
		cx := x1
		cy := y1
		grid[cy][cx]++
		for {
			if cx != x2 {
				xinc := 0
				if x1 < x2 {
					xinc = 1
				} else if x2 < x1 {
					xinc = -1
				}
				cx += xinc
			}
			if cy != y2 {
				yinc := 0
				if y1 < y2 {
					yinc = 1
				} else if y2 < y1 {
					yinc = -1
				}
				cy += yinc
			}
			grid[cy][cx]++
			if cx == x2 && cy == y2 {
				break
			}
		}
	}
	res := 0
	for y := 0; y < grid_size; y++ {
		//log.Infof("%v", grid[y])
		for x := 0; x < grid_size; x++ {
			if grid[y][x] >= 2 {
				res++
			}
		}
	}

	log.Infof("res: %d", res)
}
