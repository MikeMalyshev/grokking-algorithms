package unit7

import (
	"fmt"
	"math"
	"strings"
)

type Direction struct {
	name     string
	distance float64
}

type LocationQueueItem struct {
	next     *LocationQueueItem
	Location string
	Weight   float64
}

type LocationQueue struct {
	first *LocationQueueItem
	last  *LocationQueueItem
}

func (queue *LocationQueue) Add(location string, weight float64) {
	newItem := &LocationQueueItem{next: nil, Location: location, Weight: weight}

	if queue.first == nil {
		queue.first = newItem
		queue.last = newItem
		return
	}

	if queue.first.Weight > weight {
		newItem.next = queue.first
		queue.first = newItem
		return
	}

	item := queue.first
	for item.next != nil {
		if item.next.Weight > weight {
			newItem.next = item.next
			item.next = newItem
			return
		}
		item = item.next
	}
	queue.last.next = newItem
	queue.last = newItem
}

func (queue *LocationQueue) Pop() *LocationQueueItem {
	defer func() {
		if queue.first != nil {
			queue.first = queue.first.next
		}
	}()

	return queue.first
}

func (queue *LocationQueue) IsEmpty() bool {
	return queue.first == nil
}

func addRoad(newRoad string, dist float64, directions *map[string][]Direction) {
	dep, dest, _ := parseWay(newRoad)
	(*directions)[dep] = append((*directions)[dep], Direction{name: dest, distance: dist})
	(*directions)[dest] = append((*directions)[dest], Direction{name: dep, distance: dist})
}

func parseWay(way string) (dep, dst string, ok bool) {
	way = strings.TrimSpace(way)
	list := strings.Split(way, "-")
	if len(list) != 2 {
		return "", "", false
	}
	return list[0], list[1], true
}

func FindBestWay(dep, dest string, roads map[string]float64) string {
	directions := make(map[string][]Direction)
	for r, l := range roads {
		addRoad(r, l, &directions)
	}

	queue := LocationQueue{}
	weights := make(map[string]float64)
	marked := make(map[string]bool)
	order := make(map[string]string)

	queue.Add(dep, 0)
	weights[dep] = 0

	for !queue.IsEmpty() {
		item := queue.Pop()
		for _, dir := range directions[item.Location] {
			if marked[dir.name] {
				continue
			}
			_, ok := weights[dir.name]
			if !ok {
				weights[dir.name] = math.MaxFloat64
			}

			newDistance := dir.distance + weights[item.Location]
			if weights[dir.name] > newDistance {
				weights[dir.name] = newDistance
				queue.Add(dir.name, newDistance)
				order[dir.name] = item.Location
			}
		}
		marked[item.Location] = true
	}

	fmt.Print("\nCalculated weights: ")
	fmt.Println(weights)

	var path []string

	loc := dest
	for loc != dep {
		path = append(path, loc)
		loc = order[loc]
	}
	path = append(path, loc)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	stringPath := ""
	for i, v := range path {
		if i == 0 {
			stringPath += v
			continue
		}
		stringPath += ("->" + v)
	}
	return stringPath
}
