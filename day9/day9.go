package day9

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day9_1() {
	log.Info("day 9-1")
	f, _ := os.Open("day9/input.txt")

	scanner := bufio.NewScanner(f)
	heightmap := make([][]int, 0)
	row_i := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "")
		row := make([]int, len(nums))
		for i, n := range nums {
			row[i], _ = strconv.Atoi(n)
		}
		heightmap = append(heightmap, row)
		row_i++
	}
	res := 0
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			left := 10
			if (i - 1) >= 0 {
				left = heightmap[i-1][j]
			}
			right := 10
			if (i + 1) < len(heightmap) {
				right = heightmap[i+1][j]
			}
			bottom := 10
			if (j + 1) < len(heightmap[i]) {
				bottom = heightmap[i][j+1]
			}
			top := 10
			if (j - 1) >= 0 {
				top = heightmap[i][j-1]
			}
			num := heightmap[i][j]
			if num < left && num < right && num < top && num < bottom {
				res += num + 1
			}
		}
	}
	log.Infof("res: %d", res)
}

func Day9_2() {
	log.Info("day 9-2")
	f, _ := os.Open("day9/input.txt")

	scanner := bufio.NewScanner(f)
	heightmap := make([][]int, 0)
	row_i := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "")
		row := make([]int, len(nums))
		for i, n := range nums {
			row[i], _ = strconv.Atoi(n)
		}
		heightmap = append(heightmap, row)
		row_i++
	}
	basins := make([]int, 0)
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			left := 10
			if (i - 1) >= 0 {
				left = heightmap[i-1][j]
			}
			right := 10
			if (i + 1) < len(heightmap) {
				right = heightmap[i+1][j]
			}
			bottom := 10
			if (j + 1) < len(heightmap[i]) {
				bottom = heightmap[i][j+1]
			}
			top := 10
			if (j - 1) >= 0 {
				top = heightmap[i][j-1]
			}
			num := heightmap[i][j]
			if num < left && num < right && num < top && num < bottom {
				values := make([][]int, len(heightmap))
				for k := 0; k < len(values); k++ {
					values[k] = make([]int, len(heightmap[k]))
				}
				for k := 0; k < len(values); k++ {
					for l := 0; l < len(values[k]); l++ {
						values[k][l] = 0
					}
				}
				calcBasinSize(i, j, heightmap, &values)
				size := 0
				for k := 0; k < len(values); k++ {
					for l := 0; l < len(values[k]); l++ {
						size += values[k][l]
					}
				}

				basins = append(basins, size)
			}
		}
	}
	sort.Ints(basins)
	len := len(basins)
	log.Infof("res: %d", basins[len-1]*basins[len-2]*basins[len-3])
}

func calcBasinSize(i int, j int, heightmap [][]int, values *[][]int) {
	(*values)[i][j] = 1

	left := 10
	if (i - 1) >= 0 {
		left = heightmap[i-1][j]
	}
	right := 10
	if (i + 1) < len(heightmap) {
		right = heightmap[i+1][j]
	}
	bottom := 10
	if (j + 1) < len(heightmap[i]) {
		bottom = heightmap[i][j+1]
	}
	top := 10
	if (j - 1) >= 0 {
		top = heightmap[i][j-1]
	}
	num := heightmap[i][j]
	if num < left && left < 9 {
		calcBasinSize(i-1, j, heightmap, values)
	}
	if num < right && right < 9 {
		calcBasinSize(i+1, j, heightmap, values)
	}
	if num < bottom && bottom < 9 {
		calcBasinSize(i, j+1, heightmap, values)
	}
	if num < top && top < 9 {
		calcBasinSize(i, j-1, heightmap, values)
	}
}
