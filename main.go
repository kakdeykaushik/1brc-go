package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("filename must be provided")
		os.Exit(1)
	}
	filename := args[1]

	// city: min/max/avg
	store := NewStore()

	var count int
	for line := range emitRows(filename) {
		count++
		if count%10_00_000 == 0 {
			fmt.Println(count, " rows done")
		}

		datapoints := strings.Split(line, ";")

		city := datapoints[0]
		temperature, err := strconv.ParseFloat(strings.TrimSpace(datapoints[1]), 64)
		if err != nil {
			log.Fatal(err)
		}
		store.Add(city, temperature)
	}

	// assertTest(store, filename)
}

func emitRows(filename string) <-chan string {
	out := make(chan string, 1_000)

	go func() {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		r := bufio.NewScanner(f)
		for r.Scan() {
			out <- r.Text()
		}

		close(out)
	}()

	return out
}

func assertTest(store *Store, filename string) {

	outfile := filename[:len(filename)-3] + "out"

	f, err := os.Open(outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	expected := strings.TrimRight(string(data), "\n")
	got := store.String()

	if strings.Compare(expected, got) == 0 {
		fmt.Println("correct")
	} else {
		fmt.Println("incorrect")
		fmt.Println("Expected: ")
		fmt.Println(expected)
		fmt.Println("Got: ")
		fmt.Println(got)
	}

}
