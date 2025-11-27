package main

import "fmt"

type EmailNotifier struct {
	to   string
	from string
}

func (e EmailNotifier) Send(message string) error {
	// mock
	fmt.Printf("Sending email from %s to %s\n", e.from, e.to)
	return nil
}
