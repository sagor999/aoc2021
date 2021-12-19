package day18

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

type snailPair struct {
	x       int
	y       int
	sx      *snailPair
	sy      *snailPair
	explode bool
}

func printSnailNum(s snailPair) string {
	res := ""
	if s.explode {
		res += "*"
	}
	res += "["
	if s.sx != nil {
		res += printSnailNum(*s.sx)
	} else {
		res += fmt.Sprint(s.x)
	}
	res += ","
	if s.sy != nil {
		res += printSnailNum(*s.sy)
	} else {
		res += fmt.Sprint(s.y)
	}
	res += "]"
	if s.explode {
		res += "*"
	}
	return res
}

func deepCopy(pair snailPair) snailPair {
	res := snailPair{}
	if pair.sx != nil {
		sx := deepCopy(*pair.sx)
		res.sx = &sx
	} else {
		res.x = pair.x
	}
	if pair.sy != nil {
		sy := deepCopy(*pair.sy)
		res.sy = &sy
	} else {
		res.y = pair.y
	}
	return res
}

func Day18_1() {
	log.Info("day 18-1")
	f, _ := os.Open("day18/input.txt")

	scanner := bufio.NewScanner(f)
	snailList := make([]snailPair, 0)
	for scanner.Scan() {
		text := scanner.Text()
		snail, _ := parseSnail(text[1:])
		//log.Infof("res: %s", printSnailNum(snail))
		snailList = append(snailList, snail)
	}

	for {
		if len(snailList) == 1 {
			break
		}
		// addition of pairs
		newPair := snailPair{}
		sx := snailList[0]
		sy := snailList[1]
		newPair.sx = &sx
		newPair.sy = &sy
		snailList = snailList[1:]
		snailList[0] = newPair

		//log.Infof("after addition: %s", printSnailNum(snailList[0]))

		for {
			did_explode := explode(&snailList[0])
			if did_explode {
				//log.Infof("after explode: %s", printSnailNum(snailList[0]))
				continue
			}

			did_split := split(&snailList[0])
			if did_split {
				//log.Infof("after split: %s", printSnailNum(snailList[0]))
				continue
			}

			break
		}
	}
	log.Infof("result: %s", printSnailNum(snailList[0]))
	log.Infof("magnitude: %d", calcMagnitude(snailList[0]))
}

func Day18_2() {
	log.Info("day 18-2")
	f, _ := os.Open("day18/input.txt")

	scanner := bufio.NewScanner(f)
	snailList := make([]snailPair, 0)
	for scanner.Scan() {
		text := scanner.Text()
		snail, _ := parseSnail(text[1:])
		snailList = append(snailList, snail)
	}

	maxMagnitude := math.MinInt32
	for i := 0; i < len(snailList); i++ {
		for j := i + 1; j < len(snailList); j++ {
			newPair := snailPair{}
			sx := deepCopy(snailList[i])
			sy := deepCopy(snailList[j])
			newPair.sx = &sx
			newPair.sy = &sy

			for {
				did_explode := explode(&newPair)
				if did_explode {
					continue
				}

				did_split := split(&newPair)
				if did_split {
					continue
				}

				break
			}

			magn := calcMagnitude(newPair)
			if magn > maxMagnitude {
				maxMagnitude = magn
			}

			// try reverse pair
			sx = deepCopy(snailList[j])
			sy = deepCopy(snailList[i])
			newPair.sx = &sx
			newPair.sy = &sy

			for {
				did_explode := explode(&newPair)
				if did_explode {
					continue
				}

				did_split := split(&newPair)
				if did_split {
					continue
				}

				break
			}

			magn = calcMagnitude(newPair)
			if magn > maxMagnitude {
				maxMagnitude = magn
			}
		}
	}

	log.Infof("result: %d", maxMagnitude)
}

func calcMagnitude(pair snailPair) int {
	left := 0
	if pair.sx != nil {
		left = 3 * calcMagnitude(*pair.sx)
	} else {
		left = 3 * pair.x
	}
	right := 0
	if pair.sy != nil {
		right = 2 * calcMagnitude(*pair.sy)
	} else {
		right = 2 * pair.y
	}
	return left + right
}

func split(pair *snailPair) bool {
	if pair.sx != nil {
		if split(pair.sx) {
			return true
		}
	} else if pair.x > 9 {
		x := int(math.Floor(float64(pair.x) / 2.0))
		y := int(math.Ceil(float64(pair.x) / 2.0))
		(*pair).sx = &snailPair{x: x, y: y}
		(*pair).x = 0
		return true
	}
	if pair.sy != nil {
		if split(pair.sy) {
			return true
		}
	} else if pair.y > 9 {
		x := int(math.Floor(float64(pair.y) / 2.0))
		y := int(math.Ceil(float64(pair.y) / 2.0))
		(*pair).sy = &snailPair{x: x, y: y}
		(*pair).y = 0
		return true
	}
	return false
}

func findExplode(pair *snailPair, nest int) (bool, int, int) {
	if nest >= 4 && pair.sx == nil && pair.sy == nil {
		(*pair).explode = true
		return true, pair.x, pair.y
	} else {
		did_explode := false
		var x, y int
		if pair.sx != nil {
			did_explode, x, y = findExplode(pair.sx, nest+1)
			if did_explode {
				return did_explode, x, y
			}
		}
		if pair.sy != nil {
			did_explode, x, y = findExplode(pair.sy, nest+1)
			if did_explode {
				return did_explode, x, y
			}
		}
	}
	return false, -1, -1
}

func explode(pair *snailPair) bool {
	did_explode, x, y := findExplode(pair, 0)
	if did_explode {
		snail := printSnailNum(*pair)
		first_star := strings.Index(snail, "*")
		// propagate X to the left
		for i := first_star; i >= 0; i-- {
			if !unicode.IsDigit(rune(snail[i])) {
				continue
			}
			first := i
			last := i
			num_str := string(snail[i])
			if unicode.IsDigit(rune(snail[i-1])) {
				num_str = string(snail[i-1]) + string(snail[i])
				first = i - 1
			}
			n, _ := strconv.Atoi(num_str)
			n = n + x
			snail = snail[:first] + fmt.Sprintf("%d", n) + snail[last+1:]
			break
		}
		last_star := strings.LastIndex(snail, "*")
		// propagate Y to the right
		for i := last_star; i < len(snail); i++ {
			if !unicode.IsDigit(rune(snail[i])) {
				continue
			}
			first := i
			last := i
			num_str := string(snail[i])
			if unicode.IsDigit(rune(snail[i+1])) {
				num_str = string(snail[i]) + string(snail[i+1])
				last = i + 1
			}
			n, _ := strconv.Atoi(num_str)
			n = n + y
			snail = snail[:first] + fmt.Sprintf("%d", n) + snail[last+1:]
			break
		}
		first_star = strings.Index(snail, "*")
		last_star = strings.LastIndex(snail, "*")
		if first_star == -1 || last_star == -1 {
			log.Fatalf("Failed to find * in string: %s", snail)
		}
		//replace exploded pair with 0
		snail = snail[:first_star] + "0" + snail[last_star+1:]

		(*pair), _ = parseSnail(snail[1:])
	}

	return did_explode
}

// always pass a string that skips first [
func parseSnail(text string) (snailPair, int) {
	//log.Infof("parseSnail: %s", text)
	res := snailPair{}
	isleft := true
	i := 0
	for ; i < len(text); i++ {
		if text[i] == '[' {
			s, j := parseSnail(text[i+1:])
			i += j + 1
			if isleft {
				res.sx = &s
			} else {
				res.sy = &s
			}
		} else if text[i] == ']' {
			break
		} else if text[i] == ',' {
			isleft = false
		} else {
			var n int
			var err error
			if unicode.IsDigit(rune(text[i+1])) {
				n, err = strconv.Atoi(string(text[i]) + string(text[i+1]))
				i = i + 1
			} else {
				n, err = strconv.Atoi(string(text[i]))
			}
			if err != nil {
				log.Fatalf("got something wrong. %s, %s, %v", text, string(text[i]), err)
			}
			if isleft {
				res.x = n
			} else {
				res.y = n
			}
		}
	}
	if isleft {
		log.Fatalf("returning unclosed pair: %s", text)
	}
	return res, i
}
