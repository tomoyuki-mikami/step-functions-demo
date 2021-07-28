package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	TargetYmd string	`json:"targetYmd"`
	RandNum int			`json:"randNum"`
}

func handler(event Event) error {
	fmt.Println("StartFunction start")
	fmt.Println(event)
	fmt.Println("StartFunction end")

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(event.RandNum) == 0 {
		return errors.New("StartFunction Failed")
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
