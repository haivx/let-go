package util

import (
	"github.com/matoous/go-nanoid/v2"
)

/*
 *  Generate unique key
 */

func NewID(length ...int) (id string) {
	if len(length) == 0 {
		id, _ = gonanoid.New(8)
	} else {
		id, _ = gonanoid.New(length[0])
	}

	return
}
