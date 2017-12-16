package grains

import "errors"

const squares = 64

// Square returns how many grains are on the nTh square,
// where the first square has 1 and every square thereafter has double that.
// Could implement via math.Pow(2, n), but bitshifting is equivalent and much faster
func Square(n int) (uint64, error) {

	if n <= 0 || n > squares {
		return 0, errors.New("Square out of bounds")
	}

	return uint64(1 << uint(n-1)), nil
}

// Total returns the total amount of grains on a checkboard,
// where the first square has 1 and every square thereafter has double that.
func Total() uint64 {
	return uint64((1 << squares) - 1)
}
