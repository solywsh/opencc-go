package occ

import (
	core "github.com/ApesPlan/prefixtree-core"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *core.PrefixTree
	Values [][]string
}
