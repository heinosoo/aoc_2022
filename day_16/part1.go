package main

import (
	"regexp"
	"strconv"
	"strings"

	. "github.com/heinosoo/aoc_2022"
)

type valves map[string]valve
type valve struct {
	rate    int
	tunnels []string
	path    map[string]int
}

func part1(lines Channel[string]) {
	V := readValves(lines)
	open := make(map[string]bool)
	for label, v := range V {
		if v.rate == 0 {
			open[label] = true
		}
	}

	cache := make(map[string]int)
	maxRelease := tick("AA", "AA", open, V, 0, 30, 0, cache)
	Log(maxRelease)
}

func tick(guyA, guyB string, open map[string]bool, V valves, released, timeLeftA, timeLeftB int, cache map[string]int) int {
	cacheKey := string(released) + string(timeLeftA) + string(timeLeftB) // Fingers crossed here. : p
	if cache[cacheKey] != 0 {
		return cache[cacheKey]
	}
	max := MaxF(0)
	for next, distance := range V[guyA].path {
		if open[next] {
			continue
		}
		time := distance + 1
		if timeLeftA > time {
			nextReleased := released + V[next].rate*(timeLeftA-time)
			open[next] = true
			max(nextReleased, tick(next, guyB, open, V, nextReleased, timeLeftA-time, timeLeftB, cache))
			max(nextReleased, tick(guyB, next, open, V, nextReleased, timeLeftB, timeLeftA-time, cache))
			open[next] = false
		}
	}
	cache[cacheKey] = max(0)
	return max(0)
}

func readValves(lines Channel[string]) valves {
	V := make(valves)
	for line := range lines {
		r := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)`).FindStringSubmatch(line)
		label, tunnels := r[1], strings.Split(r[3], ", ")
		rate, _ := strconv.Atoi(r[2])
		V[label] = valve{rate, tunnels, make(map[string]int)}
	}
	for label := range V {
		calculatePaths(label, V)
	}
	return V
}

func calculatePaths(v1 string, V valves) {
	edge, visited := make(map[string]bool), make(map[string]bool)
	for _, v := range V[v1].tunnels {
		edge[v] = true
	}
	visited[v1] = true
	for d := 1; len(edge) != 0; d++ {
		newEdge := make(map[string]bool)
		for tunnel := range edge {
			if visited[tunnel] {
				continue
			} else {
				visited[tunnel] = true
			}
			V[tunnel].path[v1] = d
			for _, newTunnel := range V[tunnel].tunnels {
				if !edge[newTunnel] {
					newEdge[newTunnel] = true
				}
			}
		}
		edge = newEdge
	}
}
