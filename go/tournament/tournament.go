package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

// Tally the tournament's results so far
func Tally(in io.Reader, out io.Writer) error {
	// Read results
	csvReader := csv.NewReader(in)
	csvReader.Comma = ';'
	csvReader.Comment = '#'

	// Turn records into results
	resultsMap := make(map[string]Results)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		match, err := NewMatch(record)
		if err != nil {
			return err
		}

		homeResults, ok := resultsMap[match.home]
		if !ok {
			homeResults = Results{team: match.home}
		}

		awayResults, ok := resultsMap[match.away]
		if !ok {
			awayResults = Results{team: match.away}
		}

		switch match.result {
		case "win":
			homeResults.wins++
			awayResults.losses++

		case "draw":
			homeResults.draws++
			awayResults.draws++

		case "loss":
			homeResults.losses++
			awayResults.wins++
		}

		resultsMap[homeResults.team] = homeResults
		resultsMap[awayResults.team] = awayResults
	}

	// Sort by points
	results := make([]Results, 0, len(resultsMap))
	for _, result := range resultsMap {
		results = append(results, result)
	}
	sort.Slice(results, func(i, j int) bool {
		iPoints, jPoints := results[i].Points(), results[j].Points()

		// When points match, sort alphabetically ascending
		if iPoints == jPoints {
			return results[i].team < results[j].team
		}

		// Otherwise, descending by points
		return iPoints > jPoints
	})

	// Output as table
	fmt.Fprintf(out, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, result := range results {
		fmt.Fprintf(out, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			result.team,
			result.Matches(),
			result.wins,
			result.draws,
			result.losses,
			result.Points(),
		)
	}

	return nil
}
