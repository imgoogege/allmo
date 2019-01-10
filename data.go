package allmo

import (
	"sync"
)

var (
	sy      sync.Mutex
	RestMap map[string]int
)

type Result struct {
	Imports []string
}

