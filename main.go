// This program demonstrates a Bubble Sort of an array of integers.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	data := []int{4, 2, 5, 6, 10, 9, 1, 8, 3, 7}
	intro()
	bubbleSort(data)
	fmt.Println("\n*** Bubble Sort Complete ***")
}

func intro() {
	clearTerm()
	fmt.Println("Bubble Sort Demonstration")
	fmt.Println("****James Hill, 2023*****")
	time.Sleep(3 * time.Second)
	clearTerm()
}

// bubbleSort() carries out the sort as described in the in-line comments.
func bubbleSort(arr []int) {
	var isDone = false
	for !isDone { // Open while loop.
		showGraph(arr)
		isDone = true
		var i = 0
		for i < len(arr)-1 { // Iterate over array "data".
			if arr[i] > arr[i+1] { // If the leading element is larger than the following element,
				arr[i], arr[i+1] = arr[i+1], arr[i] // Swap their positions.
				isDone = false                      // value remains false until end of procedure.
			}
			i++
		}
	}
}

// clearTerm() clears the terminal display between updates.
func clearTerm() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// showGraph() displays the current state of the sort to the console as a graph of * characters.
func showGraph(arr []int) {
	star := "*"
	clearTerm()
	fmt.Println()
	for _, val := range arr {
		fmt.Printf(strings.Repeat(star, val) + "\n")
	}
	time.Sleep(1 * time.Second)

}
