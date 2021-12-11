package day11

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day11_1() {
	log.Info("day 11-1")
	f, _ := os.Open("day11/input.txt")

	scanner := bufio.NewScanner(f)
	octopii := make([][]int, 10)
	for i := range octopii {
		octopii[i] = make([]int, 10)
	}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		for j, c := range chars {
			n, _ := strconv.Atoi(c)
			octopii[i][j] = n
		}
		i++
	}
	num_flashes := 0
	for step := 0; step < 100; step++ {
		for i := range octopii {
			for j := range octopii[i] {
				octopii[i][j]++
			}
		}
		for {
			had_flash := false
			for i := range octopii {
				for j := range octopii[i] {
					if octopii[i][j] > 9 {
						had_flash = true
						num_flashes++
						flash(&octopii, i, j)
					}
				}
			}
			if !had_flash {
				break
			}
		}
		for i := range octopii {
			for j := range octopii[i] {
				if octopii[i][j] == -1 {
					octopii[i][j] = 0
				}
			}
		}
	}
	log.Infof("num_flashes: %d", num_flashes)
}

func Day11_2() {
	log.Info("day 11-2")
	f, _ := os.Open("day11/input.txt")

	scanner := bufio.NewScanner(f)
	octopii := make([][]int, 10)
	for i := range octopii {
		octopii[i] = make([]int, 10)
	}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		for j, c := range chars {
			n, _ := strconv.Atoi(c)
			octopii[i][j] = n
		}
		i++
	}
	res := 0
	for step := 0; step < 10000; step++ {
		num_flashes := 0
		for i := range octopii {
			for j := range octopii[i] {
				octopii[i][j]++
			}
		}
		for {
			had_flash := false
			for i := range octopii {
				for j := range octopii[i] {
					if octopii[i][j] > 9 {
						had_flash = true
						num_flashes++
						flash(&octopii, i, j)
					}
				}
			}
			if !had_flash {
				break
			}
		}
		for i := range octopii {
			for j := range octopii[i] {
				if octopii[i][j] == -1 {
					octopii[i][j] = 0
				}
			}
		}
		if num_flashes == 100 {
			res = step
			break
		}
	}
	log.Infof("res: %d", res+1)
}

func flash(octopii *[][]int, i int, j int) {
	(*octopii)[i][j] = -1
	if i > 0 {
		if (*octopii)[i-1][j] != -1 {
			(*octopii)[i-1][j]++
		}
		if j > 0 {
			if (*octopii)[i-1][j-1] != -1 {
				(*octopii)[i-1][j-1]++
			}
		}
		if j < 9 {
			if (*octopii)[i-1][j+1] != -1 {
				(*octopii)[i-1][j+1]++
			}
		}
	}
	if i < 9 {
		if (*octopii)[i+1][j] != -1 {
			(*octopii)[i+1][j]++
		}
		if j > 0 {
			if (*octopii)[i+1][j-1] != -1 {
				(*octopii)[i+1][j-1]++
			}
		}
		if j < 9 {
			if (*octopii)[i+1][j+1] != -1 {
				(*octopii)[i+1][j+1]++
			}
		}
	}
	if j > 0 {
		if (*octopii)[i][j-1] != -1 {
			(*octopii)[i][j-1]++
		}
	}
	if j < 9 {
		if (*octopii)[i][j+1] != -1 {
			(*octopii)[i][j+1]++
		}
	}
}
