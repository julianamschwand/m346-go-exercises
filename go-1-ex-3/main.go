package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "The dice shows: ", eyes)

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "The dice was rolled at: ", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO
	// write into files with "go run main.go > file.txt"
}
