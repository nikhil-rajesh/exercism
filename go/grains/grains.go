// Package grains solves Grains problem in exercism.io go-track
package grains

import "errors"

// Square returns number of grains on the input square number
func Square(inp int) (uint64, error) {
	if inp < 1 || inp > 64 {
		return 0, errors.New("Square number out of bounds")
	}
	return 1 << (inp - 1), nil
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return 1<<64 - 1
}
