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

	switch len(strings) {
	case 0:
		return FreqMap{}

	case 1:
		return Frequency(strings[0])
	}

	// Recursively divide-and-conquer
	c := make(chan FreqMap, len(strings))

	f := func(strings []string) {
		c <- ConcurrentFrequency(strings)
	}

	half := len(strings) / 2
	go f(strings[half:])
	go f(strings[:half])

	m := <-c
	for r, c := range <-c {
		m[r] += c
	}

	return m
}
