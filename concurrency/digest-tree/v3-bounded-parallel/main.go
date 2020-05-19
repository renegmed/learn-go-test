package main

/*
Bounded Parallelism

walkFiles starts a goroutine to walk the directory tree at roolt and send
the path of each regular file on the string channel. It sends the result
of the walk on the error channel. If done is closed, alkFiles abandons its
work.

*/

import (
	"crypto/md5"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// result is the product of reading and sumchecking a file using MD5
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

/*
digester reads path names from paths slice and sends digests of the corresponding
files on c channel until either paths or done channels are closed
*/
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

/*
 MD5All reads all the files in the file tree rooted at root and returns a map
 from file path to the MD5 sum of the file's contents. If the directory walk
 fails or any read operation fails, MD5All returns an error. In that case,
 MD5All does not wait for inflight read operations to complete.
*/
func MD5All(root string) (map[string][md5.Size]byte, error) {
	// MD5All closes the done channel when it returns; it may do so before
	// receiving all the values from c and errc.
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)

	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result)

	var wg sync.WaitGroup
	const numDigesters = 20

	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}() // end of pipeline

	// collect data
	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	// Check whether the Walk failed.
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

/*

walkFiles starts a goroutine to walk the directory tree at roolt and send
the path of each regular file on the string channel. It sends the result
of the walk on the error channel. If done is closed, alkFiles abandons its
work.

*/
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		// Close the paths channed after Walk returns
		defer close(paths)

		// No selec needed for this send since errc is buffered
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err // channel paths will be closed
			}
			if !info.Mode().IsRegular() {
				return nil // channel paths will be closed ???? or just continue
			}
			select {
			case paths <- path: // send path string to channel
			case <-done: // sender signals to stop
				return errors.New("walk canceled")
			}
			return nil // channel paths will be closed
		})
	}()
	return paths, errc
}
