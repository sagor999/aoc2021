package day19

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
)

func Day19_1() {
	log.Info("day 19-1")
	f, _ := os.Open("day19/input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
}
