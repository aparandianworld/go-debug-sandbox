package main

import "fmt"

type SlackNotifier struct {
	channel string
}

func (s SlackNotifier) Send(message string) error {
	// mock
	fmt.Printf("Sending Slack message to %s\n", s.channel)
	return nil
}
