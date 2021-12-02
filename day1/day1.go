package day1

import (
	"bufio"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func Day1() {
	f, _ := os.Open("day1/input.txt")

	scanner := bufio.NewScanner(f)

	// Read and print each line in the file

	prev_i := 10000000
	num_incr := 0
	arr := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		arr = append(arr, i)
		if i > prev_i {
			num_incr++
		}
		prev_i = i
	}
	log.Info("Number of incr: ", num_incr)

	num_incr = 0
	// Part 2
	for i := 0; i < len(arr)-3; i++ {
		j := i + 1
		a := arr[i] + arr[i+1] + arr[i+2]
		b := arr[j] + arr[j+1] + arr[j+2]
		if b > a {
			num_incr++
		}
	}
	log.Info("Number of incr: ", num_incr)
}
