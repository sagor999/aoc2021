package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day13_1() {
	log.Info("day 13-1")
	f, _ := os.Open("day13/input.txt")

	scanner := bufio.NewScanner(f)
	paper_size := 2000
	paper := make([][]int, paper_size)
	for i := range paper {
		paper[i] = make([]int, paper_size)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "fold") {
			sep := strings.Split(line, "=")
			if sep[0][len(sep[0])-1] == 'x' {
				x, _ := strconv.Atoi(sep[1])
				foldX(&paper, x)
			} else {
				y, _ := strconv.Atoi(sep[1])
				foldY(&paper, y)
			}
		} else {
			sep := strings.Split(line, ",")
			x, _ := strconv.Atoi(sep[0])
			y, _ := strconv.Atoi(sep[1])
			paper[y][x] = 1
		}
	}
	num_dots := 0
	max_y := 0
	max_x := 0
	for y := 0; y < paper_size; y++ {
		for x := 0; x < paper_size; x++ {
			if paper[y][x] == 1 {
				num_dots++
				if y > max_y {
					max_y = y
				}
				if x > max_x {
					max_x = x
				}
			}
		}
	}
	log.Infof("%v", num_dots)
	for y := 0; y < max_y+1; y++ {
		for x := 0; x < max_x+1; x++ {
			char := "x"
			if paper[y][x] == 0 {
				char = " "
			}
			fmt.Printf("%s", char)
		}
		fmt.Println("")
	}
}

func foldY(paper *[][]int, f int) {
	for i := 1; i < 1000; i++ {
		if f-i < 0 || f+i >= len((*paper)[0]) {
			break
		}
		for x := 0; x < len((*paper)[0]); x++ {
			if (*paper)[f+i][x] == 1 {
				(*paper)[f-i][x] = (*paper)[f+i][x]
				(*paper)[f+i][x] = 0
			}
		}
	}
}

func foldX(paper *[][]int, f int) {
	for i := 1; i < 1000; i++ {
		if f-i < 0 || f+i >= len((*paper)[0]) {
			break
		}
		for y := 0; y < len((*paper)[0]); y++ {
			if (*paper)[y][f+i] == 1 {
				(*paper)[y][f-i] = (*paper)[y][f+i]
				(*paper)[y][f+i] = 0
			}
		}
	}
}
