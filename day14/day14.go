package day14

import (
	"bufio"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day14_1() {
	log.Info("day 14-1")
	f, _ := os.Open("day14/input.txt")

	scanner := bufio.NewScanner(f)
	pairs := make(map[string]string)
	scanner.Scan()
	polymer := strings.Split(scanner.Text(), "")
	scanner.Scan() // skip empty line
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " -> ")
		pairs[sep[0]] = sep[1]
	}

	steps := 10
	for step := 0; step < steps; step++ {
		new_polymer := []string{}
		for i := 0; i < len(polymer)-1; i++ {
			pair := polymer[i] + polymer[i+1]
			elem := pairs[pair]
			if i == 0 {
				new_polymer = append(new_polymer, polymer[i])
			}
			new_polymer = append(new_polymer, elem)
			new_polymer = append(new_polymer, polymer[i+1])
		}
		polymer = new_polymer
		log.Infof("step: %d, polymer len: %d", step, len(polymer))
	}

	//log.Infof("%v", polymer)

	elements := make(map[string]int64)
	for i := 0; i < len(polymer); i++ {
		elements[polymer[i]]++
	}
	//log.Infof("%v", elements)

	most_common_elem_q := int64(0)
	least_common_elem_q := int64(math.MaxInt64)
	for _, v := range elements {
		if v > most_common_elem_q {
			most_common_elem_q = v
		}
		if v < least_common_elem_q {
			least_common_elem_q = v
		}
	}
	log.Infof("res: %d", most_common_elem_q-least_common_elem_q)

}

func Day14_2() {
	log.Info("day 14-2")
	f, _ := os.Open("day14/input.txt")

	scanner := bufio.NewScanner(f)
	pairs := make(map[string]string)
	scanner.Scan()
	polymer := strings.Split(scanner.Text(), "")
	scanner.Scan() // skip empty line
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, " -> ")
		pairs[sep[0]] = sep[1]
	}

	pairs_q := make(map[string]int)
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i] + polymer[i+1]
		pairs_q[pair]++
	}
	//log.Infof("%v", pairs_q)

	steps := 40
	for step := 0; step < steps; step++ {
		new_pairs_q := make(map[string]int)
		for pair, pair_q := range pairs_q {
			new_elem := pairs[pair]
			new_pairs_q[string(pair[0])+new_elem] += pair_q
			new_pairs_q[new_elem+string(pair[1])] += pair_q
		}
		pairs_q = new_pairs_q
		//log.Infof("step: %d", step)
		//log.Infof("%v", pairs_q)
	}

	elements := make(map[string]int64)
	for pair, pair_q := range pairs_q {
		elements[string(pair[0])] += int64(pair_q)
	}
	elements[polymer[len(polymer)-1]]++
	//log.Infof("%v", elements)

	most_common_elem_q := int64(0)
	least_common_elem_q := int64(math.MaxInt64)
	for _, v := range elements {
		if v > most_common_elem_q {
			most_common_elem_q = v
		}
		if v < least_common_elem_q {
			least_common_elem_q = v
		}
	}
	log.Infof("res: %d", most_common_elem_q-least_common_elem_q)

}
