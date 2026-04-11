package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Kylo33/group-assign/match"
)

// Every problem should be solved by two people.
const coverage = 2

type groupMember struct {
	name     string
	problems []int
}

func main() {
	// TODO: Clean up input, consider Bubbletea/huh, pflags, and/or Cobra
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v <problem count> <group members>", os.Args[0])
	}

	problemCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid problem count: %v", err)
	}
	problems := make([]int, problemCount)
	for i := range problems {
		problems[i] = i + 1
	}

	names := strings.Split(os.Args[2], ",")
	slices.Sort(names)

	var members []*groupMember
	for _, name := range names {
		members = append(members, &groupMember{name: name})
	}

	matches, err := match.Fair(problems, members, coverage)
	if err != nil {
		log.Fatalf("Error matching group members: %v", err)
	}

	for _, m := range matches {
		for _, memberPtr := range m.To {
			memberPtr.problems = append(memberPtr.problems, m.From)
		}
	}

	for _, member := range members {
		fmt.Printf("%v: %v\n", member.name, member.problems)
	}
}
