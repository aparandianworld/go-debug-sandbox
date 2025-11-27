package main

import "fmt"

func main() {

	notifier := []Notifier{
		EmailNotifier{
			to:   "to@example.com",
			from: "from@example.com",
		},
		SMSNotifier{
			phoneNumber: "1234567890",
		},
		SlackNotifier{
			channel: "#general",
		},
	}

	for _, n := range notifier {
		describe(n)
	}

	message := "Hello! This is an important message."
	for _, n := range notifier {
		if err := n.Send(message); err != nil {
			fmt.Printf("Failed to send notification: %v\n", err)
		}
	}

	fmt.Println("All notifications sent.")
}
