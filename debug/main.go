package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Person struct {
	name string
	id   int
	age  int
}

var logger *slog.Logger

func init() {
	level := slog.LevelInfo // default level
	if envLevel := os.Getenv("LOG_LEVEL"); envLevel != "" {
		if parsedLevel, err := strconv.Atoi(envLevel); err == nil {
			level = slog.Level(parsedLevel)
		}
	}
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	logger = slog.New(handler)
}

func findPerson(person []*Person, name string) *Person {

	if len(person) == 0 {
		return nil
	}

	for _, p := range person {
		if p != nil && p.name == name {
			return p
		}
	}

	return nil
}

func populate(people []*Person) []*Person {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	names := []string{"John", "Jane", "Bob", "Alice", "Tom", "Sally", "Mike", "Sue", "Bill", "Lily", "Aaron"}
	for i := range people {
		nameIndex := i % len(names)
		people[i] = &Person{
			name: names[nameIndex],
			id:   r.Intn(999),
			age:  r.Intn(100),
		}
	}
	return people
}

func main() {
	people := make([]*Person, 25)
	people = populate(people)

	fmt.Printf("People:\n")
	for _, p := range people {
		// fmt.Printf("%v\n", *p)
		// log.Printf("[INFO] %v\n", *p)
		logger.Info("Person", "person", *p)
	}
	fmt.Printf("\n")

	p := findPerson(people, "Sue")
	if p != nil {
		// fmt.Printf("Found: %v\n", *p)
		logger.Info("Found", "person", *p)
	} else {
		// fmt.Printf("Not found: %+v\n", *p)
		logger.Error("Not found", "person", *p)
	}
}
