package robotname

import "math/rand"
import "fmt"

const nameLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var assignedNames []string

// Robot beep boops
type Robot struct {
	name string
}

// Name gives the roboto a name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = GenerateUniqueName()
		assignedNames = append(assignedNames, r.name)
	}

	return r.name, nil
}

// Reset clears the robot's name
func (r *Robot) Reset() {

	// The collision test doesn't allow recycling names 😕
	// defer ReleaseName(r.name)
	r.name = ""
}

// GenerateUniqueName generates an unassigned name for the robot
func GenerateUniqueName() string {
	name := GenerateName()

	// If there's a collision, recurse!
	// This doesn't handle the scenario where all robot names have been assigned 😨
	for _, n := range assignedNames {
		if name == n {
			name = GenerateUniqueName()
			break
		}
	}

	return name
}

// GenerateName generates a name for the robot
func GenerateName() string {
	return fmt.Sprintf("%c%c%3d",
		nameLetters[rand.Intn(len(nameLetters))],
		nameLetters[rand.Intn(len(nameLetters))],
		rand.Intn(1000))
}

// ReleaseName releases the name!
func ReleaseName(name string) {
	var filteredNamed []string

	for _, n := range assignedNames {
		if n == name {
			continue
		}

		filteredNamed = append(filteredNamed, n)
	}

	assignedNames = filteredNamed
}
