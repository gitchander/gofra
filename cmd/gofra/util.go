package main

import (
	"log"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func cropInt(x, min, max int) int {
	if min > max {
		min, max = max, min
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

func cropFloat64(x, min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
