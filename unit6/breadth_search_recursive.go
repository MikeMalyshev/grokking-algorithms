package unit6

func isMangoSeller(name string) bool {
	if name[len(name)-1] == 'm' {
		return true
	}
	return false
}
func findRecursive(name string, graph map[string][]string) string {
	var checkedPersonList = make(map[string]bool)
	p, _ := rfind(name, graph, checkedPersonList)
	return p
}

func rfind(name string, graph map[string][]string, checkedPersonList map[string]bool) (string, bool) {
	for _, person := range graph[name] {
		if checkedPersonList[person] {
			continue
		}
		if isMangoSeller(person) {
			return person, true
		}
		checkedPersonList[person] = true
	}

	for _, person := range graph[name] {
		name, ok := rfind(person, graph, checkedPersonList)
		if ok {
			return name, ok
		}
	}
	return "noone", false
}
