package main

import (
	"reflect"
	"sort"
	"testing"
)

// var wantPaths = []string{
// 	"../files/NOTES.txt",
// 	"../files/notes/awsblockchain-2020-02-27_12.17.24.mp4"}

var wantPaths = []string{
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/v1/main.go",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/v2-parallel-digestion/main.go",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/v2-parallel-digestion/main_test.go",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/files/notes/awsblockchain-2020-02-27_12.17.24.mp4",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/Makefile",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/files/NOTES.txt",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/v1/main_test.go",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/v3-bounded-parallel/main.go",
	"/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree/v3-bounded-parallel/main_test.go",
}

func TestFilesChecksum(t *testing.T) {
	t.Run("read files from directory tree and must store path names", func(t *testing.T) {

		//var root = "../files"
		var root = "/home/rmedal/programming/go-workspace/src/learn-go-test-exercises/concurrency/digest-tree"
		m, err := MD5All(root)
		if err != nil {
			t.Fatalf(err.Error())
		}

		var got_paths []string

		for path := range m {
			got_paths = append(got_paths, path)
		}
		sort.Strings(wantPaths)
		sort.Strings(got_paths)
		//fmt.Println(got_paths)
		// if len(got_paths) != 9 {
		// 	t.Errorf("Path slice should have a length of 9 not %d", len(got_paths))
		// }

		if !reflect.DeepEqual(got_paths, wantPaths) {
			t.Errorf("Error: got \n%q want \n%q \n", got_paths, wantPaths)
		}

	})
}
