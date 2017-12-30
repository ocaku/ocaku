package data

import (
	
)
// Block interface
// Block is a universal block interface .
type Block interface {}

// WoodBlock .
type WoodBlock struct {
	Version   string          `json:"version"`   // version 1.0
	Previous  string          `json:"previous"`  // previous hash
	Index     int64     `json:"index"`  // block hight
	Random    int64           `json:"random"`    // random number
	Proof     int64           `json:"proof"`     // proof like len(nodes)
	Nodes     map[string]int8 `json:"nodes"`     // Now IP consensus
//	Merkle *Merkle // 
	Hash      string    `json:"hash"` // Hash signature
	Data  [][]byte        `json:"data"`   // save data infos 
	Timestamp int64           `json:"timestamp"` // time stamp
}


