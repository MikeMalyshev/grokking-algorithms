package unit6

type PersonItem struct {
	next *PersonItem
	name string
}

type PersonQueue struct {
	first *PersonItem
	last  *PersonItem
}

func (queue *PersonQueue) Add(person string) {
	newItem := &PersonItem{next: nil, name: person}

	if queue.first == nil {
		queue.first = newItem
		queue.last = newItem
		return
	}
	queue.last.next = newItem
	queue.last = newItem
}

func (queue *PersonQueue) Pop() string {
	if queue.first == nil {
		return ""
	}
	output := queue.first.name
	queue.first = queue.first.next
	return output
}

func (queue PersonQueue) String() string {
	item := queue.first
	if item == nil {
		return "PersonQueue (nil)"
	}

	output := "PersonQueue (" + item.name
	item = item.next
	for item != nil {
		output += ", " + item.name
		item = item.next
	}
	output += ")"
	return output
}

func findIterative(name string, graph map[string][]string) string {
	var queue PersonQueue
	queue.Add(name)

	var checkedPersonList = make(map[string]bool)
	for queue.first != nil {
		for _, person := range graph[queue.Pop()] {
			if checkedPersonList[person] {
				continue
			}
			if isMangoSeller(person) {
				return person
			}
			checkedPersonList[person] = true
			queue.Add(person)
		}
	}
	return "noone"
}
