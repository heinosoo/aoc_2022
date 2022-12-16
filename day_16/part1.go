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
	maxRelease := runTickers(30, 0, V)
	Log(maxRelease)
}

type cacheReadType struct {
	key     string
	outChan chan int
}
type cacheWriteType struct {
	key string
	val int
}

func createCache() (reader func(key string) int, writer func(key string, val int)) {
	cache := make(map[string]int)
	readChan, writeChan := make(chan cacheReadType, 10), make(chan cacheWriteType, 10)
	go func() {
		for {
			select {
			case w := <-writeChan:
				cache[w.key] = w.val
			case r := <-readChan:
				r.outChan <- cache[r.key]
				close(r.outChan)
			}
		}
	}()

	reader = func(key string) int {
		cacheRead := cacheReadType{key, make(chan int)}
		readChan <- cacheRead
		return <-cacheRead.outChan
	}
	writer = func(key string, val int) {
		cacheWrite := cacheWriteType{key, val}
		writeChan <- cacheWrite
	}

	return
}

func copyClosed(closed map[string]bool) map[string]bool {
	newClosed := make(map[string]bool)
	for k, v := range closed {
		if v {
			newClosed[k] = true
		}
	}
	return newClosed
}

func runTickers(time1, time2 int, V valves) int {
	reader, writer := createCache()
	var tick func(guyA, guyB string, closed map[string]bool, released, timeLeftA, timeLeftB int, maxRelease chan int)
	tick = func(guyA, guyB string, closed map[string]bool, released, timeLeftA, timeLeftB int, maxRelease chan int) {
		cacheKey := string(released) + string(timeLeftA) + string(timeLeftB) // Fingers crossed here. : p
		cacheVal := reader(cacheKey)
		if cacheVal != 0 {
			maxRelease <- cacheVal
			return
		}
		max := MaxF(0)
		for next, distance := range V[guyA].path {
			if !closed[next] {
				continue
			}
			time := distance + 1
			if timeLeftA > time {
				nextReleased := released + V[next].rate*(timeLeftA-time)
				max1, max2 := make(chan int), make(chan int)
				go func() {
					newClosed := copyClosed(closed)
					newClosed[next] = false
					tick(next, guyB, newClosed, nextReleased, timeLeftA-time, timeLeftB, max1)
				}()
				go func() {
					newClosed := copyClosed(closed)
					newClosed[next] = false
					tick(guyB, next, newClosed, nextReleased, timeLeftB, timeLeftA-time, max2)
				}()
				max(nextReleased, <-max1, <-max2)
			}
		}
		writer(cacheKey, max(0))
		maxRelease <- max(0)
	}

	initClosed := make(map[string]bool)
	for label, v := range V {
		if v.rate != 0 {
			initClosed[label] = true
		}
	}
	maxRelease := make(chan int)
	go tick("AA", "AA", initClosed, 0, time1, time2, maxRelease)
	return <-maxRelease
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
