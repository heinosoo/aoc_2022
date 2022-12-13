package main

import (
	"sort"

	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	var packets []packet
	for line := range lines {
		if line == "" {
			continue
		}
		lineChannel := NewChannelFromSlice([]rune(line), 3)
		packets = append(packets, readPacket(lineChannel).packets[0])
	}
	packets = append(packets, readPacket(NewChannelFromSlice([]rune("[[2]]"), 3)).packets[0])
	packets = append(packets, readPacket(NewChannelFromSlice([]rune("[[6]]"), 3)).packets[0])
	sort.Sort(byMagic(packets))

	decoder_key := 1
	for i, p := range packets {
		if len(p.packets) == 1 && len(p.packets[0].packets) == 1 {
			v := p.packets[0].packets[0].value
			if v == 2 || v == 6 {
				decoder_key *= i + 1
			}
		}
	}
	Log(decoder_key)
}

type byMagic []packet

func (packets byMagic) Len() int {
	return len(packets)
}
func (packets byMagic) Swap(i, j int) {
	packets[i], packets[j] = packets[j], packets[i]
}
func (packets byMagic) Less(i, j int) bool {
	less, _ := packets[i].compare(packets[j])
	return less
}
