package data

// Block interface
// Block is a universal block interface .
type Block interface {
}

// WoodBlock .
type WoodBlock struct {
	Version   string          `json:"version"`   // version 1.0
	Timestamp int64           `json:"timestamp"` // time stamp
	Random    int64           `json:"random"`    // random number
	Proof     int64           `json:"proof"`     // proof like len(nodes)
	Nodes     map[string]int8 `json:"nodes"`     // Now IP consensus
	Previous  string          `json:"previous"`  // previous hash
	Records   [][]byte        `json:"Records"`   // save info
}
