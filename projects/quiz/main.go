package main

import (
	"fmt"
	"quiz/quiz"
)

func main() {
	file := quiz.ConfigFlag()
	fmt.Println(file)
}
