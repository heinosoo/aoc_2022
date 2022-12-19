package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	blueprints, limits := parseBlueprints(lines)
	answer := 1
	for i, blueprint := range blueprints[:Min(len(blueprints), 3)] {
		cache := make(map[[9]int16]ores)
		geodes := int(maxGeodes(32, ores{1, 0, 0, 0}, ores{0, 0, 0, 0}, limits, blueprint, cache)[3])
		answer *= geodes
		Log(i+1, geodes)
	}
	Log(answer)
}
