package occ

import (
	"bufio"
	"encoding/gob"
	core "github.com/ApesPlan/prefixtree-core"
	"os"
	"path/filepath"
)

// Load gob serialized dict from dir
// 从目录加载gob序列化的dict
func Load(dir string) (*Dict, error) {
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	trie := core.New()
	if err := trie.LoadFromFile(trieFile, "gob"); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(valueFile, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	in := bufio.NewReader(file)
	dataDecoder := gob.NewDecoder(in)
	var values [][]string
	if err = dataDecoder.Decode(&values); err != nil {
		return nil, err
	}
	return &Dict{Trie: trie, Values: values}, nil
}

// Save gob serialized dict to dir
// 保存Gob序列化字典到目录
func (d *Dict) Save(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	if err := d.Trie.SaveToFile(trieFile, "gob"); err != nil {
		return err
	}
	file, err := os.OpenFile(valueFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	out := bufio.NewWriter(file)
	defer out.Flush()
	dataEncoder := gob.NewEncoder(out)
	return dataEncoder.Encode(d.Values)
}
