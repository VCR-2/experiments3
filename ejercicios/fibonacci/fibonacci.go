package fibonacci

import "errors"

func fibonacci(x, y, i uint64) (uint64, uint64, error) {
	n1, n0, n := y, x+y, uint64(2)
	for n < i {
		n1, n0, n = n0, n1+n0, n+1
		if n0 < n1 {
			return 0, 0, errors.New("Overflow")
		}
	}
	return n1, n0, nil
}
