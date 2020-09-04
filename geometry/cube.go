package geometry

import (
	"errors"
)

//CubeVolumne func
func CubeVolumne(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	}
	return 0, errors.New("Zero length edge is not allowed")

}
