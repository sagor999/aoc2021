package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Day4_1() {
	log.Info("day 4-1")
	f, _ := os.Open("day4/input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	numbers_str_arr := strings.Split(scanner.Text(), ",")
	//log.Infof("%v", numbers_str_arr)
	lotto_sequence := make([]int, 0)
	for _, v := range numbers_str_arr {
		i, _ := strconv.Atoi(v)
		lotto_sequence = append(lotto_sequence, i)
	}
	//log.Infof("%v", lotto_sequence)
	// skip empty line
	scanner.Scan()

	// lotto board number to an array of all sequences in that board (vert & hor)
	lotto_boards := make(map[int][][]int)
	board_idx := 0
	for scanner.Scan() {
		board := make([][]int, 5)
		for i := range board {
			board[i] = make([]int, 5)
		}

		line := scanner.Text()
		for i := 0; i < 5; i++ {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "  ", " ")
			sep := strings.Split(line, " ")
			for j, v := range sep {
				n, _ := strconv.Atoi(v)
				board[i][j] = n
			}
			scanner.Scan()
			line = scanner.Text()
		}
		//log.Infof("%v", board)
		sequences := make([][]int, 10)
		for i := range board {
			sequences[i] = board[i]
		}
		for i := range board {
			vert_seq := make([]int, 5)
			for j := 0; j < 5; j++ {
				vert_seq[j] = board[j][i]
			}
			sequences[i+5] = vert_seq
		}
		//log.Infof("%v", sequences)
		lotto_boards[board_idx] = sequences
		board_idx++
	}
	last_lotto_num := 0
	bingo_board := -1
	var cur_seq []int
	var hashed_seq map[int]bool
	for i := range lotto_sequence {
		cur_seq = lotto_sequence[:i+1]
		if len(cur_seq) < 5 {
			continue
		}
		hashed_seq = make(map[int]bool)
		for _, s := range cur_seq {
			hashed_seq[s] = true
		}
		bingo := false
		for board_idx, board_seqs := range lotto_boards {
			for s := range board_seqs {
				if hasAllNumbers(hashed_seq, board_seqs[s]) {
					log.Info("Bingo!")
					bingo = true
					bingo_board = board_idx
					break
				}
			}
			if bingo {
				break
			}
		}
		last_lotto_num = lotto_sequence[i]
		if bingo {
			break
		}
	}
	log.Infof("%v", last_lotto_num)
	log.Infof("%v", lotto_boards[bingo_board])
	sum := 0
	whole_board := make(map[int]bool)
	for _, seq := range lotto_boards[bingo_board] {
		for i := range seq {
			whole_board[seq[i]] = true
		}
	}
	for i, _ := range whole_board {
		_, ok := hashed_seq[i]
		if !ok {
			sum += i
		}

	}
	log.Infof("sum: %d, res: %d", sum, sum*last_lotto_num)
}

func hasAllNumbers(hash map[int]bool, seq []int) bool {
	res := true
	for i := range seq {
		_, ok := hash[seq[i]]
		if !ok {
			res = false
			break
		}
	}
	return res
}

func Day4_2() {
	log.Info("day 4-2")
	f, _ := os.Open("day4/input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	numbers_str_arr := strings.Split(scanner.Text(), ",")
	//log.Infof("%v", numbers_str_arr)
	lotto_sequence := make([]int, 0)
	for _, v := range numbers_str_arr {
		i, _ := strconv.Atoi(v)
		lotto_sequence = append(lotto_sequence, i)
	}
	//log.Infof("%v", lotto_sequence)
	// skip empty line
	scanner.Scan()

	// lotto board number to an array of all sequences in that board (vert & hor)
	lotto_boards := make(map[int][][]int)
	board_idx := 0
	for scanner.Scan() {
		board := make([][]int, 5)
		for i := range board {
			board[i] = make([]int, 5)
		}

		line := scanner.Text()
		for i := 0; i < 5; i++ {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "  ", " ")
			sep := strings.Split(line, " ")
			for j, v := range sep {
				n, _ := strconv.Atoi(v)
				board[i][j] = n
			}
			scanner.Scan()
			line = scanner.Text()
		}
		//log.Infof("%v", board)
		sequences := make([][]int, 10)
		for i := range board {
			sequences[i] = board[i]
		}
		for i := range board {
			vert_seq := make([]int, 5)
			for j := 0; j < 5; j++ {
				vert_seq[j] = board[j][i]
			}
			sequences[i+5] = vert_seq
		}
		//log.Infof("%v", sequences)
		lotto_boards[board_idx] = sequences
		board_idx++
	}
	last_lotto_num := 0
	var bingo_board [][]int
	var cur_seq []int
	var hashed_seq map[int]bool
	for i := range lotto_sequence {
		cur_seq = lotto_sequence[:i+1]
		if len(cur_seq) < 5 {
			continue
		}
		hash_seq := make(map[int]bool)
		for _, s := range cur_seq {
			hash_seq[s] = true
		}
		for board_idx, board_seqs := range lotto_boards {
			bingo := false
			for s := range board_seqs {
				if hasAllNumbers(hash_seq, board_seqs[s]) {
					log.Infof("Bingo %d!", board_idx)
					bingo = true
					bingo_board = lotto_boards[board_idx]
					last_lotto_num = lotto_sequence[i]
					hashed_seq = hash_seq

					// remove this board
					delete(lotto_boards, board_idx)
					break
				}
			}
			if bingo {
				continue
			}
		}
	}
	log.Infof("last lotto: %v", last_lotto_num)
	log.Infof("board: %v", bingo_board)
	sum := 0
	whole_board := make(map[int]bool)
	for _, seq := range bingo_board {
		for i := range seq {
			whole_board[seq[i]] = true
		}
	}
	for i, _ := range whole_board {
		_, ok := hashed_seq[i]
		if !ok {
			sum += i
		}

	}
	log.Infof("sum: %d, res: %d", sum, sum*last_lotto_num)
}
