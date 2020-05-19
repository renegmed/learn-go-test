package main

/*
	Calculate the MD5 sum of all files under the
	specified directory, then print the results sroted
	by path name.

*/
import (
	"reflect"
	"testing"
)

var wantPaths = []string{
	"../files/NOTES.txt",
	"../files/notes/awsblockchain-2020-02-27_12.17.24.mp4"}

func TestFilesChecksum(t *testing.T) {
	t.Run("read files from directory tree and must store path names", func(t *testing.T) {

		var root = "../files"
		m, err := MD5All(root)
		if err != nil {
			t.Fatalf(err.Error())
		}

		var got_paths []string

		for path := range m {
			got_paths = append(got_paths, path)
		}
		//fmt.Println(got_paths)
		if len(got_paths) != 2 {
			t.Errorf("Path slice should have a length of 5 not %d", len(got_paths))
		}

		if !reflect.DeepEqual(got_paths, wantPaths) {
			t.Errorf("Error: got \n%q want \n%q \n", got_paths, wantPaths)
		}

	})
}
