package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type ErrorCause struct {
	ErrorMessage string
	ErrorType string
}

type ErrorEvent struct {
	Error string		`json:"error"`
	Cause string		`json:"cause"`
}
type Event struct {
	TargetYmd string			`json:"targetYmd"`
	RandNum int					`json:"randNum"`
	ErrorEvent *ErrorEvent		`json:"error"`
	EndErrorEvent *ErrorEvent	`json:"end_error"`
}

func handler(event Event) error {
	fmt.Println("NotifyFunction start")

	if event.ErrorEvent != nil {
		errorMessage, err := getErrorMessage(event.ErrorEvent)
		if err != nil {
			return err
		}
		fmt.Println(errorMessage)
	}
	if event.EndErrorEvent != nil {
		errorMessage, err := getErrorMessage(event.EndErrorEvent)
		if err != nil {
			return err
		}
		fmt.Println(errorMessage)
	}

	fmt.Println("NotifyFunction end")

	return nil
}

func main() {
	lambda.Start(handler)
}

func getErrorMessage(errorEvent *ErrorEvent) (string, error) {
	b := []byte(errorEvent.Cause)
	var e ErrorCause
	if err := json.Unmarshal(b, &e); err != nil {
		return "", err
	}
	return e.ErrorMessage, nil
}
