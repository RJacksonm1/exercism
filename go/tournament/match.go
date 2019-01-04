package tournament

import "errors"

// Match within a tournament
type Match struct {
	home   string
	away   string
	result string
}

// NewMatch creates a Match from a record
func NewMatch(record []string) (Match, error) {
	if recordLength := len(record); recordLength != 3 {
		if recordLength < 3 {
			return Match{}, errors.New("Invalid record: Row too short")
		}
		return Match{}, errors.New("Invalid record: Row too long")
	}

	home, away, result := record[0], record[1], record[2]

	if result != "win" && result != "draw" && result != "loss" {
		return Match{}, errors.New("Invalid result")
	}

	return Match{home, away, result}, nil
}
