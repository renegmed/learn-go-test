package main

/*
Digesting Tree

MD5All reads all the files in the file tree rooted at root and returns a
map from file path to the MD5 sum of the file's contents. if the directory
walk fails or any read operation fails, MD5All returns an error.

*/

import (
	"crypto/md5"
	"io/ioutil"
	"os"
	"path/filepath"
)

func MD5All(root string) (map[string][md5.Size]byte, error) {

	m := make(map[string][md5.Size]byte)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		m[path] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}
