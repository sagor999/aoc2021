package day3

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day3_1() {
	f, _ := os.Open("day3/input.txt")

	scanner := bufio.NewScanner(f)

	len := 11
	num_0 := make([]int, len+1)
	num_1 := make([]int, len+1)
	gamma_rate := 0
	epsilon_rate := 0

	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, "")
		//log.Infof("%v", sep)
		for i, v := range sep {
			if v == "0" {
				num_0[i]++
			} else {
				num_1[i]++
			}
		}
	}
	for i := len; i >= 0; i-- {
		bit := 0
		if num_1[i] > num_0[i] {
			bit = 1
		}
		log.Info("bit: ", bit)
		gamma_rate |= bit << uint(len-i)
		bit = 1 - bit
		epsilon_rate |= bit << uint(len-i)
	}
	log.Info("gamma: ", gamma_rate)
	log.Info("epsilon: ", epsilon_rate)
	log.Info("result: ", gamma_rate*epsilon_rate)
}

func filter(input []string, bit_pos int, oxygen bool) []string {
	var result []string
	num_1 := 0
	num_0 := 0
	for _, v := range input {
		sep := strings.Split(v, "")
		if sep[bit_pos] == "1" {
			num_1++
		} else {
			num_0++
		}
	}
	target_bit := ""
	if oxygen {
		target_bit = "0"
		if num_1 >= num_0 {
			target_bit = "1"
		}
	} else {
		target_bit = "1"
		if num_0 <= num_1 {
			target_bit = "0"
		}
	}
	for _, v := range input {
		sep := strings.Split(v, "")
		if sep[bit_pos] == target_bit {
			result = append(result, v)
		}
	}
	return result
}

func Day3_2() {
	f, _ := os.Open("day3/input.txt")

	scanner := bufio.NewScanner(f)

	array := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line)
	}
	oxygen := 0
	co2 := 0
	oxygen_array := array
	co2_array := array
	target_bit := 0
	for {
		oxygen_array = filter(oxygen_array, target_bit, true)
		target_bit++
		if len(oxygen_array) == 1 {
			break
		}
	}
	target_bit = 0
	for {
		co2_array = filter(co2_array, target_bit, false)
		target_bit++
		if len(co2_array) == 1 {
			break
		}
	}
	log.Info("oxygen: ", oxygen_array[0])
	log.Info("co2: ", co2_array[0])
	for i := len(oxygen_array[0]) - 1; i >= 0; i-- {
		oxygen |= ((int(oxygen_array[0][i] - '0')) << uint(len(oxygen_array[0])-1-i))
	}
	for i := len(co2_array[0]) - 1; i >= 0; i-- {
		co2 |= ((int(co2_array[0][i] - '0')) << uint(len(co2_array[0])-1-i))
	}
	log.Info("result oxygen: ", oxygen)
	log.Info("result co2: ", co2)
	log.Info("result: ", oxygen*co2)
}
