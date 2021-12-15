package day15

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day15_1() {
	log.Info("day 15-1")
	f, _ := os.Open("day15/input.txt")

	cave := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		cave = append(cave, []int{})
		line := strings.Split(scanner.Text(), "")
		for l := range line {
			n, _ := strconv.Atoi(line[l])
			cave[counter] = append(cave[counter], n)
		}
		counter++
	}

	cost := search(cave)
	log.Infof("%v", cost)
}

func Day15_2() {
	log.Info("day 15-2")
	f, _ := os.Open("day15/input.txt")

	cave := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		cave = append(cave, []int{})
		line := strings.Split(scanner.Text(), "")
		for l := range line {
			n, _ := strconv.Atoi(line[l])
			cave[counter] = append(cave[counter], n)
		}
		counter++
	}
	// expand cave
	new_cave := make([][]int, len(cave)*5)
	for i := range new_cave {
		new_cave[i] = make([]int, len(cave)*5)
	}
	log.Infof("expanding cave")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for y := 0; y < len(cave); y++ {
				for x := 0; x < len(cave); x++ {
					risk := cave[y][x]
					risk += i
					if risk > 9 {
						risk -= 9
					}
					risk += j
					if risk > 9 {
						risk -= 9
					}
					new_cave[y+i*len(cave)][x+j*len(cave)] = risk
				}
			}
		}
	}
	log.Infof("expanding cave - done")
	//log.Infof("%v", new_cave)
	cave = new_cave
	cost := search(cave)
	log.Infof("%v", cost)
}
func search(cave [][]int) int {
	// Dijkstra algorithm
	max_nodes := len(cave) * len(cave)
	stride := len(cave)
	target_idx := stride*stride - 1 // bottom right corner is target
	dist := make([]int, max_nodes)
	prev := make([]int, max_nodes)
	Q := make([]int, max_nodes)
	for i := 0; i < max_nodes; i++ {
		dist[i] = math.MaxInt32
		prev[i] = -1
		Q[i] = i
	}
	dist[0] = 0

	for {
		min_dist := math.MaxInt32
		u := -1
		for i := 0; i < len(Q); i++ {
			idx := Q[i]
			if idx == -1 {
				continue
			}
			if dist[idx] < min_dist {
				min_dist = dist[idx]
				u = idx
			}
		}
		if u == -1 {
			log.Infof("reached end")
			break
		}
		Q[u] = -1
		if u == target_idx {
			log.Infof("reached target")
			break
		}
		neighbors := make([]int, 0)
		if u+1 < max_nodes && Q[u+1] != -1 {
			// check for wrap around
			y1 := u / len(cave)
			y2 := (u + 1) / len(cave)
			if y1 == y2 {
				neighbors = append(neighbors, u+1)
			}
		}
		if u-1 >= 0 && Q[u-1] != -1 {
			// check for wrap around
			y1 := u / len(cave)
			y2 := (u - 1) / len(cave)
			if y1 == y2 {
				neighbors = append(neighbors, u-1)
			}
		}
		if u+stride < max_nodes && Q[u+stride] != -1 {
			neighbors = append(neighbors, u+stride)
		}
		if u-stride >= 0 && Q[u-stride] != -1 {
			neighbors = append(neighbors, u-stride)
		}
		for n := range neighbors {
			idx := neighbors[n]
			x := idx % len(cave)
			y := idx / len(cave)

			alt := dist[u] + cave[y][x]
			if alt < dist[idx] {
				dist[idx] = alt
				prev[idx] = u
			}
		}
	}
	/*S := make([]int, 0)
	u := target_idx
	for u != -1 {
		S = append(S, u)
		u = prev[u]
	}
	for i := range S {
		idx := S[i]
		x := idx % len(cave)
		y := idx / len(cave)
		log.Infof("%d: %d,%d->%d", idx, y, x, cave[y][x])
	}*/
	return dist[target_idx]
}
