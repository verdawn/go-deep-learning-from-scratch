package main

import (
	"fmt"

	"github.com/verdawn/go-deep-learning-from-scratch/pkg/gate"
)

func main() {
	data := [][]float64{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	}

	titles := []string{"AND", "OR", "NAND", "XOR"}

	for _, title := range titles {
		println(title)
		for _, xs := range data {
			var y int
			if title == "AND" {
				y = gate.And(xs[0], xs[1])
			} else if title == "OR" {
				y = gate.Or(xs[0], xs[1])
			} else if title == "NAND" {
				y = gate.Nand(xs[0], xs[1])
			} else if title == "XOR" {
				y = gate.Xor(xs[0], xs[1])
			}

			fmt.Printf("{%d %d} -> %d\n", int(xs[0]), int(xs[1]), y)
		}
	}
}
