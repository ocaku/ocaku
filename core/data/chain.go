package data

// Chain interface
// Chain is a universal chain interface.
type Chain interface {
}

// Container .
type Container struct {
	Height   int64  `json:"height"`
	Previous string `json:"previous"` // previous hash
	Block    []*Block
	Nodes    map[string]int8 `json:"nodes"` // Now IP consensus
}
