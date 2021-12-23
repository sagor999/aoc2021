package day20

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
)

func Day20_1() {
	log.Info("day 20-1")
	f, _ := os.Open("day20/input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	img_alg := scanner.Text()
	scanner.Scan() // skip empty line
	img := make([]string, 0)
	for scanner.Scan() {
		img = append(img, scanner.Text())
	}

	last_bit := 511

	//show("orig", img)
	for i := 0; i < 50; i++ {
		bit := 0
		if i%2 == 0 {
			bit = last_bit
		}
		//log.Infof("expanding with bit: %d", bit)
		img = expand(img, string(img_alg[bit]), 2)
		//show("expand1", img)
		img = transform(img, img_alg)
		//show("transform1", img)
	}

	//show("final", img)

	log.Infof("res: %d", calc_lit_pixels(img))
}

func calc_lit_pixels(img []string) int {
	res := 0
	for y := range img {
		for x := range img[0] {
			if img[y][x] == '#' {
				res += 1
			}
		}
	}
	return res
}

func transform(img []string, img_alg string) []string {
	res := make([]string, len(img))
	copy(res, img)
	// take care of infinite borders first
	tmp_enc := ""
	for i := 0; i < 9; i++ {
		tmp_enc += string(img[0][0])
	}
	tmp_ind := conv_to_dec(tmp_enc)
	tmp := []byte(res[0])
	for x := range res[0] {
		tmp[x] = img_alg[tmp_ind]
	}
	// convert first and last row
	res[0] = string(tmp)
	res[len(res)-1] = string(tmp)

	// convert first and last column
	for y := range res {
		tmp = []byte(res[y])
		tmp[0] = img_alg[tmp_ind]
		tmp[len(tmp)-1] = img_alg[tmp_ind]
		res[y] = string(tmp)
	}

	for y := 1; y < len(img)-1; y++ {
		tmp := []byte(res[y])
		for x := 1; x < len(img[0])-1; x++ {
			enc := ""
			enc += string(img[y-1][x-1]) + string(img[y-1][x]) + string(img[y-1][x+1])
			enc += string(img[y][x-1]) + string(img[y][x]) + string(img[y][x+1])
			enc += string(img[y+1][x-1]) + string(img[y+1][x]) + string(img[y+1][x+1])
			dec := conv_to_dec(enc)
			//log.Infof("%d:%d:%s:%d:%s", x, y, enc, dec, string(img_alg[dec]))
			tmp[x] = img_alg[dec]
		}
		res[y] = string(tmp)
	}
	return res
}

func conv_to_dec(enc string) int {
	dec := 0
	for e := len(enc) - 1; e >= 0; e-- {
		if enc[e] == '#' {
			dec |= 1 << (len(enc) - 1 - e)
		}
	}
	return dec

}

func show(text string, img []string) {
	log.Info(text)
	for i := range img {
		log.Infof("%s", img[i])
	}
}

func expand(img []string, ch string, expand int) []string {
	res := make([]string, len(img)+expand*2)
	empty_str := ""
	for i := 0; i < len(img[0]); i++ {
		empty_str += ch
	}
	empty_str += ch + ch + ch + ch
	for i := 0; i < expand; i++ {
		res[i] = empty_str
	}
	for i := range img {
		res[i+expand] = ch + ch + img[i] + ch + ch
	}
	for i := expand; i < expand*2; i++ {
		res[len(img)+i] = empty_str
	}

	return res
}
