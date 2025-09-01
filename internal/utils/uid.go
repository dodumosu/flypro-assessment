package utils

import "github.com/nrednav/cuid2"

func NewID() string {
	return cuid2.Generate()
}
