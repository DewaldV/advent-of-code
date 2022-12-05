package main

import (
	"strings"
	"strconv"
	"fmt"
)

type Assignment struct {
	start int
	end   int
}

func (a *Assignment) contains(other Assignment) bool {
	return other.start >= a.start && other.end <= a.end
}

func newAssignment(elf string) (Assignment, error) {
	p := strings.Split(elf, "-")
	start, err := strconv.Atoi(p[0])
	if err != nil {
		return Assignment{}, fmt.Errorf("unable to create assignment, %w", err)
	}

	end, err := strconv.Atoi(p[1])
	if err != nil {
		return Assignment{}, fmt.Errorf("unable to create assignment, %w", err)
	}


	return Assignment{
		start: start,
		end:   end,
	}, nil
}

func fullyContains(e1, e2 string) bool {
	elf1, _ := newAssignment(e1)
	elf2, _ := newAssignment(e2)

	return elf1.contains(elf2) || elf2.contains(elf1)
}
