package main

import (
	"testing"
)

func TestCreateAssignment(t *testing.T) {
	assignment, err := newAssignment("1-2")
	if err != nil {
		t.Errorf("got error creating assignment, %s", err)
	}

	if assignment.start != 1 {
		t.Errorf("assignment had incorrect start value, expected: 1, got: %d", assignment.start)
	}

	if assignment.end != 2 {
		t.Errorf("assignment had incorrect end value, expected: 2, got: %d", assignment.start)
	}
}

func TestNonOverlappingAssignments(t *testing.T) {
	elf1 := "2-4"
	elf2 := "6-8"

	contains := fullyContains(elf1, elf2)
	if contains {
		t.Errorf("should not have fully contained each other\n")
	}
}

func TestPartialOverlappingAssignments(t *testing.T) {
	elf1 := "2-4"
	elf2 := "3-5"

	contains := fullyContains(elf1, elf2)
	if contains {
		t.Errorf("should not have fully contained each other\n")
	}
}
