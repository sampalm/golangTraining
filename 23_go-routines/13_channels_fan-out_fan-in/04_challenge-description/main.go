package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond)
}

// publisher pushes data into a channel
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}

/*
CHALLENGE #1
Is this fan out?
A: YES, fan out means you have multiple functions reading from the same channel until that channel is closed.

CHALLENGE #2
Is this fan in?
A: NO, fan in means you have a function  that can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel thats closes when all the inputs are closed.
*/
