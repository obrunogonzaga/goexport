package main

import (
	"fmt"
	"github.com/obrunogonzaga/fcutils/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Printf("EventDispatcher: %v\n", ed)
}
