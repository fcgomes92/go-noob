package main

import (
	"errors"
	"fmt"
)

type Person struct {
	ID    int
	Name  string
	Phone string
}

func add(a, b int) int {
	return a + b
}

func processSomething(thing string) (string, error) {
	if thing == "Cool thing" {
		return "processed successfully", nil
	}
	return "", errors.New("something bad happened")
}

func (person *Person) SayIntroduction() {
	fmt.Println("Hello! My name is " + person.Name)
}

func main() {
	var pos *int
	myOddsSlice := []int{1, 3, 5, 7, 9}
	fmt.Println(myOddsSlice)
	myOddsSlice = append(myOddsSlice, 11, 13)
	fmt.Println(myOddsSlice)
	pos = &myOddsSlice[2]
	fmt.Println(&pos)
	fmt.Println(*pos)

	numbers := map[string]int{}
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers)

	for index, value := range myOddsSlice {
		message := fmt.Sprintf("index: %d - value: %d", index, value)
		fmt.Println(message)
	}

	for k, v := range numbers {
		message := fmt.Sprintf("index: %s - value: %d", k, v)
		fmt.Println(message)
	}
}
