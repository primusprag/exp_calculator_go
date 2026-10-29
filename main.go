package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/primusprag/exp_calculator_go/calculations"
)

type user struct {
	endTime   time.Time
	endLvl    uint8
	endExp    float32
	scanner   bufio.Scanner
	startTime time.Time
	startLvl  uint8
	startExp  float32
}

func main() {
	user := user{
		startTime: time.Now().UTC(),
		scanner:   *bufio.NewScanner(os.Stdin),
	}
	fmt.Println("Welcome to the exp calculator!")
	err := user.getStartLvlAndExp()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Waiting for end input.")
	err = user.getEndLvlAndExp()

	fmt.Printf("Fun little division: %f\n", user.startExp/float32(user.startLvl))
	os.Exit(0)
}
