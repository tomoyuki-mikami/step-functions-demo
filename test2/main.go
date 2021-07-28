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
	TargetName string	`json:"targetName"`
}

func handler(event Event) error {
	fmt.Println("Test2Function start")
	fmt.Println(event)
	fmt.Println("Test2Function end")

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(event.RandNum) == 0 {
		return errors.New(fmt.Sprintf("%s Failed", event.TargetName))
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
