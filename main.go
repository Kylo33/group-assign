package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Kylo33/group-assign/match"
)

var coverage = flag.Int("coverage", 2, "number of people per problem")
var people = flag.String("people", "", "comma separated list of group members")
var problems = flag.Int("problems", 0, "number of problems to assign")

func main() {
	flag.Parse()

	if *people == "" {
		fmt.Println("Error: the -people flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if *problems <= 0 {
		fmt.Println("Error: the -problems flag is required")
		flag.Usage()
		os.Exit(1)
	}

	names := strings.Split(*people, ",")
	for i := range names {
		names[i] = strings.TrimSpace(names[i])
	}
	slices.Sort(names)
	assignments := make(map[string][]int)

	problemList := make([]int, *problems)
	for i := range problemList {
		problemList[i] = i + 1
	}

	matches := match.Fair(problemList, names, *coverage)
	for _, match := range matches {
		for _, name := range match.To {
			assignments[name] = append(assignments[name], match.From)
		}
	}
	for _, name := range names {
		problems := assignments[name]
		fmt.Printf("%v: %v\n", name, problems)
	}
}
