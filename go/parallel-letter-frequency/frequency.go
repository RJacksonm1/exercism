package letter

type FreqMap map[rune]int

// Frequency counts a-b-c's
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts a-b-c's faster thanks to the power of Modern Computing
func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap, len(strings))

	// Maps aren't concurrent safe, so we won't directly interact with `m`
	// Instead, generate a map via Frequency for each string
	for _, s := range strings {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	// And then reduce to a single map
	for range strings {
		for r, i := range <-c {
			m[r] += i
		}
	}

	return m
}
