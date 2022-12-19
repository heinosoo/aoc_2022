package main

import (
	"regexp"
	"strconv"

	. "github.com/heinosoo/aoc_2022"
)

type ores [4]int16
type robotT struct {
	cost  ores
	speed ores
}
type blueprintT [4]robotT

func part1(lines Channel[string]) {
	blueprints, limits := parseBlueprints(lines)
	quality := 0
	for i, blueprint := range blueprints {
		cache := make(map[[9]int16]ores)
		geodes := int(maxGeodes(24, ores{1, 0, 0, 0}, ores{0, 0, 0, 0}, limits, blueprint, cache)[3])
		quality += geodes * (i + 1)
		Log(i+1, quality)
	}
	Log(quality)
}

func cacheKey(timeLeft int16, S, R, limits ores, B blueprintT) (key [9]int16) {
	ore, clay, obs := Min(R[0], limits[0]), Min(R[1], limits[1]), Min(R[2], limits[2])
	return [9]int16{timeLeft, ore, clay, obs, R[3], S[0], S[1], S[2], S[3]}
}

func maxGeodes(timeLeft int16, S, R, limits ores, B blueprintT, cache map[[9]int16]ores) (newR ores) {
	key := cacheKey(timeLeft, S, R, limits, B)
	cacheVal, exists := cache[key]
	if exists {
		return cacheVal
	} else if timeLeft == 1 {
		return R.add(S)
	}

	branchS, branchR := buildOptions(S, R, B)
	for i := range branchS {
		newR2 := maxGeodes(timeLeft-1, branchS[i], branchR[i].add(S), limits, B, cache)
		if newR2[3] > newR[3] {
			newR = newR2
		}
	}

	cache[key] = newR
	return
}

func buildOptions(S, R ores, B blueprintT) (newS, newR []ores) {
	newS, newR = append(newS, S), append(newR, R)
	for _, robot := range B {
		switch {
		case R[0] < robot.cost[0] || R[1] < robot.cost[1] || R[2] < robot.cost[2] || R[3] < robot.cost[3]:
		default:
			newS = append(newS, S.add(robot.speed))
			newR = append(newR, R.sub(robot.cost))
		}
	}
	return
}

func (A ores) add(B ores) (res ores) {
	return ores{A[0] + B[0], A[1] + B[1], A[2] + B[2], A[3] + B[3]}
}
func (A ores) sub(B ores) (res ores) {
	return ores{A[0] - B[0], A[1] - B[1], A[2] - B[2], A[3] - B[3]}
}

func parseBlueprints(lines Channel[string]) (blueprints []blueprintT, limits ores) {
	maxOre, maxClay, maxObs := MaxF(int16(0)), MaxF(int16(0)), MaxF(int16(0))
	for line := range lines {
		r := regexp.MustCompile(`Blueprint \d+: Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`).FindStringSubmatch(line)
		oreOre, _ := strconv.Atoi(r[1])
		clayOre, _ := strconv.Atoi(r[2])
		obsOre, _ := strconv.Atoi(r[3])
		obsClay, _ := strconv.Atoi(r[4])
		geodeOre, _ := strconv.Atoi(r[5])
		geodeObs, _ := strconv.Atoi(r[6])
		oreR := robotT{ores{int16(oreOre), 0, 0, 0}, ores{1, 0, 0, 0}}
		clayR := robotT{ores{int16(clayOre), 0, 0, 0}, ores{0, 1, 0, 0}}
		obsR := robotT{ores{int16(obsOre), int16(obsClay), 0, 0}, ores{0, 0, 1, 0}}
		geodeR := robotT{ores{int16(geodeOre), 0, int16(geodeObs), 0}, ores{0, 0, 0, 1}}
		blueprints = append(blueprints, blueprintT{oreR, clayR, obsR, geodeR})
		maxOre(int16(oreOre), int16(clayOre), int16(obsOre), int16(geodeOre))
		maxClay(int16(obsClay))
		maxObs(int16(geodeObs))
	}
	// Not guaranteed to work. Risky again, but I ran out of memory. : p
	limits = ores{maxOre(0) * 2, maxClay(0) * 2, maxObs(0) * 2, 0}
	return
}
