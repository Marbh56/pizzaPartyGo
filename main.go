package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("How many people?")
	numPeople, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Couldn't get number of people.")
	}

	fmt.Println("How many pizzas do you have?")
	numPizzas, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Couldn't get number of pizzas")
	}
}
