package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

// some of the test is impossible according to the task, but I'll left them
func TestValidate(t *testing.T) {
	tests := map[string]struct {
		in    [][]int
		exout bool
	}{
		"success: simple exchange": {
			in: [][]int{
				{1, 2},
				{2, 1},
			},
			exout: true,
		},
		"success: no exchange": {
			in: [][]int{
				{0, 2},
				{2, 0},
			},
			exout: true,
		},
		"success: 3 6 6": {
			in: [][]int{
				{1, 2, 0},
				{1, 2, 3},
				{1, 2, 3},
			},
			exout: true,
		},
		"success: 1 1 1": {
			in: [][]int{
				{1},
				{0, 1},
				{0, 0, 1},
			},
			exout: true,
		},
		"fail: more containers with cap than colours": {
			in: [][]int{
				{1},
				{1},
				{1},
			},
			exout: false,
		},
		"fail: ": {
			in: [][]int{
				{1, 2},
				{2, 1},
				{0, 0},
			},
			exout: true,
		},
		"fail: multiple colours in one container": {
			in: [][]int{
				{1, 2, 1},
				{1, 2, 3},
				{1, 2, 3},
			},
			exout: false,
		},
		"fail: not enought conteiners": {
			in: [][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			exout: false,
		},
		"fail: not enought conteiners - single container has multiple colours": {
			in: [][]int{
				{1, 2, 3},
			},
			exout: false,
		},
		"fail: 2 containers with same colour": {
			in: [][]int{
				{0, 2},
				{2, 0},
				{0, 2},
			},
			exout: false,
		},
		"fail: impossible swap": {
			in: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			},
			exout: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := Validate(test.in)
			if test.exout != res {
				t.Errorf("expected: %v, got: %v", test.exout, res)
			}
		})
	}
}

func TestRead(t *testing.T) {
	tests := map[string]struct {
		in    string
		exout [][]int
	}{
		"example 1": {
			in: `2
			1 2
			2 1`,
			exout: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		"example 2": {
			in: `3
			10 20 30
			1 1 1
			0 0 1`,
			exout: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(test.in))
			out := Read(reader)

			if !slices.EqualFunc(test.exout, out, func(ex, out []int) bool {
				return slices.Equal(ex, out)
			}) {
				t.Errorf("expected: \n%v\n, got: \n%v\n", test.exout, out)
			}
		})
	}
}
