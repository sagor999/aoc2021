package day12

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

type node struct {
	links []string
}

func Day12_1() {
	log.Info("day 12-1")
	f, _ := os.Open("day12/input.txt")

	scanner := bufio.NewScanner(f)
	caves := make(map[string]node)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, "-")
		c1 := sep[0]
		c2 := sep[1]
		if _, ok := caves[c1]; !ok {
			caves[c1] = node{}
		}
		if _, ok := caves[c2]; !ok {
			caves[c2] = node{}
		}
		cave1 := caves[c1]
		cave2 := caves[c2]
		cave1.links = append(cave1.links, c2)
		cave2.links = append(cave2.links, c1)
		caves[c1] = cave1
		caves[c2] = cave2
	}
	numUniqPath := 0
	for i := range caves["start"].links {
		curPath := []string{"start", caves["start"].links[i]}
		search(caves, curPath, &numUniqPath)
	}
	log.Infof("%v", numUniqPath)
}

func search(caves map[string]node, curPath []string, numUniqPath *int) {
	curNode := curPath[len(curPath)-1]
	if curNode == "end" {
		log.Infof("found path: %v", curPath)
		(*numUniqPath)++
		return
	}
	links := caves[curNode].links
	for _, link := range links {
		// check if we already visited this cave and it is not large cave
		if hasVisited(curPath, link) && !unicode.IsUpper(rune(link[0])) {
			continue
		}
		newPath := append(curPath, link)
		search(caves, newPath, numUniqPath)
	}
}

func hasVisited(curPath []string, node string) bool {
	for _, c := range curPath {
		if c == node {
			return true
		}
	}
	return false
}

func Day12_2() {
	log.Info("day 12-2")
	f, _ := os.Open("day12/input.txt")

	scanner := bufio.NewScanner(f)
	caves := make(map[string]node)
	for scanner.Scan() {
		line := scanner.Text()
		sep := strings.Split(line, "-")
		c1 := sep[0]
		c2 := sep[1]
		if _, ok := caves[c1]; !ok {
			caves[c1] = node{}
		}
		if _, ok := caves[c2]; !ok {
			caves[c2] = node{}
		}
		cave1 := caves[c1]
		cave2 := caves[c2]
		cave1.links = append(cave1.links, c2)
		cave2.links = append(cave2.links, c1)
		caves[c1] = cave1
		caves[c2] = cave2
	}
	numUniqPath := 0
	for i := range caves["start"].links {
		curPath := []string{"start", caves["start"].links[i]}
		search2(caves, curPath, &numUniqPath)
	}
	log.Infof("%v", numUniqPath)
}

func search2(caves map[string]node, curPath []string, numUniqPath *int) {
	curNode := curPath[len(curPath)-1]
	if curNode == "end" {
		log.Infof("found path: %v", curPath)
		(*numUniqPath)++
		return
	}
	links := caves[curNode].links
	for _, link := range links {
		if link == "start" {
			continue
		}
		// check if we already visited this cave
		if hasVisited2(curPath, link) {
			continue
		}
		newPath := append(curPath, link)
		search2(caves, newPath, numUniqPath)
	}
}

func hasVisited2(curPath []string, node string) bool {
	// always allow to visit big cave as many times as necessary
	if unicode.IsUpper(rune(node[0])) {
		return false
	}
	// now look if we have visited small cave twice already
	visSmallCaveTwice := false
	for i := 1; i < len(curPath)-1; i++ {
		if unicode.IsUpper(rune(curPath[i][0])) {
			continue
		}
		for j := i + 1; j < len(curPath); j++ {
			if curPath[i] == curPath[j] {
				visSmallCaveTwice = true
				break
			}
		}
		if visSmallCaveTwice {
			break
		}
	}
	for _, c := range curPath {
		if c == node {
			return visSmallCaveTwice
		}
	}
	return false
}
