package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

func describe(n Notifier) {
	fmt.Printf("Notifier type: %T\n", n)
}
