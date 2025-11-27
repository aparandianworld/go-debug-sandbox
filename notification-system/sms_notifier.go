package main

import "fmt"

type SMSNotifier struct {
	phoneNumber string
}

func (s SMSNotifier) Send(message string) error {
	// mock
	fmt.Printf("Sending SMS to %s\n", s.phoneNumber)
	return nil
}
