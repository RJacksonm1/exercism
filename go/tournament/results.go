package tournament

type Results struct {
	team   string
	wins   int
	losses int
	draws  int
}

// Value of each outcome
const (
	WinPoints  int = 3
	DrawPoints int = 1
	LossPoints int = 0
)

// Points earned by these results
func (r *Results) Points() int {
	return (r.wins * WinPoints) + (r.draws * DrawPoints) + (r.losses * LossPoints)
}

// Matches played
func (r *Results) Matches() int {
	return r.wins + r.draws + r.losses
}
