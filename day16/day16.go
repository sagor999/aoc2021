package day16

import (
	"bufio"
	"math"
	"os"

	log "github.com/sirupsen/logrus"
)

type packet struct {
	version int
	type_id int
	result  int
}

var version_sum int = 0

func Day16_1() {
	log.Info("day 16-1")
	f, _ := os.Open("day16/input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	hex_input := scanner.Text()
	log.Infof("hex: %s", hex_input)
	// convert into binary stream
	hex_2_bin := make(map[string]string)
	hex_2_bin["0"] = "0000"
	hex_2_bin["1"] = "0001"
	hex_2_bin["2"] = "0010"
	hex_2_bin["3"] = "0011"
	hex_2_bin["4"] = "0100"
	hex_2_bin["5"] = "0101"
	hex_2_bin["6"] = "0110"
	hex_2_bin["7"] = "0111"
	hex_2_bin["8"] = "1000"
	hex_2_bin["9"] = "1001"
	hex_2_bin["A"] = "1010"
	hex_2_bin["B"] = "1011"
	hex_2_bin["C"] = "1100"
	hex_2_bin["D"] = "1101"
	hex_2_bin["E"] = "1110"
	hex_2_bin["F"] = "1111"
	bin_stream := ""
	for i := range hex_input {
		bin := hex_2_bin[string(hex_input[i])]
		bin_stream = bin_stream + bin
	}
	log.Infof("binary stream: %s", bin_stream)

	i := 0
	pkt, read_bits := extractPacket(bin_stream[i:])
	log.Infof("packet version: %d, type: %d, read_bits: %d", pkt.version, pkt.type_id, read_bits)

	//log.Infof("%d, %d", i, len(bin_stream))
	log.Infof("version sum: %d", version_sum)
	log.Infof("final packet: %d", pkt.result)
}

// input is binary stram, output is parsed packet and how many bits did we read
func extractPacket(binary_stream string) (packet, int) {
	version_bin_str := binary_stream[:3]
	typeid_bin_str := binary_stream[3:6]
	//log.Infof("%s, %s", version_bin_str, typeid_bin_str)
	version := convert_bin_str_to_number(version_bin_str)
	typeid := convert_bin_str_to_number(typeid_bin_str)
	log.Infof("version: %d, type: %d", version, typeid)
	version_sum += version

	result_packet := packet{version: version, type_id: typeid}

	last_read_bit := 6
	if typeid == 4 {
		// literal value packet
		literal_bin_str := ""
		for {
			bits := binary_stream[last_read_bit:(last_read_bit + 5)]
			last_read_bit += 5
			//log.Infof("%s", bits)
			literal_bin_str = literal_bin_str + bits[1:]
			if bits[0] == '0' {
				break
			}
		}
		literal := convert_bin_str_to_number(literal_bin_str)
		result_packet.result = literal
		//log.Infof("literal: %d", literal)
	} else {
		length_bit := string(binary_stream[6])
		last_read_bit = 7
		log.Infof("length_bit: %s", length_bit)
		subpackets := make([]packet, 0)
		if length_bit == "0" {
			len_str := binary_stream[last_read_bit:(last_read_bit + 15)]
			last_read_bit += 15
			subpacket_len := convert_bin_str_to_number(len_str)
			log.Infof("subpacket len: %d", subpacket_len)
			subpacket_read_len := 0
			for {
				subpacket, read_len := extractPacket(binary_stream[last_read_bit+subpacket_read_len:])
				log.Infof("subpacket: %v", subpacket)
				subpackets = append(subpackets, subpacket)
				subpacket_read_len += read_len
				if subpacket_read_len == subpacket_len {
					break
				}
			}
			last_read_bit += subpacket_read_len
		} else {
			len_str := binary_stream[last_read_bit:(last_read_bit + 11)]
			last_read_bit += 11
			num_packets := convert_bin_str_to_number(len_str)
			log.Infof("numpackets: %d", num_packets)
			num_pkts_read := 0
			subpacket_read_len := 0
			for {
				subpacket, read_len := extractPacket(binary_stream[last_read_bit+subpacket_read_len:])
				log.Infof("subpacket: %v", subpacket)
				subpackets = append(subpackets, subpacket)
				subpacket_read_len += read_len
				num_pkts_read++
				if num_pkts_read == num_packets {
					break
				}
			}
			last_read_bit += subpacket_read_len
		}
		res := 0
		switch typeid {
		case 0:
			for p := range subpackets {
				res += subpackets[p].result
			}
		case 1:
			res = 1
			for p := range subpackets {
				res *= subpackets[p].result
			}
		case 2:
			res = math.MaxInt32
			for p := range subpackets {
				if subpackets[p].result < res {
					res = subpackets[p].result
				}
			}
		case 3:
			res = math.MinInt32
			for p := range subpackets {
				if subpackets[p].result > res {
					res = subpackets[p].result
				}
			}
		case 5:
			if subpackets[0].result > subpackets[1].result {
				res = 1
			} else {
				res = 0
			}
		case 6:
			if subpackets[0].result < subpackets[1].result {
				res = 1
			} else {
				res = 0
			}
		case 7:
			if subpackets[0].result == subpackets[1].result {
				res = 1
			} else {
				res = 0
			}
		}
		result_packet.result = res
	}
	return result_packet, last_read_bit
}

func convert_bin_str_to_number(bin_data string) int {
	res := 0
	for i := range bin_data {
		n := 0
		if bin_data[i] == '1' {
			n = 1
		}
		res |= n << (len(bin_data) - i - 1)
	}
	return res
}
