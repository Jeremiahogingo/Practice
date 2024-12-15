package main

import "fmt"

func main() {
	// Single-line comments
	//var messagesFromDoris string
	//var numMessages float64
	//var totalCost float64
	//var costPerMessage float64

	messagesFromDoris := []string{
		"You doing anything later?",
		"Did you get my last messages",
		"Please respond I'm lonely",
	}
	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	totalCost := costPerMessage * numMessages
	fmt.Printf("Doris spent %.2f on text message today \n", totalCost)
}
