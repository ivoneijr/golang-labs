package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qt int) {
	for i := 0; i < qt; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("%s: %s (i = %d)\n", person, text, i+1)
	}
}

func main() {
	// sync code
	// speak("Ivo", "Oi oi", 3)
	// speak("Anne", "Oi também", 3)

	//async
	// go speak("Ivo", "Oi oi", 5)
	// go speak("Ari", "ooshh", 5)
	// time.Sleep(time.Second * 5)

	// sync and async
	// go speak("Ivo", "Oi oi", 10)
	// speak("Ari", "ooshh", 5)
}
