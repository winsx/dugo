package learn

import (
	"fmt"
	"sort"
)

func TryMap() {
	ages := make(map[string]int) // mapping from strings to ints
	fmt.Printf("ages: %v len:%d\n", ages, len(ages))

	ages2 := map[string]int {
		"alice": 41,
		"chalie": 34,
		"alimy": 29,
	}
	fmt.Printf("ages2: %v len:%d\n", ages2, len(ages2))

	ages3 := make(map[string]int)
	ages3["alimy"] = 29
	ages3["alice"] =23
	fmt.Printf("ages3: %v len:%d\n", ages3, len(ages3))
}

func TryMapDelete() {
	ages := map[string]int{
		"alice":  41,
		"alimy":  29,
		"chalie": 34,
		"tim": 12,
		"kid": 21,
		"cathy": 25,
	}
	fmt.Printf("ages:%v len:%d alimy:%d tom:%d\n", ages, len(ages), ages["alimy"], ages["tom"])

	delete(ages, "alice")
	fmt.Println(ages)

	delete(ages, "tom")
	fmt.Println(ages)

	ages["alimy"] += 1
	fmt.Println("alimy:" , ages["alimy"])

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	if age, ok := ages["kimy"]; !ok {
		fmt.Println("not exist name of kimy in ages map")
	} else {
		fmt.Printf("kimy:%d\n", age)
	}
}

func mapEqual(x, y map[string]int) bool {
	if (len(x) != len(y)) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

var graph = make(map[string]map[string]bool)
func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}