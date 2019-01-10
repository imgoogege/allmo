package allmo

import (
	"flag"
	"sync"
)

var (
	varS    string
	sy      sync.Mutex
	RestMap map[string]int
)

type Result struct {
	Imports []string
}

func init() {
	flag.StringVar(&varS, "mo", "github.com/googege/allmo", "search all mo")
	flag.Parse()
}
