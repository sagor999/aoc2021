package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day2_1() {
	f, _ := os.Open("day2/input.txt")

	scanner := bufio.NewScanner(f)

	x := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " ")
		log.Infof("%v", sep)
		i, _ := strconv.Atoi(sep[1])
		switch sep[0] {
		case "forward":
			x += i
		case "down":
			y += i
		case "up":
			y -= i
		}
	}
	log.Info("res: ", x*y)
}

func Day2_2() {
	f, _ := os.Open("day2/input.txt")

	scanner := bufio.NewScanner(f)

	x := 0
	aim := 0
	d := 0

	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " ")
		log.Infof("%v", sep)
		i, _ := strconv.Atoi(sep[1])
		switch sep[0] {
		case "forward":
			x += i
			d += aim * i
		case "down":
			aim += i
		case "up":
			aim -= i
		}
	}
	log.Info("res: ", x*d)
}
