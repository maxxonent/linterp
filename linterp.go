package linterp

import (
	"fmt"
	"slices"
	"sort"
)

type Float interface {
	~float32 | ~float64
}

func Func[T Float](X, Y []T) (func(T) T, error) {
	lenX, lenY := len(X), len(Y)
	if lenX != lenY {
		return nil, fmt.Errorf("different length of slices: len(x) is %v, but len(y) is %v", lenX, lenY)
	}
	if lenX < 2 {
		return nil, fmt.Errorf("data slices should contain at least 2 elements, but their length is %v", lenX)
	}
	if !slices.IsSorted(X) {
		return nil, fmt.Errorf("slice of x must be sorted")
	}

	A, B := make([]T, lenX-1), make([]T, lenY-1)

	for i := 0; i < lenX-1; i++ {
		y1 := Y[i]
		y2 := Y[i+1]
		x1 := X[i]
		x2 := X[i+1]
		A[i] = (y1 - y2) / (x1 - x2)
		fmt.Printf("Added B[%v] == %v\n", i, B[i])
	}

	result := func(x T) T {
		var a, b T
		if x <= X[0] {
			a, b = A[0], B[0]
		} else if x >= X[lenX-1] {
			a, b = A[len(A)-1], B[len(B)-1]
		} else {
			i := sort.Search(len(X), func(i int) bool {
				return X[i] > x
			})
			i--
			a, b = A[i-1], B[i-1]
		}
		return a*x + b
	}
	return result, nil
}
