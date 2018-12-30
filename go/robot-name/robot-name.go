package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

const nameLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var assignedNames map[string]bool

// Is there a better way of instantiating assignedNames?
func init() {
	assignedNames = make(map[string]bool)
}

// Robot beep boops
type Robot struct {
	name string
}

// Name gives the roboto a name
func (r *Robot) Name() (name string, err error) {
	if r.name != "" {
		return r.name, err
	}

	// Let's have 1000 attempts at trying find an unused name
	// If we're past this, we've used enough names that trying to find the remaining unique combinations is prohibitively expensive
	for i := 0; i < 1000; i++ {
		name := fmt.Sprintf("%c%c%3d", nameLetters[rand.Intn(len(nameLetters))], nameLetters[rand.Intn(len(nameLetters))], rand.Intn(1000))

		// Find an unused name
		if _, previouslyUsed := assignedNames[name]; !previouslyUsed {
			r.name = name
			assignedNames[name] = true

			return name, nil
		}
	}

	return "", errors.New("All names have (probably) been allocated")
}

// Reset clears the robot's name
func (r *Robot) Reset() {
	// If we were to allow name recycling we'd need to consume this bool in the r.Name() check
	assignedNames[r.name] = false

	r.name = ""
}
